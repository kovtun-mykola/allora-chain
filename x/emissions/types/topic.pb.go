// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: emissions/v1/topic.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	github_com_allora_network_allora_chain_math "github.com/allora-network/allora-chain/math"
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

type Topic struct {
	Id              uint64                                          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Creator         string                                          `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	Metadata        string                                          `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	LossLogic       string                                          `protobuf:"bytes,4,opt,name=loss_logic,json=lossLogic,proto3" json:"loss_logic,omitempty"`
	LossMethod      string                                          `protobuf:"bytes,5,opt,name=loss_method,json=lossMethod,proto3" json:"loss_method,omitempty"`
	InferenceLogic  string                                          `protobuf:"bytes,6,opt,name=inference_logic,json=inferenceLogic,proto3" json:"inference_logic,omitempty"`
	InferenceMethod string                                          `protobuf:"bytes,7,opt,name=inference_method,json=inferenceMethod,proto3" json:"inference_method,omitempty"`
	EpochLastEnded  int64                                           `protobuf:"varint,8,opt,name=epoch_last_ended,json=epochLastEnded,proto3" json:"epoch_last_ended,omitempty"`
	EpochLength     int64                                           `protobuf:"varint,9,opt,name=epoch_length,json=epochLength,proto3" json:"epoch_length,omitempty"`
	GroundTruthLag  int64                                           `protobuf:"varint,10,opt,name=ground_truth_lag,json=groundTruthLag,proto3" json:"ground_truth_lag,omitempty"`
	DefaultArg      string                                          `protobuf:"bytes,11,opt,name=default_arg,json=defaultArg,proto3" json:"default_arg,omitempty"`
	PNorm           github_com_allora_network_allora_chain_math.Dec `protobuf:"bytes,12,opt,name=p_norm,json=pNorm,proto3,customtype=github.com/allora-network/allora-chain/math.Dec" json:"p_norm"`
	AlphaRegret     github_com_allora_network_allora_chain_math.Dec `protobuf:"bytes,13,opt,name=alpha_regret,json=alphaRegret,proto3,customtype=github.com/allora-network/allora-chain/math.Dec" json:"alpha_regret"`
	AllowNegative   bool                                            `protobuf:"varint,14,opt,name=allow_negative,json=allowNegative,proto3" json:"allow_negative,omitempty"`
}

func (m *Topic) Reset()         { *m = Topic{} }
func (m *Topic) String() string { return proto.CompactTextString(m) }
func (*Topic) ProtoMessage()    {}
func (*Topic) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae5610c9d5deb158, []int{0}
}
func (m *Topic) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Topic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Topic.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Topic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Topic.Merge(m, src)
}
func (m *Topic) XXX_Size() int {
	return m.Size()
}
func (m *Topic) XXX_DiscardUnknown() {
	xxx_messageInfo_Topic.DiscardUnknown(m)
}

var xxx_messageInfo_Topic proto.InternalMessageInfo

func (m *Topic) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Topic) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Topic) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

func (m *Topic) GetLossLogic() string {
	if m != nil {
		return m.LossLogic
	}
	return ""
}

func (m *Topic) GetLossMethod() string {
	if m != nil {
		return m.LossMethod
	}
	return ""
}

func (m *Topic) GetInferenceLogic() string {
	if m != nil {
		return m.InferenceLogic
	}
	return ""
}

func (m *Topic) GetInferenceMethod() string {
	if m != nil {
		return m.InferenceMethod
	}
	return ""
}

func (m *Topic) GetEpochLastEnded() int64 {
	if m != nil {
		return m.EpochLastEnded
	}
	return 0
}

func (m *Topic) GetEpochLength() int64 {
	if m != nil {
		return m.EpochLength
	}
	return 0
}

func (m *Topic) GetGroundTruthLag() int64 {
	if m != nil {
		return m.GroundTruthLag
	}
	return 0
}

func (m *Topic) GetDefaultArg() string {
	if m != nil {
		return m.DefaultArg
	}
	return ""
}

func (m *Topic) GetAllowNegative() bool {
	if m != nil {
		return m.AllowNegative
	}
	return false
}

type TopicList struct {
	Topics []*Topic `protobuf:"bytes,1,rep,name=topics,proto3" json:"topics,omitempty"`
}

func (m *TopicList) Reset()         { *m = TopicList{} }
func (m *TopicList) String() string { return proto.CompactTextString(m) }
func (*TopicList) ProtoMessage()    {}
func (*TopicList) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae5610c9d5deb158, []int{1}
}
func (m *TopicList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TopicList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TopicList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TopicList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicList.Merge(m, src)
}
func (m *TopicList) XXX_Size() int {
	return m.Size()
}
func (m *TopicList) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicList.DiscardUnknown(m)
}

var xxx_messageInfo_TopicList proto.InternalMessageInfo

func (m *TopicList) GetTopics() []*Topic {
	if m != nil {
		return m.Topics
	}
	return nil
}

// stores the amount of fees collected by a topic in the last reward epoch
type TopicFeeRevenue struct {
	Epoch   int64                 `protobuf:"varint,1,opt,name=epoch,proto3" json:"epoch,omitempty"`
	Revenue cosmossdk_io_math.Int `protobuf:"bytes,2,opt,name=revenue,proto3,customtype=cosmossdk.io/math.Int" json:"revenue"`
}

func (m *TopicFeeRevenue) Reset()         { *m = TopicFeeRevenue{} }
func (m *TopicFeeRevenue) String() string { return proto.CompactTextString(m) }
func (*TopicFeeRevenue) ProtoMessage()    {}
func (*TopicFeeRevenue) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae5610c9d5deb158, []int{2}
}
func (m *TopicFeeRevenue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TopicFeeRevenue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TopicFeeRevenue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TopicFeeRevenue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicFeeRevenue.Merge(m, src)
}
func (m *TopicFeeRevenue) XXX_Size() int {
	return m.Size()
}
func (m *TopicFeeRevenue) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicFeeRevenue.DiscardUnknown(m)
}

var xxx_messageInfo_TopicFeeRevenue proto.InternalMessageInfo

func (m *TopicFeeRevenue) GetEpoch() int64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func init() {
	proto.RegisterType((*Topic)(nil), "emissions.v1.Topic")
	proto.RegisterType((*TopicList)(nil), "emissions.v1.TopicList")
	proto.RegisterType((*TopicFeeRevenue)(nil), "emissions.v1.TopicFeeRevenue")
}

func init() { proto.RegisterFile("emissions/v1/topic.proto", fileDescriptor_ae5610c9d5deb158) }

var fileDescriptor_ae5610c9d5deb158 = []byte{
	// 576 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xc1, 0x6e, 0xd3, 0x4e,
	0x10, 0xc6, 0xe3, 0xa6, 0x49, 0x9b, 0x4d, 0x9a, 0xf6, 0xbf, 0xff, 0x22, 0x2d, 0x95, 0x70, 0x43,
	0x25, 0x84, 0x01, 0xd5, 0xa6, 0x70, 0xa0, 0x57, 0x2a, 0x40, 0x2a, 0x0a, 0x3d, 0x58, 0x3d, 0xf5,
	0x62, 0x6d, 0xed, 0xe9, 0x7a, 0x55, 0x7b, 0xd7, 0xda, 0xdd, 0xa4, 0xf0, 0x16, 0x3c, 0x06, 0x47,
	0x0e, 0x3c, 0x44, 0x0f, 0x1c, 0x2a, 0x4e, 0x88, 0x43, 0x85, 0xda, 0x03, 0xaf, 0x81, 0xbc, 0xeb,
	0x24, 0x1c, 0x91, 0xb8, 0x58, 0x9e, 0xdf, 0xf7, 0xf9, 0xb3, 0x66, 0x34, 0x83, 0x08, 0x94, 0x5c,
	0x6b, 0x2e, 0x85, 0x8e, 0xa6, 0x7b, 0x91, 0x91, 0x15, 0x4f, 0xc3, 0x4a, 0x49, 0x23, 0xf1, 0x60,
	0xae, 0x84, 0xd3, 0xbd, 0xad, 0xbb, 0xa9, 0xd4, 0xa5, 0xd4, 0x89, 0xd5, 0x22, 0x57, 0x38, 0xe3,
	0xd6, 0x7f, 0xb4, 0xe4, 0x42, 0x46, 0xf6, 0xd9, 0xa0, 0x4d, 0x26, 0x99, 0x74, 0xd6, 0xfa, 0xcd,
	0xd1, 0x9d, 0xaf, 0xcb, 0xa8, 0x73, 0x5c, 0xff, 0x01, 0x0f, 0xd1, 0x12, 0xcf, 0x88, 0x37, 0xf2,
	0x82, 0xe5, 0x78, 0x89, 0x67, 0x98, 0xa0, 0x95, 0x54, 0x01, 0x35, 0x52, 0x91, 0xa5, 0x91, 0x17,
	0xf4, 0xe2, 0x59, 0x89, 0xb7, 0xd0, 0x6a, 0x09, 0x86, 0x66, 0xd4, 0x50, 0xd2, 0xb6, 0xd2, 0xbc,
	0xc6, 0xf7, 0x10, 0x2a, 0xa4, 0xd6, 0x49, 0x21, 0x19, 0x4f, 0xc9, 0xb2, 0x55, 0x7b, 0x35, 0x19,
	0xd7, 0x00, 0x6f, 0xa3, 0xbe, 0x95, 0x4b, 0x30, 0xb9, 0xcc, 0x48, 0xc7, 0xea, 0xf6, 0x8b, 0x77,
	0x96, 0xe0, 0x87, 0x68, 0x9d, 0x8b, 0x33, 0x50, 0x20, 0x52, 0x68, 0x42, 0xba, 0xd6, 0x34, 0x9c,
	0x63, 0x97, 0xf4, 0x08, 0x6d, 0x2c, 0x8c, 0x4d, 0xdc, 0x8a, 0x75, 0x2e, 0x02, 0x9a, 0xcc, 0x00,
	0x6d, 0x40, 0x25, 0xd3, 0x3c, 0x29, 0xa8, 0x36, 0x09, 0x88, 0x0c, 0x32, 0xb2, 0x3a, 0xf2, 0x82,
	0x76, 0x3c, 0xb4, 0x7c, 0x4c, 0xb5, 0x79, 0x5d, 0x53, 0x7c, 0x1f, 0x0d, 0x1a, 0x27, 0x08, 0x66,
	0x72, 0xd2, 0xb3, 0xae, 0xbe, 0x73, 0x59, 0x54, 0x87, 0x31, 0x25, 0x27, 0x22, 0x4b, 0x8c, 0x9a,
	0x98, 0x3a, 0x93, 0x11, 0xe4, 0xc2, 0x1c, 0x3f, 0xae, 0xf1, 0x98, 0xb2, 0xba, 0xd7, 0x0c, 0xce,
	0xe8, 0xa4, 0x30, 0x09, 0x55, 0x8c, 0xf4, 0x5d, 0xaf, 0x0d, 0x7a, 0xa9, 0x18, 0x3e, 0x42, 0xdd,
	0x2a, 0x11, 0x52, 0x95, 0x64, 0x50, 0x6b, 0x07, 0x2f, 0x2e, 0xaf, 0xb7, 0x5b, 0x3f, 0xae, 0xb7,
	0x23, 0xc6, 0x4d, 0x3e, 0x39, 0x0d, 0x53, 0x59, 0x46, 0xb4, 0x28, 0xa4, 0xa2, 0xbb, 0x02, 0xcc,
	0x85, 0x54, 0xe7, 0xb3, 0x32, 0xcd, 0x29, 0x17, 0x51, 0x49, 0x4d, 0x1e, 0xbe, 0x82, 0x34, 0xee,
	0x54, 0x47, 0x52, 0x95, 0xf8, 0x04, 0x0d, 0x68, 0x51, 0xe5, 0x34, 0x51, 0xc0, 0x14, 0x18, 0xb2,
	0xf6, 0x6f, 0xa9, 0x7d, 0x1b, 0x16, 0xdb, 0x2c, 0xfc, 0x00, 0x0d, 0x6b, 0xd7, 0x45, 0x22, 0x80,
	0x51, 0xc3, 0xa7, 0x40, 0x86, 0x23, 0x2f, 0x58, 0x8d, 0xd7, 0x2c, 0x3d, 0x6a, 0xe0, 0xce, 0x3e,
	0xea, 0xd9, 0x6d, 0x1a, 0x73, 0x6d, 0xf0, 0x13, 0xd4, 0xb5, 0xcb, 0xab, 0x89, 0x37, 0x6a, 0x07,
	0xfd, 0x67, 0xff, 0x87, 0x7f, 0xae, 0x6f, 0x68, 0x8d, 0x71, 0x63, 0xd9, 0xd1, 0x68, 0xdd, 0x82,
	0x37, 0x00, 0x31, 0x4c, 0x41, 0x4c, 0x00, 0x6f, 0xa2, 0x8e, 0x9d, 0xbc, 0x5d, 0xca, 0x76, 0xec,
	0x0a, 0xfc, 0x16, 0xad, 0x28, 0x67, 0x70, 0x7b, 0x79, 0xf0, 0xb4, 0x69, 0xf0, 0x8e, 0xbb, 0x00,
	0x9d, 0x9d, 0x87, 0x5c, 0xba, 0x36, 0x0e, 0x85, 0xf9, 0xf6, 0x65, 0x17, 0x35, 0xa7, 0x71, 0x28,
	0xcc, 0xa7, 0x5f, 0x9f, 0x1f, 0x7b, 0xf1, 0x2c, 0xe0, 0x20, 0xbe, 0xbc, 0xf1, 0xbd, 0xab, 0x1b,
	0xdf, 0xfb, 0x79, 0xe3, 0x7b, 0x1f, 0x6f, 0xfd, 0xd6, 0xd5, 0xad, 0xdf, 0xfa, 0x7e, 0xeb, 0xb7,
	0x4e, 0xf6, 0xff, 0x72, 0x5a, 0xef, 0xa3, 0xc5, 0xb1, 0x9a, 0x0f, 0x15, 0xe8, 0xd3, 0xae, 0x3d,
	0xac, 0xe7, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x72, 0x59, 0x3f, 0x19, 0xc6, 0x03, 0x00, 0x00,
}

func (m *Topic) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Topic) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Topic) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AllowNegative {
		i--
		if m.AllowNegative {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x70
	}
	{
		size := m.AlphaRegret.Size()
		i -= size
		if _, err := m.AlphaRegret.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTopic(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x6a
	{
		size := m.PNorm.Size()
		i -= size
		if _, err := m.PNorm.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTopic(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x62
	if len(m.DefaultArg) > 0 {
		i -= len(m.DefaultArg)
		copy(dAtA[i:], m.DefaultArg)
		i = encodeVarintTopic(dAtA, i, uint64(len(m.DefaultArg)))
		i--
		dAtA[i] = 0x5a
	}
	if m.GroundTruthLag != 0 {
		i = encodeVarintTopic(dAtA, i, uint64(m.GroundTruthLag))
		i--
		dAtA[i] = 0x50
	}
	if m.EpochLength != 0 {
		i = encodeVarintTopic(dAtA, i, uint64(m.EpochLength))
		i--
		dAtA[i] = 0x48
	}
	if m.EpochLastEnded != 0 {
		i = encodeVarintTopic(dAtA, i, uint64(m.EpochLastEnded))
		i--
		dAtA[i] = 0x40
	}
	if len(m.InferenceMethod) > 0 {
		i -= len(m.InferenceMethod)
		copy(dAtA[i:], m.InferenceMethod)
		i = encodeVarintTopic(dAtA, i, uint64(len(m.InferenceMethod)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.InferenceLogic) > 0 {
		i -= len(m.InferenceLogic)
		copy(dAtA[i:], m.InferenceLogic)
		i = encodeVarintTopic(dAtA, i, uint64(len(m.InferenceLogic)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.LossMethod) > 0 {
		i -= len(m.LossMethod)
		copy(dAtA[i:], m.LossMethod)
		i = encodeVarintTopic(dAtA, i, uint64(len(m.LossMethod)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.LossLogic) > 0 {
		i -= len(m.LossLogic)
		copy(dAtA[i:], m.LossLogic)
		i = encodeVarintTopic(dAtA, i, uint64(len(m.LossLogic)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Metadata) > 0 {
		i -= len(m.Metadata)
		copy(dAtA[i:], m.Metadata)
		i = encodeVarintTopic(dAtA, i, uint64(len(m.Metadata)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTopic(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintTopic(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TopicList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TopicList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TopicList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Topics) > 0 {
		for iNdEx := len(m.Topics) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Topics[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTopic(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *TopicFeeRevenue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TopicFeeRevenue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TopicFeeRevenue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Revenue.Size()
		i -= size
		if _, err := m.Revenue.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTopic(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Epoch != 0 {
		i = encodeVarintTopic(dAtA, i, uint64(m.Epoch))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTopic(dAtA []byte, offset int, v uint64) int {
	offset -= sovTopic(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Topic) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovTopic(uint64(m.Id))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTopic(uint64(l))
	}
	l = len(m.Metadata)
	if l > 0 {
		n += 1 + l + sovTopic(uint64(l))
	}
	l = len(m.LossLogic)
	if l > 0 {
		n += 1 + l + sovTopic(uint64(l))
	}
	l = len(m.LossMethod)
	if l > 0 {
		n += 1 + l + sovTopic(uint64(l))
	}
	l = len(m.InferenceLogic)
	if l > 0 {
		n += 1 + l + sovTopic(uint64(l))
	}
	l = len(m.InferenceMethod)
	if l > 0 {
		n += 1 + l + sovTopic(uint64(l))
	}
	if m.EpochLastEnded != 0 {
		n += 1 + sovTopic(uint64(m.EpochLastEnded))
	}
	if m.EpochLength != 0 {
		n += 1 + sovTopic(uint64(m.EpochLength))
	}
	if m.GroundTruthLag != 0 {
		n += 1 + sovTopic(uint64(m.GroundTruthLag))
	}
	l = len(m.DefaultArg)
	if l > 0 {
		n += 1 + l + sovTopic(uint64(l))
	}
	l = m.PNorm.Size()
	n += 1 + l + sovTopic(uint64(l))
	l = m.AlphaRegret.Size()
	n += 1 + l + sovTopic(uint64(l))
	if m.AllowNegative {
		n += 2
	}
	return n
}

func (m *TopicList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Topics) > 0 {
		for _, e := range m.Topics {
			l = e.Size()
			n += 1 + l + sovTopic(uint64(l))
		}
	}
	return n
}

func (m *TopicFeeRevenue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Epoch != 0 {
		n += 1 + sovTopic(uint64(m.Epoch))
	}
	l = m.Revenue.Size()
	n += 1 + l + sovTopic(uint64(l))
	return n
}

func sovTopic(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTopic(x uint64) (n int) {
	return sovTopic(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Topic) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTopic
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
			return fmt.Errorf("proto: Topic: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Topic: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopic
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopic
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
				return ErrInvalidLengthTopic
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTopic
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopic
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
				return ErrInvalidLengthTopic
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTopic
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Metadata = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LossLogic", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopic
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
				return ErrInvalidLengthTopic
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTopic
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LossLogic = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LossMethod", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopic
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
				return ErrInvalidLengthTopic
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTopic
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LossMethod = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InferenceLogic", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopic
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
				return ErrInvalidLengthTopic
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTopic
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InferenceLogic = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InferenceMethod", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopic
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
				return ErrInvalidLengthTopic
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTopic
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InferenceMethod = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochLastEnded", wireType)
			}
			m.EpochLastEnded = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopic
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochLastEnded |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochLength", wireType)
			}
			m.EpochLength = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopic
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochLength |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroundTruthLag", wireType)
			}
			m.GroundTruthLag = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopic
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GroundTruthLag |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultArg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopic
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
				return ErrInvalidLengthTopic
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTopic
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DefaultArg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PNorm", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopic
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
				return ErrInvalidLengthTopic
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTopic
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PNorm.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AlphaRegret", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopic
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
				return ErrInvalidLengthTopic
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTopic
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AlphaRegret.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowNegative", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopic
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AllowNegative = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTopic(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTopic
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
func (m *TopicList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTopic
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
			return fmt.Errorf("proto: TopicList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TopicList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Topics", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopic
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
				return ErrInvalidLengthTopic
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTopic
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Topics = append(m.Topics, &Topic{})
			if err := m.Topics[len(m.Topics)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTopic(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTopic
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
func (m *TopicFeeRevenue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTopic
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
			return fmt.Errorf("proto: TopicFeeRevenue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TopicFeeRevenue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epoch", wireType)
			}
			m.Epoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopic
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Epoch |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Revenue", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopic
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
				return ErrInvalidLengthTopic
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTopic
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Revenue.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTopic(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTopic
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
func skipTopic(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTopic
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
					return 0, ErrIntOverflowTopic
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
					return 0, ErrIntOverflowTopic
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
				return 0, ErrInvalidLengthTopic
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTopic
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTopic
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTopic        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTopic          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTopic = fmt.Errorf("proto: unexpected end of group")
)
