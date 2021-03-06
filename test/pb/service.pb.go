// Code generated by protoc-gen-go.
// source: service.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	service.proto

It has these top-level messages:
	MethodNameRequest
	MethodNameResponse
*/
package pb

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

type MethodNameRequest struct {
	Param1 string `protobuf:"bytes,1,opt,name=param1" json:"param1,omitempty"`
}

func (m *MethodNameRequest) Reset()                    { *m = MethodNameRequest{} }
func (m *MethodNameRequest) String() string            { return proto.CompactTextString(m) }
func (*MethodNameRequest) ProtoMessage()               {}
func (*MethodNameRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MethodNameRequest) GetParam1() string {
	if m != nil {
		return m.Param1
	}
	return ""
}

type MethodNameResponse struct {
	Result string `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
	Err    string `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *MethodNameResponse) Reset()                    { *m = MethodNameResponse{} }
func (m *MethodNameResponse) String() string            { return proto.CompactTextString(m) }
func (*MethodNameResponse) ProtoMessage()               {}
func (*MethodNameResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MethodNameResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func (m *MethodNameResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*MethodNameRequest)(nil), "pb.MethodNameRequest")
	proto.RegisterType((*MethodNameResponse)(nil), "pb.MethodNameResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Name service

type NameClient interface {
	MethodName(ctx context.Context, in *MethodNameRequest, opts ...grpc.CallOption) (*MethodNameResponse, error)
}

type nameClient struct {
	cc *grpc.ClientConn
}

func NewNameClient(cc *grpc.ClientConn) NameClient {
	return &nameClient{cc}
}

func (c *nameClient) MethodName(ctx context.Context, in *MethodNameRequest, opts ...grpc.CallOption) (*MethodNameResponse, error) {
	out := new(MethodNameResponse)
	err := grpc.Invoke(ctx, "/pb.Name/MethodName", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Name service

type NameServer interface {
	MethodName(context.Context, *MethodNameRequest) (*MethodNameResponse, error)
}

func RegisterNameServer(s *grpc.Server, srv NameServer) {
	s.RegisterService(&_Name_serviceDesc, srv)
}

func _Name_MethodName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MethodNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameServer).MethodName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Name/MethodName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameServer).MethodName(ctx, req.(*MethodNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Name_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Name",
	HandlerType: (*NameServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MethodName",
			Handler:    _Name_MethodName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xd2, 0xe6,
	0x12, 0xf4, 0x4d, 0x2d, 0xc9, 0xc8, 0x4f, 0xf1, 0x4b, 0xcc, 0x4d, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d,
	0x2d, 0x2e, 0x11, 0x12, 0xe3, 0x62, 0x2b, 0x48, 0x2c, 0x4a, 0xcc, 0x35, 0x94, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x0c, 0x82, 0xf2, 0x94, 0xec, 0xb8, 0x84, 0x90, 0x15, 0x17, 0x17, 0xe4, 0xe7, 0x15,
	0xa7, 0x82, 0x54, 0x17, 0xa5, 0x16, 0x97, 0xe6, 0x94, 0xc0, 0x54, 0x43, 0x78, 0x42, 0x02, 0x5c,
	0xcc, 0xa9, 0x45, 0x45, 0x12, 0x4c, 0x60, 0x41, 0x10, 0xd3, 0xc8, 0x95, 0x8b, 0x05, 0xa4, 0x53,
	0xc8, 0x96, 0x8b, 0x0b, 0x61, 0x8e, 0x90, 0xa8, 0x5e, 0x41, 0x92, 0x1e, 0x86, 0x23, 0xa4, 0xc4,
	0xd0, 0x85, 0x21, 0xd6, 0x29, 0x31, 0x38, 0x15, 0x72, 0xa9, 0x26, 0xe7, 0xe7, 0xea, 0xa5, 0xa4,
	0x96, 0x65, 0xe6, 0x42, 0xbc, 0x92, 0x54, 0x9a, 0xa6, 0x97, 0x58, 0x94, 0x9c, 0x91, 0x99, 0x5c,
	0x94, 0x9a, 0x92, 0x59, 0x92, 0x91, 0x59, 0x5c, 0x92, 0x5f, 0x54, 0xe9, 0x24, 0xee, 0x08, 0x12,
	0x73, 0x06, 0x8b, 0x79, 0x40, 0xc4, 0x02, 0x40, 0xca, 0xa3, 0x98, 0x0a, 0x92, 0x7e, 0x30, 0x32,
	0x2e, 0x62, 0x62, 0x76, 0x0f, 0x70, 0x5a, 0xc5, 0xa4, 0xe8, 0x02, 0x36, 0x29, 0x00, 0x66, 0x12,
	0xa6, 0xae, 0x24, 0x36, 0xb0, 0x35, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf4, 0x84, 0x19,
	0x5c, 0x42, 0x01, 0x00, 0x00,
}
