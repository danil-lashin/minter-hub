// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: minter/v1/minter_signer.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// SignType defines messages that have been signed by an orchestrator
type SignType int32

const (
	SIGN_TYPE_UNKNOWN                              SignType = 0
	SIGN_TYPE_ORCHESTRATOR_SIGNED_MULTI_SIG_UPDATE SignType = 1
	SIGN_TYPE_ORCHESTRATOR_SIGNED_WITHDRAW_BATCH   SignType = 2
)

var SignType_name = map[int32]string{
	0: "SIGN_TYPE_UNKNOWN",
	1: "SIGN_TYPE_ORCHESTRATOR_SIGNED_MULTI_SIG_UPDATE",
	2: "SIGN_TYPE_ORCHESTRATOR_SIGNED_WITHDRAW_BATCH",
}

var SignType_value = map[string]int32{
	"SIGN_TYPE_UNKNOWN": 0,
	"SIGN_TYPE_ORCHESTRATOR_SIGNED_MULTI_SIG_UPDATE": 1,
	"SIGN_TYPE_ORCHESTRATOR_SIGNED_WITHDRAW_BATCH":   2,
}

func (SignType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_aadd3595ca170a59, []int{0}
}

func init() {
	proto.RegisterEnum("minter.v1.SignType", SignType_name, SignType_value)
}

func init() { proto.RegisterFile("minter/v1/minter_signer.proto", fileDescriptor_aadd3595ca170a59) }

var fileDescriptor_aadd3595ca170a59 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcd, 0xcd, 0xcc, 0x2b,
	0x49, 0x2d, 0xd2, 0x2f, 0x33, 0xd4, 0x87, 0xb0, 0xe2, 0x8b, 0x33, 0xd3, 0xf3, 0x52, 0x8b, 0xf4,
	0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x38, 0x21, 0x82, 0x7a, 0x65, 0x86, 0x52, 0x22, 0xe9, 0xf9,
	0xe9, 0xf9, 0x60, 0x51, 0x7d, 0x10, 0x0b, 0xa2, 0x40, 0x6b, 0x22, 0x23, 0x17, 0x47, 0x70, 0x66,
	0x7a, 0x5e, 0x48, 0x65, 0x41, 0xaa, 0x90, 0x28, 0x97, 0x60, 0xb0, 0xa7, 0xbb, 0x5f, 0x7c, 0x48,
	0x64, 0x80, 0x6b, 0x7c, 0xa8, 0x9f, 0xb7, 0x9f, 0x7f, 0xb8, 0x9f, 0x00, 0x83, 0x90, 0x11, 0x97,
	0x1e, 0x42, 0xd8, 0x3f, 0xc8, 0xd9, 0xc3, 0x35, 0x38, 0x24, 0xc8, 0x31, 0xc4, 0x3f, 0x28, 0x1e,
	0x24, 0xec, 0xea, 0x12, 0xef, 0x1b, 0xea, 0x13, 0xe2, 0x09, 0xe2, 0xc4, 0x87, 0x06, 0xb8, 0x38,
	0x86, 0xb8, 0x0a, 0x30, 0x0a, 0x19, 0x70, 0xe9, 0xe0, 0xd7, 0x13, 0xee, 0x19, 0xe2, 0xe1, 0x12,
	0xe4, 0x18, 0x1e, 0xef, 0xe4, 0x18, 0xe2, 0xec, 0x21, 0xc0, 0x24, 0xc5, 0xd1, 0xb1, 0x58, 0x8e,
	0x61, 0xc5, 0x12, 0x39, 0x06, 0x27, 0xef, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c,
	0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63,
	0x88, 0x32, 0x4c, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcc, 0x29,
	0xc9, 0x48, 0x4d, 0xd4, 0xcd, 0x4b, 0x2d, 0xd1, 0x2f, 0x48, 0x4d, 0x4f, 0xaf, 0xd4, 0xcf, 0xcd,
	0x4f, 0x29, 0xcd, 0x49, 0xd5, 0xaf, 0x80, 0x06, 0x84, 0x7e, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12,
	0x1b, 0xd8, 0x9f, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x9f, 0x45, 0x00, 0x29, 0x01,
	0x00, 0x00,
}