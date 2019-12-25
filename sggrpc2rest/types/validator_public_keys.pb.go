// Code generated by protoc-gen-go. DO NOT EDIT.
// source: validator_public_keys.proto

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

// Protobuf definition for the Rust struct ValidatorPublicKeys
type ValidatorPublicKeys struct {
	// Validator account address
	AccountAddress []byte `protobuf:"bytes,1,opt,name=account_address,json=accountAddress,proto3" json:"account_address,omitempty"`
	// Consensus public key
	ConsensusPublicKey []byte `protobuf:"bytes,2,opt,name=consensus_public_key,json=consensusPublicKey,proto3" json:"consensus_public_key,omitempty"`
	// Validator voting power for consensus
	ConsensusVotingPower uint64 `protobuf:"varint,3,opt,name=consensus_voting_power,json=consensusVotingPower,proto3" json:"consensus_voting_power,omitempty"`
	// Network signing publick key
	NetworkSigningPublicKey []byte `protobuf:"bytes,4,opt,name=network_signing_public_key,json=networkSigningPublicKey,proto3" json:"network_signing_public_key,omitempty"`
	/// Network identity publick key
	NetworkIdentityPublicKey []byte   `protobuf:"bytes,5,opt,name=network_identity_public_key,json=networkIdentityPublicKey,proto3" json:"network_identity_public_key,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *ValidatorPublicKeys) Reset()         { *m = ValidatorPublicKeys{} }
func (m *ValidatorPublicKeys) String() string { return proto.CompactTextString(m) }
func (*ValidatorPublicKeys) ProtoMessage()    {}
func (*ValidatorPublicKeys) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e933cd90ba9d127, []int{0}
}

func (m *ValidatorPublicKeys) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidatorPublicKeys.Unmarshal(m, b)
}
func (m *ValidatorPublicKeys) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidatorPublicKeys.Marshal(b, m, deterministic)
}
func (m *ValidatorPublicKeys) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorPublicKeys.Merge(m, src)
}
func (m *ValidatorPublicKeys) XXX_Size() int {
	return xxx_messageInfo_ValidatorPublicKeys.Size(m)
}
func (m *ValidatorPublicKeys) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorPublicKeys.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorPublicKeys proto.InternalMessageInfo

func (m *ValidatorPublicKeys) GetAccountAddress() []byte {
	if m != nil {
		return m.AccountAddress
	}
	return nil
}

func (m *ValidatorPublicKeys) GetConsensusPublicKey() []byte {
	if m != nil {
		return m.ConsensusPublicKey
	}
	return nil
}

func (m *ValidatorPublicKeys) GetConsensusVotingPower() uint64 {
	if m != nil {
		return m.ConsensusVotingPower
	}
	return 0
}

func (m *ValidatorPublicKeys) GetNetworkSigningPublicKey() []byte {
	if m != nil {
		return m.NetworkSigningPublicKey
	}
	return nil
}

func (m *ValidatorPublicKeys) GetNetworkIdentityPublicKey() []byte {
	if m != nil {
		return m.NetworkIdentityPublicKey
	}
	return nil
}

func init() {
	proto.RegisterType((*ValidatorPublicKeys)(nil), "types.ValidatorPublicKeys")
}

func init() { proto.RegisterFile("validator_public_keys.proto", fileDescriptor_5e933cd90ba9d127) }

var fileDescriptor_5e933cd90ba9d127 = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0xd9, 0xda, 0x7a, 0x08, 0xa2, 0x10, 0x45, 0x17, 0x7b, 0x29, 0x5e, 0xec, 0x49, 0x04,
	0xbd, 0x89, 0x07, 0x8f, 0xe2, 0xa5, 0x54, 0xe8, 0x35, 0xa4, 0xc9, 0x50, 0x42, 0x4b, 0x26, 0x64,
	0x66, 0x5b, 0xf2, 0x77, 0xfc, 0xa5, 0xd2, 0x34, 0xdd, 0xec, 0x75, 0xde, 0xf7, 0xbd, 0x07, 0x23,
	0xa6, 0x7b, 0xbd, 0x73, 0x56, 0x33, 0x46, 0x15, 0xba, 0xf5, 0xce, 0x19, 0xb5, 0x85, 0x44, 0x2f,
	0x21, 0x22, 0xa3, 0x9c, 0x70, 0x0a, 0x40, 0x4f, 0x7f, 0x23, 0x71, 0xbb, 0x3a, 0x63, 0x8b, 0x4c,
	0xfd, 0x40, 0x22, 0xf9, 0x2c, 0x6e, 0xb4, 0x31, 0xd8, 0x79, 0x56, 0xda, 0xda, 0x08, 0x44, 0x6d,
	0x33, 0x6b, 0xe6, 0x57, 0xcb, 0xeb, 0x72, 0xfe, 0x3a, 0x5d, 0xe5, 0xab, 0xb8, 0x33, 0xe8, 0x09,
	0x3c, 0x75, 0x34, 0x98, 0x69, 0x47, 0x99, 0x96, 0x7d, 0xd6, 0x77, 0xcb, 0x77, 0x71, 0x5f, 0x8d,
	0x3d, 0xb2, 0xf3, 0x1b, 0x15, 0xf0, 0x00, 0xb1, 0xbd, 0x98, 0x35, 0xf3, 0xf1, 0xb2, 0xf6, 0xad,
	0x72, 0xb8, 0x38, 0x66, 0xf2, 0x43, 0x3c, 0x7a, 0xe0, 0x03, 0xc6, 0xad, 0x22, 0xb7, 0xf1, 0x59,
	0xaa, 0x6b, 0xe3, 0xbc, 0xf6, 0x50, 0x88, 0xdf, 0x13, 0x50, 0x27, 0x3f, 0xc5, 0xf4, 0x2c, 0x3b,
	0x0b, 0x9e, 0x1d, 0xa7, 0xa1, 0x3d, 0xc9, 0x76, 0x5b, 0x90, 0xef, 0x42, 0xf4, 0xfa, 0xfa, 0x32,
	0xbf, 0xec, 0xed, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x49, 0x35, 0x5e, 0xe3, 0x51, 0x01, 0x00, 0x00,
}
