package metricsdata

import (
	"fmt"
	"math"

	"github.com/lindb/roaring"

	"github.com/lindb/lindb/aggregation"
	"github.com/lindb/lindb/flow"
	"github.com/lindb/lindb/pkg/encoding"
	"github.com/lindb/lindb/pkg/stream"
	"github.com/lindb/lindb/series"
	"github.com/lindb/lindb/series/field"
)

//go:generate mockgen -source ./reader.go -destination=./reader_mock.go -package metricsdata

// for testing
var (
	getOffsetFunc = getOffset
)

const (
	dataFooterSize = 2 + // start time slot
		2 + // end time slot
		4 + // field metas position
		4 + // series ids position
		4 + // high offsets position
		4 // crc32 checksum
)

// Reader represents the metric block reader
type Reader interface {
	// Path returns file path
	Path() string
	// GetSeriesIDs returns the series ids in this sst file
	GetSeriesIDs() *roaring.Bitmap
	// GetFields returns the field metas in this sst file
	GetFields() field.Metas
	// GetTimeRange returns the time range in this sst file
	GetTimeRange() (start, end uint16)
	// Load loads the data from sst file, then returns the file metric scanner.
	Load(queryFlow flow.StorageQueryFlow, familyTime int64, fieldIDs []field.ID, highKey uint16, seriesID roaring.Container) flow.Scanner
	// readSeriesData reads series data from file by given position.
	readSeriesData(position int, tsd *encoding.TSDDecoder, fieldAggs []*fieldAggregator)
}

// fieldAggregator represents the field aggregator that does file data scan and aggregates
type fieldAggregator struct {
	fieldMeta field.Meta
	block     series.Block
}

// newFieldAggregator creates a field aggregator
func newFieldAggregator(fieldMeta field.Meta, block series.Block) *fieldAggregator {
	return &fieldAggregator{
		fieldMeta: fieldMeta,
		block:     block,
	}
}

// reader implements Reader interface that reads metric block
type reader struct {
	path          string
	buf           []byte
	highOffsets   *encoding.FixedOffsetDecoder
	seriesIDs     *roaring.Bitmap
	fields        field.Metas
	crc32CheckSum uint32
	start, end    uint16

	readFieldIndexes []int // read field indexes be used when query metric data
}

// NewReader creates a metric block reader
func NewReader(path string, buf []byte) (Reader, error) {
	r := &reader{
		path: path,
		buf:  buf,
	}
	if err := r.initReader(); err != nil {
		return nil, err
	}
	return r, nil
}

// Path returns the file path
func (r *reader) Path() string {
	return r.path
}

// GetSeriesIDs returns the series ids in this sst file
func (r *reader) GetSeriesIDs() *roaring.Bitmap {
	return r.seriesIDs
}

// GetFields returns the field metas in this sst file
func (r *reader) GetFields() field.Metas {
	return r.fields
}

// GetTimeRange returns the time range in this sst file
func (r *reader) GetTimeRange() (start, end uint16) {
	return r.start, r.end
}

// prepare prepares the field aggregator based on query condition
func (r *reader) prepare(familyTime int64, fieldIDs []field.ID, aggregator aggregation.ContainerAggregator) (aggs []*fieldAggregator) {
	fieldMap := make(map[field.ID]int)
	for idx, fieldMeta := range r.fields {
		fieldMap[fieldMeta.ID] = idx
	}
	for idx, fieldID := range fieldIDs { // sort by field ids
		fieldIdx, ok := fieldMap[fieldID]
		if !ok {
			continue
		}
		block, ok := aggregator.GetFieldAggregates()[idx].GetAggregateBlock(familyTime)
		if !ok {
			continue
		}
		r.readFieldIndexes = append(r.readFieldIndexes, fieldIdx)
		aggs = append(aggs, newFieldAggregator(r.fields[fieldIdx], block))
	}
	return
}

// Load loads the data from sst file, then returns the file metric scanner.
func (r *reader) Load(queryFlow flow.StorageQueryFlow,
	familyTime int64, fieldIDs []field.ID,
	highKey uint16, seriesID roaring.Container,
) flow.Scanner {
	// 1. get high container index by the high key of series ID
	highContainerIdx := r.seriesIDs.GetContainerIndex(highKey)
	if highContainerIdx < 0 {
		// if high container index < 0(series IDs not exist) return it
		return nil
	}
	// 2. get low container include all low keys by the high container index, delete op will clean empty low container
	lowContainer := r.seriesIDs.GetContainerAtIndex(highContainerIdx)
	foundSeriesIDs := lowContainer.And(seriesID)
	if foundSeriesIDs.GetCardinality() == 0 {
		return nil
	}
	offset, _ := r.highOffsets.Get(highContainerIdx)
	seriesOffsets := encoding.NewFixedOffsetDecoder(r.buf[offset:])

	aggregator := queryFlow.GetAggregator(highKey)
	fieldAggs := r.prepare(familyTime, fieldIDs, aggregator)
	if len(fieldAggs) == 0 {
		return nil
	}
	// must use lowContainer from store, because get series index based on container
	return newMetricScanner(r, fieldAggs, lowContainer, seriesOffsets)
}

// readSeriesData reads series data from file by given position.
func (r *reader) readSeriesData(position int, tsd *encoding.TSDDecoder, fieldAggs []*fieldAggregator) {
	fieldCount := r.fields.Len()
	if fieldCount == 1 {
		// metric has one field, just read the data
		tsd.ResetWithTimeRange(r.buf[position:], r.start, r.end)
		// read field data
		r.readField(fieldAggs[0].block, tsd)
		return
	}
	// read data for mutli-fields
	seriesData := r.buf[position:]
	fieldOffsets := encoding.NewFixedOffsetDecoder(seriesData)
	fieldsData := seriesData[fieldOffsets.Header()+fieldCount*fieldOffsets.ValueWidth():]
	for i, idx := range r.readFieldIndexes {
		offset, ok := fieldOffsets.Get(idx)
		if ok {
			tsd.ResetWithTimeRange(fieldsData[offset:], r.start, r.end)
			// read field data
			r.readField(fieldAggs[i].block, tsd)
		}
	}
}

// readField reads field data and aggregates it
func (r *reader) readField(block series.Block, tsd *encoding.TSDDecoder) {
	for tsd.Next() {
		if tsd.HasValue() {
			timeSlot := tsd.Slot()
			val := tsd.Value()

			if block.Append(int(timeSlot), math.Float64frombits(val)) {
				return
			}
		}
	}
}

// initReader initializes the reader context includes tag value ids/high offsets
func (r *reader) initReader() error {
	if len(r.buf) <= dataFooterSize {
		return fmt.Errorf("block length not ok")
	}
	// read footer(2+2+4+4+4+4)
	footerPos := len(r.buf) - dataFooterSize
	r.start = stream.ReadUint16(r.buf, footerPos)
	r.end = stream.ReadUint16(r.buf, footerPos+2)

	fieldMetaStartPos := int(stream.ReadUint32(r.buf, footerPos+4))
	seriesIDsStartPos := int(stream.ReadUint32(r.buf, footerPos+8))
	highOffsetsPos := int(stream.ReadUint32(r.buf, footerPos+12))
	r.crc32CheckSum = stream.ReadUint32(r.buf, footerPos+16)
	// validate offsets
	if fieldMetaStartPos > footerPos || seriesIDsStartPos > highOffsetsPos {
		return fmt.Errorf("bad offsets")
	}

	// read field metas
	offset := fieldMetaStartPos
	fieldCount := r.buf[offset]
	offset++
	r.fields = make(field.Metas, fieldCount)
	for i := byte(0); i < fieldCount; i++ {
		r.fields[i] = field.Meta{
			ID:   field.ID(r.buf[offset]),
			Type: field.Type(r.buf[offset+1]),
		}
		offset += 2
	}

	// read series ids
	seriesIDs := roaring.New()
	if err := encoding.BitmapUnmarshal(seriesIDs, r.buf[seriesIDsStartPos:]); err != nil {
		return err
	}
	r.seriesIDs = seriesIDs
	// read high offsets
	r.highOffsets = encoding.NewFixedOffsetDecoder(r.buf[highOffsetsPos:])
	return nil
}

// fieldIndexes returns field indexes of metric level
func (r *reader) fieldIndexes() map[field.ID]int {
	result := make(map[field.ID]int)
	for idx, f := range r.fields {
		result[f.ID] = idx
	}
	return result
}

// dataScanner represents the metric data scanner which scans the series data when merge operation
type dataScanner struct {
	reader        *reader
	container     roaring.Container
	seriesOffsets *encoding.FixedOffsetDecoder

	highKeys  []uint16
	highKey   uint16
	seriesPos int
}

// newDataScanner creates a data scanner for data merge
func newDataScanner(r Reader) *dataScanner {
	reader := r.(*reader)
	s := &dataScanner{
		reader:   reader,
		highKeys: reader.seriesIDs.GetHighKeys(),
	}
	s.nextContainer()
	return s
}

// fieldIndexes returns field indexes of metric level
func (s *dataScanner) fieldIndexes() map[field.ID]int {
	return s.reader.fieldIndexes()
}

// nextContainer goes next container context for scanner
func (s *dataScanner) nextContainer() {
	s.highKey = s.highKeys[s.seriesPos]
	s.container = s.reader.seriesIDs.GetContainerAtIndex(s.seriesPos)
	offset, _ := s.reader.highOffsets.Get(s.seriesPos)
	s.seriesOffsets = encoding.NewFixedOffsetDecoder(s.reader.buf[offset:])
	s.seriesPos++
}

// slotRange returns the slot range of metric level in current sst file
func (s *dataScanner) slotRange() (start, end uint16) {
	return s.reader.GetTimeRange()
}

// scan scans the data and returns series position if series id exist, else returns -1
func (s *dataScanner) scan(highKey, lowSeriesID uint16) int {
	if s.highKey < highKey {
		if s.seriesPos >= len(s.highKeys) {
			// current tag inverted no data can read
			return -1
		}
		s.nextContainer()
	}
	if highKey != s.highKey {
		// high key not match, return it
		return -1
	}
	// find data by low series id
	if s.container.Contains(lowSeriesID) {
		// get the index of low series id in container
		idx := s.container.Rank(lowSeriesID)
		// get series data data position
		offset, ok := getOffsetFunc(s.seriesOffsets, idx-1)
		if !ok {
			return -1
		}
		return offset
	}
	return -1
}

// getOffset returns the offset by idx
func getOffset(seriesOffsets *encoding.FixedOffsetDecoder, idx int) (int, bool) {
	return seriesOffsets.Get(idx)
}
