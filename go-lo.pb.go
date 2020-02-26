// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go-lo.proto

package golo

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

type Context struct {
	JobName              string   `protobuf:"bytes,1,opt,name=jobName,proto3" json:"jobName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Context) Reset()         { *m = Context{} }
func (m *Context) String() string { return proto.CompactTextString(m) }
func (*Context) ProtoMessage()    {}
func (*Context) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e939d8b2166d8ec, []int{0}
}

func (m *Context) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Context.Unmarshal(m, b)
}
func (m *Context) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Context.Marshal(b, m, deterministic)
}
func (m *Context) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Context.Merge(m, src)
}
func (m *Context) XXX_Size() int {
	return xxx_messageInfo_Context.Size(m)
}
func (m *Context) XXX_DiscardUnknown() {
	xxx_messageInfo_Context.DiscardUnknown(m)
}

var xxx_messageInfo_Context proto.InternalMessageInfo

func (m *Context) GetJobName() string {
	if m != nil {
		return m.JobName
	}
	return ""
}

type ResponseTag struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseTag) Reset()         { *m = ResponseTag{} }
func (m *ResponseTag) String() string { return proto.CompactTextString(m) }
func (*ResponseTag) ProtoMessage()    {}
func (*ResponseTag) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e939d8b2166d8ec, []int{1}
}

func (m *ResponseTag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseTag.Unmarshal(m, b)
}
func (m *ResponseTag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseTag.Marshal(b, m, deterministic)
}
func (m *ResponseTag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseTag.Merge(m, src)
}
func (m *ResponseTag) XXX_Size() int {
	return xxx_messageInfo_ResponseTag.Size(m)
}
func (m *ResponseTag) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseTag.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseTag proto.InternalMessageInfo

func (m *ResponseTag) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ResponseTag) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Response struct {
	Id                   string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	JobName              string         `protobuf:"bytes,2,opt,name=jobName,proto3" json:"jobName,omitempty"`
	Error                bool           `protobuf:"varint,3,opt,name=error,proto3" json:"error,omitempty"`
	Output               string         `protobuf:"bytes,4,opt,name=output,proto3" json:"output,omitempty"`
	Tags                 []*ResponseTag `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e939d8b2166d8ec, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Response) GetJobName() string {
	if m != nil {
		return m.JobName
	}
	return ""
}

func (m *Response) GetError() bool {
	if m != nil {
		return m.Error
	}
	return false
}

func (m *Response) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

func (m *Response) GetTags() []*ResponseTag {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterType((*Context)(nil), "golo.Context")
	proto.RegisterType((*ResponseTag)(nil), "golo.ResponseTag")
	proto.RegisterType((*Response)(nil), "golo.Response")
}

func init() { proto.RegisterFile("go-lo.proto", fileDescriptor_7e939d8b2166d8ec) }

var fileDescriptor_7e939d8b2166d8ec = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xd1, 0x4a, 0xc3, 0x30,
	0x14, 0x86, 0x6d, 0xd3, 0xad, 0xf3, 0x14, 0x87, 0x1e, 0x44, 0x82, 0x57, 0x25, 0x22, 0x14, 0xc1,
	0x82, 0x13, 0x9f, 0xc0, 0x3b, 0x2f, 0xbc, 0x08, 0x7b, 0x81, 0x94, 0x1d, 0x42, 0xb5, 0xee, 0x94,
	0x34, 0x15, 0x7d, 0x04, 0xdf, 0x5a, 0x9a, 0x66, 0xb0, 0xdd, 0xe5, 0x4b, 0xfe, 0x3f, 0xff, 0xf9,
	0x0f, 0x14, 0x96, 0x1f, 0x3b, 0xae, 0x7b, 0xc7, 0x9e, 0x31, 0xb3, 0xdc, 0xb1, 0xba, 0x83, 0xfc,
	0x95, 0xf7, 0x9e, 0x7e, 0x3c, 0x4a, 0xc8, 0x3f, 0xb8, 0x79, 0x37, 0x5f, 0x24, 0x93, 0x32, 0xa9,
	0xce, 0xf5, 0x01, 0xd5, 0x0b, 0x14, 0x9a, 0x86, 0x9e, 0xf7, 0x03, 0x6d, 0x8d, 0xc5, 0x4b, 0x10,
	0x9f, 0xf4, 0x1b, 0x45, 0xd3, 0x11, 0xaf, 0x61, 0xf1, 0x6d, 0xba, 0x91, 0x64, 0x1a, 0xee, 0x66,
	0x50, 0x7f, 0x09, 0xac, 0x0e, 0x3e, 0x5c, 0x43, 0xda, 0xee, 0xa2, 0x27, 0x6d, 0x77, 0xc7, 0x69,
	0xe9, 0x49, 0xda, 0xf4, 0x19, 0x39, 0xc7, 0x4e, 0x8a, 0x32, 0xa9, 0x56, 0x7a, 0x06, 0xbc, 0x81,
	0x25, 0x8f, 0xbe, 0x1f, 0xbd, 0xcc, 0x82, 0x3c, 0x12, 0xde, 0x43, 0xe6, 0x8d, 0x1d, 0xe4, 0xa2,
	0x14, 0x55, 0xb1, 0xb9, 0xaa, 0xa7, 0x56, 0xf5, 0xd1, 0xb4, 0x3a, 0x3c, 0x6f, 0x9e, 0x40, 0xbc,
	0x71, 0x83, 0x0f, 0x90, 0x6f, 0x5d, 0x6b, 0x2d, 0x39, 0xbc, 0x98, 0xa5, 0xb1, 0xfd, 0xed, 0xfa,
	0xd4, 0xa9, 0xce, 0x9a, 0x65, 0xd8, 0xd3, 0xf3, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x39, 0xbc,
	0x1c, 0x5d, 0x36, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// JobClient is the client API for Job service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JobClient interface {
	Trigger(ctx context.Context, in *Context, opts ...grpc.CallOption) (*Response, error)
}

type jobClient struct {
	cc *grpc.ClientConn
}

func NewJobClient(cc *grpc.ClientConn) JobClient {
	return &jobClient{cc}
}

func (c *jobClient) Trigger(ctx context.Context, in *Context, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/golo.Job/Trigger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobServer is the server API for Job service.
type JobServer interface {
	Trigger(context.Context, *Context) (*Response, error)
}

// UnimplementedJobServer can be embedded to have forward compatible implementations.
type UnimplementedJobServer struct {
}

func (*UnimplementedJobServer) Trigger(ctx context.Context, req *Context) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Trigger not implemented")
}

func RegisterJobServer(s *grpc.Server, srv JobServer) {
	s.RegisterService(&_Job_serviceDesc, srv)
}

func _Job_Trigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Context)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).Trigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/golo.Job/Trigger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).Trigger(ctx, req.(*Context))
	}
	return interceptor(ctx, in, info, handler)
}

var _Job_serviceDesc = grpc.ServiceDesc{
	ServiceName: "golo.Job",
	HandlerType: (*JobServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Trigger",
			Handler:    _Job_Trigger_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go-lo.proto",
}