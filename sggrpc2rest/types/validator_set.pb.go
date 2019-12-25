// Code generated by protoc-gen-go. DO NOT EDIT.
// source: validator_set.proto

package types

import (
	fmt "fmt"
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

// Protobuf definition for the Rust struct ValidatorSet.
type ValidatorSet struct {
	ValidatorPublicKeys  []*ValidatorPublicKeys `protobuf:"bytes,1,rep,name=validator_public_keys,json=validatorPublicKeys,proto3" json:"validator_public_keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ValidatorSet) Reset()         { *m = ValidatorSet{} }
func (m *ValidatorSet) String() string { return proto.CompactTextString(m) }
func (*ValidatorSet) ProtoMessage()    {}
func (*ValidatorSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b7856de7657d9d, []int{0}
}

func (m *ValidatorSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidatorSet.Unmarshal(m, b)
}
func (m *ValidatorSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidatorSet.Marshal(b, m, deterministic)
}
func (m *ValidatorSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorSet.Merge(m, src)
}
func (m *ValidatorSet) XXX_Size() int {
	return xxx_messageInfo_ValidatorSet.Size(m)
}
func (m *ValidatorSet) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorSet.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorSet proto.InternalMessageInfo

func (m *ValidatorSet) GetValidatorPublicKeys() []*ValidatorPublicKeys {
	if m != nil {
		return m.ValidatorPublicKeys
	}
	return nil
}

func init() {
	proto.RegisterType((*ValidatorSet)(nil), "types.ValidatorSet")
}

func init() { proto.RegisterFile("validator_set.proto", fileDescriptor_e0b7856de7657d9d) }

var fileDescriptor_e0b7856de7657d9d = []byte{
	// 116 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x4b, 0xcc, 0xc9,
	0x4c, 0x49, 0x2c, 0xc9, 0x2f, 0x8a, 0x2f, 0x4e, 0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x2d, 0xa9, 0x2c, 0x48, 0x2d, 0x96, 0x92, 0x46, 0xc8, 0x15, 0x94, 0x26, 0xe5, 0x64, 0x26,
	0xc7, 0x67, 0xa7, 0x56, 0x16, 0x43, 0xd4, 0x28, 0xc5, 0x71, 0xf1, 0x84, 0xc1, 0xa4, 0x83, 0x53,
	0x4b, 0x84, 0xfc, 0xb8, 0x44, 0xb1, 0x2a, 0x97, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x36, 0x92, 0xd2,
	0x03, 0x9b, 0xa9, 0x07, 0xd7, 0x13, 0x00, 0x56, 0xe2, 0x9d, 0x5a, 0x59, 0x1c, 0x84, 0x70, 0x03,
	0x42, 0x30, 0x89, 0x0d, 0x6c, 0x8d, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x36, 0xc1, 0x3c, 0xca,
	0xa1, 0x00, 0x00, 0x00,
}
