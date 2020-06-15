// Code generated by ICEBERG protoc-gen-go. DO NOT EDIT EXCEPET SERVER VERSION.
// source: hi.proto

/*
Package hi is a generated protocol buffer package.

It is generated from these files:
	hi.proto

It has these top-level messages:
	HiRequest
	HiResponse
*/
package hi

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	"context"

	"github.com/kwins/iceberg/frame"
	"github.com/kwins/iceberg/frame/config"
	"github.com/kwins/iceberg/frame/protocol"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// HelloRequest 请求结构
type HiRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name" xml:"name,omitempty"`
}

func (m *HiRequest) Reset()                    { *m = HiRequest{} }
func (m *HiRequest) String() string            { return proto.CompactTextString(m) }
func (*HiRequest) ProtoMessage()               {}
func (*HiRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HiRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// HelloResponse 响应结构
type HiResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message" xml:"message,omitempty"`
}

func (m *HiResponse) Reset()                    { *m = HiResponse{} }
func (m *HiResponse) String() string            { return proto.CompactTextString(m) }
func (*HiResponse) ProtoMessage()               {}
func (*HiResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HiResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HiRequest)(nil), "hi.HiRequest")
	proto.RegisterType((*HiResponse)(nil), "hi.HiResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.

// Client API for Hi service
// iceberg server version,relation to server uri.
var hiVersion = frame.SrvVersionName[frame.SV1]

// SayHi 定义SayHello方法
func SayHi(ctx frame.Context, in *HiRequest, opts ...frame.CallOption) (*HiResponse, error) {
	task, err := frame.ReadyTask(ctx, "sayhi", "hi", hiVersion, in, opts...)
	if err != nil {
		return nil, err
	}
	back, err := frame.DeliverTo(task)
	if err != nil {
		return nil, err
	}

	var out HiResponse
	if err := protocol.Unpack(back.GetFormat(), back.GetBody(), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// HiServer Server API for Hello service
type HiServer interface {
	SayHi(c frame.Context) error
}

// RegisterHiServer register HiServer with etcd info
func RegisterHiServer(srv HiServer, cfg *config.BaseCfg) {
	frame.RegisterAndServe(&hiServerDesc, srv, cfg)
}

// hi server SayHi handler
func hiSayHiHandler(srv interface{}, ctx frame.Context) error {
	return srv.(HiServer).SayHi(ctx)
}

// hi server describe
var hiServerDesc = frame.ServiceDesc{
	Version:     hiVersion,
	ServiceName: "Hi",
	HandlerType: (*HiServer)(nil),
	Methods: []frame.MethodDesc{
		{
			Allowed:    "false",
			MethodName: "sayhi",
			Handler:    hiSayHiHandler,
		},
	},
	ServiceURI: []string{
		"/services/" + hiVersion + "/hi/hi",
	},
	Metadata: "hi.Hi",
}

func init() { proto.RegisterFile("hi.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc8, 0xc8, 0xd4, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xca, 0xc8, 0x54, 0x92, 0xe7, 0xe2, 0xf4, 0xc8, 0x0c, 0x4a,
	0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95, 0xd4, 0xb8, 0xb8, 0x40, 0x0a, 0x8a, 0x0b, 0xf2, 0xf3,
	0x8a, 0x53, 0x85, 0x24, 0xb8, 0xd8, 0x73, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x61, 0x8a, 0x60, 0x5c,
	0x23, 0x3d, 0x2e, 0x26, 0x8f, 0x4c, 0x21, 0x0d, 0x2e, 0xd6, 0xe0, 0xc4, 0x4a, 0x8f, 0x4c, 0x21,
	0x5e, 0xbd, 0x8c, 0x4c, 0x3d, 0xb8, 0xc9, 0x52, 0x7c, 0x30, 0x2e, 0xc4, 0x1c, 0x25, 0x06, 0x27,
	0x96, 0x28, 0xa6, 0x8c, 0xcc, 0x24, 0x36, 0xb0, 0x4b, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x4a, 0x07, 0x5f, 0xff, 0x95, 0x00, 0x00, 0x00,
}