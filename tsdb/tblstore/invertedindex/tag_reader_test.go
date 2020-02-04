package invertedindex

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/lindb/roaring"
	"github.com/stretchr/testify/assert"

	"github.com/lindb/lindb/kv"
	"github.com/lindb/lindb/kv/table"
	"github.com/lindb/lindb/sql/stmt"
)

func buildInvertedIndexBlock() (zoneBlock []byte, ipBlock []byte, hostBlock []byte) {
	nopKVFlusher := kv.NewNopFlusher()
	seriesFlusher := NewTagFlusher(nopKVFlusher)
	// disable auto reset to pick the entrySetBuffer
	/////////////////////////
	// tag id mapping relation
	/////////////////////////
	ipMapping := map[uint32]string{
		1: "192.168.1.1",
		2: "192.168.1.2",
		3: "192.168.1.3",
		4: "192.168.2.4",
		5: "192.168.2.5",
		6: "192.168.2.6",
		7: "192.168.3.7",
		8: "192.168.3.8",
		9: "192.168.3.9"}
	zoneMapping := map[uint32]string{
		1: "nj",
		2: "sh",
		3: "bj"}
	hostMapping := map[uint32]string{
		1: "eleme-dev-nj-1",
		2: "eleme-dev-nj-2",
		3: "eleme-dev-nj-3",
		4: "eleme-dev-sh-4",
		5: "eleme-dev-sh-5",
		6: "eleme-dev-sh-6",
		7: "eleme-dev-bj-7",
		8: "eleme-dev-bj-8",
		9: "eleme-dev-bj-9"}
	flush := func(mapping map[uint32]string) {
		for id, value := range mapping {
			seriesFlusher.FlushTagValue(value, id)
		}
	}
	/////////////////////////
	// flush zone tag, tagID: 20
	/////////////////////////
	flush(zoneMapping)
	// pick the zoneBlock buffer
	_ = seriesFlusher.FlushTagKeyID(20)
	zoneBlock = append(zoneBlock, nopKVFlusher.Bytes()...)

	/////////////////////////
	// flush ip tag, tagID: 21
	/////////////////////////
	flush(ipMapping)
	// pick the ipBlock buffer
	_ = seriesFlusher.FlushTagKeyID(21)
	ipBlock = append(ipBlock, nopKVFlusher.Bytes()...)

	/////////////////////////
	// flush host tag, tagID: 22
	/////////////////////////
	flush(hostMapping)
	// pick the hostBlock buffer
	_ = seriesFlusher.FlushTagKeyID(22)
	hostBlock = append(hostBlock, nopKVFlusher.Bytes()...)
	return zoneBlock, ipBlock, hostBlock
}

func buildSeriesIndexReader(ctrl *gomock.Controller) TagReader {
	zoneBlock, ipBlock, hostBlock := buildInvertedIndexBlock()
	// mock readers
	mockReader := table.NewMockReader(ctrl)
	mockReader.EXPECT().Get(uint32(19)).Return(nil).AnyTimes()
	mockReader.EXPECT().Get(uint32(20)).Return(zoneBlock).AnyTimes()
	mockReader.EXPECT().Get(uint32(21)).Return(ipBlock).AnyTimes()
	mockReader.EXPECT().Get(uint32(22)).Return(hostBlock).AnyTimes()
	// build series index reader
	return NewReader([]table.Reader{mockReader})
}

func TestTagReader_FindValueIDsForTagKeyID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	reader := buildSeriesIndexReader(ctrl)
	// read not tagID key
	idSet, err := reader.GetTagValueIDsForTagKeyID(19)
	assert.NotNil(t, err)
	assert.Nil(t, idSet)

	// read zone block
	idSet, err = reader.GetTagValueIDsForTagKeyID(20)
	assert.Nil(t, err)
	assert.NotNil(t, idSet)
	assert.Equal(t, roaring.BitmapOf(1, 2, 3), idSet)
}

func TestTagReader_FindValueIDsByExprForTagKeyID_bad_case(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	defer func() {
		newTagKVEntrySetFunc = newTagKVEntrySet
	}()

	reader := buildSeriesIndexReader(ctrl)

	// tagID not exist
	idSet, err := reader.FindValueIDsByExprForTagKeyID(19, nil)
	assert.Error(t, err)
	assert.Nil(t, idSet)

	// find zone with bad expression
	idSet, err = reader.FindValueIDsByExprForTagKeyID(20, nil)
	assert.Error(t, err)
	assert.Nil(t, idSet)

	mockEntry := NewMockTagKVEntrySetINTF(ctrl)
	newTagKVEntrySetFunc = func(block []byte) (intf TagKVEntrySetINTF, err error) {
		return mockEntry, err
	}
	mockEntry.EXPECT().TrieTree().Return(nil, fmt.Errorf("err"))
	idSet, err = reader.FindValueIDsByExprForTagKeyID(22, &stmt.EqualsExpr{Key: "host", Value: "eleme-dev-sh-4"})
	assert.Error(t, err)
	assert.Nil(t, idSet)

}

func TestTagReader_FindSeriesIDsByExprForTagID_EqualExpr(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	reader := buildSeriesIndexReader(ctrl)

	idSet, err := reader.FindValueIDsByExprForTagKeyID(22, &stmt.EqualsExpr{Key: "host", Value: "eleme-dev-sh-4"})
	assert.NoError(t, err)
	assert.Equal(t, roaring.BitmapOf(4), idSet)
	// find not existed host
	_, err = reader.FindValueIDsByExprForTagKeyID(22, &stmt.EqualsExpr{Key: "host", Value: "eleme-dev-sh-41"})
	assert.Error(t, err)
}

func TestTagReader_FindValueIDsByExprForTagKeyID_InExpr(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	reader := buildSeriesIndexReader(ctrl)

	// find existed host
	idSet, err := reader.FindValueIDsByExprForTagKeyID(22, &stmt.InExpr{
		Key: "host", Values: []string{"eleme-dev-sh-4", "eleme-dev-sh-5", "eleme-dev-sh-55"}},
	)
	assert.NoError(t, err)
	assert.Equal(t, roaring.BitmapOf(4, 5), idSet)

	// find not existed host
	_, err = reader.FindValueIDsByExprForTagKeyID(22, &stmt.InExpr{
		Key: "host", Values: []string{"eleme-dev-sh-55"}},
	)
	assert.Error(t, err)
}

func TestTagReader_FindSeriesIDsByExprForTagID_LikeExpr(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	reader := buildSeriesIndexReader(ctrl)

	// find existed host
	idSet, err := reader.FindValueIDsByExprForTagKeyID(22, &stmt.LikeExpr{Key: "host", Value: "eleme-dev-sh-"})
	assert.NoError(t, err)
	assert.Equal(t, roaring.BitmapOf(4, 5, 6), idSet)
	// find not existed host
	_, err = reader.FindValueIDsByExprForTagKeyID(22, &stmt.InExpr{Key: "host", Values: []string{"eleme-dev-sh---"}})
	assert.Error(t, err)
}

func TestTagReader_FindSeriesIDsByExprForTagID_RegexExpr(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	reader := buildSeriesIndexReader(ctrl)

	idSet, err := reader.FindValueIDsByExprForTagKeyID(22, &stmt.RegexExpr{Key: "host", Regexp: "eleme-dev-sh-"})
	assert.NoError(t, err)
	assert.Equal(t, roaring.BitmapOf(4, 5, 6), idSet)

	// find not existed host
	_, err = reader.FindValueIDsByExprForTagKeyID(22, &stmt.RegexExpr{Key: "host", Regexp: "eleme-prod-sh-"})
	assert.Error(t, err)
}

func TestTagReader_SuggestTagValues(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	defer func() {
		newTagKVEntrySetFunc = newTagKVEntrySet
	}()
	reader := buildSeriesIndexReader(ctrl)

	// tagID not exist
	assert.Nil(t, reader.SuggestTagValues(19, "", 10000000))
	// search ip
	assert.Len(t, reader.SuggestTagValues(21, "192", 1000), 9)
	assert.Len(t, reader.SuggestTagValues(21, "192", 3), 3)

	// mock corruption
	mockReader := table.NewMockReader(ctrl)
	mockReader.EXPECT().Get(uint32(18)).Return([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}).AnyTimes()
	corruptedReader := NewReader([]table.Reader{mockReader})
	assert.Nil(t, corruptedReader.SuggestTagValues(18, "", 10000000))

	mockEntry := NewMockTagKVEntrySetINTF(ctrl)
	newTagKVEntrySetFunc = func(block []byte) (intf TagKVEntrySetINTF, err error) {
		return mockEntry, err
	}
	mockEntry.EXPECT().TrieTree().Return(nil, fmt.Errorf("err"))
	assert.Len(t, reader.SuggestTagValues(21, "192", 1000), 0)
}

func Test_InvertedIndexReader_WalkTagValues(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	reader := buildSeriesIndexReader(ctrl)

	// tagID not exist
	assert.NotPanics(t, func() {
		_ = reader.WalkTagValues(
			19,
			"",
			func(tagValue []byte, tagValueID uint32) bool {
				panic("tagID doesn't exist!")
			})
	})

	// search ip
	var ipCount1 int
	assert.Nil(t, reader.WalkTagValues(
		21,
		"192",
		func(tagValue []byte, tagValueID uint32) bool {
			ipCount1++
			return true
		}))
	assert.Equal(t, 9, ipCount1)

	// break case
	var ipCount2 int
	assert.Nil(t, reader.WalkTagValues(
		21,
		"192",
		func(tagValue []byte, tagValueID uint32) bool {
			ipCount2++
			return ipCount2 != 3
		}))
	assert.Equal(t, 3, ipCount2)
}

func TestTagReader_WalkTagValues_err(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	defer func() {
		newTagKVEntrySetFunc = newTagKVEntrySet
	}()
	reader := buildSeriesIndexReader(ctrl)

	mockEntry := NewMockTagKVEntrySetINTF(ctrl)
	newTagKVEntrySetFunc = func(block []byte) (intf TagKVEntrySetINTF, err error) {
		return mockEntry, err
	}
	mockEntry.EXPECT().TrieTree().Return(nil, fmt.Errorf("err"))
	ipCount2 := 0
	assert.Nil(t, reader.WalkTagValues(
		21,
		"192",
		func(tagValue []byte, tagValueID uint32) bool {
			ipCount2++
			return true
		}))
	assert.Equal(t, 0, ipCount2)
}