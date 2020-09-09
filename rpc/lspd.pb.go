// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lspd.proto

package lspd

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ChannelInformationRequest struct {
	/// The identity pubkey of the Lightning node
	Pubkey               string   `protobuf:"bytes,1,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChannelInformationRequest) Reset()         { *m = ChannelInformationRequest{} }
func (m *ChannelInformationRequest) String() string { return proto.CompactTextString(m) }
func (*ChannelInformationRequest) ProtoMessage()    {}
func (*ChannelInformationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c69a0f5a734bca26, []int{0}
}

func (m *ChannelInformationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChannelInformationRequest.Unmarshal(m, b)
}
func (m *ChannelInformationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChannelInformationRequest.Marshal(b, m, deterministic)
}
func (m *ChannelInformationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelInformationRequest.Merge(m, src)
}
func (m *ChannelInformationRequest) XXX_Size() int {
	return xxx_messageInfo_ChannelInformationRequest.Size(m)
}
func (m *ChannelInformationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelInformationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelInformationRequest proto.InternalMessageInfo

func (m *ChannelInformationRequest) GetPubkey() string {
	if m != nil {
		return m.Pubkey
	}
	return ""
}

type ChannelInformationReply struct {
	/// The name of of lsp
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	/// The identity pubkey of the Lightning node
	Pubkey string `protobuf:"bytes,2,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	/// The network location of the lightning node, e.g. `12.34.56.78:9012` or
	/// `localhost:10011`
	Host string `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	/// The channel capacity in satoshis
	ChannelCapacity int64 `protobuf:"varint,4,opt,name=channel_capacity,proto3" json:"channel_capacity,omitempty"`
	/// The target number of blocks that the funding transaction should be
	/// confirmed by.
	TargetConf int32 `protobuf:"varint,5,opt,name=target_conf,proto3" json:"target_conf,omitempty"`
	/// The base fee charged regardless of the number of milli-satoshis sent.
	BaseFeeMsat int64 `protobuf:"varint,6,opt,name=base_fee_msat,proto3" json:"base_fee_msat,omitempty"`
	/// The effective fee rate in milli-satoshis. The precision of this value goes
	/// up to 6 decimal places, so 1e-6.
	FeeRate float64 `protobuf:"fixed64,7,opt,name=fee_rate,proto3" json:"fee_rate,omitempty"`
	/// The required timelock delta for HTLCs forwarded over the channel.
	TimeLockDelta uint32 `protobuf:"varint,8,opt,name=time_lock_delta,proto3" json:"time_lock_delta,omitempty"`
	/// The minimum value in millisatoshi we will require for incoming HTLCs on
	/// the channel.
	MinHtlcMsat          int64    `protobuf:"varint,9,opt,name=min_htlc_msat,proto3" json:"min_htlc_msat,omitempty"`
	ChannelFeePermyriad  int64    `protobuf:"varint,10,opt,name=channel_fee_permyriad,json=channelFeePermyriad,proto3" json:"channel_fee_permyriad,omitempty"`
	LspPubkey            []byte   `protobuf:"bytes,11,opt,name=lsp_pubkey,json=lspPubkey,proto3" json:"lsp_pubkey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChannelInformationReply) Reset()         { *m = ChannelInformationReply{} }
func (m *ChannelInformationReply) String() string { return proto.CompactTextString(m) }
func (*ChannelInformationReply) ProtoMessage()    {}
func (*ChannelInformationReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_c69a0f5a734bca26, []int{1}
}

func (m *ChannelInformationReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChannelInformationReply.Unmarshal(m, b)
}
func (m *ChannelInformationReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChannelInformationReply.Marshal(b, m, deterministic)
}
func (m *ChannelInformationReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelInformationReply.Merge(m, src)
}
func (m *ChannelInformationReply) XXX_Size() int {
	return xxx_messageInfo_ChannelInformationReply.Size(m)
}
func (m *ChannelInformationReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelInformationReply.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelInformationReply proto.InternalMessageInfo

func (m *ChannelInformationReply) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ChannelInformationReply) GetPubkey() string {
	if m != nil {
		return m.Pubkey
	}
	return ""
}

func (m *ChannelInformationReply) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *ChannelInformationReply) GetChannelCapacity() int64 {
	if m != nil {
		return m.ChannelCapacity
	}
	return 0
}

func (m *ChannelInformationReply) GetTargetConf() int32 {
	if m != nil {
		return m.TargetConf
	}
	return 0
}

func (m *ChannelInformationReply) GetBaseFeeMsat() int64 {
	if m != nil {
		return m.BaseFeeMsat
	}
	return 0
}

func (m *ChannelInformationReply) GetFeeRate() float64 {
	if m != nil {
		return m.FeeRate
	}
	return 0
}

func (m *ChannelInformationReply) GetTimeLockDelta() uint32 {
	if m != nil {
		return m.TimeLockDelta
	}
	return 0
}

func (m *ChannelInformationReply) GetMinHtlcMsat() int64 {
	if m != nil {
		return m.MinHtlcMsat
	}
	return 0
}

func (m *ChannelInformationReply) GetChannelFeePermyriad() int64 {
	if m != nil {
		return m.ChannelFeePermyriad
	}
	return 0
}

func (m *ChannelInformationReply) GetLspPubkey() []byte {
	if m != nil {
		return m.LspPubkey
	}
	return nil
}

type OpenChannelRequest struct {
	/// The identity pubkey of the Lightning node
	Pubkey               string   `protobuf:"bytes,1,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpenChannelRequest) Reset()         { *m = OpenChannelRequest{} }
func (m *OpenChannelRequest) String() string { return proto.CompactTextString(m) }
func (*OpenChannelRequest) ProtoMessage()    {}
func (*OpenChannelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c69a0f5a734bca26, []int{2}
}

func (m *OpenChannelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenChannelRequest.Unmarshal(m, b)
}
func (m *OpenChannelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenChannelRequest.Marshal(b, m, deterministic)
}
func (m *OpenChannelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenChannelRequest.Merge(m, src)
}
func (m *OpenChannelRequest) XXX_Size() int {
	return xxx_messageInfo_OpenChannelRequest.Size(m)
}
func (m *OpenChannelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenChannelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OpenChannelRequest proto.InternalMessageInfo

func (m *OpenChannelRequest) GetPubkey() string {
	if m != nil {
		return m.Pubkey
	}
	return ""
}

type OpenChannelReply struct {
	/// The transaction hash
	TxHash string `protobuf:"bytes,1,opt,name=tx_hash,proto3" json:"tx_hash,omitempty"`
	/// The output index
	OutputIndex          uint32   `protobuf:"varint,2,opt,name=output_index,proto3" json:"output_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpenChannelReply) Reset()         { *m = OpenChannelReply{} }
func (m *OpenChannelReply) String() string { return proto.CompactTextString(m) }
func (*OpenChannelReply) ProtoMessage()    {}
func (*OpenChannelReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_c69a0f5a734bca26, []int{3}
}

func (m *OpenChannelReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenChannelReply.Unmarshal(m, b)
}
func (m *OpenChannelReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenChannelReply.Marshal(b, m, deterministic)
}
func (m *OpenChannelReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenChannelReply.Merge(m, src)
}
func (m *OpenChannelReply) XXX_Size() int {
	return xxx_messageInfo_OpenChannelReply.Size(m)
}
func (m *OpenChannelReply) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenChannelReply.DiscardUnknown(m)
}

var xxx_messageInfo_OpenChannelReply proto.InternalMessageInfo

func (m *OpenChannelReply) GetTxHash() string {
	if m != nil {
		return m.TxHash
	}
	return ""
}

func (m *OpenChannelReply) GetOutputIndex() uint32 {
	if m != nil {
		return m.OutputIndex
	}
	return 0
}

type RegisterPaymentRequest struct {
	Blob                 []byte   `protobuf:"bytes,3,opt,name=blob,proto3" json:"blob,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterPaymentRequest) Reset()         { *m = RegisterPaymentRequest{} }
func (m *RegisterPaymentRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterPaymentRequest) ProtoMessage()    {}
func (*RegisterPaymentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c69a0f5a734bca26, []int{4}
}

func (m *RegisterPaymentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterPaymentRequest.Unmarshal(m, b)
}
func (m *RegisterPaymentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterPaymentRequest.Marshal(b, m, deterministic)
}
func (m *RegisterPaymentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterPaymentRequest.Merge(m, src)
}
func (m *RegisterPaymentRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterPaymentRequest.Size(m)
}
func (m *RegisterPaymentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterPaymentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterPaymentRequest proto.InternalMessageInfo

func (m *RegisterPaymentRequest) GetBlob() []byte {
	if m != nil {
		return m.Blob
	}
	return nil
}

type RegisterPaymentReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterPaymentReply) Reset()         { *m = RegisterPaymentReply{} }
func (m *RegisterPaymentReply) String() string { return proto.CompactTextString(m) }
func (*RegisterPaymentReply) ProtoMessage()    {}
func (*RegisterPaymentReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_c69a0f5a734bca26, []int{5}
}

func (m *RegisterPaymentReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterPaymentReply.Unmarshal(m, b)
}
func (m *RegisterPaymentReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterPaymentReply.Marshal(b, m, deterministic)
}
func (m *RegisterPaymentReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterPaymentReply.Merge(m, src)
}
func (m *RegisterPaymentReply) XXX_Size() int {
	return xxx_messageInfo_RegisterPaymentReply.Size(m)
}
func (m *RegisterPaymentReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterPaymentReply.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterPaymentReply proto.InternalMessageInfo

type PaymentInformation struct {
	PaymentHash          []byte   `protobuf:"bytes,1,opt,name=payment_hash,json=paymentHash,proto3" json:"payment_hash,omitempty"`
	PaymentSecret        []byte   `protobuf:"bytes,2,opt,name=payment_secret,json=paymentSecret,proto3" json:"payment_secret,omitempty"`
	Destination          []byte   `protobuf:"bytes,3,opt,name=destination,proto3" json:"destination,omitempty"`
	IncomingAmountMsat   int64    `protobuf:"varint,4,opt,name=incoming_amount_msat,json=incomingAmountMsat,proto3" json:"incoming_amount_msat,omitempty"`
	OutgoingAmountMsat   int64    `protobuf:"varint,5,opt,name=outgoing_amount_msat,json=outgoingAmountMsat,proto3" json:"outgoing_amount_msat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaymentInformation) Reset()         { *m = PaymentInformation{} }
func (m *PaymentInformation) String() string { return proto.CompactTextString(m) }
func (*PaymentInformation) ProtoMessage()    {}
func (*PaymentInformation) Descriptor() ([]byte, []int) {
	return fileDescriptor_c69a0f5a734bca26, []int{6}
}

func (m *PaymentInformation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentInformation.Unmarshal(m, b)
}
func (m *PaymentInformation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentInformation.Marshal(b, m, deterministic)
}
func (m *PaymentInformation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentInformation.Merge(m, src)
}
func (m *PaymentInformation) XXX_Size() int {
	return xxx_messageInfo_PaymentInformation.Size(m)
}
func (m *PaymentInformation) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentInformation.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentInformation proto.InternalMessageInfo

func (m *PaymentInformation) GetPaymentHash() []byte {
	if m != nil {
		return m.PaymentHash
	}
	return nil
}

func (m *PaymentInformation) GetPaymentSecret() []byte {
	if m != nil {
		return m.PaymentSecret
	}
	return nil
}

func (m *PaymentInformation) GetDestination() []byte {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (m *PaymentInformation) GetIncomingAmountMsat() int64 {
	if m != nil {
		return m.IncomingAmountMsat
	}
	return 0
}

func (m *PaymentInformation) GetOutgoingAmountMsat() int64 {
	if m != nil {
		return m.OutgoingAmountMsat
	}
	return 0
}

func init() {
	proto.RegisterType((*ChannelInformationRequest)(nil), "lspd.ChannelInformationRequest")
	proto.RegisterType((*ChannelInformationReply)(nil), "lspd.ChannelInformationReply")
	proto.RegisterType((*OpenChannelRequest)(nil), "lspd.OpenChannelRequest")
	proto.RegisterType((*OpenChannelReply)(nil), "lspd.OpenChannelReply")
	proto.RegisterType((*RegisterPaymentRequest)(nil), "lspd.RegisterPaymentRequest")
	proto.RegisterType((*RegisterPaymentReply)(nil), "lspd.RegisterPaymentReply")
	proto.RegisterType((*PaymentInformation)(nil), "lspd.PaymentInformation")
}

func init() {
	proto.RegisterFile("lspd.proto", fileDescriptor_c69a0f5a734bca26)
}

var fileDescriptor_c69a0f5a734bca26 = []byte{
	// 568 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xff, 0x6e, 0xd3, 0x30,
	0x10, 0x26, 0xac, 0xfb, 0xd1, 0x6b, 0xcb, 0xa6, 0xa3, 0x94, 0x50, 0x31, 0x11, 0x0a, 0x48, 0x11,
	0x9a, 0xa6, 0x69, 0x7b, 0x82, 0x0d, 0x09, 0x81, 0xc4, 0x44, 0x65, 0x24, 0xfe, 0x8d, 0xdc, 0xf4,
	0xd6, 0x5a, 0x4b, 0x6c, 0x13, 0xbb, 0x68, 0x7d, 0x07, 0x1e, 0x93, 0x97, 0xe0, 0x3f, 0x64, 0x3b,
	0x61, 0xed, 0xba, 0x89, 0xff, 0xce, 0xdf, 0x7d, 0xf7, 0xdd, 0xc5, 0xdf, 0x39, 0x00, 0x85, 0xd1,
	0xd3, 0x63, 0x5d, 0x29, 0xab, 0xb0, 0xe5, 0xe2, 0xd1, 0x19, 0xbc, 0xf8, 0x30, 0xe7, 0x52, 0x52,
	0xf1, 0x59, 0x5e, 0xa9, 0xaa, 0xe4, 0x56, 0x28, 0xc9, 0xe8, 0xc7, 0x82, 0x8c, 0xc5, 0x01, 0xec,
	0xe8, 0xc5, 0xe4, 0x9a, 0x96, 0x71, 0x94, 0x44, 0x69, 0x9b, 0xd5, 0xa7, 0xd1, 0xaf, 0x2d, 0x78,
	0x7e, 0x5f, 0x95, 0x2e, 0x96, 0x88, 0xd0, 0x92, 0xbc, 0xa4, 0xba, 0xc2, 0xc7, 0x2b, 0x3a, 0x8f,
	0x57, 0x75, 0x1c, 0x77, 0xae, 0x8c, 0x8d, 0xb7, 0x02, 0xd7, 0xc5, 0xf8, 0x1e, 0x0e, 0xf2, 0x20,
	0x9d, 0xe5, 0x5c, 0xf3, 0x5c, 0xd8, 0x65, 0xdc, 0x4a, 0xa2, 0x74, 0x8b, 0x6d, 0xe0, 0x98, 0x40,
	0xc7, 0xf2, 0x6a, 0x46, 0x36, 0xcb, 0x95, 0xbc, 0x8a, 0xb7, 0x93, 0x28, 0xdd, 0x66, 0xab, 0x10,
	0xbe, 0x85, 0xde, 0x84, 0x1b, 0xca, 0xae, 0x88, 0xb2, 0xd2, 0x70, 0x1b, 0xef, 0x78, 0xa9, 0x75,
	0x10, 0x87, 0xb0, 0xe7, 0xe2, 0x8a, 0x5b, 0x8a, 0x77, 0x93, 0x28, 0x8d, 0xd8, 0xbf, 0x33, 0xa6,
	0xb0, 0x6f, 0x45, 0x49, 0x59, 0xa1, 0xf2, 0xeb, 0x6c, 0x4a, 0x85, 0xe5, 0xf1, 0x5e, 0x12, 0xa5,
	0x3d, 0x76, 0x17, 0x76, 0xbd, 0x4a, 0x21, 0xb3, 0xb9, 0x2d, 0xf2, 0xd0, 0xab, 0x1d, 0x7a, 0xad,
	0x81, 0x78, 0x0a, 0xcf, 0x9a, 0xef, 0x70, 0x3d, 0x34, 0x55, 0xe5, 0xb2, 0x12, 0x7c, 0x1a, 0x83,
	0x67, 0x3f, 0xad, 0x93, 0x1f, 0x89, 0xc6, 0x4d, 0x0a, 0x0f, 0xbd, 0x71, 0x59, 0x7d, 0x87, 0x9d,
	0x24, 0x4a, 0xbb, 0xac, 0x5d, 0x18, 0x3d, 0x0e, 0x76, 0x1c, 0x01, 0x7e, 0xd5, 0x24, 0x6b, 0x47,
	0xfe, 0x67, 0xde, 0x18, 0x0e, 0xd6, 0xd8, 0xce, 0xb4, 0x18, 0x76, 0xed, 0x4d, 0x36, 0xe7, 0x66,
	0x5e, 0x93, 0x9b, 0x23, 0x8e, 0xa0, 0xab, 0x16, 0x56, 0x2f, 0x6c, 0x26, 0xe4, 0x94, 0x6e, 0xbc,
	0x81, 0x3d, 0xb6, 0x86, 0x8d, 0x8e, 0x60, 0xc0, 0x68, 0x26, 0x8c, 0xa5, 0x6a, 0xcc, 0x97, 0x25,
	0x49, 0xdb, 0xcc, 0x80, 0xd0, 0x9a, 0x14, 0x6a, 0xe2, 0x0d, 0xee, 0x32, 0x1f, 0x8f, 0x06, 0xd0,
	0xdf, 0x60, 0xeb, 0x62, 0x39, 0xfa, 0x1d, 0x01, 0xd6, 0xc0, 0xca, 0x52, 0xe1, 0x6b, 0xe8, 0xea,
	0x80, 0xde, 0xce, 0xd7, 0x65, 0x9d, 0x1a, 0xfb, 0xe4, 0x66, 0x7c, 0x07, 0x4f, 0x1a, 0x8a, 0xa1,
	0xbc, 0x22, 0xeb, 0xa7, 0xec, 0xb2, 0x5e, 0x8d, 0x7e, 0xf3, 0xa0, 0xdb, 0x96, 0x29, 0x19, 0x2b,
	0xa4, 0x17, 0xae, 0x67, 0x5a, 0x85, 0xf0, 0x04, 0xfa, 0x42, 0xe6, 0xaa, 0x14, 0x72, 0x96, 0xf1,
	0x52, 0x2d, 0xa4, 0x0d, 0x46, 0x86, 0xfd, 0xc3, 0x26, 0x77, 0xee, 0x53, 0x97, 0xce, 0xcd, 0x13,
	0xe8, 0xab, 0x85, 0x9d, 0xa9, 0xbb, 0x15, 0xdb, 0xa1, 0xa2, 0xc9, 0xdd, 0x56, 0x9c, 0xfe, 0x89,
	0xa0, 0x57, 0xdf, 0xbd, 0xb3, 0x81, 0x2a, 0xfc, 0x0e, 0xb8, 0xf9, 0x98, 0xf0, 0xd5, 0xb1, 0x7f,
	0xab, 0x0f, 0x3e, 0xce, 0xe1, 0xe1, 0xc3, 0x04, 0x77, 0x9d, 0x8f, 0xf0, 0x1c, 0x3a, 0x2b, 0x46,
	0x63, 0x1c, 0xf8, 0x9b, 0x9b, 0x32, 0x1c, 0xdc, 0x93, 0x09, 0x12, 0x97, 0xb0, 0x7f, 0xc7, 0x2b,
	0x7c, 0x19, 0xc8, 0xf7, 0x1b, 0x3e, 0x1c, 0x3e, 0x90, 0xf5, 0x72, 0x17, 0x6f, 0xa0, 0x2f, 0xd4,
	0xf1, 0xac, 0xd2, 0x79, 0xa0, 0x19, 0xaa, 0x7e, 0x8a, 0x9c, 0x2e, 0xda, 0x5f, 0x8c, 0x9e, 0x8e,
	0xdd, 0x5f, 0x69, 0x1c, 0x4d, 0x76, 0xfc, 0xef, 0xe9, 0xec, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x41, 0x89, 0xc1, 0x2c, 0xac, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChannelOpenerClient is the client API for ChannelOpener service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChannelOpenerClient interface {
	ChannelInformation(ctx context.Context, in *ChannelInformationRequest, opts ...grpc.CallOption) (*ChannelInformationReply, error)
	OpenChannel(ctx context.Context, in *OpenChannelRequest, opts ...grpc.CallOption) (*OpenChannelReply, error)
	RegisterPayment(ctx context.Context, in *RegisterPaymentRequest, opts ...grpc.CallOption) (*RegisterPaymentReply, error)
}

type channelOpenerClient struct {
	cc grpc.ClientConnInterface
}

func NewChannelOpenerClient(cc grpc.ClientConnInterface) ChannelOpenerClient {
	return &channelOpenerClient{cc}
}

func (c *channelOpenerClient) ChannelInformation(ctx context.Context, in *ChannelInformationRequest, opts ...grpc.CallOption) (*ChannelInformationReply, error) {
	out := new(ChannelInformationReply)
	err := c.cc.Invoke(ctx, "/lspd.ChannelOpener/ChannelInformation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelOpenerClient) OpenChannel(ctx context.Context, in *OpenChannelRequest, opts ...grpc.CallOption) (*OpenChannelReply, error) {
	out := new(OpenChannelReply)
	err := c.cc.Invoke(ctx, "/lspd.ChannelOpener/OpenChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelOpenerClient) RegisterPayment(ctx context.Context, in *RegisterPaymentRequest, opts ...grpc.CallOption) (*RegisterPaymentReply, error) {
	out := new(RegisterPaymentReply)
	err := c.cc.Invoke(ctx, "/lspd.ChannelOpener/RegisterPayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChannelOpenerServer is the server API for ChannelOpener service.
type ChannelOpenerServer interface {
	ChannelInformation(context.Context, *ChannelInformationRequest) (*ChannelInformationReply, error)
	OpenChannel(context.Context, *OpenChannelRequest) (*OpenChannelReply, error)
	RegisterPayment(context.Context, *RegisterPaymentRequest) (*RegisterPaymentReply, error)
}

// UnimplementedChannelOpenerServer can be embedded to have forward compatible implementations.
type UnimplementedChannelOpenerServer struct {
}

func (*UnimplementedChannelOpenerServer) ChannelInformation(ctx context.Context, req *ChannelInformationRequest) (*ChannelInformationReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChannelInformation not implemented")
}
func (*UnimplementedChannelOpenerServer) OpenChannel(ctx context.Context, req *OpenChannelRequest) (*OpenChannelReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenChannel not implemented")
}
func (*UnimplementedChannelOpenerServer) RegisterPayment(ctx context.Context, req *RegisterPaymentRequest) (*RegisterPaymentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterPayment not implemented")
}

func RegisterChannelOpenerServer(s *grpc.Server, srv ChannelOpenerServer) {
	s.RegisterService(&_ChannelOpener_serviceDesc, srv)
}

func _ChannelOpener_ChannelInformation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChannelInformationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelOpenerServer).ChannelInformation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lspd.ChannelOpener/ChannelInformation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelOpenerServer).ChannelInformation(ctx, req.(*ChannelInformationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChannelOpener_OpenChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelOpenerServer).OpenChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lspd.ChannelOpener/OpenChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelOpenerServer).OpenChannel(ctx, req.(*OpenChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChannelOpener_RegisterPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelOpenerServer).RegisterPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lspd.ChannelOpener/RegisterPayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelOpenerServer).RegisterPayment(ctx, req.(*RegisterPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChannelOpener_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lspd.ChannelOpener",
	HandlerType: (*ChannelOpenerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ChannelInformation",
			Handler:    _ChannelOpener_ChannelInformation_Handler,
		},
		{
			MethodName: "OpenChannel",
			Handler:    _ChannelOpener_OpenChannel_Handler,
		},
		{
			MethodName: "RegisterPayment",
			Handler:    _ChannelOpener_RegisterPayment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lspd.proto",
}
