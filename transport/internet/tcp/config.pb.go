// Code generated by protoc-gen-go.
// source: v2ray.com/core/transport/internet/tcp/config.proto
// DO NOT EDIT!

/*
Package tcp is a generated protocol buffer package.

It is generated from these files:
	v2ray.com/core/transport/internet/tcp/config.proto

It has these top-level messages:
	ConnectionReuse
	Config
*/
package tcp

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v2ray_core_common_serial "v2ray.com/core/common/serial"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ConnectionReuse struct {
	Enable bool `protobuf:"varint,1,opt,name=enable" json:"enable,omitempty"`
}

func (m *ConnectionReuse) Reset()                    { *m = ConnectionReuse{} }
func (m *ConnectionReuse) String() string            { return proto.CompactTextString(m) }
func (*ConnectionReuse) ProtoMessage()               {}
func (*ConnectionReuse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Config struct {
	ConnectionReuse *ConnectionReuse                       `protobuf:"bytes,1,opt,name=connection_reuse,json=connectionReuse" json:"connection_reuse,omitempty"`
	HeaderSettings  *v2ray_core_common_serial.TypedMessage `protobuf:"bytes,2,opt,name=header_settings,json=headerSettings" json:"header_settings,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Config) GetConnectionReuse() *ConnectionReuse {
	if m != nil {
		return m.ConnectionReuse
	}
	return nil
}

func (m *Config) GetHeaderSettings() *v2ray_core_common_serial.TypedMessage {
	if m != nil {
		return m.HeaderSettings
	}
	return nil
}

func init() {
	proto.RegisterType((*ConnectionReuse)(nil), "v2ray.core.transport.internet.tcp.ConnectionReuse")
	proto.RegisterType((*Config)(nil), "v2ray.core.transport.internet.tcp.Config")
}

func init() { proto.RegisterFile("v2ray.com/core/transport/internet/tcp/config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x89, 0x42, 0x90, 0x2d, 0x18, 0xc9, 0x41, 0x8a, 0x27, 0x2d, 0x28, 0x7a, 0xd9, 0x95,
	0x78, 0xf2, 0xda, 0x9e, 0x45, 0x89, 0x9e, 0x04, 0x09, 0xdb, 0xe9, 0x18, 0x03, 0xdd, 0x9d, 0x65,
	0x76, 0x14, 0xfa, 0xcf, 0xfc, 0x79, 0x92, 0x6c, 0x5b, 0x24, 0x97, 0x1e, 0x17, 0xde, 0xf7, 0xf6,
	0x7d, 0xa3, 0xaa, 0x9f, 0x8a, 0xed, 0x46, 0x03, 0x39, 0x03, 0xc4, 0x68, 0x84, 0xad, 0x8f, 0x81,
	0x58, 0x4c, 0xe7, 0x05, 0xd9, 0xa3, 0x18, 0x81, 0x60, 0x80, 0xfc, 0x67, 0xd7, 0xea, 0xc0, 0x24,
	0x54, 0x5e, 0xed, 0x18, 0x46, 0xbd, 0xcf, 0xeb, 0x5d, 0x5e, 0x0b, 0x84, 0x8b, 0xfb, 0x51, 0x2d,
	0x90, 0x73, 0xe4, 0x4d, 0x44, 0xee, 0xec, 0xda, 0xc8, 0x26, 0xe0, 0xaa, 0x71, 0x18, 0xa3, 0x6d,
	0x31, 0x95, 0xce, 0xee, 0x54, 0xb1, 0x20, 0xef, 0x11, 0xa4, 0x23, 0x5f, 0xe3, 0x77, 0xc4, 0xf2,
	0x5c, 0xe5, 0xe8, 0xed, 0x72, 0x8d, 0xd3, 0xec, 0x32, 0xbb, 0x3d, 0xa9, 0xb7, 0xaf, 0xd9, 0x6f,
	0xa6, 0xf2, 0xc5, 0x30, 0xa8, 0xfc, 0x50, 0x67, 0xb0, 0xa7, 0x1a, 0xee, 0xb1, 0x21, 0x3c, 0xa9,
	0x2a, 0x7d, 0x70, 0xa5, 0x1e, 0x7d, 0x58, 0x17, 0x30, 0x5a, 0xf0, 0xac, 0x8a, 0x2f, 0xb4, 0x2b,
	0xe4, 0x26, 0xa2, 0x48, 0xe7, 0xdb, 0x38, 0x3d, 0x1a, 0xda, 0x6f, 0xfe, 0xb7, 0x27, 0x39, 0x9d,
	0xe4, 0xf4, 0x5b, 0x2f, 0xf7, 0x94, 0xdc, 0xea, 0xd3, 0x84, 0xbf, 0x6e, 0xe9, 0xf9, 0xa3, 0xba,
	0x06, 0x72, 0x87, 0xa7, 0xcd, 0x27, 0x49, 0xf0, 0xa5, 0xbf, 0xcd, 0xfb, 0xb1, 0x40, 0x58, 0xe6,
	0xc3, 0x9d, 0x1e, 0xfe, 0x02, 0x00, 0x00, 0xff, 0xff, 0x15, 0x9d, 0x12, 0x13, 0xb2, 0x01, 0x00,
	0x00,
}
