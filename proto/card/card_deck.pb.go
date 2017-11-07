// Code generated by protoc-gen-go. DO NOT EDIT.
// source: card_deck.proto

package card

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CardDeck struct {
	Index            *uint32                 `protobuf:"varint,1,req,name=index" json:"index,omitempty"`
	Elems            []*CardDeckCardDeckElem `protobuf:"bytes,2,rep,name=elems" json:"elems,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *CardDeck) Reset()                    { *m = CardDeck{} }
func (m *CardDeck) String() string            { return proto.CompactTextString(m) }
func (*CardDeck) ProtoMessage()               {}
func (*CardDeck) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *CardDeck) GetIndex() uint32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

func (m *CardDeck) GetElems() []*CardDeckCardDeckElem {
	if m != nil {
		return m.Elems
	}
	return nil
}

type CardDeckCardDeckElem struct {
	Id               *uint32 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Pos              *uint32 `protobuf:"varint,2,opt,name=pos" json:"pos,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CardDeckCardDeckElem) Reset()                    { *m = CardDeckCardDeckElem{} }
func (m *CardDeckCardDeckElem) String() string            { return proto.CompactTextString(m) }
func (*CardDeckCardDeckElem) ProtoMessage()               {}
func (*CardDeckCardDeckElem) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

func (m *CardDeckCardDeckElem) GetId() uint32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *CardDeckCardDeckElem) GetPos() uint32 {
	if m != nil && m.Pos != nil {
		return *m.Pos
	}
	return 0
}

func init() {
	proto.RegisterType((*CardDeck)(nil), "protocol.card_deck")
	proto.RegisterType((*CardDeckCardDeckElem)(nil), "protocol.card_deck.card_deck_elem")
}

func init() { proto.RegisterFile("card_deck.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x4e, 0x2c, 0x4a,
	0x89, 0x4f, 0x49, 0x4d, 0xce, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0xc9,
	0xf9, 0x39, 0x4a, 0x15, 0x5c, 0x9c, 0x70, 0x49, 0x21, 0x5e, 0x2e, 0xd6, 0xcc, 0xbc, 0x94, 0xd4,
	0x0a, 0x09, 0x46, 0x05, 0x26, 0x0d, 0x5e, 0x21, 0x43, 0x2e, 0xd6, 0xd4, 0x9c, 0xd4, 0xdc, 0x62,
	0x09, 0x26, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x25, 0x3d, 0x98, 0x2e, 0x3d, 0x84, 0x79, 0x70, 0x56,
	0x3c, 0x48, 0xa9, 0x94, 0x26, 0x17, 0x1f, 0xaa, 0x88, 0x10, 0x17, 0x17, 0x53, 0x66, 0x0a, 0xd4,
	0x40, 0x6e, 0x2e, 0xe6, 0x82, 0x7c, 0x90, 0x71, 0x8c, 0x1a, 0xbc, 0x4e, 0xfa, 0x51, 0xba, 0xe9,
	0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x29, 0x89, 0x65, 0x99, 0x29, 0x19,
	0xa9, 0x79, 0x45, 0x95, 0xe9, 0x89, 0xf9, 0xfa, 0x3e, 0xf9, 0xe9, 0x99, 0x79, 0x21, 0xa9, 0xc5,
	0x25, 0xfa, 0x60, 0x2b, 0xf5, 0x41, 0x66, 0x02, 0x02, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x62, 0xb2,
	0x18, 0xc6, 0x00, 0x00, 0x00,
}
