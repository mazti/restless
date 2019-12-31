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

type CreateMetaRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateMetaRequest) Reset()         { *m = CreateMetaRequest{} }
func (m *CreateMetaRequest) String() string { return proto.CompactTextString(m) }
func (*CreateMetaRequest) ProtoMessage()    {}
func (*CreateMetaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdd022eb384fa38d, []int{0}
}

func (m *CreateMetaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMetaRequest.Unmarshal(m, b)
}
func (m *CreateMetaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMetaRequest.Marshal(b, m, deterministic)
}
func (m *CreateMetaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMetaRequest.Merge(m, src)
}
func (m *CreateMetaRequest) XXX_Size() int {
	return xxx_messageInfo_CreateMetaRequest.Size(m)
}
func (m *CreateMetaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMetaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMetaRequest proto.InternalMessageInfo

func (m *CreateMetaRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateMetaReply struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateMetaReply) Reset()         { *m = CreateMetaReply{} }
func (m *CreateMetaReply) String() string { return proto.CompactTextString(m) }
func (*CreateMetaReply) ProtoMessage()    {}
func (*CreateMetaReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdd022eb384fa38d, []int{1}
}

func (m *CreateMetaReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMetaReply.Unmarshal(m, b)
}
func (m *CreateMetaReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMetaReply.Marshal(b, m, deterministic)
}
func (m *CreateMetaReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMetaReply.Merge(m, src)
}
func (m *CreateMetaReply) XXX_Size() int {
	return xxx_messageInfo_CreateMetaReply.Size(m)
}
func (m *CreateMetaReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMetaReply.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMetaReply proto.InternalMessageInfo

func (m *CreateMetaReply) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateMetaReply) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetMetaRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMetaRequest) Reset()         { *m = GetMetaRequest{} }
func (m *GetMetaRequest) String() string { return proto.CompactTextString(m) }
func (*GetMetaRequest) ProtoMessage()    {}
func (*GetMetaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdd022eb384fa38d, []int{2}
}

func (m *GetMetaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMetaRequest.Unmarshal(m, b)
}
func (m *GetMetaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMetaRequest.Marshal(b, m, deterministic)
}
func (m *GetMetaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMetaRequest.Merge(m, src)
}
func (m *GetMetaRequest) XXX_Size() int {
	return xxx_messageInfo_GetMetaRequest.Size(m)
}
func (m *GetMetaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMetaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMetaRequest proto.InternalMessageInfo

func (m *GetMetaRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetMetaReply struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMetaReply) Reset()         { *m = GetMetaReply{} }
func (m *GetMetaReply) String() string { return proto.CompactTextString(m) }
func (*GetMetaReply) ProtoMessage()    {}
func (*GetMetaReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdd022eb384fa38d, []int{3}
}

func (m *GetMetaReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMetaReply.Unmarshal(m, b)
}
func (m *GetMetaReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMetaReply.Marshal(b, m, deterministic)
}
func (m *GetMetaReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMetaReply.Merge(m, src)
}
func (m *GetMetaReply) XXX_Size() int {
	return xxx_messageInfo_GetMetaReply.Size(m)
}
func (m *GetMetaReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMetaReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetMetaReply proto.InternalMessageInfo

func (m *GetMetaReply) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetMetaReply) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListMetaRequest struct {
	Limit                int32    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               int32    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListMetaRequest) Reset()         { *m = ListMetaRequest{} }
func (m *ListMetaRequest) String() string { return proto.CompactTextString(m) }
func (*ListMetaRequest) ProtoMessage()    {}
func (*ListMetaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdd022eb384fa38d, []int{4}
}

func (m *ListMetaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMetaRequest.Unmarshal(m, b)
}
func (m *ListMetaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMetaRequest.Marshal(b, m, deterministic)
}
func (m *ListMetaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMetaRequest.Merge(m, src)
}
func (m *ListMetaRequest) XXX_Size() int {
	return xxx_messageInfo_ListMetaRequest.Size(m)
}
func (m *ListMetaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMetaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListMetaRequest proto.InternalMessageInfo

func (m *ListMetaRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListMetaRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type ListMetaReply struct {
	Metas                []*MetaModel `protobuf:"bytes,1,rep,name=metas,proto3" json:"metas,omitempty"`
	Metadata             *Metadata    `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ListMetaReply) Reset()         { *m = ListMetaReply{} }
func (m *ListMetaReply) String() string { return proto.CompactTextString(m) }
func (*ListMetaReply) ProtoMessage()    {}
func (*ListMetaReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdd022eb384fa38d, []int{5}
}

func (m *ListMetaReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMetaReply.Unmarshal(m, b)
}
func (m *ListMetaReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMetaReply.Marshal(b, m, deterministic)
}
func (m *ListMetaReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMetaReply.Merge(m, src)
}
func (m *ListMetaReply) XXX_Size() int {
	return xxx_messageInfo_ListMetaReply.Size(m)
}
func (m *ListMetaReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMetaReply.DiscardUnknown(m)
}

var xxx_messageInfo_ListMetaReply proto.InternalMessageInfo

func (m *ListMetaReply) GetMetas() []*MetaModel {
	if m != nil {
		return m.Metas
	}
	return nil
}

func (m *ListMetaReply) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type UpdateMetaRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateMetaRequest) Reset()         { *m = UpdateMetaRequest{} }
func (m *UpdateMetaRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateMetaRequest) ProtoMessage()    {}
func (*UpdateMetaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdd022eb384fa38d, []int{6}
}

func (m *UpdateMetaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateMetaRequest.Unmarshal(m, b)
}
func (m *UpdateMetaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateMetaRequest.Marshal(b, m, deterministic)
}
func (m *UpdateMetaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateMetaRequest.Merge(m, src)
}
func (m *UpdateMetaRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateMetaRequest.Size(m)
}
func (m *UpdateMetaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateMetaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateMetaRequest proto.InternalMessageInfo

func (m *UpdateMetaRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateMetaRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type UpdateMetaReply struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateMetaReply) Reset()         { *m = UpdateMetaReply{} }
func (m *UpdateMetaReply) String() string { return proto.CompactTextString(m) }
func (*UpdateMetaReply) ProtoMessage()    {}
func (*UpdateMetaReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdd022eb384fa38d, []int{7}
}

func (m *UpdateMetaReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateMetaReply.Unmarshal(m, b)
}
func (m *UpdateMetaReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateMetaReply.Marshal(b, m, deterministic)
}
func (m *UpdateMetaReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateMetaReply.Merge(m, src)
}
func (m *UpdateMetaReply) XXX_Size() int {
	return xxx_messageInfo_UpdateMetaReply.Size(m)
}
func (m *UpdateMetaReply) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateMetaReply.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateMetaReply proto.InternalMessageInfo

func (m *UpdateMetaReply) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateMetaReply) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DeleteMetaRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteMetaRequest) Reset()         { *m = DeleteMetaRequest{} }
func (m *DeleteMetaRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteMetaRequest) ProtoMessage()    {}
func (*DeleteMetaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdd022eb384fa38d, []int{8}
}

func (m *DeleteMetaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteMetaRequest.Unmarshal(m, b)
}
func (m *DeleteMetaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteMetaRequest.Marshal(b, m, deterministic)
}
func (m *DeleteMetaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteMetaRequest.Merge(m, src)
}
func (m *DeleteMetaRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteMetaRequest.Size(m)
}
func (m *DeleteMetaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteMetaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteMetaRequest proto.InternalMessageInfo

func (m *DeleteMetaRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteMetaReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteMetaReply) Reset()         { *m = DeleteMetaReply{} }
func (m *DeleteMetaReply) String() string { return proto.CompactTextString(m) }
func (*DeleteMetaReply) ProtoMessage()    {}
func (*DeleteMetaReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdd022eb384fa38d, []int{9}
}

func (m *DeleteMetaReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteMetaReply.Unmarshal(m, b)
}
func (m *DeleteMetaReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteMetaReply.Marshal(b, m, deterministic)
}
func (m *DeleteMetaReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteMetaReply.Merge(m, src)
}
func (m *DeleteMetaReply) XXX_Size() int {
	return xxx_messageInfo_DeleteMetaReply.Size(m)
}
func (m *DeleteMetaReply) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteMetaReply.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteMetaReply proto.InternalMessageInfo

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
	return fileDescriptor_bdd022eb384fa38d, []int{10}
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

type Metadata struct {
	Count                int32    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Limit                int32    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               int32    `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Total                int64    `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdd022eb384fa38d, []int{11}
}

func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *Metadata) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *Metadata) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *Metadata) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateMetaRequest)(nil), "pb.CreateMetaRequest")
	proto.RegisterType((*CreateMetaReply)(nil), "pb.CreateMetaReply")
	proto.RegisterType((*GetMetaRequest)(nil), "pb.GetMetaRequest")
	proto.RegisterType((*GetMetaReply)(nil), "pb.GetMetaReply")
	proto.RegisterType((*ListMetaRequest)(nil), "pb.ListMetaRequest")
	proto.RegisterType((*ListMetaReply)(nil), "pb.ListMetaReply")
	proto.RegisterType((*UpdateMetaRequest)(nil), "pb.UpdateMetaRequest")
	proto.RegisterType((*UpdateMetaReply)(nil), "pb.UpdateMetaReply")
	proto.RegisterType((*DeleteMetaRequest)(nil), "pb.DeleteMetaRequest")
	proto.RegisterType((*DeleteMetaReply)(nil), "pb.DeleteMetaReply")
	proto.RegisterType((*MetaModel)(nil), "pb.MetaModel")
	proto.RegisterType((*Metadata)(nil), "pb.Metadata")
}

func init() { proto.RegisterFile("pb/meta.proto", fileDescriptor_bdd022eb384fa38d) }

var fileDescriptor_bdd022eb384fa38d = []byte{
	// 462 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x41, 0x6b, 0xd4, 0x40,
	0x14, 0x26, 0xd9, 0x4d, 0x6c, 0x5f, 0x77, 0x9b, 0xe6, 0x69, 0x65, 0x09, 0x1e, 0xc2, 0xf4, 0x60,
	0xe8, 0x61, 0x03, 0x91, 0x22, 0x08, 0xd2, 0x83, 0x62, 0x05, 0xed, 0x25, 0xe0, 0x55, 0x98, 0x98,
	0x69, 0x19, 0xc8, 0x66, 0xc6, 0x66, 0x7a, 0x28, 0xe2, 0xc5, 0xbf, 0xe0, 0x4f, 0xf3, 0x2f, 0xe8,
	0xff, 0x90, 0x99, 0xd9, 0xa6, 0x93, 0xac, 0x48, 0x6e, 0x99, 0x37, 0xdf, 0xf7, 0xbd, 0xf9, 0xde,
	0xf7, 0x76, 0x61, 0x29, 0xab, 0x7c, 0xc3, 0x14, 0x5d, 0xcb, 0x1b, 0xa1, 0x04, 0xfa, 0xb2, 0x4a,
	0x9e, 0x5d, 0x0b, 0x71, 0xdd, 0xb0, 0x9c, 0x4a, 0x9e, 0xd3, 0xb6, 0x15, 0x8a, 0x2a, 0x2e, 0xda,
	0xce, 0x22, 0xc8, 0x73, 0x88, 0xdf, 0xdc, 0x30, 0xaa, 0xd8, 0x25, 0x53, 0xb4, 0x64, 0x5f, 0x6f,
	0x59, 0xa7, 0x10, 0x61, 0xde, 0xd2, 0x0d, 0x5b, 0x79, 0xa9, 0x97, 0xed, 0x97, 0xe6, 0x9b, 0x9c,
	0x41, 0xe4, 0x02, 0x65, 0x73, 0x87, 0x87, 0xe0, 0xf3, 0x7a, 0x0b, 0xf2, 0x79, 0xdd, 0xd3, 0x7c,
	0x87, 0x96, 0xc2, 0xe1, 0x05, 0x53, 0xae, 0xf8, 0x88, 0x45, 0x0a, 0x58, 0xf4, 0x88, 0xa9, 0xaa,
	0xe7, 0x10, 0x7d, 0xe4, 0xdd, 0x40, 0xf6, 0x09, 0x04, 0x0d, 0xdf, 0x70, 0x65, 0x98, 0x41, 0x69,
	0x0f, 0xf8, 0x14, 0x42, 0x71, 0x75, 0xd5, 0x31, 0x65, 0xe8, 0x41, 0xb9, 0x3d, 0x91, 0xcf, 0xb0,
	0x7c, 0x10, 0xd0, 0x5d, 0x4f, 0x20, 0xd0, 0x73, 0xeb, 0x56, 0x5e, 0x3a, 0xcb, 0x0e, 0x8a, 0xe5,
	0x5a, 0x56, 0x6b, 0x7d, 0x7b, 0x29, 0x6a, 0xd6, 0x94, 0xf6, 0x0e, 0x33, 0xd8, 0xd3, 0x1f, 0x35,
	0x55, 0xd4, 0xe8, 0x1d, 0x14, 0x8b, 0x7b, 0x9c, 0xae, 0x95, 0xfd, 0x2d, 0x79, 0x09, 0xf1, 0x27,
	0x59, 0x8f, 0xc6, 0x3a, 0xc5, 0xd9, 0x19, 0x44, 0x2e, 0x71, 0xea, 0x40, 0x4e, 0x20, 0x7e, 0xcb,
	0x1a, 0xf6, 0xdf, 0x7e, 0x24, 0x86, 0xc8, 0x05, 0xc9, 0xe6, 0x8e, 0xe4, 0xb0, 0xdf, 0xbb, 0x9c,
	0xd4, 0xa8, 0x86, 0xbd, 0x7b, 0xbb, 0x7a, 0xe4, 0x5f, 0xc4, 0x6d, 0xdb, 0x8f, 0xdc, 0x1c, 0x1e,
	0x82, 0xf0, 0xff, 0x1d, 0xc4, 0xcc, 0x0d, 0x42, 0xa3, 0x95, 0x50, 0xb4, 0x59, 0xcd, 0x53, 0x2f,
	0x9b, 0x95, 0xf6, 0x50, 0xfc, 0xf1, 0x61, 0xae, 0xdb, 0xe0, 0x3b, 0x08, 0xed, 0xd6, 0xe1, 0xb1,
	0x9e, 0xf4, 0xce, 0xaa, 0x26, 0x8f, 0xc7, 0x65, 0xed, 0xea, 0xe8, 0xc7, 0xaf, 0xdf, 0x3f, 0x7d,
	0x20, 0x81, 0xf9, 0x2d, 0xbc, 0xf2, 0x4e, 0xf1, 0x1c, 0x66, 0x17, 0x4c, 0x21, 0x6a, 0xf4, 0x70,
	0x1f, 0x93, 0xa3, 0x41, 0x4d, 0xd3, 0xd1, 0xd0, 0x17, 0x08, 0x86, 0x9e, 0x7f, 0xe3, 0xf5, 0x77,
	0x7c, 0x0d, 0x73, 0xbd, 0x30, 0x68, 0xfa, 0x8d, 0x76, 0x2f, 0x89, 0x87, 0x45, 0xad, 0xb1, 0x34,
	0x1a, 0x8f, 0xd0, 0x3e, 0x01, 0x3f, 0x40, 0x68, 0x63, 0xb5, 0x3e, 0x76, 0x76, 0xc3, 0xfa, 0x18,
	0x25, 0x4f, 0x8e, 0x8d, 0x48, 0x54, 0x38, 0x0f, 0xd1, 0x66, 0xde, 0x43, 0x68, 0x73, 0xb4, 0x62,
	0x3b, 0xc1, 0x5b, 0xb1, 0x71, 0xd4, 0x5b, 0x57, 0xa7, 0x8e, 0x58, 0x15, 0x9a, 0x3f, 0x81, 0x17,
	0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x38, 0xfa, 0xe4, 0x63, 0x37, 0x04, 0x00, 0x00,
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
	Create(ctx context.Context, in *CreateMetaRequest, opts ...grpc.CallOption) (*CreateMetaReply, error)
	Get(ctx context.Context, in *GetMetaRequest, opts ...grpc.CallOption) (*GetMetaReply, error)
	List(ctx context.Context, in *ListMetaRequest, opts ...grpc.CallOption) (*ListMetaReply, error)
	Update(ctx context.Context, in *UpdateMetaRequest, opts ...grpc.CallOption) (*UpdateMetaReply, error)
	Delete(ctx context.Context, in *DeleteMetaRequest, opts ...grpc.CallOption) (*DeleteMetaReply, error)
}

type metaClient struct {
	cc *grpc.ClientConn
}

func NewMetaClient(cc *grpc.ClientConn) MetaClient {
	return &metaClient{cc}
}

func (c *metaClient) Create(ctx context.Context, in *CreateMetaRequest, opts ...grpc.CallOption) (*CreateMetaReply, error) {
	out := new(CreateMetaReply)
	err := c.cc.Invoke(ctx, "/pb.Meta/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaClient) Get(ctx context.Context, in *GetMetaRequest, opts ...grpc.CallOption) (*GetMetaReply, error) {
	out := new(GetMetaReply)
	err := c.cc.Invoke(ctx, "/pb.Meta/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaClient) List(ctx context.Context, in *ListMetaRequest, opts ...grpc.CallOption) (*ListMetaReply, error) {
	out := new(ListMetaReply)
	err := c.cc.Invoke(ctx, "/pb.Meta/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaClient) Update(ctx context.Context, in *UpdateMetaRequest, opts ...grpc.CallOption) (*UpdateMetaReply, error) {
	out := new(UpdateMetaReply)
	err := c.cc.Invoke(ctx, "/pb.Meta/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaClient) Delete(ctx context.Context, in *DeleteMetaRequest, opts ...grpc.CallOption) (*DeleteMetaReply, error) {
	out := new(DeleteMetaReply)
	err := c.cc.Invoke(ctx, "/pb.Meta/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetaServer is the server API for Meta service.
type MetaServer interface {
	Create(context.Context, *CreateMetaRequest) (*CreateMetaReply, error)
	Get(context.Context, *GetMetaRequest) (*GetMetaReply, error)
	List(context.Context, *ListMetaRequest) (*ListMetaReply, error)
	Update(context.Context, *UpdateMetaRequest) (*UpdateMetaReply, error)
	Delete(context.Context, *DeleteMetaRequest) (*DeleteMetaReply, error)
}

// UnimplementedMetaServer can be embedded to have forward compatible implementations.
type UnimplementedMetaServer struct {
}

func (*UnimplementedMetaServer) Create(ctx context.Context, req *CreateMetaRequest) (*CreateMetaReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedMetaServer) Get(ctx context.Context, req *GetMetaRequest) (*GetMetaReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedMetaServer) List(ctx context.Context, req *ListMetaRequest) (*ListMetaReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedMetaServer) Update(ctx context.Context, req *UpdateMetaRequest) (*UpdateMetaReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedMetaServer) Delete(ctx context.Context, req *DeleteMetaRequest) (*DeleteMetaReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterMetaServer(s *grpc.Server, srv MetaServer) {
	s.RegisterService(&_Meta_serviceDesc, srv)
}

func _Meta_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMetaRequest)
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
		return srv.(MetaServer).Create(ctx, req.(*CreateMetaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meta_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMetaRequest)
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
		return srv.(MetaServer).Get(ctx, req.(*GetMetaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meta_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMetaRequest)
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
		return srv.(MetaServer).List(ctx, req.(*ListMetaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meta_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMetaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Meta/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaServer).Update(ctx, req.(*UpdateMetaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meta_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMetaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Meta/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaServer).Delete(ctx, req.(*DeleteMetaRequest))
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
		{
			MethodName: "Update",
			Handler:    _Meta_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Meta_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/meta.proto",
}
