// Code generated by protoc-gen-go. DO NOT EDIT.
// source: queue.proto

package v1

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

type Message struct {
	Payload              string   `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_96e4d7d76a734cd8, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

type Ok struct {
	Response             string   `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ok) Reset()         { *m = Ok{} }
func (m *Ok) String() string { return proto.CompactTextString(m) }
func (*Ok) ProtoMessage()    {}
func (*Ok) Descriptor() ([]byte, []int) {
	return fileDescriptor_96e4d7d76a734cd8, []int{1}
}

func (m *Ok) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ok.Unmarshal(m, b)
}
func (m *Ok) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ok.Marshal(b, m, deterministic)
}
func (m *Ok) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ok.Merge(m, src)
}
func (m *Ok) XXX_Size() int {
	return xxx_messageInfo_Ok.Size(m)
}
func (m *Ok) XXX_DiscardUnknown() {
	xxx_messageInfo_Ok.DiscardUnknown(m)
}

var xxx_messageInfo_Ok proto.InternalMessageInfo

func (m *Ok) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "v1.Message")
	proto.RegisterType((*Ok)(nil), "v1.Ok")
}

func init() {
	proto.RegisterFile("queue.proto", fileDescriptor_96e4d7d76a734cd8)
}

var fileDescriptor_96e4d7d76a734cd8 = []byte{
	// 150 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x2c, 0x4d, 0x2d,
	0x4d, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x33, 0x54, 0x52, 0xe6, 0x62, 0xf7,
	0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0x92, 0xe0, 0x62, 0x2f, 0x48, 0xac, 0xcc, 0xc9, 0x4f,
	0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x71, 0x95, 0x14, 0xb8, 0x98, 0xfc, 0xb3,
	0x85, 0xa4, 0xb8, 0x38, 0x8a, 0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0xa1, 0x0a, 0xe0, 0x7c,
	0x23, 0x4f, 0x2e, 0xd6, 0x40, 0x90, 0xc9, 0x42, 0x72, 0x5c, 0xec, 0x01, 0xa5, 0x49, 0x39, 0x99,
	0xc5, 0x19, 0x42, 0xdc, 0x7a, 0x65, 0x86, 0x7a, 0x50, 0xc3, 0xa5, 0xd8, 0x40, 0x1c, 0xff, 0x6c,
	0x21, 0x05, 0x2e, 0xce, 0xe0, 0xd2, 0xa4, 0xe2, 0xe4, 0xa2, 0xcc, 0xa4, 0x54, 0x21, 0xa8, 0xa0,
	0x14, 0xb2, 0xca, 0x24, 0x36, 0xb0, 0xe3, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x74,
	0xfc, 0x97, 0xab, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// QueueClient is the client API for Queue service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueueClient interface {
	Publish(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Ok, error)
	Subscribe(ctx context.Context, in *Ok, opts ...grpc.CallOption) (*Message, error)
}

type queueClient struct {
	cc grpc.ClientConnInterface
}

func NewQueueClient(cc grpc.ClientConnInterface) QueueClient {
	return &queueClient{cc}
}

func (c *queueClient) Publish(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Ok, error) {
	out := new(Ok)
	err := c.cc.Invoke(ctx, "/v1.Queue/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueClient) Subscribe(ctx context.Context, in *Ok, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/v1.Queue/Subscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueueServer is the server API for Queue service.
type QueueServer interface {
	Publish(context.Context, *Message) (*Ok, error)
	Subscribe(context.Context, *Ok) (*Message, error)
}

// UnimplementedQueueServer can be embedded to have forward compatible implementations.
type UnimplementedQueueServer struct {
}

func (*UnimplementedQueueServer) Publish(ctx context.Context, req *Message) (*Ok, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (*UnimplementedQueueServer) Subscribe(ctx context.Context, req *Ok) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}

func RegisterQueueServer(s *grpc.Server, srv QueueServer) {
	s.RegisterService(&_Queue_serviceDesc, srv)
}

func _Queue_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Queue/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueServer).Publish(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Queue_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ok)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueServer).Subscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Queue/Subscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueServer).Subscribe(ctx, req.(*Ok))
	}
	return interceptor(ctx, in, info, handler)
}

var _Queue_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.Queue",
	HandlerType: (*QueueServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _Queue_Publish_Handler,
		},
		{
			MethodName: "Subscribe",
			Handler:    _Queue_Subscribe_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "queue.proto",
}
