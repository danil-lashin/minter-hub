// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: minter/v1/pool.proto

package types

import (
	fmt "fmt"
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

// OutgoingTx is a withdrawal on the bridged contract
// TODO: can this type be replaced by outgoing transfer tx
type OutgoingTx struct {
	Sender   string     `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	DestAddr string     `protobuf:"bytes,2,opt,name=dest_addr,json=destAddr,proto3" json:"dest_addr,omitempty"`
	Amount   types.Coin `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount"`
}

func (m *OutgoingTx) Reset()         { *m = OutgoingTx{} }
func (m *OutgoingTx) String() string { return proto.CompactTextString(m) }
func (*OutgoingTx) ProtoMessage()    {}
func (*OutgoingTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_86823eeec411a049, []int{0}
}
func (m *OutgoingTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OutgoingTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OutgoingTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OutgoingTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutgoingTx.Merge(m, src)
}
func (m *OutgoingTx) XXX_Size() int {
	return m.Size()
}
func (m *OutgoingTx) XXX_DiscardUnknown() {
	xxx_messageInfo_OutgoingTx.DiscardUnknown(m)
}

var xxx_messageInfo_OutgoingTx proto.InternalMessageInfo

func (m *OutgoingTx) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *OutgoingTx) GetDestAddr() string {
	if m != nil {
		return m.DestAddr
	}
	return ""
}

func (m *OutgoingTx) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

// IDSet represents a set of IDs
type IDSet struct {
	Ids []uint64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (m *IDSet) Reset()         { *m = IDSet{} }
func (m *IDSet) String() string { return proto.CompactTextString(m) }
func (*IDSet) ProtoMessage()    {}
func (*IDSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_86823eeec411a049, []int{1}
}
func (m *IDSet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IDSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IDSet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IDSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IDSet.Merge(m, src)
}
func (m *IDSet) XXX_Size() int {
	return m.Size()
}
func (m *IDSet) XXX_DiscardUnknown() {
	xxx_messageInfo_IDSet.DiscardUnknown(m)
}

var xxx_messageInfo_IDSet proto.InternalMessageInfo

func (m *IDSet) GetIds() []uint64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func init() {
	proto.RegisterType((*OutgoingTx)(nil), "minter.v1.OutgoingTx")
	proto.RegisterType((*IDSet)(nil), "minter.v1.IDSet")
}

func init() { proto.RegisterFile("minter/v1/pool.proto", fileDescriptor_86823eeec411a049) }

var fileDescriptor_86823eeec411a049 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x2c, 0x90, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0x63, 0x5a, 0x2a, 0x6a, 0x16, 0x14, 0x55, 0x28, 0x2d, 0x92, 0xa9, 0x3a, 0x75, 0xc1,
	0x56, 0x60, 0x60, 0xa6, 0xb0, 0x20, 0x06, 0xa4, 0xc2, 0xc4, 0x82, 0x9c, 0xfa, 0xe4, 0x5a, 0x6a,
	0x7c, 0x51, 0xec, 0x54, 0x2d, 0x4f, 0xc1, 0x63, 0x75, 0xec, 0xc8, 0x84, 0x50, 0xf2, 0x22, 0x28,
	0x7f, 0xb6, 0xef, 0x7e, 0xf7, 0x0d, 0xbf, 0x3b, 0x3a, 0x4a, 0x8d, 0xf5, 0x90, 0x8b, 0x6d, 0x2c,
	0x32, 0xc4, 0x0d, 0xcf, 0x72, 0xf4, 0x18, 0x0e, 0x5b, 0xca, 0xb7, 0xf1, 0x84, 0xad, 0xd0, 0xa5,
	0xe8, 0x44, 0x22, 0x1d, 0x88, 0x6d, 0x9c, 0x80, 0x97, 0xb1, 0x58, 0xa1, 0xb1, 0x6d, 0x75, 0x32,
	0xd2, 0xa8, 0xb1, 0x89, 0xa2, 0x4e, 0x2d, 0x9d, 0x7d, 0x51, 0xfa, 0x5a, 0x78, 0x8d, 0xc6, 0xea,
	0xf7, 0x5d, 0x78, 0x49, 0x07, 0x0e, 0xac, 0x82, 0x3c, 0x22, 0x53, 0x32, 0x1f, 0x2e, 0xbb, 0x29,
	0xbc, 0xa2, 0x43, 0x05, 0xce, 0x7f, 0x4a, 0xa5, 0xf2, 0xe8, 0xa4, 0x59, 0x9d, 0xd5, 0xe0, 0x41,
	0xa9, 0x3c, 0xbc, 0xa7, 0x03, 0x99, 0x62, 0x61, 0x7d, 0xd4, 0x9b, 0x92, 0xf9, 0xf9, 0xed, 0x98,
	0xb7, 0x26, 0xbc, 0x36, 0xe1, 0x9d, 0x09, 0x7f, 0x44, 0x63, 0x17, 0xfd, 0xc3, 0xef, 0x75, 0xb0,
	0xec, 0xea, 0xb3, 0x31, 0x3d, 0x7d, 0x7e, 0x7a, 0x03, 0x1f, 0x5e, 0xd0, 0x9e, 0x51, 0x2e, 0x22,
	0xd3, 0xde, 0xbc, 0xbf, 0xac, 0xe3, 0xe2, 0xe5, 0x50, 0x32, 0x72, 0x2c, 0x19, 0xf9, 0x2b, 0x19,
	0xf9, 0xae, 0x58, 0x70, 0xac, 0x58, 0xf0, 0x53, 0xb1, 0xe0, 0x23, 0xd6, 0xc6, 0xaf, 0x8b, 0x84,
	0xaf, 0x30, 0x15, 0x72, 0xe3, 0xd7, 0x20, 0x6f, 0x2c, 0x78, 0x91, 0x81, 0xd6, 0x7b, 0x91, 0xa2,
	0x2a, 0x36, 0x20, 0x76, 0xa2, 0x7b, 0x96, 0xdf, 0x67, 0xe0, 0x92, 0x41, 0x73, 0xea, 0xdd, 0x7f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x57, 0x89, 0x5e, 0x43, 0x01, 0x00, 0x00,
}

func (m *OutgoingTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OutgoingTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OutgoingTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.DestAddr) > 0 {
		i -= len(m.DestAddr)
		copy(dAtA[i:], m.DestAddr)
		i = encodeVarintPool(dAtA, i, uint64(len(m.DestAddr)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintPool(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IDSet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IDSet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IDSet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Ids) > 0 {
		dAtA3 := make([]byte, len(m.Ids)*10)
		var j2 int
		for _, num := range m.Ids {
			for num >= 1<<7 {
				dAtA3[j2] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j2++
			}
			dAtA3[j2] = uint8(num)
			j2++
		}
		i -= j2
		copy(dAtA[i:], dAtA3[:j2])
		i = encodeVarintPool(dAtA, i, uint64(j2))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPool(dAtA []byte, offset int, v uint64) int {
	offset -= sovPool(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OutgoingTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	l = len(m.DestAddr)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovPool(uint64(l))
	return n
}

func (m *IDSet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Ids) > 0 {
		l = 0
		for _, e := range m.Ids {
			l += sovPool(uint64(e))
		}
		n += 1 + sovPool(uint64(l)) + l
	}
	return n
}

func sovPool(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPool(x uint64) (n int) {
	return sovPool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OutgoingTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPool
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
			return fmt.Errorf("proto: OutgoingTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OutgoingTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPool
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPool
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
func (m *IDSet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPool
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
			return fmt.Errorf("proto: IDSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IDSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowPool
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Ids = append(m.Ids, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowPool
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthPool
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthPool
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Ids) == 0 {
					m.Ids = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowPool
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Ids = append(m.Ids, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Ids", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPool
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPool
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
func skipPool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
				return 0, ErrInvalidLengthPool
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPool
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPool
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPool        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPool          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPool = fmt.Errorf("proto: unexpected end of group")
)
