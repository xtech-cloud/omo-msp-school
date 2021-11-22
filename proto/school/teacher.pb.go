// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/school/teacher.proto

package omo_msp_school

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type TeacherInfo struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Entity               string   `protobuf:"bytes,4,opt,name=entity,proto3" json:"entity,omitempty"`
	User                 string   `protobuf:"bytes,5,opt,name=user,proto3" json:"user,omitempty"`
	Created              uint64   `protobuf:"varint,6,opt,name=created,proto3" json:"created,omitempty"`
	Updated              uint64   `protobuf:"varint,7,opt,name=updated,proto3" json:"updated,omitempty"`
	Creator              string   `protobuf:"bytes,8,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string   `protobuf:"bytes,9,opt,name=operator,proto3" json:"operator,omitempty"`
	Owner                string   `protobuf:"bytes,10,opt,name=owner,proto3" json:"owner,omitempty"`
	Classes              []string `protobuf:"bytes,11,rep,name=classes,proto3" json:"classes,omitempty"`
	Subjects             []string `protobuf:"bytes,12,rep,name=subjects,proto3" json:"subjects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeacherInfo) Reset()         { *m = TeacherInfo{} }
func (m *TeacherInfo) String() string { return proto.CompactTextString(m) }
func (*TeacherInfo) ProtoMessage()    {}
func (*TeacherInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_93ae7c86bf27648d, []int{0}
}

func (m *TeacherInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeacherInfo.Unmarshal(m, b)
}
func (m *TeacherInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeacherInfo.Marshal(b, m, deterministic)
}
func (m *TeacherInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeacherInfo.Merge(m, src)
}
func (m *TeacherInfo) XXX_Size() int {
	return xxx_messageInfo_TeacherInfo.Size(m)
}
func (m *TeacherInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TeacherInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TeacherInfo proto.InternalMessageInfo

func (m *TeacherInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TeacherInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *TeacherInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TeacherInfo) GetEntity() string {
	if m != nil {
		return m.Entity
	}
	return ""
}

func (m *TeacherInfo) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *TeacherInfo) GetCreated() uint64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *TeacherInfo) GetUpdated() uint64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *TeacherInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *TeacherInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *TeacherInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *TeacherInfo) GetClasses() []string {
	if m != nil {
		return m.Classes
	}
	return nil
}

func (m *TeacherInfo) GetSubjects() []string {
	if m != nil {
		return m.Subjects
	}
	return nil
}

type ReqTeacherAdd struct {
	Owner                string   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Entity               string   `protobuf:"bytes,3,opt,name=entity,proto3" json:"entity,omitempty"`
	User                 string   `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	Operator             string   `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
	Subjects             []string `protobuf:"bytes,6,rep,name=subjects,proto3" json:"subjects,omitempty"`
	Classes              []string `protobuf:"bytes,7,rep,name=classes,proto3" json:"classes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqTeacherAdd) Reset()         { *m = ReqTeacherAdd{} }
func (m *ReqTeacherAdd) String() string { return proto.CompactTextString(m) }
func (*ReqTeacherAdd) ProtoMessage()    {}
func (*ReqTeacherAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_93ae7c86bf27648d, []int{1}
}

func (m *ReqTeacherAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqTeacherAdd.Unmarshal(m, b)
}
func (m *ReqTeacherAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqTeacherAdd.Marshal(b, m, deterministic)
}
func (m *ReqTeacherAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqTeacherAdd.Merge(m, src)
}
func (m *ReqTeacherAdd) XXX_Size() int {
	return xxx_messageInfo_ReqTeacherAdd.Size(m)
}
func (m *ReqTeacherAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqTeacherAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqTeacherAdd proto.InternalMessageInfo

func (m *ReqTeacherAdd) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ReqTeacherAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqTeacherAdd) GetEntity() string {
	if m != nil {
		return m.Entity
	}
	return ""
}

func (m *ReqTeacherAdd) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *ReqTeacherAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqTeacherAdd) GetSubjects() []string {
	if m != nil {
		return m.Subjects
	}
	return nil
}

func (m *ReqTeacherAdd) GetClasses() []string {
	if m != nil {
		return m.Classes
	}
	return nil
}

type ReqTeacherUpdate struct {
	Owner                string   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Operator             string   `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	Classes              []string `protobuf:"bytes,5,rep,name=classes,proto3" json:"classes,omitempty"`
	Subjects             []string `protobuf:"bytes,6,rep,name=subjects,proto3" json:"subjects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqTeacherUpdate) Reset()         { *m = ReqTeacherUpdate{} }
func (m *ReqTeacherUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqTeacherUpdate) ProtoMessage()    {}
func (*ReqTeacherUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_93ae7c86bf27648d, []int{2}
}

func (m *ReqTeacherUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqTeacherUpdate.Unmarshal(m, b)
}
func (m *ReqTeacherUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqTeacherUpdate.Marshal(b, m, deterministic)
}
func (m *ReqTeacherUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqTeacherUpdate.Merge(m, src)
}
func (m *ReqTeacherUpdate) XXX_Size() int {
	return xxx_messageInfo_ReqTeacherUpdate.Size(m)
}
func (m *ReqTeacherUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqTeacherUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqTeacherUpdate proto.InternalMessageInfo

func (m *ReqTeacherUpdate) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ReqTeacherUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqTeacherUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqTeacherUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqTeacherUpdate) GetClasses() []string {
	if m != nil {
		return m.Classes
	}
	return nil
}

func (m *ReqTeacherUpdate) GetSubjects() []string {
	if m != nil {
		return m.Subjects
	}
	return nil
}

type ReqTeacherBatch struct {
	Owner                string           `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Operator             string           `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
	List                 []*ReqTeacherAdd `protobuf:"bytes,3,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ReqTeacherBatch) Reset()         { *m = ReqTeacherBatch{} }
func (m *ReqTeacherBatch) String() string { return proto.CompactTextString(m) }
func (*ReqTeacherBatch) ProtoMessage()    {}
func (*ReqTeacherBatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_93ae7c86bf27648d, []int{3}
}

func (m *ReqTeacherBatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqTeacherBatch.Unmarshal(m, b)
}
func (m *ReqTeacherBatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqTeacherBatch.Marshal(b, m, deterministic)
}
func (m *ReqTeacherBatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqTeacherBatch.Merge(m, src)
}
func (m *ReqTeacherBatch) XXX_Size() int {
	return xxx_messageInfo_ReqTeacherBatch.Size(m)
}
func (m *ReqTeacherBatch) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqTeacherBatch.DiscardUnknown(m)
}

var xxx_messageInfo_ReqTeacherBatch proto.InternalMessageInfo

func (m *ReqTeacherBatch) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ReqTeacherBatch) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqTeacherBatch) GetList() []*ReqTeacherAdd {
	if m != nil {
		return m.List
	}
	return nil
}

type ReplyTeacherInfo struct {
	Info                 *TeacherInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyTeacherInfo) Reset()         { *m = ReplyTeacherInfo{} }
func (m *ReplyTeacherInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyTeacherInfo) ProtoMessage()    {}
func (*ReplyTeacherInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_93ae7c86bf27648d, []int{4}
}

func (m *ReplyTeacherInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyTeacherInfo.Unmarshal(m, b)
}
func (m *ReplyTeacherInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyTeacherInfo.Marshal(b, m, deterministic)
}
func (m *ReplyTeacherInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyTeacherInfo.Merge(m, src)
}
func (m *ReplyTeacherInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyTeacherInfo.Size(m)
}
func (m *ReplyTeacherInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyTeacherInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyTeacherInfo proto.InternalMessageInfo

func (m *ReplyTeacherInfo) GetInfo() *TeacherInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ReplyTeacherInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReplyTeacherList struct {
	Total                uint32         `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Pages                uint32         `protobuf:"varint,2,opt,name=pages,proto3" json:"pages,omitempty"`
	Page                 uint32         `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Number               uint32         `protobuf:"varint,4,opt,name=number,proto3" json:"number,omitempty"`
	Status               *ReplyStatus   `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	List                 []*TeacherInfo `protobuf:"bytes,6,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ReplyTeacherList) Reset()         { *m = ReplyTeacherList{} }
func (m *ReplyTeacherList) String() string { return proto.CompactTextString(m) }
func (*ReplyTeacherList) ProtoMessage()    {}
func (*ReplyTeacherList) Descriptor() ([]byte, []int) {
	return fileDescriptor_93ae7c86bf27648d, []int{5}
}

func (m *ReplyTeacherList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyTeacherList.Unmarshal(m, b)
}
func (m *ReplyTeacherList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyTeacherList.Marshal(b, m, deterministic)
}
func (m *ReplyTeacherList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyTeacherList.Merge(m, src)
}
func (m *ReplyTeacherList) XXX_Size() int {
	return xxx_messageInfo_ReplyTeacherList.Size(m)
}
func (m *ReplyTeacherList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyTeacherList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyTeacherList proto.InternalMessageInfo

func (m *ReplyTeacherList) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReplyTeacherList) GetPages() uint32 {
	if m != nil {
		return m.Pages
	}
	return 0
}

func (m *ReplyTeacherList) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ReplyTeacherList) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *ReplyTeacherList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyTeacherList) GetList() []*TeacherInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*TeacherInfo)(nil), "omo.msp.school.TeacherInfo")
	proto.RegisterType((*ReqTeacherAdd)(nil), "omo.msp.school.ReqTeacherAdd")
	proto.RegisterType((*ReqTeacherUpdate)(nil), "omo.msp.school.ReqTeacherUpdate")
	proto.RegisterType((*ReqTeacherBatch)(nil), "omo.msp.school.ReqTeacherBatch")
	proto.RegisterType((*ReplyTeacherInfo)(nil), "omo.msp.school.ReplyTeacherInfo")
	proto.RegisterType((*ReplyTeacherList)(nil), "omo.msp.school.ReplyTeacherList")
}

func init() {
	proto.RegisterFile("proto/school/teacher.proto", fileDescriptor_93ae7c86bf27648d)
}

var fileDescriptor_93ae7c86bf27648d = []byte{
	// 604 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x5e, 0x7e, 0x9a, 0xae, 0xa7, 0x74, 0x4c, 0x16, 0x42, 0x5e, 0xa7, 0x89, 0x28, 0x57, 0xbb,
	0xea, 0xc4, 0x76, 0xc9, 0xd5, 0x26, 0xa1, 0x69, 0x80, 0x34, 0xc8, 0xe0, 0x86, 0x9b, 0x29, 0x4b,
	0xbc, 0x35, 0xa8, 0x89, 0xb3, 0xd8, 0x29, 0xec, 0x6d, 0xb8, 0xe1, 0x19, 0x78, 0x0d, 0xde, 0x80,
	0x57, 0x41, 0x3e, 0x4e, 0x5b, 0xa7, 0x4a, 0x4a, 0xef, 0xfc, 0xf9, 0x7c, 0xfe, 0xce, 0x77, 0x3e,
	0xc7, 0x0a, 0x8c, 0x8b, 0x92, 0x4b, 0x7e, 0x22, 0xe2, 0x29, 0xe7, 0xb3, 0x13, 0xc9, 0xa2, 0x78,
	0xca, 0xca, 0x09, 0x6e, 0x92, 0x3d, 0x9e, 0xf1, 0x49, 0x26, 0x8a, 0x89, 0xae, 0x8e, 0x0f, 0x1a,
	0xdc, 0x98, 0x67, 0x19, 0xcf, 0x35, 0x35, 0xf8, 0x65, 0xc3, 0xf0, 0xb3, 0x3e, 0x7c, 0x95, 0xdf,
	0x73, 0xb2, 0x07, 0x76, 0x9a, 0x50, 0xcb, 0xb7, 0x8e, 0xdd, 0xd0, 0x4e, 0x13, 0xb2, 0x0f, 0x4e,
	0x95, 0x26, 0xd4, 0xf6, 0xad, 0xe3, 0x41, 0xa8, 0x96, 0x84, 0x80, 0x9b, 0x47, 0x19, 0xa3, 0x0e,
	0x6e, 0xe1, 0x9a, 0xbc, 0x04, 0x8f, 0xe5, 0x32, 0x95, 0x4f, 0xd4, 0xc5, 0xdd, 0x1a, 0x29, 0x6e,
	0x25, 0x58, 0x49, 0x7b, 0x9a, 0xab, 0xd6, 0x84, 0x42, 0x3f, 0x2e, 0x59, 0x24, 0x59, 0x42, 0x3d,
	0x6c, 0xb3, 0x80, 0xaa, 0x52, 0x15, 0x09, 0x56, 0xfa, 0xba, 0x52, 0xc3, 0xe5, 0x19, 0x5e, 0xd2,
	0x5d, 0x94, 0x5a, 0x40, 0x32, 0x86, 0x5d, 0x5e, 0xb0, 0x12, 0x4b, 0x03, 0x2c, 0x2d, 0x31, 0x79,
	0x01, 0x3d, 0xfe, 0x3d, 0x67, 0x25, 0x05, 0x2c, 0x68, 0x80, 0x5a, 0xb3, 0x48, 0x08, 0x26, 0xe8,
	0xd0, 0x77, 0x50, 0x4b, 0x43, 0xa5, 0x25, 0xaa, 0xbb, 0x6f, 0x2c, 0x96, 0x82, 0x3e, 0xc3, 0xd2,
	0x12, 0x07, 0xbf, 0x2d, 0x18, 0x85, 0xec, 0xb1, 0x8e, 0xea, 0x3c, 0x49, 0x56, 0xea, 0x96, 0xa9,
	0xbe, 0x48, 0xc7, 0x6e, 0x4d, 0xc7, 0x69, 0x4d, 0xc7, 0x35, 0xd2, 0x31, 0xe7, 0xe9, 0xad, 0xcd,
	0x63, 0xfa, 0xf3, 0x9a, 0xfe, 0xcc, 0xa9, 0xfa, 0x8d, 0xa9, 0x82, 0x9f, 0x16, 0xec, 0xaf, 0x9c,
	0x7f, 0xc1, 0x44, 0x3b, 0xcc, 0x6f, 0x77, 0xd9, 0xa6, 0x45, 0x77, 0xcd, 0xa2, 0x61, 0xa3, 0xd7,
	0x1d, 0xee, 0x9a, 0xf9, 0x60, 0x0e, 0xcf, 0x57, 0x0e, 0x2f, 0x22, 0x19, 0x4f, 0x3b, 0x0c, 0x9a,
	0xad, 0xed, 0xb5, 0xd6, 0xaf, 0xc1, 0x9d, 0xa5, 0x42, 0x52, 0xc7, 0x77, 0x8e, 0x87, 0xa7, 0x47,
	0x93, 0xe6, 0x1b, 0x98, 0x34, 0x2e, 0x2f, 0x44, 0x6a, 0xf0, 0x43, 0x25, 0x53, 0xcc, 0x9e, 0xcc,
	0x07, 0x70, 0x02, 0x6e, 0x9a, 0xdf, 0x73, 0xec, 0x3b, 0x3c, 0x3d, 0x5c, 0x97, 0x31, 0xa8, 0x21,
	0x12, 0xc9, 0x19, 0x78, 0x42, 0x46, 0xb2, 0x12, 0xe8, 0xa8, 0xe5, 0x08, 0xb6, 0xb8, 0x41, 0x4a,
	0x58, 0x53, 0x83, 0x3f, 0x56, 0xb3, 0xf5, 0x87, 0x54, 0x48, 0x35, 0xb3, 0xe4, 0x32, 0x9a, 0x61,
	0xef, 0x51, 0xa8, 0x81, 0xda, 0x2d, 0xa2, 0x07, 0xa6, 0xe5, 0x47, 0xa1, 0x06, 0xea, 0x62, 0xd4,
	0x02, 0x2f, 0x66, 0x14, 0xe2, 0x5a, 0x7d, 0x67, 0x79, 0x95, 0xdd, 0xd5, 0x5f, 0xd4, 0x28, 0xac,
	0x91, 0xe1, 0xb0, 0xb7, 0xb5, 0x43, 0x95, 0x03, 0xc6, 0xe9, 0x61, 0x9c, 0x9b, 0x73, 0x50, 0xc4,
	0xd3, 0xbf, 0x0e, 0xec, 0xd5, 0xbb, 0x37, 0xac, 0x9c, 0xa7, 0x31, 0x23, 0xef, 0xc1, 0x3b, 0x4f,
	0x92, 0xeb, 0x9c, 0x91, 0xcd, 0xd7, 0x31, 0xf6, 0x5b, 0x1d, 0x19, 0x3d, 0x82, 0x1d, 0x72, 0x05,
	0xde, 0x25, 0x93, 0x4a, 0xac, 0xc5, 0xff, 0x63, 0xc5, 0x84, 0x54, 0xc4, 0xad, 0xa4, 0xde, 0x41,
	0xff, 0x92, 0x49, 0xcc, 0xbc, 0x4b, 0xeb, 0x63, 0xf4, 0xc0, 0x36, 0x6b, 0xa9, 0xe3, 0xc1, 0x0e,
	0xf9, 0x04, 0x03, 0xfd, 0xa6, 0x94, 0x33, 0xbf, 0x7b, 0x4c, 0x4d, 0xda, 0xca, 0xde, 0x5b, 0x18,
	0x84, 0x2c, 0xe3, 0x73, 0xf6, 0xdf, 0x61, 0x0f, 0x5a, 0xd5, 0x6a, 0x99, 0x6b, 0xd8, 0x3d, 0x4f,
	0x12, 0xfd, 0x9c, 0x5e, 0x75, 0x1b, 0x43, 0xc2, 0x36, 0xa3, 0x5e, 0x1c, 0x7d, 0x3d, 0x34, 0x7f,
	0x24, 0x6f, 0x78, 0xc6, 0x6f, 0x33, 0x51, 0xdc, 0x6a, 0x78, 0xe7, 0x61, 0xf1, 0xec, 0x5f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x77, 0x1a, 0xa0, 0xc2, 0x9a, 0x06, 0x00, 0x00,
}
