// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ipv6_rib_edm_route.proto

/*
Package cisco_ios_xr_ip_rib_ipv6_oper_ipv6_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_routes_route is a generated protocol buffer package.

It is generated from these files:
	ipv6_rib_edm_route.proto

It has these top-level messages:
	Ipv6RibEdmRoute_KEYS
	Ipv6RibEdmRoute
	Ipv6RibEdmAddr
	Ipv6RibEdmPath
	Ipv6RibEdmPathItem
	RibEdmLocalLabel
*/
package cisco_ios_xr_ip_rib_ipv6_oper_ipv6_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_routes_route

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

// Information of a rib route head and rib proto route
type Ipv6RibEdmRoute_KEYS struct {
	VrfName        string `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
	AfName         string `protobuf:"bytes,2,opt,name=af_name,json=afName" json:"af_name,omitempty"`
	SafName        string `protobuf:"bytes,3,opt,name=saf_name,json=safName" json:"saf_name,omitempty"`
	RouteTableName string `protobuf:"bytes,4,opt,name=route_table_name,json=routeTableName" json:"route_table_name,omitempty"`
	Address        string `protobuf:"bytes,5,opt,name=address" json:"address,omitempty"`
	PrefixLength   uint32 `protobuf:"varint,6,opt,name=prefix_length,json=prefixLength" json:"prefix_length,omitempty"`
	NextHopAddress string `protobuf:"bytes,7,opt,name=next_hop_address,json=nextHopAddress" json:"next_hop_address,omitempty"`
	InterfaceName  string `protobuf:"bytes,8,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
}

func (m *Ipv6RibEdmRoute_KEYS) Reset()                    { *m = Ipv6RibEdmRoute_KEYS{} }
func (m *Ipv6RibEdmRoute_KEYS) String() string            { return proto.CompactTextString(m) }
func (*Ipv6RibEdmRoute_KEYS) ProtoMessage()               {}
func (*Ipv6RibEdmRoute_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Ipv6RibEdmRoute_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ipv6RibEdmRoute_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *Ipv6RibEdmRoute_KEYS) GetSafName() string {
	if m != nil {
		return m.SafName
	}
	return ""
}

func (m *Ipv6RibEdmRoute_KEYS) GetRouteTableName() string {
	if m != nil {
		return m.RouteTableName
	}
	return ""
}

func (m *Ipv6RibEdmRoute_KEYS) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Ipv6RibEdmRoute_KEYS) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

func (m *Ipv6RibEdmRoute_KEYS) GetNextHopAddress() string {
	if m != nil {
		return m.NextHopAddress
	}
	return ""
}

func (m *Ipv6RibEdmRoute_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type Ipv6RibEdmRoute struct {
	// Route prefix
	Prefix string `protobuf:"bytes,50,opt,name=prefix" json:"prefix,omitempty"`
	// Length of prefix
	PrefixLength uint32 `protobuf:"varint,51,opt,name=prefix_length,json=prefixLength" json:"prefix_length,omitempty"`
	// Route version
	RouteVersion uint32 `protobuf:"varint,52,opt,name=route_version,json=routeVersion" json:"route_version,omitempty"`
	// Protocol advertising the route
	ProtocolId uint32 `protobuf:"varint,53,opt,name=protocol_id,json=protocolId" json:"protocol_id,omitempty"`
	//  Name of Protocol
	ProtocolName string `protobuf:"bytes,54,opt,name=protocol_name,json=protocolName" json:"protocol_name,omitempty"`
	// Instance name
	Instance string `protobuf:"bytes,55,opt,name=instance" json:"instance,omitempty"`
	// Client adding the route to RIB
	ClientId uint32 `protobuf:"varint,56,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	// Route type
	RouteType uint32 `protobuf:"varint,57,opt,name=route_type,json=routeType" json:"route_type,omitempty"`
	// Route priority
	Priority uint32 `protobuf:"varint,58,opt,name=priority" json:"priority,omitempty"`
	// SVD Type of route
	SvdType uint32 `protobuf:"varint,59,opt,name=svd_type,json=svdType" json:"svd_type,omitempty"`
	// Route flags
	Flags uint32 `protobuf:"varint,60,opt,name=flags" json:"flags,omitempty"`
	// Extended Route flags
	ExtendedFlags uint64 `protobuf:"varint,61,opt,name=extended_flags,json=extendedFlags" json:"extended_flags,omitempty"`
	// Opaque proto specific info
	Tag uint32 `protobuf:"varint,62,opt,name=tag" json:"tag,omitempty"`
	// Distance of the route
	Distance uint32 `protobuf:"varint,63,opt,name=distance" json:"distance,omitempty"`
	// Diversion distance of the route
	DiversionDistance uint32 `protobuf:"varint,64,opt,name=diversion_distance,json=diversionDistance" json:"diversion_distance,omitempty"`
	// Route metric
	Metric uint32 `protobuf:"varint,65,opt,name=metric" json:"metric,omitempty"`
	// Number of paths
	PathsCount uint32 `protobuf:"varint,66,opt,name=paths_count,json=pathsCount" json:"paths_count,omitempty"`
	// BGP Attribute ID
	AttributeIdentity uint32 `protobuf:"varint,67,opt,name=attribute_identity,json=attributeIdentity" json:"attribute_identity,omitempty"`
	// BGP Traffic Index
	TrafficIndex uint32 `protobuf:"varint,68,opt,name=traffic_index,json=trafficIndex" json:"traffic_index,omitempty"`
	// Route ip precedence
	RoutePrecedence uint32 `protobuf:"varint,69,opt,name=route_precedence,json=routePrecedence" json:"route_precedence,omitempty"`
	// Route qos group
	QosGroup uint32 `protobuf:"varint,70,opt,name=qos_group,json=qosGroup" json:"qos_group,omitempty"`
	// Flow tag
	FlowTag uint32 `protobuf:"varint,71,opt,name=flow_tag,json=flowTag" json:"flow_tag,omitempty"`
	// Forward Class
	FwdClass uint32 `protobuf:"varint,72,opt,name=fwd_class,json=fwdClass" json:"fwd_class,omitempty"`
	// Number of pic paths in this route
	PicCount uint32 `protobuf:"varint,73,opt,name=pic_count,json=picCount" json:"pic_count,omitempty"`
	// Is the route active or backup
	Active bool `protobuf:"varint,74,opt,name=active" json:"active,omitempty"`
	// Route has a diversion path
	Diversion bool `protobuf:"varint,75,opt,name=diversion" json:"diversion,omitempty"`
	// Diversion route protocol name
	DiversionProtoName string `protobuf:"bytes,76,opt,name=diversion_proto_name,json=diversionProtoName" json:"diversion_proto_name,omitempty"`
	// Age of route (seconds)
	RouteAge uint32 `protobuf:"varint,77,opt,name=route_age,json=routeAge" json:"route_age,omitempty"`
	// Local label of the route
	RouteLabel uint32 `protobuf:"varint,78,opt,name=route_label,json=routeLabel" json:"route_label,omitempty"`
	// Route Version
	Version uint32 `protobuf:"varint,79,opt,name=version" json:"version,omitempty"`
	// Table Version
	TblVersion uint64 `protobuf:"varint,80,opt,name=tbl_version,json=tblVersion" json:"tbl_version,omitempty"`
	// Route modification time(nanoseconds)
	RouteModifyTime uint64 `protobuf:"varint,81,opt,name=route_modify_time,json=routeModifyTime" json:"route_modify_time,omitempty"`
	// Path(s) of the route
	RoutePath *Ipv6RibEdmPath `protobuf:"bytes,82,opt,name=route_path,json=routePath" json:"route_path,omitempty"`
}

func (m *Ipv6RibEdmRoute) Reset()                    { *m = Ipv6RibEdmRoute{} }
func (m *Ipv6RibEdmRoute) String() string            { return proto.CompactTextString(m) }
func (*Ipv6RibEdmRoute) ProtoMessage()               {}
func (*Ipv6RibEdmRoute) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Ipv6RibEdmRoute) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *Ipv6RibEdmRoute) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetRouteVersion() uint32 {
	if m != nil {
		return m.RouteVersion
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetProtocolId() uint32 {
	if m != nil {
		return m.ProtocolId
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetProtocolName() string {
	if m != nil {
		return m.ProtocolName
	}
	return ""
}

func (m *Ipv6RibEdmRoute) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

func (m *Ipv6RibEdmRoute) GetClientId() uint32 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetRouteType() uint32 {
	if m != nil {
		return m.RouteType
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetSvdType() uint32 {
	if m != nil {
		return m.SvdType
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetExtendedFlags() uint64 {
	if m != nil {
		return m.ExtendedFlags
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetTag() uint32 {
	if m != nil {
		return m.Tag
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetDistance() uint32 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetDiversionDistance() uint32 {
	if m != nil {
		return m.DiversionDistance
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetMetric() uint32 {
	if m != nil {
		return m.Metric
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetPathsCount() uint32 {
	if m != nil {
		return m.PathsCount
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetAttributeIdentity() uint32 {
	if m != nil {
		return m.AttributeIdentity
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetTrafficIndex() uint32 {
	if m != nil {
		return m.TrafficIndex
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetRoutePrecedence() uint32 {
	if m != nil {
		return m.RoutePrecedence
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetQosGroup() uint32 {
	if m != nil {
		return m.QosGroup
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetFlowTag() uint32 {
	if m != nil {
		return m.FlowTag
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetFwdClass() uint32 {
	if m != nil {
		return m.FwdClass
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetPicCount() uint32 {
	if m != nil {
		return m.PicCount
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *Ipv6RibEdmRoute) GetDiversion() bool {
	if m != nil {
		return m.Diversion
	}
	return false
}

func (m *Ipv6RibEdmRoute) GetDiversionProtoName() string {
	if m != nil {
		return m.DiversionProtoName
	}
	return ""
}

func (m *Ipv6RibEdmRoute) GetRouteAge() uint32 {
	if m != nil {
		return m.RouteAge
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetRouteLabel() uint32 {
	if m != nil {
		return m.RouteLabel
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetTblVersion() uint64 {
	if m != nil {
		return m.TblVersion
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetRouteModifyTime() uint64 {
	if m != nil {
		return m.RouteModifyTime
	}
	return 0
}

func (m *Ipv6RibEdmRoute) GetRoutePath() *Ipv6RibEdmPath {
	if m != nil {
		return m.RoutePath
	}
	return nil
}

type Ipv6RibEdmAddr struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *Ipv6RibEdmAddr) Reset()                    { *m = Ipv6RibEdmAddr{} }
func (m *Ipv6RibEdmAddr) String() string            { return proto.CompactTextString(m) }
func (*Ipv6RibEdmAddr) ProtoMessage()               {}
func (*Ipv6RibEdmAddr) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Ipv6RibEdmAddr) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Information of a rib path
type Ipv6RibEdmPath struct {
	// Next path
	Ipv6RibEdmPath []*Ipv6RibEdmPathItem `protobuf:"bytes,1,rep,name=ipv6_rib_edm_path,json=ipv6RibEdmPath" json:"ipv6_rib_edm_path,omitempty"`
}

func (m *Ipv6RibEdmPath) Reset()                    { *m = Ipv6RibEdmPath{} }
func (m *Ipv6RibEdmPath) String() string            { return proto.CompactTextString(m) }
func (*Ipv6RibEdmPath) ProtoMessage()               {}
func (*Ipv6RibEdmPath) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Ipv6RibEdmPath) GetIpv6RibEdmPath() []*Ipv6RibEdmPathItem {
	if m != nil {
		return m.Ipv6RibEdmPath
	}
	return nil
}

type Ipv6RibEdmPathItem struct {
	// Nexthop
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	// Infosource
	InformationSource string `protobuf:"bytes,2,opt,name=information_source,json=informationSource" json:"information_source,omitempty"`
	// V6 nexthop
	V6Nexthop string `protobuf:"bytes,3,opt,name=v6_nexthop,json=v6Nexthop" json:"v6_nexthop,omitempty"`
	// Interface Name
	InterfaceName string `protobuf:"bytes,4,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
	// Metrics
	Metric uint32 `protobuf:"varint,5,opt,name=metric" json:"metric,omitempty"`
	// Load Metrics
	LoadMetric uint32 `protobuf:"varint,6,opt,name=load_metric,json=loadMetric" json:"load_metric,omitempty"`
	// Flags extended to 64 bits
	Flags64 uint64 `protobuf:"varint,7,opt,name=flags64" json:"flags64,omitempty"`
	// Flags
	Flags uint32 `protobuf:"varint,8,opt,name=flags" json:"flags,omitempty"`
	// Private Flags
	PrivateFlags uint32 `protobuf:"varint,9,opt,name=private_flags,json=privateFlags" json:"private_flags,omitempty"`
	// Looping path
	Looped bool `protobuf:"varint,10,opt,name=looped" json:"looped,omitempty"`
	// Nexthop tableid
	NextHopTableId uint32 `protobuf:"varint,11,opt,name=next_hop_table_id,json=nextHopTableId" json:"next_hop_table_id,omitempty"`
	// VRF Name of the nh table
	NextHopVrfName string `protobuf:"bytes,12,opt,name=next_hop_vrf_name,json=nextHopVrfName" json:"next_hop_vrf_name,omitempty"`
	// NH table name
	NextHopTableName string `protobuf:"bytes,13,opt,name=next_hop_table_name,json=nextHopTableName" json:"next_hop_table_name,omitempty"`
	// NH afi
	NextHopAfi uint32 `protobuf:"varint,14,opt,name=next_hop_afi,json=nextHopAfi" json:"next_hop_afi,omitempty"`
	// NH safi
	NextHopSafi uint32 `protobuf:"varint,15,opt,name=next_hop_safi,json=nextHopSafi" json:"next_hop_safi,omitempty"`
	// Label associated with this path
	RouteLabel uint32 `protobuf:"varint,16,opt,name=route_label,json=routeLabel" json:"route_label,omitempty"`
	// Tunnel ID associated with this path
	TunnelId uint32 `protobuf:"varint,17,opt,name=tunnel_id,json=tunnelId" json:"tunnel_id,omitempty"`
	// Path id of this path
	Pathid uint32 `protobuf:"varint,18,opt,name=pathid" json:"pathid,omitempty"`
	// Path id of this path's backup
	BackupPathid uint32 `protobuf:"varint,19,opt,name=backup_pathid,json=backupPathid" json:"backup_pathid,omitempty"`
	// Refcnt of backup
	RefCntOfBackup uint32 `protobuf:"varint,20,opt,name=ref_cnt_of_backup,json=refCntOfBackup" json:"ref_cnt_of_backup,omitempty"`
	// Number of extended communities
	NumberOfExtendedCommunities uint32 `protobuf:"varint,21,opt,name=number_of_extended_communities,json=numberOfExtendedCommunities" json:"number_of_extended_communities,omitempty"`
	// MVPN attribute present
	MvpnPresent bool `protobuf:"varint,22,opt,name=mvpn_present,json=mvpnPresent" json:"mvpn_present,omitempty"`
	// Path RT present
	PathrtPresent           bool `protobuf:"varint,23,opt,name=pathrt_present,json=pathrtPresent" json:"pathrt_present,omitempty"`
	VrfimportrtPresent      bool `protobuf:"varint,24,opt,name=vrfimportrt_present,json=vrfimportrtPresent" json:"vrfimportrt_present,omitempty"`
	SourceasrtPresent       bool `protobuf:"varint,25,opt,name=sourceasrt_present,json=sourceasrtPresent" json:"sourceasrt_present,omitempty"`
	SourcerdPresent         bool `protobuf:"varint,26,opt,name=sourcerd_present,json=sourcerdPresent" json:"sourcerd_present,omitempty"`
	SegmentedNexthopPresent bool `protobuf:"varint,27,opt,name=segmented_nexthop_present,json=segmentedNexthopPresent" json:"segmented_nexthop_present,omitempty"`
	// NHID associated with this path
	NextHopId uint32 `protobuf:"varint,28,opt,name=next_hop_id,json=nextHopId" json:"next_hop_id,omitempty"`
	// NHID references
	NextHopIdRefcount uint32 `protobuf:"varint,29,opt,name=next_hop_id_refcount,json=nextHopIdRefcount" json:"next_hop_id_refcount,omitempty"`
	// OSPF area associated with the path
	OspfAreaId string `protobuf:"bytes,30,opt,name=ospf_area_id,json=ospfAreaId" json:"ospf_area_id,omitempty"`
	// Remote backup node address
	RemoteBackupAddr []*Ipv6RibEdmAddr `protobuf:"bytes,31,rep,name=remote_backup_addr,json=remoteBackupAddr" json:"remote_backup_addr,omitempty"`
	// Path has a label stack
	HasLabelstk bool `protobuf:"varint,32,opt,name=has_labelstk,json=hasLabelstk" json:"has_labelstk,omitempty"`
	// Number of labels in stack
	NumLabels uint32 `protobuf:"varint,33,opt,name=num_labels,json=numLabels" json:"num_labels,omitempty"`
	// Labels for this path
	Labelstk []uint32 `protobuf:"varint,34,rep,packed,name=labelstk" json:"labelstk,omitempty"`
	// binding Label for this path
	BindingLabel uint32 `protobuf:"varint,35,opt,name=binding_label,json=bindingLabel" json:"binding_label,omitempty"`
	// Fib nhid encap id
	NhidFeid uint64 `protobuf:"varint,36,opt,name=nhid_feid,json=nhidFeid" json:"nhid_feid,omitempty"`
	// Fib mpls encap id
	MplsFeid uint64 `protobuf:"varint,37,opt,name=mpls_feid,json=mplsFeid" json:"mpls_feid,omitempty"`
}

func (m *Ipv6RibEdmPathItem) Reset()                    { *m = Ipv6RibEdmPathItem{} }
func (m *Ipv6RibEdmPathItem) String() string            { return proto.CompactTextString(m) }
func (*Ipv6RibEdmPathItem) ProtoMessage()               {}
func (*Ipv6RibEdmPathItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Ipv6RibEdmPathItem) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Ipv6RibEdmPathItem) GetInformationSource() string {
	if m != nil {
		return m.InformationSource
	}
	return ""
}

func (m *Ipv6RibEdmPathItem) GetV6Nexthop() string {
	if m != nil {
		return m.V6Nexthop
	}
	return ""
}

func (m *Ipv6RibEdmPathItem) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *Ipv6RibEdmPathItem) GetMetric() uint32 {
	if m != nil {
		return m.Metric
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetLoadMetric() uint32 {
	if m != nil {
		return m.LoadMetric
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetFlags64() uint64 {
	if m != nil {
		return m.Flags64
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetPrivateFlags() uint32 {
	if m != nil {
		return m.PrivateFlags
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetLooped() bool {
	if m != nil {
		return m.Looped
	}
	return false
}

func (m *Ipv6RibEdmPathItem) GetNextHopTableId() uint32 {
	if m != nil {
		return m.NextHopTableId
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetNextHopVrfName() string {
	if m != nil {
		return m.NextHopVrfName
	}
	return ""
}

func (m *Ipv6RibEdmPathItem) GetNextHopTableName() string {
	if m != nil {
		return m.NextHopTableName
	}
	return ""
}

func (m *Ipv6RibEdmPathItem) GetNextHopAfi() uint32 {
	if m != nil {
		return m.NextHopAfi
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetNextHopSafi() uint32 {
	if m != nil {
		return m.NextHopSafi
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetRouteLabel() uint32 {
	if m != nil {
		return m.RouteLabel
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetTunnelId() uint32 {
	if m != nil {
		return m.TunnelId
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetPathid() uint32 {
	if m != nil {
		return m.Pathid
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetBackupPathid() uint32 {
	if m != nil {
		return m.BackupPathid
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetRefCntOfBackup() uint32 {
	if m != nil {
		return m.RefCntOfBackup
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetNumberOfExtendedCommunities() uint32 {
	if m != nil {
		return m.NumberOfExtendedCommunities
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetMvpnPresent() bool {
	if m != nil {
		return m.MvpnPresent
	}
	return false
}

func (m *Ipv6RibEdmPathItem) GetPathrtPresent() bool {
	if m != nil {
		return m.PathrtPresent
	}
	return false
}

func (m *Ipv6RibEdmPathItem) GetVrfimportrtPresent() bool {
	if m != nil {
		return m.VrfimportrtPresent
	}
	return false
}

func (m *Ipv6RibEdmPathItem) GetSourceasrtPresent() bool {
	if m != nil {
		return m.SourceasrtPresent
	}
	return false
}

func (m *Ipv6RibEdmPathItem) GetSourcerdPresent() bool {
	if m != nil {
		return m.SourcerdPresent
	}
	return false
}

func (m *Ipv6RibEdmPathItem) GetSegmentedNexthopPresent() bool {
	if m != nil {
		return m.SegmentedNexthopPresent
	}
	return false
}

func (m *Ipv6RibEdmPathItem) GetNextHopId() uint32 {
	if m != nil {
		return m.NextHopId
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetNextHopIdRefcount() uint32 {
	if m != nil {
		return m.NextHopIdRefcount
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetOspfAreaId() string {
	if m != nil {
		return m.OspfAreaId
	}
	return ""
}

func (m *Ipv6RibEdmPathItem) GetRemoteBackupAddr() []*Ipv6RibEdmAddr {
	if m != nil {
		return m.RemoteBackupAddr
	}
	return nil
}

func (m *Ipv6RibEdmPathItem) GetHasLabelstk() bool {
	if m != nil {
		return m.HasLabelstk
	}
	return false
}

func (m *Ipv6RibEdmPathItem) GetNumLabels() uint32 {
	if m != nil {
		return m.NumLabels
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetLabelstk() []uint32 {
	if m != nil {
		return m.Labelstk
	}
	return nil
}

func (m *Ipv6RibEdmPathItem) GetBindingLabel() uint32 {
	if m != nil {
		return m.BindingLabel
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetNhidFeid() uint64 {
	if m != nil {
		return m.NhidFeid
	}
	return 0
}

func (m *Ipv6RibEdmPathItem) GetMplsFeid() uint64 {
	if m != nil {
		return m.MplsFeid
	}
	return 0
}

// Information of local label for route head
type RibEdmLocalLabel struct {
	// Protocol Name
	ProtocolName string `protobuf:"bytes,1,opt,name=protocol_name,json=protocolName" json:"protocol_name,omitempty"`
	// Client ID
	ClientId uint32 `protobuf:"varint,2,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	// Stale
	Stale uint32 `protobuf:"varint,3,opt,name=stale" json:"stale,omitempty"`
	// Mirrored
	Mirrored uint32 `protobuf:"varint,4,opt,name=mirrored" json:"mirrored,omitempty"`
	// Merge disable
	MergeDisable uint32 `protobuf:"varint,5,opt,name=merge_disable,json=mergeDisable" json:"merge_disable,omitempty"`
	// Redist only
	RedistOnly uint32 `protobuf:"varint,6,opt,name=redist_only,json=redistOnly" json:"redist_only,omitempty"`
	// Local label
	Label uint32 `protobuf:"varint,7,opt,name=label" json:"label,omitempty"`
	// Distance
	Distance uint32 `protobuf:"varint,8,opt,name=distance" json:"distance,omitempty"`
}

func (m *RibEdmLocalLabel) Reset()                    { *m = RibEdmLocalLabel{} }
func (m *RibEdmLocalLabel) String() string            { return proto.CompactTextString(m) }
func (*RibEdmLocalLabel) ProtoMessage()               {}
func (*RibEdmLocalLabel) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RibEdmLocalLabel) GetProtocolName() string {
	if m != nil {
		return m.ProtocolName
	}
	return ""
}

func (m *RibEdmLocalLabel) GetClientId() uint32 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *RibEdmLocalLabel) GetStale() uint32 {
	if m != nil {
		return m.Stale
	}
	return 0
}

func (m *RibEdmLocalLabel) GetMirrored() uint32 {
	if m != nil {
		return m.Mirrored
	}
	return 0
}

func (m *RibEdmLocalLabel) GetMergeDisable() uint32 {
	if m != nil {
		return m.MergeDisable
	}
	return 0
}

func (m *RibEdmLocalLabel) GetRedistOnly() uint32 {
	if m != nil {
		return m.RedistOnly
	}
	return 0
}

func (m *RibEdmLocalLabel) GetLabel() uint32 {
	if m != nil {
		return m.Label
	}
	return 0
}

func (m *RibEdmLocalLabel) GetDistance() uint32 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func init() {
	proto.RegisterType((*Ipv6RibEdmRoute_KEYS)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.routes.route.ipv6_rib_edm_route_KEYS")
	proto.RegisterType((*Ipv6RibEdmRoute)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.routes.route.ipv6_rib_edm_route")
	proto.RegisterType((*Ipv6RibEdmAddr)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.routes.route.ipv6_rib_edm_addr")
	proto.RegisterType((*Ipv6RibEdmPath)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.routes.route.ipv6_rib_edm_path")
	proto.RegisterType((*Ipv6RibEdmPathItem)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.routes.route.ipv6_rib_edm_path_item")
	proto.RegisterType((*RibEdmLocalLabel)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.routes.route.rib_edm_local_label")
}

func init() { proto.RegisterFile("ipv6_rib_edm_route.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1528 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x57, 0xcd, 0x72, 0x1b, 0xc7,
	0x11, 0x2e, 0xf0, 0x17, 0x18, 0x12, 0x14, 0xb1, 0x64, 0xc8, 0xa1, 0x69, 0xcb, 0x30, 0x15, 0x57,
	0x91, 0xa9, 0x22, 0x9c, 0x92, 0x1d, 0x26, 0x71, 0x7e, 0x69, 0x8a, 0xb2, 0x11, 0x53, 0x22, 0x03,
	0xa9, 0x54, 0x95, 0xd3, 0xd4, 0x60, 0x67, 0x16, 0x98, 0xd2, 0xfe, 0x69, 0x66, 0x00, 0x91, 0x87,
	0x1c, 0x72, 0xc9, 0x39, 0x97, 0xbc, 0x40, 0x9e, 0x20, 0x6f, 0x90, 0xca, 0x13, 0xe4, 0x98, 0xd7,
	0x49, 0x75, 0xf7, 0xec, 0x02, 0xfc, 0xc9, 0xd9, 0xba, 0x08, 0xea, 0xaf, 0xbf, 0x9e, 0xe9, 0xed,
	0xdf, 0x21, 0xe3, 0xa6, 0x9c, 0x9e, 0x08, 0x6b, 0x86, 0x42, 0xab, 0x4c, 0xd8, 0x62, 0xe2, 0x75,
	0xaf, 0xb4, 0x85, 0x2f, 0xa2, 0x3f, 0xc7, 0xc6, 0xc5, 0x85, 0x30, 0x85, 0x13, 0xd7, 0x56, 0x98,
	0x12, 0x49, 0xc8, 0x2e, 0x4a, 0x6d, 0x7b, 0x95, 0x5d, 0x6f, 0x6a, 0x13, 0x07, 0xff, 0xf4, 0x64,
	0xe2, 0x7a, 0x32, 0xe9, 0x39, 0xf8, 0x75, 0x32, 0xe9, 0x05, 0x13, 0x3c, 0x53, 0x78, 0x39, 0x4c,
	0xb5, 0xc8, 0x65, 0xa6, 0xdd, 0xff, 0x53, 0xf4, 0x10, 0x70, 0xf4, 0x73, 0xf0, 0x8f, 0x05, 0xb6,
	0x7b, 0xdf, 0x37, 0xf1, 0xfd, 0xf9, 0x9f, 0x5e, 0x45, 0x7b, 0xac, 0x39, 0xb5, 0x09, 0x1a, 0xf1,
	0x46, 0xb7, 0x71, 0xd8, 0x1a, 0xac, 0x4e, 0x6d, 0xf2, 0x52, 0x66, 0x3a, 0xda, 0x65, 0xab, 0x32,
	0x68, 0x16, 0x50, 0xb3, 0x22, 0x49, 0xb1, 0xc7, 0x9a, 0xae, 0xd2, 0x2c, 0x92, 0x8d, 0x0b, 0xaa,
	0x43, 0xb6, 0x79, 0xd7, 0x17, 0xbe, 0x84, 0x94, 0x0d, 0xc4, 0x5f, 0x03, 0x8c, 0x4c, 0xce, 0x56,
	0xa5, 0x52, 0x56, 0x3b, 0xc7, 0x97, 0xe9, 0x8c, 0x20, 0x46, 0x4f, 0x58, 0xbb, 0xb4, 0x3a, 0x31,
	0xd7, 0x22, 0xd5, 0xf9, 0xc8, 0x8f, 0xf9, 0x4a, 0xb7, 0x71, 0xd8, 0x1e, 0xac, 0x13, 0x78, 0x81,
	0x18, 0x5c, 0x94, 0xeb, 0x6b, 0x2f, 0xc6, 0x45, 0x29, 0xaa, 0x73, 0x56, 0xe9, 0x22, 0xc0, 0xbf,
	0x2b, 0xca, 0xd3, 0x70, 0xdc, 0xe7, 0x6c, 0xc3, 0xe4, 0x5e, 0xdb, 0x44, 0xc6, 0xc1, 0xa1, 0x26,
	0xf2, 0xda, 0x35, 0x0a, 0xfe, 0x1c, 0xfc, 0xa7, 0xc5, 0xa2, 0xfb, 0x41, 0x8a, 0x76, 0xd8, 0x0a,
	0xdd, 0xcb, 0x9f, 0x52, 0x0c, 0x48, 0xba, 0xef, 0xe4, 0x97, 0x0f, 0x38, 0xf9, 0x84, 0xb5, 0x29,
	0x1a, 0x53, 0x6d, 0x9d, 0x29, 0x72, 0xfe, 0x15, 0x91, 0x10, 0x7c, 0x43, 0x58, 0xf4, 0x29, 0x5b,
	0xc3, 0x2a, 0x89, 0x8b, 0x54, 0x18, 0xc5, 0x7f, 0x86, 0x14, 0x56, 0x41, 0x7d, 0x45, 0x57, 0x05,
	0x02, 0xfa, 0x7f, 0x82, 0x9e, 0xac, 0x57, 0x20, 0x86, 0xf3, 0x23, 0xd6, 0x34, 0xb9, 0xf3, 0x32,
	0x8f, 0x35, 0xff, 0x39, 0xea, 0x6b, 0x39, 0xda, 0x67, 0xad, 0x38, 0x35, 0x3a, 0xf7, 0x70, 0xfe,
	0x2f, 0xf0, 0xfc, 0x26, 0x01, 0x7d, 0x15, 0x7d, 0xc2, 0x58, 0xc8, 0xd8, 0x4d, 0xa9, 0xf9, 0x2f,
	0x51, 0xdb, 0xa2, 0x5c, 0xdd, 0x94, 0x78, 0x6e, 0x69, 0x4d, 0x61, 0x8d, 0xbf, 0xe1, 0x5f, 0x93,
	0x69, 0x25, 0x63, 0x1d, 0x4c, 0x15, 0x19, 0xfe, 0x0a, 0x75, 0xab, 0x6e, 0xaa, 0xd0, 0x6c, 0x9b,
	0x2d, 0x27, 0xa9, 0x1c, 0x39, 0xfe, 0x6b, 0xc4, 0x49, 0x80, 0x54, 0xe8, 0x6b, 0xaf, 0x73, 0xa5,
	0x95, 0x20, 0xf5, 0x6f, 0xba, 0x8d, 0xc3, 0xa5, 0x41, 0xbb, 0x42, 0x9f, 0x23, 0x6d, 0x93, 0x2d,
	0x7a, 0x39, 0xe2, 0xbf, 0x45, 0x53, 0xf8, 0x2f, 0x78, 0xa1, 0x4c, 0xf8, 0xba, 0xdf, 0x91, 0x17,
	0x95, 0x1c, 0x1d, 0xb3, 0x48, 0x99, 0x10, 0x60, 0x51, 0xb3, 0x7e, 0x8f, 0xac, 0x4e, 0xad, 0x79,
	0x56, 0xd1, 0x77, 0xd8, 0x4a, 0xa6, 0xbd, 0x35, 0x31, 0x3f, 0x45, 0x4a, 0x90, 0x30, 0x0d, 0xd2,
	0x8f, 0x9d, 0x88, 0x8b, 0x49, 0xee, 0xf9, 0x37, 0x21, 0x0d, 0x00, 0x9d, 0x01, 0x02, 0xf7, 0x48,
	0xef, 0xad, 0x19, 0x42, 0xb0, 0x8c, 0xd2, 0xb9, 0x87, 0x98, 0x9c, 0xd1, 0x3d, 0xb5, 0xa6, 0x1f,
	0x14, 0x90, 0x35, 0x6f, 0x65, 0x92, 0x98, 0x58, 0x98, 0x5c, 0xe9, 0x6b, 0xfe, 0x8c, 0x72, 0x1f,
	0xc0, 0x3e, 0x60, 0xd1, 0x51, 0xd5, 0x2e, 0xa5, 0xd5, 0xb1, 0x56, 0x1a, 0x3c, 0x3f, 0x47, 0xde,
	0x23, 0xc4, 0xaf, 0x6a, 0x18, 0x92, 0xf8, 0xae, 0x70, 0x62, 0x64, 0x8b, 0x49, 0xc9, 0x9f, 0x53,
	0x0c, 0xde, 0x15, 0xee, 0x5b, 0x90, 0x21, 0x13, 0x49, 0x5a, 0xbc, 0x17, 0x10, 0xb6, 0x6f, 0x29,
	0x13, 0x20, 0xbf, 0x96, 0x23, 0xb0, 0x4b, 0xde, 0x2b, 0x11, 0xa7, 0xd2, 0x39, 0xfe, 0x1d, 0xd9,
	0x25, 0xef, 0xd5, 0x19, 0xc8, 0xa0, 0x2c, 0x4d, 0x1c, 0x3e, 0xb9, 0x1f, 0xd2, 0x6b, 0x62, 0xfa,
	0xe0, 0x1d, 0xb6, 0x22, 0x63, 0x6f, 0xa6, 0x9a, 0xff, 0xa1, 0xdb, 0x38, 0x6c, 0x0e, 0x82, 0x14,
	0x7d, 0xcc, 0x5a, 0x75, 0x58, 0xf9, 0xf7, 0xa8, 0x9a, 0x01, 0xd1, 0x4f, 0xd9, 0xf6, 0x2c, 0x1d,
	0x58, 0xa2, 0x54, 0xb4, 0x17, 0x58, 0x94, 0xb3, 0x54, 0x5d, 0x81, 0x0a, 0x4b, 0x77, 0x9f, 0x51,
	0xbd, 0x09, 0x39, 0xd2, 0xfc, 0x05, 0x39, 0x81, 0xc0, 0xe9, 0x48, 0x43, 0x5a, 0x48, 0x99, 0xca,
	0xa1, 0x4e, 0xf9, 0x4b, 0x4a, 0x0b, 0x42, 0x17, 0x80, 0xc0, 0x1c, 0xa9, 0x7c, 0xb9, 0xa4, 0x2f,
	0x9f, 0xce, 0x1a, 0xcb, 0x0f, 0xd3, 0xba, 0xf7, 0xae, 0xb0, 0xd4, 0x98, 0x1f, 0xa6, 0x55, 0xe7,
	0xfd, 0x84, 0x75, 0xe8, 0xec, 0xac, 0x50, 0x26, 0xb9, 0x11, 0xde, 0x64, 0x9a, 0xff, 0x11, 0x69,
	0x14, 0xfe, 0x17, 0x88, 0xbf, 0x36, 0x99, 0x8e, 0xfe, 0xd9, 0xa8, 0xfa, 0x04, 0x4a, 0x82, 0x0f,
	0xba, 0x8d, 0xc3, 0xb5, 0xa7, 0x7f, 0x6b, 0xf4, 0x7e, 0xd0, 0xc9, 0xde, 0xbb, 0x35, 0xb0, 0xc0,
	0xb1, 0xd0, 0xba, 0x57, 0xd2, 0x8f, 0x0f, 0x8e, 0x58, 0xe7, 0x96, 0x1e, 0xc6, 0x24, 0x34, 0xe6,
	0x54, 0xa6, 0x93, 0x6a, 0xd8, 0x93, 0x70, 0xf0, 0xdf, 0xc6, 0x1d, 0x2e, 0x9c, 0x15, 0xfd, 0xfb,
	0x21, 0x94, 0x37, 0xba, 0x8b, 0x87, 0x6b, 0x4f, 0xff, 0xfe, 0xc1, 0x7d, 0xba, 0x30, 0x5e, 0x67,
	0x83, 0x0d, 0xc0, 0x07, 0x66, 0x78, 0xae, 0x32, 0x0c, 0xc2, 0x5f, 0xd7, 0xd9, 0xce, 0xc3, 0xd4,
	0xf9, 0x0d, 0xd4, 0xb8, 0xbd, 0x81, 0x8e, 0x59, 0x64, 0xf2, 0xa4, 0xb0, 0x99, 0xf4, 0x50, 0xc5,
	0xae, 0x98, 0xd8, 0xb8, 0x5a, 0x82, 0x9d, 0x39, 0xcd, 0x2b, 0x54, 0xc0, 0x08, 0x9d, 0x9e, 0x08,
	0x58, 0x3b, 0xe3, 0xa2, 0x0c, 0x1b, 0xb1, 0x35, 0x3d, 0x79, 0x49, 0xc0, 0x03, 0x0b, 0x68, 0xe9,
	0x81, 0x05, 0x34, 0x37, 0x98, 0x96, 0xef, 0x0e, 0xa6, 0xb4, 0x90, 0x4a, 0x04, 0x25, 0x2d, 0x43,
	0x06, 0xd0, 0x0b, 0x22, 0x70, 0xb6, 0x8a, 0xc3, 0xf4, 0xe4, 0x2b, 0xdc, 0x80, 0x4b, 0x83, 0x4a,
	0x9c, 0x4d, 0xe1, 0xe6, 0xfc, 0x14, 0xc6, 0x7d, 0x62, 0xa6, 0xd2, 0xeb, 0x30, 0x84, 0x5b, 0xd5,
	0xea, 0x42, 0x90, 0x66, 0xf0, 0x0e, 0x5b, 0x49, 0x8b, 0xa2, 0xd4, 0x8a, 0x33, 0x6a, 0x7e, 0x92,
	0xa2, 0x23, 0xd6, 0xa9, 0xf7, 0x2e, 0xa5, 0xc6, 0x28, 0xbe, 0x86, 0x07, 0x54, 0x8b, 0x17, 0x77,
	0x7c, 0xff, 0x36, 0xb5, 0x7e, 0x63, 0xac, 0xdf, 0xda, 0xd1, 0x6f, 0xc2, 0x53, 0xe3, 0x98, 0x6d,
	0xdd, 0x39, 0x15, 0xc9, 0x6d, 0x24, 0x6f, 0xce, 0x9f, 0x8b, 0xf4, 0x2e, 0x5b, 0x9f, 0x2d, 0xff,
	0xc4, 0xf0, 0x0d, 0x8a, 0x49, 0xb5, 0xf8, 0x13, 0x13, 0x1d, 0xb0, 0x76, 0xcd, 0x70, 0x40, 0x79,
	0x84, 0x94, 0xb5, 0x40, 0x79, 0x25, 0x13, 0x73, 0x77, 0xb4, 0x6c, 0xde, 0x1b, 0x2d, 0xfb, 0xac,
	0xe5, 0x27, 0x79, 0xae, 0x71, 0x2f, 0x77, 0x68, 0x30, 0x11, 0xd0, 0x57, 0xf8, 0x30, 0x90, 0x7e,
	0x6c, 0x14, 0x8f, 0x28, 0x5d, 0x24, 0x41, 0x74, 0x87, 0x32, 0x7e, 0x3b, 0x29, 0x45, 0x50, 0x6f,
	0x51, 0x74, 0x09, 0xbc, 0x22, 0xd2, 0x11, 0xeb, 0x58, 0x9d, 0x88, 0x38, 0xf7, 0xa2, 0x48, 0x04,
	0xa9, 0xf8, 0x36, 0x45, 0xd1, 0xea, 0xe4, 0x2c, 0xf7, 0x97, 0xc9, 0x37, 0x88, 0x46, 0x67, 0xec,
	0x71, 0x3e, 0xc9, 0x86, 0xda, 0x02, 0xb3, 0xde, 0x9e, 0x71, 0x91, 0x65, 0x93, 0xdc, 0x78, 0xa3,
	0x1d, 0xff, 0x11, 0xda, 0xed, 0x13, 0xeb, 0x32, 0x39, 0x0f, 0x9c, 0xb3, 0x19, 0x25, 0xfa, 0x8c,
	0xad, 0x67, 0xd3, 0x12, 0xe6, 0xb1, 0x76, 0x3a, 0xf7, 0x7c, 0x07, 0x73, 0xba, 0x06, 0xd8, 0x15,
	0x41, 0x50, 0xa5, 0xe0, 0xb0, 0xf5, 0x35, 0x69, 0x17, 0x49, 0x6d, 0x42, 0x2b, 0xda, 0x17, 0x6c,
	0x6b, 0x6a, 0x13, 0x93, 0x95, 0x85, 0xf5, 0x73, 0x5c, 0x8e, 0xdc, 0x68, 0x4e, 0x55, 0x19, 0x1c,
	0xb3, 0x88, 0xfa, 0x47, 0xba, 0x39, 0xfe, 0x1e, 0xf2, 0x3b, 0x33, 0x4d, 0x45, 0x3f, 0x62, 0x9b,
	0x04, 0x5a, 0x55, 0x93, 0x3f, 0x42, 0xf2, 0xa3, 0x0a, 0xaf, 0xa8, 0x5f, 0xb3, 0x3d, 0xa7, 0x47,
	0x99, 0xce, 0xbd, 0x56, 0x55, 0xf7, 0xd5, 0x36, 0xfb, 0x68, 0xb3, 0x5b, 0x13, 0x42, 0x33, 0x56,
	0xb6, 0x8f, 0xd9, 0x5a, 0x5d, 0x1f, 0x46, 0xf1, 0x8f, 0xe9, 0xd9, 0x13, 0xaa, 0xa3, 0xaf, 0xa2,
	0x2f, 0xd8, 0xf6, 0x9c, 0x5e, 0x58, 0x9d, 0xd0, 0x8e, 0xfc, 0x84, 0xd6, 0x7d, 0x4d, 0x1c, 0x04,
	0x05, 0x94, 0x64, 0xe1, 0xca, 0x44, 0x48, 0xab, 0x25, 0x9c, 0xf8, 0x18, 0x4b, 0x97, 0x01, 0x76,
	0x6a, 0xb5, 0xec, 0xab, 0xe8, 0x5f, 0x0d, 0x16, 0x59, 0x9d, 0x15, 0x5e, 0x87, 0x84, 0xe3, 0x40,
	0xe6, 0x9f, 0xe2, 0x38, 0xfd, 0xb0, 0x36, 0x09, 0x38, 0x36, 0xd8, 0x24, 0x67, 0xa9, 0x0c, 0xe1,
	0x31, 0x0d, 0x55, 0x34, 0x96, 0x8e, 0xda, 0xc5, 0xf9, 0xb7, 0xbc, 0x4b, 0x55, 0x34, 0x96, 0xee,
	0x22, 0x40, 0x30, 0x0a, 0xf3, 0x49, 0x16, 0x28, 0xfc, 0xb3, 0x10, 0xd6, 0x49, 0x46, 0x04, 0x78,
	0xc7, 0xd5, 0xd6, 0x07, 0xdd, 0x45, 0x68, 0xa8, 0x4a, 0xc6, 0xc6, 0x31, 0xb9, 0x32, 0xf9, 0x28,
	0x34, 0xe4, 0x93, 0xd0, 0x38, 0x04, 0xd6, 0x2d, 0x99, 0x8f, 0x8d, 0x12, 0x89, 0x36, 0x8a, 0xff,
	0x18, 0xa7, 0x5d, 0x13, 0x80, 0xe7, 0xda, 0x28, 0x50, 0x66, 0x65, 0xea, 0x48, 0xf9, 0x39, 0x29,
	0x01, 0x00, 0xe5, 0xc1, 0x5f, 0x16, 0xd8, 0x56, 0xf5, 0x7d, 0x69, 0x11, 0xcb, 0x94, 0x6e, 0xb9,
	0xff, 0xba, 0x6e, 0x3c, 0xf0, 0xba, 0xbe, 0xf5, 0x82, 0x5e, 0xb8, 0xf3, 0x82, 0xde, 0x66, 0xcb,
	0xce, 0xcb, 0x94, 0xfe, 0x16, 0x6a, 0x0f, 0x48, 0x80, 0x4f, 0xcd, 0x8c, 0xb5, 0x85, 0xd5, 0x0a,
	0xe7, 0x7d, 0x7b, 0x50, 0xcb, 0x70, 0x67, 0xa6, 0xed, 0x48, 0xc3, 0x73, 0x15, 0xf2, 0x11, 0x26,
	0xfe, 0x3a, 0x82, 0xcf, 0x08, 0xc3, 0xf1, 0xa4, 0xe1, 0x3d, 0x2b, 0x8a, 0x3c, 0xbd, 0xa9, 0xe6,
	0x3e, 0x41, 0x97, 0x79, 0x7a, 0x03, 0xf7, 0x52, 0xa0, 0x56, 0xe9, 0x5e, 0xfa, 0x9e, 0xf9, 0xa7,
	0x72, 0xf3, 0xf6, 0x53, 0x79, 0xb8, 0x82, 0x1f, 0xf5, 0xe5, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x20, 0x18, 0x26, 0x97, 0xaa, 0x0e, 0x00, 0x00,
}
