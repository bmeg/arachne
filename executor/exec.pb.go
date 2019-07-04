// Code generated by protoc-gen-go. DO NOT EDIT.
// source: exec.proto

package executor

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type Code struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Function             string   `protobuf:"bytes,2,opt,name=function,proto3" json:"function,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Code) Reset()         { *m = Code{} }
func (m *Code) String() string { return proto.CompactTextString(m) }
func (*Code) ProtoMessage()    {}
func (*Code) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d737c7315c25422, []int{0}
}

func (m *Code) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Code.Unmarshal(m, b)
}
func (m *Code) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Code.Marshal(b, m, deterministic)
}
func (m *Code) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Code.Merge(m, src)
}
func (m *Code) XXX_Size() int {
	return xxx_messageInfo_Code.Size(m)
}
func (m *Code) XXX_DiscardUnknown() {
	xxx_messageInfo_Code.DiscardUnknown(m)
}

var xxx_messageInfo_Code proto.InternalMessageInfo

func (m *Code) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Code) GetFunction() string {
	if m != nil {
		return m.Function
	}
	return ""
}

type CompileResult struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CompileResult) Reset()         { *m = CompileResult{} }
func (m *CompileResult) String() string { return proto.CompactTextString(m) }
func (*CompileResult) ProtoMessage()    {}
func (*CompileResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d737c7315c25422, []int{1}
}

func (m *CompileResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompileResult.Unmarshal(m, b)
}
func (m *CompileResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompileResult.Marshal(b, m, deterministic)
}
func (m *CompileResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompileResult.Merge(m, src)
}
func (m *CompileResult) XXX_Size() int {
	return xxx_messageInfo_CompileResult.Size(m)
}
func (m *CompileResult) XXX_DiscardUnknown() {
	xxx_messageInfo_CompileResult.DiscardUnknown(m)
}

var xxx_messageInfo_CompileResult proto.InternalMessageInfo

func (m *CompileResult) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CompileResult) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type Input struct {
	//repeated google.protobuf.Value data = 1;
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Code                 uint32   `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Input) Reset()         { *m = Input{} }
func (m *Input) String() string { return proto.CompactTextString(m) }
func (*Input) ProtoMessage()    {}
func (*Input) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d737c7315c25422, []int{2}
}

func (m *Input) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Input.Unmarshal(m, b)
}
func (m *Input) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Input.Marshal(b, m, deterministic)
}
func (m *Input) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Input.Merge(m, src)
}
func (m *Input) XXX_Size() int {
	return xxx_messageInfo_Input.Size(m)
}
func (m *Input) XXX_DiscardUnknown() {
	xxx_messageInfo_Input.DiscardUnknown(m)
}

var xxx_messageInfo_Input proto.InternalMessageInfo

func (m *Input) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *Input) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

type Result struct {
	//google.protobuf.Value data = 1;
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d737c7315c25422, []int{3}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *Result) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*Code)(nil), "executor.Code")
	proto.RegisterType((*CompileResult)(nil), "executor.CompileResult")
	proto.RegisterType((*Input)(nil), "executor.Input")
	proto.RegisterType((*Result)(nil), "executor.Result")
}

func init() { proto.RegisterFile("exec.proto", fileDescriptor_4d737c7315c25422) }

var fileDescriptor_4d737c7315c25422 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0xcb, 0x4a, 0x04, 0x31,
	0x10, 0x34, 0x61, 0x1f, 0x63, 0xc3, 0xae, 0xd2, 0x08, 0x2e, 0x73, 0x92, 0x9c, 0xf6, 0x34, 0xca,
	0xf8, 0xf8, 0x81, 0xc5, 0x83, 0x37, 0x99, 0x3f, 0x58, 0x93, 0x16, 0x02, 0xeb, 0xf4, 0x90, 0x49,
	0xc0, 0xcf, 0x97, 0xc4, 0xcc, 0xe3, 0x30, 0xb7, 0xaa, 0xa6, 0x2a, 0x55, 0x15, 0x00, 0xfa, 0x25,
	0x5d, 0x75, 0x8e, 0x3d, 0x63, 0x11, 0x71, 0xf0, 0xec, 0xd4, 0x1b, 0xac, 0x4e, 0x6c, 0x08, 0x11,
	0x56, 0x9a, 0x0d, 0x1d, 0xc4, 0x83, 0x38, 0x5e, 0x37, 0x09, 0x63, 0x09, 0xc5, 0x77, 0x68, 0xb5,
	0xb7, 0xdc, 0x1e, 0x64, 0xba, 0x8f, 0x5c, 0xbd, 0xc2, 0xee, 0xc4, 0x3f, 0x9d, 0xbd, 0x50, 0x43,
	0x7d, 0xb8, 0x78, 0xdc, 0x83, 0xb4, 0x26, 0xd9, 0x77, 0x8d, 0xb4, 0x06, 0xef, 0x60, 0x4d, 0xce,
	0xb1, 0xcb, 0xce, 0x7f, 0xa2, 0x1e, 0x61, 0xfd, 0xd1, 0x76, 0xc1, 0xc7, 0x3c, 0x73, 0xf6, 0xe7,
	0x21, 0x2f, 0xe2, 0xb1, 0x83, 0x4c, 0x8f, 0x24, 0xac, 0x6a, 0xd8, 0xe4, 0x80, 0x25, 0xc7, 0x62,
	0x48, 0xed, 0xa1, 0x78, 0xcf, 0xfb, 0xf0, 0x05, 0xb6, 0xb9, 0x27, 0xee, 0xab, 0x61, 0x75, 0x15,
	0x27, 0x97, 0xf7, 0x73, 0x3e, 0x9b, 0xa2, 0xae, 0xb0, 0x86, 0xed, 0xa7, 0x63, 0x4d, 0x7d, 0x8f,
	0x37, 0x93, 0x2a, 0x35, 0x2f, 0x6f, 0xa7, 0xc3, 0xa0, 0x3f, 0x8a, 0x27, 0xf1, 0xb5, 0x49, 0x5f,
	0xfb, 0xfc, 0x17, 0x00, 0x00, 0xff, 0xff, 0x75, 0x3a, 0x2e, 0xd6, 0x68, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExecutorClient is the client API for Executor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExecutorClient interface {
	Compile(ctx context.Context, in *Code, opts ...grpc.CallOption) (*CompileResult, error)
	Process(ctx context.Context, opts ...grpc.CallOption) (Executor_ProcessClient, error)
}

type executorClient struct {
	cc *grpc.ClientConn
}

func NewExecutorClient(cc *grpc.ClientConn) ExecutorClient {
	return &executorClient{cc}
}

func (c *executorClient) Compile(ctx context.Context, in *Code, opts ...grpc.CallOption) (*CompileResult, error) {
	out := new(CompileResult)
	err := c.cc.Invoke(ctx, "/executor.Executor/Compile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executorClient) Process(ctx context.Context, opts ...grpc.CallOption) (Executor_ProcessClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Executor_serviceDesc.Streams[0], "/executor.Executor/Process", opts...)
	if err != nil {
		return nil, err
	}
	x := &executorProcessClient{stream}
	return x, nil
}

type Executor_ProcessClient interface {
	Send(*Input) error
	Recv() (*Result, error)
	grpc.ClientStream
}

type executorProcessClient struct {
	grpc.ClientStream
}

func (x *executorProcessClient) Send(m *Input) error {
	return x.ClientStream.SendMsg(m)
}

func (x *executorProcessClient) Recv() (*Result, error) {
	m := new(Result)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ExecutorServer is the server API for Executor service.
type ExecutorServer interface {
	Compile(context.Context, *Code) (*CompileResult, error)
	Process(Executor_ProcessServer) error
}

func RegisterExecutorServer(s *grpc.Server, srv ExecutorServer) {
	s.RegisterService(&_Executor_serviceDesc, srv)
}

func _Executor_Compile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Code)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).Compile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/executor.Executor/Compile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).Compile(ctx, req.(*Code))
	}
	return interceptor(ctx, in, info, handler)
}

func _Executor_Process_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ExecutorServer).Process(&executorProcessServer{stream})
}

type Executor_ProcessServer interface {
	Send(*Result) error
	Recv() (*Input, error)
	grpc.ServerStream
}

type executorProcessServer struct {
	grpc.ServerStream
}

func (x *executorProcessServer) Send(m *Result) error {
	return x.ServerStream.SendMsg(m)
}

func (x *executorProcessServer) Recv() (*Input, error) {
	m := new(Input)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Executor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "executor.Executor",
	HandlerType: (*ExecutorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Compile",
			Handler:    _Executor_Compile_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Process",
			Handler:       _Executor_Process_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "exec.proto",
}
