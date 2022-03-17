// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/school/classes.proto

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

type MemberInfo struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Student              string   `protobuf:"bytes,2,opt,name=student,proto3" json:"student,omitempty"`
	Status               uint32   `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	Remark               string   `protobuf:"bytes,4,opt,name=remark,proto3" json:"remark,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MemberInfo) Reset()         { *m = MemberInfo{} }
func (m *MemberInfo) String() string { return proto.CompactTextString(m) }
func (*MemberInfo) ProtoMessage()    {}
func (*MemberInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_49492fe1160d7d70, []int{0}
}

func (m *MemberInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemberInfo.Unmarshal(m, b)
}
func (m *MemberInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemberInfo.Marshal(b, m, deterministic)
}
func (m *MemberInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberInfo.Merge(m, src)
}
func (m *MemberInfo) XXX_Size() int {
	return xxx_messageInfo_MemberInfo.Size(m)
}
func (m *MemberInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MemberInfo proto.InternalMessageInfo

func (m *MemberInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *MemberInfo) GetStudent() string {
	if m != nil {
		return m.Student
	}
	return ""
}

func (m *MemberInfo) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *MemberInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

type ClassInfo struct {
	Id                   uint64        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid                  string        `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string        `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Created              uint64        `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	Updated              uint64        `protobuf:"varint,5,opt,name=updated,proto3" json:"updated,omitempty"`
	Creator              string        `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string        `protobuf:"bytes,7,opt,name=operator,proto3" json:"operator,omitempty"`
	Owner                string        `protobuf:"bytes,8,opt,name=owner,proto3" json:"owner,omitempty"`
	Enrol                string        `protobuf:"bytes,9,opt,name=enrol,proto3" json:"enrol,omitempty"`
	Grade                uint32        `protobuf:"varint,10,opt,name=grade,proto3" json:"grade,omitempty"`
	Type                 uint32        `protobuf:"varint,11,opt,name=type,proto3" json:"type,omitempty"`
	No                   uint32        `protobuf:"varint,12,opt,name=no,proto3" json:"no,omitempty"`
	Master               string        `protobuf:"bytes,13,opt,name=master,proto3" json:"master,omitempty"`
	Assistant            string        `protobuf:"bytes,14,opt,name=assistant,proto3" json:"assistant,omitempty"`
	Teachers             []string      `protobuf:"bytes,15,rep,name=teachers,proto3" json:"teachers,omitempty"`
	Students             []*MemberInfo `protobuf:"bytes,16,rep,name=students,proto3" json:"students,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ClassInfo) Reset()         { *m = ClassInfo{} }
func (m *ClassInfo) String() string { return proto.CompactTextString(m) }
func (*ClassInfo) ProtoMessage()    {}
func (*ClassInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_49492fe1160d7d70, []int{1}
}

func (m *ClassInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClassInfo.Unmarshal(m, b)
}
func (m *ClassInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClassInfo.Marshal(b, m, deterministic)
}
func (m *ClassInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassInfo.Merge(m, src)
}
func (m *ClassInfo) XXX_Size() int {
	return xxx_messageInfo_ClassInfo.Size(m)
}
func (m *ClassInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ClassInfo proto.InternalMessageInfo

func (m *ClassInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ClassInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ClassInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ClassInfo) GetCreated() uint64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *ClassInfo) GetUpdated() uint64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *ClassInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *ClassInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ClassInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ClassInfo) GetEnrol() string {
	if m != nil {
		return m.Enrol
	}
	return ""
}

func (m *ClassInfo) GetGrade() uint32 {
	if m != nil {
		return m.Grade
	}
	return 0
}

func (m *ClassInfo) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ClassInfo) GetNo() uint32 {
	if m != nil {
		return m.No
	}
	return 0
}

func (m *ClassInfo) GetMaster() string {
	if m != nil {
		return m.Master
	}
	return ""
}

func (m *ClassInfo) GetAssistant() string {
	if m != nil {
		return m.Assistant
	}
	return ""
}

func (m *ClassInfo) GetTeachers() []string {
	if m != nil {
		return m.Teachers
	}
	return nil
}

func (m *ClassInfo) GetStudents() []*MemberInfo {
	if m != nil {
		return m.Students
	}
	return nil
}

type ReqClassAdd struct {
	Scene                string   `protobuf:"bytes,1,opt,name=scene,proto3" json:"scene,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Enrol                string   `protobuf:"bytes,3,opt,name=enrol,proto3" json:"enrol,omitempty"`
	Count                uint32   `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	Operator             string   `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
	Type                 uint32   `protobuf:"varint,6,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqClassAdd) Reset()         { *m = ReqClassAdd{} }
func (m *ReqClassAdd) String() string { return proto.CompactTextString(m) }
func (*ReqClassAdd) ProtoMessage()    {}
func (*ReqClassAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_49492fe1160d7d70, []int{2}
}

func (m *ReqClassAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqClassAdd.Unmarshal(m, b)
}
func (m *ReqClassAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqClassAdd.Marshal(b, m, deterministic)
}
func (m *ReqClassAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqClassAdd.Merge(m, src)
}
func (m *ReqClassAdd) XXX_Size() int {
	return xxx_messageInfo_ReqClassAdd.Size(m)
}
func (m *ReqClassAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqClassAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqClassAdd proto.InternalMessageInfo

func (m *ReqClassAdd) GetScene() string {
	if m != nil {
		return m.Scene
	}
	return ""
}

func (m *ReqClassAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqClassAdd) GetEnrol() string {
	if m != nil {
		return m.Enrol
	}
	return ""
}

func (m *ReqClassAdd) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ReqClassAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqClassAdd) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

type ReqClassUpdate struct {
	Parent               string   `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Operator             string   `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqClassUpdate) Reset()         { *m = ReqClassUpdate{} }
func (m *ReqClassUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqClassUpdate) ProtoMessage()    {}
func (*ReqClassUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_49492fe1160d7d70, []int{3}
}

func (m *ReqClassUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqClassUpdate.Unmarshal(m, b)
}
func (m *ReqClassUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqClassUpdate.Marshal(b, m, deterministic)
}
func (m *ReqClassUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqClassUpdate.Merge(m, src)
}
func (m *ReqClassUpdate) XXX_Size() int {
	return xxx_messageInfo_ReqClassUpdate.Size(m)
}
func (m *ReqClassUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqClassUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqClassUpdate proto.InternalMessageInfo

func (m *ReqClassUpdate) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ReqClassUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqClassUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqClassUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type ReqClassMaster struct {
	Parent               string   `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Teacher              string   `protobuf:"bytes,3,opt,name=teacher,proto3" json:"teacher,omitempty"`
	Operator             string   `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqClassMaster) Reset()         { *m = ReqClassMaster{} }
func (m *ReqClassMaster) String() string { return proto.CompactTextString(m) }
func (*ReqClassMaster) ProtoMessage()    {}
func (*ReqClassMaster) Descriptor() ([]byte, []int) {
	return fileDescriptor_49492fe1160d7d70, []int{4}
}

func (m *ReqClassMaster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqClassMaster.Unmarshal(m, b)
}
func (m *ReqClassMaster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqClassMaster.Marshal(b, m, deterministic)
}
func (m *ReqClassMaster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqClassMaster.Merge(m, src)
}
func (m *ReqClassMaster) XXX_Size() int {
	return xxx_messageInfo_ReqClassMaster.Size(m)
}
func (m *ReqClassMaster) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqClassMaster.DiscardUnknown(m)
}

var xxx_messageInfo_ReqClassMaster proto.InternalMessageInfo

func (m *ReqClassMaster) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ReqClassMaster) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqClassMaster) GetTeacher() string {
	if m != nil {
		return m.Teacher
	}
	return ""
}

func (m *ReqClassMaster) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type ReqClassStudent struct {
	Parent               string   `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Student              string   `protobuf:"bytes,3,opt,name=student,proto3" json:"student,omitempty"`
	Remark               string   `protobuf:"bytes,4,opt,name=remark,proto3" json:"remark,omitempty"`
	Operator             string   `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
	Status               uint32   `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqClassStudent) Reset()         { *m = ReqClassStudent{} }
func (m *ReqClassStudent) String() string { return proto.CompactTextString(m) }
func (*ReqClassStudent) ProtoMessage()    {}
func (*ReqClassStudent) Descriptor() ([]byte, []int) {
	return fileDescriptor_49492fe1160d7d70, []int{5}
}

func (m *ReqClassStudent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqClassStudent.Unmarshal(m, b)
}
func (m *ReqClassStudent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqClassStudent.Marshal(b, m, deterministic)
}
func (m *ReqClassStudent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqClassStudent.Merge(m, src)
}
func (m *ReqClassStudent) XXX_Size() int {
	return xxx_messageInfo_ReqClassStudent.Size(m)
}
func (m *ReqClassStudent) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqClassStudent.DiscardUnknown(m)
}

var xxx_messageInfo_ReqClassStudent proto.InternalMessageInfo

func (m *ReqClassStudent) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ReqClassStudent) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqClassStudent) GetStudent() string {
	if m != nil {
		return m.Student
	}
	return ""
}

func (m *ReqClassStudent) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqClassStudent) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqClassStudent) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type ReqClassTeacher struct {
	School               string   `protobuf:"bytes,1,opt,name=school,proto3" json:"school,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Teacher              string   `protobuf:"bytes,3,opt,name=teacher,proto3" json:"teacher,omitempty"`
	Operator             string   `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqClassTeacher) Reset()         { *m = ReqClassTeacher{} }
func (m *ReqClassTeacher) String() string { return proto.CompactTextString(m) }
func (*ReqClassTeacher) ProtoMessage()    {}
func (*ReqClassTeacher) Descriptor() ([]byte, []int) {
	return fileDescriptor_49492fe1160d7d70, []int{6}
}

func (m *ReqClassTeacher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqClassTeacher.Unmarshal(m, b)
}
func (m *ReqClassTeacher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqClassTeacher.Marshal(b, m, deterministic)
}
func (m *ReqClassTeacher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqClassTeacher.Merge(m, src)
}
func (m *ReqClassTeacher) XXX_Size() int {
	return xxx_messageInfo_ReqClassTeacher.Size(m)
}
func (m *ReqClassTeacher) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqClassTeacher.DiscardUnknown(m)
}

var xxx_messageInfo_ReqClassTeacher proto.InternalMessageInfo

func (m *ReqClassTeacher) GetSchool() string {
	if m != nil {
		return m.School
	}
	return ""
}

func (m *ReqClassTeacher) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqClassTeacher) GetTeacher() string {
	if m != nil {
		return m.Teacher
	}
	return ""
}

func (m *ReqClassTeacher) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type ReplyClassInfo struct {
	Info                 *ClassInfo   `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyClassInfo) Reset()         { *m = ReplyClassInfo{} }
func (m *ReplyClassInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyClassInfo) ProtoMessage()    {}
func (*ReplyClassInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_49492fe1160d7d70, []int{7}
}

func (m *ReplyClassInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyClassInfo.Unmarshal(m, b)
}
func (m *ReplyClassInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyClassInfo.Marshal(b, m, deterministic)
}
func (m *ReplyClassInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyClassInfo.Merge(m, src)
}
func (m *ReplyClassInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyClassInfo.Size(m)
}
func (m *ReplyClassInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyClassInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyClassInfo proto.InternalMessageInfo

func (m *ReplyClassInfo) GetInfo() *ClassInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ReplyClassInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReplyClassStudents struct {
	Status               *ReplyStatus  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Students             []*MemberInfo `protobuf:"bytes,2,rep,name=students,proto3" json:"students,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ReplyClassStudents) Reset()         { *m = ReplyClassStudents{} }
func (m *ReplyClassStudents) String() string { return proto.CompactTextString(m) }
func (*ReplyClassStudents) ProtoMessage()    {}
func (*ReplyClassStudents) Descriptor() ([]byte, []int) {
	return fileDescriptor_49492fe1160d7d70, []int{8}
}

func (m *ReplyClassStudents) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyClassStudents.Unmarshal(m, b)
}
func (m *ReplyClassStudents) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyClassStudents.Marshal(b, m, deterministic)
}
func (m *ReplyClassStudents) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyClassStudents.Merge(m, src)
}
func (m *ReplyClassStudents) XXX_Size() int {
	return xxx_messageInfo_ReplyClassStudents.Size(m)
}
func (m *ReplyClassStudents) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyClassStudents.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyClassStudents proto.InternalMessageInfo

func (m *ReplyClassStudents) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyClassStudents) GetStudents() []*MemberInfo {
	if m != nil {
		return m.Students
	}
	return nil
}

type ReplyClassList struct {
	Total                uint32       `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Page                 uint32       `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Pages                uint32       `protobuf:"varint,3,opt,name=pages,proto3" json:"pages,omitempty"`
	Number               uint32       `protobuf:"varint,4,opt,name=number,proto3" json:"number,omitempty"`
	List                 []*ClassInfo `protobuf:"bytes,5,rep,name=list,proto3" json:"list,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyClassList) Reset()         { *m = ReplyClassList{} }
func (m *ReplyClassList) String() string { return proto.CompactTextString(m) }
func (*ReplyClassList) ProtoMessage()    {}
func (*ReplyClassList) Descriptor() ([]byte, []int) {
	return fileDescriptor_49492fe1160d7d70, []int{9}
}

func (m *ReplyClassList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyClassList.Unmarshal(m, b)
}
func (m *ReplyClassList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyClassList.Marshal(b, m, deterministic)
}
func (m *ReplyClassList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyClassList.Merge(m, src)
}
func (m *ReplyClassList) XXX_Size() int {
	return xxx_messageInfo_ReplyClassList.Size(m)
}
func (m *ReplyClassList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyClassList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyClassList proto.InternalMessageInfo

func (m *ReplyClassList) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReplyClassList) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ReplyClassList) GetPages() uint32 {
	if m != nil {
		return m.Pages
	}
	return 0
}

func (m *ReplyClassList) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *ReplyClassList) GetList() []*ClassInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *ReplyClassList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*MemberInfo)(nil), "omo.msp.school.MemberInfo")
	proto.RegisterType((*ClassInfo)(nil), "omo.msp.school.ClassInfo")
	proto.RegisterType((*ReqClassAdd)(nil), "omo.msp.school.ReqClassAdd")
	proto.RegisterType((*ReqClassUpdate)(nil), "omo.msp.school.ReqClassUpdate")
	proto.RegisterType((*ReqClassMaster)(nil), "omo.msp.school.ReqClassMaster")
	proto.RegisterType((*ReqClassStudent)(nil), "omo.msp.school.ReqClassStudent")
	proto.RegisterType((*ReqClassTeacher)(nil), "omo.msp.school.ReqClassTeacher")
	proto.RegisterType((*ReplyClassInfo)(nil), "omo.msp.school.ReplyClassInfo")
	proto.RegisterType((*ReplyClassStudents)(nil), "omo.msp.school.ReplyClassStudents")
	proto.RegisterType((*ReplyClassList)(nil), "omo.msp.school.ReplyClassList")
}

func init() {
	proto.RegisterFile("proto/school/classes.proto", fileDescriptor_49492fe1160d7d70)
}

var fileDescriptor_49492fe1160d7d70 = []byte{
	// 869 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x41, 0x8f, 0xdb, 0x44,
	0x14, 0x5e, 0x3b, 0xd9, 0x24, 0x9e, 0x6c, 0xb2, 0x95, 0x85, 0xd0, 0x6c, 0x0a, 0x25, 0xf2, 0x69,
	0x2f, 0xa4, 0xd2, 0x56, 0xe2, 0xc2, 0x29, 0xad, 0x20, 0x54, 0x34, 0x2a, 0x72, 0x00, 0x21, 0x2e,
	0xd5, 0xac, 0xfd, 0xba, 0x6b, 0x88, 0x3d, 0xee, 0xcc, 0xb8, 0x28, 0x47, 0x7e, 0x03, 0x47, 0x4e,
	0xfd, 0x3f, 0xfc, 0x28, 0xf4, 0x66, 0xc6, 0x76, 0x12, 0xe2, 0x60, 0x96, 0x9c, 0xe2, 0x6f, 0xde,
	0x9b, 0xef, 0xbd, 0xf9, 0xde, 0x7b, 0x33, 0x21, 0x93, 0x5c, 0x70, 0xc5, 0x9f, 0xca, 0xe8, 0x9e,
	0xf3, 0xf5, 0xd3, 0x68, 0xcd, 0xa4, 0x04, 0x39, 0xd3, 0x8b, 0xfe, 0x98, 0xa7, 0x7c, 0x96, 0xca,
	0x7c, 0x66, 0xac, 0x93, 0xab, 0x5d, 0x5f, 0x9e, 0xa6, 0x3c, 0x33, 0xae, 0xc1, 0x3d, 0x21, 0x4b,
	0x48, 0x6f, 0x41, 0xbc, 0xcc, 0xde, 0x72, 0xff, 0x11, 0xe9, 0x14, 0x49, 0x4c, 0x9d, 0xa9, 0x73,
	0xed, 0x85, 0xf8, 0xe9, 0x53, 0xd2, 0x97, 0xaa, 0x88, 0x21, 0x53, 0xd4, 0xd5, 0xab, 0x25, 0xf4,
	0x3f, 0x26, 0x3d, 0xa9, 0x98, 0x2a, 0x24, 0xed, 0x4c, 0x9d, 0xeb, 0x51, 0x68, 0x11, 0xae, 0x0b,
	0x48, 0x99, 0xf8, 0x95, 0x76, 0xf5, 0x06, 0x8b, 0x82, 0x3f, 0x3b, 0xc4, 0x7b, 0x81, 0x69, 0xea,
	0x48, 0x63, 0xe2, 0xda, 0x40, 0xdd, 0xd0, 0x4d, 0xe2, 0x32, 0xb2, 0x5b, 0x47, 0xf6, 0x49, 0x37,
	0x63, 0x29, 0x68, 0x76, 0x2f, 0xd4, 0xdf, 0x98, 0x4d, 0x24, 0x80, 0x29, 0x88, 0x35, 0x79, 0x37,
	0x2c, 0x21, 0x5a, 0x8a, 0x3c, 0xd6, 0x96, 0x73, 0x63, 0xb1, 0xb0, 0xda, 0xc3, 0x05, 0xed, 0x99,
	0x13, 0x58, 0xe8, 0x4f, 0xc8, 0x80, 0xe7, 0x20, 0xb4, 0xa9, 0xaf, 0x4d, 0x15, 0xf6, 0x3f, 0x22,
	0xe7, 0xfc, 0xb7, 0x0c, 0x04, 0x1d, 0x68, 0x83, 0x01, 0xb8, 0x0a, 0x99, 0xe0, 0x6b, 0xea, 0x99,
	0x55, 0x0d, 0x70, 0xf5, 0x4e, 0xb0, 0x18, 0x28, 0xd1, 0x42, 0x18, 0x80, 0xf9, 0xab, 0x4d, 0x0e,
	0x74, 0xa8, 0x17, 0xf5, 0x37, 0x9e, 0x3a, 0xe3, 0xf4, 0x42, 0xaf, 0xb8, 0x19, 0x47, 0xad, 0x52,
	0x26, 0x15, 0x08, 0x3a, 0x32, 0x5a, 0x19, 0xe4, 0x7f, 0x42, 0x3c, 0x26, 0x65, 0x22, 0x15, 0xcb,
	0x14, 0x1d, 0x6b, 0x53, 0xbd, 0x80, 0x79, 0x2b, 0x60, 0xd1, 0x3d, 0x08, 0x49, 0x2f, 0xa7, 0x1d,
	0xcc, 0xbb, 0xc4, 0xfe, 0x17, 0x64, 0x60, 0x0b, 0x24, 0xe9, 0xa3, 0x69, 0xe7, 0x7a, 0x78, 0x33,
	0x99, 0xed, 0x76, 0xc3, 0xac, 0xae, 0x77, 0x58, 0xf9, 0x06, 0x7f, 0x38, 0x64, 0x18, 0xc2, 0x3b,
	0x5d, 0xa0, 0x79, 0x1c, 0xe3, 0x99, 0x64, 0x04, 0x19, 0xd8, 0x5e, 0x30, 0xa0, 0xaa, 0x89, 0xbb,
	0x55, 0x93, 0x4a, 0x93, 0xce, 0x9e, 0x26, 0x11, 0x2f, 0x32, 0xa5, 0xeb, 0x34, 0x0a, 0x0d, 0xd8,
	0x51, 0xfc, 0x7c, 0x4f, 0xf1, 0x52, 0xaf, 0x5e, 0xad, 0x57, 0xf0, 0x0b, 0x19, 0x97, 0x49, 0xfd,
	0xa0, 0xcb, 0x89, 0x8a, 0xe5, 0x4c, 0x60, 0x3b, 0x9a, 0xc4, 0x2c, 0x6a, 0xd9, 0x3f, 0xdb, 0xf1,
	0xbb, 0xbb, 0xf1, 0x83, 0xbc, 0x8e, 0xb5, 0x34, 0x55, 0x68, 0x1f, 0x8b, 0x92, 0xbe, 0xad, 0x80,
	0x0d, 0x57, 0xc2, 0xa3, 0x11, 0x3f, 0x38, 0xe4, 0xb2, 0x0c, 0xb9, 0xaa, 0xa7, 0xaa, 0x7d, 0xcc,
	0x72, 0x32, 0x3b, 0xff, 0x98, 0xcc, 0x43, 0x13, 0x78, 0x54, 0xfd, 0x7a, 0x9a, 0x7b, 0xdb, 0xd3,
	0x1c, 0xbc, 0xab, 0x53, 0xfc, 0xde, 0x1e, 0x09, 0x5d, 0x75, 0x27, 0x95, 0x29, 0x1a, 0x74, 0x32,
	0x59, 0x14, 0x16, 0x22, 0x5f, 0x6f, 0xea, 0xcb, 0xe2, 0x73, 0xd2, 0x4d, 0xb2, 0xb7, 0x5c, 0xc7,
	0x1b, 0xde, 0x5c, 0xed, 0x37, 0x74, 0xe5, 0x18, 0x6a, 0x37, 0xff, 0x59, 0x75, 0x16, 0x57, 0x6f,
	0x78, 0xbc, 0xbf, 0x41, 0xd3, 0xaf, 0xb4, 0x4b, 0x75, 0xd0, 0xdf, 0x1d, 0xe2, 0xd7, 0x61, 0x6d,
	0x39, 0xe4, 0x16, 0x97, 0xd3, 0x9a, 0x6b, 0x67, 0x08, 0xdd, 0xff, 0x30, 0x84, 0x7f, 0x39, 0xdb,
	0x47, 0x7f, 0x95, 0x48, 0x85, 0x73, 0xa4, 0xb8, 0x62, 0x46, 0xeb, 0x51, 0x68, 0x00, 0xf6, 0x76,
	0xce, 0xee, 0xcc, 0x1c, 0x8e, 0x42, 0xfd, 0x8d, 0x9e, 0xf8, 0x5b, 0x5e, 0xc7, 0x06, 0x60, 0xb1,
	0xb2, 0x02, 0x43, 0xd9, 0x41, 0xb4, 0x08, 0x25, 0x5d, 0x27, 0x52, 0xd1, 0x73, 0x9d, 0xde, 0x31,
	0x49, 0xd1, 0x6d, 0x4b, 0x86, 0x5e, 0x6b, 0x19, 0x6e, 0x3e, 0x78, 0x64, 0xfc, 0xc2, 0x3c, 0x4c,
	0x2b, 0x10, 0xef, 0x93, 0x08, 0xfc, 0x05, 0xe9, 0xcd, 0xe3, 0xf8, 0x75, 0x06, 0xfe, 0x01, 0x86,
	0xea, 0xf6, 0x99, 0x3c, 0x39, 0x48, 0x5f, 0xa9, 0x12, 0x9c, 0x21, 0xd1, 0x02, 0x54, 0x13, 0x51,
	0x01, 0x52, 0x61, 0xf6, 0xc7, 0x88, 0xd0, 0x1e, 0x9c, 0xf9, 0xdf, 0x90, 0xfe, 0x02, 0x94, 0xd6,
	0xfa, 0xa1, 0x4c, 0x36, 0xa5, 0x97, 0x64, 0xb0, 0x00, 0x35, 0x17, 0x82, 0x6d, 0x1a, 0xa9, 0xd0,
	0xb5, 0x05, 0xd5, 0x2b, 0x32, 0x5c, 0x80, 0x7a, 0xbe, 0xf9, 0x3a, 0x59, 0xe3, 0x45, 0xd4, 0xc4,
	0xf6, 0x1d, 0xbb, 0x83, 0x76, 0x6c, 0xab, 0x13, 0xb0, 0x59, 0xc1, 0x96, 0xe4, 0x62, 0x01, 0x0a,
	0x4b, 0x9d, 0x48, 0x95, 0x44, 0x0f, 0xa1, 0xab, 0x36, 0x6b, 0x3a, 0xcf, 0x5c, 0xed, 0x58, 0xcb,
	0x27, 0x4d, 0x4d, 0x61, 0x5c, 0x5a, 0x64, 0xf7, 0x15, 0xf1, 0x42, 0x48, 0xf9, 0x7b, 0xf8, 0xd7,
	0xd6, 0xb8, 0x3a, 0xc8, 0x55, 0x75, 0x85, 0xb7, 0x02, 0x65, 0xdf, 0x81, 0xc6, 0xac, 0x8c, 0xfd,
	0x38, 0xd3, 0xb7, 0xe4, 0x62, 0x05, 0x6a, 0x5e, 0x3d, 0xde, 0xff, 0x8b, 0xec, 0x47, 0x32, 0x9a,
	0xe7, 0x39, 0x64, 0x71, 0xf9, 0x5c, 0x7c, 0xd6, 0xc4, 0x66, 0x1d, 0x26, 0x41, 0xb3, 0x62, 0xe5,
	0x1d, 0x17, 0x9c, 0xf9, 0x3f, 0x91, 0xcb, 0x55, 0x71, 0xab, 0x04, 0x8b, 0xd4, 0x89, 0x99, 0x97,
	0x65, 0xc6, 0xe5, 0xeb, 0xd1, 0xc8, 0x6b, 0x1d, 0x1a, 0x04, 0xb0, 0xad, 0xfc, 0xba, 0x4e, 0xf4,
	0x24, 0x84, 0xcf, 0x3f, 0xfd, 0xf9, 0xf1, 0xf6, 0x9f, 0xe3, 0x2f, 0x79, 0xca, 0xdf, 0xa4, 0x32,
	0x7f, 0x63, 0xe0, 0x6d, 0x4f, 0x1b, 0x9f, 0xfd, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x5b, 0x9b,
	0xdb, 0x6e, 0x0b, 0x00, 0x00,
}
