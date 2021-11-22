// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/school/common.proto

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

type OptionType int32

const (
	OptionType_Add    OptionType = 0
	OptionType_Update OptionType = 1
	OptionType_Remove OptionType = 2
	OptionType_Get    OptionType = 3
)

var OptionType_name = map[int32]string{
	0: "Add",
	1: "Update",
	2: "Remove",
	3: "Get",
}

var OptionType_value = map[string]int32{
	"Add":    0,
	"Update": 1,
	"Remove": 2,
	"Get":    3,
}

func (x OptionType) String() string {
	return proto.EnumName(OptionType_name, int32(x))
}

func (OptionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b102dd72cb6e728a, []int{0}
}

type TargetType int32

const (
	TargetType_TSelf       TargetType = 0
	TargetType_TStudent    TargetType = 1
	TargetType_TTeacher    TargetType = 2
	TargetType_TClass      TargetType = 3
	TargetType_TStudent_GH TargetType = 4
	TargetType_TTeacher_GH TargetType = 5
	TargetType_TSubject    TargetType = 6
)

var TargetType_name = map[int32]string{
	0: "TSelf",
	1: "TStudent",
	2: "TTeacher",
	3: "TClass",
	4: "TStudent_GH",
	5: "TTeacher_GH",
	6: "TSubject",
}

var TargetType_value = map[string]int32{
	"TSelf":       0,
	"TStudent":    1,
	"TTeacher":    2,
	"TClass":      3,
	"TStudent_GH": 4,
	"TTeacher_GH": 5,
	"TSubject":    6,
}

func (x TargetType) String() string {
	return proto.EnumName(TargetType_name, int32(x))
}

func (TargetType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b102dd72cb6e728a, []int{1}
}

type PairInfo struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PairInfo) Reset()         { *m = PairInfo{} }
func (m *PairInfo) String() string { return proto.CompactTextString(m) }
func (*PairInfo) ProtoMessage()    {}
func (*PairInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_b102dd72cb6e728a, []int{0}
}

func (m *PairInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PairInfo.Unmarshal(m, b)
}
func (m *PairInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PairInfo.Marshal(b, m, deterministic)
}
func (m *PairInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PairInfo.Merge(m, src)
}
func (m *PairInfo) XXX_Size() int {
	return xxx_messageInfo_PairInfo.Size(m)
}
func (m *PairInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PairInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PairInfo proto.InternalMessageInfo

func (m *PairInfo) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *PairInfo) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type RequestInfo struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Operator             string   `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
	Flag                 string   `protobuf:"bytes,3,opt,name=flag,proto3" json:"flag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestInfo) Reset()         { *m = RequestInfo{} }
func (m *RequestInfo) String() string { return proto.CompactTextString(m) }
func (*RequestInfo) ProtoMessage()    {}
func (*RequestInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_b102dd72cb6e728a, []int{1}
}

func (m *RequestInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestInfo.Unmarshal(m, b)
}
func (m *RequestInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestInfo.Marshal(b, m, deterministic)
}
func (m *RequestInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestInfo.Merge(m, src)
}
func (m *RequestInfo) XXX_Size() int {
	return xxx_messageInfo_RequestInfo.Size(m)
}
func (m *RequestInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RequestInfo proto.InternalMessageInfo

func (m *RequestInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RequestInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *RequestInfo) GetFlag() string {
	if m != nil {
		return m.Flag
	}
	return ""
}

type RequestPage struct {
	Page                 uint32   `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Number               uint32   `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	Flag                 string   `protobuf:"bytes,3,opt,name=flag,proto3" json:"flag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestPage) Reset()         { *m = RequestPage{} }
func (m *RequestPage) String() string { return proto.CompactTextString(m) }
func (*RequestPage) ProtoMessage()    {}
func (*RequestPage) Descriptor() ([]byte, []int) {
	return fileDescriptor_b102dd72cb6e728a, []int{2}
}

func (m *RequestPage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestPage.Unmarshal(m, b)
}
func (m *RequestPage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestPage.Marshal(b, m, deterministic)
}
func (m *RequestPage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestPage.Merge(m, src)
}
func (m *RequestPage) XXX_Size() int {
	return xxx_messageInfo_RequestPage.Size(m)
}
func (m *RequestPage) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestPage.DiscardUnknown(m)
}

var xxx_messageInfo_RequestPage proto.InternalMessageInfo

func (m *RequestPage) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *RequestPage) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *RequestPage) GetFlag() string {
	if m != nil {
		return m.Flag
	}
	return ""
}

type ReplyStatus struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReplyStatus) Reset()         { *m = ReplyStatus{} }
func (m *ReplyStatus) String() string { return proto.CompactTextString(m) }
func (*ReplyStatus) ProtoMessage()    {}
func (*ReplyStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_b102dd72cb6e728a, []int{3}
}

func (m *ReplyStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyStatus.Unmarshal(m, b)
}
func (m *ReplyStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyStatus.Marshal(b, m, deterministic)
}
func (m *ReplyStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyStatus.Merge(m, src)
}
func (m *ReplyStatus) XXX_Size() int {
	return xxx_messageInfo_ReplyStatus.Size(m)
}
func (m *ReplyStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyStatus proto.InternalMessageInfo

func (m *ReplyStatus) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ReplyStatus) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type ReplyInfo struct {
	Uid                  string       `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyInfo) Reset()         { *m = ReplyInfo{} }
func (m *ReplyInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyInfo) ProtoMessage()    {}
func (*ReplyInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_b102dd72cb6e728a, []int{4}
}

func (m *ReplyInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyInfo.Unmarshal(m, b)
}
func (m *ReplyInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyInfo.Marshal(b, m, deterministic)
}
func (m *ReplyInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyInfo.Merge(m, src)
}
func (m *ReplyInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyInfo.Size(m)
}
func (m *ReplyInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyInfo proto.InternalMessageInfo

func (m *ReplyInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReplyInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterEnum("omo.msp.school.OptionType", OptionType_name, OptionType_value)
	proto.RegisterEnum("omo.msp.school.TargetType", TargetType_name, TargetType_value)
	proto.RegisterType((*PairInfo)(nil), "omo.msp.school.PairInfo")
	proto.RegisterType((*RequestInfo)(nil), "omo.msp.school.RequestInfo")
	proto.RegisterType((*RequestPage)(nil), "omo.msp.school.RequestPage")
	proto.RegisterType((*ReplyStatus)(nil), "omo.msp.school.ReplyStatus")
	proto.RegisterType((*ReplyInfo)(nil), "omo.msp.school.ReplyInfo")
}

func init() {
	proto.RegisterFile("proto/school/common.proto", fileDescriptor_b102dd72cb6e728a)
}

var fileDescriptor_b102dd72cb6e728a = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xc1, 0xaf, 0xd2, 0x40,
	0x18, 0xc4, 0x5f, 0xe9, 0xa3, 0xc2, 0x87, 0x4f, 0x37, 0x1b, 0x63, 0xd0, 0x17, 0x13, 0xd3, 0x93,
	0xe1, 0x50, 0x12, 0x48, 0xbc, 0x78, 0x52, 0x0f, 0xe8, 0xc1, 0x40, 0x4a, 0xbd, 0x78, 0x21, 0x4b,
	0xfb, 0x51, 0xaa, 0xdd, 0x7e, 0x6b, 0x77, 0x97, 0x84, 0xff, 0xde, 0xec, 0xb6, 0x10, 0x4c, 0xf4,
	0x36, 0x33, 0x3b, 0xf3, 0xeb, 0xb6, 0x29, 0xbc, 0x52, 0x2d, 0x19, 0x9a, 0xeb, 0xfc, 0x48, 0x54,
	0xcf, 0x73, 0x92, 0x92, 0x9a, 0xc4, 0x67, 0xfc, 0x19, 0x49, 0x4a, 0xa4, 0x56, 0x49, 0x77, 0x18,
	0x2f, 0x60, 0xb4, 0x11, 0x55, 0xfb, 0xb5, 0x39, 0x10, 0x67, 0x10, 0xfe, 0xc2, 0xf3, 0x34, 0x78,
	0x1b, 0xbc, 0x1b, 0xa7, 0x4e, 0xf2, 0x17, 0x30, 0x3c, 0x89, 0xda, 0xe2, 0x74, 0xe0, 0xb3, 0xce,
	0xc4, 0x6b, 0x98, 0xa4, 0xf8, 0xdb, 0xa2, 0x36, 0x97, 0x99, 0xad, 0x8a, 0xcb, 0xcc, 0x56, 0x05,
	0x7f, 0x0d, 0x23, 0x52, 0xd8, 0x0a, 0x43, 0x6d, 0xbf, 0xbc, 0x7a, 0xce, 0xe1, 0xfe, 0x50, 0x8b,
	0x72, 0x1a, 0xfa, 0xdc, 0xeb, 0xf8, 0xdb, 0x15, 0xb8, 0x11, 0x25, 0xba, 0x8a, 0x12, 0x25, 0x7a,
	0xe2, 0x43, 0xea, 0x35, 0x7f, 0x09, 0x51, 0x63, 0xe5, 0x1e, 0x3b, 0xe0, 0x43, 0xda, 0xbb, 0x7f,
	0xe2, 0x96, 0x0e, 0xa7, 0xea, 0xf3, 0xd6, 0x08, 0x63, 0xb5, 0xab, 0xe4, 0x54, 0x5c, 0x71, 0x4e,
	0xbb, 0x3b, 0x4b, 0x5d, 0xf6, 0x97, 0x73, 0x32, 0x4e, 0x61, 0xec, 0x47, 0xff, 0x79, 0xa5, 0x25,
	0x44, 0xda, 0xe3, 0xfc, 0x66, 0xb2, 0x78, 0x4c, 0xfe, 0xfe, 0x90, 0xc9, 0xcd, 0x13, 0xd3, 0xbe,
	0x3a, 0x7b, 0x0f, 0xb0, 0x56, 0xa6, 0xa2, 0x26, 0x3b, 0x2b, 0xe4, 0x4f, 0x20, 0xfc, 0x58, 0x14,
	0xec, 0x8e, 0x03, 0x44, 0xdf, 0x55, 0x21, 0x0c, 0xb2, 0xc0, 0xe9, 0x14, 0x25, 0x9d, 0x90, 0x0d,
	0x5c, 0x61, 0x85, 0x86, 0x85, 0x33, 0x02, 0xc8, 0x44, 0x5b, 0xa2, 0xf1, 0xbb, 0x31, 0x0c, 0xb3,
	0x2d, 0xd6, 0x07, 0x76, 0xc7, 0x9f, 0xc2, 0x28, 0xdb, 0x1a, 0x5b, 0x60, 0x63, 0x58, 0xe0, 0x5d,
	0x86, 0x22, 0x3f, 0x62, 0xcb, 0x06, 0x8e, 0x94, 0x7d, 0xae, 0x85, 0xd6, 0x2c, 0xe4, 0xcf, 0x61,
	0x72, 0xe9, 0xed, 0x56, 0x5f, 0xd8, 0xbd, 0x0f, 0xfa, 0xaa, 0x0b, 0x86, 0x1d, 0xc9, 0xee, 0x7f,
	0x62, 0x6e, 0x58, 0xf4, 0xe9, 0xcd, 0x8f, 0xc7, 0xdb, 0x5f, 0xe6, 0x03, 0x49, 0xda, 0x49, 0xad,
	0x76, 0x9d, 0xdd, 0x47, 0xfe, 0x70, 0xf9, 0x27, 0x00, 0x00, 0xff, 0xff, 0xe0, 0x9a, 0xfc, 0x62,
	0x58, 0x02, 0x00, 0x00,
}
