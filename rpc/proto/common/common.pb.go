// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: common.proto

package common

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TaskType int32

const (
	TaskType_UNKNOWN      TaskType = 0
	TaskType_Intermediate TaskType = 1
	TaskType_Leaf         TaskType = 2
)

var TaskType_name = map[int32]string{
	0: "UNKNOWN",
	1: "Intermediate",
	2: "Leaf",
}

var TaskType_value = map[string]int32{
	"UNKNOWN":      0,
	"Intermediate": 1,
	"Leaf":         2,
}

func (x TaskType) String() string {
	return proto.EnumName(TaskType_name, int32(x))
}

func (TaskType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

type TaskRequest struct {
	JobID                int64    `protobuf:"varint,1,opt,name=jobID,proto3" json:"jobID,omitempty"`
	ParentTaskID         string   `protobuf:"bytes,2,opt,name=parentTaskID,proto3" json:"parentTaskID,omitempty"`
	Type                 int32    `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	PhysicalPlan         []byte   `protobuf:"bytes,4,opt,name=physicalPlan,proto3" json:"physicalPlan,omitempty"`
	Payload              []byte   `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskRequest) Reset()         { *m = TaskRequest{} }
func (m *TaskRequest) String() string { return proto.CompactTextString(m) }
func (*TaskRequest) ProtoMessage()    {}
func (*TaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}
func (m *TaskRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TaskRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskRequest.Merge(m, src)
}
func (m *TaskRequest) XXX_Size() int {
	return m.Size()
}
func (m *TaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TaskRequest proto.InternalMessageInfo

func (m *TaskRequest) GetJobID() int64 {
	if m != nil {
		return m.JobID
	}
	return 0
}

func (m *TaskRequest) GetParentTaskID() string {
	if m != nil {
		return m.ParentTaskID
	}
	return ""
}

func (m *TaskRequest) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *TaskRequest) GetPhysicalPlan() []byte {
	if m != nil {
		return m.PhysicalPlan
	}
	return nil
}

func (m *TaskRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type TaskResponse struct {
	JobID                int64    `protobuf:"varint,1,opt,name=jobID,proto3" json:"jobID,omitempty"`
	TaskID               string   `protobuf:"bytes,2,opt,name=TaskID,proto3" json:"TaskID,omitempty"`
	Completed            bool     `protobuf:"varint,3,opt,name=completed,proto3" json:"completed,omitempty"`
	ErrMsg               string   `protobuf:"bytes,4,opt,name=errMsg,proto3" json:"errMsg,omitempty"`
	Payload              []byte   `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskResponse) Reset()         { *m = TaskResponse{} }
func (m *TaskResponse) String() string { return proto.CompactTextString(m) }
func (*TaskResponse) ProtoMessage()    {}
func (*TaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}
func (m *TaskResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TaskResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskResponse.Merge(m, src)
}
func (m *TaskResponse) XXX_Size() int {
	return m.Size()
}
func (m *TaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TaskResponse proto.InternalMessageInfo

func (m *TaskResponse) GetJobID() int64 {
	if m != nil {
		return m.JobID
	}
	return 0
}

func (m *TaskResponse) GetTaskID() string {
	if m != nil {
		return m.TaskID
	}
	return ""
}

func (m *TaskResponse) GetCompleted() bool {
	if m != nil {
		return m.Completed
	}
	return false
}

func (m *TaskResponse) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *TaskResponse) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterEnum("common.TaskType", TaskType_name, TaskType_value)
	proto.RegisterType((*TaskRequest)(nil), "common.TaskRequest")
	proto.RegisterType((*TaskResponse)(nil), "common.TaskResponse")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0x87, 0x3b, 0xfd, 0x93, 0xb6, 0xd3, 0x1c, 0xc2, 0x5a, 0x24, 0x88, 0x84, 0x90, 0x53, 0xf0,
	0x50, 0xc4, 0x9e, 0xbc, 0x4a, 0x0f, 0x2d, 0x6a, 0x95, 0xb5, 0xe2, 0x79, 0x9b, 0x8c, 0x5a, 0x4d,
	0xb2, 0x31, 0x59, 0x85, 0xbe, 0x83, 0x0f, 0xa0, 0x6f, 0xe4, 0xd1, 0x47, 0x90, 0xfa, 0x22, 0x92,
	0x4d, 0x44, 0x73, 0xe8, 0x6d, 0xbf, 0xdf, 0xee, 0x2c, 0xdf, 0xcc, 0xa0, 0x19, 0xc8, 0x38, 0x96,
	0xc9, 0x28, 0xcd, 0xa4, 0x92, 0xcc, 0x28, 0xc9, 0x7b, 0x07, 0x1c, 0x2c, 0x44, 0xfe, 0xc8, 0xe9,
	0xe9, 0x99, 0x72, 0xc5, 0x86, 0xd8, 0x79, 0x90, 0xcb, 0xd9, 0xc4, 0x06, 0x17, 0xfc, 0x16, 0x2f,
	0x81, 0x79, 0x68, 0xa6, 0x22, 0xa3, 0x44, 0x15, 0x4f, 0x67, 0x13, 0xbb, 0xe9, 0x82, 0xdf, 0xe7,
	0xb5, 0x8c, 0x31, 0x6c, 0xab, 0x75, 0x4a, 0x76, 0xcb, 0x05, 0xbf, 0xc3, 0xf5, 0x59, 0xd7, 0xdd,
	0xaf, 0xf3, 0x55, 0x20, 0xa2, 0xcb, 0x48, 0x24, 0x76, 0xdb, 0x05, 0xdf, 0xe4, 0xb5, 0x8c, 0xd9,
	0xd8, 0x4d, 0xc5, 0x3a, 0x92, 0x22, 0xb4, 0x3b, 0xfa, 0xfa, 0x17, 0xbd, 0x57, 0x40, 0xb3, 0x74,
	0xcb, 0x53, 0x99, 0xe4, 0xb4, 0x45, 0x6e, 0x17, 0x8d, 0x9a, 0x56, 0x45, 0x6c, 0x1f, 0xfb, 0x81,
	0x8c, 0xd3, 0x88, 0x14, 0x85, 0xda, 0xaa, 0xc7, 0xff, 0x82, 0xa2, 0x8a, 0xb2, 0xec, 0x3c, 0xbf,
	0xd3, 0x52, 0x7d, 0x5e, 0xd1, 0x76, 0x9d, 0x83, 0x31, 0xf6, 0x8a, 0x9f, 0x17, 0x45, 0x63, 0x03,
	0xec, 0x5e, 0xcf, 0x4f, 0xe7, 0x17, 0x37, 0x73, 0xab, 0xc1, 0x2c, 0x34, 0x67, 0x89, 0xa2, 0x2c,
	0xa6, 0x70, 0x25, 0x14, 0x59, 0xc0, 0x7a, 0xd8, 0x3e, 0x23, 0x71, 0x6b, 0x35, 0x8f, 0xa6, 0xe5,
	0x78, 0xaf, 0x28, 0x7b, 0x59, 0x05, 0xc4, 0x8e, 0xd1, 0x98, 0x8a, 0x24, 0x8c, 0x88, 0xed, 0x8c,
	0xaa, 0x7d, 0xfc, 0x9b, 0xfe, 0xde, 0xb0, 0x1e, 0x96, 0x6d, 0x7b, 0x0d, 0x1f, 0x0e, 0xe1, 0xc4,
	0xfa, 0xd8, 0x38, 0xf0, 0xb9, 0x71, 0xe0, 0x6b, 0xe3, 0xc0, 0xdb, 0xb7, 0xd3, 0x58, 0x1a, 0x7a,
	0x95, 0xe3, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x93, 0x76, 0xd8, 0xda, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TaskServiceClient is the client API for TaskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TaskServiceClient interface {
	Handle(ctx context.Context, opts ...grpc.CallOption) (TaskService_HandleClient, error)
}

type taskServiceClient struct {
	cc *grpc.ClientConn
}

func NewTaskServiceClient(cc *grpc.ClientConn) TaskServiceClient {
	return &taskServiceClient{cc}
}

func (c *taskServiceClient) Handle(ctx context.Context, opts ...grpc.CallOption) (TaskService_HandleClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TaskService_serviceDesc.Streams[0], "/common.TaskService/Handle", opts...)
	if err != nil {
		return nil, err
	}
	x := &taskServiceHandleClient{stream}
	return x, nil
}

type TaskService_HandleClient interface {
	Send(*TaskRequest) error
	Recv() (*TaskResponse, error)
	grpc.ClientStream
}

type taskServiceHandleClient struct {
	grpc.ClientStream
}

func (x *taskServiceHandleClient) Send(m *TaskRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *taskServiceHandleClient) Recv() (*TaskResponse, error) {
	m := new(TaskResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TaskServiceServer is the server API for TaskService service.
type TaskServiceServer interface {
	Handle(TaskService_HandleServer) error
}

// UnimplementedTaskServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTaskServiceServer struct {
}

func (*UnimplementedTaskServiceServer) Handle(srv TaskService_HandleServer) error {
	return status.Errorf(codes.Unimplemented, "method Handle not implemented")
}

func RegisterTaskServiceServer(s *grpc.Server, srv TaskServiceServer) {
	s.RegisterService(&_TaskService_serviceDesc, srv)
}

func _TaskService_Handle_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TaskServiceServer).Handle(&taskServiceHandleServer{stream})
}

type TaskService_HandleServer interface {
	Send(*TaskResponse) error
	Recv() (*TaskRequest, error)
	grpc.ServerStream
}

type taskServiceHandleServer struct {
	grpc.ServerStream
}

func (x *taskServiceHandleServer) Send(m *TaskResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *taskServiceHandleServer) Recv() (*TaskRequest, error) {
	m := new(TaskRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TaskService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "common.TaskService",
	HandlerType: (*TaskServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Handle",
			Handler:       _TaskService_Handle_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "common.proto",
}

func (m *TaskRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TaskRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TaskRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.PhysicalPlan) > 0 {
		i -= len(m.PhysicalPlan)
		copy(dAtA[i:], m.PhysicalPlan)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.PhysicalPlan)))
		i--
		dAtA[i] = 0x22
	}
	if m.Type != 0 {
		i = encodeVarintCommon(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ParentTaskID) > 0 {
		i -= len(m.ParentTaskID)
		copy(dAtA[i:], m.ParentTaskID)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.ParentTaskID)))
		i--
		dAtA[i] = 0x12
	}
	if m.JobID != 0 {
		i = encodeVarintCommon(dAtA, i, uint64(m.JobID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TaskResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TaskResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TaskResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ErrMsg) > 0 {
		i -= len(m.ErrMsg)
		copy(dAtA[i:], m.ErrMsg)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.ErrMsg)))
		i--
		dAtA[i] = 0x22
	}
	if m.Completed {
		i--
		if m.Completed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.TaskID) > 0 {
		i -= len(m.TaskID)
		copy(dAtA[i:], m.TaskID)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.TaskID)))
		i--
		dAtA[i] = 0x12
	}
	if m.JobID != 0 {
		i = encodeVarintCommon(dAtA, i, uint64(m.JobID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCommon(dAtA []byte, offset int, v uint64) int {
	offset -= sovCommon(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TaskRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.JobID != 0 {
		n += 1 + sovCommon(uint64(m.JobID))
	}
	l = len(m.ParentTaskID)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovCommon(uint64(m.Type))
	}
	l = len(m.PhysicalPlan)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TaskResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.JobID != 0 {
		n += 1 + sovCommon(uint64(m.JobID))
	}
	l = len(m.TaskID)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	if m.Completed {
		n += 2
	}
	l = len(m.ErrMsg)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCommon(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCommon(x uint64) (n int) {
	return sovCommon(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TaskRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TaskRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TaskRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field JobID", wireType)
			}
			m.JobID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.JobID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParentTaskID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ParentTaskID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PhysicalPlan", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PhysicalPlan = append(m.PhysicalPlan[:0], dAtA[iNdEx:postIndex]...)
			if m.PhysicalPlan == nil {
				m.PhysicalPlan = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCommon
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCommon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TaskResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TaskResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TaskResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field JobID", wireType)
			}
			m.JobID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.JobID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaskID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TaskID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Completed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Completed = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrMsg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ErrMsg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCommon
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCommon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCommon(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCommon
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthCommon
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthCommon
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCommon
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipCommon(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthCommon
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthCommon = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCommon   = fmt.Errorf("proto: integer overflow")
)
