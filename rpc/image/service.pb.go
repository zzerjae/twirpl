// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/image/service.proto

/*
Package image is a generated protocol buffer package.

It is generated from these files:
	rpc/image/service.proto

It has these top-level messages:
	Search
	Giphy
*/
package image

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

// Ways to search for a Giphy
type Search struct {
	Term string `protobuf:"bytes,1,opt,name=term" json:"term,omitempty"`
}

func (m *Search) Reset()                    { *m = Search{} }
func (m *Search) String() string            { return proto.CompactTextString(m) }
func (*Search) ProtoMessage()               {}
func (*Search) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Search) GetTerm() string {
	if m != nil {
		return m.Term
	}
	return ""
}

// Giphy that was found for search
type Giphy struct {
	Url string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
}

func (m *Giphy) Reset()                    { *m = Giphy{} }
func (m *Giphy) String() string            { return proto.CompactTextString(m) }
func (*Giphy) ProtoMessage()               {}
func (*Giphy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Giphy) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func init() {
	proto.RegisterType((*Search)(nil), "com.rynop.twirpl.image.Search")
	proto.RegisterType((*Giphy)(nil), "com.rynop.twirpl.image.Giphy")
}

func init() { proto.RegisterFile("rpc/image/service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x2a, 0x48, 0xd6,
	0xcf, 0xcc, 0x4d, 0x4c, 0x4f, 0xd5, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4b, 0xce, 0xcf, 0xd5, 0x2b, 0xaa, 0xcc, 0xcb, 0x2f, 0xd0, 0x2b,
	0x29, 0xcf, 0x2c, 0x2a, 0xc8, 0xd1, 0x03, 0xab, 0x52, 0x92, 0xe1, 0x62, 0x0b, 0x4e, 0x4d, 0x2c,
	0x4a, 0xce, 0x10, 0x12, 0xe2, 0x62, 0x29, 0x49, 0x2d, 0xca, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0,
	0x0c, 0x02, 0xb3, 0x95, 0x24, 0xb9, 0x58, 0xdd, 0x33, 0x0b, 0x32, 0x2a, 0x85, 0x04, 0xb8, 0x98,
	0x4b, 0x8b, 0x72, 0xa0, 0x72, 0x20, 0xa6, 0x51, 0x28, 0x17, 0xab, 0x27, 0xc8, 0x04, 0x21, 0x1f,
	0x2e, 0x6e, 0xe7, 0xa2, 0xd4, 0xc4, 0x92, 0x54, 0x88, 0x4a, 0x39, 0x3d, 0xec, 0x36, 0xe9, 0x41,
	0xac, 0x91, 0x92, 0xc5, 0x25, 0x0f, 0xd6, 0xee, 0xc4, 0x1e, 0xc5, 0x0a, 0xe6, 0x26, 0xb1, 0x81,
	0xdd, 0x6d, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x88, 0xe8, 0x30, 0xb1, 0xd2, 0x00, 0x00, 0x00,
}
