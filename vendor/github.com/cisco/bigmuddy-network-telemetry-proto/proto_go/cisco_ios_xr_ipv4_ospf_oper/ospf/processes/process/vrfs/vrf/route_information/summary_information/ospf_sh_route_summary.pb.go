// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ospf_sh_route_summary.proto

/*
Package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_vrfs_vrf_route_information_summary_information is a generated protocol buffer package.

It is generated from these files:
	ospf_sh_route_summary.proto

It has these top-level messages:
	OspfShRouteSummary_KEYS
	OspfShRouteSummary
	OspfShTime
	OspfShRouteSummCommon
*/
package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_vrfs_vrf_route_information_summary_information

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

// OSPF Route Summary Information
type OspfShRouteSummary_KEYS struct {
	ProcessName string `protobuf:"bytes,1,opt,name=process_name,json=processName" json:"process_name,omitempty"`
	VrfName     string `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
}

func (m *OspfShRouteSummary_KEYS) Reset()                    { *m = OspfShRouteSummary_KEYS{} }
func (m *OspfShRouteSummary_KEYS) String() string            { return proto.CompactTextString(m) }
func (*OspfShRouteSummary_KEYS) ProtoMessage()               {}
func (*OspfShRouteSummary_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *OspfShRouteSummary_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *OspfShRouteSummary_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

type OspfShRouteSummary struct {
	// RIB failure count
	Failures uint32 `protobuf:"varint,50,opt,name=failures" json:"failures,omitempty"`
	// Last time failed
	FailureTime *OspfShTime `protobuf:"bytes,51,opt,name=failure_time,json=failureTime" json:"failure_time,omitempty"`
	// Last failed address
	FailureAddress string `protobuf:"bytes,52,opt,name=failure_address,json=failureAddress" json:"failure_address,omitempty"`
	// OSPF route summary Information
	Common *OspfShRouteSummCommon `protobuf:"bytes,53,opt,name=common" json:"common,omitempty"`
}

func (m *OspfShRouteSummary) Reset()                    { *m = OspfShRouteSummary{} }
func (m *OspfShRouteSummary) String() string            { return proto.CompactTextString(m) }
func (*OspfShRouteSummary) ProtoMessage()               {}
func (*OspfShRouteSummary) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *OspfShRouteSummary) GetFailures() uint32 {
	if m != nil {
		return m.Failures
	}
	return 0
}

func (m *OspfShRouteSummary) GetFailureTime() *OspfShTime {
	if m != nil {
		return m.FailureTime
	}
	return nil
}

func (m *OspfShRouteSummary) GetFailureAddress() string {
	if m != nil {
		return m.FailureAddress
	}
	return ""
}

func (m *OspfShRouteSummary) GetCommon() *OspfShRouteSummCommon {
	if m != nil {
		return m.Common
	}
	return nil
}

type OspfShTime struct {
	Second     uint32 `protobuf:"varint,1,opt,name=second" json:"second,omitempty"`
	Nanosecond uint32 `protobuf:"varint,2,opt,name=nanosecond" json:"nanosecond,omitempty"`
}

func (m *OspfShTime) Reset()                    { *m = OspfShTime{} }
func (m *OspfShTime) String() string            { return proto.CompactTextString(m) }
func (*OspfShTime) ProtoMessage()               {}
func (*OspfShTime) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *OspfShTime) GetSecond() uint32 {
	if m != nil {
		return m.Second
	}
	return 0
}

func (m *OspfShTime) GetNanosecond() uint32 {
	if m != nil {
		return m.Nanosecond
	}
	return 0
}

// OSPF Route Summary Common Information
type OspfShRouteSummCommon struct {
	// External type 1 count
	ExternalType1S uint32 `protobuf:"varint,1,opt,name=external_type1s,json=externalType1s" json:"external_type1s,omitempty"`
	// External type 2 count
	ExternalType2S uint32 `protobuf:"varint,2,opt,name=external_type2s,json=externalType2s" json:"external_type2s,omitempty"`
	// External NSSA type 1 count
	ExternalNssaType1S uint32 `protobuf:"varint,3,opt,name=external_nssa_type1s,json=externalNssaType1s" json:"external_nssa_type1s,omitempty"`
	// External NSSA type 2 count
	ExternalNssaType2S uint32 `protobuf:"varint,4,opt,name=external_nssa_type2s,json=externalNssaType2s" json:"external_nssa_type2s,omitempty"`
	// Inter-area count
	InterAreas uint32 `protobuf:"varint,5,opt,name=inter_areas,json=interAreas" json:"inter_areas,omitempty"`
	// Intra-area count
	IntraAreas uint32 `protobuf:"varint,6,opt,name=intra_areas,json=intraAreas" json:"intra_areas,omitempty"`
	// Total count
	Total uint32 `protobuf:"varint,7,opt,name=total" json:"total,omitempty"`
}

func (m *OspfShRouteSummCommon) Reset()                    { *m = OspfShRouteSummCommon{} }
func (m *OspfShRouteSummCommon) String() string            { return proto.CompactTextString(m) }
func (*OspfShRouteSummCommon) ProtoMessage()               {}
func (*OspfShRouteSummCommon) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *OspfShRouteSummCommon) GetExternalType1S() uint32 {
	if m != nil {
		return m.ExternalType1S
	}
	return 0
}

func (m *OspfShRouteSummCommon) GetExternalType2S() uint32 {
	if m != nil {
		return m.ExternalType2S
	}
	return 0
}

func (m *OspfShRouteSummCommon) GetExternalNssaType1S() uint32 {
	if m != nil {
		return m.ExternalNssaType1S
	}
	return 0
}

func (m *OspfShRouteSummCommon) GetExternalNssaType2S() uint32 {
	if m != nil {
		return m.ExternalNssaType2S
	}
	return 0
}

func (m *OspfShRouteSummCommon) GetInterAreas() uint32 {
	if m != nil {
		return m.InterAreas
	}
	return 0
}

func (m *OspfShRouteSummCommon) GetIntraAreas() uint32 {
	if m != nil {
		return m.IntraAreas
	}
	return 0
}

func (m *OspfShRouteSummCommon) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func init() {
	proto.RegisterType((*OspfShRouteSummary_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.route_information.summary_information.ospf_sh_route_summary_KEYS")
	proto.RegisterType((*OspfShRouteSummary)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.route_information.summary_information.ospf_sh_route_summary")
	proto.RegisterType((*OspfShTime)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.route_information.summary_information.ospf_sh_time")
	proto.RegisterType((*OspfShRouteSummCommon)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.route_information.summary_information.ospf_sh_route_summ_common")
}

func init() { proto.RegisterFile("ospf_sh_route_summary.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0xd5, 0xa6, 0x34, 0x2d, 0x93, 0xb4, 0x48, 0x56, 0x41, 0x6e, 0x91, 0x20, 0xe4, 0x42, 0x4e,
	0x2b, 0xd8, 0x96, 0x0f, 0xe8, 0x01, 0x2e, 0x48, 0x3d, 0x2c, 0xbd, 0xc0, 0xc5, 0x32, 0x9b, 0x59,
	0x61, 0x29, 0xb6, 0x57, 0x33, 0x4e, 0xd4, 0x48, 0x7c, 0x00, 0x67, 0xc4, 0x4f, 0xf0, 0x97, 0x68,
	0xbd, 0xde, 0x10, 0xc4, 0x72, 0xa4, 0x97, 0x95, 0xe7, 0xbd, 0xe7, 0x37, 0x6f, 0x66, 0x65, 0x78,
	0xea, 0xb9, 0xa9, 0x15, 0x7f, 0x51, 0xe4, 0xd7, 0x01, 0x15, 0xaf, 0xad, 0xd5, 0xb4, 0xcd, 0x1b,
	0xf2, 0xc1, 0x0b, 0x5d, 0x19, 0xae, 0xbc, 0x32, 0x9e, 0xd5, 0x1d, 0x29, 0xd3, 0x6c, 0xae, 0x54,
	0x94, 0xfb, 0x06, 0x29, 0x6f, 0x4f, 0xad, 0xae, 0x42, 0x66, 0xe4, 0xfe, 0x94, 0x6f, 0xa8, 0x8e,
	0x9f, 0xbc, 0x33, 0x34, 0xae, 0xf6, 0x64, 0x75, 0x30, 0xde, 0xe5, 0xc9, 0x7c, 0x1f, 0x9b, 0x7f,
	0x82, 0x8b, 0xc1, 0x04, 0xea, 0xfd, 0xdb, 0x8f, 0x1f, 0xc4, 0x0b, 0x98, 0x26, 0x5f, 0xe5, 0xb4,
	0x45, 0x99, 0xcd, 0xb2, 0xc5, 0xc3, 0x72, 0x92, 0xb0, 0x1b, 0x6d, 0x51, 0x9c, 0xc3, 0xf1, 0x86,
	0xea, 0x8e, 0x1e, 0x45, 0xfa, 0x68, 0x43, 0x75, 0x4b, 0xcd, 0xbf, 0x1d, 0xc0, 0xe3, 0x41, 0x73,
	0x71, 0x01, 0xc7, 0xb5, 0x36, 0xab, 0x35, 0x21, 0xcb, 0x62, 0x96, 0x2d, 0x4e, 0xca, 0x5d, 0x2d,
	0xbe, 0x67, 0x30, 0x4d, 0x85, 0x0a, 0xc6, 0xa2, 0xbc, 0x9c, 0x65, 0x8b, 0x49, 0xe1, 0xf3, 0xff,
	0xbe, 0x8c, 0xbc, 0x0f, 0xdb, 0xb6, 0x2d, 0x27, 0x29, 0xc4, 0xad, 0xb1, 0x28, 0x5e, 0xc2, 0xa3,
	0x3e, 0x93, 0x5e, 0x2e, 0x09, 0x99, 0xe5, 0x55, 0x1c, 0xf6, 0x34, 0xc1, 0xd7, 0x1d, 0x2a, 0x7e,
	0x64, 0x30, 0xae, 0xbc, 0xb5, 0xde, 0xc9, 0x37, 0x31, 0xf7, 0xd7, 0x7b, 0xcc, 0xfd, 0x7b, 0xc9,
	0xaa, 0xcb, 0x50, 0xa6, 0x2c, 0xf3, 0x77, 0x30, 0xdd, 0x1f, 0x4e, 0x3c, 0x81, 0x31, 0x63, 0xe5,
	0xdd, 0x32, 0xfe, 0xd2, 0x93, 0x32, 0x55, 0xe2, 0x19, 0x80, 0xd3, 0xce, 0x27, 0x6e, 0x14, 0xb9,
	0x3d, 0x64, 0xfe, 0x73, 0x04, 0xe7, 0xff, 0xec, 0xd6, 0x6e, 0x09, 0xef, 0x02, 0x92, 0xd3, 0x2b,
	0x15, 0xb6, 0x0d, 0xbe, 0xe6, 0x64, 0x7f, 0xda, 0xc3, 0xb7, 0x11, 0xfd, 0x4b, 0x58, 0x70, 0xea,
	0xf5, 0x87, 0xb0, 0x60, 0xf1, 0x0a, 0xce, 0x76, 0x42, 0xc7, 0xac, 0x7b, 0xdb, 0x83, 0xa8, 0x16,
	0x3d, 0x77, 0xc3, 0xac, 0x93, 0xf5, 0xe0, 0x8d, 0x82, 0xe5, 0x83, 0xe1, 0x1b, 0x05, 0x8b, 0xe7,
	0x30, 0x31, 0x2e, 0x20, 0x29, 0x4d, 0xa8, 0x59, 0x1e, 0x76, 0x43, 0x47, 0xe8, 0xba, 0x45, 0x92,
	0x80, 0x74, 0x12, 0x8c, 0x77, 0x02, 0xd2, 0x9d, 0xe0, 0x0c, 0x0e, 0x83, 0x0f, 0x7a, 0x25, 0x8f,
	0x22, 0xd5, 0x15, 0x9f, 0xc7, 0xf1, 0x11, 0x5f, 0xfe, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x9c, 0xc4,
	0x8c, 0xff, 0xe3, 0x03, 0x00, 0x00,
}
