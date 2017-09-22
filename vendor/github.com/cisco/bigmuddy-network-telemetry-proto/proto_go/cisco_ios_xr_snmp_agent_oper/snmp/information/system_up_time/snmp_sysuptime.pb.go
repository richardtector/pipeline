// Code generated by protoc-gen-go. DO NOT EDIT.
// source: snmp_sysuptime.proto

/*
Package cisco_ios_xr_snmp_agent_oper_snmp_information_system_up_time is a generated protocol buffer package.

It is generated from these files:
	snmp_sysuptime.proto

It has these top-level messages:
	SnmpSysuptime_KEYS
	SnmpSysuptime
*/
package cisco_ios_xr_snmp_agent_oper_snmp_information_system_up_time

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// SNMP sysUpTime in centi seconds
type SnmpSysuptime_KEYS struct {
}

func (m *SnmpSysuptime_KEYS) Reset()                    { *m = SnmpSysuptime_KEYS{} }
func (m *SnmpSysuptime_KEYS) String() string            { return proto.CompactTextString(m) }
func (*SnmpSysuptime_KEYS) ProtoMessage()               {}
func (*SnmpSysuptime_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type SnmpSysuptime struct {
	// sysUpTime  1.3.6.1.2.1.1.3
	SystemUpTime uint32 `protobuf:"varint,50,opt,name=system_up_time,json=systemUpTime" json:"system_up_time,omitempty"`
}

func (m *SnmpSysuptime) Reset()                    { *m = SnmpSysuptime{} }
func (m *SnmpSysuptime) String() string            { return proto.CompactTextString(m) }
func (*SnmpSysuptime) ProtoMessage()               {}
func (*SnmpSysuptime) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SnmpSysuptime) GetSystemUpTime() uint32 {
	if m != nil {
		return m.SystemUpTime
	}
	return 0
}

func init() {
	proto.RegisterType((*SnmpSysuptime_KEYS)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.information.system_up_time.snmp_sysuptime_KEYS")
	proto.RegisterType((*SnmpSysuptime)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.information.system_up_time.snmp_sysuptime")
}

func init() { proto.RegisterFile("snmp_sysuptime.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0xce, 0xcb, 0x2d,
	0x88, 0x2f, 0xae, 0x2c, 0x2e, 0x2d, 0x28, 0xc9, 0xcc, 0x4d, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0xb2, 0x49, 0xce, 0x2c, 0x4e, 0xce, 0x8f, 0xcf, 0xcc, 0x2f, 0x8e, 0xaf, 0x28, 0x8a, 0x07,
	0x2b, 0x49, 0x4c, 0x4f, 0xcd, 0x2b, 0x89, 0xcf, 0x2f, 0x48, 0x2d, 0xd2, 0x03, 0xf1, 0xf5, 0x32,
	0xf3, 0xd2, 0xf2, 0x8b, 0x72, 0x13, 0x4b, 0x32, 0xf3, 0xf3, 0xf4, 0x8a, 0x2b, 0x8b, 0x4b, 0x52,
	0x73, 0xe3, 0x4b, 0x0b, 0xe2, 0x41, 0x66, 0x28, 0x89, 0x72, 0x09, 0xa3, 0x9a, 0x1a, 0xef, 0xed,
	0x1a, 0x19, 0xac, 0x64, 0xc6, 0xc5, 0x87, 0x2a, 0x2c, 0xa4, 0xc2, 0xc5, 0x87, 0xaa, 0x55, 0xc2,
	0x48, 0x81, 0x51, 0x83, 0x37, 0x88, 0x07, 0x22, 0x1a, 0x5a, 0x10, 0x92, 0x99, 0x9b, 0x9a, 0xc4,
	0x06, 0x76, 0x93, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x75, 0x20, 0x44, 0x7f, 0xab, 0x00, 0x00,
	0x00,
}
