// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ipv4_acl_edm_oor_detail.proto

/*
Package cisco_ios_xr_ipv4_acl_oper_ipv4_acl_and_prefix_list_oor_details is a generated protocol buffer package.

It is generated from these files:
	ipv4_acl_edm_oor_detail.proto

It has these top-level messages:
	Ipv4AclEdmOorDetail_KEYS
	Ipv4AclEdmOorDetail
*/
package cisco_ios_xr_ipv4_acl_oper_ipv4_acl_and_prefix_list_oor_details

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

// Oor deatil config BAG
type Ipv4AclEdmOorDetail_KEYS struct {
}

func (m *Ipv4AclEdmOorDetail_KEYS) Reset()                    { *m = Ipv4AclEdmOorDetail_KEYS{} }
func (m *Ipv4AclEdmOorDetail_KEYS) String() string            { return proto.CompactTextString(m) }
func (*Ipv4AclEdmOorDetail_KEYS) ProtoMessage()               {}
func (*Ipv4AclEdmOorDetail_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Ipv4AclEdmOorDetail struct {
	// default max configurable acls
	DefaultMaxAcLs uint32 `protobuf:"varint,50,opt,name=default_max_ac_ls,json=defaultMaxAcLs" json:"default_max_ac_ls,omitempty"`
	// default max configurable aces
	DefaultMaxAcEs uint32 `protobuf:"varint,51,opt,name=default_max_ac_es,json=defaultMaxAcEs" json:"default_max_ac_es,omitempty"`
	// Current configured acls
	CurrentConfiguredAcLs uint32 `protobuf:"varint,52,opt,name=current_configured_ac_ls,json=currentConfiguredAcLs" json:"current_configured_ac_ls,omitempty"`
	// Current configured aces
	CurrentConfiguredAcEs uint32 `protobuf:"varint,53,opt,name=current_configured_ac_es,json=currentConfiguredAcEs" json:"current_configured_ac_es,omitempty"`
	// Current max configurable acls
	CurrentMaxConfigurableAcLs uint32 `protobuf:"varint,54,opt,name=current_max_configurable_ac_ls,json=currentMaxConfigurableAcLs" json:"current_max_configurable_ac_ls,omitempty"`
	// Current max configurable aces
	CurrentMaxConfigurableAcEs uint32 `protobuf:"varint,55,opt,name=current_max_configurable_ac_es,json=currentMaxConfigurableAcEs" json:"current_max_configurable_ac_es,omitempty"`
	// max configurable acls
	MaxConfigurableAcLs uint32 `protobuf:"varint,56,opt,name=max_configurable_ac_ls,json=maxConfigurableAcLs" json:"max_configurable_ac_ls,omitempty"`
	// max configurable aces
	MaxConfigurableAcEs uint32 `protobuf:"varint,57,opt,name=max_configurable_ac_es,json=maxConfigurableAcEs" json:"max_configurable_ac_es,omitempty"`
}

func (m *Ipv4AclEdmOorDetail) Reset()                    { *m = Ipv4AclEdmOorDetail{} }
func (m *Ipv4AclEdmOorDetail) String() string            { return proto.CompactTextString(m) }
func (*Ipv4AclEdmOorDetail) ProtoMessage()               {}
func (*Ipv4AclEdmOorDetail) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Ipv4AclEdmOorDetail) GetDefaultMaxAcLs() uint32 {
	if m != nil {
		return m.DefaultMaxAcLs
	}
	return 0
}

func (m *Ipv4AclEdmOorDetail) GetDefaultMaxAcEs() uint32 {
	if m != nil {
		return m.DefaultMaxAcEs
	}
	return 0
}

func (m *Ipv4AclEdmOorDetail) GetCurrentConfiguredAcLs() uint32 {
	if m != nil {
		return m.CurrentConfiguredAcLs
	}
	return 0
}

func (m *Ipv4AclEdmOorDetail) GetCurrentConfiguredAcEs() uint32 {
	if m != nil {
		return m.CurrentConfiguredAcEs
	}
	return 0
}

func (m *Ipv4AclEdmOorDetail) GetCurrentMaxConfigurableAcLs() uint32 {
	if m != nil {
		return m.CurrentMaxConfigurableAcLs
	}
	return 0
}

func (m *Ipv4AclEdmOorDetail) GetCurrentMaxConfigurableAcEs() uint32 {
	if m != nil {
		return m.CurrentMaxConfigurableAcEs
	}
	return 0
}

func (m *Ipv4AclEdmOorDetail) GetMaxConfigurableAcLs() uint32 {
	if m != nil {
		return m.MaxConfigurableAcLs
	}
	return 0
}

func (m *Ipv4AclEdmOorDetail) GetMaxConfigurableAcEs() uint32 {
	if m != nil {
		return m.MaxConfigurableAcEs
	}
	return 0
}

func init() {
	proto.RegisterType((*Ipv4AclEdmOorDetail_KEYS)(nil), "cisco_ios_xr_ipv4_acl_oper.ipv4_acl_and_prefix_list.oor.details.ipv4_acl_edm_oor_detail_KEYS")
	proto.RegisterType((*Ipv4AclEdmOorDetail)(nil), "cisco_ios_xr_ipv4_acl_oper.ipv4_acl_and_prefix_list.oor.details.ipv4_acl_edm_oor_detail")
}

func init() { proto.RegisterFile("ipv4_acl_edm_oor_detail.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0x4b, 0xfc, 0x30,
	0x10, 0x47, 0xf9, 0xf3, 0x07, 0x0f, 0x01, 0x05, 0x2b, 0x6a, 0x10, 0x5d, 0x64, 0x4f, 0x7a, 0xe9,
	0xc1, 0xae, 0xae, 0x9e, 0x44, 0x97, 0x9e, 0x74, 0x2f, 0x7a, 0xf2, 0x34, 0x64, 0xd3, 0xa9, 0x04,
	0xd2, 0x4e, 0xc9, 0xb4, 0x92, 0xef, 0xe9, 0x17, 0x12, 0x63, 0xb6, 0x08, 0x6e, 0xf4, 0x58, 0xe6,
	0xbd, 0xdf, 0xeb, 0x21, 0xe2, 0xc4, 0x74, 0x6f, 0x33, 0x50, 0xda, 0x02, 0x56, 0x0d, 0x10, 0x39,
	0xa8, 0xb0, 0x57, 0xc6, 0xe6, 0x9d, 0xa3, 0x9e, 0xb2, 0x5b, 0x6d, 0x58, 0x13, 0x18, 0x62, 0xf0,
	0x0e, 0x46, 0x96, 0x3a, 0x74, 0xf9, 0xf8, 0xa5, 0xda, 0x0a, 0x3a, 0x87, 0xb5, 0xf1, 0x60, 0x0d,
	0xf7, 0x39, 0x91, 0xcb, 0xbf, 0x56, 0x78, 0x3a, 0x11, 0xc7, 0x89, 0x02, 0x3c, 0x94, 0x2f, 0xcf,
	0xd3, 0xf7, 0xff, 0xe2, 0x30, 0x01, 0x64, 0xe7, 0x62, 0xb7, 0xc2, 0x5a, 0x0d, 0xb6, 0x87, 0x46,
	0x79, 0x50, 0x1a, 0x2c, 0xcb, 0x8b, 0xd3, 0x7f, 0x67, 0xdb, 0x4f, 0x3b, 0xf1, 0xb0, 0x54, 0xfe,
	0x4e, 0x3f, 0xf2, 0x06, 0x14, 0x59, 0x16, 0x3f, 0xd1, 0x92, 0xb3, 0xb9, 0x90, 0x7a, 0x70, 0x0e,
	0xdb, 0x1e, 0x34, 0xb5, 0xb5, 0x79, 0x1d, 0x1c, 0x56, 0x71, 0x7c, 0x16, 0x8c, 0xfd, 0x78, 0x5f,
	0x8c, 0xe7, 0xd0, 0x48, 0x8a, 0xc8, 0xf2, 0x32, 0x29, 0x96, 0x9c, 0xdd, 0x8b, 0xc9, 0x5a, 0xfc,
	0xfc, 0xb9, 0xb5, 0xac, 0x56, 0x16, 0x63, 0xf7, 0x2a, 0xe8, 0x47, 0x91, 0x5a, 0x2a, 0xbf, 0xf8,
	0xc6, 0x84, 0xf8, 0x1f, 0x1b, 0xc8, 0x72, 0xfe, 0xfb, 0x46, 0xc9, 0x59, 0x21, 0x0e, 0x12, 0xfd,
	0xeb, 0xe0, 0xee, 0x35, 0x1b, 0xc2, 0x09, 0x09, 0x59, 0xde, 0x24, 0xa4, 0x92, 0x57, 0x5b, 0xe1,
	0xf5, 0x14, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x59, 0xed, 0x44, 0x8c, 0x5e, 0x02, 0x00, 0x00,
}
