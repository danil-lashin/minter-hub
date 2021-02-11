// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: oracle/v1/msgs.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Prices struct {
	List []*Price `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (m *Prices) Reset()         { *m = Prices{} }
func (m *Prices) String() string { return proto.CompactTextString(m) }
func (*Prices) ProtoMessage()    {}
func (*Prices) Descriptor() ([]byte, []int) {
	return fileDescriptor_6dda9defd295d067, []int{0}
}
func (m *Prices) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Prices) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Prices.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Prices) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Prices.Merge(m, src)
}
func (m *Prices) XXX_Size() int {
	return m.Size()
}
func (m *Prices) XXX_DiscardUnknown() {
	xxx_messageInfo_Prices.DiscardUnknown(m)
}

var xxx_messageInfo_Prices proto.InternalMessageInfo

func (m *Prices) GetList() []*Price {
	if m != nil {
		return m.List
	}
	return nil
}

type Price struct {
	Name  string                                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=value,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"value"`
}

func (m *Price) Reset()         { *m = Price{} }
func (m *Price) String() string { return proto.CompactTextString(m) }
func (*Price) ProtoMessage()    {}
func (*Price) Descriptor() ([]byte, []int) {
	return fileDescriptor_6dda9defd295d067, []int{1}
}
func (m *Price) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Price) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Price.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Price) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Price.Merge(m, src)
}
func (m *Price) XXX_Size() int {
	return m.Size()
}
func (m *Price) XXX_DiscardUnknown() {
	xxx_messageInfo_Price.DiscardUnknown(m)
}

var xxx_messageInfo_Price proto.InternalMessageInfo

func (m *Price) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// WithdrawClaim claims that a batch of withdrawal
// operations on the bridge contract was executed.
type MsgPriceClaim struct {
	Epoch        uint64  `protobuf:"varint,1,opt,name=epoch,proto3" json:"epoch,omitempty"`
	Prices       *Prices `protobuf:"bytes,2,opt,name=prices,proto3" json:"prices,omitempty"`
	Orchestrator string  `protobuf:"bytes,4,opt,name=orchestrator,proto3" json:"orchestrator,omitempty"`
}

func (m *MsgPriceClaim) Reset()         { *m = MsgPriceClaim{} }
func (m *MsgPriceClaim) String() string { return proto.CompactTextString(m) }
func (*MsgPriceClaim) ProtoMessage()    {}
func (*MsgPriceClaim) Descriptor() ([]byte, []int) {
	return fileDescriptor_6dda9defd295d067, []int{2}
}
func (m *MsgPriceClaim) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgPriceClaim) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgPriceClaim.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgPriceClaim) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPriceClaim.Merge(m, src)
}
func (m *MsgPriceClaim) XXX_Size() int {
	return m.Size()
}
func (m *MsgPriceClaim) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPriceClaim.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPriceClaim proto.InternalMessageInfo

func (m *MsgPriceClaim) GetEpoch() uint64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func (m *MsgPriceClaim) GetPrices() *Prices {
	if m != nil {
		return m.Prices
	}
	return nil
}

func (m *MsgPriceClaim) GetOrchestrator() string {
	if m != nil {
		return m.Orchestrator
	}
	return ""
}

type MsgPriceClaimResponse struct {
}

func (m *MsgPriceClaimResponse) Reset()         { *m = MsgPriceClaimResponse{} }
func (m *MsgPriceClaimResponse) String() string { return proto.CompactTextString(m) }
func (*MsgPriceClaimResponse) ProtoMessage()    {}
func (*MsgPriceClaimResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6dda9defd295d067, []int{3}
}
func (m *MsgPriceClaimResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgPriceClaimResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgPriceClaimResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgPriceClaimResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPriceClaimResponse.Merge(m, src)
}
func (m *MsgPriceClaimResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgPriceClaimResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPriceClaimResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPriceClaimResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Prices)(nil), "oracle.v1.Prices")
	proto.RegisterType((*Price)(nil), "oracle.v1.Price")
	proto.RegisterType((*MsgPriceClaim)(nil), "oracle.v1.MsgPriceClaim")
	proto.RegisterType((*MsgPriceClaimResponse)(nil), "oracle.v1.MsgPriceClaimResponse")
}

func init() { proto.RegisterFile("oracle/v1/msgs.proto", fileDescriptor_6dda9defd295d067) }

var fileDescriptor_6dda9defd295d067 = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x51, 0x41, 0x4b, 0x02, 0x41,
	0x14, 0xde, 0xd5, 0x55, 0x70, 0x2c, 0xa8, 0xc1, 0x6a, 0x91, 0x58, 0x65, 0x89, 0xb0, 0x83, 0x33,
	0x68, 0xff, 0xc0, 0xba, 0x44, 0x08, 0xb1, 0xc7, 0x2e, 0x31, 0xae, 0xc3, 0xec, 0xd2, 0xee, 0xbe,
	0x65, 0x67, 0x94, 0xbc, 0xf6, 0x0b, 0x82, 0xfe, 0x94, 0x47, 0xa1, 0x4b, 0x74, 0x90, 0xd0, 0x7e,
	0x48, 0x38, 0xab, 0xa5, 0x41, 0xa7, 0x79, 0xf3, 0xbd, 0xef, 0x7d, 0xdf, 0xc7, 0x7b, 0xa8, 0x06,
	0x19, 0xf3, 0x23, 0x4e, 0xc7, 0x1d, 0x1a, 0x4b, 0x21, 0x49, 0x9a, 0x81, 0x02, 0x5c, 0xc9, 0x51,
	0x32, 0xee, 0xd4, 0x4f, 0x05, 0x80, 0x88, 0x38, 0x65, 0x69, 0x48, 0x59, 0x92, 0x80, 0x62, 0x2a,
	0x84, 0x64, 0x4d, 0xac, 0xd7, 0x04, 0x08, 0xd0, 0x25, 0x5d, 0x55, 0x39, 0xea, 0x12, 0x54, 0xbe,
	0xcb, 0x42, 0x9f, 0x4b, 0x7c, 0x86, 0xac, 0x28, 0x94, 0xca, 0x2e, 0x34, 0x8b, 0xad, 0x6a, 0xf7,
	0x80, 0xfc, 0xe8, 0x12, 0x4d, 0xf0, 0x74, 0xd7, 0x65, 0xa8, 0xa4, 0xbf, 0x18, 0x23, 0x2b, 0x61,
	0x31, 0xb7, 0xcd, 0xa6, 0xd9, 0xaa, 0x78, 0xba, 0xc6, 0xd7, 0xa8, 0x34, 0x66, 0xd1, 0x88, 0xdb,
	0x85, 0x15, 0xd8, 0x23, 0xd3, 0x79, 0xc3, 0xf8, 0x98, 0x37, 0xce, 0x45, 0xa8, 0x82, 0xd1, 0x80,
	0xf8, 0x10, 0x53, 0x1f, 0x64, 0x0c, 0x72, 0xfd, 0xb4, 0xe5, 0xf0, 0x91, 0xaa, 0x49, 0xca, 0x25,
	0xb9, 0x49, 0x94, 0x97, 0x0f, 0xbb, 0x0a, 0xed, 0xf7, 0xa5, 0xd0, 0x2e, 0x57, 0x11, 0x0b, 0x63,
	0x5c, 0x43, 0x25, 0x9e, 0x82, 0x1f, 0x68, 0x2f, 0xcb, 0xcb, 0x3f, 0xf8, 0x02, 0x95, 0x53, 0x9d,
	0x5c, 0xbb, 0x55, 0xbb, 0x87, 0x7f, 0x13, 0x4b, 0x6f, 0x4d, 0xc0, 0x2e, 0xda, 0x83, 0xcc, 0x0f,
	0xb8, 0x54, 0x19, 0x53, 0x90, 0xd9, 0x96, 0xce, 0xbc, 0x83, 0xb9, 0x27, 0xe8, 0x68, 0xc7, 0xd5,
	0xe3, 0x32, 0x85, 0x44, 0xf2, 0x2e, 0xa0, 0x62, 0x5f, 0x0a, 0x1c, 0x20, 0xb4, 0x15, 0xc9, 0xde,
	0x32, 0xdb, 0x19, 0xab, 0x37, 0xff, 0xeb, 0x6c, 0x04, 0x5d, 0xe7, 0xf9, 0xed, 0xeb, 0xb5, 0x60,
	0xbb, 0xc7, 0xf4, 0xf7, 0xa0, 0x3a, 0xe8, 0x83, 0xbf, 0xe2, 0xf5, 0x6e, 0xa7, 0x0b, 0xc7, 0x9c,
	0x2d, 0x1c, 0xf3, 0x73, 0xe1, 0x98, 0x2f, 0x4b, 0xc7, 0x98, 0x2d, 0x1d, 0xe3, 0x7d, 0xe9, 0x18,
	0xf7, 0x9d, 0xad, 0x45, 0xb2, 0x48, 0x05, 0x9c, 0xb5, 0x13, 0xae, 0x68, 0xca, 0x85, 0x98, 0xd0,
	0x18, 0x86, 0xa3, 0x88, 0xd3, 0xa7, 0x8d, 0xaa, 0xde, 0xeb, 0xa0, 0xac, 0xcf, 0x7c, 0xf9, 0x1d,
	0x00, 0x00, 0xff, 0xff, 0x80, 0x1c, 0x74, 0x6f, 0x3d, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	PriceClaim(ctx context.Context, in *MsgPriceClaim, opts ...grpc.CallOption) (*MsgPriceClaimResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) PriceClaim(ctx context.Context, in *MsgPriceClaim, opts ...grpc.CallOption) (*MsgPriceClaimResponse, error) {
	out := new(MsgPriceClaimResponse)
	err := c.cc.Invoke(ctx, "/oracle.v1.Msg/PriceClaim", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	PriceClaim(context.Context, *MsgPriceClaim) (*MsgPriceClaimResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) PriceClaim(ctx context.Context, req *MsgPriceClaim) (*MsgPriceClaimResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PriceClaim not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_PriceClaim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgPriceClaim)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).PriceClaim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oracle.v1.Msg/PriceClaim",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).PriceClaim(ctx, req.(*MsgPriceClaim))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "oracle.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PriceClaim",
			Handler:    _Msg_PriceClaim_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "oracle/v1/msgs.proto",
}

func (m *Prices) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Prices) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Prices) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.List) > 0 {
		for iNdEx := len(m.List) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.List[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMsgs(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	return len(dAtA) - i, nil
}

func (m *Price) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Price) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Price) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Value.Size()
		i -= size
		if _, err := m.Value.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMsgs(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintMsgs(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgPriceClaim) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgPriceClaim) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgPriceClaim) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Orchestrator) > 0 {
		i -= len(m.Orchestrator)
		copy(dAtA[i:], m.Orchestrator)
		i = encodeVarintMsgs(dAtA, i, uint64(len(m.Orchestrator)))
		i--
		dAtA[i] = 0x22
	}
	if m.Prices != nil {
		{
			size, err := m.Prices.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMsgs(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Epoch != 0 {
		i = encodeVarintMsgs(dAtA, i, uint64(m.Epoch))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgPriceClaimResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgPriceClaimResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgPriceClaimResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintMsgs(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsgs(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Prices) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.List) > 0 {
		for _, e := range m.List {
			l = e.Size()
			n += 1 + l + sovMsgs(uint64(l))
		}
	}
	return n
}

func (m *Price) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovMsgs(uint64(l))
	}
	l = m.Value.Size()
	n += 1 + l + sovMsgs(uint64(l))
	return n
}

func (m *MsgPriceClaim) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Epoch != 0 {
		n += 1 + sovMsgs(uint64(m.Epoch))
	}
	if m.Prices != nil {
		l = m.Prices.Size()
		n += 1 + l + sovMsgs(uint64(l))
	}
	l = len(m.Orchestrator)
	if l > 0 {
		n += 1 + l + sovMsgs(uint64(l))
	}
	return n
}

func (m *MsgPriceClaimResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovMsgs(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsgs(x uint64) (n int) {
	return sovMsgs(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Prices) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgs
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
			return fmt.Errorf("proto: Prices: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Prices: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field List", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
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
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.List = append(m.List, &Price{})
			if err := m.List[len(m.List)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsgs
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMsgs
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
func (m *Price) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgs
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
			return fmt.Errorf("proto: Price: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Price: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
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
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
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
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Value.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsgs
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMsgs
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
func (m *MsgPriceClaim) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgs
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
			return fmt.Errorf("proto: MsgPriceClaim: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgPriceClaim: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epoch", wireType)
			}
			m.Epoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Epoch |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prices", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
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
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Prices == nil {
				m.Prices = &Prices{}
			}
			if err := m.Prices.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Orchestrator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
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
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Orchestrator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsgs
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMsgs
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
func (m *MsgPriceClaimResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgs
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
			return fmt.Errorf("proto: MsgPriceClaimResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgPriceClaimResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsgs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsgs
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMsgs
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
func skipMsgs(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsgs
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
					return 0, ErrIntOverflowMsgs
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
					return 0, ErrIntOverflowMsgs
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
				return 0, ErrInvalidLengthMsgs
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsgs
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsgs
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsgs        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsgs          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsgs = fmt.Errorf("proto: unexpected end of group")
)
