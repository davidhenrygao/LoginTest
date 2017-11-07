// Code generated by protoc-gen-go. DO NOT EDIT.
// source: load_cards.proto

package card

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type C2SLoadCards struct {
	BeginIndex       *uint32 `protobuf:"varint,1,opt,name=begin_index,def=1" json:"begin_index,omitempty"`
	PageSize         *uint32 `protobuf:"varint,2,opt,name=page_size" json:"page_size,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *C2SLoadCards) Reset()                    { *m = C2SLoadCards{} }
func (m *C2SLoadCards) String() string            { return proto.CompactTextString(m) }
func (*C2SLoadCards) ProtoMessage()               {}
func (*C2SLoadCards) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

const Default_C2SLoadCards_BeginIndex uint32 = 1

func (m *C2SLoadCards) GetBeginIndex() uint32 {
	if m != nil && m.BeginIndex != nil {
		return *m.BeginIndex
	}
	return Default_C2SLoadCards_BeginIndex
}

func (m *C2SLoadCards) GetPageSize() uint32 {
	if m != nil && m.PageSize != nil {
		return *m.PageSize
	}
	return 0
}

type S2CLoadCards struct {
	Code             *int32  `protobuf:"varint,1,req,name=code" json:"code,omitempty"`
	Cards            []*Card `protobuf:"bytes,2,rep,name=cards" json:"cards,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *S2CLoadCards) Reset()                    { *m = S2CLoadCards{} }
func (m *S2CLoadCards) String() string            { return proto.CompactTextString(m) }
func (*S2CLoadCards) ProtoMessage()               {}
func (*S2CLoadCards) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *S2CLoadCards) GetCode() int32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return 0
}

func (m *S2CLoadCards) GetCards() []*Card {
	if m != nil {
		return m.Cards
	}
	return nil
}

func init() {
	proto.RegisterType((*C2SLoadCards)(nil), "protocol.c2s_load_cards")
	proto.RegisterType((*S2CLoadCards)(nil), "protocol.s2c_load_cards")
}

func init() { proto.RegisterFile("load_cards.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0xc9, 0x4f, 0x4c,
	0x89, 0x4f, 0x4e, 0x2c, 0x4a, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53,
	0xc9, 0xf9, 0x39, 0x52, 0x5c, 0x20, 0x61, 0x88, 0xa8, 0x92, 0x35, 0x17, 0x5f, 0xb2, 0x51, 0x71,
	0x3c, 0x42, 0xb5, 0x90, 0x18, 0x17, 0x77, 0x52, 0x6a, 0x7a, 0x66, 0x5e, 0x7c, 0x66, 0x5e, 0x4a,
	0x6a, 0x85, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xaf, 0x15, 0xa3, 0xa1, 0x90, 0x20, 0x17, 0x67, 0x41,
	0x62, 0x7a, 0x6a, 0x7c, 0x71, 0x66, 0x55, 0xaa, 0x04, 0x13, 0x48, 0x54, 0xc9, 0x96, 0x8b, 0xaf,
	0xd8, 0x28, 0x19, 0x59, 0x33, 0x0f, 0x17, 0x4b, 0x72, 0x7e, 0x4a, 0xaa, 0x04, 0xa3, 0x02, 0x93,
	0x06, 0xab, 0x90, 0x2c, 0x17, 0x2b, 0x58, 0x58, 0x82, 0x49, 0x81, 0x59, 0x83, 0xdb, 0x88, 0x4f,
	0x0f, 0xe6, 0x04, 0x3d, 0x90, 0xb0, 0x93, 0x7e, 0x94, 0x6e, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92,
	0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x4a, 0x62, 0x59, 0x66, 0x4a, 0x46, 0x6a, 0x5e, 0x51, 0x65, 0x7a,
	0x62, 0xbe, 0xbe, 0x4f, 0x7e, 0x7a, 0x66, 0x5e, 0x48, 0x6a, 0x71, 0x89, 0x3e, 0x58, 0x8f, 0x3e,
	0x48, 0x03, 0x20, 0x00, 0x00, 0xff, 0xff, 0xc7, 0x2a, 0xa0, 0x25, 0xd5, 0x00, 0x00, 0x00,
}
