// Code generated by protoc-gen-go. DO NOT EDIT.
// source: snmp_q_stats_b.proto

/*
Package cisco_ios_xr_snmp_agent_oper_snmp_information_request_type_detail_nms_addresses_nms_address is a generated protocol buffer package.

It is generated from these files:
	snmp_q_stats_b.proto

It has these top-level messages:
	SnmpQStatsB_KEYS
	SnmpQStatsB
*/
package cisco_ios_xr_snmp_agent_oper_snmp_information_request_type_detail_nms_addresses_nms_address

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

// SNMP incomming/processing queue statistics
type SnmpQStatsB_KEYS struct {
	NmsAddr string `protobuf:"bytes,1,opt,name=nms_addr,json=nmsAddr" json:"nms_addr,omitempty"`
}

func (m *SnmpQStatsB_KEYS) Reset()                    { *m = SnmpQStatsB_KEYS{} }
func (m *SnmpQStatsB_KEYS) String() string            { return proto.CompactTextString(m) }
func (*SnmpQStatsB_KEYS) ProtoMessage()               {}
func (*SnmpQStatsB_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SnmpQStatsB_KEYS) GetNmsAddr() string {
	if m != nil {
		return m.NmsAddr
	}
	return ""
}

type SnmpQStatsB struct {
	// Total request count for each managment station or client
	TotalCount uint32 `protobuf:"varint,50,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
	// Processing agent request count for each client for particluar managment station
	AgentRequestCount uint32 `protobuf:"varint,51,opt,name=agent_request_count,json=agentRequestCount" json:"agent_request_count,omitempty"`
	// Processing interfce request count for each client for particluar managment station
	InterfaceRequestCount uint32 `protobuf:"varint,52,opt,name=interface_request_count,json=interfaceRequestCount" json:"interface_request_count,omitempty"`
	// Processing entity request count for each client for particluar managment station
	EntityRequestCount uint32 `protobuf:"varint,53,opt,name=entity_request_count,json=entityRequestCount" json:"entity_request_count,omitempty"`
	// Processing route request count for each client for particluar Managment station
	RouteRequestCount uint32 `protobuf:"varint,54,opt,name=route_request_count,json=routeRequestCount" json:"route_request_count,omitempty"`
	// Processing infra request count for each client for particluar Managment station
	InfraRequestCount uint32 `protobuf:"varint,55,opt,name=infra_request_count,json=infraRequestCount" json:"infra_request_count,omitempty"`
}

func (m *SnmpQStatsB) Reset()                    { *m = SnmpQStatsB{} }
func (m *SnmpQStatsB) String() string            { return proto.CompactTextString(m) }
func (*SnmpQStatsB) ProtoMessage()               {}
func (*SnmpQStatsB) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SnmpQStatsB) GetTotalCount() uint32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *SnmpQStatsB) GetAgentRequestCount() uint32 {
	if m != nil {
		return m.AgentRequestCount
	}
	return 0
}

func (m *SnmpQStatsB) GetInterfaceRequestCount() uint32 {
	if m != nil {
		return m.InterfaceRequestCount
	}
	return 0
}

func (m *SnmpQStatsB) GetEntityRequestCount() uint32 {
	if m != nil {
		return m.EntityRequestCount
	}
	return 0
}

func (m *SnmpQStatsB) GetRouteRequestCount() uint32 {
	if m != nil {
		return m.RouteRequestCount
	}
	return 0
}

func (m *SnmpQStatsB) GetInfraRequestCount() uint32 {
	if m != nil {
		return m.InfraRequestCount
	}
	return 0
}

func init() {
	proto.RegisterType((*SnmpQStatsB_KEYS)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.information.request_type_detail.nms_addresses.nms_address.snmp_q_stats_b_KEYS")
	proto.RegisterType((*SnmpQStatsB)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.information.request_type_detail.nms_addresses.nms_address.snmp_q_stats_b")
}

func init() { proto.RegisterFile("snmp_q_stats_b.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x40, 0x89, 0x07, 0x3f, 0x56, 0x14, 0x4c, 0x2a, 0xd6, 0x93, 0xa5, 0xa7, 0x9e, 0x96, 0x62,
	0xb5, 0x9e, 0x45, 0x3c, 0x79, 0x8b, 0x27, 0xf1, 0x30, 0x6c, 0x93, 0x89, 0x2c, 0x34, 0xbb, 0xe9,
	0xcc, 0x04, 0xec, 0x6f, 0xf2, 0x4f, 0x4a, 0x77, 0x69, 0x69, 0xf6, 0xb8, 0xf3, 0xde, 0x63, 0x06,
	0x56, 0x8d, 0xd8, 0xb5, 0x1d, 0x6c, 0x80, 0xc5, 0x08, 0xc3, 0x4a, 0x77, 0xe4, 0xc5, 0xe7, 0xdf,
	0x95, 0xe5, 0xca, 0x83, 0xf5, 0x0c, 0xbf, 0x04, 0x41, 0x31, 0x3f, 0xe8, 0x04, 0x7c, 0x87, 0xa4,
	0x77, 0x6f, 0x6d, 0x5d, 0xe3, 0xa9, 0x35, 0x62, 0xbd, 0xd3, 0x84, 0x9b, 0x1e, 0x59, 0x40, 0xb6,
	0x1d, 0x42, 0x8d, 0x62, 0xec, 0x5a, 0xbb, 0x96, 0xc1, 0xd4, 0x35, 0x21, 0x33, 0xf2, 0xf1, 0x6b,
	0x3a, 0x57, 0xc5, 0x70, 0x29, 0x7c, 0xbc, 0x7f, 0x7d, 0xe6, 0xf7, 0xea, 0x7c, 0x6f, 0x8d, 0xb3,
	0x49, 0x36, 0xbb, 0x28, 0xcf, 0x5c, 0xcb, 0xaf, 0x75, 0x4d, 0xd3, 0xbf, 0x13, 0x75, 0x3d, 0x4c,
	0xf2, 0x07, 0x75, 0x29, 0x5e, 0xcc, 0x1a, 0x2a, 0xdf, 0x3b, 0x19, 0x3f, 0x4e, 0xb2, 0xd9, 0x55,
	0xa9, 0xc2, 0xe8, 0x6d, 0x37, 0xc9, 0xb5, 0x2a, 0xe2, 0xc9, 0xfb, 0xe3, 0xa2, 0xb8, 0x08, 0xe2,
	0x4d, 0x40, 0x65, 0x24, 0xd1, 0x5f, 0xaa, 0x3b, 0xeb, 0x04, 0xa9, 0x31, 0x15, 0x26, 0xcd, 0x53,
	0x68, 0x6e, 0x0f, 0x78, 0xd0, 0xcd, 0xd5, 0x08, 0x9d, 0x58, 0xd9, 0x26, 0xd1, 0x73, 0x88, 0xf2,
	0xc8, 0x06, 0x85, 0x56, 0x05, 0xf9, 0x5e, 0xd2, 0x2d, 0xcb, 0x78, 0x59, 0x40, 0xa9, 0x6f, 0x5d,
	0x43, 0x26, 0xf1, 0x5f, 0xa2, 0x1f, 0xd0, 0xb1, 0xbf, 0x3a, 0x0d, 0x7f, 0xb8, 0xf8, 0x0f, 0x00,
	0x00, 0xff, 0xff, 0x51, 0x0b, 0xd5, 0xa3, 0xdb, 0x01, 0x00, 0x00,
}
