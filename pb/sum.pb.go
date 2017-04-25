// Code generated by protoc-gen-go.
// source: pb/sum.proto
// DO NOT EDIT!

/*
Package sum is a generated protocol buffer package.

It is generated from these files:
	pb/sum.proto

It has these top-level messages:
	SumRequest
	SumResponse
*/
package sum

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type SumRequest struct {
	A int64 `protobuf:"varint,1,opt,name=a" json:"a,omitempty"`
	B int64 `protobuf:"varint,2,opt,name=b" json:"b,omitempty"`
}

func (m *SumRequest) Reset()                    { *m = SumRequest{} }
func (m *SumRequest) String() string            { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()               {}
func (*SumRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SumRequest) GetA() int64 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *SumRequest) GetB() int64 {
	if m != nil {
		return m.B
	}
	return 0
}

type SumResponse struct {
	Result int64 `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
}

func (m *SumResponse) Reset()                    { *m = SumResponse{} }
func (m *SumResponse) String() string            { return proto.CompactTextString(m) }
func (*SumResponse) ProtoMessage()               {}
func (*SumResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SumResponse) GetResult() int64 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*SumRequest)(nil), "SumRequest")
	proto.RegisterType((*SumResponse)(nil), "SumResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Summator service

type SummatorClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
}

type summatorClient struct {
	cc *grpc.ClientConn
}

func NewSummatorClient(cc *grpc.ClientConn) SummatorClient {
	return &summatorClient{cc}
}

func (c *summatorClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := grpc.Invoke(ctx, "/Summator/Sum", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Summator service

type SummatorServer interface {
	Sum(context.Context, *SumRequest) (*SumResponse, error)
}

func RegisterSummatorServer(s *grpc.Server, srv SummatorServer) {
	s.RegisterService(&_Summator_serviceDesc, srv)
}

func _Summator_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SummatorServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Summator/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SummatorServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Summator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Summator",
	HandlerType: (*SummatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _Summator_Sum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/sum.proto",
}

func init() { proto.RegisterFile("pb/sum.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x48, 0xd2, 0x2f,
	0x2e, 0xcd, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xd2, 0xe0, 0xe2, 0x0a, 0x2e, 0xcd, 0x0d,
	0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0xe2, 0xe1, 0x62, 0x4c, 0x94, 0x60, 0x54, 0x60, 0xd4,
	0x60, 0x0e, 0x62, 0x4c, 0x04, 0xf1, 0x92, 0x24, 0x98, 0x20, 0xbc, 0x24, 0x25, 0x55, 0x2e, 0x6e,
	0xb0, 0xca, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x31, 0x2e, 0xb6, 0xa2, 0xd4, 0xe2, 0xd2,
	0x9c, 0x12, 0xa8, 0x7a, 0x28, 0xcf, 0x48, 0x8f, 0x8b, 0x23, 0xb8, 0x34, 0x37, 0x37, 0xb1, 0x24,
	0xbf, 0x48, 0x48, 0x89, 0x8b, 0x39, 0xb8, 0x34, 0x57, 0x88, 0x5b, 0x0f, 0x61, 0x85, 0x14, 0x8f,
	0x1e, 0x92, 0x29, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0x77, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xbf, 0xdb, 0x13, 0xf8, 0x97, 0x00, 0x00, 0x00,
}
