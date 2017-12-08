// Code generated by protoc-gen-go. DO NOT EDIT.
// source: match_start.proto

package battle

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type C2SMatchStart struct {
	//
	// 战斗类型：
	// 竞技场1v1 : 1
	// 竞技场3v3 : 2
	// ...
	Type             *uint32 `protobuf:"varint,1,req,name=type" json:"type,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *C2SMatchStart) Reset()                    { *m = C2SMatchStart{} }
func (m *C2SMatchStart) String() string            { return proto.CompactTextString(m) }
func (*C2SMatchStart) ProtoMessage()               {}
func (*C2SMatchStart) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *C2SMatchStart) GetType() uint32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

type S2CMatchStart struct {
	Code             *int32 `protobuf:"varint,1,req,name=code" json:"code,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *S2CMatchStart) Reset()                    { *m = S2CMatchStart{} }
func (m *S2CMatchStart) String() string            { return proto.CompactTextString(m) }
func (*S2CMatchStart) ProtoMessage()               {}
func (*S2CMatchStart) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *S2CMatchStart) GetCode() int32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return 0
}

func init() {
	proto.RegisterType((*C2SMatchStart)(nil), "protocol.c2s_match_start")
	proto.RegisterType((*S2CMatchStart)(nil), "protocol.s2c_match_start")
}

func init() { proto.RegisterFile("match_start.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 135 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcc, 0x4d, 0x2c, 0x49,
	0xce, 0x88, 0x2f, 0x2e, 0x49, 0x2c, 0x2a, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00,
	0x53, 0xc9, 0xf9, 0x39, 0x4a, 0xf2, 0x5c, 0xfc, 0xc9, 0x46, 0xc5, 0xf1, 0x48, 0x4a, 0x84, 0x78,
	0xb8, 0x58, 0x4a, 0x2a, 0x0b, 0x52, 0x25, 0x18, 0x15, 0x98, 0x34, 0x78, 0x41, 0x0a, 0x8a, 0x8d,
	0x92, 0xd1, 0x15, 0x24, 0xe7, 0xa7, 0x40, 0x14, 0xb0, 0x3a, 0x19, 0x46, 0xe9, 0xa7, 0x67, 0x96,
	0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xa7, 0x24, 0x96, 0x65, 0xa6, 0x64, 0xa4, 0xe6,
	0x15, 0x55, 0xa6, 0x27, 0xe6, 0xeb, 0xfb, 0xe4, 0xa7, 0x67, 0xe6, 0x85, 0xa4, 0x16, 0x97, 0xe8,
	0x83, 0x2d, 0xd4, 0x4f, 0x4a, 0x2c, 0x29, 0xc9, 0x49, 0x05, 0x04, 0x00, 0x00, 0xff, 0xff, 0xe9,
	0xd6, 0xe6, 0xee, 0x92, 0x00, 0x00, 0x00,
}
