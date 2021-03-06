// Code generated by protoc-gen-go. DO NOT EDIT.
// source: array_elem.proto

/*
Package common is a generated protocol buffer package.

It is generated from these files:
	array_elem.proto
	cmd.proto
	error.proto
	heartbeat.proto

It has these top-level messages:
	IarrayElem
	UarrayElem
	C2SHeartbeat
	S2CHeartbeat
*/
package common

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

type IarrayElem struct {
	Value            *int32  `protobuf:"varint,1,req,name=value" json:"value,omitempty"`
	Pos              *uint32 `protobuf:"varint,2,req,name=pos" json:"pos,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *IarrayElem) Reset()                    { *m = IarrayElem{} }
func (m *IarrayElem) String() string            { return proto.CompactTextString(m) }
func (*IarrayElem) ProtoMessage()               {}
func (*IarrayElem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IarrayElem) GetValue() int32 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

func (m *IarrayElem) GetPos() uint32 {
	if m != nil && m.Pos != nil {
		return *m.Pos
	}
	return 0
}

type UarrayElem struct {
	Value            *uint32 `protobuf:"varint,1,req,name=value" json:"value,omitempty"`
	Pos              *uint32 `protobuf:"varint,2,req,name=pos" json:"pos,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UarrayElem) Reset()                    { *m = UarrayElem{} }
func (m *UarrayElem) String() string            { return proto.CompactTextString(m) }
func (*UarrayElem) ProtoMessage()               {}
func (*UarrayElem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UarrayElem) GetValue() uint32 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

func (m *UarrayElem) GetPos() uint32 {
	if m != nil && m.Pos != nil {
		return *m.Pos
	}
	return 0
}

func init() {
	proto.RegisterType((*IarrayElem)(nil), "protocol.iarray_elem")
	proto.RegisterType((*UarrayElem)(nil), "protocol.uarray_elem")
}

func init() { proto.RegisterFile("array_elem.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x2c, 0x2a, 0x4a,
	0xac, 0x8c, 0x4f, 0xcd, 0x49, 0xcd, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53,
	0xc9, 0xf9, 0x39, 0x4a, 0x9a, 0x5c, 0xdc, 0x99, 0x08, 0x69, 0x21, 0x5e, 0x2e, 0xd6, 0xb2, 0xc4,
	0x9c, 0xd2, 0x54, 0x09, 0x46, 0x05, 0x26, 0x0d, 0x56, 0x21, 0x6e, 0x2e, 0xe6, 0x82, 0xfc, 0x62,
	0x09, 0x26, 0x05, 0x26, 0x0d, 0x5e, 0x90, 0xd2, 0x52, 0x5c, 0x4a, 0x79, 0x51, 0x94, 0x3a, 0x19,
	0x46, 0xe9, 0xa7, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xa7, 0x24, 0x96,
	0x65, 0xa6, 0x64, 0xa4, 0xe6, 0x15, 0x55, 0xa6, 0x27, 0xe6, 0xeb, 0xfb, 0xe4, 0xa7, 0x67, 0xe6,
	0x85, 0xa4, 0x16, 0x97, 0xe8, 0x83, 0x1d, 0xa1, 0x9f, 0x9c, 0x9f, 0x9b, 0x9b, 0x9f, 0x07, 0x08,
	0x00, 0x00, 0xff, 0xff, 0xbe, 0x82, 0x37, 0x09, 0xa5, 0x00, 0x00, 0x00,
}
