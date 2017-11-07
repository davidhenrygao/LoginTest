// Code generated by protoc-gen-go. DO NOT EDIT.
// source: launch.proto

package login

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type C2SLaunch struct {
	//
	// index begin from 1, and increase every re-launch.
	// hashkey(a) transfer a to 8 byte key.
	//
	// etoken = base64(openid)+"@"+base64(subid)+":"+base64(index)
	// hmac = base64(hmac64_md5(hashkey(etoken), secret))
	// signature = etoken+":"+hmac
	Signature        *string `protobuf:"bytes,1,req,name=signature" json:"signature,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *C2SLaunch) Reset()                    { *m = C2SLaunch{} }
func (m *C2SLaunch) String() string            { return proto.CompactTextString(m) }
func (*C2SLaunch) ProtoMessage()               {}
func (*C2SLaunch) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *C2SLaunch) GetSignature() string {
	if m != nil && m.Signature != nil {
		return *m.Signature
	}
	return ""
}

type S2CLaunch struct {
	Code             *int32               `protobuf:"varint,1,req,name=code" json:"code,omitempty"`
	Player           *S2CLaunchPlayerInfo `protobuf:"bytes,2,opt,name=player" json:"player,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *S2CLaunch) Reset()                    { *m = S2CLaunch{} }
func (m *S2CLaunch) String() string            { return proto.CompactTextString(m) }
func (*S2CLaunch) ProtoMessage()               {}
func (*S2CLaunch) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *S2CLaunch) GetCode() int32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return 0
}

func (m *S2CLaunch) GetPlayer() *S2CLaunchPlayerInfo {
	if m != nil {
		return m.Player
	}
	return nil
}

type S2CLaunchPlayerInfo struct {
	Id               *uint64 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Name             *string `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	Level            *uint32 `protobuf:"varint,3,opt,name=level" json:"level,omitempty"`
	Exp              *uint32 `protobuf:"varint,4,opt,name=exp" json:"exp,omitempty"`
	Gold             *uint32 `protobuf:"varint,5,opt,name=gold" json:"gold,omitempty"`
	Headimgurl       *string `protobuf:"bytes,6,opt,name=headimgurl" json:"headimgurl,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *S2CLaunchPlayerInfo) Reset()                    { *m = S2CLaunchPlayerInfo{} }
func (m *S2CLaunchPlayerInfo) String() string            { return proto.CompactTextString(m) }
func (*S2CLaunchPlayerInfo) ProtoMessage()               {}
func (*S2CLaunchPlayerInfo) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1, 0} }

func (m *S2CLaunchPlayerInfo) GetId() uint64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *S2CLaunchPlayerInfo) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *S2CLaunchPlayerInfo) GetLevel() uint32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *S2CLaunchPlayerInfo) GetExp() uint32 {
	if m != nil && m.Exp != nil {
		return *m.Exp
	}
	return 0
}

func (m *S2CLaunchPlayerInfo) GetGold() uint32 {
	if m != nil && m.Gold != nil {
		return *m.Gold
	}
	return 0
}

func (m *S2CLaunchPlayerInfo) GetHeadimgurl() string {
	if m != nil && m.Headimgurl != nil {
		return *m.Headimgurl
	}
	return ""
}

func init() {
	proto.RegisterType((*C2SLaunch)(nil), "protocol.c2s_launch")
	proto.RegisterType((*S2CLaunch)(nil), "protocol.s2c_launch")
	proto.RegisterType((*S2CLaunchPlayerInfo)(nil), "protocol.s2c_launch.player_info")
}

func init() { proto.RegisterFile("launch.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0xb1, 0x6e, 0x83, 0x30,
	0x10, 0x86, 0x65, 0x02, 0x51, 0x39, 0xc8, 0x50, 0x4f, 0x56, 0x97, 0x5a, 0x99, 0x98, 0x4c, 0xc4,
	0x23, 0x74, 0xee, 0xd8, 0xa9, 0x4b, 0xe4, 0xda, 0x57, 0x63, 0xc9, 0xd8, 0xc8, 0x40, 0xd4, 0x3c,
	0x57, 0x5f, 0xb0, 0xc2, 0x08, 0x35, 0xd3, 0xe9, 0xd3, 0xff, 0xdf, 0x77, 0x07, 0xb5, 0x93, 0x8b,
	0x57, 0xbd, 0x18, 0x63, 0x98, 0x03, 0x7d, 0x4a, 0x43, 0x05, 0x77, 0x7e, 0x05, 0x50, 0xdd, 0x74,
	0xdd, 0x52, 0xfa, 0x0c, 0xe5, 0x64, 0x8d, 0x97, 0xf3, 0x12, 0x91, 0x11, 0x9e, 0x35, 0xe5, 0xf9,
	0x97, 0x00, 0x4c, 0x9d, 0xda, 0x1b, 0x35, 0xe4, 0x2a, 0xe8, 0x2d, 0x2c, 0xe8, 0x05, 0x8e, 0xa3,
	0x93, 0x77, 0x8c, 0x2c, 0xe3, 0xa4, 0xa9, 0x3a, 0x2e, 0x76, 0xb1, 0xf8, 0xdf, 0x11, 0x5b, 0xe5,
	0x6a, 0xfd, 0x77, 0x78, 0x41, 0xa8, 0x1e, 0x90, 0x02, 0x64, 0x56, 0x27, 0x59, 0xbe, 0xaa, 0xbd,
	0x1c, 0x90, 0x65, 0xeb, 0x5d, 0x7a, 0x82, 0xc2, 0xe1, 0x0d, 0x1d, 0x3b, 0x70, 0xd2, 0x9c, 0x68,
	0x05, 0x07, 0xfc, 0x19, 0x59, 0x9e, 0xa0, 0x86, 0xdc, 0x04, 0xa7, 0x59, 0x91, 0x88, 0x02, 0xf4,
	0x28, 0xb5, 0x1d, 0xcc, 0x12, 0x1d, 0x3b, 0x72, 0xd2, 0x94, 0x6f, 0x97, 0x4f, 0x61, 0xec, 0xdc,
	0x2f, 0x5f, 0x42, 0x85, 0xa1, 0xd5, 0xf2, 0x66, 0x75, 0x8f, 0x3e, 0xde, 0x8d, 0x0c, 0xed, 0x7b,
	0x30, 0xd6, 0x7f, 0xe0, 0x34, 0xb7, 0xe9, 0xd9, 0xd6, 0xad, 0xfc, 0x17, 0x00, 0x00, 0xff, 0xff,
	0xbe, 0xfe, 0x0d, 0x34, 0x21, 0x01, 0x00, 0x00,
}
