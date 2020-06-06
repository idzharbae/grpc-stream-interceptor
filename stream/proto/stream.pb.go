// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stream.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SortIntegerReq struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SortIntegerReq) Reset()         { *m = SortIntegerReq{} }
func (m *SortIntegerReq) String() string { return proto.CompactTextString(m) }
func (*SortIntegerReq) ProtoMessage()    {}
func (*SortIntegerReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb17ef3f514bfe54, []int{0}
}

func (m *SortIntegerReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SortIntegerReq.Unmarshal(m, b)
}
func (m *SortIntegerReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SortIntegerReq.Marshal(b, m, deterministic)
}
func (m *SortIntegerReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SortIntegerReq.Merge(m, src)
}
func (m *SortIntegerReq) XXX_Size() int {
	return xxx_messageInfo_SortIntegerReq.Size(m)
}
func (m *SortIntegerReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SortIntegerReq.DiscardUnknown(m)
}

var xxx_messageInfo_SortIntegerReq proto.InternalMessageInfo

func (m *SortIntegerReq) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type SortIntegerResp struct {
	SortedNumbers        []int32  `protobuf:"varint,1,rep,packed,name=sorted_numbers,json=sortedNumbers,proto3" json:"sorted_numbers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SortIntegerResp) Reset()         { *m = SortIntegerResp{} }
func (m *SortIntegerResp) String() string { return proto.CompactTextString(m) }
func (*SortIntegerResp) ProtoMessage()    {}
func (*SortIntegerResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb17ef3f514bfe54, []int{1}
}

func (m *SortIntegerResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SortIntegerResp.Unmarshal(m, b)
}
func (m *SortIntegerResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SortIntegerResp.Marshal(b, m, deterministic)
}
func (m *SortIntegerResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SortIntegerResp.Merge(m, src)
}
func (m *SortIntegerResp) XXX_Size() int {
	return xxx_messageInfo_SortIntegerResp.Size(m)
}
func (m *SortIntegerResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SortIntegerResp.DiscardUnknown(m)
}

var xxx_messageInfo_SortIntegerResp proto.InternalMessageInfo

func (m *SortIntegerResp) GetSortedNumbers() []int32 {
	if m != nil {
		return m.SortedNumbers
	}
	return nil
}

type FibonacciReq struct {
	Offset               int32    `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                int32    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FibonacciReq) Reset()         { *m = FibonacciReq{} }
func (m *FibonacciReq) String() string { return proto.CompactTextString(m) }
func (*FibonacciReq) ProtoMessage()    {}
func (*FibonacciReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb17ef3f514bfe54, []int{2}
}

func (m *FibonacciReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibonacciReq.Unmarshal(m, b)
}
func (m *FibonacciReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibonacciReq.Marshal(b, m, deterministic)
}
func (m *FibonacciReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibonacciReq.Merge(m, src)
}
func (m *FibonacciReq) XXX_Size() int {
	return xxx_messageInfo_FibonacciReq.Size(m)
}
func (m *FibonacciReq) XXX_DiscardUnknown() {
	xxx_messageInfo_FibonacciReq.DiscardUnknown(m)
}

var xxx_messageInfo_FibonacciReq proto.InternalMessageInfo

func (m *FibonacciReq) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *FibonacciReq) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type FibonacciResp struct {
	Num                  int32    `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FibonacciResp) Reset()         { *m = FibonacciResp{} }
func (m *FibonacciResp) String() string { return proto.CompactTextString(m) }
func (*FibonacciResp) ProtoMessage()    {}
func (*FibonacciResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb17ef3f514bfe54, []int{3}
}

func (m *FibonacciResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibonacciResp.Unmarshal(m, b)
}
func (m *FibonacciResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibonacciResp.Marshal(b, m, deterministic)
}
func (m *FibonacciResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibonacciResp.Merge(m, src)
}
func (m *FibonacciResp) XXX_Size() int {
	return xxx_messageInfo_FibonacciResp.Size(m)
}
func (m *FibonacciResp) XXX_DiscardUnknown() {
	xxx_messageInfo_FibonacciResp.DiscardUnknown(m)
}

var xxx_messageInfo_FibonacciResp proto.InternalMessageInfo

func (m *FibonacciResp) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

func init() {
	proto.RegisterType((*SortIntegerReq)(nil), "proto.SortIntegerReq")
	proto.RegisterType((*SortIntegerResp)(nil), "proto.SortIntegerResp")
	proto.RegisterType((*FibonacciReq)(nil), "proto.FibonacciReq")
	proto.RegisterType((*FibonacciResp)(nil), "proto.FibonacciResp")
}

func init() {
	proto.RegisterFile("stream.proto", fileDescriptor_bb17ef3f514bfe54)
}

var fileDescriptor_bb17ef3f514bfe54 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x5d, 0x4b, 0x7a, 0x18, 0x9b, 0x28, 0x63, 0x0d, 0xa1, 0xa7, 0x1a, 0x10, 0x72, 0x2a,
	0x45, 0x2f, 0x1e, 0x04, 0x3d, 0x09, 0x1e, 0x14, 0x49, 0x1e, 0x40, 0x92, 0x38, 0x95, 0x05, 0xb3,
	0x1b, 0x67, 0xa7, 0x7d, 0x5f, 0xdf, 0x44, 0x92, 0x8d, 0x92, 0xa2, 0x27, 0x4f, 0xb3, 0xff, 0xb7,
	0xff, 0xcc, 0xec, 0xbf, 0x30, 0x73, 0xc2, 0x54, 0x36, 0xab, 0x96, 0xad, 0x58, 0x0c, 0xfa, 0x92,
	0x66, 0x10, 0x15, 0x96, 0xe5, 0xc1, 0x08, 0xbd, 0x11, 0xe7, 0xf4, 0x81, 0x31, 0x4c, 0xcd, 0xb6,
	0xa9, 0x88, 0x13, 0xb5, 0x54, 0x59, 0x90, 0x0f, 0x2a, 0xbd, 0x86, 0xe3, 0x3d, 0xa7, 0x6b, 0xf1,
	0x02, 0x22, 0x67, 0x59, 0xe8, 0xf5, 0xc5, 0x7b, 0x5c, 0xa2, 0x96, 0x93, 0x2c, 0xc8, 0x43, 0x4f,
	0x9f, 0x3c, 0x4c, 0x6f, 0x60, 0x76, 0xaf, 0x2b, 0x6b, 0xca, 0xba, 0xd6, 0xc3, 0x06, 0xbb, 0xd9,
	0x38, 0x92, 0xef, 0x0d, 0x5e, 0xe1, 0x1c, 0x82, 0x77, 0xdd, 0x68, 0x49, 0x0e, 0x7b, 0xec, 0x45,
	0x7a, 0x0e, 0xe1, 0xa8, 0xdb, 0xb5, 0x78, 0x02, 0x13, 0xb3, 0x6d, 0x86, 0xde, 0xee, 0x78, 0xf9,
	0xa9, 0x20, 0x2c, 0xfa, 0x70, 0x05, 0xf1, 0x4e, 0xd7, 0x84, 0x77, 0x70, 0x34, 0x7a, 0x2c, 0x9e,
	0xf9, 0xd0, 0xab, 0xfd, 0xa8, 0x8b, 0xf8, 0x2f, 0xec, 0xda, 0xf4, 0x20, 0x53, 0xf8, 0x08, 0xf1,
	0x08, 0x77, 0x85, 0xcb, 0x5a, 0xf4, 0x8e, 0xfe, 0x31, 0x6c, 0xad, 0xf0, 0x16, 0xa2, 0x67, 0xd6,
	0x46, 0x7e, 0xa2, 0xe0, 0xe9, 0xe0, 0x1f, 0x7f, 0xcd, 0x62, 0xfe, 0x1b, 0x76, 0x23, 0xd6, 0xaa,
	0x9a, 0xf6, 0x17, 0x57, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x16, 0x3f, 0x48, 0x94, 0xc6, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StreamServiceClient is the client API for StreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StreamServiceClient interface {
	// Client side stream
	SortInteger(ctx context.Context, opts ...grpc.CallOption) (StreamService_SortIntegerClient, error)
	// Bi directional stream
	SortIntegerInteractive(ctx context.Context, opts ...grpc.CallOption) (StreamService_SortIntegerInteractiveClient, error)
	// Server side stream
	PrintFibonacci(ctx context.Context, in *FibonacciReq, opts ...grpc.CallOption) (StreamService_PrintFibonacciClient, error)
}

type streamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamServiceClient(cc grpc.ClientConnInterface) StreamServiceClient {
	return &streamServiceClient{cc}
}

func (c *streamServiceClient) SortInteger(ctx context.Context, opts ...grpc.CallOption) (StreamService_SortIntegerClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamService_serviceDesc.Streams[0], "/proto.StreamService/SortInteger", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceSortIntegerClient{stream}
	return x, nil
}

type StreamService_SortIntegerClient interface {
	Send(*SortIntegerReq) error
	CloseAndRecv() (*SortIntegerResp, error)
	grpc.ClientStream
}

type streamServiceSortIntegerClient struct {
	grpc.ClientStream
}

func (x *streamServiceSortIntegerClient) Send(m *SortIntegerReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamServiceSortIntegerClient) CloseAndRecv() (*SortIntegerResp, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SortIntegerResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamServiceClient) SortIntegerInteractive(ctx context.Context, opts ...grpc.CallOption) (StreamService_SortIntegerInteractiveClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamService_serviceDesc.Streams[1], "/proto.StreamService/SortIntegerInteractive", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceSortIntegerInteractiveClient{stream}
	return x, nil
}

type StreamService_SortIntegerInteractiveClient interface {
	Send(*SortIntegerReq) error
	Recv() (*SortIntegerResp, error)
	grpc.ClientStream
}

type streamServiceSortIntegerInteractiveClient struct {
	grpc.ClientStream
}

func (x *streamServiceSortIntegerInteractiveClient) Send(m *SortIntegerReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamServiceSortIntegerInteractiveClient) Recv() (*SortIntegerResp, error) {
	m := new(SortIntegerResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamServiceClient) PrintFibonacci(ctx context.Context, in *FibonacciReq, opts ...grpc.CallOption) (StreamService_PrintFibonacciClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamService_serviceDesc.Streams[2], "/proto.StreamService/PrintFibonacci", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServicePrintFibonacciClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamService_PrintFibonacciClient interface {
	Recv() (*FibonacciResp, error)
	grpc.ClientStream
}

type streamServicePrintFibonacciClient struct {
	grpc.ClientStream
}

func (x *streamServicePrintFibonacciClient) Recv() (*FibonacciResp, error) {
	m := new(FibonacciResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamServiceServer is the server API for StreamService service.
type StreamServiceServer interface {
	// Client side stream
	SortInteger(StreamService_SortIntegerServer) error
	// Bi directional stream
	SortIntegerInteractive(StreamService_SortIntegerInteractiveServer) error
	// Server side stream
	PrintFibonacci(*FibonacciReq, StreamService_PrintFibonacciServer) error
}

// UnimplementedStreamServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStreamServiceServer struct {
}

func (*UnimplementedStreamServiceServer) SortInteger(srv StreamService_SortIntegerServer) error {
	return status.Errorf(codes.Unimplemented, "method SortInteger not implemented")
}
func (*UnimplementedStreamServiceServer) SortIntegerInteractive(srv StreamService_SortIntegerInteractiveServer) error {
	return status.Errorf(codes.Unimplemented, "method SortIntegerInteractive not implemented")
}
func (*UnimplementedStreamServiceServer) PrintFibonacci(req *FibonacciReq, srv StreamService_PrintFibonacciServer) error {
	return status.Errorf(codes.Unimplemented, "method PrintFibonacci not implemented")
}

func RegisterStreamServiceServer(s *grpc.Server, srv StreamServiceServer) {
	s.RegisterService(&_StreamService_serviceDesc, srv)
}

func _StreamService_SortInteger_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServiceServer).SortInteger(&streamServiceSortIntegerServer{stream})
}

type StreamService_SortIntegerServer interface {
	SendAndClose(*SortIntegerResp) error
	Recv() (*SortIntegerReq, error)
	grpc.ServerStream
}

type streamServiceSortIntegerServer struct {
	grpc.ServerStream
}

func (x *streamServiceSortIntegerServer) SendAndClose(m *SortIntegerResp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamServiceSortIntegerServer) Recv() (*SortIntegerReq, error) {
	m := new(SortIntegerReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _StreamService_SortIntegerInteractive_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServiceServer).SortIntegerInteractive(&streamServiceSortIntegerInteractiveServer{stream})
}

type StreamService_SortIntegerInteractiveServer interface {
	Send(*SortIntegerResp) error
	Recv() (*SortIntegerReq, error)
	grpc.ServerStream
}

type streamServiceSortIntegerInteractiveServer struct {
	grpc.ServerStream
}

func (x *streamServiceSortIntegerInteractiveServer) Send(m *SortIntegerResp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamServiceSortIntegerInteractiveServer) Recv() (*SortIntegerReq, error) {
	m := new(SortIntegerReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _StreamService_PrintFibonacci_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FibonacciReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamServiceServer).PrintFibonacci(m, &streamServicePrintFibonacciServer{stream})
}

type StreamService_PrintFibonacciServer interface {
	Send(*FibonacciResp) error
	grpc.ServerStream
}

type streamServicePrintFibonacciServer struct {
	grpc.ServerStream
}

func (x *streamServicePrintFibonacciServer) Send(m *FibonacciResp) error {
	return x.ServerStream.SendMsg(m)
}

var _StreamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.StreamService",
	HandlerType: (*StreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SortInteger",
			Handler:       _StreamService_SortInteger_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "SortIntegerInteractive",
			Handler:       _StreamService_SortIntegerInteractive_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "PrintFibonacci",
			Handler:       _StreamService_PrintFibonacci_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "stream.proto",
}