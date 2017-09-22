// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ospf_sh_topology.proto

/*
Package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_vrfs_vrf_route_information_route_table_route is a generated protocol buffer package.

It is generated from these files:
	ospf_sh_topology.proto

It has these top-level messages:
	OspfShTopology_KEYS
	OspfShTopology
	OspfShTime
	OspfShTopCommon
	OspfShTopPath
*/
package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_vrfs_vrf_route_information_route_table_route

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
	VrfName      string `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
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

func (m *OspfShTopology_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
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
	proto.RegisterType((*OspfShTopology_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.route_information.route_table.route.ospf_sh_topology_KEYS")
	proto.RegisterType((*OspfShTopology)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.route_information.route_table.route.ospf_sh_topology")
	proto.RegisterType((*OspfShTime)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.route_information.route_table.route.ospf_sh_time")
	proto.RegisterType((*OspfShTopCommon)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.route_information.route_table.route.ospf_sh_top_common")
	proto.RegisterType((*OspfShTopPath)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.route_information.route_table.route.ospf_sh_top_path")
}

func init() { proto.RegisterFile("ospf_sh_topology.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 840 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0xdf, 0x6f, 0x23, 0x35,
	0x10, 0x56, 0xae, 0x25, 0x49, 0x9d, 0x06, 0x0e, 0x73, 0xf4, 0xb6, 0x08, 0xb8, 0x10, 0x24, 0x88,
	0x10, 0x8a, 0x50, 0x5b, 0x7e, 0xbe, 0xa0, 0x0a, 0xae, 0xa2, 0xa2, 0x57, 0x9d, 0xb6, 0x3d, 0x24,
	0x9e, 0x2c, 0xc7, 0x3b, 0xdb, 0x58, 0xda, 0x5d, 0x2f, 0xb6, 0x13, 0xd2, 0x47, 0xfe, 0x85, 0x93,
	0x78, 0xe0, 0x89, 0x47, 0xfe, 0x4d, 0xe4, 0x19, 0x6f, 0x77, 0xef, 0xf2, 0x4c, 0x79, 0x89, 0xbc,
	0xdf, 0x37, 0x9e, 0xf9, 0x3c, 0x33, 0x1e, 0x87, 0x1d, 0x18, 0x57, 0xe7, 0xc2, 0x2d, 0x85, 0x37,
	0xb5, 0x29, 0xcc, 0xcd, 0xed, 0xbc, 0xb6, 0xc6, 0x1b, 0x2e, 0x94, 0x76, 0xca, 0x08, 0x6d, 0x9c,
	0xd8, 0x58, 0xa1, 0xeb, 0xf5, 0x89, 0x40, 0x4b, 0x53, 0x83, 0x9d, 0x87, 0x55, 0xb0, 0x53, 0xe0,
	0x1c, 0xb8, 0x66, 0x35, 0x5f, 0xdb, 0x1c, 0x7f, 0xe6, 0xd6, 0xac, 0x3c, 0x08, 0x5d, 0xe5, 0xc6,
	0x96, 0xd2, 0x6b, 0x53, 0x45, 0xc4, 0xcb, 0x45, 0x01, 0xb4, 0x9e, 0xbe, 0xec, 0xb1, 0x77, 0x5f,
	0x8f, 0x2d, 0x7e, 0x7e, 0xfa, 0xeb, 0x15, 0xff, 0x88, 0xed, 0x47, 0x8f, 0xa2, 0x92, 0x25, 0x24,
	0xbd, 0x49, 0x6f, 0xb6, 0x97, 0x8e, 0x22, 0x76, 0x29, 0x4b, 0xe0, 0x87, 0x6c, 0xb8, 0xb6, 0x39,
	0xd1, 0x0f, 0x90, 0x1e, 0xac, 0x6d, 0x8e, 0xd4, 0x01, 0xeb, 0xd7, 0x16, 0x72, 0xbd, 0x49, 0x76,
	0x90, 0x88, 0x5f, 0xfc, 0x63, 0x36, 0xa6, 0x95, 0x28, 0xa0, 0xba, 0xf1, 0xcb, 0x64, 0x77, 0xd2,
	0x9b, 0x8d, 0xd3, 0x7d, 0x02, 0x2f, 0x10, 0x9b, 0xfe, 0xb1, 0xcb, 0x1e, 0xbe, 0x2e, 0x2a, 0xe8,
	0x21, 0xf9, 0xd1, 0xef, 0x11, 0xe9, 0x41, 0xec, 0x39, 0x39, 0x9f, 0xb3, 0x77, 0xba, 0x26, 0x4d,
	0x88, 0x63, 0x0c, 0xf1, 0x76, 0xc7, 0x92, 0xe2, 0xb4, 0x2e, 0x4b, 0xf0, 0x56, 0xab, 0xe4, 0x04,
	0x0d, 0xc9, 0xe5, 0x33, 0x84, 0xf8, 0x07, 0x8c, 0xc5, 0xa4, 0xdd, 0xd6, 0x90, 0x7c, 0x89, 0x31,
	0xf7, 0x10, 0xb9, 0xbe, 0xad, 0x81, 0x7f, 0xca, 0xde, 0x22, 0x5a, 0x99, 0xaa, 0x02, 0xe5, 0x21,
	0x4b, 0xbe, 0x9a, 0xf4, 0x66, 0xc3, 0xf4, 0x4d, 0x84, 0x7f, 0x68, 0x50, 0xfe, 0xb2, 0xd7, 0x38,
	0x0a, 0xf5, 0x48, 0xbe, 0x9e, 0xf4, 0x66, 0xa3, 0x23, 0x37, 0xff, 0x8f, 0xcb, 0x3b, 0xef, 0x64,
	0x51, 0x28, 0x53, 0x96, 0xa6, 0x8a, 0xea, 0xcf, 0xab, 0xdc, 0xf0, 0xbf, 0x7a, 0x8d, 0xfc, 0x5a,
	0xfa, 0xa5, 0x28, 0xb4, 0xf3, 0xc9, 0x37, 0x93, 0x9d, 0xd9, 0xe8, 0xe8, 0xb7, 0x7b, 0x55, 0x16,
	0xa2, 0xa7, 0x63, 0xaa, 0x8f, 0xf4, 0xcb, 0x0b, 0xed, 0xfc, 0xf4, 0x8c, 0xed, 0xdf, 0x99, 0x68,
	0x6a, 0x28, 0x07, 0xca, 0x54, 0x19, 0x36, 0xe2, 0x38, 0x8d, 0x5f, 0xfc, 0x43, 0xc6, 0x2a, 0x59,
	0x99, 0xc8, 0x3d, 0x40, 0xae, 0x83, 0x4c, 0xff, 0xee, 0x33, 0xbe, 0x9d, 0x05, 0x3e, 0x65, 0x14,
	0x4f, 0x48, 0x0b, 0x52, 0xe8, 0xc6, 0x2b, 0xd5, 0xfe, 0xd4, 0x82, 0x3c, 0xcf, 0xf8, 0x27, 0x4d,
	0x76, 0xda, 0x0e, 0x21, 0xff, 0xb4, 0xf5, 0xba, 0xe9, 0x91, 0xcf, 0x18, 0xf5, 0x96, 0xb0, 0x7a,
	0x21, 0xd6, 0x60, 0x9d, 0x36, 0x15, 0xb6, 0xfd, 0x38, 0x25, 0x07, 0xa9, 0x5e, 0xfc, 0x42, 0x70,
	0x6b, 0x1b, 0x24, 0x35, 0xb6, 0xe1, 0x0e, 0xec, 0x46, 0xdb, 0xab, 0x3a, 0x6f, 0x6c, 0x4f, 0xd8,
	0x01, 0xd9, 0xe6, 0xc6, 0xfe, 0x2e, 0x6d, 0x26, 0x32, 0xed, 0xbc, 0xac, 0x14, 0x24, 0x6f, 0xa0,
	0xf3, 0x47, 0xc8, 0x9e, 0x11, 0xf9, 0x63, 0xe4, 0xda, 0xa6, 0x76, 0x66, 0x65, 0x15, 0x24, 0xfd,
	0xce, 0xc1, 0xae, 0x10, 0x0a, 0x75, 0x8f, 0x2a, 0x56, 0x75, 0x26, 0xc3, 0x01, 0x75, 0x09, 0xc9,
	0x00, 0x7b, 0xb2, 0xbc, 0xbf, 0xca, 0xeb, 0x12, 0xe2, 0xa1, 0x5f, 0xa0, 0x8c, 0xeb, 0x50, 0xe7,
	0x3f, 0xef, 0x7a, 0x32, 0x97, 0xba, 0x20, 0x65, 0xc3, 0xff, 0x43, 0x19, 0x15, 0xf9, 0x4c, 0xea,
	0x02, 0x75, 0x7d, 0xce, 0x78, 0x5b, 0xb8, 0xda, 0x6a, 0x63, 0xb5, 0xbf, 0x4d, 0xf6, 0x30, 0xb9,
	0x0f, 0x9b, 0xca, 0x3d, 0x8f, 0x78, 0x3b, 0x89, 0xe4, 0xca, 0x1b, 0x01, 0x1b, 0x55, 0xac, 0x32,
	0xc8, 0x12, 0x86, 0xb3, 0x81, 0x72, 0x7f, 0xba, 0xf2, 0xe6, 0x69, 0x24, 0xf8, 0xf7, 0xec, 0xfd,
	0xe8, 0xdd, 0xb6, 0xe3, 0xcb, 0xc2, 0x8d, 0x76, 0x1e, 0x2c, 0x64, 0xc9, 0x08, 0x37, 0x1e, 0x52,
	0x1c, 0xdb, 0x8c, 0xb1, 0xf4, 0xce, 0x80, 0x7f, 0xcb, 0x0e, 0x3b, 0x0e, 0xaa, 0x85, 0xed, 0xee,
	0xde, 0x47, 0x95, 0x07, 0x77, 0xbb, 0x2f, 0x17, 0xb6, 0xdd, 0x3a, 0xfd, 0x67, 0xe7, 0x95, 0x69,
	0x8b, 0xb7, 0x91, 0x7f, 0xc1, 0x1e, 0x35, 0x19, 0xf3, 0x60, 0x73, 0xa9, 0xa0, 0xfb, 0x0a, 0xf0,
	0x38, 0x43, 0x22, 0x85, 0x13, 0xff, 0xb8, 0xe9, 0xd6, 0x0a, 0x36, 0x5e, 0x2c, 0x4d, 0x2d, 0x64,
	0x96, 0x59, 0x70, 0x2e, 0x3e, 0x0d, 0x94, 0x90, 0x4b, 0xd8, 0xf8, 0x9f, 0x4c, 0x7d, 0x4a, 0xd4,
	0x56, 0xb3, 0xee, 0x74, 0x86, 0x7a, 0x6c, 0xd6, 0x27, 0x8c, 0x3e, 0x45, 0xe1, 0xa4, 0xce, 0xf0,
	0xae, 0xec, 0xa5, 0x34, 0x4b, 0x2f, 0x02, 0xc2, 0xbf, 0x63, 0xef, 0x75, 0x86, 0x98, 0x76, 0xa2,
	0x54, 0xd2, 0xf9, 0x20, 0x5c, 0x2a, 0x8f, 0x57, 0x65, 0x18, 0xcf, 0x1e, 0x86, 0xcb, 0xb9, 0x7b,
	0x16, 0xe8, 0x73, 0x64, 0xf9, 0x09, 0x7b, 0xfc, 0xea, 0xde, 0x95, 0x2a, 0x29, 0x03, 0x78, 0x6f,
	0x86, 0x51, 0x35, 0x6d, 0x7c, 0xa1, 0xca, 0x3a, 0xac, 0xb6, 0xde, 0x8d, 0xc1, 0xf6, 0xbb, 0x71,
	0xc8, 0x86, 0x85, 0x93, 0xf4, 0x6a, 0x0c, 0x91, 0x1e, 0x14, 0x4e, 0xe2, 0x9b, 0xf1, 0x98, 0x0d,
	0x9a, 0xa1, 0x43, 0xed, 0xd3, 0x97, 0x34, 0x6f, 0x9e, 0xb0, 0x11, 0x12, 0xd4, 0xa2, 0xb1, 0x59,
	0x58, 0x80, 0xce, 0x10, 0x59, 0xf4, 0xf1, 0x4f, 0xc1, 0xf1, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x31, 0xda, 0xf5, 0x59, 0x2e, 0x08, 0x00, 0x00,
}
