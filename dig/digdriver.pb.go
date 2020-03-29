// Code generated by protoc-gen-go. DO NOT EDIT.
// source: digdriver.proto

package dig

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_09c56c430a23bc8e, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Collection struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Collection) Reset()         { *m = Collection{} }
func (m *Collection) String() string { return proto.CompactTextString(m) }
func (*Collection) ProtoMessage()    {}
func (*Collection) Descriptor() ([]byte, []int) {
	return fileDescriptor_09c56c430a23bc8e, []int{1}
}

func (m *Collection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Collection.Unmarshal(m, b)
}
func (m *Collection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Collection.Marshal(b, m, deterministic)
}
func (m *Collection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Collection.Merge(m, src)
}
func (m *Collection) XXX_Size() int {
	return xxx_messageInfo_Collection.Size(m)
}
func (m *Collection) XXX_DiscardUnknown() {
	xxx_messageInfo_Collection.DiscardUnknown(m)
}

var xxx_messageInfo_Collection proto.InternalMessageInfo

func (m *Collection) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type RowID struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RowID) Reset()         { *m = RowID{} }
func (m *RowID) String() string { return proto.CompactTextString(m) }
func (*RowID) ProtoMessage()    {}
func (*RowID) Descriptor() ([]byte, []int) {
	return fileDescriptor_09c56c430a23bc8e, []int{2}
}

func (m *RowID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RowID.Unmarshal(m, b)
}
func (m *RowID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RowID.Marshal(b, m, deterministic)
}
func (m *RowID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RowID.Merge(m, src)
}
func (m *RowID) XXX_Size() int {
	return xxx_messageInfo_RowID.Size(m)
}
func (m *RowID) XXX_DiscardUnknown() {
	xxx_messageInfo_RowID.DiscardUnknown(m)
}

var xxx_messageInfo_RowID proto.InternalMessageInfo

func (m *RowID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type RowRequest struct {
	Collection           string   `protobuf:"bytes,1,opt,name=collection,proto3" json:"collection,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	RequestID            uint64   `protobuf:"varint,3,opt,name=requestID,proto3" json:"requestID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RowRequest) Reset()         { *m = RowRequest{} }
func (m *RowRequest) String() string { return proto.CompactTextString(m) }
func (*RowRequest) ProtoMessage()    {}
func (*RowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_09c56c430a23bc8e, []int{3}
}

func (m *RowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RowRequest.Unmarshal(m, b)
}
func (m *RowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RowRequest.Marshal(b, m, deterministic)
}
func (m *RowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RowRequest.Merge(m, src)
}
func (m *RowRequest) XXX_Size() int {
	return xxx_messageInfo_RowRequest.Size(m)
}
func (m *RowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RowRequest proto.InternalMessageInfo

func (m *RowRequest) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

func (m *RowRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RowRequest) GetRequestID() uint64 {
	if m != nil {
		return m.RequestID
	}
	return 0
}

type Row struct {
	Id                   string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Data                 *_struct.Struct `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	RequestID            uint64          `protobuf:"varint,3,opt,name=requestID,proto3" json:"requestID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Row) Reset()         { *m = Row{} }
func (m *Row) String() string { return proto.CompactTextString(m) }
func (*Row) ProtoMessage()    {}
func (*Row) Descriptor() ([]byte, []int) {
	return fileDescriptor_09c56c430a23bc8e, []int{4}
}

func (m *Row) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Row.Unmarshal(m, b)
}
func (m *Row) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Row.Marshal(b, m, deterministic)
}
func (m *Row) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Row.Merge(m, src)
}
func (m *Row) XXX_Size() int {
	return xxx_messageInfo_Row.Size(m)
}
func (m *Row) XXX_DiscardUnknown() {
	xxx_messageInfo_Row.DiscardUnknown(m)
}

var xxx_messageInfo_Row proto.InternalMessageInfo

func (m *Row) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Row) GetData() *_struct.Struct {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Row) GetRequestID() uint64 {
	if m != nil {
		return m.RequestID
	}
	return 0
}

type CollectionInfo struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CollectionInfo) Reset()         { *m = CollectionInfo{} }
func (m *CollectionInfo) String() string { return proto.CompactTextString(m) }
func (*CollectionInfo) ProtoMessage()    {}
func (*CollectionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_09c56c430a23bc8e, []int{5}
}

func (m *CollectionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectionInfo.Unmarshal(m, b)
}
func (m *CollectionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectionInfo.Marshal(b, m, deterministic)
}
func (m *CollectionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionInfo.Merge(m, src)
}
func (m *CollectionInfo) XXX_Size() int {
	return xxx_messageInfo_CollectionInfo.Size(m)
}
func (m *CollectionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CollectionInfo proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Empty)(nil), "dig.Empty")
	proto.RegisterType((*Collection)(nil), "dig.Collection")
	proto.RegisterType((*RowID)(nil), "dig.RowID")
	proto.RegisterType((*RowRequest)(nil), "dig.RowRequest")
	proto.RegisterType((*Row)(nil), "dig.Row")
	proto.RegisterType((*CollectionInfo)(nil), "dig.CollectionInfo")
}

func init() { proto.RegisterFile("digdriver.proto", fileDescriptor_09c56c430a23bc8e) }

var fileDescriptor_09c56c430a23bc8e = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xcf, 0x4f, 0xfa, 0x40,
	0x10, 0xc5, 0xb3, 0xfc, 0xfc, 0x32, 0x24, 0xf0, 0x75, 0x3c, 0x40, 0x08, 0x31, 0xa4, 0x07, 0x25,
	0xd1, 0x94, 0x06, 0x4f, 0x5e, 0xb5, 0x86, 0xf4, 0xba, 0xdc, 0x3c, 0x09, 0xdd, 0xa5, 0xd9, 0x04,
	0x18, 0xdc, 0x6e, 0x6d, 0xf8, 0xc7, 0x3d, 0x1b, 0xb7, 0xc5, 0xd5, 0x6a, 0xe2, 0x6d, 0xf2, 0xe6,
	0xf3, 0x66, 0xf2, 0x66, 0xa0, 0x2f, 0x54, 0x22, 0xb4, 0x7a, 0x95, 0xda, 0x3f, 0x68, 0x32, 0x84,
	0x75, 0xa1, 0x92, 0xd1, 0x38, 0x21, 0x4a, 0xb6, 0x72, 0x66, 0xa5, 0x75, 0xb6, 0x99, 0xa5, 0x46,
	0x67, 0xb1, 0x29, 0x10, 0xaf, 0x0d, 0xcd, 0xc7, 0xdd, 0xc1, 0x1c, 0xbd, 0x09, 0xc0, 0x03, 0x6d,
	0xb7, 0x32, 0x36, 0x8a, 0xf6, 0x88, 0xd0, 0xd8, 0xaf, 0x76, 0x72, 0xc8, 0x26, 0x6c, 0xda, 0xe1,
	0xb6, 0xf6, 0x06, 0xd0, 0xe4, 0x94, 0x47, 0x21, 0xf6, 0xa0, 0xa6, 0x44, 0xd9, 0xaa, 0x29, 0xe1,
	0x3d, 0x01, 0x70, 0xca, 0xb9, 0x7c, 0xc9, 0x64, 0x6a, 0xf0, 0x02, 0x20, 0xfe, 0x1c, 0x54, 0x52,
	0x5f, 0x94, 0xd2, 0x5d, 0x3b, 0xb9, 0x71, 0x0c, 0x1d, 0x5d, 0x58, 0xa3, 0x70, 0x58, 0x9f, 0xb0,
	0x69, 0x83, 0x3b, 0xc1, 0x7b, 0x86, 0x3a, 0xa7, 0xbc, 0xba, 0x12, 0xaf, 0xa1, 0x21, 0x56, 0x66,
	0x65, 0xc7, 0x74, 0xe7, 0x03, 0xbf, 0xc8, 0xe8, 0x9f, 0x32, 0xfa, 0x4b, 0x9b, 0x91, 0x5b, 0xe8,
	0x8f, 0x0d, 0xff, 0xa1, 0xe7, 0x82, 0x47, 0xfb, 0x0d, 0xcd, 0xdf, 0x18, 0x74, 0x42, 0x95, 0x2c,
	0x29, 0xd3, 0xb1, 0xc4, 0x19, 0xf4, 0x16, 0xd2, 0x38, 0x24, 0x45, 0xf0, 0x85, 0x4a, 0x7c, 0x7b,
	0xb6, 0x51, 0xdf, 0xd6, 0xae, 0x1b, 0x30, 0xbc, 0x83, 0xb3, 0x6f, 0x86, 0x8f, 0x99, 0x58, 0xe5,
	0x46, 0xe7, 0x15, 0xc1, 0x52, 0x57, 0xd0, 0x5a, 0x48, 0x13, 0x85, 0xe9, 0x4f, 0xbe, 0x58, 0x6a,
	0x1f, 0x10, 0x30, 0xbc, 0x84, 0xf6, 0x42, 0x1a, 0x4e, 0xf9, 0x2f, 0xe4, 0xbf, 0x13, 0x19, 0x30,
	0xbc, 0x81, 0x6e, 0xc9, 0xdd, 0x1f, 0xa3, 0xb0, 0x64, 0xdd, 0xb3, 0x1c, 0x3b, 0x65, 0x01, 0x5b,
	0xb7, 0xec, 0xfd, 0x6e, 0xdf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x23, 0xb2, 0x2a, 0x1d, 0x49, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DigSourceClient is the client API for DigSource service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DigSourceClient interface {
	GetCollections(ctx context.Context, in *Empty, opts ...grpc.CallOption) (DigSource_GetCollectionsClient, error)
	GetCollectionInfo(ctx context.Context, in *Collection, opts ...grpc.CallOption) (*CollectionInfo, error)
	GetIDs(ctx context.Context, in *Collection, opts ...grpc.CallOption) (DigSource_GetIDsClient, error)
	GetRows(ctx context.Context, in *Collection, opts ...grpc.CallOption) (DigSource_GetRowsClient, error)
	GetRowsByID(ctx context.Context, opts ...grpc.CallOption) (DigSource_GetRowsByIDClient, error)
}

type digSourceClient struct {
	cc *grpc.ClientConn
}

func NewDigSourceClient(cc *grpc.ClientConn) DigSourceClient {
	return &digSourceClient{cc}
}

func (c *digSourceClient) GetCollections(ctx context.Context, in *Empty, opts ...grpc.CallOption) (DigSource_GetCollectionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DigSource_serviceDesc.Streams[0], "/dig.DigSource/GetCollections", opts...)
	if err != nil {
		return nil, err
	}
	x := &digSourceGetCollectionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DigSource_GetCollectionsClient interface {
	Recv() (*Collection, error)
	grpc.ClientStream
}

type digSourceGetCollectionsClient struct {
	grpc.ClientStream
}

func (x *digSourceGetCollectionsClient) Recv() (*Collection, error) {
	m := new(Collection)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *digSourceClient) GetCollectionInfo(ctx context.Context, in *Collection, opts ...grpc.CallOption) (*CollectionInfo, error) {
	out := new(CollectionInfo)
	err := c.cc.Invoke(ctx, "/dig.DigSource/GetCollectionInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *digSourceClient) GetIDs(ctx context.Context, in *Collection, opts ...grpc.CallOption) (DigSource_GetIDsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DigSource_serviceDesc.Streams[1], "/dig.DigSource/GetIDs", opts...)
	if err != nil {
		return nil, err
	}
	x := &digSourceGetIDsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DigSource_GetIDsClient interface {
	Recv() (*RowID, error)
	grpc.ClientStream
}

type digSourceGetIDsClient struct {
	grpc.ClientStream
}

func (x *digSourceGetIDsClient) Recv() (*RowID, error) {
	m := new(RowID)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *digSourceClient) GetRows(ctx context.Context, in *Collection, opts ...grpc.CallOption) (DigSource_GetRowsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DigSource_serviceDesc.Streams[2], "/dig.DigSource/GetRows", opts...)
	if err != nil {
		return nil, err
	}
	x := &digSourceGetRowsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DigSource_GetRowsClient interface {
	Recv() (*Row, error)
	grpc.ClientStream
}

type digSourceGetRowsClient struct {
	grpc.ClientStream
}

func (x *digSourceGetRowsClient) Recv() (*Row, error) {
	m := new(Row)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *digSourceClient) GetRowsByID(ctx context.Context, opts ...grpc.CallOption) (DigSource_GetRowsByIDClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DigSource_serviceDesc.Streams[3], "/dig.DigSource/GetRowsByID", opts...)
	if err != nil {
		return nil, err
	}
	x := &digSourceGetRowsByIDClient{stream}
	return x, nil
}

type DigSource_GetRowsByIDClient interface {
	Send(*RowRequest) error
	Recv() (*Row, error)
	grpc.ClientStream
}

type digSourceGetRowsByIDClient struct {
	grpc.ClientStream
}

func (x *digSourceGetRowsByIDClient) Send(m *RowRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *digSourceGetRowsByIDClient) Recv() (*Row, error) {
	m := new(Row)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DigSourceServer is the server API for DigSource service.
type DigSourceServer interface {
	GetCollections(*Empty, DigSource_GetCollectionsServer) error
	GetCollectionInfo(context.Context, *Collection) (*CollectionInfo, error)
	GetIDs(*Collection, DigSource_GetIDsServer) error
	GetRows(*Collection, DigSource_GetRowsServer) error
	GetRowsByID(DigSource_GetRowsByIDServer) error
}

// UnimplementedDigSourceServer can be embedded to have forward compatible implementations.
type UnimplementedDigSourceServer struct {
}

func (*UnimplementedDigSourceServer) GetCollections(req *Empty, srv DigSource_GetCollectionsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetCollections not implemented")
}
func (*UnimplementedDigSourceServer) GetCollectionInfo(ctx context.Context, req *Collection) (*CollectionInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCollectionInfo not implemented")
}
func (*UnimplementedDigSourceServer) GetIDs(req *Collection, srv DigSource_GetIDsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetIDs not implemented")
}
func (*UnimplementedDigSourceServer) GetRows(req *Collection, srv DigSource_GetRowsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetRows not implemented")
}
func (*UnimplementedDigSourceServer) GetRowsByID(srv DigSource_GetRowsByIDServer) error {
	return status.Errorf(codes.Unimplemented, "method GetRowsByID not implemented")
}

func RegisterDigSourceServer(s *grpc.Server, srv DigSourceServer) {
	s.RegisterService(&_DigSource_serviceDesc, srv)
}

func _DigSource_GetCollections_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DigSourceServer).GetCollections(m, &digSourceGetCollectionsServer{stream})
}

type DigSource_GetCollectionsServer interface {
	Send(*Collection) error
	grpc.ServerStream
}

type digSourceGetCollectionsServer struct {
	grpc.ServerStream
}

func (x *digSourceGetCollectionsServer) Send(m *Collection) error {
	return x.ServerStream.SendMsg(m)
}

func _DigSource_GetCollectionInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Collection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DigSourceServer).GetCollectionInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dig.DigSource/GetCollectionInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DigSourceServer).GetCollectionInfo(ctx, req.(*Collection))
	}
	return interceptor(ctx, in, info, handler)
}

func _DigSource_GetIDs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Collection)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DigSourceServer).GetIDs(m, &digSourceGetIDsServer{stream})
}

type DigSource_GetIDsServer interface {
	Send(*RowID) error
	grpc.ServerStream
}

type digSourceGetIDsServer struct {
	grpc.ServerStream
}

func (x *digSourceGetIDsServer) Send(m *RowID) error {
	return x.ServerStream.SendMsg(m)
}

func _DigSource_GetRows_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Collection)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DigSourceServer).GetRows(m, &digSourceGetRowsServer{stream})
}

type DigSource_GetRowsServer interface {
	Send(*Row) error
	grpc.ServerStream
}

type digSourceGetRowsServer struct {
	grpc.ServerStream
}

func (x *digSourceGetRowsServer) Send(m *Row) error {
	return x.ServerStream.SendMsg(m)
}

func _DigSource_GetRowsByID_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DigSourceServer).GetRowsByID(&digSourceGetRowsByIDServer{stream})
}

type DigSource_GetRowsByIDServer interface {
	Send(*Row) error
	Recv() (*RowRequest, error)
	grpc.ServerStream
}

type digSourceGetRowsByIDServer struct {
	grpc.ServerStream
}

func (x *digSourceGetRowsByIDServer) Send(m *Row) error {
	return x.ServerStream.SendMsg(m)
}

func (x *digSourceGetRowsByIDServer) Recv() (*RowRequest, error) {
	m := new(RowRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _DigSource_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dig.DigSource",
	HandlerType: (*DigSourceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCollectionInfo",
			Handler:    _DigSource_GetCollectionInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetCollections",
			Handler:       _DigSource_GetCollections_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetIDs",
			Handler:       _DigSource_GetIDs_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetRows",
			Handler:       _DigSource_GetRows_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetRowsByID",
			Handler:       _DigSource_GetRowsByID_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "digdriver.proto",
}
