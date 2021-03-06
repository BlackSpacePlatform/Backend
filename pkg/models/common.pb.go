// Code generated by protoc-gen-go. DO NOT EDIT.
// source: models/common.proto

package models

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
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

type Gender struct {
	Male                 bool     `protobuf:"varint,1,opt,name=male,proto3" json:"male,omitempty"`
	Female               bool     `protobuf:"varint,2,opt,name=female,proto3" json:"female,omitempty"`
	Transgender          bool     `protobuf:"varint,3,opt,name=transgender,proto3" json:"transgender,omitempty"`
	GenderNonBinary      bool     `protobuf:"varint,4,opt,name=genderNonBinary,proto3" json:"genderNonBinary,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Gender) Reset()         { *m = Gender{} }
func (m *Gender) String() string { return proto.CompactTextString(m) }
func (*Gender) ProtoMessage()    {}
func (*Gender) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9283273b5f3cbe2, []int{0}
}

func (m *Gender) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Gender.Unmarshal(m, b)
}
func (m *Gender) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Gender.Marshal(b, m, deterministic)
}
func (m *Gender) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Gender.Merge(m, src)
}
func (m *Gender) XXX_Size() int {
	return xxx_messageInfo_Gender.Size(m)
}
func (m *Gender) XXX_DiscardUnknown() {
	xxx_messageInfo_Gender.DiscardUnknown(m)
}

var xxx_messageInfo_Gender proto.InternalMessageInfo

func (m *Gender) GetMale() bool {
	if m != nil {
		return m.Male
	}
	return false
}

func (m *Gender) GetFemale() bool {
	if m != nil {
		return m.Female
	}
	return false
}

func (m *Gender) GetTransgender() bool {
	if m != nil {
		return m.Transgender
	}
	return false
}

func (m *Gender) GetGenderNonBinary() bool {
	if m != nil {
		return m.GenderNonBinary
	}
	return false
}

func init() {
	proto.RegisterType((*Gender)(nil), "Gender")
}

func init() {
	proto.RegisterFile("models/common.proto", fileDescriptor_a9283273b5f3cbe2)
}

var fileDescriptor_a9283273b5f3cbe2 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xb1, 0x4b, 0x43, 0x31,
	0x10, 0xc6, 0x79, 0x5a, 0x42, 0x89, 0x83, 0x10, 0x41, 0x1e, 0x4e, 0xc5, 0xa9, 0x4b, 0x9b, 0xc1,
	0x41, 0x70, 0xec, 0xe2, 0xe6, 0xe0, 0xe8, 0x96, 0xa4, 0xf7, 0x62, 0x20, 0x77, 0x17, 0x92, 0x14,
	0x74, 0xf3, 0x4f, 0x17, 0xef, 0x15, 0x7c, 0x74, 0xfb, 0xbe, 0xdf, 0x97, 0x1f, 0x84, 0xd3, 0x77,
	0xc8, 0x47, 0xc8, 0xcd, 0x06, 0x46, 0x64, 0xda, 0x97, 0xca, 0x9d, 0x1f, 0x5e, 0x62, 0xea, 0x9f,
	0x27, 0xbf, 0x0f, 0x8c, 0x36, 0xd1, 0xc4, 0x3e, 0xf3, 0x17, 0x17, 0x20, 0x2b, 0x73, 0xd8, 0x45,
	0xa0, 0x5d, 0xe4, 0x8a, 0x96, 0x4b, 0x4f, 0x4c, 0xcd, 0xfe, 0x95, 0xb3, 0xfb, 0xbc, 0x70, 0x85,
	0xf8, 0xd3, 0x64, 0x5b, 0x0d, 0x36, 0x32, 0xc7, 0x0c, 0xff, 0xac, 0x27, 0x84, 0xd6, 0x1d, 0x96,
	0x59, 0x7c, 0xfc, 0x19, 0xb4, 0x7a, 0x05, 0x3a, 0x42, 0x35, 0x46, 0xaf, 0xd0, 0x65, 0x18, 0x87,
	0xcd, 0xb0, 0x5d, 0xbf, 0x4b, 0x36, 0xf7, 0x5a, 0x4d, 0x20, 0xf4, 0x4a, 0xe8, 0xb9, 0x99, 0x8d,
	0xbe, 0xe9, 0xd5, 0x51, 0x8b, 0xa2, 0x8e, 0xd7, 0x32, 0x2e, 0x91, 0xd9, 0xea, 0xdb, 0x39, 0xbd,
	0x31, 0x1d, 0x12, 0xb9, 0xfa, 0x3d, 0xae, 0xe4, 0xd5, 0x25, 0x3e, 0xac, 0x3f, 0xd4, 0x7c, 0x0e,
	0xaf, 0xe4, 0x4f, 0x4f, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x5a, 0x6a, 0x38, 0x1f, 0x01,
	0x00, 0x00,
}
