// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ospf_sh_topology.proto

/*
Package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_areas_route_area_connected_route_areas_connected_route_area is a generated protocol buffer package.

It is generated from these files:
	ospf_sh_topology.proto

It has these top-level messages:
	OspfShTopology_KEYS
	OspfShTopology
	OspfShTime
	OspfShTopCommon
	OspfShTopPath
*/
package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_areas_route_area_connected_route_areas_connected_route_area

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

// OSPF Route Information
type OspfShTopology_KEYS struct {
	ProcessName  string `protobuf:"bytes,1,opt,name=process_name,json=processName" json:"process_name,omitempty"`
	AreaId       uint32 `protobuf:"varint,2,opt,name=area_id,json=areaId" json:"area_id,omitempty"`
	Prefix       string `protobuf:"bytes,3,opt,name=prefix" json:"prefix,omitempty"`
	PrefixLength uint32 `protobuf:"varint,4,opt,name=prefix_length,json=prefixLength" json:"prefix_length,omitempty"`
}

func (m *OspfShTopology_KEYS) Reset()                    { *m = OspfShTopology_KEYS{} }
func (m *OspfShTopology_KEYS) String() string            { return proto.CompactTextString(m) }
func (*OspfShTopology_KEYS) ProtoMessage()               {}
func (*OspfShTopology_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *OspfShTopology_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *OspfShTopology_KEYS) GetAreaId() uint32 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

func (m *OspfShTopology_KEYS) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *OspfShTopology_KEYS) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

type OspfShTopology struct {
	// Prefix
	RoutePrefix string `protobuf:"bytes,50,opt,name=route_prefix,json=routePrefix" json:"route_prefix,omitempty"`
	// Prefix length
	RoutePrefixLength uint32 `protobuf:"varint,51,opt,name=route_prefix_length,json=routePrefixLength" json:"route_prefix_length,omitempty"`
	// Metric
	RouteMetric uint32 `protobuf:"varint,52,opt,name=route_metric,json=routeMetric" json:"route_metric,omitempty"`
	// Route type
	RouteType string `protobuf:"bytes,53,opt,name=route_type,json=routeType" json:"route_type,omitempty"`
	// If true, connected route
	RouteConnected bool `protobuf:"varint,54,opt,name=route_connected,json=routeConnected" json:"route_connected,omitempty"`
	// Route information
	RouteInfo *OspfShTopCommon `protobuf:"bytes,55,opt,name=route_info,json=routeInfo" json:"route_info,omitempty"`
	// List of paths to this route
	RoutePathList []*OspfShTopPath `protobuf:"bytes,56,rep,name=route_path_list,json=routePathList" json:"route_path_list,omitempty"`
}

func (m *OspfShTopology) Reset()                    { *m = OspfShTopology{} }
func (m *OspfShTopology) String() string            { return proto.CompactTextString(m) }
func (*OspfShTopology) ProtoMessage()               {}
func (*OspfShTopology) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *OspfShTopology) GetRoutePrefix() string {
	if m != nil {
		return m.RoutePrefix
	}
	return ""
}

func (m *OspfShTopology) GetRoutePrefixLength() uint32 {
	if m != nil {
		return m.RoutePrefixLength
	}
	return 0
}

func (m *OspfShTopology) GetRouteMetric() uint32 {
	if m != nil {
		return m.RouteMetric
	}
	return 0
}

func (m *OspfShTopology) GetRouteType() string {
	if m != nil {
		return m.RouteType
	}
	return ""
}

func (m *OspfShTopology) GetRouteConnected() bool {
	if m != nil {
		return m.RouteConnected
	}
	return false
}

func (m *OspfShTopology) GetRouteInfo() *OspfShTopCommon {
	if m != nil {
		return m.RouteInfo
	}
	return nil
}

func (m *OspfShTopology) GetRoutePathList() []*OspfShTopPath {
	if m != nil {
		return m.RoutePathList
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

// OSPF Common Route Information
type OspfShTopCommon struct {
	// Area ID
	RouteAreaId uint32 `protobuf:"varint,1,opt,name=route_area_id,json=routeAreaId" json:"route_area_id,omitempty"`
	// TE metric
	RouteTeMetric uint32 `protobuf:"varint,2,opt,name=route_te_metric,json=routeTeMetric" json:"route_te_metric,omitempty"`
	// RIB version
	RouteRibVersion uint32 `protobuf:"varint,3,opt,name=route_rib_version,json=routeRibVersion" json:"route_rib_version,omitempty"`
	// SPF version
	RouteSpfVersion uint64 `protobuf:"varint,4,opt,name=route_spf_version,json=routeSpfVersion" json:"route_spf_version,omitempty"`
	// Forward distance
	RouteForwardDistance uint32 `protobuf:"varint,5,opt,name=route_forward_distance,json=routeForwardDistance" json:"route_forward_distance,omitempty"`
	// Protocol source
	RouteSource uint32 `protobuf:"varint,6,opt,name=route_source,json=routeSource" json:"route_source,omitempty"`
	// Last time updated
	RouteUpdateTime *OspfShTime `protobuf:"bytes,7,opt,name=route_update_time,json=routeUpdateTime" json:"route_update_time,omitempty"`
	// Last time update failed
	RouteFailTime *OspfShTime `protobuf:"bytes,8,opt,name=route_fail_time,json=routeFailTime" json:"route_fail_time,omitempty"`
	// SPF priority
	RouteSpfPriority uint32 `protobuf:"varint,9,opt,name=route_spf_priority,json=routeSpfPriority" json:"route_spf_priority,omitempty"`
	// If true, exclude from TE paths
	RouteAutoExcluded bool `protobuf:"varint,10,opt,name=route_auto_excluded,json=routeAutoExcluded" json:"route_auto_excluded,omitempty"`
	// If true, SRTE registered prefix route
	RouteSrtePrefixRegistered bool `protobuf:"varint,11,opt,name=route_srte_prefix_registered,json=routeSrtePrefixRegistered" json:"route_srte_prefix_registered,omitempty"`
	// SRTE registered neigbhor count on route
	RouteSrteNbrRegistered uint32 `protobuf:"varint,12,opt,name=route_srte_nbr_registered,json=routeSrteNbrRegistered" json:"route_srte_nbr_registered,omitempty"`
}

func (m *OspfShTopCommon) Reset()                    { *m = OspfShTopCommon{} }
func (m *OspfShTopCommon) String() string            { return proto.CompactTextString(m) }
func (*OspfShTopCommon) ProtoMessage()               {}
func (*OspfShTopCommon) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *OspfShTopCommon) GetRouteAreaId() uint32 {
	if m != nil {
		return m.RouteAreaId
	}
	return 0
}

func (m *OspfShTopCommon) GetRouteTeMetric() uint32 {
	if m != nil {
		return m.RouteTeMetric
	}
	return 0
}

func (m *OspfShTopCommon) GetRouteRibVersion() uint32 {
	if m != nil {
		return m.RouteRibVersion
	}
	return 0
}

func (m *OspfShTopCommon) GetRouteSpfVersion() uint64 {
	if m != nil {
		return m.RouteSpfVersion
	}
	return 0
}

func (m *OspfShTopCommon) GetRouteForwardDistance() uint32 {
	if m != nil {
		return m.RouteForwardDistance
	}
	return 0
}

func (m *OspfShTopCommon) GetRouteSource() uint32 {
	if m != nil {
		return m.RouteSource
	}
	return 0
}

func (m *OspfShTopCommon) GetRouteUpdateTime() *OspfShTime {
	if m != nil {
		return m.RouteUpdateTime
	}
	return nil
}

func (m *OspfShTopCommon) GetRouteFailTime() *OspfShTime {
	if m != nil {
		return m.RouteFailTime
	}
	return nil
}

func (m *OspfShTopCommon) GetRouteSpfPriority() uint32 {
	if m != nil {
		return m.RouteSpfPriority
	}
	return 0
}

func (m *OspfShTopCommon) GetRouteAutoExcluded() bool {
	if m != nil {
		return m.RouteAutoExcluded
	}
	return false
}

func (m *OspfShTopCommon) GetRouteSrtePrefixRegistered() bool {
	if m != nil {
		return m.RouteSrtePrefixRegistered
	}
	return false
}

func (m *OspfShTopCommon) GetRouteSrteNbrRegistered() uint32 {
	if m != nil {
		return m.RouteSrteNbrRegistered
	}
	return 0
}

// OSPF Route Path Information
type OspfShTopPath struct {
	// Next hop Interface
	RouteInterfaceName string `protobuf:"bytes,1,opt,name=route_interface_name,json=routeInterfaceName" json:"route_interface_name,omitempty"`
	// Nexthop IP address
	RouteNextHopAddress string `protobuf:"bytes,2,opt,name=route_next_hop_address,json=routeNextHopAddress" json:"route_next_hop_address,omitempty"`
	// IP address of source of route
	RouteSource string `protobuf:"bytes,3,opt,name=route_source,json=routeSource" json:"route_source,omitempty"`
	// LSA ID, see RFC2328
	RouteLsaid string `protobuf:"bytes,4,opt,name=route_lsaid,json=routeLsaid" json:"route_lsaid,omitempty"`
	// Multicast-intact path
	RoutePathIsMcastIntact bool `protobuf:"varint,5,opt,name=route_path_is_mcast_intact,json=routePathIsMcastIntact" json:"route_path_is_mcast_intact,omitempty"`
	// UCMP path
	RoutePathIsUcmpPath bool `protobuf:"varint,6,opt,name=route_path_is_ucmp_path,json=routePathIsUcmpPath" json:"route_path_is_ucmp_path,omitempty"`
	// Metric
	RouteMetric uint32 `protobuf:"varint,7,opt,name=route_metric,json=routeMetric" json:"route_metric,omitempty"`
	// LSA type, see RFC2328 etc.
	LsaType uint32 `protobuf:"varint,8,opt,name=lsa_type,json=lsaType" json:"lsa_type,omitempty"`
	// Area ID
	AreaId uint32 `protobuf:"varint,9,opt,name=area_id,json=areaId" json:"area_id,omitempty"`
	// Area format IP or uint32
	AreaFormat bool `protobuf:"varint,10,opt,name=area_format,json=areaFormat" json:"area_format,omitempty"`
}

func (m *OspfShTopPath) Reset()                    { *m = OspfShTopPath{} }
func (m *OspfShTopPath) String() string            { return proto.CompactTextString(m) }
func (*OspfShTopPath) ProtoMessage()               {}
func (*OspfShTopPath) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *OspfShTopPath) GetRouteInterfaceName() string {
	if m != nil {
		return m.RouteInterfaceName
	}
	return ""
}

func (m *OspfShTopPath) GetRouteNextHopAddress() string {
	if m != nil {
		return m.RouteNextHopAddress
	}
	return ""
}

func (m *OspfShTopPath) GetRouteSource() string {
	if m != nil {
		return m.RouteSource
	}
	return ""
}

func (m *OspfShTopPath) GetRouteLsaid() string {
	if m != nil {
		return m.RouteLsaid
	}
	return ""
}

func (m *OspfShTopPath) GetRoutePathIsMcastIntact() bool {
	if m != nil {
		return m.RoutePathIsMcastIntact
	}
	return false
}

func (m *OspfShTopPath) GetRoutePathIsUcmpPath() bool {
	if m != nil {
		return m.RoutePathIsUcmpPath
	}
	return false
}

func (m *OspfShTopPath) GetRouteMetric() uint32 {
	if m != nil {
		return m.RouteMetric
	}
	return 0
}

func (m *OspfShTopPath) GetLsaType() uint32 {
	if m != nil {
		return m.LsaType
	}
	return 0
}

func (m *OspfShTopPath) GetAreaId() uint32 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

func (m *OspfShTopPath) GetAreaFormat() bool {
	if m != nil {
		return m.AreaFormat
	}
	return false
}

func init() {
	proto.RegisterType((*OspfShTopology_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.route_areas.route_area.connected_route_areas.connected_route_area.ospf_sh_topology_KEYS")
	proto.RegisterType((*OspfShTopology)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.route_areas.route_area.connected_route_areas.connected_route_area.ospf_sh_topology")
	proto.RegisterType((*OspfShTime)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.route_areas.route_area.connected_route_areas.connected_route_area.ospf_sh_time")
	proto.RegisterType((*OspfShTopCommon)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.route_areas.route_area.connected_route_areas.connected_route_area.ospf_sh_top_common")
	proto.RegisterType((*OspfShTopPath)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.route_areas.route_area.connected_route_areas.connected_route_area.ospf_sh_top_path")
}

func init() { proto.RegisterFile("ospf_sh_topology.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 858 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0xdd, 0x6e, 0xe4, 0x34,
	0x14, 0x56, 0x68, 0xe9, 0xcc, 0x78, 0x3a, 0xb0, 0x98, 0xa5, 0x9b, 0x22, 0x60, 0x87, 0x41, 0x82,
	0x11, 0x42, 0x11, 0x6a, 0xcb, 0xef, 0x0d, 0xaa, 0x60, 0x2b, 0x2a, 0xba, 0xd5, 0x2a, 0xed, 0x22,
	0x71, 0x65, 0x79, 0x9c, 0x93, 0x8e, 0xa5, 0x24, 0x8e, 0x6c, 0xa7, 0xcc, 0x3c, 0x01, 0xb7, 0x88,
	0x27, 0xe0, 0x02, 0xc1, 0x9b, 0x20, 0x71, 0xc1, 0x83, 0xf0, 0x16, 0xc8, 0xc7, 0xce, 0x24, 0xdd,
	0xe1, 0x01, 0xca, 0x9d, 0xfd, 0x9d, 0x1f, 0x9f, 0x73, 0xf2, 0x9d, 0x73, 0x42, 0x0e, 0x94, 0xa9,
	0x73, 0x66, 0x96, 0xcc, 0xaa, 0x5a, 0x15, 0xea, 0x66, 0x9d, 0xd4, 0x5a, 0x59, 0x45, 0x7f, 0x89,
	0x84, 0x34, 0x42, 0x31, 0xa9, 0x0c, 0x5b, 0x69, 0x26, 0xeb, 0xdb, 0x13, 0x86, 0xaa, 0xaa, 0x06,
	0x9d, 0xb8, 0x93, 0x53, 0x14, 0x60, 0x0c, 0x98, 0xf6, 0x94, 0x64, 0x90, 0xf3, 0xa6, 0xb0, 0xec,
	0x56, 0xe7, 0x89, 0x56, 0x8d, 0x05, 0x26, 0xab, 0x5c, 0xe9, 0x92, 0x5b, 0xa9, 0xaa, 0x80, 0x70,
	0x0d, 0xdc, 0xf4, 0xce, 0x89, 0x50, 0x55, 0x05, 0xc2, 0x42, 0xc6, 0xfa, 0x0a, 0xff, 0x85, 0xce,
	0x7e, 0x8e, 0xc8, 0x1b, 0x2f, 0xc6, 0xcb, 0xbe, 0x7b, 0xf2, 0xc3, 0x15, 0x7d, 0x97, 0xec, 0x87,
	0x20, 0x58, 0xc5, 0x4b, 0x88, 0xa3, 0x69, 0x34, 0x1f, 0xa5, 0xe3, 0x80, 0x5d, 0xf2, 0x12, 0xe8,
	0x23, 0x32, 0x70, 0x4e, 0x98, 0xcc, 0xe2, 0x97, 0xa6, 0xd1, 0x7c, 0x92, 0xee, 0xb9, 0xeb, 0x79,
	0x46, 0x0f, 0xc8, 0x5e, 0xad, 0x21, 0x97, 0xab, 0x78, 0x07, 0xad, 0xc2, 0x8d, 0xbe, 0x47, 0x26,
	0xfe, 0xc4, 0x0a, 0xa8, 0x6e, 0xec, 0x32, 0xde, 0x45, 0xb3, 0x7d, 0x0f, 0x5e, 0x20, 0x36, 0xfb,
	0x67, 0x97, 0x3c, 0x78, 0x31, 0x24, 0x17, 0x8d, 0x8f, 0x3a, 0xf8, 0x3d, 0xf2, 0xd1, 0x20, 0xf6,
	0xcc, 0x3b, 0x4f, 0xc8, 0xeb, 0x7d, 0x95, 0xf6, 0x89, 0x63, 0x7c, 0xe2, 0xb5, 0x9e, 0xa6, 0x7f,
	0xa7, 0x73, 0x59, 0x82, 0xd5, 0x52, 0xc4, 0x27, 0xa8, 0xe8, 0x5d, 0x3e, 0x45, 0x88, 0xbe, 0x4d,
	0x88, 0x57, 0xb1, 0xeb, 0x1a, 0xe2, 0x4f, 0xf0, 0xcd, 0x11, 0x22, 0xd7, 0xeb, 0x1a, 0xe8, 0x07,
	0xe4, 0x55, 0x2f, 0xde, 0x94, 0x36, 0xfe, 0x74, 0x1a, 0xcd, 0x87, 0xe9, 0x2b, 0x08, 0x7f, 0xdd,
	0xa2, 0xf4, 0xcf, 0xa8, 0x75, 0xe4, 0xbe, 0x5e, 0xfc, 0xd9, 0x34, 0x9a, 0x8f, 0x8f, 0x7e, 0x8f,
	0x92, 0xfb, 0x47, 0x88, 0xa4, 0x57, 0x79, 0x26, 0x54, 0x59, 0xaa, 0x2a, 0x64, 0x7c, 0x5e, 0xe5,
	0x8a, 0xfe, 0x1d, 0xb5, 0x29, 0xd7, 0xdc, 0x2e, 0x59, 0x21, 0x8d, 0x8d, 0x3f, 0x9f, 0xee, 0xcc,
	0xc7, 0x47, 0xbf, 0xdd, 0xfb, 0x6c, 0x5c, 0xc4, 0xe9, 0xc4, 0xf3, 0x80, 0xdb, 0xe5, 0x85, 0x34,
	0x76, 0x76, 0x46, 0xf6, 0x37, 0x2a, 0xb2, 0x04, 0x47, 0x5c, 0x03, 0x42, 0x55, 0x19, 0xd2, 0x7d,
	0x92, 0x86, 0x1b, 0x7d, 0x87, 0x90, 0x8a, 0x57, 0x2a, 0xc8, 0x3c, 0xd9, 0x7b, 0xc8, 0xec, 0xa7,
	0x01, 0xa1, 0xdb, 0x95, 0xa3, 0x33, 0x32, 0xe9, 0x82, 0x71, 0x6d, 0x12, 0xf5, 0x38, 0x76, 0xea,
	0x7b, 0xe5, 0xfd, 0xb6, 0xa2, 0x1d, 0x13, 0xbd, 0x7f, 0x6f, 0x7a, 0xdd, 0x72, 0xf1, 0x43, 0xe2,
	0x39, 0xcc, 0xb4, 0x5c, 0xb0, 0x5b, 0xd0, 0x46, 0xaa, 0x0a, 0xdb, 0x6b, 0x92, 0x7a, 0x07, 0xa9,
	0x5c, 0x7c, 0xef, 0xe1, 0x4e, 0xd7, 0x85, 0xd4, 0xea, 0xba, 0x5e, 0xdb, 0x0d, 0xba, 0x57, 0x75,
	0xde, 0xea, 0x9e, 0x90, 0x03, 0xaf, 0x9b, 0x2b, 0xfd, 0x23, 0xd7, 0x19, 0xcb, 0xa4, 0xb1, 0xbc,
	0x12, 0x10, 0xbf, 0x8c, 0xce, 0x1f, 0xa2, 0xf4, 0xcc, 0x0b, 0xbf, 0x09, 0xb2, 0xae, 0x79, 0x8c,
	0x6a, 0xb4, 0x80, 0x78, 0xaf, 0x97, 0xd8, 0x15, 0x42, 0x8e, 0x2b, 0x21, 0x8a, 0xa6, 0xce, 0xb8,
	0x4b, 0x50, 0x96, 0x10, 0x0f, 0x90, 0xfb, 0xbf, 0xde, 0x6f, 0xb6, 0xc8, 0x12, 0x42, 0xa1, 0x9e,
	0x63, 0xe8, 0xd7, 0x8e, 0x1b, 0x7f, 0x6d, 0xb8, 0x9f, 0x73, 0x59, 0xf8, 0x6c, 0x86, 0xff, 0x97,
	0x6c, 0x3c, 0x99, 0xce, 0xb8, 0x2c, 0x30, 0x97, 0x8f, 0x08, 0xed, 0x08, 0x52, 0x6b, 0xa9, 0xb4,
	0xb4, 0xeb, 0x78, 0x84, 0x1f, 0xf1, 0x41, 0xcb, 0x90, 0x67, 0x01, 0xef, 0x26, 0x2b, 0x6f, 0xac,
	0x62, 0xb0, 0x12, 0x45, 0x93, 0x41, 0x16, 0x13, 0x9c, 0x75, 0xfe, 0x1b, 0x9f, 0x36, 0x56, 0x3d,
	0x09, 0x02, 0xfa, 0x15, 0x79, 0x2b, 0x78, 0xd7, 0xdd, 0x38, 0xd6, 0x70, 0x23, 0x8d, 0x05, 0x0d,
	0x59, 0x3c, 0x46, 0xc3, 0x43, 0xff, 0x8e, 0x6e, 0xc7, 0x72, 0xba, 0x51, 0xa0, 0x5f, 0x90, 0xc3,
	0x9e, 0x83, 0x6a, 0xa1, 0xfb, 0xd6, 0xfb, 0x18, 0xe5, 0xc1, 0xc6, 0xfa, 0x72, 0xa1, 0x3b, 0xd3,
	0xd9, 0x1f, 0x3b, 0x77, 0xb6, 0x07, 0x76, 0x3d, 0xfd, 0x98, 0x3c, 0x6c, 0x2b, 0x6c, 0x41, 0xe7,
	0x5c, 0x40, 0x7f, 0xa7, 0xd1, 0x30, 0xdf, 0x82, 0x08, 0x57, 0xdb, 0x71, 0xdb, 0x15, 0x15, 0xac,
	0x2c, 0x5b, 0xaa, 0x9a, 0xf1, 0x2c, 0xd3, 0x60, 0x0c, 0x36, 0xe7, 0x28, 0xf5, 0x05, 0xb9, 0x84,
	0x95, 0xfd, 0x56, 0xd5, 0xa7, 0x5e, 0xb4, 0xd5, 0x14, 0x3b, 0xbd, 0x25, 0x15, 0x9a, 0xe2, 0x31,
	0xf1, 0x57, 0x56, 0x18, 0x2e, 0x33, 0xec, 0xc9, 0x51, 0xea, 0x77, 0xc3, 0x85, 0x43, 0xe8, 0x97,
	0xe4, 0xcd, 0xde, 0x80, 0x95, 0x86, 0x95, 0x82, 0x1b, 0xeb, 0x02, 0xe7, 0xc2, 0x62, 0x4b, 0x0e,
	0x43, 0xee, 0x6e, 0x88, 0x9d, 0x9b, 0xa7, 0x4e, 0x7c, 0x8e, 0x52, 0x7a, 0x42, 0x1e, 0xdd, 0xb5,
	0x6d, 0x44, 0xe9, 0x2b, 0x80, 0xfd, 0x39, 0x0c, 0x51, 0x7b, 0xc3, 0xe7, 0xa2, 0xac, 0xdd, 0x69,
	0x6b, 0x0f, 0x0e, 0xb6, 0xf7, 0xe0, 0x21, 0x19, 0x16, 0x86, 0xfb, 0x2d, 0x38, 0x44, 0xf1, 0xa0,
	0x30, 0x1c, 0x77, 0x60, 0xef, 0x1f, 0x60, 0x74, 0xe7, 0x1f, 0xe0, 0x31, 0x19, 0xa3, 0xc0, 0x53,
	0x3a, 0x90, 0x85, 0x38, 0xe8, 0x0c, 0x91, 0xc5, 0x1e, 0xfe, 0x16, 0x1d, 0xff, 0x1b, 0x00, 0x00,
	0xff, 0xff, 0x3c, 0x74, 0xfc, 0x8a, 0x30, 0x09, 0x00, 0x00,
}
