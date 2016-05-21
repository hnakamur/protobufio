// Code generated by protoc-gen-go.
// source: protobuf_test.proto
// DO NOT EDIT!

/*
Package protobufio is a generated protocol buffer package.

It is generated from these files:
	protobuf_test.proto

It has these top-level messages:
	ExampleMessage
*/
package protobufio

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type ExampleMessage struct {
	Label string  `protobuf:"bytes,1,opt,name=label" json:"label,omitempty"`
	Type  int32   `protobuf:"zigzag32,2,opt,name=type" json:"type,omitempty"`
	Reps  []int64 `protobuf:"zigzag64,3,rep,name=reps" json:"reps,omitempty"`
}

func (m *ExampleMessage) Reset()                    { *m = ExampleMessage{} }
func (m *ExampleMessage) String() string            { return proto.CompactTextString(m) }
func (*ExampleMessage) ProtoMessage()               {}
func (*ExampleMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*ExampleMessage)(nil), "protobufio.ExampleMessage")
}

var fileDescriptor0 = []byte{
	// 119 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0x2a, 0x4d, 0x8b, 0x2f, 0x49, 0x2d, 0x2e, 0xd1, 0x03, 0xf3, 0x84, 0xb8, 0x60, 0x82,
	0x99, 0xf9, 0x4a, 0x7e, 0x5c, 0x7c, 0xae, 0x15, 0x89, 0xb9, 0x05, 0x39, 0xa9, 0xbe, 0xa9, 0xc5,
	0xc5, 0x89, 0xe9, 0xa9, 0x42, 0x22, 0x5c, 0xac, 0x39, 0x89, 0x49, 0xa9, 0x39, 0x12, 0x8c, 0x0a,
	0x8c, 0x1a, 0x9c, 0x41, 0x10, 0x8e, 0x90, 0x10, 0x17, 0x4b, 0x49, 0x65, 0x41, 0xaa, 0x04, 0x13,
	0x50, 0x50, 0x30, 0x08, 0xcc, 0x06, 0x89, 0x15, 0xa5, 0x16, 0x14, 0x4b, 0x30, 0x2b, 0x30, 0x6b,
	0x08, 0x05, 0x81, 0xd9, 0x49, 0x6c, 0x60, 0xb3, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf7,
	0x9a, 0x62, 0x84, 0x79, 0x00, 0x00, 0x00,
}
