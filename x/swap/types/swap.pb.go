// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gert/swap/v1beta1/swap.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Params defines the parameters for the swap module.
type Params struct {
	// allowed_pools defines that pools that are allowed to be created
	AllowedPools AllowedPools `protobuf:"bytes,1,rep,name=allowed_pools,json=allowedPools,proto3,castrepeated=AllowedPools" json:"allowed_pools"`
	// swap_fee defines the swap fee for all pools
	SwapFee github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=swap_fee,json=swapFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"swap_fee"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_9df359be90eb28cb, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetAllowedPools() AllowedPools {
	if m != nil {
		return m.AllowedPools
	}
	return nil
}

// AllowedPool defines a pool that is allowed to be created
type AllowedPool struct {
	// token_a represents the a token allowed
	TokenA string `protobuf:"bytes,1,opt,name=token_a,json=tokenA,proto3" json:"token_a,omitempty"`
	// token_b represents the b token allowed
	TokenB string `protobuf:"bytes,2,opt,name=token_b,json=tokenB,proto3" json:"token_b,omitempty"`
}

func (m *AllowedPool) Reset()      { *m = AllowedPool{} }
func (*AllowedPool) ProtoMessage() {}
func (*AllowedPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_9df359be90eb28cb, []int{1}
}
func (m *AllowedPool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AllowedPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AllowedPool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AllowedPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllowedPool.Merge(m, src)
}
func (m *AllowedPool) XXX_Size() int {
	return m.Size()
}
func (m *AllowedPool) XXX_DiscardUnknown() {
	xxx_messageInfo_AllowedPool.DiscardUnknown(m)
}

var xxx_messageInfo_AllowedPool proto.InternalMessageInfo

func (m *AllowedPool) GetTokenA() string {
	if m != nil {
		return m.TokenA
	}
	return ""
}

func (m *AllowedPool) GetTokenB() string {
	if m != nil {
		return m.TokenB
	}
	return ""
}

// PoolRecord represents the state of a liquidity pool
// and is used to store the state of a denominated pool
type PoolRecord struct {
	// pool_id represents the unique id of the pool
	PoolID string `protobuf:"bytes,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	// reserves_a is the a token coin reserves
	ReservesA types.Coin `protobuf:"bytes,2,opt,name=reserves_a,json=reservesA,proto3" json:"reserves_a"`
	// reserves_b is the a token coin reserves
	ReservesB types.Coin `protobuf:"bytes,3,opt,name=reserves_b,json=reservesB,proto3" json:"reserves_b"`
	// total_shares is the total distrubuted shares of the pool
	TotalShares github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=total_shares,json=totalShares,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"total_shares"`
}

func (m *PoolRecord) Reset()         { *m = PoolRecord{} }
func (m *PoolRecord) String() string { return proto.CompactTextString(m) }
func (*PoolRecord) ProtoMessage()    {}
func (*PoolRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_9df359be90eb28cb, []int{2}
}
func (m *PoolRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoolRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoolRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoolRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolRecord.Merge(m, src)
}
func (m *PoolRecord) XXX_Size() int {
	return m.Size()
}
func (m *PoolRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolRecord.DiscardUnknown(m)
}

var xxx_messageInfo_PoolRecord proto.InternalMessageInfo

func (m *PoolRecord) GetPoolID() string {
	if m != nil {
		return m.PoolID
	}
	return ""
}

func (m *PoolRecord) GetReservesA() types.Coin {
	if m != nil {
		return m.ReservesA
	}
	return types.Coin{}
}

func (m *PoolRecord) GetReservesB() types.Coin {
	if m != nil {
		return m.ReservesB
	}
	return types.Coin{}
}

// ShareRecord stores the shares owned for a depositor and pool
type ShareRecord struct {
	// depositor represents the owner of the shares
	Depositor github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=depositor,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"depositor,omitempty"`
	// pool_id represents the pool the shares belong to
	PoolID string `protobuf:"bytes,2,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	// shares_owned represents the number of shares owned by depsoitor for the pool_id
	SharesOwned github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=shares_owned,json=sharesOwned,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"shares_owned"`
}

func (m *ShareRecord) Reset()         { *m = ShareRecord{} }
func (m *ShareRecord) String() string { return proto.CompactTextString(m) }
func (*ShareRecord) ProtoMessage()    {}
func (*ShareRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_9df359be90eb28cb, []int{3}
}
func (m *ShareRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ShareRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ShareRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ShareRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShareRecord.Merge(m, src)
}
func (m *ShareRecord) XXX_Size() int {
	return m.Size()
}
func (m *ShareRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ShareRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ShareRecord proto.InternalMessageInfo

func (m *ShareRecord) GetDepositor() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Depositor
	}
	return nil
}

func (m *ShareRecord) GetPoolID() string {
	if m != nil {
		return m.PoolID
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "gert.swap.v1beta1.Params")
	proto.RegisterType((*AllowedPool)(nil), "gert.swap.v1beta1.AllowedPool")
	proto.RegisterType((*PoolRecord)(nil), "gert.swap.v1beta1.PoolRecord")
	proto.RegisterType((*ShareRecord)(nil), "gert.swap.v1beta1.ShareRecord")
}

func init() { proto.RegisterFile("gert/swap/v1beta1/swap.proto", fileDescriptor_9df359be90eb28cb) }

var fileDescriptor_9df359be90eb28cb = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0x3f, 0x6f, 0xd3, 0x4e,
	0x18, 0xc7, 0xe3, 0x34, 0x4a, 0x7e, 0xb9, 0xe4, 0x37, 0x60, 0x2a, 0x91, 0x56, 0xc8, 0xae, 0x82,
	0x84, 0xba, 0xc4, 0x56, 0xcb, 0x86, 0x10, 0x22, 0x26, 0x20, 0x32, 0x51, 0x99, 0x01, 0xc1, 0x72,
	0x3a, 0xdb, 0x4f, 0x53, 0x2b, 0x8e, 0x1f, 0xcb, 0x77, 0x24, 0xf4, 0x5d, 0x20, 0x26, 0x46, 0x66,
	0xe6, 0xbe, 0x02, 0xa6, 0x8e, 0x55, 0x27, 0xc4, 0x10, 0x50, 0xf2, 0x2e, 0x60, 0x41, 0xf7, 0x87,
	0xd6, 0x08, 0x21, 0x51, 0x31, 0xf9, 0x9e, 0xfb, 0xde, 0xf7, 0xf3, 0x3c, 0xf7, 0xb5, 0x8e, 0xdc,
	0x9c, 0xb2, 0x39, 0xf3, 0xf9, 0x82, 0x15, 0xfe, 0x7c, 0x2f, 0x02, 0xc1, 0xf6, 0x54, 0xe1, 0x15,
	0x25, 0x0a, 0xb4, 0xaf, 0x49, 0xd5, 0x53, 0x1b, 0x46, 0xdd, 0xde, 0x9c, 0xe0, 0x04, 0x95, 0xea,
	0xcb, 0x95, 0x3e, 0xb8, 0xed, 0xc4, 0xc8, 0x67, 0xc8, 0xfd, 0x88, 0x71, 0xb8, 0x00, 0xc5, 0x98,
	0xe6, 0x46, 0xdf, 0xd2, 0x3a, 0xd5, 0x46, 0x5d, 0x68, 0xa9, 0xff, 0xd1, 0x22, 0xcd, 0x03, 0x56,
	0xb2, 0x19, 0xb7, 0x5f, 0x90, 0xff, 0x59, 0x96, 0xe1, 0x02, 0x12, 0x5a, 0x20, 0x66, 0xbc, 0x67,
	0xed, 0x6c, 0xec, 0x76, 0xf6, 0x1d, 0xef, 0xb7, 0x31, 0xbc, 0xa1, 0x3e, 0x77, 0x80, 0x98, 0x05,
	0x9b, 0xa7, 0x4b, 0xb7, 0xf6, 0xe1, 0x8b, 0xdb, 0xad, 0x6c, 0xf2, 0xb0, 0xcb, 0x2a, 0x95, 0xfd,
	0x9c, 0xfc, 0x27, 0xfd, 0xf4, 0x10, 0xa0, 0x57, 0xdf, 0xb1, 0x76, 0xdb, 0xc1, 0x3d, 0xe9, 0xfa,
	0xbc, 0x74, 0x6f, 0x4f, 0x52, 0x71, 0xf4, 0x2a, 0xf2, 0x62, 0x9c, 0x99, 0xc1, 0xcc, 0x67, 0xc0,
	0x93, 0xa9, 0x2f, 0x8e, 0x0b, 0xe0, 0xde, 0x08, 0xe2, 0xf3, 0x93, 0x01, 0x31, 0x73, 0x8f, 0x20,
	0x0e, 0x5b, 0x92, 0xf6, 0x18, 0xe0, 0x6e, 0xe3, 0xdd, 0x7b, 0xb7, 0xd6, 0x7f, 0x44, 0x3a, 0x95,
	0xe6, 0xf6, 0x0d, 0xd2, 0x12, 0x38, 0x85, 0x9c, 0xb2, 0x9e, 0x25, 0x9b, 0x85, 0x4d, 0x55, 0x0e,
	0x2f, 0x85, 0x48, 0x4f, 0x61, 0x84, 0xc0, 0x60, 0xde, 0xd6, 0x09, 0x91, 0x80, 0x10, 0x62, 0x2c,
	0x13, 0xfb, 0x16, 0x69, 0xc9, 0x1c, 0x68, 0x9a, 0x68, 0x4c, 0x40, 0x56, 0x4b, 0xb7, 0x29, 0x0f,
	0x8c, 0x47, 0x61, 0x53, 0x4a, 0xe3, 0xc4, 0xbe, 0x4f, 0x48, 0x09, 0x1c, 0xca, 0x39, 0x70, 0xca,
	0x14, 0xb5, 0xb3, 0xbf, 0xe5, 0x99, 0x51, 0xe5, 0xff, 0xb8, 0xc8, 0xec, 0x21, 0xa6, 0x79, 0xd0,
	0x90, 0xd7, 0x0e, 0xdb, 0x3f, 0x2d, 0xc3, 0x5f, 0xfc, 0x51, 0x6f, 0xe3, 0x8a, 0xfe, 0xc0, 0xa6,
	0xa4, 0x2b, 0x50, 0xb0, 0x8c, 0xf2, 0x23, 0x56, 0x02, 0xef, 0x35, 0xae, 0x9c, 0xee, 0x38, 0x17,
	0x95, 0x74, 0xc7, 0xb9, 0x08, 0x3b, 0x8a, 0xf8, 0x4c, 0x01, 0xfb, 0xdf, 0x2d, 0xd2, 0x51, 0x4b,
	0x93, 0xca, 0x21, 0x69, 0x27, 0x50, 0x20, 0x4f, 0x05, 0x96, 0x2a, 0x97, 0x6e, 0xf0, 0xe4, 0xdb,
	0xd2, 0x1d, 0xfc, 0x45, 0xa7, 0x61, 0x1c, 0x0f, 0x93, 0xa4, 0x04, 0xce, 0xcf, 0x4f, 0x06, 0xd7,
	0x4d, 0x43, 0xb3, 0x13, 0x1c, 0x0b, 0xe0, 0xe1, 0x25, 0xba, 0x9a, 0x7e, 0xfd, 0x8f, 0xe9, 0x53,
	0xd2, 0xd5, 0xf7, 0xa6, 0xb8, 0xc8, 0x21, 0x51, 0xf9, 0xfd, 0xf3, 0xed, 0x35, 0xf1, 0xa9, 0x04,
	0x06, 0x0f, 0x4e, 0x57, 0x8e, 0x75, 0xb6, 0x72, 0xac, 0xaf, 0x2b, 0xc7, 0x7a, 0xb3, 0x76, 0x6a,
	0x67, 0x6b, 0xa7, 0xf6, 0x69, 0xed, 0xd4, 0x5e, 0x56, 0xe1, 0xf2, 0x81, 0x0c, 0x32, 0x16, 0x71,
	0xb5, 0xf2, 0x5f, 0xeb, 0x17, 0xad, 0x1a, 0x44, 0x4d, 0xf5, 0xce, 0xee, 0xfc, 0x08, 0x00, 0x00,
	0xff, 0xff, 0xbb, 0x33, 0x82, 0xdb, 0xeb, 0x03, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.SwapFee.Size()
		i -= size
		if _, err := m.SwapFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSwap(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.AllowedPools) > 0 {
		for iNdEx := len(m.AllowedPools) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AllowedPools[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSwap(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *AllowedPool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AllowedPool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AllowedPool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TokenB) > 0 {
		i -= len(m.TokenB)
		copy(dAtA[i:], m.TokenB)
		i = encodeVarintSwap(dAtA, i, uint64(len(m.TokenB)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.TokenA) > 0 {
		i -= len(m.TokenA)
		copy(dAtA[i:], m.TokenA)
		i = encodeVarintSwap(dAtA, i, uint64(len(m.TokenA)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PoolRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoolRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoolRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.TotalShares.Size()
		i -= size
		if _, err := m.TotalShares.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSwap(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.ReservesB.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintSwap(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.ReservesA.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintSwap(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.PoolID) > 0 {
		i -= len(m.PoolID)
		copy(dAtA[i:], m.PoolID)
		i = encodeVarintSwap(dAtA, i, uint64(len(m.PoolID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ShareRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ShareRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ShareRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.SharesOwned.Size()
		i -= size
		if _, err := m.SharesOwned.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSwap(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.PoolID) > 0 {
		i -= len(m.PoolID)
		copy(dAtA[i:], m.PoolID)
		i = encodeVarintSwap(dAtA, i, uint64(len(m.PoolID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Depositor) > 0 {
		i -= len(m.Depositor)
		copy(dAtA[i:], m.Depositor)
		i = encodeVarintSwap(dAtA, i, uint64(len(m.Depositor)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSwap(dAtA []byte, offset int, v uint64) int {
	offset -= sovSwap(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.AllowedPools) > 0 {
		for _, e := range m.AllowedPools {
			l = e.Size()
			n += 1 + l + sovSwap(uint64(l))
		}
	}
	l = m.SwapFee.Size()
	n += 1 + l + sovSwap(uint64(l))
	return n
}

func (m *AllowedPool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TokenA)
	if l > 0 {
		n += 1 + l + sovSwap(uint64(l))
	}
	l = len(m.TokenB)
	if l > 0 {
		n += 1 + l + sovSwap(uint64(l))
	}
	return n
}

func (m *PoolRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PoolID)
	if l > 0 {
		n += 1 + l + sovSwap(uint64(l))
	}
	l = m.ReservesA.Size()
	n += 1 + l + sovSwap(uint64(l))
	l = m.ReservesB.Size()
	n += 1 + l + sovSwap(uint64(l))
	l = m.TotalShares.Size()
	n += 1 + l + sovSwap(uint64(l))
	return n
}

func (m *ShareRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Depositor)
	if l > 0 {
		n += 1 + l + sovSwap(uint64(l))
	}
	l = len(m.PoolID)
	if l > 0 {
		n += 1 + l + sovSwap(uint64(l))
	}
	l = m.SharesOwned.Size()
	n += 1 + l + sovSwap(uint64(l))
	return n
}

func sovSwap(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSwap(x uint64) (n int) {
	return sovSwap(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSwap
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowedPools", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSwap
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AllowedPools = append(m.AllowedPools, AllowedPool{})
			if err := m.AllowedPools[len(m.AllowedPools)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSwap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SwapFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSwap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSwap
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AllowedPool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSwap
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AllowedPool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AllowedPool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenA", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSwap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenA = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenB", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSwap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenB = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSwap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSwap
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PoolRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSwap
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PoolRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoolRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSwap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReservesA", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSwap
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ReservesA.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReservesB", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSwap
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ReservesB.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalShares", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSwap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalShares.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSwap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSwap
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ShareRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSwap
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ShareRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ShareRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Depositor", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSwap
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Depositor = append(m.Depositor[:0], dAtA[iNdEx:postIndex]...)
			if m.Depositor == nil {
				m.Depositor = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSwap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SharesOwned", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSwap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SharesOwned.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSwap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSwap
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSwap(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSwap
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSwap
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSwap
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthSwap
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSwap
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSwap
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSwap        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSwap          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSwap = fmt.Errorf("proto: unexpected end of group")
)
