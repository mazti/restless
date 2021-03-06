// Code generated by protoc-gen-go. DO NOT EDIT.
// source: base/pb/table.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	proto1 "github.com/mazti/restless/shared/proto"
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

type CreateTableRequest struct {
	Schema               string         `protobuf:"bytes,1,opt,name=schema,proto3" json:"schema,omitempty"`
	Name                 string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Columns              []*TableColumn `protobuf:"bytes,3,rep,name=columns,proto3" json:"columns,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateTableRequest) Reset()         { *m = CreateTableRequest{} }
func (m *CreateTableRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTableRequest) ProtoMessage()    {}
func (*CreateTableRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_275758ac75acb2d5, []int{0}
}

func (m *CreateTableRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTableRequest.Unmarshal(m, b)
}
func (m *CreateTableRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTableRequest.Marshal(b, m, deterministic)
}
func (m *CreateTableRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTableRequest.Merge(m, src)
}
func (m *CreateTableRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTableRequest.Size(m)
}
func (m *CreateTableRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTableRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTableRequest proto.InternalMessageInfo

func (m *CreateTableRequest) GetSchema() string {
	if m != nil {
		return m.Schema
	}
	return ""
}

func (m *CreateTableRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateTableRequest) GetColumns() []*TableColumn {
	if m != nil {
		return m.Columns
	}
	return nil
}

type TableColumn struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Attributes           string   `protobuf:"bytes,2,opt,name=attributes,proto3" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TableColumn) Reset()         { *m = TableColumn{} }
func (m *TableColumn) String() string { return proto.CompactTextString(m) }
func (*TableColumn) ProtoMessage()    {}
func (*TableColumn) Descriptor() ([]byte, []int) {
	return fileDescriptor_275758ac75acb2d5, []int{1}
}

func (m *TableColumn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TableColumn.Unmarshal(m, b)
}
func (m *TableColumn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TableColumn.Marshal(b, m, deterministic)
}
func (m *TableColumn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TableColumn.Merge(m, src)
}
func (m *TableColumn) XXX_Size() int {
	return xxx_messageInfo_TableColumn.Size(m)
}
func (m *TableColumn) XXX_DiscardUnknown() {
	xxx_messageInfo_TableColumn.DiscardUnknown(m)
}

var xxx_messageInfo_TableColumn proto.InternalMessageInfo

func (m *TableColumn) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TableColumn) GetAttributes() string {
	if m != nil {
		return m.Attributes
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateTableRequest)(nil), "pb.CreateTableRequest")
	proto.RegisterType((*TableColumn)(nil), "pb.TableColumn")
}

func init() { proto.RegisterFile("base/pb/table.proto", fileDescriptor_275758ac75acb2d5) }

var fileDescriptor_275758ac75acb2d5 = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xc1, 0x4a, 0xf4, 0x30,
	0x14, 0x85, 0x69, 0xe7, 0xff, 0x2b, 0x66, 0x16, 0x23, 0x51, 0x86, 0x52, 0x44, 0x86, 0xae, 0x46,
	0x17, 0x0d, 0xcc, 0xec, 0xdc, 0xc9, 0xe0, 0x52, 0x17, 0x45, 0x70, 0x7d, 0x53, 0x2f, 0x6d, 0xb1,
	0x69, 0x62, 0xef, 0xed, 0x42, 0xc5, 0x8d, 0xaf, 0xe0, 0xa3, 0xf9, 0x0a, 0x3e, 0x88, 0x4c, 0x5a,
	0xa5, 0xe0, 0x2e, 0xf7, 0xe4, 0xe6, 0x3b, 0x39, 0x47, 0x1c, 0x6b, 0x20, 0x54, 0x4e, 0x2b, 0x06,
	0xdd, 0x60, 0xe6, 0x3a, 0xcb, 0x56, 0x86, 0x4e, 0x27, 0xa7, 0xa5, 0xb5, 0x65, 0x83, 0x0a, 0x5c,
	0xad, 0xa0, 0x6d, 0x2d, 0x03, 0xd7, 0xb6, 0xa5, 0x61, 0x23, 0xd9, 0x96, 0x35, 0x57, 0xbd, 0xce,
	0x0a, 0x6b, 0x94, 0x81, 0x17, 0xae, 0x55, 0x87, 0xc4, 0x0d, 0x12, 0x29, 0xaa, 0xa0, 0xc3, 0x07,
	0xe5, 0xf7, 0xc6, 0x61, 0x78, 0x94, 0x3e, 0x0a, 0xb9, 0xeb, 0x10, 0x18, 0xef, 0xf6, 0x5e, 0x39,
	0x3e, 0xf5, 0x48, 0x2c, 0x97, 0x22, 0xa2, 0xa2, 0x42, 0x03, 0x71, 0xb0, 0x0a, 0xd6, 0x87, 0xf9,
	0x38, 0x49, 0x29, 0xfe, 0xb5, 0x60, 0x30, 0x0e, 0xbd, 0xea, 0xcf, 0xf2, 0x5c, 0x1c, 0x14, 0xb6,
	0xe9, 0x4d, 0x4b, 0xf1, 0x6c, 0x35, 0x5b, 0xcf, 0x37, 0x8b, 0xcc, 0xe9, 0xcc, 0xe3, 0x76, 0x5e,
	0xcf, 0x7f, 0xee, 0xd3, 0x2b, 0x31, 0x9f, 0xe8, 0xbf, 0xb4, 0x60, 0x42, 0x3b, 0x13, 0x02, 0x98,
	0xbb, 0x5a, 0xf7, 0x8c, 0x34, 0xfa, 0x4c, 0x94, 0xcd, 0xbd, 0xf8, 0xef, 0x11, 0xf2, 0x56, 0x44,
	0xc3, 0xc7, 0xe5, 0x72, 0xef, 0xf7, 0x37, 0x44, 0x72, 0x94, 0x8d, 0x49, 0xaf, 0x8d, 0xe3, 0xe7,
	0x1b, 0x2a, 0xd3, 0xe4, 0xfd, 0xf3, 0xeb, 0x23, 0x3c, 0x49, 0x17, 0x43, 0xb3, 0xea, 0x75, 0x88,
	0xf5, 0x76, 0x19, 0x5c, 0xe8, 0xc8, 0xf7, 0xb1, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xea, 0xd1,
	0x4d, 0x06, 0x7d, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TableClient is the client API for Table service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TableClient interface {
	Create(ctx context.Context, in *CreateTableRequest, opts ...grpc.CallOption) (*proto1.EmptyMsg, error)
}

type tableClient struct {
	cc *grpc.ClientConn
}

func NewTableClient(cc *grpc.ClientConn) TableClient {
	return &tableClient{cc}
}

func (c *tableClient) Create(ctx context.Context, in *CreateTableRequest, opts ...grpc.CallOption) (*proto1.EmptyMsg, error) {
	out := new(proto1.EmptyMsg)
	err := c.cc.Invoke(ctx, "/pb.Table/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TableServer is the server API for Table service.
type TableServer interface {
	Create(context.Context, *CreateTableRequest) (*proto1.EmptyMsg, error)
}

// UnimplementedTableServer can be embedded to have forward compatible implementations.
type UnimplementedTableServer struct {
}

func (*UnimplementedTableServer) Create(ctx context.Context, req *CreateTableRequest) (*proto1.EmptyMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

func RegisterTableServer(s *grpc.Server, srv TableServer) {
	s.RegisterService(&_Table_serviceDesc, srv)
}

func _Table_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TableServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Table/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TableServer).Create(ctx, req.(*CreateTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Table_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Table",
	HandlerType: (*TableServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Table_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "base/pb/table.proto",
}
