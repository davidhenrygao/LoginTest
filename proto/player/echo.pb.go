// Code generated by protoc-gen-go. DO NOT EDIT.
// source: echo.proto

/*
Package player is a generated protocol buffer package.

It is generated from these files:
	echo.proto
	gm_change_player_property.proto
	kick_out_player.proto
	logout.proto
	update_player_property.proto

It has these top-level messages:
	C2SEcho
	S2CEcho
	C2SGmChangePlayerProperty
	S2CGmChangePlayerProperty
	S2CKickOutPlayer
	C2SLogout
	S2CLogout
	S2CUpdatePlayerProperty
*/
package player

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

type C2SEcho struct {
	Msg              *string `protobuf:"bytes,1,req,name=msg" json:"msg,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *C2SEcho) Reset()                    { *m = C2SEcho{} }
func (m *C2SEcho) String() string            { return proto.CompactTextString(m) }
func (*C2SEcho) ProtoMessage()               {}
func (*C2SEcho) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *C2SEcho) GetMsg() string {
	if m != nil && m.Msg != nil {
		return *m.Msg
	}
	return ""
}

type S2CEcho struct {
	Msg              *string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *S2CEcho) Reset()                    { *m = S2CEcho{} }
func (m *S2CEcho) String() string            { return proto.CompactTextString(m) }
func (*S2CEcho) ProtoMessage()               {}
func (*S2CEcho) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *S2CEcho) GetMsg() string {
	if m != nil && m.Msg != nil {
		return *m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*C2SEcho)(nil), "protocol.c2s_echo")
	proto.RegisterType((*S2CEcho)(nil), "protocol.s2c_echo")
}

func init() { proto.RegisterFile("echo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 124 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4d, 0xce, 0xc8,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0xc9, 0xf9, 0x39, 0x4a, 0xe2, 0x5c,
	0x1c, 0xc9, 0x46, 0xc5, 0xf1, 0x20, 0x39, 0x21, 0x6e, 0x2e, 0xe6, 0xdc, 0xe2, 0x74, 0x09, 0x46,
	0x05, 0x26, 0x0d, 0x4e, 0x90, 0x44, 0xb1, 0x51, 0x32, 0x9a, 0x04, 0xa3, 0x06, 0xa7, 0x93, 0x61,
	0x94, 0x7e, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x4a, 0x62, 0x59,
	0x66, 0x4a, 0x46, 0x6a, 0x5e, 0x51, 0x65, 0x7a, 0x62, 0xbe, 0xbe, 0x4f, 0x7e, 0x7a, 0x66, 0x5e,
	0x48, 0x6a, 0x71, 0x89, 0x3e, 0xd8, 0x02, 0xfd, 0x82, 0x9c, 0xc4, 0xca, 0xd4, 0x22, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x1b, 0x50, 0x12, 0x07, 0x7b, 0x00, 0x00, 0x00,
}
