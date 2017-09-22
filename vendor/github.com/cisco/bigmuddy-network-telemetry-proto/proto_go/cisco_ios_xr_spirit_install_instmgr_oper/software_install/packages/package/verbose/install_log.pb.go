// Code generated by protoc-gen-go. DO NOT EDIT.
// source: install_log.proto

/*
Package cisco_ios_xr_spirit_install_instmgr_oper_software_install_packages_package_verbose is a generated protocol buffer package.

It is generated from these files:
	install_log.proto

It has these top-level messages:
	InstallLog_KEYS
	InstallLog
*/
package cisco_ios_xr_spirit_install_instmgr_oper_software_install_packages_package_verbose

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

// Install Log
type InstallLog_KEYS struct {
	PackageName string `protobuf:"bytes,1,opt,name=package_name,json=packageName" json:"package_name,omitempty"`
}

func (m *InstallLog_KEYS) Reset()                    { *m = InstallLog_KEYS{} }
func (m *InstallLog_KEYS) String() string            { return proto.CompactTextString(m) }
func (*InstallLog_KEYS) ProtoMessage()               {}
func (*InstallLog_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *InstallLog_KEYS) GetPackageName() string {
	if m != nil {
		return m.PackageName
	}
	return ""
}

type InstallLog struct {
	// log
	Log string `protobuf:"bytes,50,opt,name=log" json:"log,omitempty"`
}

func (m *InstallLog) Reset()                    { *m = InstallLog{} }
func (m *InstallLog) String() string            { return proto.CompactTextString(m) }
func (*InstallLog) ProtoMessage()               {}
func (*InstallLog) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *InstallLog) GetLog() string {
	if m != nil {
		return m.Log
	}
	return ""
}

func init() {
	proto.RegisterType((*InstallLog_KEYS)(nil), "cisco_ios_xr_spirit_install_instmgr_oper.software_install.packages.package.verbose.install_log_KEYS")
	proto.RegisterType((*InstallLog)(nil), "cisco_ios_xr_spirit_install_instmgr_oper.software_install.packages.package.verbose.install_log")
}

func init() { proto.RegisterFile("install_log.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcc, 0xcc, 0x2b, 0x2e,
	0x49, 0xcc, 0xc9, 0x89, 0xcf, 0xc9, 0x4f, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x0a, 0x4a,
	0xce, 0x2c, 0x4e, 0xce, 0x8f, 0xcf, 0xcc, 0x2f, 0x8e, 0xaf, 0x28, 0x8a, 0x2f, 0x2e, 0xc8, 0x2c,
	0xca, 0x2c, 0x89, 0x87, 0x29, 0x03, 0xd1, 0xb9, 0xe9, 0x45, 0xf1, 0xf9, 0x05, 0xa9, 0x45, 0x7a,
	0xc5, 0xf9, 0x69, 0x25, 0xe5, 0x89, 0x45, 0xa9, 0x30, 0x59, 0xbd, 0x82, 0xc4, 0xe4, 0xec, 0xc4,
	0xf4, 0xd4, 0x62, 0x18, 0x43, 0xaf, 0x2c, 0xb5, 0x28, 0x29, 0xbf, 0x38, 0x55, 0xc9, 0x94, 0x4b,
	0x00, 0xc9, 0xa2, 0x78, 0x6f, 0xd7, 0xc8, 0x60, 0x21, 0x45, 0x2e, 0x1e, 0xa8, 0xb2, 0xf8, 0xbc,
	0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x6e, 0xa8, 0x98, 0x5f, 0x62, 0x6e,
	0xaa, 0x92, 0x3c, 0x17, 0x37, 0x92, 0x36, 0x21, 0x01, 0x2e, 0xe6, 0x9c, 0xfc, 0x74, 0x09, 0x23,
	0xb0, 0x42, 0x10, 0x33, 0x89, 0x0d, 0xec, 0x64, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8a,
	0x59, 0xe4, 0x2a, 0xc7, 0x00, 0x00, 0x00,
}
