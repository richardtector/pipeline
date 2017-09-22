// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vrrp_mgo_info.proto

/*
Package cisco_ios_xr_ipv4_vrrp_oper_vrrp_mgo_sessions_mgo_session is a generated protocol buffer package.

It is generated from these files:
	vrrp_mgo_info.proto

It has these top-level messages:
	VrrpMgoInfo_KEYS
	VrrpMgoInfo
	VrrpSlaveInfoType
*/
package cisco_ios_xr_ipv4_vrrp_oper_vrrp_mgo_sessions_mgo_session

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

// MGO group data
type VrrpMgoInfo_KEYS struct {
	SessionName string `protobuf:"bytes,1,opt,name=session_name,json=sessionName" json:"session_name,omitempty"`
}

func (m *VrrpMgoInfo_KEYS) Reset()                    { *m = VrrpMgoInfo_KEYS{} }
func (m *VrrpMgoInfo_KEYS) String() string            { return proto.CompactTextString(m) }
func (*VrrpMgoInfo_KEYS) ProtoMessage()               {}
func (*VrrpMgoInfo_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *VrrpMgoInfo_KEYS) GetSessionName() string {
	if m != nil {
		return m.SessionName
	}
	return ""
}

type VrrpMgoInfo struct {
	// Session Name
	PrimarySessionName string `protobuf:"bytes,50,opt,name=primary_session_name,json=primarySessionName" json:"primary_session_name,omitempty"`
	// Interface of primary session
	PrimarySessionInterface string `protobuf:"bytes,51,opt,name=primary_session_interface,json=primarySessionInterface" json:"primary_session_interface,omitempty"`
	// Address family of primary session
	PrimaryAfName string `protobuf:"bytes,52,opt,name=primary_af_name,json=primaryAfName" json:"primary_af_name,omitempty"`
	// VRID of primary session
	PrimarySessionNumber uint32 `protobuf:"varint,53,opt,name=primary_session_number,json=primarySessionNumber" json:"primary_session_number,omitempty"`
	// State of primary session
	PrimarySessionState string `protobuf:"bytes,54,opt,name=primary_session_state,json=primarySessionState" json:"primary_session_state,omitempty"`
	// List of slaves following this primary session
	SlaveList []*VrrpSlaveInfoType `protobuf:"bytes,55,rep,name=slave_list,json=slaveList" json:"slave_list,omitempty"`
}

func (m *VrrpMgoInfo) Reset()                    { *m = VrrpMgoInfo{} }
func (m *VrrpMgoInfo) String() string            { return proto.CompactTextString(m) }
func (*VrrpMgoInfo) ProtoMessage()               {}
func (*VrrpMgoInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *VrrpMgoInfo) GetPrimarySessionName() string {
	if m != nil {
		return m.PrimarySessionName
	}
	return ""
}

func (m *VrrpMgoInfo) GetPrimarySessionInterface() string {
	if m != nil {
		return m.PrimarySessionInterface
	}
	return ""
}

func (m *VrrpMgoInfo) GetPrimaryAfName() string {
	if m != nil {
		return m.PrimaryAfName
	}
	return ""
}

func (m *VrrpMgoInfo) GetPrimarySessionNumber() uint32 {
	if m != nil {
		return m.PrimarySessionNumber
	}
	return 0
}

func (m *VrrpMgoInfo) GetPrimarySessionState() string {
	if m != nil {
		return m.PrimarySessionState
	}
	return ""
}

func (m *VrrpMgoInfo) GetSlaveList() []*VrrpSlaveInfoType {
	if m != nil {
		return m.SlaveList
	}
	return nil
}

// Slave info
type VrrpSlaveInfoType struct {
	// Interface of slave
	SlaveInterface string `protobuf:"bytes,1,opt,name=slave_interface,json=slaveInterface" json:"slave_interface,omitempty"`
	// VRID of slave
	SlaveVirtualRouterId uint32 `protobuf:"varint,2,opt,name=slave_virtual_router_id,json=slaveVirtualRouterId" json:"slave_virtual_router_id,omitempty"`
}

func (m *VrrpSlaveInfoType) Reset()                    { *m = VrrpSlaveInfoType{} }
func (m *VrrpSlaveInfoType) String() string            { return proto.CompactTextString(m) }
func (*VrrpSlaveInfoType) ProtoMessage()               {}
func (*VrrpSlaveInfoType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *VrrpSlaveInfoType) GetSlaveInterface() string {
	if m != nil {
		return m.SlaveInterface
	}
	return ""
}

func (m *VrrpSlaveInfoType) GetSlaveVirtualRouterId() uint32 {
	if m != nil {
		return m.SlaveVirtualRouterId
	}
	return 0
}

func init() {
	proto.RegisterType((*VrrpMgoInfo_KEYS)(nil), "cisco_ios_xr_ipv4_vrrp_oper.vrrp.mgo_sessions.mgo_session.vrrp_mgo_info_KEYS")
	proto.RegisterType((*VrrpMgoInfo)(nil), "cisco_ios_xr_ipv4_vrrp_oper.vrrp.mgo_sessions.mgo_session.vrrp_mgo_info")
	proto.RegisterType((*VrrpSlaveInfoType)(nil), "cisco_ios_xr_ipv4_vrrp_oper.vrrp.mgo_sessions.mgo_session.vrrp_slave_info_type")
}

func init() { proto.RegisterFile("vrrp_mgo_info.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x69, 0x05, 0xa1, 0x53, 0x6b, 0x61, 0x5b, 0x6d, 0xbc, 0xd5, 0x1e, 0xb4, 0xa7, 0x20,
	0xfd, 0x63, 0xd1, 0x9b, 0x07, 0x0f, 0x45, 0x51, 0x48, 0x41, 0xf0, 0x34, 0x6c, 0xdb, 0x8d, 0x2c,
	0x34, 0xd9, 0x30, 0xbb, 0x0d, 0xf6, 0xd3, 0x2b, 0x9d, 0x24, 0xd4, 0x44, 0x4f, 0xde, 0x92, 0xf7,
	0xde, 0x6f, 0x66, 0x98, 0x59, 0xe8, 0xa4, 0x44, 0x09, 0x46, 0x1f, 0x06, 0x75, 0x1c, 0x1a, 0x3f,
	0x21, 0xe3, 0x8c, 0xb8, 0x5b, 0x69, 0xbb, 0x32, 0xa8, 0x8d, 0xc5, 0x4f, 0x42, 0x9d, 0xa4, 0x13,
	0xe4, 0x98, 0x49, 0x14, 0xf9, 0xfb, 0x2f, 0x7f, 0x0f, 0x58, 0x65, 0xad, 0x36, 0xb1, 0xfd, 0xf9,
	0x33, 0x98, 0x81, 0x28, 0x55, 0xc4, 0xa7, 0xc7, 0xf7, 0x85, 0xb8, 0x84, 0x93, 0x3c, 0x80, 0xb1,
	0x8c, 0x94, 0x57, 0xeb, 0xd7, 0x86, 0x8d, 0xa0, 0x99, 0x6b, 0x2f, 0x32, 0x52, 0x83, 0xaf, 0x3a,
	0xb4, 0x4a, 0xa4, 0xb8, 0x81, 0x6e, 0x42, 0x3a, 0x92, 0xb4, 0xc3, 0x12, 0x3c, 0x62, 0x58, 0xe4,
	0xde, 0xe2, 0x50, 0x43, 0xdc, 0xc3, 0x45, 0x95, 0xd0, 0xb1, 0x53, 0x14, 0xca, 0x95, 0xf2, 0xc6,
	0x8c, 0xf5, 0xca, 0xd8, 0xbc, 0xb0, 0xc5, 0x15, 0xb4, 0x0b, 0x56, 0x86, 0x59, 0xa3, 0x09, 0x13,
	0xad, 0x5c, 0x7e, 0x08, 0xb9, 0xc7, 0x04, 0xce, 0x7f, 0x4d, 0xb5, 0x8d, 0x96, 0x8a, 0xbc, 0x69,
	0xbf, 0x36, 0x6c, 0x05, 0xdd, 0xca, 0x5c, 0xec, 0x89, 0x11, 0x9c, 0x55, 0x29, 0xeb, 0xa4, 0x53,
	0xde, 0x2d, 0xf7, 0xe8, 0x94, 0xa1, 0xc5, 0xde, 0x12, 0x31, 0x80, 0xdd, 0xc8, 0x54, 0xe1, 0x46,
	0x5b, 0xe7, 0xcd, 0xfa, 0x47, 0xc3, 0xe6, 0xe8, 0xd5, 0xff, 0xf7, 0x69, 0xd8, 0xc5, 0xac, 0x22,
	0x5f, 0xc6, 0xed, 0x12, 0x15, 0x34, 0x58, 0x78, 0xd6, 0xd6, 0x0d, 0x52, 0xe8, 0xfe, 0x15, 0x11,
	0xd7, 0xd0, 0x2e, 0xa4, 0x62, 0x97, 0xd9, 0xfd, 0x4e, 0x59, 0x3e, 0xac, 0x70, 0x0a, 0xbd, 0x2c,
	0x98, 0x6a, 0x72, 0x5b, 0xb9, 0x41, 0x32, 0x5b, 0xa7, 0x08, 0xf5, 0xda, 0xab, 0x67, 0xbb, 0x61,
	0xfb, 0x2d, 0x73, 0x03, 0x36, 0xe7, 0xeb, 0xe5, 0x31, 0x3f, 0xba, 0xf1, 0x77, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x44, 0x77, 0x92, 0xb4, 0x8b, 0x02, 0x00, 0x00,
}
