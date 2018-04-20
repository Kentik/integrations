// Code generated by protoc-gen-go.
// source: fdw/result.proto
// DO NOT EDIT!

/*
Package fdw is a generated protocol buffer package.

It is generated from these files:
	fdw/result.proto

It has these top-level messages:
	Result
*/
package fdw

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Result struct {
	RequestId        *string  `protobuf:"bytes,1,req,name=request_id" json:"request_id,omitempty"`
	Query            *string  `protobuf:"bytes,2,req,name=query" json:"query,omitempty"`
	Tn               *string  `protobuf:"bytes,3,req,name=tn" json:"tn,omitempty"`
	UserId           *uint32  `protobuf:"varint,4,req,name=user_id" json:"user_id,omitempty"`
	ServerId         *uint32  `protobuf:"varint,5,req,name=server_id" json:"server_id,omitempty"`
	RemoteHost       *string  `protobuf:"bytes,6,req,name=remote_host" json:"remote_host,omitempty"`
	Aggs             []string `protobuf:"bytes,7,rep,name=aggs" json:"aggs,omitempty"`
	Orderby          []string `protobuf:"bytes,8,rep,name=orderby" json:"orderby,omitempty"`
	Groupby          []string `protobuf:"bytes,9,rep,name=groupby" json:"groupby,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Result) GetRequestId() string {
	if m != nil && m.RequestId != nil {
		return *m.RequestId
	}
	return ""
}

func (m *Result) GetQuery() string {
	if m != nil && m.Query != nil {
		return *m.Query
	}
	return ""
}

func (m *Result) GetTn() string {
	if m != nil && m.Tn != nil {
		return *m.Tn
	}
	return ""
}

func (m *Result) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *Result) GetServerId() uint32 {
	if m != nil && m.ServerId != nil {
		return *m.ServerId
	}
	return 0
}

func (m *Result) GetRemoteHost() string {
	if m != nil && m.RemoteHost != nil {
		return *m.RemoteHost
	}
	return ""
}

func (m *Result) GetAggs() []string {
	if m != nil {
		return m.Aggs
	}
	return nil
}

func (m *Result) GetOrderby() []string {
	if m != nil {
		return m.Orderby
	}
	return nil
}

func (m *Result) GetGroupby() []string {
	if m != nil {
		return m.Groupby
	}
	return nil
}

func init() {
	proto.RegisterType((*Result)(nil), "fdw.Result")
}

func init() { proto.RegisterFile("fdw/result.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x34, 0xcc, 0x4b, 0x0e, 0x82, 0x40,
	0x0c, 0x80, 0xe1, 0xf0, 0x96, 0x2a, 0x51, 0xc7, 0x4d, 0x97, 0xc4, 0x95, 0x2b, 0xbd, 0x8b, 0x17,
	0x20, 0x1a, 0x2a, 0x9a, 0xa8, 0xc5, 0x4e, 0x47, 0xc3, 0x6d, 0x3c, 0xaa, 0xc3, 0x10, 0x77, 0xfd,
	0xbf, 0xb4, 0x85, 0xd5, 0xa5, 0xfd, 0x1c, 0x84, 0xac, 0xbb, 0xeb, 0xbe, 0x17, 0x56, 0x36, 0x89,
	0x97, 0xed, 0x37, 0x82, 0xfc, 0x18, 0xd4, 0x18, 0x00, 0xa1, 0x97, 0x23, 0xab, 0xcd, 0xad, 0xc5,
	0xa8, 0x8e, 0x77, 0xa5, 0xa9, 0x20, 0xf3, 0x22, 0x03, 0xc6, 0x21, 0x01, 0x62, 0x7d, 0x62, 0x12,
	0xe6, 0x25, 0x14, 0xce, 0x92, 0x8c, 0xbb, 0xa9, 0x87, 0xca, 0xac, 0xa1, 0xf4, 0xfd, 0x9e, 0x28,
	0x0b, 0xb4, 0x81, 0xb9, 0xd0, 0x83, 0x95, 0x9a, 0x2b, 0x5b, 0xc5, 0x3c, 0x1c, 0x2e, 0x20, 0x3d,
	0x75, 0x9d, 0xc5, 0xa2, 0x4e, 0xa6, 0x37, 0x2c, 0x2d, 0xc9, 0x79, 0xc0, 0xd9, 0x1f, 0x3a, 0x61,
	0xd7, 0x7b, 0x28, 0x47, 0xf8, 0x05, 0x00, 0x00, 0xff, 0xff, 0x15, 0x9c, 0x4c, 0xa1, 0xba, 0x00,
	0x00, 0x00,
}
