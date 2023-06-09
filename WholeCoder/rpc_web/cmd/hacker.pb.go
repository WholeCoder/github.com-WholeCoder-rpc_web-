// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hacker.proto

package cmd

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

type Hacker struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hacker) Reset()         { *m = Hacker{} }
func (m *Hacker) String() string { return proto.CompactTextString(m) }
func (*Hacker) ProtoMessage()    {}
func (*Hacker) Descriptor() ([]byte, []int) {
	return fileDescriptor_3aa5634b99676f1c, []int{0}
}

func (m *Hacker) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hacker.Unmarshal(m, b)
}
func (m *Hacker) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hacker.Marshal(b, m, deterministic)
}
func (m *Hacker) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hacker.Merge(m, src)
}
func (m *Hacker) XXX_Size() int {
	return xxx_messageInfo_Hacker.Size(m)
}
func (m *Hacker) XXX_DiscardUnknown() {
	xxx_messageInfo_Hacker.DiscardUnknown(m)
}

var xxx_messageInfo_Hacker proto.InternalMessageInfo

func (m *Hacker) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Hacker)(nil), "main.Hacker")
}

func init() {
	proto.RegisterFile("hacker.proto", fileDescriptor_3aa5634b99676f1c)
}

var fileDescriptor_3aa5634b99676f1c = []byte{
	// 120 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0x48, 0x4c, 0xce,
	0x4e, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d, 0xcc, 0xcc, 0x53, 0x92,
	0xe1, 0x62, 0xf3, 0x00, 0x8b, 0x0a, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a,
	0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x46, 0x96, 0x5c, 0xbc, 0x10, 0xd9, 0xe0, 0xd4, 0xa2, 0xb2,
	0xcc, 0xe4, 0x54, 0x21, 0x0d, 0x2e, 0x8e, 0xe0, 0xc4, 0x4a, 0x8f, 0xd4, 0x9c, 0x9c, 0x7c, 0x21,
	0x1e, 0x3d, 0x90, 0x09, 0x7a, 0x10, 0x05, 0x52, 0x28, 0x3c, 0x25, 0x06, 0x27, 0xf6, 0x28, 0x56,
	0x3d, 0xeb, 0xe4, 0xdc, 0x94, 0x24, 0x36, 0xb0, 0x75, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x8d, 0xc7, 0xc1, 0xc1, 0x7e, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HackerServiceClient is the client API for HackerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HackerServiceClient interface {
	SayHello(ctx context.Context, in *Hacker, opts ...grpc.CallOption) (*Hacker, error)
}

type hackerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHackerServiceClient(cc grpc.ClientConnInterface) HackerServiceClient {
	return &hackerServiceClient{cc}
}

func (c *hackerServiceClient) SayHello(ctx context.Context, in *Hacker, opts ...grpc.CallOption) (*Hacker, error) {
	out := new(Hacker)
	err := c.cc.Invoke(ctx, "/main.HackerService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HackerServiceServer is the server API for HackerService service.
type HackerServiceServer interface {
	SayHello(context.Context, *Hacker) (*Hacker, error)
}

// UnimplementedHackerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHackerServiceServer struct {
}

func (*UnimplementedHackerServiceServer) SayHello(ctx context.Context, req *Hacker) (*Hacker, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}

func RegisterHackerServiceServer(s *grpc.Server, srv HackerServiceServer) {
	s.RegisterService(&_HackerService_serviceDesc, srv)
}

func _HackerService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Hacker)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HackerServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.HackerService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HackerServiceServer).SayHello(ctx, req.(*Hacker))
	}
	return interceptor(ctx, in, info, handler)
}

var _HackerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "main.HackerService",
	HandlerType: (*HackerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _HackerService_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hacker.proto",
}
