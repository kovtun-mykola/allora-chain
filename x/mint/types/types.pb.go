// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mint/v1beta1/types.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// Params defines the parameters for the x/mint module.
type Params struct {
	// type of coin to mint
	MintDenom string `protobuf:"bytes,1,opt,name=mint_denom,json=mintDenom,proto3" json:"mint_denom,omitempty"`
	// timestep to recalibrate the emission rate
	// in units of "times per month" e.g. 30 for daily
	// or 4 for weekly
	EmissionCalibrationsTimestepPerMonth uint64 `protobuf:"varint,2,opt,name=emission_calibrations_timestep_per_month,json=emissionCalibrationsTimestepPerMonth,proto3" json:"emission_calibrations_timestep_per_month,omitempty"`
	// maximum total supply of the coin
	MaxSupply cosmossdk_io_math.Int `protobuf:"bytes,3,opt,name=max_supply,json=maxSupply,proto3,customtype=cosmossdk.io/math.Int" json:"max_supply"`
	// ecosystem treasury fraction ideally emitted per unit time
	// this value includes a denominator, see f_emission_denominator
	FEmissionNumerator cosmossdk_io_math.Int `protobuf:"bytes,4,opt,name=f_emission_numerator,json=fEmissionNumerator,proto3,customtype=cosmossdk.io/math.Int" json:"f_emission_numerator"`
	// f_emission is passed as a fraction, this is the denominator
	FEmissionDenominator cosmossdk_io_math.Int `protobuf:"bytes,5,opt,name=f_emission_denominator,json=fEmissionDenominator,proto3,customtype=cosmossdk.io/math.Int" json:"f_emission_denominator"`
	// exponential moving average smoothing factor for one month
	// this value is a fraction, see one_month_smoothing_degree_denominator
	OneMonthSmoothingDegreeNumerator cosmossdk_io_math.Int `protobuf:"bytes,6,opt,name=one_month_smoothing_degree_numerator,json=oneMonthSmoothingDegreeNumerator,proto3,customtype=cosmossdk.io/math.Int" json:"one_month_smoothing_degree_numerator"`
	// denominator for the one_month_smoothing_degree_numerator
	// pass the value as a fraction due to precision issues
	OneMonthSmoothingDegreeDenominator cosmossdk_io_math.Int `protobuf:"bytes,7,opt,name=one_month_smoothing_degree_denominator,json=oneMonthSmoothingDegreeDenominator,proto3,customtype=cosmossdk.io/math.Int" json:"one_month_smoothing_degree_denominator"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_010015e812760429, []int{0}
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

func (m *Params) GetMintDenom() string {
	if m != nil {
		return m.MintDenom
	}
	return ""
}

func (m *Params) GetEmissionCalibrationsTimestepPerMonth() uint64 {
	if m != nil {
		return m.EmissionCalibrationsTimestepPerMonth
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "mint.v1beta1.Params")
}

func init() { proto.RegisterFile("mint/v1beta1/types.proto", fileDescriptor_010015e812760429) }

var fileDescriptor_010015e812760429 = []byte{
	// 458 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x73, 0xd0, 0x06, 0xe5, 0xc4, 0xc2, 0x29, 0x20, 0x13, 0x09, 0x27, 0xaa, 0x2a, 0x14,
	0x55, 0xaa, 0x8f, 0xaa, 0x1b, 0x63, 0x09, 0x43, 0x87, 0x42, 0x95, 0x22, 0x06, 0x96, 0xd3, 0x39,
	0xb9, 0xd8, 0xa7, 0xfa, 0xee, 0x59, 0x77, 0x17, 0x48, 0x37, 0x06, 0x26, 0x26, 0x3e, 0x06, 0x63,
	0x07, 0xbe, 0x03, 0x1d, 0x2b, 0x26, 0xc4, 0x50, 0xa1, 0x64, 0xe8, 0xd7, 0x40, 0x3e, 0x3b, 0xad,
	0x07, 0x60, 0xf0, 0x12, 0xc5, 0xef, 0xfd, 0xf5, 0xfb, 0xff, 0xdf, 0xb3, 0x1f, 0x0e, 0x94, 0xd4,
	0x8e, 0xbe, 0xdf, 0x8b, 0x85, 0xe3, 0x7b, 0xd4, 0x9d, 0xe5, 0xc2, 0x46, 0xb9, 0x01, 0x07, 0xe4,
	0x7e, 0xd1, 0x89, 0xaa, 0x4e, 0xaf, 0x9b, 0x40, 0x02, 0xbe, 0x41, 0x8b, 0x7f, 0xa5, 0xa6, 0xf7,
	0x78, 0x02, 0x56, 0x81, 0x65, 0x65, 0xa3, 0x7c, 0xa8, 0x5a, 0x0f, 0xb8, 0x92, 0x1a, 0xa8, 0xff,
	0x2d, 0x4b, 0x5b, 0xdf, 0x37, 0x71, 0xfb, 0x98, 0x1b, 0xae, 0x2c, 0x79, 0x82, 0x71, 0x81, 0x67,
	0x53, 0xa1, 0x41, 0x05, 0x68, 0x80, 0x86, 0x9d, 0x71, 0xa7, 0xa8, 0x8c, 0x8a, 0x02, 0x79, 0x8b,
	0x87, 0x42, 0x49, 0x6b, 0x25, 0x68, 0x36, 0xe1, 0x99, 0x8c, 0x0d, 0x77, 0x12, 0xb4, 0x65, 0x4e,
	0x2a, 0x61, 0x9d, 0xc8, 0x59, 0x2e, 0x0c, 0x53, 0xa0, 0x5d, 0x1a, 0xdc, 0x19, 0xa0, 0xe1, 0xc6,
	0x78, 0x7b, 0xad, 0x7f, 0x51, 0x93, 0xbf, 0xa9, 0xd4, 0xc7, 0xc2, 0x1c, 0x15, 0x5a, 0xf2, 0x1a,
	0x63, 0xc5, 0x17, 0xcc, 0xce, 0xf3, 0x3c, 0x3b, 0x0b, 0xee, 0x16, 0xb6, 0x07, 0xcf, 0x2e, 0xae,
	0xfa, 0xad, 0x5f, 0x57, 0xfd, 0x87, 0x65, 0x7c, 0x3b, 0x3d, 0x8d, 0x24, 0x50, 0xc5, 0x5d, 0x1a,
	0x1d, 0x6a, 0xf7, 0xe3, 0xdb, 0x2e, 0xae, 0xe6, 0x3a, 0xd4, 0xee, 0xeb, 0xf5, 0xf9, 0x0e, 0x1a,
	0x77, 0x14, 0x5f, 0x9c, 0x78, 0x04, 0x89, 0x71, 0x77, 0xc6, 0x6e, 0xa2, 0xea, 0xb9, 0x12, 0x86,
	0x3b, 0x30, 0xc1, 0x46, 0x43, 0x34, 0x99, 0xbd, 0xac, 0x60, 0xaf, 0xd6, 0x2c, 0x32, 0xc3, 0x8f,
	0x6a, 0x1e, 0x7e, 0x63, 0x52, 0x7b, 0x97, 0xcd, 0x86, 0x2e, 0xdd, 0x1b, 0x97, 0xd1, 0x2d, 0x8d,
	0x7c, 0x44, 0x78, 0x1b, 0xb4, 0x28, 0xd7, 0xca, 0xac, 0x02, 0x70, 0xa9, 0xd4, 0x09, 0x9b, 0x8a,
	0xc4, 0x08, 0x51, 0x1b, 0xae, 0xdd, 0xd0, 0x76, 0x00, 0x5a, 0xf8, 0xf7, 0x70, 0xb2, 0x66, 0x8f,
	0x3c, 0xfa, 0x76, 0xd4, 0x4f, 0x08, 0x3f, 0xfd, 0x4f, 0x84, 0xfa, 0xec, 0xf7, 0x1a, 0x86, 0xd8,
	0xfa, 0x47, 0x88, 0xda, 0x26, 0x9e, 0xf7, 0x3f, 0x5f, 0x9f, 0xef, 0xf4, 0x78, 0x96, 0x81, 0xe1,
	0xbb, 0x93, 0x94, 0x4b, 0x4d, 0x17, 0xd4, 0x1f, 0x4a, 0xf9, 0xf9, 0x1e, 0x1c, 0x5d, 0x2c, 0x43,
	0x74, 0xb9, 0x0c, 0xd1, 0xef, 0x65, 0x88, 0xbe, 0xac, 0xc2, 0xd6, 0xe5, 0x2a, 0x6c, 0xfd, 0x5c,
	0x85, 0xad, 0x77, 0xfb, 0x89, 0x74, 0xe9, 0x3c, 0x8e, 0x26, 0xa0, 0x68, 0x05, 0xd0, 0xc2, 0x7d,
	0x00, 0x73, 0x4a, 0xff, 0xc6, 0xf3, 0x07, 0x17, 0xb7, 0xfd, 0x7d, 0xec, 0xff, 0x09, 0x00, 0x00,
	0xff, 0xff, 0x75, 0xfa, 0xb8, 0x67, 0x8d, 0x03, 0x00, 0x00,
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
		size := m.OneMonthSmoothingDegreeDenominator.Size()
		i -= size
		if _, err := m.OneMonthSmoothingDegreeDenominator.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.OneMonthSmoothingDegreeNumerator.Size()
		i -= size
		if _, err := m.OneMonthSmoothingDegreeNumerator.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.FEmissionDenominator.Size()
		i -= size
		if _, err := m.FEmissionDenominator.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.FEmissionNumerator.Size()
		i -= size
		if _, err := m.FEmissionNumerator.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.MaxSupply.Size()
		i -= size
		if _, err := m.MaxSupply.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.EmissionCalibrationsTimestepPerMonth != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.EmissionCalibrationsTimestepPerMonth))
		i--
		dAtA[i] = 0x10
	}
	if len(m.MintDenom) > 0 {
		i -= len(m.MintDenom)
		copy(dAtA[i:], m.MintDenom)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.MintDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
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
	l = len(m.MintDenom)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.EmissionCalibrationsTimestepPerMonth != 0 {
		n += 1 + sovTypes(uint64(m.EmissionCalibrationsTimestepPerMonth))
	}
	l = m.MaxSupply.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.FEmissionNumerator.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.FEmissionDenominator.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.OneMonthSmoothingDegreeNumerator.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.OneMonthSmoothingDegreeDenominator.Size()
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
				return fmt.Errorf("proto: wrong wireType = %d for field MintDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EmissionCalibrationsTimestepPerMonth", wireType)
			}
			m.EmissionCalibrationsTimestepPerMonth = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EmissionCalibrationsTimestepPerMonth |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSupply", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxSupply.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FEmissionNumerator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FEmissionNumerator.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FEmissionDenominator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FEmissionDenominator.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OneMonthSmoothingDegreeNumerator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OneMonthSmoothingDegreeNumerator.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OneMonthSmoothingDegreeDenominator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OneMonthSmoothingDegreeDenominator.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
