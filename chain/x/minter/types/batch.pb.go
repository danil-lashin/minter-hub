// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: minter/v1/batch.proto

package types

import (
	fmt "fmt"
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

// OutgoingTxBatch represents a batch of transactions going from Peggy to ETH
type OutgoingTxBatch struct {
	BatchNonce   uint64                `protobuf:"varint,1,opt,name=batch_nonce,json=batchNonce,proto3" json:"batch_nonce,omitempty"`
	MinterNonce  uint64                `protobuf:"varint,2,opt,name=minter_nonce,json=minterNonce,proto3" json:"minter_nonce,omitempty"`
	Transactions []*OutgoingTransferTx `protobuf:"bytes,3,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (m *OutgoingTxBatch) Reset()         { *m = OutgoingTxBatch{} }
func (m *OutgoingTxBatch) String() string { return proto.CompactTextString(m) }
func (*OutgoingTxBatch) ProtoMessage()    {}
func (*OutgoingTxBatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfc5e5d918be8aee, []int{0}
}
func (m *OutgoingTxBatch) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OutgoingTxBatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OutgoingTxBatch.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OutgoingTxBatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutgoingTxBatch.Merge(m, src)
}
func (m *OutgoingTxBatch) XXX_Size() int {
	return m.Size()
}
func (m *OutgoingTxBatch) XXX_DiscardUnknown() {
	xxx_messageInfo_OutgoingTxBatch.DiscardUnknown(m)
}

var xxx_messageInfo_OutgoingTxBatch proto.InternalMessageInfo

func (m *OutgoingTxBatch) GetBatchNonce() uint64 {
	if m != nil {
		return m.BatchNonce
	}
	return 0
}

func (m *OutgoingTxBatch) GetMinterNonce() uint64 {
	if m != nil {
		return m.MinterNonce
	}
	return 0
}

func (m *OutgoingTxBatch) GetTransactions() []*OutgoingTransferTx {
	if m != nil {
		return m.Transactions
	}
	return nil
}

// OutgoingTransferTx represents an individual send from Peggy to ETH
type OutgoingTransferTx struct {
	Id          uint64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Sender      string      `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	DestAddress string      `protobuf:"bytes,3,opt,name=dest_address,json=destAddress,proto3" json:"dest_address,omitempty"`
	MinterToken *MinterCoin `protobuf:"bytes,4,opt,name=minter_token,json=minterToken,proto3" json:"minter_token,omitempty"`
}

func (m *OutgoingTransferTx) Reset()         { *m = OutgoingTransferTx{} }
func (m *OutgoingTransferTx) String() string { return proto.CompactTextString(m) }
func (*OutgoingTransferTx) ProtoMessage()    {}
func (*OutgoingTransferTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfc5e5d918be8aee, []int{1}
}
func (m *OutgoingTransferTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OutgoingTransferTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OutgoingTransferTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OutgoingTransferTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutgoingTransferTx.Merge(m, src)
}
func (m *OutgoingTransferTx) XXX_Size() int {
	return m.Size()
}
func (m *OutgoingTransferTx) XXX_DiscardUnknown() {
	xxx_messageInfo_OutgoingTransferTx.DiscardUnknown(m)
}

var xxx_messageInfo_OutgoingTransferTx proto.InternalMessageInfo

func (m *OutgoingTransferTx) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *OutgoingTransferTx) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *OutgoingTransferTx) GetDestAddress() string {
	if m != nil {
		return m.DestAddress
	}
	return ""
}

func (m *OutgoingTransferTx) GetMinterToken() *MinterCoin {
	if m != nil {
		return m.MinterToken
	}
	return nil
}

func init() {
	proto.RegisterType((*OutgoingTxBatch)(nil), "minter.v1.OutgoingTxBatch")
	proto.RegisterType((*OutgoingTransferTx)(nil), "minter.v1.OutgoingTransferTx")
}

func init() { proto.RegisterFile("minter/v1/batch.proto", fileDescriptor_bfc5e5d918be8aee) }

var fileDescriptor_bfc5e5d918be8aee = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xc1, 0x4e, 0x32, 0x31,
	0x10, 0xc7, 0x29, 0x10, 0x12, 0xba, 0xe4, 0xfb, 0x92, 0x26, 0x18, 0xa2, 0x71, 0x45, 0x4e, 0x5c,
	0xdc, 0x06, 0xbc, 0x78, 0x05, 0x8f, 0x46, 0x4d, 0x08, 0x27, 0x2f, 0xa4, 0x6c, 0xc7, 0xa5, 0x11,
	0x5a, 0xb2, 0x1d, 0x08, 0xbc, 0x85, 0x47, 0x7d, 0x23, 0x8f, 0x1c, 0x3d, 0x1a, 0x78, 0x11, 0xd3,
	0xed, 0xba, 0x6a, 0xbc, 0xb5, 0xff, 0xfe, 0x3a, 0xf3, 0x9b, 0x0c, 0x6d, 0x2e, 0x94, 0x46, 0x48,
	0xf9, 0xba, 0xc7, 0xa7, 0x02, 0xe3, 0x59, 0xb4, 0x4c, 0x0d, 0x1a, 0x56, 0xf7, 0x71, 0xb4, 0xee,
	0x1d, 0x9f, 0x7c, 0x13, 0x02, 0x11, 0x2c, 0x0a, 0x54, 0x46, 0x7b, 0xae, 0xf3, 0x42, 0xe8, 0xff,
	0xfb, 0x15, 0x26, 0x46, 0xe9, 0x64, 0xbc, 0x19, 0xba, 0x0a, 0xec, 0x8c, 0x06, 0x59, 0xa9, 0x89,
	0x36, 0x3a, 0x86, 0x16, 0x69, 0x93, 0x6e, 0x75, 0x44, 0xb3, 0xe8, 0xce, 0x25, 0xec, 0x9c, 0x36,
	0x7c, 0xcd, 0x9c, 0x28, 0x67, 0x44, 0xe0, 0x33, 0x8f, 0x0c, 0x68, 0x03, 0x53, 0xa1, 0xad, 0x88,
	0x5d, 0x33, 0xdb, 0xaa, 0xb4, 0x2b, 0xdd, 0xa0, 0x7f, 0x1a, 0x15, 0x5a, 0x51, 0xd1, 0xd5, 0x61,
	0x8f, 0x90, 0x8e, 0x37, 0xa3, 0x5f, 0x5f, 0x3a, 0xaf, 0x84, 0xb2, 0xbf, 0x10, 0xfb, 0x47, 0xcb,
	0x4a, 0xe6, 0x52, 0x65, 0x25, 0xd9, 0x11, 0xad, 0x59, 0xd0, 0x12, 0xd2, 0x4c, 0xa3, 0x3e, 0xca,
	0x6f, 0x4e, 0x52, 0x82, 0xc5, 0x89, 0x90, 0x32, 0x05, 0xeb, 0x0c, 0xdc, 0x6b, 0xe0, 0xb2, 0x81,
	0x8f, 0xd8, 0x55, 0x31, 0x07, 0x9a, 0x27, 0xd0, 0xad, 0x6a, 0x9b, 0x74, 0x83, 0x7e, 0xf3, 0x87,
	0xe4, 0x6d, 0x76, 0xba, 0x36, 0x4a, 0x7f, 0x8d, 0x37, 0x76, 0xe4, 0xf0, 0xe6, 0x6d, 0x1f, 0x92,
	0xdd, 0x3e, 0x24, 0x1f, 0xfb, 0x90, 0x3c, 0x1f, 0xc2, 0xd2, 0xee, 0x10, 0x96, 0xde, 0x0f, 0x61,
	0xe9, 0xa1, 0x97, 0x28, 0x9c, 0xad, 0xa6, 0x51, 0x6c, 0x16, 0x5c, 0xcc, 0x71, 0x06, 0xe2, 0x42,
	0x03, 0xf2, 0x25, 0x24, 0xc9, 0x96, 0x2f, 0x8c, 0x5c, 0xcd, 0x81, 0x6f, 0x78, 0xbe, 0x12, 0xdc,
	0x2e, 0xc1, 0x4e, 0x6b, 0xd9, 0x2a, 0x2e, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x4c, 0x50,
	0x4e, 0xcb, 0x01, 0x00, 0x00,
}

func (m *OutgoingTxBatch) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OutgoingTxBatch) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OutgoingTxBatch) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Transactions) > 0 {
		for iNdEx := len(m.Transactions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Transactions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBatch(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.MinterNonce != 0 {
		i = encodeVarintBatch(dAtA, i, uint64(m.MinterNonce))
		i--
		dAtA[i] = 0x10
	}
	if m.BatchNonce != 0 {
		i = encodeVarintBatch(dAtA, i, uint64(m.BatchNonce))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *OutgoingTransferTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OutgoingTransferTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OutgoingTransferTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MinterToken != nil {
		{
			size, err := m.MinterToken.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBatch(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.DestAddress) > 0 {
		i -= len(m.DestAddress)
		copy(dAtA[i:], m.DestAddress)
		i = encodeVarintBatch(dAtA, i, uint64(len(m.DestAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintBatch(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintBatch(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintBatch(dAtA []byte, offset int, v uint64) int {
	offset -= sovBatch(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OutgoingTxBatch) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BatchNonce != 0 {
		n += 1 + sovBatch(uint64(m.BatchNonce))
	}
	if m.MinterNonce != 0 {
		n += 1 + sovBatch(uint64(m.MinterNonce))
	}
	if len(m.Transactions) > 0 {
		for _, e := range m.Transactions {
			l = e.Size()
			n += 1 + l + sovBatch(uint64(l))
		}
	}
	return n
}

func (m *OutgoingTransferTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovBatch(uint64(m.Id))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovBatch(uint64(l))
	}
	l = len(m.DestAddress)
	if l > 0 {
		n += 1 + l + sovBatch(uint64(l))
	}
	if m.MinterToken != nil {
		l = m.MinterToken.Size()
		n += 1 + l + sovBatch(uint64(l))
	}
	return n
}

func sovBatch(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBatch(x uint64) (n int) {
	return sovBatch(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OutgoingTxBatch) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBatch
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
			return fmt.Errorf("proto: OutgoingTxBatch: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OutgoingTxBatch: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchNonce", wireType)
			}
			m.BatchNonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BatchNonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinterNonce", wireType)
			}
			m.MinterNonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinterNonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transactions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
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
				return ErrInvalidLengthBatch
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Transactions = append(m.Transactions, &OutgoingTransferTx{})
			if err := m.Transactions[len(m.Transactions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBatch(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBatch
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBatch
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
func (m *OutgoingTransferTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBatch
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
			return fmt.Errorf("proto: OutgoingTransferTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OutgoingTransferTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
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
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
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
				return ErrInvalidLengthBatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
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
				return ErrInvalidLengthBatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinterToken", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
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
				return ErrInvalidLengthBatch
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MinterToken == nil {
				m.MinterToken = &MinterCoin{}
			}
			if err := m.MinterToken.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBatch(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBatch
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBatch
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
func skipBatch(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBatch
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
					return 0, ErrIntOverflowBatch
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
					return 0, ErrIntOverflowBatch
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
				return 0, ErrInvalidLengthBatch
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBatch
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBatch
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBatch        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBatch          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBatch = fmt.Errorf("proto: unexpected end of group")
)
