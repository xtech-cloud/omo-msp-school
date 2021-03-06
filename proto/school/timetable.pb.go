// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/school/timetable.proto

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

type TimetableInfo struct {
	Id                   uint64           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid                  string           `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string           `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Created              uint64           `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	Updated              uint64           `protobuf:"varint,5,opt,name=updated,proto3" json:"updated,omitempty"`
	Creator              string           `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string           `protobuf:"bytes,7,opt,name=operator,proto3" json:"operator,omitempty"`
	School               string           `protobuf:"bytes,8,opt,name=school,proto3" json:"school,omitempty"`
	Class                string           `protobuf:"bytes,9,opt,name=class,proto3" json:"class,omitempty"`
	Year                 uint32           `protobuf:"varint,10,opt,name=year,proto3" json:"year,omitempty"`
	Items                []*TimetableItem `protobuf:"bytes,11,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TimetableInfo) Reset()         { *m = TimetableInfo{} }
func (m *TimetableInfo) String() string { return proto.CompactTextString(m) }
func (*TimetableInfo) ProtoMessage()    {}
func (*TimetableInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5794f740e098e8c, []int{0}
}

func (m *TimetableInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimetableInfo.Unmarshal(m, b)
}
func (m *TimetableInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimetableInfo.Marshal(b, m, deterministic)
}
func (m *TimetableInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimetableInfo.Merge(m, src)
}
func (m *TimetableInfo) XXX_Size() int {
	return xxx_messageInfo_TimetableInfo.Size(m)
}
func (m *TimetableInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TimetableInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TimetableInfo proto.InternalMessageInfo

func (m *TimetableInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TimetableInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *TimetableInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TimetableInfo) GetCreated() uint64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *TimetableInfo) GetUpdated() uint64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *TimetableInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *TimetableInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *TimetableInfo) GetSchool() string {
	if m != nil {
		return m.School
	}
	return ""
}

func (m *TimetableInfo) GetClass() string {
	if m != nil {
		return m.Class
	}
	return ""
}

func (m *TimetableInfo) GetYear() uint32 {
	if m != nil {
		return m.Year
	}
	return 0
}

func (m *TimetableInfo) GetItems() []*TimetableItem {
	if m != nil {
		return m.Items
	}
	return nil
}

type TimetableItem struct {
	Weekday              uint32   `protobuf:"varint,1,opt,name=weekday,proto3" json:"weekday,omitempty"`
	Number               uint32   `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Class                string   `protobuf:"bytes,4,opt,name=class,proto3" json:"class,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimetableItem) Reset()         { *m = TimetableItem{} }
func (m *TimetableItem) String() string { return proto.CompactTextString(m) }
func (*TimetableItem) ProtoMessage()    {}
func (*TimetableItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5794f740e098e8c, []int{1}
}

func (m *TimetableItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimetableItem.Unmarshal(m, b)
}
func (m *TimetableItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimetableItem.Marshal(b, m, deterministic)
}
func (m *TimetableItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimetableItem.Merge(m, src)
}
func (m *TimetableItem) XXX_Size() int {
	return xxx_messageInfo_TimetableItem.Size(m)
}
func (m *TimetableItem) XXX_DiscardUnknown() {
	xxx_messageInfo_TimetableItem.DiscardUnknown(m)
}

var xxx_messageInfo_TimetableItem proto.InternalMessageInfo

func (m *TimetableItem) GetWeekday() uint32 {
	if m != nil {
		return m.Weekday
	}
	return 0
}

func (m *TimetableItem) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *TimetableItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TimetableItem) GetClass() string {
	if m != nil {
		return m.Class
	}
	return ""
}

type ReqTimetableAdd struct {
	School               string           `protobuf:"bytes,1,opt,name=school,proto3" json:"school,omitempty"`
	Class                string           `protobuf:"bytes,2,opt,name=class,proto3" json:"class,omitempty"`
	Year                 uint32           `protobuf:"varint,3,opt,name=year,proto3" json:"year,omitempty"`
	Operator             string           `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	Items                []*TimetableItem `protobuf:"bytes,5,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ReqTimetableAdd) Reset()         { *m = ReqTimetableAdd{} }
func (m *ReqTimetableAdd) String() string { return proto.CompactTextString(m) }
func (*ReqTimetableAdd) ProtoMessage()    {}
func (*ReqTimetableAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5794f740e098e8c, []int{2}
}

func (m *ReqTimetableAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqTimetableAdd.Unmarshal(m, b)
}
func (m *ReqTimetableAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqTimetableAdd.Marshal(b, m, deterministic)
}
func (m *ReqTimetableAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqTimetableAdd.Merge(m, src)
}
func (m *ReqTimetableAdd) XXX_Size() int {
	return xxx_messageInfo_ReqTimetableAdd.Size(m)
}
func (m *ReqTimetableAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqTimetableAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqTimetableAdd proto.InternalMessageInfo

func (m *ReqTimetableAdd) GetSchool() string {
	if m != nil {
		return m.School
	}
	return ""
}

func (m *ReqTimetableAdd) GetClass() string {
	if m != nil {
		return m.Class
	}
	return ""
}

func (m *ReqTimetableAdd) GetYear() uint32 {
	if m != nil {
		return m.Year
	}
	return 0
}

func (m *ReqTimetableAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqTimetableAdd) GetItems() []*TimetableItem {
	if m != nil {
		return m.Items
	}
	return nil
}

type ReqTimetableBatch struct {
	School               string           `protobuf:"bytes,1,opt,name=school,proto3" json:"school,omitempty"`
	Operator             string           `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
	Year                 uint32           `protobuf:"varint,3,opt,name=year,proto3" json:"year,omitempty"`
	Items                []*TimetableItem `protobuf:"bytes,4,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ReqTimetableBatch) Reset()         { *m = ReqTimetableBatch{} }
func (m *ReqTimetableBatch) String() string { return proto.CompactTextString(m) }
func (*ReqTimetableBatch) ProtoMessage()    {}
func (*ReqTimetableBatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5794f740e098e8c, []int{3}
}

func (m *ReqTimetableBatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqTimetableBatch.Unmarshal(m, b)
}
func (m *ReqTimetableBatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqTimetableBatch.Marshal(b, m, deterministic)
}
func (m *ReqTimetableBatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqTimetableBatch.Merge(m, src)
}
func (m *ReqTimetableBatch) XXX_Size() int {
	return xxx_messageInfo_ReqTimetableBatch.Size(m)
}
func (m *ReqTimetableBatch) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqTimetableBatch.DiscardUnknown(m)
}

var xxx_messageInfo_ReqTimetableBatch proto.InternalMessageInfo

func (m *ReqTimetableBatch) GetSchool() string {
	if m != nil {
		return m.School
	}
	return ""
}

func (m *ReqTimetableBatch) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqTimetableBatch) GetYear() uint32 {
	if m != nil {
		return m.Year
	}
	return 0
}

func (m *ReqTimetableBatch) GetItems() []*TimetableItem {
	if m != nil {
		return m.Items
	}
	return nil
}

type ReplyTimetableInfo struct {
	Status               *ReplyStatus   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Info                 *TimetableInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ReplyTimetableInfo) Reset()         { *m = ReplyTimetableInfo{} }
func (m *ReplyTimetableInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyTimetableInfo) ProtoMessage()    {}
func (*ReplyTimetableInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5794f740e098e8c, []int{4}
}

func (m *ReplyTimetableInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyTimetableInfo.Unmarshal(m, b)
}
func (m *ReplyTimetableInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyTimetableInfo.Marshal(b, m, deterministic)
}
func (m *ReplyTimetableInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyTimetableInfo.Merge(m, src)
}
func (m *ReplyTimetableInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyTimetableInfo.Size(m)
}
func (m *ReplyTimetableInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyTimetableInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyTimetableInfo proto.InternalMessageInfo

func (m *ReplyTimetableInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyTimetableInfo) GetInfo() *TimetableInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type ReplyTimetableList struct {
	Status               *ReplyStatus     `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	List                 []*TimetableInfo `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ReplyTimetableList) Reset()         { *m = ReplyTimetableList{} }
func (m *ReplyTimetableList) String() string { return proto.CompactTextString(m) }
func (*ReplyTimetableList) ProtoMessage()    {}
func (*ReplyTimetableList) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5794f740e098e8c, []int{5}
}

func (m *ReplyTimetableList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyTimetableList.Unmarshal(m, b)
}
func (m *ReplyTimetableList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyTimetableList.Marshal(b, m, deterministic)
}
func (m *ReplyTimetableList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyTimetableList.Merge(m, src)
}
func (m *ReplyTimetableList) XXX_Size() int {
	return xxx_messageInfo_ReplyTimetableList.Size(m)
}
func (m *ReplyTimetableList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyTimetableList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyTimetableList proto.InternalMessageInfo

func (m *ReplyTimetableList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyTimetableList) GetList() []*TimetableInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*TimetableInfo)(nil), "omo.msp.school.TimetableInfo")
	proto.RegisterType((*TimetableItem)(nil), "omo.msp.school.TimetableItem")
	proto.RegisterType((*ReqTimetableAdd)(nil), "omo.msp.school.ReqTimetableAdd")
	proto.RegisterType((*ReqTimetableBatch)(nil), "omo.msp.school.ReqTimetableBatch")
	proto.RegisterType((*ReplyTimetableInfo)(nil), "omo.msp.school.ReplyTimetableInfo")
	proto.RegisterType((*ReplyTimetableList)(nil), "omo.msp.school.ReplyTimetableList")
}

func init() {
	proto.RegisterFile("proto/school/timetable.proto", fileDescriptor_e5794f740e098e8c)
}

var fileDescriptor_e5794f740e098e8c = []byte{
	// 566 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x41, 0x6f, 0xd3, 0x4c,
	0x10, 0xad, 0x1d, 0x27, 0x6d, 0x27, 0x5f, 0xfa, 0x95, 0x15, 0x42, 0x6e, 0x4a, 0x21, 0xf8, 0x94,
	0x53, 0x2a, 0x9a, 0x23, 0xa7, 0x44, 0x82, 0x08, 0xa9, 0xa8, 0x95, 0xc3, 0x89, 0x4b, 0xe5, 0x78,
	0x27, 0x74, 0x55, 0xaf, 0xd7, 0xf5, 0x6e, 0x8a, 0x22, 0xf1, 0x23, 0xf8, 0x11, 0x88, 0x9f, 0xc7,
	0x6f, 0x40, 0x1e, 0x27, 0xc6, 0x49, 0x9d, 0x94, 0x22, 0x6e, 0xfb, 0x76, 0x66, 0xdf, 0x7b, 0xf3,
	0xc6, 0x32, 0x3c, 0x4f, 0x52, 0x65, 0xd4, 0xa9, 0x0e, 0xaf, 0x95, 0x8a, 0x4e, 0x8d, 0x90, 0x68,
	0x82, 0x49, 0x84, 0x3d, 0xba, 0x66, 0x07, 0x4a, 0xaa, 0x9e, 0xd4, 0x49, 0x2f, 0xaf, 0xb7, 0x8f,
	0x56, 0xba, 0x43, 0x25, 0xa5, 0x8a, 0xf3, 0x56, 0xef, 0x87, 0x0d, 0xad, 0x8f, 0xcb, 0xe7, 0xef,
	0xe3, 0xa9, 0x62, 0x07, 0x60, 0x0b, 0xee, 0x5a, 0x1d, 0xab, 0xeb, 0xf8, 0xb6, 0xe0, 0xec, 0x10,
	0x6a, 0x33, 0xc1, 0x5d, 0xbb, 0x63, 0x75, 0xf7, 0xfd, 0xec, 0xc8, 0x18, 0x38, 0x71, 0x20, 0xd1,
	0xad, 0xd1, 0x15, 0x9d, 0x99, 0x0b, 0xbb, 0x61, 0x8a, 0x81, 0x41, 0xee, 0x3a, 0xf4, 0x74, 0x09,
	0xb3, 0xca, 0x2c, 0xe1, 0x54, 0xa9, 0xe7, 0x95, 0x05, 0x2c, 0xde, 0xa8, 0xd4, 0x6d, 0x10, 0xd5,
	0x12, 0xb2, 0x36, 0xec, 0xa9, 0x04, 0x53, 0x2a, 0xed, 0x52, 0xa9, 0xc0, 0xec, 0x19, 0x34, 0xf2,
	0x41, 0xdc, 0x3d, 0xaa, 0x2c, 0x10, 0x7b, 0x0a, 0xf5, 0x30, 0x0a, 0xb4, 0x76, 0xf7, 0xe9, 0x3a,
	0x07, 0x99, 0xd7, 0x39, 0x06, 0xa9, 0x0b, 0x1d, 0xab, 0xdb, 0xf2, 0xe9, 0xcc, 0xfa, 0x50, 0x17,
	0x06, 0xa5, 0x76, 0x9b, 0x9d, 0x5a, 0xb7, 0x79, 0x76, 0xd2, 0x5b, 0x8d, 0xab, 0xf7, 0x3b, 0x0f,
	0x83, 0xd2, 0xcf, 0x7b, 0xbd, 0x9b, 0x72, 0x4e, 0x06, 0x65, 0xe6, 0xfe, 0x0b, 0xe2, 0x0d, 0x0f,
	0xe6, 0x14, 0x56, 0xcb, 0x5f, 0xc2, 0xcc, 0x61, 0x3c, 0x93, 0x13, 0x4c, 0x29, 0xb4, 0x96, 0xbf,
	0x40, 0x95, 0xb9, 0x15, 0xae, 0x9d, 0x92, 0x6b, 0xef, 0xbb, 0x05, 0xff, 0xfb, 0x78, 0x5b, 0x08,
	0x0e, 0x38, 0x2f, 0xcd, 0x6d, 0x55, 0xcf, 0x6d, 0x57, 0xcd, 0x5d, 0x2b, 0xcd, 0x5d, 0x4e, 0xd5,
	0x59, 0x4b, 0xb5, 0xc8, 0xa4, 0xfe, 0x88, 0x4c, 0xbe, 0x59, 0xf0, 0xa4, 0x6c, 0x73, 0x18, 0x98,
	0xf0, 0x7a, 0xa3, 0xd1, 0xb2, 0xbc, 0xbd, 0x26, 0x5f, 0x65, 0xb7, 0xb0, 0xe4, 0x3c, 0xc2, 0xd2,
	0x57, 0x60, 0x3e, 0x26, 0xd1, 0x7c, 0xf5, 0x9b, 0xee, 0x43, 0x43, 0x9b, 0xc0, 0xcc, 0x34, 0x59,
	0x6a, 0x9e, 0x1d, 0xaf, 0x73, 0xd1, 0x9b, 0x31, 0xb5, 0xf8, 0x8b, 0x56, 0xf6, 0x1a, 0x1c, 0x11,
	0x4f, 0x15, 0x79, 0xdd, 0x2a, 0x1f, 0x4f, 0x95, 0x4f, 0xad, 0xf7, 0xd5, 0xcf, 0x85, 0x36, 0x7f,
	0xad, 0x1e, 0x09, 0x6d, 0x5c, 0xfb, 0xa1, 0xe1, 0x49, 0x3d, 0x6b, 0x3d, 0xfb, 0x59, 0x83, 0xc3,
	0xe2, 0x7e, 0x8c, 0xe9, 0x9d, 0x08, 0x91, 0x5d, 0x40, 0x63, 0xc0, 0xf9, 0x45, 0x8c, 0xec, 0xe5,
	0x7d, 0xd9, 0x95, 0x2f, 0xac, 0xed, 0x55, 0xfa, 0x5a, 0x51, 0xf2, 0x76, 0xd8, 0x18, 0xf6, 0x06,
	0x9c, 0xe7, 0xab, 0x7e, 0xb5, 0x8d, 0x92, 0x5a, 0x1e, 0x22, 0xcd, 0x02, 0xf2, 0x76, 0xd8, 0x39,
	0xec, 0x8e, 0xd0, 0x50, 0x5a, 0x15, 0xe9, 0xdc, 0xce, 0x50, 0x9b, 0xcb, 0xe0, 0x33, 0xfe, 0x21,
	0xdb, 0x25, 0x34, 0x47, 0x68, 0x86, 0xf3, 0x77, 0x22, 0x32, 0x98, 0xfe, 0x0b, 0xc6, 0x0f, 0xf0,
	0xdf, 0x08, 0x4d, 0xb6, 0x22, 0xa1, 0x8d, 0x08, 0xb7, 0x53, 0xbe, 0xd8, 0xb8, 0x5f, 0x7a, 0xec,
	0xed, 0xb0, 0xb7, 0xb0, 0xef, 0xa3, 0x54, 0x77, 0x98, 0xed, 0x65, 0x13, 0x57, 0x96, 0x77, 0xfb,
	0xa8, 0x92, 0x2b, 0x5f, 0xc5, 0xf0, 0xe4, 0xd3, 0x71, 0xf9, 0xcf, 0xfe, 0x46, 0x49, 0x75, 0x25,
	0x75, 0x72, 0x95, 0xc3, 0x49, 0x83, 0x8a, 0xfd, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x59, 0xe1,
	0xac, 0x55, 0x2d, 0x06, 0x00, 0x00,
}
