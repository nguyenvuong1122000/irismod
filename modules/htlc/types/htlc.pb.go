// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: htlc/htlc.proto

package types

import (
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

// HTLCState defines the state of an HTLC
type HTLCState int32

const (
	// HTLC_STATE_OPEN defines an open state.
	Open HTLCState = 0
	// HTLC_STATE_COMPLETED defines a completed state.
	Completed HTLCState = 1
	// HTLC_STATE_EXPIRED defines an expired state.
	Expired HTLCState = 2
	// HTLC_STATE_REFUNDED defines a refunded state.
	Refunded HTLCState = 3
)

var HTLCState_name = map[int32]string{
	0: "HTLC_STATE_OPEN",
	1: "HTLC_STATE_COMPLETED",
	2: "HTLC_STATE_EXPIRED",
	3: "HTLC_STATE_REFUNDED",
}

var HTLCState_value = map[string]int32{
	"HTLC_STATE_OPEN":      0,
	"HTLC_STATE_COMPLETED": 1,
	"HTLC_STATE_EXPIRED":   2,
	"HTLC_STATE_REFUNDED":  3,
}

func (x HTLCState) String() string {
	return proto.EnumName(HTLCState_name, int32(x))
}

func (HTLCState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c03699801a204f8b, []int{0}
}

// HTLC defines the struct of an HTLC
type HTLC struct {
	Sender               string                                   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	To                   string                                   `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	ReceiverOnOtherChain string                                   `protobuf:"bytes,3,opt,name=receiver_on_other_chain,json=receiverOnOtherChain,proto3" json:"receiver_on_other_chain,omitempty" yaml:"receiver_on_other_chain"`
	Amount               github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
	Secret               string                                   `protobuf:"bytes,5,opt,name=secret,proto3" json:"secret,omitempty"`
	Timestamp            uint64                                   `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	ExpirationHeight     uint64                                   `protobuf:"varint,7,opt,name=expiration_height,json=expirationHeight,proto3" json:"expiration_height,omitempty" yaml:"expiration_height"`
	State                HTLCState                                `protobuf:"varint,8,opt,name=state,proto3,enum=irismod.htlc.HTLCState" json:"state,omitempty"`
}

func (m *HTLC) Reset()         { *m = HTLC{} }
func (m *HTLC) String() string { return proto.CompactTextString(m) }
func (*HTLC) ProtoMessage()    {}
func (*HTLC) Descriptor() ([]byte, []int) {
	return fileDescriptor_c03699801a204f8b, []int{0}
}
func (m *HTLC) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HTLC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HTLC.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HTLC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HTLC.Merge(m, src)
}
func (m *HTLC) XXX_Size() int {
	return m.Size()
}
func (m *HTLC) XXX_DiscardUnknown() {
	xxx_messageInfo_HTLC.DiscardUnknown(m)
}

var xxx_messageInfo_HTLC proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("irismod.htlc.HTLCState", HTLCState_name, HTLCState_value)
	proto.RegisterType((*HTLC)(nil), "irismod.htlc.HTLC")
}

func init() { proto.RegisterFile("htlc/htlc.proto", fileDescriptor_c03699801a204f8b) }

var fileDescriptor_c03699801a204f8b = []byte{
	// 545 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xed, 0xc4, 0x4d, 0x93, 0x6d, 0x69, 0xc3, 0x12, 0x51, 0x63, 0x15, 0xc7, 0x0a, 0x42,
	0x44, 0x48, 0xb1, 0x69, 0xb9, 0xf5, 0x46, 0x12, 0xa3, 0x56, 0x94, 0x26, 0x72, 0x83, 0x04, 0x5c,
	0x2c, 0xc7, 0x1e, 0x62, 0x8b, 0xd8, 0x6b, 0x79, 0x37, 0x15, 0x7d, 0x03, 0x94, 0x13, 0x2f, 0x10,
	0x09, 0x09, 0xb8, 0xf0, 0x24, 0x39, 0xf6, 0xc8, 0x29, 0x40, 0x72, 0xe1, 0xdc, 0x27, 0x40, 0x6b,
	0xa7, 0x34, 0x12, 0xea, 0xc5, 0xde, 0xf9, 0xe7, 0xb3, 0xb5, 0x9f, 0x66, 0xd0, 0xb6, 0xcf, 0x86,
	0xae, 0xc1, 0x1f, 0x7a, 0x9c, 0x10, 0x46, 0xf0, 0x66, 0x90, 0x04, 0x34, 0x24, 0x9e, 0xce, 0x33,
	0x45, 0x75, 0x09, 0x0d, 0x09, 0x35, 0xfa, 0x0e, 0x05, 0xe3, 0x6c, 0xaf, 0x0f, 0xcc, 0xd9, 0x33,
	0x5c, 0x12, 0x44, 0x19, 0xad, 0x54, 0x06, 0x64, 0x40, 0xd2, 0xa3, 0xc1, 0x4f, 0x59, 0x5a, 0xfb,
	0x9a, 0x47, 0xd2, 0x61, 0xef, 0xb8, 0x85, 0xef, 0xa2, 0x02, 0x85, 0xc8, 0x83, 0x44, 0x16, 0x35,
	0xb1, 0x5e, 0xb2, 0x96, 0x15, 0xde, 0x42, 0x39, 0x46, 0xe4, 0x5c, 0x9a, 0xe5, 0x18, 0xc1, 0x6f,
	0xd0, 0x4e, 0x02, 0x2e, 0x04, 0x67, 0x90, 0xd8, 0x24, 0xb2, 0x09, 0xf3, 0x21, 0xb1, 0x5d, 0xdf,
	0x09, 0x22, 0x39, 0xcf, 0xa1, 0x66, 0xed, 0x72, 0x56, 0x55, 0xcf, 0x9d, 0x70, 0x78, 0x50, 0xbb,
	0x01, 0xac, 0x59, 0x95, 0xab, 0x4e, 0x27, 0xea, 0xf0, 0xbc, 0xc5, 0x63, 0xec, 0xa2, 0x82, 0x13,
	0x92, 0x51, 0xc4, 0x64, 0x49, 0xcb, 0xd7, 0x37, 0xf6, 0xef, 0xe9, 0x99, 0x92, 0xce, 0x95, 0xf4,
	0xa5, 0x92, 0xde, 0x22, 0x41, 0xd4, 0x7c, 0x32, 0x9d, 0x55, 0x85, 0xef, 0x3f, 0xab, 0xf5, 0x41,
	0xc0, 0xfc, 0x51, 0x5f, 0x77, 0x49, 0x68, 0x2c, 0xfd, 0xb3, 0x57, 0x83, 0x7a, 0xef, 0x0d, 0x76,
	0x1e, 0x03, 0x4d, 0x3f, 0xa0, 0xd6, 0xf2, 0xd7, 0x99, 0xa7, 0x9b, 0x00, 0x93, 0xd7, 0xae, 0x3c,
	0x79, 0x85, 0x77, 0x51, 0x89, 0x05, 0x21, 0x50, 0xe6, 0x84, 0xb1, 0x5c, 0xd0, 0xc4, 0xba, 0x64,
	0x5d, 0x07, 0xf8, 0x08, 0xdd, 0x86, 0x0f, 0x71, 0x90, 0x38, 0x2c, 0x20, 0x91, 0xed, 0x43, 0x30,
	0xf0, 0x99, 0xbc, 0xce, 0xa9, 0xe6, 0xee, 0xe5, 0xac, 0x2a, 0x67, 0xbe, 0xff, 0x21, 0x35, 0xab,
	0x7c, 0x9d, 0x1d, 0xa6, 0x11, 0x6e, 0xa0, 0x35, 0xca, 0x1c, 0x06, 0x72, 0x51, 0x13, 0xeb, 0x5b,
	0xfb, 0x3b, 0xfa, 0xea, 0x14, 0x75, 0x3e, 0x8b, 0x53, 0xde, 0xb6, 0x32, 0xea, 0x40, 0xfa, 0xf3,
	0xb9, 0x2a, 0x3e, 0xfe, 0x26, 0xa2, 0xd2, 0xbf, 0x16, 0xbe, 0x8f, 0xb6, 0x79, 0x61, 0x9f, 0xf6,
	0x9e, 0xf5, 0x4c, 0xbb, 0xd3, 0x35, 0x4f, 0xca, 0x82, 0x52, 0x1c, 0x4f, 0x34, 0xa9, 0x13, 0x43,
	0x84, 0x1f, 0xa1, 0xca, 0x4a, 0xbb, 0xd5, 0x79, 0xd9, 0x3d, 0x36, 0x7b, 0x66, 0xbb, 0x2c, 0x2a,
	0xb7, 0xc6, 0x13, 0xad, 0xd4, 0x22, 0x61, 0x3c, 0x04, 0x06, 0x1e, 0x7e, 0x80, 0xf0, 0x0a, 0x68,
	0xbe, 0xee, 0x1e, 0x59, 0x66, 0xbb, 0x9c, 0x53, 0x36, 0xc6, 0x13, 0x6d, 0xdd, 0xe4, 0x17, 0x07,
	0x0f, 0x3f, 0x44, 0x77, 0x56, 0x20, 0xcb, 0x7c, 0xfe, 0xea, 0xa4, 0x6d, 0xb6, 0xcb, 0x79, 0x65,
	0x73, 0x3c, 0xd1, 0x8a, 0x16, 0xbc, 0x1b, 0x45, 0x1e, 0x78, 0x8a, 0xf4, 0xf1, 0x8b, 0x2a, 0x34,
	0x5f, 0x4c, 0x7f, 0xab, 0xc2, 0x74, 0xae, 0x8a, 0x17, 0x73, 0x55, 0xfc, 0x35, 0x57, 0xc5, 0x4f,
	0x0b, 0x55, 0xb8, 0x58, 0xa8, 0xc2, 0x8f, 0x85, 0x2a, 0xbc, 0x6d, 0xac, 0x4c, 0x8b, 0x5b, 0x47,
	0xc0, 0x8c, 0xa5, 0xbd, 0x11, 0x12, 0x6f, 0x34, 0x04, 0x9a, 0xee, 0x77, 0x36, 0xb8, 0x7e, 0x21,
	0x5d, 0xd1, 0xa7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x62, 0x9f, 0xa4, 0xf9, 0x02, 0x00,
	0x00,
}

func (this *HTLC) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HTLC)
	if !ok {
		that2, ok := that.(HTLC)
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
	if this.Sender != that1.Sender {
		return false
	}
	if this.To != that1.To {
		return false
	}
	if this.ReceiverOnOtherChain != that1.ReceiverOnOtherChain {
		return false
	}
	if len(this.Amount) != len(that1.Amount) {
		return false
	}
	for i := range this.Amount {
		if !this.Amount[i].Equal(&that1.Amount[i]) {
			return false
		}
	}
	if this.Secret != that1.Secret {
		return false
	}
	if this.Timestamp != that1.Timestamp {
		return false
	}
	if this.ExpirationHeight != that1.ExpirationHeight {
		return false
	}
	if this.State != that1.State {
		return false
	}
	return true
}
func (m *HTLC) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HTLC) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HTLC) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.State != 0 {
		i = encodeVarintHtlc(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x40
	}
	if m.ExpirationHeight != 0 {
		i = encodeVarintHtlc(dAtA, i, uint64(m.ExpirationHeight))
		i--
		dAtA[i] = 0x38
	}
	if m.Timestamp != 0 {
		i = encodeVarintHtlc(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Secret) > 0 {
		i -= len(m.Secret)
		copy(dAtA[i:], m.Secret)
		i = encodeVarintHtlc(dAtA, i, uint64(len(m.Secret)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintHtlc(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.ReceiverOnOtherChain) > 0 {
		i -= len(m.ReceiverOnOtherChain)
		copy(dAtA[i:], m.ReceiverOnOtherChain)
		i = encodeVarintHtlc(dAtA, i, uint64(len(m.ReceiverOnOtherChain)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.To) > 0 {
		i -= len(m.To)
		copy(dAtA[i:], m.To)
		i = encodeVarintHtlc(dAtA, i, uint64(len(m.To)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintHtlc(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintHtlc(dAtA []byte, offset int, v uint64) int {
	offset -= sovHtlc(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HTLC) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovHtlc(uint64(l))
	}
	l = len(m.To)
	if l > 0 {
		n += 1 + l + sovHtlc(uint64(l))
	}
	l = len(m.ReceiverOnOtherChain)
	if l > 0 {
		n += 1 + l + sovHtlc(uint64(l))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovHtlc(uint64(l))
		}
	}
	l = len(m.Secret)
	if l > 0 {
		n += 1 + l + sovHtlc(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sovHtlc(uint64(m.Timestamp))
	}
	if m.ExpirationHeight != 0 {
		n += 1 + sovHtlc(uint64(m.ExpirationHeight))
	}
	if m.State != 0 {
		n += 1 + sovHtlc(uint64(m.State))
	}
	return n
}

func sovHtlc(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHtlc(x uint64) (n int) {
	return sovHtlc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HTLC) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHtlc
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
			return fmt.Errorf("proto: HTLC: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HTLC: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHtlc
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
				return ErrInvalidLengthHtlc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHtlc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHtlc
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
				return ErrInvalidLengthHtlc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHtlc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.To = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReceiverOnOtherChain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHtlc
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
				return ErrInvalidLengthHtlc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHtlc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReceiverOnOtherChain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHtlc
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
				return ErrInvalidLengthHtlc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHtlc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Secret", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHtlc
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
				return ErrInvalidLengthHtlc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHtlc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Secret = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHtlc
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
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpirationHeight", wireType)
			}
			m.ExpirationHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHtlc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpirationHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHtlc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= HTLCState(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHtlc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHtlc
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
func skipHtlc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHtlc
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
					return 0, ErrIntOverflowHtlc
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
					return 0, ErrIntOverflowHtlc
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
				return 0, ErrInvalidLengthHtlc
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHtlc
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHtlc
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHtlc        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHtlc          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHtlc = fmt.Errorf("proto: unexpected end of group")
)
