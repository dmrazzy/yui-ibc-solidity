// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: client/ibft2/IBFT2.proto

package ibft2

import (
	_ "."
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	client "pkg/ibc/client"
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

type ClientState struct {
	ChainId         string         `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	IbcStoreAddress []byte         `protobuf:"bytes,2,opt,name=ibc_store_address,json=ibcStoreAddress,proto3" json:"ibc_store_address,omitempty"`
	LatestHeight    *client.Height `protobuf:"bytes,3,opt,name=latest_height,json=latestHeight,proto3" json:"latest_height,omitempty"`
}

func (m *ClientState) Reset()         { *m = ClientState{} }
func (m *ClientState) String() string { return proto.CompactTextString(m) }
func (*ClientState) ProtoMessage()    {}
func (*ClientState) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cc22388094e2414, []int{0}
}
func (m *ClientState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClientState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClientState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClientState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientState.Merge(m, src)
}
func (m *ClientState) XXX_Size() int {
	return m.Size()
}
func (m *ClientState) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientState.DiscardUnknown(m)
}

var xxx_messageInfo_ClientState proto.InternalMessageInfo

func (m *ClientState) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *ClientState) GetIbcStoreAddress() []byte {
	if m != nil {
		return m.IbcStoreAddress
	}
	return nil
}

func (m *ClientState) GetLatestHeight() *client.Height {
	if m != nil {
		return m.LatestHeight
	}
	return nil
}

type ConsensusState struct {
	Timestamp  uint64   `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Root       []byte   `protobuf:"bytes,2,opt,name=root,proto3" json:"root,omitempty"`
	Validators [][]byte `protobuf:"bytes,3,rep,name=validators,proto3" json:"validators,omitempty"`
}

func (m *ConsensusState) Reset()         { *m = ConsensusState{} }
func (m *ConsensusState) String() string { return proto.CompactTextString(m) }
func (*ConsensusState) ProtoMessage()    {}
func (*ConsensusState) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cc22388094e2414, []int{1}
}
func (m *ConsensusState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConsensusState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConsensusState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConsensusState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsensusState.Merge(m, src)
}
func (m *ConsensusState) XXX_Size() int {
	return m.Size()
}
func (m *ConsensusState) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsensusState.DiscardUnknown(m)
}

var xxx_messageInfo_ConsensusState proto.InternalMessageInfo

func (m *ConsensusState) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *ConsensusState) GetRoot() []byte {
	if m != nil {
		return m.Root
	}
	return nil
}

func (m *ConsensusState) GetValidators() [][]byte {
	if m != nil {
		return m.Validators
	}
	return nil
}

type Header struct {
	BesuHeaderRlp     []byte         `protobuf:"bytes,1,opt,name=besu_header_rlp,json=besuHeaderRlp,proto3" json:"besu_header_rlp,omitempty"`
	Seals             [][]byte       `protobuf:"bytes,2,rep,name=seals,proto3" json:"seals,omitempty"`
	TrustedHeight     *client.Height `protobuf:"bytes,3,opt,name=trusted_height,json=trustedHeight,proto3" json:"trusted_height,omitempty"`
	AccountStateProof []byte         `protobuf:"bytes,4,opt,name=account_state_proof,json=accountStateProof,proto3" json:"account_state_proof,omitempty"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cc22388094e2414, []int{2}
}
func (m *Header) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Header.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(m, src)
}
func (m *Header) XXX_Size() int {
	return m.Size()
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func (m *Header) GetBesuHeaderRlp() []byte {
	if m != nil {
		return m.BesuHeaderRlp
	}
	return nil
}

func (m *Header) GetSeals() [][]byte {
	if m != nil {
		return m.Seals
	}
	return nil
}

func (m *Header) GetTrustedHeight() *client.Height {
	if m != nil {
		return m.TrustedHeight
	}
	return nil
}

func (m *Header) GetAccountStateProof() []byte {
	if m != nil {
		return m.AccountStateProof
	}
	return nil
}

func init() {
	proto.RegisterType((*ClientState)(nil), "ibc.lightclients.ibft2.v1.ClientState")
	proto.RegisterType((*ConsensusState)(nil), "ibc.lightclients.ibft2.v1.ConsensusState")
	proto.RegisterType((*Header)(nil), "ibc.lightclients.ibft2.v1.Header")
}

func init() { proto.RegisterFile("client/ibft2/IBFT2.proto", fileDescriptor_0cc22388094e2414) }

var fileDescriptor_0cc22388094e2414 = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x41, 0x6e, 0xd3, 0x40,
	0x14, 0x86, 0x33, 0x24, 0xb4, 0x74, 0x9a, 0xb4, 0xea, 0x34, 0x42, 0x6e, 0x85, 0xac, 0x28, 0x0b,
	0x14, 0x21, 0x6a, 0x8b, 0x70, 0x02, 0x5a, 0x09, 0xa5, 0x3b, 0xe4, 0xb2, 0x62, 0x63, 0xcd, 0x8c,
	0x5f, 0x9a, 0x11, 0xee, 0x8c, 0x35, 0xef, 0xb9, 0xa2, 0xdb, 0x9e, 0x80, 0x53, 0x70, 0x03, 0xee,
	0xc0, 0xb2, 0x4b, 0x96, 0x28, 0xb9, 0x08, 0xf2, 0x4c, 0x24, 0xb2, 0x61, 0x65, 0xff, 0xdf, 0xff,
	0x7b, 0xde, 0xfb, 0xe5, 0xe1, 0x89, 0xae, 0x0d, 0x58, 0xca, 0x8d, 0x5a, 0xd2, 0x3c, 0xbf, 0xbe,
	0xfc, 0xf8, 0x79, 0x9e, 0x35, 0xde, 0x91, 0x13, 0x67, 0x46, 0xe9, 0xac, 0x36, 0xb7, 0x2b, 0x8a,
	0x11, 0xcc, 0x42, 0x26, 0xbb, 0x7f, 0x77, 0x3e, 0x45, 0x57, 0x9b, 0xca, 0xd0, 0xc3, 0x45, 0x88,
	0xaa, 0x76, 0x79, 0x01, 0xdf, 0x08, 0x2c, 0x1a, 0x67, 0x31, 0x7e, 0x7e, 0x7e, 0xba, 0x3d, 0xf8,
	0x2a, 0x3c, 0x22, 0x9c, 0x3e, 0x32, 0x7e, 0x18, 0xc1, 0x0d, 0x49, 0x02, 0x71, 0xc6, 0x5f, 0xe8,
	0x95, 0x34, 0xb6, 0x34, 0x55, 0xc2, 0x26, 0x6c, 0x76, 0x50, 0xec, 0x07, 0x7d, 0x5d, 0x89, 0x37,
	0xfc, 0xc4, 0x28, 0x5d, 0x22, 0x39, 0x0f, 0xa5, 0xac, 0x2a, 0x0f, 0x88, 0xc9, 0xb3, 0x09, 0x9b,
	0x0d, 0x8b, 0x63, 0xa3, 0xf4, 0x4d, 0xc7, 0x3f, 0x44, 0x2c, 0xde, 0xf2, 0x51, 0x2d, 0x09, 0x90,
	0xca, 0x15, 0x74, 0x0b, 0x27, 0xfd, 0x09, 0x9b, 0x1d, 0xce, 0xf7, 0xb3, 0x45, 0x90, 0xc5, 0x30,
	0xba, 0x51, 0x4d, 0x15, 0x3f, 0xba, 0x72, 0x16, 0xc1, 0x62, 0x8b, 0x71, 0x8d, 0x57, 0xfc, 0x80,
	0xcc, 0x1d, 0x20, 0xc9, 0xbb, 0x26, 0xec, 0x31, 0x28, 0xfe, 0x01, 0x21, 0xf8, 0xc0, 0x3b, 0x47,
	0xdb, 0xe1, 0xe1, 0x5d, 0xa4, 0x9c, 0xdf, 0xcb, 0xda, 0x54, 0x92, 0x9c, 0xc7, 0xa4, 0x3f, 0xe9,
	0xcf, 0x86, 0xc5, 0x0e, 0x99, 0xfe, 0x60, 0x7c, 0x6f, 0x01, 0xb2, 0x02, 0x2f, 0x5e, 0xf3, 0x63,
	0x05, 0xd8, 0x96, 0xab, 0x20, 0x4b, 0x5f, 0xc7, 0x11, 0xc3, 0x62, 0xd4, 0xe1, 0x18, 0x2a, 0xea,
	0x46, 0x8c, 0xf9, 0x73, 0x04, 0x59, 0x77, 0x25, 0xbb, 0xd3, 0xa2, 0x10, 0x19, 0x3f, 0x22, 0xdf,
	0x22, 0x41, 0xf5, 0x9f, 0x6e, 0xa3, 0xad, 0x1d, 0xa5, 0xc8, 0xf8, 0xa9, 0xd4, 0xda, 0xb5, 0x96,
	0x4a, 0xec, 0xba, 0x95, 0x8d, 0x77, 0x6e, 0x99, 0x0c, 0xc2, 0xc4, 0x93, 0xad, 0x15, 0x5a, 0x7f,
	0xea, 0x8c, 0xcb, 0xc5, 0xe3, 0xcf, 0xe4, 0x25, 0x1f, 0x6b, 0x67, 0xc9, 0x4b, 0x4d, 0x98, 0x6b,
	0xe7, 0x21, 0xa7, 0x87, 0x06, 0xf0, 0xd7, 0x3a, 0x65, 0x4f, 0xeb, 0x94, 0xfd, 0x59, 0xa7, 0xec,
	0xfb, 0x26, 0xed, 0x3d, 0x6d, 0xd2, 0xde, 0xef, 0x4d, 0xda, 0xfb, 0x32, 0x6e, 0xbe, 0xde, 0xe6,
	0x46, 0xe9, 0x7c, 0xf7, 0xee, 0xa8, 0xbd, 0xf0, 0x8b, 0xdf, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff,
	0xaa, 0x72, 0x97, 0x21, 0x52, 0x02, 0x00, 0x00,
}

func (m *ClientState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClientState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClientState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LatestHeight != nil {
		{
			size, err := m.LatestHeight.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintIBFT2(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.IbcStoreAddress) > 0 {
		i -= len(m.IbcStoreAddress)
		copy(dAtA[i:], m.IbcStoreAddress)
		i = encodeVarintIBFT2(dAtA, i, uint64(len(m.IbcStoreAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintIBFT2(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ConsensusState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConsensusState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConsensusState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Validators) > 0 {
		for iNdEx := len(m.Validators) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Validators[iNdEx])
			copy(dAtA[i:], m.Validators[iNdEx])
			i = encodeVarintIBFT2(dAtA, i, uint64(len(m.Validators[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Root) > 0 {
		i -= len(m.Root)
		copy(dAtA[i:], m.Root)
		i = encodeVarintIBFT2(dAtA, i, uint64(len(m.Root)))
		i--
		dAtA[i] = 0x12
	}
	if m.Timestamp != 0 {
		i = encodeVarintIBFT2(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Header) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Header) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Header) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AccountStateProof) > 0 {
		i -= len(m.AccountStateProof)
		copy(dAtA[i:], m.AccountStateProof)
		i = encodeVarintIBFT2(dAtA, i, uint64(len(m.AccountStateProof)))
		i--
		dAtA[i] = 0x22
	}
	if m.TrustedHeight != nil {
		{
			size, err := m.TrustedHeight.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintIBFT2(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Seals) > 0 {
		for iNdEx := len(m.Seals) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Seals[iNdEx])
			copy(dAtA[i:], m.Seals[iNdEx])
			i = encodeVarintIBFT2(dAtA, i, uint64(len(m.Seals[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.BesuHeaderRlp) > 0 {
		i -= len(m.BesuHeaderRlp)
		copy(dAtA[i:], m.BesuHeaderRlp)
		i = encodeVarintIBFT2(dAtA, i, uint64(len(m.BesuHeaderRlp)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintIBFT2(dAtA []byte, offset int, v uint64) int {
	offset -= sovIBFT2(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClientState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovIBFT2(uint64(l))
	}
	l = len(m.IbcStoreAddress)
	if l > 0 {
		n += 1 + l + sovIBFT2(uint64(l))
	}
	if m.LatestHeight != nil {
		l = m.LatestHeight.Size()
		n += 1 + l + sovIBFT2(uint64(l))
	}
	return n
}

func (m *ConsensusState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Timestamp != 0 {
		n += 1 + sovIBFT2(uint64(m.Timestamp))
	}
	l = len(m.Root)
	if l > 0 {
		n += 1 + l + sovIBFT2(uint64(l))
	}
	if len(m.Validators) > 0 {
		for _, b := range m.Validators {
			l = len(b)
			n += 1 + l + sovIBFT2(uint64(l))
		}
	}
	return n
}

func (m *Header) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BesuHeaderRlp)
	if l > 0 {
		n += 1 + l + sovIBFT2(uint64(l))
	}
	if len(m.Seals) > 0 {
		for _, b := range m.Seals {
			l = len(b)
			n += 1 + l + sovIBFT2(uint64(l))
		}
	}
	if m.TrustedHeight != nil {
		l = m.TrustedHeight.Size()
		n += 1 + l + sovIBFT2(uint64(l))
	}
	l = len(m.AccountStateProof)
	if l > 0 {
		n += 1 + l + sovIBFT2(uint64(l))
	}
	return n
}

func sovIBFT2(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIBFT2(x uint64) (n int) {
	return sovIBFT2(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClientState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIBFT2
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
			return fmt.Errorf("proto: ClientState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIBFT2
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
				return ErrInvalidLengthIBFT2
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIBFT2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcStoreAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIBFT2
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
				return ErrInvalidLengthIBFT2
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIBFT2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IbcStoreAddress = append(m.IbcStoreAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.IbcStoreAddress == nil {
				m.IbcStoreAddress = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestHeight", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIBFT2
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
				return ErrInvalidLengthIBFT2
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIBFT2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LatestHeight == nil {
				m.LatestHeight = &client.Height{}
			}
			if err := m.LatestHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIBFT2(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIBFT2
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
func (m *ConsensusState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIBFT2
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
			return fmt.Errorf("proto: ConsensusState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConsensusState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIBFT2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Root", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIBFT2
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
				return ErrInvalidLengthIBFT2
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIBFT2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Root = append(m.Root[:0], dAtA[iNdEx:postIndex]...)
			if m.Root == nil {
				m.Root = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validators", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIBFT2
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
				return ErrInvalidLengthIBFT2
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIBFT2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validators = append(m.Validators, make([]byte, postIndex-iNdEx))
			copy(m.Validators[len(m.Validators)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIBFT2(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIBFT2
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
func (m *Header) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIBFT2
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
			return fmt.Errorf("proto: Header: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Header: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BesuHeaderRlp", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIBFT2
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
				return ErrInvalidLengthIBFT2
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIBFT2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BesuHeaderRlp = append(m.BesuHeaderRlp[:0], dAtA[iNdEx:postIndex]...)
			if m.BesuHeaderRlp == nil {
				m.BesuHeaderRlp = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seals", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIBFT2
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
				return ErrInvalidLengthIBFT2
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIBFT2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Seals = append(m.Seals, make([]byte, postIndex-iNdEx))
			copy(m.Seals[len(m.Seals)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrustedHeight", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIBFT2
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
				return ErrInvalidLengthIBFT2
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIBFT2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TrustedHeight == nil {
				m.TrustedHeight = &client.Height{}
			}
			if err := m.TrustedHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountStateProof", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIBFT2
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
				return ErrInvalidLengthIBFT2
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIBFT2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccountStateProof = append(m.AccountStateProof[:0], dAtA[iNdEx:postIndex]...)
			if m.AccountStateProof == nil {
				m.AccountStateProof = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIBFT2(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIBFT2
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
func skipIBFT2(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIBFT2
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
					return 0, ErrIntOverflowIBFT2
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
					return 0, ErrIntOverflowIBFT2
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
				return 0, ErrInvalidLengthIBFT2
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIBFT2
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIBFT2
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIBFT2        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIBFT2          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIBFT2 = fmt.Errorf("proto: unexpected end of group")
)
