// Code generated by protoc-gen-go. DO NOT EDIT.
// source: challenge.proto

/*
Package login is a generated protocol buffer package.

It is generated from these files:
	challenge.proto
	exchangekey.proto
	handshake.proto
	launch.proto
	login.proto

It has these top-level messages:
	S2CChallenge
	C2SExchangekey
	S2CExchangekey
	C2SHandshake
	S2CHandshake
	C2SLaunch
	S2CLaunch
	C2SLogin
	S2CLogin
*/
package login

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

type S2CChallenge struct {
	Challenge        *string `protobuf:"bytes,1,req,name=challenge" json:"challenge,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *S2CChallenge) Reset()                    { *m = S2CChallenge{} }
func (m *S2CChallenge) String() string            { return proto.CompactTextString(m) }
func (*S2CChallenge) ProtoMessage()               {}
func (*S2CChallenge) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *S2CChallenge) GetChallenge() string {
	if m != nil && m.Challenge != nil {
		return *m.Challenge
	}
	return ""
}

func init() {
	proto.RegisterType((*S2CChallenge)(nil), "login.s2c_challenge")
}

func init() { proto.RegisterFile("challenge.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 112 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0xce, 0x48, 0xcc,
	0xc9, 0x49, 0xcd, 0x4b, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcd, 0xc9, 0x4f,
	0xcf, 0xcc, 0x53, 0x52, 0xe2, 0xe2, 0x2d, 0x36, 0x4a, 0x8e, 0x87, 0xcb, 0x0a, 0x09, 0x72, 0x71,
	0xc2, 0x39, 0x12, 0x8c, 0x0a, 0x4c, 0x1a, 0x9c, 0x4e, 0x06, 0x51, 0x7a, 0xe9, 0x99, 0x25, 0x19,
	0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x29, 0x89, 0x65, 0x99, 0x29, 0x19, 0xa9, 0x79, 0x45,
	0x95, 0xe9, 0x89, 0xf9, 0xfa, 0x3e, 0x20, 0x53, 0x42, 0x52, 0x8b, 0x4b, 0xf4, 0xc1, 0xc6, 0xea,
	0x83, 0x4d, 0x05, 0x04, 0x00, 0x00, 0xff, 0xff, 0x68, 0xaf, 0x56, 0x0a, 0x6e, 0x00, 0x00, 0x00,
}