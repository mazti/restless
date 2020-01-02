// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/base/base.proto

package base

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

type CreateBaseRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBaseRequest) Reset()         { *m = CreateBaseRequest{} }
func (m *CreateBaseRequest) String() string { return proto.CompactTextString(m) }
func (*CreateBaseRequest) ProtoMessage()    {}
func (*CreateBaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_68c032de796ae9b1, []int{0}
}

func (m *CreateBaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBaseRequest.Unmarshal(m, b)
}
func (m *CreateBaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBaseRequest.Marshal(b, m, deterministic)
}
func (m *CreateBaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBaseRequest.Merge(m, src)
}
func (m *CreateBaseRequest) XXX_Size() int {
	return xxx_messageInfo_CreateBaseRequest.Size(m)
}
func (m *CreateBaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBaseRequest proto.InternalMessageInfo

func (m *CreateBaseRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateBaseReply struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBaseReply) Reset()         { *m = CreateBaseReply{} }
func (m *CreateBaseReply) String() string { return proto.CompactTextString(m) }
func (*CreateBaseReply) ProtoMessage()    {}
func (*CreateBaseReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_68c032de796ae9b1, []int{1}
}

func (m *CreateBaseReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBaseReply.Unmarshal(m, b)
}
func (m *CreateBaseReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBaseReply.Marshal(b, m, deterministic)
}
func (m *CreateBaseReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBaseReply.Merge(m, src)
}
func (m *CreateBaseReply) XXX_Size() int {
	return xxx_messageInfo_CreateBaseReply.Size(m)
}
func (m *CreateBaseReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBaseReply.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBaseReply proto.InternalMessageInfo

func (m *CreateBaseReply) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateBaseReply) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetBaseRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBaseRequest) Reset()         { *m = GetBaseRequest{} }
func (m *GetBaseRequest) String() string { return proto.CompactTextString(m) }
func (*GetBaseRequest) ProtoMessage()    {}
func (*GetBaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_68c032de796ae9b1, []int{2}
}

func (m *GetBaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBaseRequest.Unmarshal(m, b)
}
func (m *GetBaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBaseRequest.Marshal(b, m, deterministic)
}
func (m *GetBaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBaseRequest.Merge(m, src)
}
func (m *GetBaseRequest) XXX_Size() int {
	return xxx_messageInfo_GetBaseRequest.Size(m)
}
func (m *GetBaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBaseRequest proto.InternalMessageInfo

func (m *GetBaseRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetBaseReply struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBaseReply) Reset()         { *m = GetBaseReply{} }
func (m *GetBaseReply) String() string { return proto.CompactTextString(m) }
func (*GetBaseReply) ProtoMessage()    {}
func (*GetBaseReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_68c032de796ae9b1, []int{3}
}

func (m *GetBaseReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBaseReply.Unmarshal(m, b)
}
func (m *GetBaseReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBaseReply.Marshal(b, m, deterministic)
}
func (m *GetBaseReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBaseReply.Merge(m, src)
}
func (m *GetBaseReply) XXX_Size() int {
	return xxx_messageInfo_GetBaseReply.Size(m)
}
func (m *GetBaseReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBaseReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetBaseReply proto.InternalMessageInfo

func (m *GetBaseReply) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetBaseReply) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListBaseRequest struct {
	Limit                int32    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               int32    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBaseRequest) Reset()         { *m = ListBaseRequest{} }
func (m *ListBaseRequest) String() string { return proto.CompactTextString(m) }
func (*ListBaseRequest) ProtoMessage()    {}
func (*ListBaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_68c032de796ae9b1, []int{4}
}

func (m *ListBaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBaseRequest.Unmarshal(m, b)
}
func (m *ListBaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBaseRequest.Marshal(b, m, deterministic)
}
func (m *ListBaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBaseRequest.Merge(m, src)
}
func (m *ListBaseRequest) XXX_Size() int {
	return xxx_messageInfo_ListBaseRequest.Size(m)
}
func (m *ListBaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListBaseRequest proto.InternalMessageInfo

func (m *ListBaseRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListBaseRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type ListBaseReply struct {
	Bases                []*BaseModel  `protobuf:"bytes,1,rep,name=bases,proto3" json:"bases,omitempty"`
	Metadata             *ListMetadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ListBaseReply) Reset()         { *m = ListBaseReply{} }
func (m *ListBaseReply) String() string { return proto.CompactTextString(m) }
func (*ListBaseReply) ProtoMessage()    {}
func (*ListBaseReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_68c032de796ae9b1, []int{5}
}

func (m *ListBaseReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBaseReply.Unmarshal(m, b)
}
func (m *ListBaseReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBaseReply.Marshal(b, m, deterministic)
}
func (m *ListBaseReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBaseReply.Merge(m, src)
}
func (m *ListBaseReply) XXX_Size() int {
	return xxx_messageInfo_ListBaseReply.Size(m)
}
func (m *ListBaseReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBaseReply.DiscardUnknown(m)
}

var xxx_messageInfo_ListBaseReply proto.InternalMessageInfo

func (m *ListBaseReply) GetBases() []*BaseModel {
	if m != nil {
		return m.Bases
	}
	return nil
}

func (m *ListBaseReply) GetMetadata() *ListMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type UpdateBaseRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateBaseRequest) Reset()         { *m = UpdateBaseRequest{} }
func (m *UpdateBaseRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateBaseRequest) ProtoMessage()    {}
func (*UpdateBaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_68c032de796ae9b1, []int{6}
}

func (m *UpdateBaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateBaseRequest.Unmarshal(m, b)
}
func (m *UpdateBaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateBaseRequest.Marshal(b, m, deterministic)
}
func (m *UpdateBaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBaseRequest.Merge(m, src)
}
func (m *UpdateBaseRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateBaseRequest.Size(m)
}
func (m *UpdateBaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBaseRequest proto.InternalMessageInfo

func (m *UpdateBaseRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateBaseRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type UpdateBaseReply struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateBaseReply) Reset()         { *m = UpdateBaseReply{} }
func (m *UpdateBaseReply) String() string { return proto.CompactTextString(m) }
func (*UpdateBaseReply) ProtoMessage()    {}
func (*UpdateBaseReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_68c032de796ae9b1, []int{7}
}

func (m *UpdateBaseReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateBaseReply.Unmarshal(m, b)
}
func (m *UpdateBaseReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateBaseReply.Marshal(b, m, deterministic)
}
func (m *UpdateBaseReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBaseReply.Merge(m, src)
}
func (m *UpdateBaseReply) XXX_Size() int {
	return xxx_messageInfo_UpdateBaseReply.Size(m)
}
func (m *UpdateBaseReply) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBaseReply.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBaseReply proto.InternalMessageInfo

func (m *UpdateBaseReply) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateBaseReply) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DeleteBaseRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteBaseRequest) Reset()         { *m = DeleteBaseRequest{} }
func (m *DeleteBaseRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteBaseRequest) ProtoMessage()    {}
func (*DeleteBaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_68c032de796ae9b1, []int{8}
}

func (m *DeleteBaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBaseRequest.Unmarshal(m, b)
}
func (m *DeleteBaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBaseRequest.Marshal(b, m, deterministic)
}
func (m *DeleteBaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBaseRequest.Merge(m, src)
}
func (m *DeleteBaseRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteBaseRequest.Size(m)
}
func (m *DeleteBaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBaseRequest proto.InternalMessageInfo

func (m *DeleteBaseRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type EmptyMsg struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyMsg) Reset()         { *m = EmptyMsg{} }
func (m *EmptyMsg) String() string { return proto.CompactTextString(m) }
func (*EmptyMsg) ProtoMessage()    {}
func (*EmptyMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_68c032de796ae9b1, []int{9}
}

func (m *EmptyMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyMsg.Unmarshal(m, b)
}
func (m *EmptyMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyMsg.Marshal(b, m, deterministic)
}
func (m *EmptyMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyMsg.Merge(m, src)
}
func (m *EmptyMsg) XXX_Size() int {
	return xxx_messageInfo_EmptyMsg.Size(m)
}
func (m *EmptyMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyMsg.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyMsg proto.InternalMessageInfo

type BaseModel struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BaseModel) Reset()         { *m = BaseModel{} }
func (m *BaseModel) String() string { return proto.CompactTextString(m) }
func (*BaseModel) ProtoMessage()    {}
func (*BaseModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_68c032de796ae9b1, []int{10}
}

func (m *BaseModel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseModel.Unmarshal(m, b)
}
func (m *BaseModel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseModel.Marshal(b, m, deterministic)
}
func (m *BaseModel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseModel.Merge(m, src)
}
func (m *BaseModel) XXX_Size() int {
	return xxx_messageInfo_BaseModel.Size(m)
}
func (m *BaseModel) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseModel.DiscardUnknown(m)
}

var xxx_messageInfo_BaseModel proto.InternalMessageInfo

func (m *BaseModel) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *BaseModel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListMetadata struct {
	Count                int32    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Limit                int32    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               int32    `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Total                int64    `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListMetadata) Reset()         { *m = ListMetadata{} }
func (m *ListMetadata) String() string { return proto.CompactTextString(m) }
func (*ListMetadata) ProtoMessage()    {}
func (*ListMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_68c032de796ae9b1, []int{11}
}

func (m *ListMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMetadata.Unmarshal(m, b)
}
func (m *ListMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMetadata.Marshal(b, m, deterministic)
}
func (m *ListMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMetadata.Merge(m, src)
}
func (m *ListMetadata) XXX_Size() int {
	return xxx_messageInfo_ListMetadata.Size(m)
}
func (m *ListMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ListMetadata proto.InternalMessageInfo

func (m *ListMetadata) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ListMetadata) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListMetadata) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ListMetadata) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

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
	return fileDescriptor_68c032de796ae9b1, []int{12}
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
	return fileDescriptor_68c032de796ae9b1, []int{13}
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
	proto.RegisterType((*CreateBaseRequest)(nil), "base.CreateBaseRequest")
	proto.RegisterType((*CreateBaseReply)(nil), "base.CreateBaseReply")
	proto.RegisterType((*GetBaseRequest)(nil), "base.GetBaseRequest")
	proto.RegisterType((*GetBaseReply)(nil), "base.GetBaseReply")
	proto.RegisterType((*ListBaseRequest)(nil), "base.ListBaseRequest")
	proto.RegisterType((*ListBaseReply)(nil), "base.ListBaseReply")
	proto.RegisterType((*UpdateBaseRequest)(nil), "base.UpdateBaseRequest")
	proto.RegisterType((*UpdateBaseReply)(nil), "base.UpdateBaseReply")
	proto.RegisterType((*DeleteBaseRequest)(nil), "base.DeleteBaseRequest")
	proto.RegisterType((*EmptyMsg)(nil), "base.EmptyMsg")
	proto.RegisterType((*BaseModel)(nil), "base.BaseModel")
	proto.RegisterType((*ListMetadata)(nil), "base.ListMetadata")
	proto.RegisterType((*CreateTableRequest)(nil), "base.CreateTableRequest")
	proto.RegisterType((*TableColumn)(nil), "base.TableColumn")
}

func init() { proto.RegisterFile("pb/base/base.proto", fileDescriptor_68c032de796ae9b1) }

var fileDescriptor_68c032de796ae9b1 = []byte{
	// 573 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xd1, 0x8a, 0xd3, 0x4c,
	0x14, 0xa6, 0x69, 0x93, 0xdd, 0x3d, 0xdb, 0x6d, 0xfe, 0x9e, 0x7f, 0xbb, 0x96, 0x22, 0x52, 0x46,
	0xc4, 0xb2, 0x42, 0x03, 0x91, 0x45, 0xf0, 0x46, 0xd6, 0xba, 0x2c, 0x88, 0x55, 0x08, 0x8a, 0xd7,
	0xd3, 0x66, 0xda, 0x0d, 0x24, 0x99, 0xd8, 0x99, 0x5e, 0x94, 0xc5, 0x1b, 0x5f, 0xc1, 0x47, 0xf3,
	0x15, 0x7c, 0x02, 0x9f, 0x40, 0x66, 0x26, 0x4d, 0x67, 0xdb, 0x22, 0xbd, 0x29, 0x9d, 0x39, 0xdf,
	0xf9, 0xce, 0x39, 0xdf, 0xf9, 0x26, 0x80, 0xc5, 0x24, 0x98, 0x50, 0xc1, 0xf4, 0xcf, 0xb0, 0x58,
	0x70, 0xc9, 0xb1, 0xa1, 0xfe, 0xf7, 0x1e, 0xcf, 0x39, 0x9f, 0xa7, 0x2c, 0xa0, 0x45, 0x12, 0xd0,
	0x3c, 0xe7, 0x92, 0xca, 0x84, 0xe7, 0xc2, 0x60, 0xc8, 0x73, 0x68, 0x8f, 0x16, 0x8c, 0x4a, 0xf6,
	0x96, 0x0a, 0x16, 0xb1, 0x6f, 0x4b, 0x26, 0x24, 0x22, 0x34, 0x72, 0x9a, 0xb1, 0x6e, 0xad, 0x5f,
	0x1b, 0x9c, 0x44, 0xfa, 0x3f, 0xb9, 0x02, 0xdf, 0x06, 0x16, 0xe9, 0x0a, 0x5b, 0xe0, 0x24, 0x71,
	0x09, 0x72, 0x92, 0xb8, 0x4a, 0x73, 0xac, 0xb4, 0x3e, 0xb4, 0x6e, 0x99, 0xb4, 0xc9, 0xb7, 0xb2,
	0x48, 0x08, 0xcd, 0x0a, 0x71, 0x28, 0xeb, 0x1b, 0xf0, 0x3f, 0x24, 0xe2, 0x01, 0xed, 0x39, 0xb8,
	0x69, 0x92, 0x25, 0x52, 0x67, 0xba, 0x91, 0x39, 0xe0, 0x05, 0x78, 0x7c, 0x36, 0x13, 0x4c, 0xea,
	0x74, 0x37, 0x2a, 0x4f, 0x64, 0x06, 0x67, 0x1b, 0x02, 0x55, 0xf5, 0x19, 0xb8, 0x4a, 0x2d, 0xd1,
	0xad, 0xf5, 0xeb, 0x83, 0xd3, 0xd0, 0x1f, 0x6a, 0x1d, 0x55, 0x7c, 0xcc, 0x63, 0x96, 0x46, 0x26,
	0x8a, 0x43, 0x38, 0xce, 0x98, 0xa4, 0x31, 0x95, 0x54, 0x33, 0x9e, 0x86, 0x68, 0x90, 0x8a, 0x6d,
	0x5c, 0x46, 0xa2, 0x0a, 0x43, 0x5e, 0x41, 0xfb, 0x4b, 0x11, 0x6f, 0xc9, 0x7b, 0xc8, 0x84, 0x57,
	0xe0, 0xdb, 0x89, 0x87, 0x0a, 0xf3, 0x14, 0xda, 0xef, 0x58, 0xca, 0xfe, 0x59, 0x8f, 0x00, 0x1c,
	0xdf, 0x64, 0x85, 0x5c, 0x8d, 0xc5, 0x9c, 0x04, 0x70, 0x52, 0x0d, 0x79, 0x50, 0x85, 0x3b, 0x68,
	0xda, 0xb3, 0x2a, 0xdd, 0xa7, 0x7c, 0x99, 0x57, 0xba, 0xeb, 0xc3, 0x66, 0x1b, 0xce, 0xfe, 0x6d,
	0xd4, 0xed, 0x6d, 0x28, 0xb4, 0xe4, 0x92, 0xa6, 0xdd, 0x46, 0xbf, 0x36, 0xa8, 0x47, 0xe6, 0x40,
	0x32, 0x40, 0xe3, 0xb8, 0xcf, 0x74, 0x92, 0x56, 0xc3, 0x5c, 0x80, 0x27, 0xa6, 0x77, 0x2c, 0xa3,
	0x65, 0x9f, 0xe5, 0x69, 0x5f, 0xaf, 0xf8, 0x02, 0x8e, 0xa6, 0x3c, 0x5d, 0x66, 0xb9, 0xe8, 0xd6,
	0xf5, 0x5a, 0xdb, 0x66, 0x59, 0x9a, 0x70, 0xa4, 0x23, 0xd1, 0x1a, 0x41, 0xae, 0xe1, 0xd4, 0xba,
	0xdf, 0xf7, 0x06, 0xf0, 0x09, 0x00, 0x95, 0x72, 0x91, 0x4c, 0x96, 0x92, 0x89, 0xb2, 0x92, 0x75,
	0x13, 0xfe, 0x71, 0xa0, 0xa1, 0xd4, 0xc4, 0xf7, 0xe0, 0x99, 0xd6, 0xf1, 0x91, 0xa9, 0xb8, 0xf3,
	0xc6, 0x7a, 0x9d, 0xdd, 0x40, 0x91, 0xae, 0xc8, 0x7f, 0x3f, 0x7e, 0xfd, 0xfe, 0xe9, 0x00, 0x71,
	0xf5, 0x43, 0x7e, 0x5d, 0xbb, 0xc4, 0x11, 0xd4, 0x6f, 0x95, 0x46, 0x06, 0xff, 0xf0, 0x31, 0xf5,
	0x70, 0xeb, 0x56, 0x51, 0xa0, 0xa6, 0x68, 0x22, 0x98, 0x0f, 0xc2, 0x7d, 0x12, 0x7f, 0xc7, 0x6b,
	0x68, 0xa8, 0xad, 0x61, 0x67, 0xe3, 0x56, 0x9b, 0xe6, 0xff, 0xed, 0x6b, 0xc5, 0x73, 0xa6, 0x79,
	0x8e, 0xd0, 0xb4, 0x82, 0x9f, 0xc0, 0x33, 0x8e, 0x5c, 0xcf, 0xb4, 0x63, 0xec, 0xf5, 0x4c, 0x5b,
	0xc6, 0x25, 0x1d, 0x4d, 0xe4, 0x87, 0x56, 0x43, 0x6a, 0xb0, 0x1b, 0xf0, 0x8c, 0x57, 0xd7, 0x84,
	0x3b, 0xce, 0xed, 0xb5, 0x4c, 0xa0, 0x72, 0x6b, 0x39, 0xda, 0xa5, 0xc5, 0x14, 0x7e, 0x05, 0x57,
	0xef, 0x0d, 0x3f, 0x56, 0xa2, 0x77, 0x6d, 0x6d, 0x6d, 0xf7, 0xec, 0x10, 0xf6, 0x34, 0xe1, 0x39,
	0xf1, 0x03, 0xa9, 0x60, 0xc1, 0xbd, 0x71, 0x93, 0xea, 0x6f, 0xe2, 0xe9, 0x2f, 0xe4, 0xcb, 0xbf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xf7, 0x75, 0xbd, 0x21, 0x5b, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BaseClient is the client API for Base service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BaseClient interface {
	Create(ctx context.Context, in *CreateBaseRequest, opts ...grpc.CallOption) (*CreateBaseReply, error)
	Get(ctx context.Context, in *GetBaseRequest, opts ...grpc.CallOption) (*GetBaseReply, error)
	List(ctx context.Context, in *ListBaseRequest, opts ...grpc.CallOption) (*ListBaseReply, error)
	Update(ctx context.Context, in *UpdateBaseRequest, opts ...grpc.CallOption) (*UpdateBaseReply, error)
	Delete(ctx context.Context, in *DeleteBaseRequest, opts ...grpc.CallOption) (*EmptyMsg, error)
}

type baseClient struct {
	cc *grpc.ClientConn
}

func NewBaseClient(cc *grpc.ClientConn) BaseClient {
	return &baseClient{cc}
}

func (c *baseClient) Create(ctx context.Context, in *CreateBaseRequest, opts ...grpc.CallOption) (*CreateBaseReply, error) {
	out := new(CreateBaseReply)
	err := c.cc.Invoke(ctx, "/base.Base/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseClient) Get(ctx context.Context, in *GetBaseRequest, opts ...grpc.CallOption) (*GetBaseReply, error) {
	out := new(GetBaseReply)
	err := c.cc.Invoke(ctx, "/base.Base/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseClient) List(ctx context.Context, in *ListBaseRequest, opts ...grpc.CallOption) (*ListBaseReply, error) {
	out := new(ListBaseReply)
	err := c.cc.Invoke(ctx, "/base.Base/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseClient) Update(ctx context.Context, in *UpdateBaseRequest, opts ...grpc.CallOption) (*UpdateBaseReply, error) {
	out := new(UpdateBaseReply)
	err := c.cc.Invoke(ctx, "/base.Base/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseClient) Delete(ctx context.Context, in *DeleteBaseRequest, opts ...grpc.CallOption) (*EmptyMsg, error) {
	out := new(EmptyMsg)
	err := c.cc.Invoke(ctx, "/base.Base/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BaseServer is the server API for Base service.
type BaseServer interface {
	Create(context.Context, *CreateBaseRequest) (*CreateBaseReply, error)
	Get(context.Context, *GetBaseRequest) (*GetBaseReply, error)
	List(context.Context, *ListBaseRequest) (*ListBaseReply, error)
	Update(context.Context, *UpdateBaseRequest) (*UpdateBaseReply, error)
	Delete(context.Context, *DeleteBaseRequest) (*EmptyMsg, error)
}

// UnimplementedBaseServer can be embedded to have forward compatible implementations.
type UnimplementedBaseServer struct {
}

func (*UnimplementedBaseServer) Create(ctx context.Context, req *CreateBaseRequest) (*CreateBaseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedBaseServer) Get(ctx context.Context, req *GetBaseRequest) (*GetBaseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedBaseServer) List(ctx context.Context, req *ListBaseRequest) (*ListBaseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedBaseServer) Update(ctx context.Context, req *UpdateBaseRequest) (*UpdateBaseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedBaseServer) Delete(ctx context.Context, req *DeleteBaseRequest) (*EmptyMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterBaseServer(s *grpc.Server, srv BaseServer) {
	s.RegisterService(&_Base_serviceDesc, srv)
}

func _Base_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/base.Base/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseServer).Create(ctx, req.(*CreateBaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Base_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/base.Base/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseServer).Get(ctx, req.(*GetBaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Base_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/base.Base/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseServer).List(ctx, req.(*ListBaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Base_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/base.Base/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseServer).Update(ctx, req.(*UpdateBaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Base_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/base.Base/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseServer).Delete(ctx, req.(*DeleteBaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Base_serviceDesc = grpc.ServiceDesc{
	ServiceName: "base.Base",
	HandlerType: (*BaseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Base_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Base_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Base_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Base_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Base_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/base/base.proto",
}

// TableClient is the client API for Table service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TableClient interface {
	Create(ctx context.Context, in *CreateTableRequest, opts ...grpc.CallOption) (*EmptyMsg, error)
}

type tableClient struct {
	cc *grpc.ClientConn
}

func NewTableClient(cc *grpc.ClientConn) TableClient {
	return &tableClient{cc}
}

func (c *tableClient) Create(ctx context.Context, in *CreateTableRequest, opts ...grpc.CallOption) (*EmptyMsg, error) {
	out := new(EmptyMsg)
	err := c.cc.Invoke(ctx, "/base.Table/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TableServer is the server API for Table service.
type TableServer interface {
	Create(context.Context, *CreateTableRequest) (*EmptyMsg, error)
}

// UnimplementedTableServer can be embedded to have forward compatible implementations.
type UnimplementedTableServer struct {
}

func (*UnimplementedTableServer) Create(ctx context.Context, req *CreateTableRequest) (*EmptyMsg, error) {
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
		FullMethod: "/base.Table/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TableServer).Create(ctx, req.(*CreateTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Table_serviceDesc = grpc.ServiceDesc{
	ServiceName: "base.Table",
	HandlerType: (*TableServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Table_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/base/base.proto",
}
