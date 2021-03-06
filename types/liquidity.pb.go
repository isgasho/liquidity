// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: liquidity.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
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

type LiquidityPoolType struct {
	PoolTypeIndex     uint32 `protobuf:"varint,1,opt,name=poolTypeIndex,proto3" json:"poolTypeIndex,omitempty" yaml:"pool_type_index"`
	Name              string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty" yaml:"name"`
	MinReserveCoinNum uint32 `protobuf:"varint,3,opt,name=MinReserveCoinNum,proto3" json:"MinReserveCoinNum,omitempty" yaml:"min_reserve_coin_num"`
	MaxReserveCoinNum uint32 `protobuf:"varint,4,opt,name=MaxReserveCoinNum,proto3" json:"MaxReserveCoinNum,omitempty" yaml:"max_reserve_coin_num"`
}

func (m *LiquidityPoolType) Reset()         { *m = LiquidityPoolType{} }
func (m *LiquidityPoolType) String() string { return proto.CompactTextString(m) }
func (*LiquidityPoolType) ProtoMessage()    {}
func (*LiquidityPoolType) Descriptor() ([]byte, []int) {
	return fileDescriptor_65b46dab34d3c00e, []int{0}
}
func (m *LiquidityPoolType) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LiquidityPoolType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LiquidityPoolType.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LiquidityPoolType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LiquidityPoolType.Merge(m, src)
}
func (m *LiquidityPoolType) XXX_Size() int {
	return m.Size()
}
func (m *LiquidityPoolType) XXX_DiscardUnknown() {
	xxx_messageInfo_LiquidityPoolType.DiscardUnknown(m)
}

var xxx_messageInfo_LiquidityPoolType proto.InternalMessageInfo

type Params struct {
	LiquidityPoolTypes       []LiquidityPoolType                    `protobuf:"bytes,1,rep,name=LiquidityPoolTypes,proto3" json:"LiquidityPoolTypes" yaml:"liquidity_pool_types"`
	MinInitDepositToPool     github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=MinInitDepositToPool,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"MinInitDepositToPool" yaml:"min_init_deposit_to_pool"`
	InitPoolCoinMintAmount   github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=InitPoolCoinMintAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"InitPoolCoinMintAmount" yaml:"init_pool_coin_mint_amount"`
	SwapFeeRate              github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=SwapFeeRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"SwapFeeRate" yaml:"swap_fee_rate"`
	LiquidityPoolFeeRate     github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=LiquidityPoolFeeRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"LiquidityPoolFeeRate" yaml:"liquidity_pool_fee_rate"`
	LiquidityPoolCreationFee []types.Coin                           `protobuf:"bytes,6,rep,name=LiquidityPoolCreationFee,proto3" json:"LiquidityPoolCreationFee" yaml:"liquidity_pool_creation_fee"`
	UnitBatchSize            uint32                                 `protobuf:"varint,7,opt,name=UnitBatchSize,proto3" json:"UnitBatchSize,omitempty" yaml:"unit_batch_size"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_65b46dab34d3c00e, []int{1}
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

type LiquidityPool struct {
	PoolID            uint64                                        `protobuf:"varint,1,opt,name=PoolID,proto3" json:"PoolID,omitempty" yaml:"pool_id"`
	PoolTypeIndex     uint32                                        `protobuf:"varint,2,opt,name=poolTypeIndex,proto3" json:"poolTypeIndex,omitempty" yaml:"pool_type_index"`
	ReserveCoinDenoms []string                                      `protobuf:"bytes,3,rep,name=ReserveCoinDenoms,proto3" json:"ReserveCoinDenoms,omitempty" yaml:"reserve_coin_denoms"`
	ReserveAccount    github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,4,opt,name=ReserveAccount,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"ReserveAccount,omitempty" yaml:"reserve_account"`
	PoolCoinDenom     string                                        `protobuf:"bytes,5,opt,name=PoolCoinDenom,proto3" json:"PoolCoinDenom,omitempty" yaml:"pool_coin_denom"`
}

func (m *LiquidityPool) Reset()      { *m = LiquidityPool{} }
func (*LiquidityPool) ProtoMessage() {}
func (*LiquidityPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_65b46dab34d3c00e, []int{2}
}
func (m *LiquidityPool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LiquidityPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LiquidityPool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LiquidityPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LiquidityPool.Merge(m, src)
}
func (m *LiquidityPool) XXX_Size() int {
	return m.Size()
}
func (m *LiquidityPool) XXX_DiscardUnknown() {
	xxx_messageInfo_LiquidityPool.DiscardUnknown(m)
}

var xxx_messageInfo_LiquidityPool proto.InternalMessageInfo

func init() {
	proto.RegisterType((*LiquidityPoolType)(nil), "cosmos.liquidity.LiquidityPoolType")
	proto.RegisterType((*Params)(nil), "cosmos.liquidity.Params")
	proto.RegisterType((*LiquidityPool)(nil), "cosmos.liquidity.LiquidityPool")
}

func init() { proto.RegisterFile("liquidity.proto", fileDescriptor_65b46dab34d3c00e) }

var fileDescriptor_65b46dab34d3c00e = []byte{
	// 744 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xbf, 0x4f, 0xdb, 0x4c,
	0x18, 0x8e, 0x09, 0x5f, 0xbe, 0x2f, 0xc7, 0x07, 0x14, 0x2b, 0xa2, 0x6e, 0x2a, 0xd9, 0xe9, 0xa1,
	0x56, 0x11, 0x12, 0x8e, 0x68, 0x37, 0xa6, 0xc6, 0x44, 0x48, 0x51, 0x09, 0xa2, 0x86, 0x2e, 0x5d,
	0xac, 0x8b, 0x7d, 0x85, 0x53, 0xf1, 0x5d, 0x6a, 0x5f, 0x20, 0x61, 0xab, 0xd4, 0x4e, 0xed, 0xd0,
	0x91, 0x91, 0x3f, 0x87, 0x91, 0xb1, 0xed, 0x60, 0xb5, 0xb0, 0x74, 0xce, 0xd8, 0xa9, 0xba, 0xbb,
	0xfc, 0x4e, 0x90, 0x9a, 0xc9, 0x97, 0x7b, 0xdf, 0xf7, 0x79, 0x1e, 0x3d, 0xef, 0x7b, 0x6f, 0xc0,
	0xf2, 0x09, 0x79, 0xd7, 0x24, 0x01, 0xe1, 0x6d, 0xbb, 0x11, 0x31, 0xce, 0xf4, 0x7b, 0x3e, 0x8b,
	0x43, 0x16, 0xdb, 0xfd, 0xfb, 0x7c, 0xee, 0x88, 0x1d, 0x31, 0x19, 0x2c, 0x89, 0x93, 0xca, 0xcb,
	0xdf, 0x57, 0x79, 0x9e, 0x0a, 0xf8, 0x8c, 0x50, 0x15, 0x80, 0x17, 0x73, 0x60, 0x65, 0xb7, 0x57,
	0xbc, 0xcf, 0xd8, 0xc9, 0x61, 0xbb, 0x81, 0xf5, 0xe7, 0x60, 0xb1, 0xd1, 0x3d, 0x57, 0x69, 0x80,
	0x5b, 0x86, 0x56, 0xd0, 0x8a, 0x8b, 0x4e, 0xbe, 0x93, 0x58, 0xab, 0x6d, 0x14, 0x9e, 0x6c, 0x41,
	0x11, 0xf6, 0x78, 0xbb, 0x81, 0x3d, 0x22, 0x12, 0xa0, 0x3b, 0x5a, 0xa0, 0xaf, 0x81, 0xf9, 0x3d,
	0x14, 0x62, 0x63, 0xae, 0xa0, 0x15, 0xb3, 0xce, 0x72, 0x27, 0xb1, 0x16, 0x54, 0x21, 0x45, 0x21,
	0x86, 0xae, 0x0c, 0xea, 0x35, 0xb0, 0x52, 0x23, 0xd4, 0xc5, 0x31, 0x8e, 0x4e, 0xf1, 0x36, 0x23,
	0x74, 0xaf, 0x19, 0x1a, 0x69, 0x49, 0x65, 0x75, 0x12, 0xeb, 0xa1, 0xaa, 0x08, 0x09, 0xf5, 0x22,
	0x95, 0xe3, 0x09, 0xe5, 0x1e, 0x6d, 0x86, 0xd0, 0x9d, 0xac, 0x94, 0x70, 0xa8, 0x35, 0x06, 0x37,
	0x3f, 0x01, 0x87, 0x5a, 0x53, 0xe1, 0xc6, 0x2b, 0xe1, 0xb7, 0x0c, 0xc8, 0xec, 0xa3, 0x08, 0x85,
	0xb1, 0xde, 0x02, 0xfa, 0x84, 0x49, 0xb1, 0xa1, 0x15, 0xd2, 0xc5, 0x85, 0xa7, 0x6b, 0xf6, 0x78,
	0x0f, 0xec, 0x89, 0x5c, 0x67, 0xed, 0x2a, 0xb1, 0x52, 0x03, 0x0d, 0xfd, 0x54, 0xaf, 0xef, 0x63,
	0x0c, 0xdd, 0x29, 0x1c, 0xfa, 0x47, 0x0d, 0xe4, 0x6a, 0x84, 0x56, 0x29, 0xe1, 0x15, 0xdc, 0x60,
	0x31, 0xe1, 0x87, 0x4c, 0x44, 0xbb, 0xc6, 0xbe, 0x14, 0xb8, 0xdf, 0x13, 0xeb, 0xc9, 0x11, 0xe1,
	0xc7, 0xcd, 0xba, 0xed, 0xb3, 0xb0, 0xa4, 0xe4, 0x74, 0x3f, 0x1b, 0x71, 0xf0, 0xb6, 0x24, 0xf1,
	0xed, 0x2a, 0xe5, 0x9d, 0xc4, 0xb2, 0x06, 0xa6, 0x12, 0x4a, 0xb8, 0x17, 0x28, 0x54, 0x8f, 0x33,
	0xa9, 0x05, 0xba, 0x53, 0xe9, 0xf4, 0x4f, 0x1a, 0x58, 0x15, 0xb7, 0xe2, 0x87, 0x30, 0xa8, 0x46,
	0x28, 0x2f, 0x87, 0xac, 0x49, 0xb9, 0x6c, 0x58, 0xd6, 0x39, 0x98, 0x59, 0xc9, 0x23, 0xa5, 0x44,
	0xaa, 0x90, 0x36, 0xc8, 0x6e, 0x84, 0x84, 0x72, 0x0f, 0x49, 0x64, 0xe8, 0xde, 0x41, 0xa9, 0x1f,
	0x83, 0x85, 0x83, 0x33, 0xd4, 0xd8, 0xc1, 0xd8, 0x45, 0x1c, 0xcb, 0x1e, 0xff, 0xef, 0xec, 0xcc,
	0xa0, 0xa0, 0x82, 0xfd, 0x4e, 0x62, 0xe5, 0x94, 0x82, 0xf8, 0x0c, 0x35, 0xbc, 0x37, 0x18, 0x7b,
	0x11, 0xe2, 0x18, 0xba, 0xc3, 0xd0, 0xfa, 0x07, 0x0d, 0xe4, 0x46, 0xda, 0xd2, 0xe3, 0xfc, 0x47,
	0x72, 0xee, 0xcf, 0xcc, 0x69, 0x4e, 0x9d, 0x80, 0x01, 0xfb, 0x54, 0x36, 0xfd, 0xbd, 0x06, 0x8c,
	0x91, 0xc0, 0x76, 0x84, 0x11, 0x27, 0x8c, 0xee, 0x60, 0x6c, 0x64, 0xe4, 0x1c, 0x3e, 0xe8, 0xcd,
	0x61, 0x1d, 0xc5, 0xd8, 0x3e, 0xdd, 0xac, 0x63, 0x8e, 0x36, 0x6d, 0x61, 0x9c, 0xb3, 0xde, 0x9d,
	0x3e, 0x38, 0x95, 0xdb, 0xef, 0x42, 0x09, 0x11, 0xd0, 0xbd, 0x93, 0x46, 0x2c, 0x85, 0x57, 0x94,
	0x70, 0x07, 0x71, 0xff, 0xf8, 0x80, 0x9c, 0x63, 0xe3, 0xdf, 0xf1, 0xa5, 0xd0, 0x14, 0xad, 0xac,
	0x8b, 0xb8, 0x17, 0x93, 0x73, 0x0c, 0xdd, 0xd1, 0x82, 0xad, 0xff, 0x2e, 0x2e, 0xad, 0xd4, 0xaf,
	0x4b, 0x4b, 0x83, 0x9f, 0xd3, 0x60, 0x71, 0x84, 0x48, 0x5f, 0x07, 0x19, 0xf1, 0xad, 0x56, 0xe4,
	0xae, 0x99, 0x77, 0xf4, 0x4e, 0x62, 0x2d, 0x0d, 0xed, 0x1a, 0x12, 0x40, 0xb7, 0x9b, 0x31, 0xb9,
	0x9e, 0xe6, 0x66, 0x5d, 0x4f, 0xbb, 0x60, 0x65, 0xe8, 0xb5, 0x57, 0x30, 0x65, 0x61, 0x6c, 0xa4,
	0x0b, 0xe9, 0x62, 0xd6, 0x31, 0x3b, 0x89, 0x95, 0x57, 0x28, 0x23, 0x6b, 0x22, 0x90, 0x49, 0xd0,
	0x9d, 0x2c, 0xd4, 0x63, 0xb0, 0xd4, 0xbd, 0x2c, 0xfb, 0xbe, 0x7c, 0x13, 0x6a, 0x22, 0x5f, 0x0c,
	0x04, 0xf5, 0xa0, 0x90, 0x4a, 0x80, 0xbf, 0x13, 0x6b, 0xe3, 0x2f, 0x66, 0xa6, 0xec, 0xfb, 0xe5,
	0x20, 0x88, 0x70, 0x1c, 0xbb, 0x63, 0x14, 0xc2, 0x84, 0xde, 0xcb, 0x90, 0x32, 0xe4, 0x44, 0x66,
	0x27, 0x4c, 0x18, 0x68, 0x87, 0xee, 0x68, 0xc1, 0xa0, 0x1d, 0xce, 0xf6, 0xd5, 0x4f, 0x33, 0x75,
	0x75, 0x63, 0x6a, 0xd7, 0x37, 0xa6, 0xf6, 0xe3, 0xc6, 0xd4, 0xbe, 0xdc, 0x9a, 0xa9, 0xeb, 0x5b,
	0x33, 0xf5, 0xf5, 0xd6, 0x4c, 0xbd, 0x7e, 0x3c, 0x24, 0x94, 0x63, 0x1a, 0xe0, 0x48, 0x3c, 0xd2,
	0x52, 0x7f, 0x8c, 0x94, 0xd6, 0x7a, 0x46, 0xfe, 0xa3, 0x3c, 0xfb, 0x13, 0x00, 0x00, 0xff, 0xff,
	0x06, 0x0c, 0xe0, 0xb7, 0xa5, 0x06, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.LiquidityPoolTypes) != len(that1.LiquidityPoolTypes) {
		return false
	}
	for i := range this.LiquidityPoolTypes {
		if !this.LiquidityPoolTypes[i].Equal(&that1.LiquidityPoolTypes[i]) {
			return false
		}
	}
	if !this.MinInitDepositToPool.Equal(that1.MinInitDepositToPool) {
		return false
	}
	if !this.InitPoolCoinMintAmount.Equal(that1.InitPoolCoinMintAmount) {
		return false
	}
	if !this.SwapFeeRate.Equal(that1.SwapFeeRate) {
		return false
	}
	if !this.LiquidityPoolFeeRate.Equal(that1.LiquidityPoolFeeRate) {
		return false
	}
	if len(this.LiquidityPoolCreationFee) != len(that1.LiquidityPoolCreationFee) {
		return false
	}
	for i := range this.LiquidityPoolCreationFee {
		if !this.LiquidityPoolCreationFee[i].Equal(&that1.LiquidityPoolCreationFee[i]) {
			return false
		}
	}
	if this.UnitBatchSize != that1.UnitBatchSize {
		return false
	}
	return true
}
func (this *LiquidityPool) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LiquidityPool)
	if !ok {
		that2, ok := that.(LiquidityPool)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.PoolID != that1.PoolID {
		return false
	}
	if this.PoolTypeIndex != that1.PoolTypeIndex {
		return false
	}
	if len(this.ReserveCoinDenoms) != len(that1.ReserveCoinDenoms) {
		return false
	}
	for i := range this.ReserveCoinDenoms {
		if this.ReserveCoinDenoms[i] != that1.ReserveCoinDenoms[i] {
			return false
		}
	}
	if !bytes.Equal(this.ReserveAccount, that1.ReserveAccount) {
		return false
	}
	if this.PoolCoinDenom != that1.PoolCoinDenom {
		return false
	}
	return true
}
func (m *LiquidityPoolType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LiquidityPoolType) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LiquidityPoolType) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxReserveCoinNum != 0 {
		i = encodeVarintLiquidity(dAtA, i, uint64(m.MaxReserveCoinNum))
		i--
		dAtA[i] = 0x20
	}
	if m.MinReserveCoinNum != 0 {
		i = encodeVarintLiquidity(dAtA, i, uint64(m.MinReserveCoinNum))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintLiquidity(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.PoolTypeIndex != 0 {
		i = encodeVarintLiquidity(dAtA, i, uint64(m.PoolTypeIndex))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
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
	if m.UnitBatchSize != 0 {
		i = encodeVarintLiquidity(dAtA, i, uint64(m.UnitBatchSize))
		i--
		dAtA[i] = 0x38
	}
	if len(m.LiquidityPoolCreationFee) > 0 {
		for iNdEx := len(m.LiquidityPoolCreationFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LiquidityPoolCreationFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLiquidity(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	{
		size := m.LiquidityPoolFeeRate.Size()
		i -= size
		if _, err := m.LiquidityPoolFeeRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLiquidity(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.SwapFeeRate.Size()
		i -= size
		if _, err := m.SwapFeeRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLiquidity(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.InitPoolCoinMintAmount.Size()
		i -= size
		if _, err := m.InitPoolCoinMintAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLiquidity(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.MinInitDepositToPool.Size()
		i -= size
		if _, err := m.MinInitDepositToPool.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLiquidity(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.LiquidityPoolTypes) > 0 {
		for iNdEx := len(m.LiquidityPoolTypes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LiquidityPoolTypes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLiquidity(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *LiquidityPool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LiquidityPool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LiquidityPool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PoolCoinDenom) > 0 {
		i -= len(m.PoolCoinDenom)
		copy(dAtA[i:], m.PoolCoinDenom)
		i = encodeVarintLiquidity(dAtA, i, uint64(len(m.PoolCoinDenom)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ReserveAccount) > 0 {
		i -= len(m.ReserveAccount)
		copy(dAtA[i:], m.ReserveAccount)
		i = encodeVarintLiquidity(dAtA, i, uint64(len(m.ReserveAccount)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ReserveCoinDenoms) > 0 {
		for iNdEx := len(m.ReserveCoinDenoms) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ReserveCoinDenoms[iNdEx])
			copy(dAtA[i:], m.ReserveCoinDenoms[iNdEx])
			i = encodeVarintLiquidity(dAtA, i, uint64(len(m.ReserveCoinDenoms[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.PoolTypeIndex != 0 {
		i = encodeVarintLiquidity(dAtA, i, uint64(m.PoolTypeIndex))
		i--
		dAtA[i] = 0x10
	}
	if m.PoolID != 0 {
		i = encodeVarintLiquidity(dAtA, i, uint64(m.PoolID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintLiquidity(dAtA []byte, offset int, v uint64) int {
	offset -= sovLiquidity(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LiquidityPoolType) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolTypeIndex != 0 {
		n += 1 + sovLiquidity(uint64(m.PoolTypeIndex))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovLiquidity(uint64(l))
	}
	if m.MinReserveCoinNum != 0 {
		n += 1 + sovLiquidity(uint64(m.MinReserveCoinNum))
	}
	if m.MaxReserveCoinNum != 0 {
		n += 1 + sovLiquidity(uint64(m.MaxReserveCoinNum))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.LiquidityPoolTypes) > 0 {
		for _, e := range m.LiquidityPoolTypes {
			l = e.Size()
			n += 1 + l + sovLiquidity(uint64(l))
		}
	}
	l = m.MinInitDepositToPool.Size()
	n += 1 + l + sovLiquidity(uint64(l))
	l = m.InitPoolCoinMintAmount.Size()
	n += 1 + l + sovLiquidity(uint64(l))
	l = m.SwapFeeRate.Size()
	n += 1 + l + sovLiquidity(uint64(l))
	l = m.LiquidityPoolFeeRate.Size()
	n += 1 + l + sovLiquidity(uint64(l))
	if len(m.LiquidityPoolCreationFee) > 0 {
		for _, e := range m.LiquidityPoolCreationFee {
			l = e.Size()
			n += 1 + l + sovLiquidity(uint64(l))
		}
	}
	if m.UnitBatchSize != 0 {
		n += 1 + sovLiquidity(uint64(m.UnitBatchSize))
	}
	return n
}

func (m *LiquidityPool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolID != 0 {
		n += 1 + sovLiquidity(uint64(m.PoolID))
	}
	if m.PoolTypeIndex != 0 {
		n += 1 + sovLiquidity(uint64(m.PoolTypeIndex))
	}
	if len(m.ReserveCoinDenoms) > 0 {
		for _, s := range m.ReserveCoinDenoms {
			l = len(s)
			n += 1 + l + sovLiquidity(uint64(l))
		}
	}
	l = len(m.ReserveAccount)
	if l > 0 {
		n += 1 + l + sovLiquidity(uint64(l))
	}
	l = len(m.PoolCoinDenom)
	if l > 0 {
		n += 1 + l + sovLiquidity(uint64(l))
	}
	return n
}

func sovLiquidity(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLiquidity(x uint64) (n int) {
	return sovLiquidity(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LiquidityPoolType) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLiquidity
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
			return fmt.Errorf("proto: LiquidityPoolType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LiquidityPoolType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolTypeIndex", wireType)
			}
			m.PoolTypeIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolTypeIndex |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidity
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
				return ErrInvalidLengthLiquidity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinReserveCoinNum", wireType)
			}
			m.MinReserveCoinNum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinReserveCoinNum |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxReserveCoinNum", wireType)
			}
			m.MaxReserveCoinNum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxReserveCoinNum |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLiquidity(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLiquidity
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLiquidity
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLiquidity
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
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidityPoolTypes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidity
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
				return ErrInvalidLengthLiquidity
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LiquidityPoolTypes = append(m.LiquidityPoolTypes, LiquidityPoolType{})
			if err := m.LiquidityPoolTypes[len(m.LiquidityPoolTypes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinInitDepositToPool", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidity
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
				return ErrInvalidLengthLiquidity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinInitDepositToPool.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitPoolCoinMintAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidity
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
				return ErrInvalidLengthLiquidity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InitPoolCoinMintAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapFeeRate", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidity
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
				return ErrInvalidLengthLiquidity
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SwapFeeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidityPoolFeeRate", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidity
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
				return ErrInvalidLengthLiquidity
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LiquidityPoolFeeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidityPoolCreationFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidity
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
				return ErrInvalidLengthLiquidity
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LiquidityPoolCreationFee = append(m.LiquidityPoolCreationFee, types.Coin{})
			if err := m.LiquidityPoolCreationFee[len(m.LiquidityPoolCreationFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnitBatchSize", wireType)
			}
			m.UnitBatchSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UnitBatchSize |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLiquidity(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLiquidity
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLiquidity
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
func (m *LiquidityPool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLiquidity
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
			return fmt.Errorf("proto: LiquidityPool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LiquidityPool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolID", wireType)
			}
			m.PoolID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolTypeIndex", wireType)
			}
			m.PoolTypeIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolTypeIndex |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReserveCoinDenoms", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidity
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
				return ErrInvalidLengthLiquidity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReserveCoinDenoms = append(m.ReserveCoinDenoms, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReserveAccount", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidity
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
				return ErrInvalidLengthLiquidity
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReserveAccount = append(m.ReserveAccount[:0], dAtA[iNdEx:postIndex]...)
			if m.ReserveAccount == nil {
				m.ReserveAccount = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolCoinDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidity
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
				return ErrInvalidLengthLiquidity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolCoinDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLiquidity(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLiquidity
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLiquidity
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
func skipLiquidity(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLiquidity
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
					return 0, ErrIntOverflowLiquidity
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
					return 0, ErrIntOverflowLiquidity
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
				return 0, ErrInvalidLengthLiquidity
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLiquidity
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLiquidity
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLiquidity        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLiquidity          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLiquidity = fmt.Errorf("proto: unexpected end of group")
)
