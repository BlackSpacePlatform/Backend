// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/models/table.proto

package database

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
	_ "github.com/mwitkow/go-proto-validators"
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

type User struct {
	Id                   uint32               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	DeletedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	FirstName            string               `protobuf:"bytes,5,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName             string               `protobuf:"bytes,6,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Gender               *Gender              `protobuf:"bytes,7,opt,name=gender,proto3" json:"gender,omitempty"`
	Email                string               `protobuf:"bytes,8,opt,name=email,proto3" json:"email,omitempty"`
	Password             string               `protobuf:"bytes,9,opt,name=password,proto3" json:"password,omitempty"`
	BirthDate            string               `protobuf:"bytes,10,opt,name=birthDate,proto3" json:"birthDate,omitempty"`
	IsActive             bool                 `protobuf:"varint,11,opt,name=isActive,proto3" json:"isActive,omitempty"`
	IsOnline             bool                 `protobuf:"varint,12,opt,name=isOnline,proto3" json:"isOnline,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_a14d0565a7da4bc3, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *User) GetDeletedAt() *timestamp.Timestamp {
	if m != nil {
		return m.DeletedAt
	}
	return nil
}

func (m *User) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetGender() *Gender {
	if m != nil {
		return m.Gender
	}
	return nil
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetBirthDate() string {
	if m != nil {
		return m.BirthDate
	}
	return ""
}

func (m *User) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

func (m *User) GetIsOnline() bool {
	if m != nil {
		return m.IsOnline
	}
	return false
}

func init() {
	proto.RegisterType((*User)(nil), "User")
}

func init() {
	proto.RegisterFile("proto/models/table.proto", fileDescriptor_a14d0565a7da4bc3)
}

var fileDescriptor_a14d0565a7da4bc3 = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x41, 0x6f, 0xd4, 0x30,
	0x10, 0x85, 0x95, 0x6d, 0x9b, 0x66, 0xdd, 0xc2, 0x21, 0xe2, 0xe0, 0x5d, 0x21, 0xb1, 0xea, 0x29,
	0x97, 0x8d, 0x25, 0x90, 0x40, 0xf4, 0xb6, 0x15, 0x12, 0x9c, 0x40, 0x8a, 0xe0, 0xc2, 0x05, 0x39,
	0xf1, 0x6c, 0x6a, 0x61, 0x67, 0x22, 0x7b, 0xb6, 0xcb, 0x5f, 0x6b, 0x7f, 0x00, 0xbf, 0x0b, 0xc5,
	0x71, 0xd2, 0x72, 0xea, 0x2d, 0xef, 0xbd, 0xf9, 0xe6, 0x4d, 0xcc, 0x78, 0xef, 0x90, 0x50, 0x58,
	0x54, 0x60, 0xbc, 0x20, 0x59, 0x1b, 0x28, 0x83, 0xb5, 0xbe, 0x6e, 0x35, 0xdd, 0x1e, 0xea, 0xb2,
	0x41, 0x2b, 0x74, 0xb7, 0xc7, 0xda, 0xe0, 0x1f, 0xec, 0xa1, 0x13, 0x21, 0x6e, 0xb6, 0x2d, 0x74,
	0xdb, 0x16, 0x9d, 0x15, 0xd8, 0x93, 0xc6, 0xce, 0x8b, 0x41, 0x44, 0xf6, 0xfd, 0x13, 0xd6, 0x1e,
	0x35, 0xfd, 0xc6, 0xa3, 0x68, 0x71, 0x1b, 0xc2, 0xed, 0x9d, 0x34, 0x5a, 0x49, 0x42, 0xe7, 0xc5,
	0xfc, 0x19, 0xb9, 0x0f, 0x4f, 0xb8, 0xe0, 0xd4, 0x87, 0xbd, 0xf0, 0xae, 0x11, 0x2d, 0x62, 0x6b,
	0xe0, 0xd1, 0x23, 0x6d, 0xc1, 0x93, 0xb4, 0x7d, 0x04, 0x57, 0xff, 0xfd, 0x46, 0x83, 0xd6, 0x62,
	0x37, 0x46, 0x57, 0x7f, 0x4f, 0xd8, 0xe9, 0x0f, 0x0f, 0x2e, 0xbf, 0x62, 0x0b, 0xad, 0x78, 0xb2,
	0x49, 0x8a, 0x17, 0x37, 0xf9, 0xc3, 0xfd, 0xea, 0x25, 0xbb, 0xcc, 0x53, 0x0f, 0x4e, 0x4b, 0x53,
	0x24, 0x5f, 0x92, 0x6a, 0xa1, 0x55, 0xfe, 0x91, 0xb1, 0xc6, 0x81, 0x24, 0x50, 0xbf, 0x24, 0xf1,
	0xc5, 0x26, 0x29, 0x2e, 0xde, 0xae, 0xcb, 0xb1, 0xbd, 0x9c, 0xda, 0xcb, 0xef, 0x53, 0x7b, 0xb5,
	0x8c, 0xd3, 0x3b, 0x1a, 0x50, 0x05, 0x06, 0x22, 0x7a, 0xf2, 0x3c, 0x1a, 0xa7, 0x47, 0xf4, 0xd0,
	0xab, 0xa9, 0xf5, 0xf4, 0x79, 0x34, 0x4e, 0xef, 0x28, 0x7f, 0xcd, 0x96, 0x7b, 0xed, 0x3c, 0x7d,
	0x95, 0x16, 0xf8, 0xd9, 0x26, 0x29, 0x96, 0xd5, 0xa3, 0x91, 0xaf, 0x59, 0x66, 0x64, 0x0c, 0xd3,
	0x10, 0xce, 0x3a, 0x7f, 0xc3, 0xd2, 0x16, 0x3a, 0x05, 0x8e, 0x9f, 0x87, 0xc2, 0xf3, 0xf2, 0x73,
	0x90, 0x55, 0xb4, 0xf3, 0x57, 0xec, 0x0c, 0xac, 0xd4, 0x86, 0x67, 0x81, 0x1c, 0xc5, 0xb0, 0xb2,
	0x97, 0xde, 0x1f, 0xd1, 0x29, 0xbe, 0x1c, 0x57, 0x4e, 0x7a, 0x38, 0xa6, 0xd6, 0x8e, 0x6e, 0x3f,
	0x49, 0x02, 0xce, 0xc6, 0x63, 0x66, 0x63, 0x20, 0xb5, 0xdf, 0x35, 0xa4, 0xef, 0x80, 0x5f, 0x6c,
	0x92, 0x22, 0xab, 0x66, 0x3d, 0x66, 0xdf, 0x3a, 0xa3, 0x3b, 0xe0, 0x97, 0x53, 0x36, 0xea, 0xeb,
	0xf4, 0xe1, 0x7e, 0xb5, 0xc8, 0x92, 0x1b, 0xf6, 0x33, 0x53, 0x92, 0x64, 0x2d, 0x3d, 0xd4, 0x69,
	0x78, 0x95, 0x77, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xce, 0x19, 0x3a, 0x3e, 0xbf, 0x02, 0x00,
	0x00,
}
