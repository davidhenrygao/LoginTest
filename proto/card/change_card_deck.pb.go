// Code generated by protoc-gen-go. DO NOT EDIT.
// source: change_card_deck.proto

package card

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ChangeInfo struct {
	Id               *uint32 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Pos              *uint32 `protobuf:"varint,2,req,name=pos" json:"pos,omitempty"`
	CardDeskIndex    *uint32 `protobuf:"varint,3,opt,name=card_desk_index" json:"card_desk_index,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ChangeInfo) Reset()                    { *m = ChangeInfo{} }
func (m *ChangeInfo) String() string            { return proto.CompactTextString(m) }
func (*ChangeInfo) ProtoMessage()               {}
func (*ChangeInfo) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *ChangeInfo) GetId() uint32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *ChangeInfo) GetPos() uint32 {
	if m != nil && m.Pos != nil {
		return *m.Pos
	}
	return 0
}

func (m *ChangeInfo) GetCardDeskIndex() uint32 {
	if m != nil && m.CardDeskIndex != nil {
		return *m.CardDeskIndex
	}
	return 0
}

type C2SChangeCardDeck struct {
	Change           *ChangeInfo `protobuf:"bytes,1,req,name=change" json:"change,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *C2SChangeCardDeck) Reset()                    { *m = C2SChangeCardDeck{} }
func (m *C2SChangeCardDeck) String() string            { return proto.CompactTextString(m) }
func (*C2SChangeCardDeck) ProtoMessage()               {}
func (*C2SChangeCardDeck) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *C2SChangeCardDeck) GetChange() *ChangeInfo {
	if m != nil {
		return m.Change
	}
	return nil
}

type S2CChangeCardDeck struct {
	Code             *int32      `protobuf:"varint,1,req,name=code" json:"code,omitempty"`
	Change           *ChangeInfo `protobuf:"bytes,2,opt,name=change" json:"change,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *S2CChangeCardDeck) Reset()                    { *m = S2CChangeCardDeck{} }
func (m *S2CChangeCardDeck) String() string            { return proto.CompactTextString(m) }
func (*S2CChangeCardDeck) ProtoMessage()               {}
func (*S2CChangeCardDeck) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *S2CChangeCardDeck) GetCode() int32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return 0
}

func (m *S2CChangeCardDeck) GetChange() *ChangeInfo {
	if m != nil {
		return m.Change
	}
	return nil
}

func init() {
	proto.RegisterType((*ChangeInfo)(nil), "protocol.change_info")
	proto.RegisterType((*C2SChangeCardDeck)(nil), "protocol.c2s_change_card_deck")
	proto.RegisterType((*S2CChangeCardDeck)(nil), "protocol.s2c_change_card_deck")
}

func init() { proto.RegisterFile("change_card_deck.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8d, 0xb1, 0x4a, 0xc0, 0x30,
	0x10, 0x40, 0x69, 0xaa, 0x22, 0x57, 0x8b, 0x10, 0xaa, 0x76, 0x2c, 0x05, 0xa1, 0x8b, 0x09, 0x74,
	0x17, 0xc1, 0x55, 0x47, 0x27, 0x97, 0x10, 0x2f, 0x31, 0x0d, 0xd5, 0x5c, 0x69, 0xaa, 0xe8, 0xdf,
	0x8b, 0x69, 0x85, 0xa2, 0x38, 0x1d, 0x77, 0xdc, 0x7b, 0x0f, 0xce, 0x71, 0xd0, 0xc1, 0x59, 0x85,
	0x7a, 0x36, 0xca, 0x58, 0x1c, 0xc5, 0x34, 0xd3, 0x42, 0xfc, 0x38, 0x0d, 0xa4, 0x97, 0xf6, 0x06,
	0x8a, 0xed, 0xc7, 0x87, 0x67, 0xe2, 0x00, 0xcc, 0x9b, 0x3a, 0x6b, 0x58, 0x57, 0xf2, 0x02, 0xf2,
	0x89, 0x62, 0xcd, 0xd2, 0x72, 0x01, 0xa7, 0x9b, 0x24, 0x8e, 0xca, 0x07, 0x63, 0x3f, 0xea, 0xbc,
	0xc9, 0xba, 0xb2, 0xbd, 0x86, 0x0a, 0xfb, 0xa8, 0x7e, 0x87, 0xf8, 0x25, 0x1c, 0xad, 0xb7, 0x64,
	0x2b, 0xfa, 0x33, 0xf1, 0xd3, 0x14, 0xbb, 0x60, 0x7b, 0x07, 0x55, 0xec, 0xf1, 0x2f, 0x7e, 0x02,
	0x07, 0x48, 0x66, 0x85, 0x0f, 0x77, 0x32, 0xd6, 0x64, 0xff, 0xca, 0x6e, 0xe5, 0xe3, 0x95, 0xf3,
	0xcb, 0xf0, 0xf6, 0x24, 0x90, 0x5e, 0xa5, 0xd1, 0xef, 0xde, 0x0c, 0x36, 0xcc, 0x9f, 0x4e, 0x93,
	0xbc, 0x27, 0xe7, 0xc3, 0x83, 0x8d, 0x8b, 0x4c, 0xa8, 0xfc, 0x4e, 0x7d, 0x05, 0x00, 0x00, 0xff,
	0xff, 0xd2, 0x5f, 0xe6, 0xd5, 0x20, 0x01, 0x00, 0x00,
}
