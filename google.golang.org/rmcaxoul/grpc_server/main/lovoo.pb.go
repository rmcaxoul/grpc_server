// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lovoo.proto

package main

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type OpRequest struct {
	A                    *Number  `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    *Number  `protobuf:"bytes,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpRequest) Reset()         { *m = OpRequest{} }
func (m *OpRequest) String() string { return proto.CompactTextString(m) }
func (*OpRequest) ProtoMessage()    {}
func (*OpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3adae3df9dc6dc07, []int{0}
}

func (m *OpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpRequest.Unmarshal(m, b)
}
func (m *OpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpRequest.Marshal(b, m, deterministic)
}
func (m *OpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpRequest.Merge(m, src)
}
func (m *OpRequest) XXX_Size() int {
	return xxx_messageInfo_OpRequest.Size(m)
}
func (m *OpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OpRequest proto.InternalMessageInfo

func (m *OpRequest) GetA() *Number {
	if m != nil {
		return m.A
	}
	return nil
}

func (m *OpRequest) GetB() *Number {
	if m != nil {
		return m.B
	}
	return nil
}

type DivRequest struct {
	A                    *Number        `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    *NonZeroNumber `protobuf:"bytes,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DivRequest) Reset()         { *m = DivRequest{} }
func (m *DivRequest) String() string { return proto.CompactTextString(m) }
func (*DivRequest) ProtoMessage()    {}
func (*DivRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3adae3df9dc6dc07, []int{1}
}

func (m *DivRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DivRequest.Unmarshal(m, b)
}
func (m *DivRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DivRequest.Marshal(b, m, deterministic)
}
func (m *DivRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DivRequest.Merge(m, src)
}
func (m *DivRequest) XXX_Size() int {
	return xxx_messageInfo_DivRequest.Size(m)
}
func (m *DivRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DivRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DivRequest proto.InternalMessageInfo

func (m *DivRequest) GetA() *Number {
	if m != nil {
		return m.A
	}
	return nil
}

func (m *DivRequest) GetB() *NonZeroNumber {
	if m != nil {
		return m.B
	}
	return nil
}

type OpResponse struct {
	C                    *Number  `protobuf:"bytes,1,opt,name=c,proto3" json:"c,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpResponse) Reset()         { *m = OpResponse{} }
func (m *OpResponse) String() string { return proto.CompactTextString(m) }
func (*OpResponse) ProtoMessage()    {}
func (*OpResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3adae3df9dc6dc07, []int{2}
}

func (m *OpResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpResponse.Unmarshal(m, b)
}
func (m *OpResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpResponse.Marshal(b, m, deterministic)
}
func (m *OpResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpResponse.Merge(m, src)
}
func (m *OpResponse) XXX_Size() int {
	return xxx_messageInfo_OpResponse.Size(m)
}
func (m *OpResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OpResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OpResponse proto.InternalMessageInfo

func (m *OpResponse) GetC() *Number {
	if m != nil {
		return m.C
	}
	return nil
}

type Number struct {
	// Types that are valid to be assigned to N:
	//	*Number_Int
	//	*Number_Float
	N                    isNumber_N `protobuf_oneof:"n"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Number) Reset()         { *m = Number{} }
func (m *Number) String() string { return proto.CompactTextString(m) }
func (*Number) ProtoMessage()    {}
func (*Number) Descriptor() ([]byte, []int) {
	return fileDescriptor_3adae3df9dc6dc07, []int{3}
}

func (m *Number) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Number.Unmarshal(m, b)
}
func (m *Number) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Number.Marshal(b, m, deterministic)
}
func (m *Number) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Number.Merge(m, src)
}
func (m *Number) XXX_Size() int {
	return xxx_messageInfo_Number.Size(m)
}
func (m *Number) XXX_DiscardUnknown() {
	xxx_messageInfo_Number.DiscardUnknown(m)
}

var xxx_messageInfo_Number proto.InternalMessageInfo

type isNumber_N interface {
	isNumber_N()
}

type Number_Int struct {
	Int int32 `protobuf:"varint,1,opt,name=int,proto3,oneof"`
}

type Number_Float struct {
	Float float32 `protobuf:"fixed32,2,opt,name=float,proto3,oneof"`
}

func (*Number_Int) isNumber_N() {}

func (*Number_Float) isNumber_N() {}

func (m *Number) GetN() isNumber_N {
	if m != nil {
		return m.N
	}
	return nil
}

func (m *Number) GetInt() int32 {
	if x, ok := m.GetN().(*Number_Int); ok {
		return x.Int
	}
	return 0
}

func (m *Number) GetFloat() float32 {
	if x, ok := m.GetN().(*Number_Float); ok {
		return x.Float
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Number) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Number_Int)(nil),
		(*Number_Float)(nil),
	}
}

type NonZeroNumber struct {
	// Types that are valid to be assigned to N:
	//	*NonZeroNumber_Int
	//	*NonZeroNumber_Float
	N                    isNonZeroNumber_N `protobuf_oneof:"n"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *NonZeroNumber) Reset()         { *m = NonZeroNumber{} }
func (m *NonZeroNumber) String() string { return proto.CompactTextString(m) }
func (*NonZeroNumber) ProtoMessage()    {}
func (*NonZeroNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor_3adae3df9dc6dc07, []int{4}
}

func (m *NonZeroNumber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NonZeroNumber.Unmarshal(m, b)
}
func (m *NonZeroNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NonZeroNumber.Marshal(b, m, deterministic)
}
func (m *NonZeroNumber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NonZeroNumber.Merge(m, src)
}
func (m *NonZeroNumber) XXX_Size() int {
	return xxx_messageInfo_NonZeroNumber.Size(m)
}
func (m *NonZeroNumber) XXX_DiscardUnknown() {
	xxx_messageInfo_NonZeroNumber.DiscardUnknown(m)
}

var xxx_messageInfo_NonZeroNumber proto.InternalMessageInfo

type isNonZeroNumber_N interface {
	isNonZeroNumber_N()
}

type NonZeroNumber_Int struct {
	Int int32 `protobuf:"varint,1,opt,name=int,proto3,oneof"`
}

type NonZeroNumber_Float struct {
	Float float32 `protobuf:"fixed32,2,opt,name=float,proto3,oneof"`
}

func (*NonZeroNumber_Int) isNonZeroNumber_N() {}

func (*NonZeroNumber_Float) isNonZeroNumber_N() {}

func (m *NonZeroNumber) GetN() isNonZeroNumber_N {
	if m != nil {
		return m.N
	}
	return nil
}

func (m *NonZeroNumber) GetInt() int32 {
	if x, ok := m.GetN().(*NonZeroNumber_Int); ok {
		return x.Int
	}
	return 0
}

func (m *NonZeroNumber) GetFloat() float32 {
	if x, ok := m.GetN().(*NonZeroNumber_Float); ok {
		return x.Float
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*NonZeroNumber) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*NonZeroNumber_Int)(nil),
		(*NonZeroNumber_Float)(nil),
	}
}

func init() {
	proto.RegisterType((*OpRequest)(nil), "main.OpRequest")
	proto.RegisterType((*DivRequest)(nil), "main.DivRequest")
	proto.RegisterType((*OpResponse)(nil), "main.OpResponse")
	proto.RegisterType((*Number)(nil), "main.Number")
	proto.RegisterType((*NonZeroNumber)(nil), "main.NonZeroNumber")
}

func init() {
	proto.RegisterFile("lovoo.proto", fileDescriptor_3adae3df9dc6dc07)
}

var fileDescriptor_3adae3df9dc6dc07 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x4f, 0x83, 0x40,
	0x10, 0xc5, 0x97, 0x52, 0xda, 0x38, 0xd5, 0x68, 0xd6, 0x44, 0x1b, 0xbc, 0x28, 0x27, 0x8c, 0x11,
	0x92, 0x7a, 0xd1, 0x83, 0x07, 0xb1, 0x87, 0x26, 0xfe, 0x4b, 0x30, 0x5e, 0x7a, 0x31, 0x0b, 0x5d,
	0x09, 0xc9, 0x96, 0xc1, 0x65, 0x21, 0x7e, 0x2d, 0xbf, 0x5e, 0x4f, 0x66, 0x4b, 0x2b, 0x36, 0xf6,
	0xc0, 0x8d, 0xc7, 0x7b, 0xf3, 0x9b, 0xd9, 0xd9, 0x85, 0x81, 0xc0, 0x0a, 0xd1, 0xcb, 0x25, 0x2a,
	0xa4, 0xdd, 0x39, 0x4b, 0x33, 0xfb, 0xb8, 0x62, 0x22, 0x9d, 0x31, 0xc5, 0xfd, 0xf5, 0x47, 0x6d,
	0x3b, 0xf7, 0xb0, 0xf3, 0x92, 0x87, 0xfc, 0xb3, 0xe4, 0x85, 0xa2, 0x36, 0x18, 0x6c, 0x68, 0x9c,
	0x1a, 0xee, 0x60, 0xb4, 0xeb, 0xe9, 0x3a, 0xef, 0xb9, 0x9c, 0x47, 0x5c, 0x86, 0x06, 0xd3, 0x5e,
	0x34, 0xec, 0x6c, 0xf3, 0x22, 0xe7, 0x01, 0x60, 0x9c, 0x56, 0x6d, 0x28, 0x67, 0x0d, 0xe5, 0x70,
	0xe5, 0x61, 0x36, 0xe5, 0x12, 0x1b, 0x98, 0x0b, 0xa0, 0x27, 0x2a, 0x72, 0xcc, 0x0a, 0xae, 0x61,
	0xf1, 0x76, 0x58, 0xec, 0xdc, 0x40, 0xaf, 0x16, 0x94, 0x82, 0x99, 0x66, 0x6a, 0x99, 0xb3, 0x26,
	0x24, 0xd4, 0x82, 0x1e, 0x81, 0xf5, 0x21, 0x90, 0xa9, 0x65, 0xbb, 0xce, 0x84, 0x84, 0xb5, 0x0c,
	0x4c, 0x30, 0x32, 0xe7, 0x0d, 0xf6, 0x36, 0x1a, 0xd3, 0x93, 0x3f, 0x84, 0xa0, 0xbf, 0x08, 0xba,
	0x76, 0xe7, 0x9a, 0xac, 0x51, 0xce, 0x06, 0x2a, 0x80, 0x45, 0xd0, 0x07, 0xeb, 0x96, 0x10, 0x42,
	0x36, 0xb1, 0xa3, 0x6f, 0x03, 0xac, 0x47, 0xbd, 0x7c, 0xea, 0x82, 0x79, 0x37, 0x9b, 0xd1, 0xfd,
	0x7a, 0xe6, 0xdf, 0x15, 0xdb, 0x07, 0xcd, 0x8f, 0xd5, 0x09, 0x5d, 0x30, 0x5f, 0xcb, 0xa8, 0x4d,
	0xf2, 0x1c, 0xba, 0x4f, 0xa5, 0x50, 0xed, 0xa2, 0xe6, 0x38, 0xad, 0xe8, 0xca, 0x68, 0x2e, 0xe7,
	0x7f, 0x34, 0xb8, 0x9c, 0x5e, 0x24, 0x88, 0x89, 0xe0, 0x5e, 0x82, 0x82, 0x65, 0x89, 0x87, 0x32,
	0xf1, 0xe5, 0x3c, 0x66, 0x5f, 0x58, 0x0a, 0x3f, 0x91, 0x79, 0xfc, 0x5e, 0x70, 0x59, 0x71, 0xe9,
	0xeb, 0xca, 0xa8, 0xb7, 0x7c, 0x37, 0x57, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x16, 0x13, 0x85,
	0x8a, 0x65, 0x02, 0x00, 0x00,
}