// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/school/school.proto

package omo_msp_school

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for SchoolService service

type SchoolService interface {
	AddOne(ctx context.Context, in *ReqSchoolAdd, opts ...client.CallOption) (*ReplySchoolInfo, error)
	GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplySchoolInfo, error)
	GetList(ctx context.Context, in *RequestPage, opts ...client.CallOption) (*ReplySchoolList, error)
	UpdateOne(ctx context.Context, in *ReqSchoolUpdate, opts ...client.CallOption) (*ReplySchoolInfo, error)
	RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error)
	UpdateSubject(ctx context.Context, in *ReqSchoolSubject, opts ...client.CallOption) (*ReplySchoolSubjects, error)
	AppendTeacher(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyList, error)
	SubtractTeacher(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyList, error)
}

type schoolService struct {
	c    client.Client
	name string
}

func NewSchoolService(name string, c client.Client) SchoolService {
	return &schoolService{
		c:    c,
		name: name,
	}
}

func (c *schoolService) AddOne(ctx context.Context, in *ReqSchoolAdd, opts ...client.CallOption) (*ReplySchoolInfo, error) {
	req := c.c.NewRequest(c.name, "SchoolService.AddOne", in)
	out := new(ReplySchoolInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schoolService) GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplySchoolInfo, error) {
	req := c.c.NewRequest(c.name, "SchoolService.GetOne", in)
	out := new(ReplySchoolInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schoolService) GetList(ctx context.Context, in *RequestPage, opts ...client.CallOption) (*ReplySchoolList, error) {
	req := c.c.NewRequest(c.name, "SchoolService.GetList", in)
	out := new(ReplySchoolList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schoolService) UpdateOne(ctx context.Context, in *ReqSchoolUpdate, opts ...client.CallOption) (*ReplySchoolInfo, error) {
	req := c.c.NewRequest(c.name, "SchoolService.UpdateOne", in)
	out := new(ReplySchoolInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schoolService) RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "SchoolService.RemoveOne", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schoolService) UpdateSubject(ctx context.Context, in *ReqSchoolSubject, opts ...client.CallOption) (*ReplySchoolSubjects, error) {
	req := c.c.NewRequest(c.name, "SchoolService.UpdateSubject", in)
	out := new(ReplySchoolSubjects)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schoolService) AppendTeacher(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyList, error) {
	req := c.c.NewRequest(c.name, "SchoolService.AppendTeacher", in)
	out := new(ReplyList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schoolService) SubtractTeacher(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyList, error) {
	req := c.c.NewRequest(c.name, "SchoolService.SubtractTeacher", in)
	out := new(ReplyList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SchoolService service

type SchoolServiceHandler interface {
	AddOne(context.Context, *ReqSchoolAdd, *ReplySchoolInfo) error
	GetOne(context.Context, *RequestInfo, *ReplySchoolInfo) error
	GetList(context.Context, *RequestPage, *ReplySchoolList) error
	UpdateOne(context.Context, *ReqSchoolUpdate, *ReplySchoolInfo) error
	RemoveOne(context.Context, *RequestInfo, *ReplyInfo) error
	UpdateSubject(context.Context, *ReqSchoolSubject, *ReplySchoolSubjects) error
	AppendTeacher(context.Context, *RequestInfo, *ReplyList) error
	SubtractTeacher(context.Context, *RequestInfo, *ReplyList) error
}

func RegisterSchoolServiceHandler(s server.Server, hdlr SchoolServiceHandler, opts ...server.HandlerOption) error {
	type schoolService interface {
		AddOne(ctx context.Context, in *ReqSchoolAdd, out *ReplySchoolInfo) error
		GetOne(ctx context.Context, in *RequestInfo, out *ReplySchoolInfo) error
		GetList(ctx context.Context, in *RequestPage, out *ReplySchoolList) error
		UpdateOne(ctx context.Context, in *ReqSchoolUpdate, out *ReplySchoolInfo) error
		RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error
		UpdateSubject(ctx context.Context, in *ReqSchoolSubject, out *ReplySchoolSubjects) error
		AppendTeacher(ctx context.Context, in *RequestInfo, out *ReplyList) error
		SubtractTeacher(ctx context.Context, in *RequestInfo, out *ReplyList) error
	}
	type SchoolService struct {
		schoolService
	}
	h := &schoolServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&SchoolService{h}, opts...))
}

type schoolServiceHandler struct {
	SchoolServiceHandler
}

func (h *schoolServiceHandler) AddOne(ctx context.Context, in *ReqSchoolAdd, out *ReplySchoolInfo) error {
	return h.SchoolServiceHandler.AddOne(ctx, in, out)
}

func (h *schoolServiceHandler) GetOne(ctx context.Context, in *RequestInfo, out *ReplySchoolInfo) error {
	return h.SchoolServiceHandler.GetOne(ctx, in, out)
}

func (h *schoolServiceHandler) GetList(ctx context.Context, in *RequestPage, out *ReplySchoolList) error {
	return h.SchoolServiceHandler.GetList(ctx, in, out)
}

func (h *schoolServiceHandler) UpdateOne(ctx context.Context, in *ReqSchoolUpdate, out *ReplySchoolInfo) error {
	return h.SchoolServiceHandler.UpdateOne(ctx, in, out)
}

func (h *schoolServiceHandler) RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error {
	return h.SchoolServiceHandler.RemoveOne(ctx, in, out)
}

func (h *schoolServiceHandler) UpdateSubject(ctx context.Context, in *ReqSchoolSubject, out *ReplySchoolSubjects) error {
	return h.SchoolServiceHandler.UpdateSubject(ctx, in, out)
}

func (h *schoolServiceHandler) AppendTeacher(ctx context.Context, in *RequestInfo, out *ReplyList) error {
	return h.SchoolServiceHandler.AppendTeacher(ctx, in, out)
}

func (h *schoolServiceHandler) SubtractTeacher(ctx context.Context, in *RequestInfo, out *ReplyList) error {
	return h.SchoolServiceHandler.SubtractTeacher(ctx, in, out)
}
