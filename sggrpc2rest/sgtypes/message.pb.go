// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package sgtypes

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

type ErrorCode int32

const (
	ErrorCode_UNKNOWN               ErrorCode = 0
	ErrorCode_SEQUENCE_NUMBER_WRONG ErrorCode = 1
	ErrorCode_TIMEOUT               ErrorCode = 2
)

var ErrorCode_name = map[int32]string{
	0: "UNKNOWN",
	1: "SEQUENCE_NUMBER_WRONG",
	2: "TIMEOUT",
}

var ErrorCode_value = map[string]int32{
	"UNKNOWN":               0,
	"SEQUENCE_NUMBER_WRONG": 1,
	"TIMEOUT":               2,
}

func (x ErrorCode) String() string {
	return proto.EnumName(ErrorCode_name, int32(x))
}

func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

type RawNegotiateMessage struct {
	SenderAddr           []byte     `protobuf:"bytes,1,opt,name=sender_addr,json=senderAddr,proto3" json:"sender_addr,omitempty"`
	ResourceType         *StructTag `protobuf:"bytes,2,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	SenderAmount         int64      `protobuf:"varint,3,opt,name=sender_amount,json=senderAmount,proto3" json:"sender_amount,omitempty"`
	ReceiverAddr         []byte     `protobuf:"bytes,4,opt,name=receiver_addr,json=receiverAddr,proto3" json:"receiver_addr,omitempty"`
	ReceiverAmount       int64      `protobuf:"varint,5,opt,name=receiver_amount,json=receiverAmount,proto3" json:"receiver_amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RawNegotiateMessage) Reset()         { *m = RawNegotiateMessage{} }
func (m *RawNegotiateMessage) String() string { return proto.CompactTextString(m) }
func (*RawNegotiateMessage) ProtoMessage()    {}
func (*RawNegotiateMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *RawNegotiateMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RawNegotiateMessage.Unmarshal(m, b)
}
func (m *RawNegotiateMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RawNegotiateMessage.Marshal(b, m, deterministic)
}
func (m *RawNegotiateMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RawNegotiateMessage.Merge(m, src)
}
func (m *RawNegotiateMessage) XXX_Size() int {
	return xxx_messageInfo_RawNegotiateMessage.Size(m)
}
func (m *RawNegotiateMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_RawNegotiateMessage.DiscardUnknown(m)
}

var xxx_messageInfo_RawNegotiateMessage proto.InternalMessageInfo

func (m *RawNegotiateMessage) GetSenderAddr() []byte {
	if m != nil {
		return m.SenderAddr
	}
	return nil
}

func (m *RawNegotiateMessage) GetResourceType() *StructTag {
	if m != nil {
		return m.ResourceType
	}
	return nil
}

func (m *RawNegotiateMessage) GetSenderAmount() int64 {
	if m != nil {
		return m.SenderAmount
	}
	return 0
}

func (m *RawNegotiateMessage) GetReceiverAddr() []byte {
	if m != nil {
		return m.ReceiverAddr
	}
	return nil
}

func (m *RawNegotiateMessage) GetReceiverAmount() int64 {
	if m != nil {
		return m.ReceiverAmount
	}
	return 0
}

type OpenChannelNodeNegotiateMessage struct {
	RawMessage           *RawNegotiateMessage `protobuf:"bytes,1,opt,name=raw_message,json=rawMessage,proto3" json:"raw_message,omitempty"`
	SenderSign           []byte               `protobuf:"bytes,2,opt,name=sender_sign,json=senderSign,proto3" json:"sender_sign,omitempty"`
	ReceiverSign         []byte               `protobuf:"bytes,3,opt,name=receiver_sign,json=receiverSign,proto3" json:"receiver_sign,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *OpenChannelNodeNegotiateMessage) Reset()         { *m = OpenChannelNodeNegotiateMessage{} }
func (m *OpenChannelNodeNegotiateMessage) String() string { return proto.CompactTextString(m) }
func (*OpenChannelNodeNegotiateMessage) ProtoMessage()    {}
func (*OpenChannelNodeNegotiateMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}

func (m *OpenChannelNodeNegotiateMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenChannelNodeNegotiateMessage.Unmarshal(m, b)
}
func (m *OpenChannelNodeNegotiateMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenChannelNodeNegotiateMessage.Marshal(b, m, deterministic)
}
func (m *OpenChannelNodeNegotiateMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenChannelNodeNegotiateMessage.Merge(m, src)
}
func (m *OpenChannelNodeNegotiateMessage) XXX_Size() int {
	return xxx_messageInfo_OpenChannelNodeNegotiateMessage.Size(m)
}
func (m *OpenChannelNodeNegotiateMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenChannelNodeNegotiateMessage.DiscardUnknown(m)
}

var xxx_messageInfo_OpenChannelNodeNegotiateMessage proto.InternalMessageInfo

func (m *OpenChannelNodeNegotiateMessage) GetRawMessage() *RawNegotiateMessage {
	if m != nil {
		return m.RawMessage
	}
	return nil
}

func (m *OpenChannelNodeNegotiateMessage) GetSenderSign() []byte {
	if m != nil {
		return m.SenderSign
	}
	return nil
}

func (m *OpenChannelNodeNegotiateMessage) GetReceiverSign() []byte {
	if m != nil {
		return m.ReceiverSign
	}
	return nil
}

type AddressMessage struct {
	Addr                 []byte   `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	IpAddr               []byte   `protobuf:"bytes,2,opt,name=ip_addr,json=ipAddr,proto3" json:"ip_addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddressMessage) Reset()         { *m = AddressMessage{} }
func (m *AddressMessage) String() string { return proto.CompactTextString(m) }
func (*AddressMessage) ProtoMessage()    {}
func (*AddressMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{2}
}

func (m *AddressMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressMessage.Unmarshal(m, b)
}
func (m *AddressMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressMessage.Marshal(b, m, deterministic)
}
func (m *AddressMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressMessage.Merge(m, src)
}
func (m *AddressMessage) XXX_Size() int {
	return xxx_messageInfo_AddressMessage.Size(m)
}
func (m *AddressMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressMessage.DiscardUnknown(m)
}

var xxx_messageInfo_AddressMessage proto.InternalMessageInfo

func (m *AddressMessage) GetAddr() []byte {
	if m != nil {
		return m.Addr
	}
	return nil
}

func (m *AddressMessage) GetIpAddr() []byte {
	if m != nil {
		return m.IpAddr
	}
	return nil
}

type StructTag struct {
	AccountAddr          []byte       `protobuf:"bytes,1,opt,name=account_addr,json=accountAddr,proto3" json:"account_addr,omitempty"`
	Module               string       `protobuf:"bytes,2,opt,name=module,proto3" json:"module,omitempty"`
	Name                 string       `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	TypeParams           []*StructTag `protobuf:"bytes,4,rep,name=type_params,json=typeParams,proto3" json:"type_params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *StructTag) Reset()         { *m = StructTag{} }
func (m *StructTag) String() string { return proto.CompactTextString(m) }
func (*StructTag) ProtoMessage()    {}
func (*StructTag) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{3}
}

func (m *StructTag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StructTag.Unmarshal(m, b)
}
func (m *StructTag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StructTag.Marshal(b, m, deterministic)
}
func (m *StructTag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StructTag.Merge(m, src)
}
func (m *StructTag) XXX_Size() int {
	return xxx_messageInfo_StructTag.Size(m)
}
func (m *StructTag) XXX_DiscardUnknown() {
	xxx_messageInfo_StructTag.DiscardUnknown(m)
}

var xxx_messageInfo_StructTag proto.InternalMessageInfo

func (m *StructTag) GetAccountAddr() []byte {
	if m != nil {
		return m.AccountAddr
	}
	return nil
}

func (m *StructTag) GetModule() string {
	if m != nil {
		return m.Module
	}
	return ""
}

func (m *StructTag) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StructTag) GetTypeParams() []*StructTag {
	if m != nil {
		return m.TypeParams
	}
	return nil
}

type ErrorMessage struct {
	RawTransactionHash   []byte   `protobuf:"bytes,1,opt,name=raw_transaction_hash,json=rawTransactionHash,proto3" json:"raw_transaction_hash,omitempty"`
	ErrorCode            uint32   `protobuf:"varint,2,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	ErrorMessage         string   `protobuf:"bytes,3,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorMessage) Reset()         { *m = ErrorMessage{} }
func (m *ErrorMessage) String() string { return proto.CompactTextString(m) }
func (*ErrorMessage) ProtoMessage()    {}
func (*ErrorMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{4}
}

func (m *ErrorMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorMessage.Unmarshal(m, b)
}
func (m *ErrorMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorMessage.Marshal(b, m, deterministic)
}
func (m *ErrorMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorMessage.Merge(m, src)
}
func (m *ErrorMessage) XXX_Size() int {
	return xxx_messageInfo_ErrorMessage.Size(m)
}
func (m *ErrorMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorMessage proto.InternalMessageInfo

func (m *ErrorMessage) GetRawTransactionHash() []byte {
	if m != nil {
		return m.RawTransactionHash
	}
	return nil
}

func (m *ErrorMessage) GetErrorCode() uint32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *ErrorMessage) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

type BalanceQueryRequest struct {
	LocalAddr            []byte   `protobuf:"bytes,1,opt,name=local_addr,json=localAddr,proto3" json:"local_addr,omitempty"`
	RemoteAddr           []byte   `protobuf:"bytes,2,opt,name=remote_addr,json=remoteAddr,proto3" json:"remote_addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BalanceQueryRequest) Reset()         { *m = BalanceQueryRequest{} }
func (m *BalanceQueryRequest) String() string { return proto.CompactTextString(m) }
func (*BalanceQueryRequest) ProtoMessage()    {}
func (*BalanceQueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{5}
}

func (m *BalanceQueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BalanceQueryRequest.Unmarshal(m, b)
}
func (m *BalanceQueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BalanceQueryRequest.Marshal(b, m, deterministic)
}
func (m *BalanceQueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalanceQueryRequest.Merge(m, src)
}
func (m *BalanceQueryRequest) XXX_Size() int {
	return xxx_messageInfo_BalanceQueryRequest.Size(m)
}
func (m *BalanceQueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BalanceQueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BalanceQueryRequest proto.InternalMessageInfo

func (m *BalanceQueryRequest) GetLocalAddr() []byte {
	if m != nil {
		return m.LocalAddr
	}
	return nil
}

func (m *BalanceQueryRequest) GetRemoteAddr() []byte {
	if m != nil {
		return m.RemoteAddr
	}
	return nil
}

type BalanceQueryResponse struct {
	LocalAddr            []byte   `protobuf:"bytes,1,opt,name=local_addr,json=localAddr,proto3" json:"local_addr,omitempty"`
	LocalBalance         uint64   `protobuf:"varint,2,opt,name=local_balance,json=localBalance,proto3" json:"local_balance,omitempty"`
	RemoteAddr           []byte   `protobuf:"bytes,3,opt,name=remote_addr,json=remoteAddr,proto3" json:"remote_addr,omitempty"`
	RemoteBalance        uint64   `protobuf:"varint,4,opt,name=remote_balance,json=remoteBalance,proto3" json:"remote_balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BalanceQueryResponse) Reset()         { *m = BalanceQueryResponse{} }
func (m *BalanceQueryResponse) String() string { return proto.CompactTextString(m) }
func (*BalanceQueryResponse) ProtoMessage()    {}
func (*BalanceQueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{6}
}

func (m *BalanceQueryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BalanceQueryResponse.Unmarshal(m, b)
}
func (m *BalanceQueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BalanceQueryResponse.Marshal(b, m, deterministic)
}
func (m *BalanceQueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalanceQueryResponse.Merge(m, src)
}
func (m *BalanceQueryResponse) XXX_Size() int {
	return xxx_messageInfo_BalanceQueryResponse.Size(m)
}
func (m *BalanceQueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BalanceQueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BalanceQueryResponse proto.InternalMessageInfo

func (m *BalanceQueryResponse) GetLocalAddr() []byte {
	if m != nil {
		return m.LocalAddr
	}
	return nil
}

func (m *BalanceQueryResponse) GetLocalBalance() uint64 {
	if m != nil {
		return m.LocalBalance
	}
	return 0
}

func (m *BalanceQueryResponse) GetRemoteAddr() []byte {
	if m != nil {
		return m.RemoteAddr
	}
	return nil
}

func (m *BalanceQueryResponse) GetRemoteBalance() uint64 {
	if m != nil {
		return m.RemoteBalance
	}
	return 0
}

type NextHop struct {
	RemoteAddr           []byte   `protobuf:"bytes,1,opt,name=remote_addr,json=remoteAddr,proto3" json:"remote_addr,omitempty"`
	Amount               uint64   `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NextHop) Reset()         { *m = NextHop{} }
func (m *NextHop) String() string { return proto.CompactTextString(m) }
func (*NextHop) ProtoMessage()    {}
func (*NextHop) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{7}
}

func (m *NextHop) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NextHop.Unmarshal(m, b)
}
func (m *NextHop) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NextHop.Marshal(b, m, deterministic)
}
func (m *NextHop) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NextHop.Merge(m, src)
}
func (m *NextHop) XXX_Size() int {
	return xxx_messageInfo_NextHop.Size(m)
}
func (m *NextHop) XXX_DiscardUnknown() {
	xxx_messageInfo_NextHop.DiscardUnknown(m)
}

var xxx_messageInfo_NextHop proto.InternalMessageInfo

func (m *NextHop) GetRemoteAddr() []byte {
	if m != nil {
		return m.RemoteAddr
	}
	return nil
}

func (m *NextHop) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type MultiHopChannelRequest struct {
	Request              *ChannelTransactionRequest `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Hops                 []*NextHop                 `protobuf:"bytes,2,rep,name=hops,proto3" json:"hops,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *MultiHopChannelRequest) Reset()         { *m = MultiHopChannelRequest{} }
func (m *MultiHopChannelRequest) String() string { return proto.CompactTextString(m) }
func (*MultiHopChannelRequest) ProtoMessage()    {}
func (*MultiHopChannelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{8}
}

func (m *MultiHopChannelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiHopChannelRequest.Unmarshal(m, b)
}
func (m *MultiHopChannelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiHopChannelRequest.Marshal(b, m, deterministic)
}
func (m *MultiHopChannelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiHopChannelRequest.Merge(m, src)
}
func (m *MultiHopChannelRequest) XXX_Size() int {
	return xxx_messageInfo_MultiHopChannelRequest.Size(m)
}
func (m *MultiHopChannelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiHopChannelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MultiHopChannelRequest proto.InternalMessageInfo

func (m *MultiHopChannelRequest) GetRequest() *ChannelTransactionRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *MultiHopChannelRequest) GetHops() []*NextHop {
	if m != nil {
		return m.Hops
	}
	return nil
}

type ExchangeSeedMessageRequest struct {
	SenderSeed           []byte   `protobuf:"bytes,1,opt,name=sender_seed,json=senderSeed,proto3" json:"sender_seed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExchangeSeedMessageRequest) Reset()         { *m = ExchangeSeedMessageRequest{} }
func (m *ExchangeSeedMessageRequest) String() string { return proto.CompactTextString(m) }
func (*ExchangeSeedMessageRequest) ProtoMessage()    {}
func (*ExchangeSeedMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{9}
}

func (m *ExchangeSeedMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExchangeSeedMessageRequest.Unmarshal(m, b)
}
func (m *ExchangeSeedMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExchangeSeedMessageRequest.Marshal(b, m, deterministic)
}
func (m *ExchangeSeedMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExchangeSeedMessageRequest.Merge(m, src)
}
func (m *ExchangeSeedMessageRequest) XXX_Size() int {
	return xxx_messageInfo_ExchangeSeedMessageRequest.Size(m)
}
func (m *ExchangeSeedMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExchangeSeedMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExchangeSeedMessageRequest proto.InternalMessageInfo

func (m *ExchangeSeedMessageRequest) GetSenderSeed() []byte {
	if m != nil {
		return m.SenderSeed
	}
	return nil
}

type ExchangeSeedMessageResponse struct {
	SenderSeed           []byte   `protobuf:"bytes,1,opt,name=sender_seed,json=senderSeed,proto3" json:"sender_seed,omitempty"`
	ReceiverSeed         []byte   `protobuf:"bytes,2,opt,name=receiver_seed,json=receiverSeed,proto3" json:"receiver_seed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExchangeSeedMessageResponse) Reset()         { *m = ExchangeSeedMessageResponse{} }
func (m *ExchangeSeedMessageResponse) String() string { return proto.CompactTextString(m) }
func (*ExchangeSeedMessageResponse) ProtoMessage()    {}
func (*ExchangeSeedMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{10}
}

func (m *ExchangeSeedMessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExchangeSeedMessageResponse.Unmarshal(m, b)
}
func (m *ExchangeSeedMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExchangeSeedMessageResponse.Marshal(b, m, deterministic)
}
func (m *ExchangeSeedMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExchangeSeedMessageResponse.Merge(m, src)
}
func (m *ExchangeSeedMessageResponse) XXX_Size() int {
	return xxx_messageInfo_ExchangeSeedMessageResponse.Size(m)
}
func (m *ExchangeSeedMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExchangeSeedMessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExchangeSeedMessageResponse proto.InternalMessageInfo

func (m *ExchangeSeedMessageResponse) GetSenderSeed() []byte {
	if m != nil {
		return m.SenderSeed
	}
	return nil
}

func (m *ExchangeSeedMessageResponse) GetReceiverSeed() []byte {
	if m != nil {
		return m.ReceiverSeed
	}
	return nil
}

type AntQueryMessage struct {
	SValue                   []byte                  `protobuf:"bytes,1,opt,name=s_value,json=sValue,proto3" json:"s_value,omitempty"`
	SenderAddr               []byte                  `protobuf:"bytes,2,opt,name=sender_addr,json=senderAddr,proto3" json:"sender_addr,omitempty"`
	BalanceQueryResponseList []*BalanceQueryResponse `protobuf:"bytes,3,rep,name=balance_query_response_list,json=balanceQueryResponseList,proto3" json:"balance_query_response_list,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                `json:"-"`
	XXX_unrecognized         []byte                  `json:"-"`
	XXX_sizecache            int32                   `json:"-"`
}

func (m *AntQueryMessage) Reset()         { *m = AntQueryMessage{} }
func (m *AntQueryMessage) String() string { return proto.CompactTextString(m) }
func (*AntQueryMessage) ProtoMessage()    {}
func (*AntQueryMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{11}
}

func (m *AntQueryMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AntQueryMessage.Unmarshal(m, b)
}
func (m *AntQueryMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AntQueryMessage.Marshal(b, m, deterministic)
}
func (m *AntQueryMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AntQueryMessage.Merge(m, src)
}
func (m *AntQueryMessage) XXX_Size() int {
	return xxx_messageInfo_AntQueryMessage.Size(m)
}
func (m *AntQueryMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_AntQueryMessage.DiscardUnknown(m)
}

var xxx_messageInfo_AntQueryMessage proto.InternalMessageInfo

func (m *AntQueryMessage) GetSValue() []byte {
	if m != nil {
		return m.SValue
	}
	return nil
}

func (m *AntQueryMessage) GetSenderAddr() []byte {
	if m != nil {
		return m.SenderAddr
	}
	return nil
}

func (m *AntQueryMessage) GetBalanceQueryResponseList() []*BalanceQueryResponse {
	if m != nil {
		return m.BalanceQueryResponseList
	}
	return nil
}

type AntFinalMessage struct {
	RValue                   []byte                  `protobuf:"bytes,1,opt,name=r_value,json=rValue,proto3" json:"r_value,omitempty"`
	BalanceQueryResponseList []*BalanceQueryResponse `protobuf:"bytes,2,rep,name=balance_query_response_list,json=balanceQueryResponseList,proto3" json:"balance_query_response_list,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                `json:"-"`
	XXX_unrecognized         []byte                  `json:"-"`
	XXX_sizecache            int32                   `json:"-"`
}

func (m *AntFinalMessage) Reset()         { *m = AntFinalMessage{} }
func (m *AntFinalMessage) String() string { return proto.CompactTextString(m) }
func (*AntFinalMessage) ProtoMessage()    {}
func (*AntFinalMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{12}
}

func (m *AntFinalMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AntFinalMessage.Unmarshal(m, b)
}
func (m *AntFinalMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AntFinalMessage.Marshal(b, m, deterministic)
}
func (m *AntFinalMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AntFinalMessage.Merge(m, src)
}
func (m *AntFinalMessage) XXX_Size() int {
	return xxx_messageInfo_AntFinalMessage.Size(m)
}
func (m *AntFinalMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_AntFinalMessage.DiscardUnknown(m)
}

var xxx_messageInfo_AntFinalMessage proto.InternalMessageInfo

func (m *AntFinalMessage) GetRValue() []byte {
	if m != nil {
		return m.RValue
	}
	return nil
}

func (m *AntFinalMessage) GetBalanceQueryResponseList() []*BalanceQueryResponse {
	if m != nil {
		return m.BalanceQueryResponseList
	}
	return nil
}

func init() {
	proto.RegisterEnum("sgtypes.ErrorCode", ErrorCode_name, ErrorCode_value)
	proto.RegisterType((*RawNegotiateMessage)(nil), "sgtypes.RawNegotiateMessage")
	proto.RegisterType((*OpenChannelNodeNegotiateMessage)(nil), "sgtypes.OpenChannelNodeNegotiateMessage")
	proto.RegisterType((*AddressMessage)(nil), "sgtypes.AddressMessage")
	proto.RegisterType((*StructTag)(nil), "sgtypes.StructTag")
	proto.RegisterType((*ErrorMessage)(nil), "sgtypes.ErrorMessage")
	proto.RegisterType((*BalanceQueryRequest)(nil), "sgtypes.BalanceQueryRequest")
	proto.RegisterType((*BalanceQueryResponse)(nil), "sgtypes.BalanceQueryResponse")
	proto.RegisterType((*NextHop)(nil), "sgtypes.NextHop")
	proto.RegisterType((*MultiHopChannelRequest)(nil), "sgtypes.MultiHopChannelRequest")
	proto.RegisterType((*ExchangeSeedMessageRequest)(nil), "sgtypes.ExchangeSeedMessageRequest")
	proto.RegisterType((*ExchangeSeedMessageResponse)(nil), "sgtypes.ExchangeSeedMessageResponse")
	proto.RegisterType((*AntQueryMessage)(nil), "sgtypes.AntQueryMessage")
	proto.RegisterType((*AntFinalMessage)(nil), "sgtypes.AntFinalMessage")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 794 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x4f, 0x6f, 0xe3, 0x54,
	0x10, 0xc7, 0x49, 0x48, 0x94, 0xb1, 0xd3, 0x8d, 0xbc, 0x4b, 0x37, 0xbb, 0xcb, 0xaa, 0xc1, 0x05,
	0x51, 0x71, 0xa8, 0x50, 0x7b, 0xe0, 0x42, 0x25, 0xda, 0xca, 0x50, 0x04, 0x71, 0xa8, 0x93, 0xd0,
	0x0b, 0x92, 0xf5, 0x6a, 0x8f, 0x12, 0x4b, 0x8e, 0x9f, 0xfb, 0x9e, 0xdd, 0xb4, 0x12, 0x47, 0x0e,
	0xdc, 0xf8, 0x0e, 0x1c, 0x90, 0xf8, 0x5a, 0x7c, 0x12, 0xf4, 0xfe, 0xd8, 0x75, 0xfe, 0xa0, 0x1e,
	0xb8, 0xf9, 0xcd, 0x9b, 0xf9, 0xcd, 0x6f, 0x7e, 0x33, 0xf3, 0x0c, 0xbd, 0x25, 0x72, 0x4e, 0xe6,
	0x78, 0x9c, 0x31, 0x9a, 0x53, 0xbb, 0xc3, 0xe7, 0xf9, 0x63, 0x86, 0xfc, 0xed, 0x9b, 0x70, 0x41,
	0xd2, 0x14, 0x93, 0x20, 0x67, 0x24, 0xe5, 0x24, 0xcc, 0x63, 0x9a, 0x2a, 0x1f, 0xe7, 0x1f, 0x03,
	0x5e, 0xfa, 0x64, 0xe5, 0xe1, 0x9c, 0xe6, 0x31, 0xc9, 0x71, 0xa4, 0x10, 0xec, 0x03, 0x30, 0x39,
	0xa6, 0x11, 0xb2, 0x80, 0x44, 0x11, 0x1b, 0x18, 0x43, 0xe3, 0xc8, 0xf2, 0x41, 0x99, 0xce, 0xa3,
	0x88, 0xd9, 0x5f, 0x41, 0x8f, 0x21, 0xa7, 0x05, 0x0b, 0x31, 0x10, 0x59, 0x06, 0x8d, 0xa1, 0x71,
	0x64, 0x9e, 0xd8, 0xc7, 0x3a, 0xe9, 0xf1, 0x24, 0x67, 0x45, 0x98, 0x4f, 0xc9, 0xdc, 0xb7, 0x4a,
	0xc7, 0xe9, 0x63, 0x86, 0xf6, 0x21, 0xf4, 0x4a, 0xe4, 0x25, 0x2d, 0xd2, 0x7c, 0xd0, 0x1c, 0x1a,
	0x47, 0x4d, 0xdf, 0xd2, 0xd8, 0xd2, 0x26, 0x9c, 0x18, 0x86, 0x18, 0xdf, 0x97, 0x04, 0x5a, 0x92,
	0x80, 0x55, 0x1a, 0x25, 0x85, 0xcf, 0xe1, 0xc5, 0x93, 0x93, 0xc2, 0xfa, 0x50, 0x62, 0xed, 0x55,
	0x6e, 0xd2, 0xea, 0xfc, 0x65, 0xc0, 0xc1, 0x38, 0xc3, 0xf4, 0x52, 0xc9, 0xe0, 0xd1, 0x08, 0xb7,
	0x0a, 0x3e, 0x03, 0x93, 0x91, 0x55, 0xa0, 0x15, 0x94, 0x05, 0x9b, 0x27, 0x1f, 0x57, 0xd5, 0xec,
	0xd0, 0xc8, 0x07, 0x46, 0x56, 0xdb, 0x7a, 0xf1, 0x78, 0x9e, 0x4a, 0x31, 0x2a, 0xbd, 0x26, 0xf1,
	0x3c, 0x5d, 0xab, 0x48, 0xba, 0x34, 0xd7, 0x2b, 0x12, 0x4e, 0xce, 0x19, 0xec, 0x89, 0xca, 0x90,
	0xf3, 0x12, 0xd7, 0x86, 0x56, 0xad, 0x01, 0xf2, 0xdb, 0x7e, 0x0d, 0x9d, 0x38, 0x53, 0xb2, 0xa8,
	0x3c, 0xed, 0x38, 0x13, 0x61, 0xce, 0x1f, 0x06, 0x74, 0x2b, 0xd9, 0xed, 0x4f, 0xc0, 0x22, 0x61,
	0x28, 0x04, 0xa8, 0xf7, 0xd0, 0xd4, 0x36, 0xa9, 0xe0, 0x3e, 0xb4, 0x97, 0x34, 0x2a, 0x12, 0xd5,
	0xbd, 0xae, 0xaf, 0x4f, 0x22, 0x6b, 0x4a, 0x96, 0x28, 0x39, 0x76, 0x7d, 0xf9, 0x6d, 0x9f, 0x82,
	0x29, 0xa4, 0x08, 0x32, 0xc2, 0xc8, 0x92, 0x0f, 0x5a, 0xc3, 0xe6, 0x7f, 0xb4, 0x1b, 0x84, 0xe1,
	0x27, 0xe9, 0xe5, 0xfc, 0x66, 0x80, 0xe5, 0x32, 0x46, 0x59, 0x59, 0xcf, 0x97, 0xf0, 0x4a, 0xc8,
	0x5c, 0x1b, 0xc4, 0x60, 0x41, 0xf8, 0x42, 0x93, 0xb3, 0x19, 0x59, 0x4d, 0x9f, 0xae, 0xae, 0x08,
	0x5f, 0xd8, 0xef, 0x01, 0x50, 0x20, 0x04, 0x21, 0x8d, 0x14, 0xcf, 0x9e, 0xdf, 0x95, 0x96, 0x4b,
	0x1a, 0xc9, 0x71, 0x52, 0xd7, 0x65, 0xe7, 0x14, 0x67, 0x0b, 0x6b, 0x59, 0x9d, 0x19, 0xbc, 0xbc,
	0x20, 0x09, 0x49, 0x43, 0xbc, 0x2e, 0x90, 0x3d, 0xfa, 0x78, 0x57, 0x20, 0xcf, 0x05, 0x74, 0x42,
	0x43, 0x92, 0xd4, 0xf5, 0xe9, 0x4a, 0x8b, 0x54, 0xe7, 0x00, 0x4c, 0x86, 0x4b, 0x9a, 0x63, 0x5d,
	0x6b, 0x50, 0x26, 0xa9, 0xf7, 0x9f, 0x06, 0xbc, 0x5a, 0xc7, 0xe5, 0x19, 0x4d, 0x39, 0x3e, 0x07,
	0x7c, 0x08, 0x3d, 0x75, 0x7d, 0xab, 0x82, 0x25, 0x74, 0xcb, 0xb7, 0xa4, 0x51, 0x03, 0x6e, 0x66,
	0x6f, 0x6e, 0x66, 0xb7, 0x3f, 0x83, 0x3d, 0xed, 0x50, 0xc2, 0xb4, 0x24, 0x4c, 0x4f, 0x59, 0x35,
	0x8e, 0x73, 0x01, 0x1d, 0x0f, 0x1f, 0xf2, 0x2b, 0x9a, 0x6d, 0x42, 0x1a, 0x5b, 0x90, 0xfb, 0xd0,
	0xd6, 0x8b, 0xa4, 0x18, 0xe9, 0x93, 0xf3, 0x2b, 0xec, 0x8f, 0x8a, 0x24, 0x8f, 0xaf, 0x68, 0xa6,
	0x77, 0xa8, 0x94, 0xf0, 0x6b, 0xe8, 0x30, 0xf5, 0xa9, 0x57, 0xc6, 0xa9, 0x26, 0x42, 0x7b, 0xd6,
	0xfa, 0xa9, 0x83, 0xfc, 0x32, 0xc4, 0xfe, 0x14, 0x5a, 0x0b, 0x9a, 0xf1, 0x41, 0x43, 0x0e, 0x53,
	0xbf, 0x0a, 0xd5, 0x84, 0x7d, 0x79, 0xeb, 0x9c, 0xc1, 0x5b, 0xf7, 0x41, 0x3c, 0x61, 0x73, 0x9c,
	0x20, 0x46, 0xe5, 0xfa, 0x69, 0x8c, 0xda, 0xe6, 0x21, 0x46, 0xeb, 0x2f, 0x95, 0x70, 0x77, 0x42,
	0x78, 0xb7, 0x33, 0x5c, 0xf7, 0xea, 0xb9, 0xf8, 0xf5, 0xcd, 0x15, 0x2e, 0x8d, 0x8d, 0xcd, 0x15,
	0x49, 0xfe, 0x36, 0xe0, 0xc5, 0x79, 0x9a, 0xcb, 0x31, 0x28, 0x67, 0xfd, 0x35, 0x74, 0x78, 0x70,
	0x4f, 0x92, 0x02, 0x35, 0x6a, 0x9b, 0xff, 0x2c, 0x4e, 0x9b, 0x8f, 0x6b, 0x63, 0xeb, 0x71, 0xfd,
	0x05, 0xde, 0xe9, 0x9e, 0x06, 0x77, 0x02, 0x31, 0x60, 0x9a, 0x6d, 0x90, 0xc4, 0x5c, 0xbc, 0x98,
	0x42, 0xae, 0xf7, 0x95, 0x5c, 0xbb, 0x66, 0xd0, 0x1f, 0xdc, 0xee, 0xb0, 0xfe, 0x18, 0xf3, 0xdc,
	0xf9, 0x5d, 0x71, 0xfd, 0x36, 0x4e, 0x49, 0x52, 0xe3, 0xca, 0xd6, 0xb9, 0x32, 0xc5, 0xf5, 0x19,
	0x2a, 0x8d, 0xff, 0x45, 0xe5, 0x8b, 0x6f, 0xa0, 0xeb, 0x56, 0xab, 0x6c, 0x42, 0x67, 0xe6, 0xfd,
	0xe0, 0x8d, 0x6f, 0xbc, 0xfe, 0x07, 0xf6, 0x1b, 0xf8, 0x68, 0xe2, 0x5e, 0xcf, 0x5c, 0xef, 0xd2,
	0x0d, 0xbc, 0xd9, 0xe8, 0xc2, 0xf5, 0x83, 0x1b, 0x7f, 0xec, 0x7d, 0xd7, 0x37, 0x84, 0xdf, 0xf4,
	0xfb, 0x91, 0x3b, 0x9e, 0x4d, 0xfb, 0x8d, 0xdb, 0xb6, 0xfc, 0x8f, 0x9d, 0xfe, 0x1b, 0x00, 0x00,
	0xff, 0xff, 0x1e, 0x3d, 0x88, 0xf1, 0xfc, 0x06, 0x00, 0x00,
}
