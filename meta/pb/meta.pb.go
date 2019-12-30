// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/meta.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdd022eb384fa38d, []int{0}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdd022eb384fa38d, []int{1}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdd022eb384fa38d, []int{2}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetReply struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetReply) Reset()         { *m = GetReply{} }
func (m *GetReply) String() string { return proto.CompactTextString(m) }
func (*GetReply) ProtoMessage()    {}
func (*GetReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdd022eb384fa38d, []int{3}
}

func (m *GetReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetReply.Unmarshal(m, b)
}
func (m *GetReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetReply.Marshal(b, m, deterministic)
}
func (m *GetReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetReply.Merge(m, src)
}
func (m *GetReply) XXX_Size() int {
	return xxx_messageInfo_GetReply.Size(m)
}
func (m *GetReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetReply proto.InternalMessageInfo

func (m *GetReply) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetReply) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListRequest struct {
	Limit                int32    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               int32    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdd022eb384fa38d, []int{4}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type ListReply struct {
	Metas                []*MetaModel `protobuf:"bytes,1,rep,name=metas,proto3" json:"metas,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ListReply) Reset()         { *m = ListReply{} }
func (m *ListReply) String() string { return proto.CompactTextString(m) }
func (*ListReply) ProtoMessage()    {}
func (*ListReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdd022eb384fa38d, []int{5}
}

func (m *ListReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListReply.Unmarshal(m, b)
}
func (m *ListReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListReply.Marshal(b, m, deterministic)
}
func (m *ListReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListReply.Merge(m, src)
}
func (m *ListReply) XXX_Size() int {
	return xxx_messageInfo_ListReply.Size(m)
}
func (m *ListReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ListReply.DiscardUnknown(m)
}

var xxx_messageInfo_ListReply proto.InternalMessageInfo

func (m *ListReply) GetMetas() []*MetaModel {
	if m != nil {
		return m.Metas
	}
	return nil
}

type MetaModel struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetaModel) Reset()         { *m = MetaModel{} }
func (m *MetaModel) String() string { return proto.CompactTextString(m) }
func (*MetaModel) ProtoMessage()    {}
func (*MetaModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdd022eb384fa38d, []int{6}
}

func (m *MetaModel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetaModel.Unmarshal(m, b)
}
func (m *MetaModel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetaModel.Marshal(b, m, deterministic)
}
func (m *MetaModel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetaModel.Merge(m, src)
}
func (m *MetaModel) XXX_Size() int {
	return xxx_messageInfo_MetaModel.Size(m)
}
func (m *MetaModel) XXX_DiscardUnknown() {
	xxx_messageInfo_MetaModel.DiscardUnknown(m)
}

var xxx_messageInfo_MetaModel proto.InternalMessageInfo

func (m *MetaModel) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MetaModel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "pb.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "pb.CreateResponse")
	proto.RegisterType((*GetRequest)(nil), "pb.GetRequest")
	proto.RegisterType((*GetReply)(nil), "pb.GetReply")
	proto.RegisterType((*ListRequest)(nil), "pb.ListRequest")
	proto.RegisterType((*ListReply)(nil), "pb.ListReply")
	proto.RegisterType((*MetaModel)(nil), "pb.MetaModel")
}

func init() { proto.RegisterFile("pb/meta.proto", fileDescriptor_bdd022eb384fa38d) }

var fileDescriptor_bdd022eb384fa38d = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xb1, 0x4e, 0xf3, 0x30,
	0x14, 0x85, 0x95, 0xb4, 0xa9, 0xfe, 0xdc, 0xfe, 0x09, 0x60, 0x21, 0x54, 0x45, 0x1d, 0x2a, 0x47,
	0x42, 0x15, 0x43, 0x82, 0x0a, 0x13, 0x4c, 0x88, 0xa1, 0x0b, 0x5d, 0xf2, 0x06, 0x8e, 0x72, 0x5b,
	0x59, 0x4a, 0x62, 0x53, 0x9b, 0xa1, 0x2b, 0xaf, 0x80, 0xc4, 0x8b, 0xf1, 0x0a, 0x3c, 0x08, 0xb2,
	0xdd, 0xa6, 0x81, 0xa9, 0x5b, 0xee, 0xf1, 0xf9, 0x4e, 0xec, 0x73, 0x21, 0x92, 0x65, 0xde, 0xa0,
	0x66, 0x99, 0xdc, 0x0a, 0x2d, 0x88, 0x2f, 0xcb, 0x64, 0xba, 0x11, 0x62, 0x53, 0x63, 0xce, 0x24,
	0xcf, 0x59, 0xdb, 0x0a, 0xcd, 0x34, 0x17, 0xad, 0x72, 0x0e, 0x9a, 0x42, 0xf4, 0xbc, 0x45, 0xa6,
	0xb1, 0xc0, 0xd7, 0x37, 0x54, 0x9a, 0x10, 0x18, 0xb6, 0xac, 0xc1, 0x89, 0x37, 0xf3, 0xe6, 0x61,
	0x61, 0xbf, 0xe9, 0x3d, 0xc4, 0x07, 0x93, 0x92, 0xa2, 0x55, 0x48, 0x62, 0xf0, 0x79, 0xb5, 0xf7,
	0xf8, 0xbc, 0xea, 0x28, 0xbf, 0x47, 0x4d, 0x01, 0x96, 0xa8, 0x0f, 0xb9, 0x7f, 0x08, 0x9a, 0xc1,
	0x3f, 0x7b, 0x2a, 0xeb, 0xdd, 0x49, 0x69, 0x8f, 0x30, 0x7e, 0xe1, 0xaa, 0x8b, 0xbb, 0x84, 0xa0,
	0xe6, 0x0d, 0xd7, 0x96, 0x0a, 0x0a, 0x37, 0x90, 0x2b, 0x18, 0x89, 0xf5, 0x5a, 0xa1, 0xb6, 0x68,
	0x50, 0xec, 0x27, 0x7a, 0x0b, 0xa1, 0x83, 0xcd, 0xdf, 0x52, 0x08, 0x4c, 0x45, 0x6a, 0xe2, 0xcd,
	0x06, 0xf3, 0xf1, 0x22, 0xca, 0x64, 0x99, 0xad, 0x50, 0xb3, 0x95, 0xa8, 0xb0, 0x2e, 0xdc, 0x19,
	0xcd, 0x21, 0xec, 0xb4, 0x53, 0xee, 0xb7, 0xf8, 0xf4, 0x60, 0x68, 0x08, 0xf2, 0x04, 0x23, 0x57,
	0x16, 0xb9, 0x30, 0xc9, 0xbf, 0xda, 0x4d, 0x48, 0x5f, 0x72, 0x5d, 0xd2, 0xf3, 0xf7, 0xaf, 0xef,
	0x0f, 0x1f, 0x68, 0x60, 0x37, 0xf7, 0xe0, 0xdd, 0x90, 0x14, 0x06, 0x4b, 0xd4, 0x24, 0x36, 0xe6,
	0x63, 0x85, 0xc9, 0xff, 0x6e, 0x36, 0xcf, 0xb8, 0x86, 0xa1, 0x79, 0x13, 0x39, 0x33, 0x6a, 0xaf,
	0x9a, 0x24, 0x3a, 0x0a, 0xb2, 0xde, 0x95, 0x23, 0xbb, 0xe8, 0xbb, 0x9f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x89, 0x0b, 0xc8, 0x21, 0x1b, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MetaClient is the client API for Meta service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MetaClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error)
}

type metaClient struct {
	cc *grpc.ClientConn
}

func NewMetaClient(cc *grpc.ClientConn) MetaClient {
	return &metaClient{cc}
}

func (c *metaClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/pb.Meta/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error) {
	out := new(GetReply)
	err := c.cc.Invoke(ctx, "/pb.Meta/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error) {
	out := new(ListReply)
	err := c.cc.Invoke(ctx, "/pb.Meta/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetaServer is the server API for Meta service.
type MetaServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Get(context.Context, *GetRequest) (*GetReply, error)
	List(context.Context, *ListRequest) (*ListReply, error)
}

// UnimplementedMetaServer can be embedded to have forward compatible implementations.
type UnimplementedMetaServer struct {
}

func (*UnimplementedMetaServer) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedMetaServer) Get(ctx context.Context, req *GetRequest) (*GetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedMetaServer) List(ctx context.Context, req *ListRequest) (*ListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterMetaServer(s *grpc.Server, srv MetaServer) {
	s.RegisterService(&_Meta_serviceDesc, srv)
}

func _Meta_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Meta/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meta_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Meta/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meta_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Meta/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Meta_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Meta",
	HandlerType: (*MetaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Meta_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Meta_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Meta_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/meta.proto",
}
