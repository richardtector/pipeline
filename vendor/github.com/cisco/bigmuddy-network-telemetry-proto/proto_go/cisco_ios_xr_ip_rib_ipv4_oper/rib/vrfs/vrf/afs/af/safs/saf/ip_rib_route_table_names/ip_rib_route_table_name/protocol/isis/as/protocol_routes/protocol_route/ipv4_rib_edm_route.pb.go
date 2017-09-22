// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ipv4_rib_edm_route.proto

/*
Package cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_isis_as_protocol_routes_protocol_route is a generated protocol buffer package.

It is generated from these files:
	ipv4_rib_edm_route.proto

It has these top-level messages:
	Ipv4RibEdmRoute_KEYS
	Ipv4RibEdmRoute
	Ipv4RibEdmPath
	Ipv4RibEdmPathItem
	RibEdmLocalLabel
*/
package cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_isis_as_protocol_routes_protocol_route

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
type Ipv4RibEdmRoute_KEYS struct {
	VrfName        string `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
	AfName         string `protobuf:"bytes,2,opt,name=af_name,json=afName" json:"af_name,omitempty"`
	SafName        string `protobuf:"bytes,3,opt,name=saf_name,json=safName" json:"saf_name,omitempty"`
	RouteTableName string `protobuf:"bytes,4,opt,name=route_table_name,json=routeTableName" json:"route_table_name,omitempty"`
	As             string `protobuf:"bytes,5,opt,name=as" json:"as,omitempty"`
	Address        string `protobuf:"bytes,6,opt,name=address" json:"address,omitempty"`
	PrefixLength   uint32 `protobuf:"varint,7,opt,name=prefix_length,json=prefixLength" json:"prefix_length,omitempty"`
}

func (m *Ipv4RibEdmRoute_KEYS) Reset()                    { *m = Ipv4RibEdmRoute_KEYS{} }
func (m *Ipv4RibEdmRoute_KEYS) String() string            { return proto.CompactTextString(m) }
func (*Ipv4RibEdmRoute_KEYS) ProtoMessage()               {}
func (*Ipv4RibEdmRoute_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Ipv4RibEdmRoute_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ipv4RibEdmRoute_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *Ipv4RibEdmRoute_KEYS) GetSafName() string {
	if m != nil {
		return m.SafName
	}
	return ""
}

func (m *Ipv4RibEdmRoute_KEYS) GetRouteTableName() string {
	if m != nil {
		return m.RouteTableName
	}
	return ""
}

func (m *Ipv4RibEdmRoute_KEYS) GetAs() string {
	if m != nil {
		return m.As
	}
	return ""
}

func (m *Ipv4RibEdmRoute_KEYS) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Ipv4RibEdmRoute_KEYS) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

type Ipv4RibEdmRoute struct {
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
	RoutePath *Ipv4RibEdmPath `protobuf:"bytes,82,opt,name=route_path,json=routePath" json:"route_path,omitempty"`
}

func (m *Ipv4RibEdmRoute) Reset()                    { *m = Ipv4RibEdmRoute{} }
func (m *Ipv4RibEdmRoute) String() string            { return proto.CompactTextString(m) }
func (*Ipv4RibEdmRoute) ProtoMessage()               {}
func (*Ipv4RibEdmRoute) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Ipv4RibEdmRoute) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *Ipv4RibEdmRoute) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetRouteVersion() uint32 {
	if m != nil {
		return m.RouteVersion
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetProtocolId() uint32 {
	if m != nil {
		return m.ProtocolId
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetProtocolName() string {
	if m != nil {
		return m.ProtocolName
	}
	return ""
}

func (m *Ipv4RibEdmRoute) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

func (m *Ipv4RibEdmRoute) GetClientId() uint32 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetRouteType() uint32 {
	if m != nil {
		return m.RouteType
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetSvdType() uint32 {
	if m != nil {
		return m.SvdType
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetExtendedFlags() uint64 {
	if m != nil {
		return m.ExtendedFlags
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetTag() uint32 {
	if m != nil {
		return m.Tag
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetDistance() uint32 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetDiversionDistance() uint32 {
	if m != nil {
		return m.DiversionDistance
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetMetric() uint32 {
	if m != nil {
		return m.Metric
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetPathsCount() uint32 {
	if m != nil {
		return m.PathsCount
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetAttributeIdentity() uint32 {
	if m != nil {
		return m.AttributeIdentity
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetTrafficIndex() uint32 {
	if m != nil {
		return m.TrafficIndex
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetRoutePrecedence() uint32 {
	if m != nil {
		return m.RoutePrecedence
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetQosGroup() uint32 {
	if m != nil {
		return m.QosGroup
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetFlowTag() uint32 {
	if m != nil {
		return m.FlowTag
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetFwdClass() uint32 {
	if m != nil {
		return m.FwdClass
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetPicCount() uint32 {
	if m != nil {
		return m.PicCount
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *Ipv4RibEdmRoute) GetDiversion() bool {
	if m != nil {
		return m.Diversion
	}
	return false
}

func (m *Ipv4RibEdmRoute) GetDiversionProtoName() string {
	if m != nil {
		return m.DiversionProtoName
	}
	return ""
}

func (m *Ipv4RibEdmRoute) GetRouteAge() uint32 {
	if m != nil {
		return m.RouteAge
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetRouteLabel() uint32 {
	if m != nil {
		return m.RouteLabel
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetTblVersion() uint64 {
	if m != nil {
		return m.TblVersion
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetRouteModifyTime() uint64 {
	if m != nil {
		return m.RouteModifyTime
	}
	return 0
}

func (m *Ipv4RibEdmRoute) GetRoutePath() *Ipv4RibEdmPath {
	if m != nil {
		return m.RoutePath
	}
	return nil
}

// Information of a rib path
type Ipv4RibEdmPath struct {
	// Next path
	Ipv4RibEdmPath []*Ipv4RibEdmPathItem `protobuf:"bytes,1,rep,name=ipv4_rib_edm_path,json=ipv4RibEdmPath" json:"ipv4_rib_edm_path,omitempty"`
}

func (m *Ipv4RibEdmPath) Reset()                    { *m = Ipv4RibEdmPath{} }
func (m *Ipv4RibEdmPath) String() string            { return proto.CompactTextString(m) }
func (*Ipv4RibEdmPath) ProtoMessage()               {}
func (*Ipv4RibEdmPath) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Ipv4RibEdmPath) GetIpv4RibEdmPath() []*Ipv4RibEdmPathItem {
	if m != nil {
		return m.Ipv4RibEdmPath
	}
	return nil
}

type Ipv4RibEdmPathItem struct {
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
	RemoteBackupAddr [][]byte `protobuf:"bytes,31,rep,name=remote_backup_addr,json=remoteBackupAddr,proto3" json:"remote_backup_addr,omitempty"`
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

func (m *Ipv4RibEdmPathItem) Reset()                    { *m = Ipv4RibEdmPathItem{} }
func (m *Ipv4RibEdmPathItem) String() string            { return proto.CompactTextString(m) }
func (*Ipv4RibEdmPathItem) ProtoMessage()               {}
func (*Ipv4RibEdmPathItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Ipv4RibEdmPathItem) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetInformationSource() string {
	if m != nil {
		return m.InformationSource
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetV6Nexthop() string {
	if m != nil {
		return m.V6Nexthop
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetMetric() uint32 {
	if m != nil {
		return m.Metric
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetLoadMetric() uint32 {
	if m != nil {
		return m.LoadMetric
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetFlags64() uint64 {
	if m != nil {
		return m.Flags64
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetPrivateFlags() uint32 {
	if m != nil {
		return m.PrivateFlags
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetLooped() bool {
	if m != nil {
		return m.Looped
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetNextHopTableId() uint32 {
	if m != nil {
		return m.NextHopTableId
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetNextHopVrfName() string {
	if m != nil {
		return m.NextHopVrfName
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetNextHopTableName() string {
	if m != nil {
		return m.NextHopTableName
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetNextHopAfi() uint32 {
	if m != nil {
		return m.NextHopAfi
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetNextHopSafi() uint32 {
	if m != nil {
		return m.NextHopSafi
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetRouteLabel() uint32 {
	if m != nil {
		return m.RouteLabel
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetTunnelId() uint32 {
	if m != nil {
		return m.TunnelId
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetPathid() uint32 {
	if m != nil {
		return m.Pathid
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetBackupPathid() uint32 {
	if m != nil {
		return m.BackupPathid
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetRefCntOfBackup() uint32 {
	if m != nil {
		return m.RefCntOfBackup
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetNumberOfExtendedCommunities() uint32 {
	if m != nil {
		return m.NumberOfExtendedCommunities
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetMvpnPresent() bool {
	if m != nil {
		return m.MvpnPresent
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetPathrtPresent() bool {
	if m != nil {
		return m.PathrtPresent
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetVrfimportrtPresent() bool {
	if m != nil {
		return m.VrfimportrtPresent
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetSourceasrtPresent() bool {
	if m != nil {
		return m.SourceasrtPresent
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetSourcerdPresent() bool {
	if m != nil {
		return m.SourcerdPresent
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetSegmentedNexthopPresent() bool {
	if m != nil {
		return m.SegmentedNexthopPresent
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetNextHopId() uint32 {
	if m != nil {
		return m.NextHopId
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetNextHopIdRefcount() uint32 {
	if m != nil {
		return m.NextHopIdRefcount
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetOspfAreaId() string {
	if m != nil {
		return m.OspfAreaId
	}
	return ""
}

func (m *Ipv4RibEdmPathItem) GetRemoteBackupAddr() [][]byte {
	if m != nil {
		return m.RemoteBackupAddr
	}
	return nil
}

func (m *Ipv4RibEdmPathItem) GetHasLabelstk() bool {
	if m != nil {
		return m.HasLabelstk
	}
	return false
}

func (m *Ipv4RibEdmPathItem) GetNumLabels() uint32 {
	if m != nil {
		return m.NumLabels
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetLabelstk() []uint32 {
	if m != nil {
		return m.Labelstk
	}
	return nil
}

func (m *Ipv4RibEdmPathItem) GetBindingLabel() uint32 {
	if m != nil {
		return m.BindingLabel
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetNhidFeid() uint64 {
	if m != nil {
		return m.NhidFeid
	}
	return 0
}

func (m *Ipv4RibEdmPathItem) GetMplsFeid() uint64 {
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
func (*RibEdmLocalLabel) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

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
	proto.RegisterType((*Ipv4RibEdmRoute_KEYS)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.isis.as.protocol_routes.protocol_route.ipv4_rib_edm_route_KEYS")
	proto.RegisterType((*Ipv4RibEdmRoute)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.isis.as.protocol_routes.protocol_route.ipv4_rib_edm_route")
	proto.RegisterType((*Ipv4RibEdmPath)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.isis.as.protocol_routes.protocol_route.ipv4_rib_edm_path")
	proto.RegisterType((*Ipv4RibEdmPathItem)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.isis.as.protocol_routes.protocol_route.ipv4_rib_edm_path_item")
	proto.RegisterType((*RibEdmLocalLabel)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.isis.as.protocol_routes.protocol_route.rib_edm_local_label")
}

func init() { proto.RegisterFile("ipv4_rib_edm_route.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1499 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x57, 0xdd, 0x52, 0x5b, 0x37,
	0x10, 0x1e, 0x43, 0x30, 0xb6, 0xc0, 0x04, 0x1f, 0x28, 0x88, 0x90, 0x1f, 0x87, 0x34, 0x33, 0xa6,
	0x53, 0x9c, 0x4e, 0x92, 0xd2, 0x36, 0xfd, 0x25, 0x84, 0x24, 0x6e, 0x48, 0xa0, 0x4e, 0x26, 0x33,
	0xbd, 0xd2, 0xc8, 0x47, 0x3a, 0xb6, 0x26, 0xe7, 0x2f, 0x92, 0xec, 0xc0, 0x65, 0x9f, 0xa3, 0xef,
	0xd0, 0x9b, 0xbe, 0x47, 0x5f, 0xa1, 0xd3, 0x07, 0xe8, 0x4d, 0x67, 0x7a, 0xdf, 0xd1, 0xae, 0xce,
	0xb1, 0x81, 0xf4, 0xbe, 0xbd, 0x49, 0xd8, 0xef, 0xfb, 0xa4, 0x23, 0xed, 0xae, 0x76, 0xd7, 0x84,
	0xaa, 0x7c, 0x7c, 0x9f, 0x69, 0xd5, 0x67, 0x52, 0x24, 0x4c, 0x67, 0x23, 0x2b, 0x3b, 0xb9, 0xce,
	0x6c, 0x16, 0xfc, 0x5c, 0x09, 0x95, 0x09, 0x33, 0xa6, 0x32, 0xc3, 0x4e, 0x34, 0x53, 0x39, 0xa8,
	0x40, 0x9e, 0xe5, 0x52, 0x77, 0xb4, 0xea, 0x77, 0xc6, 0x3a, 0x32, 0xee, 0x9f, 0x0e, 0x8f, 0x4c,
	0x87, 0x47, 0x1d, 0xe3, 0xfe, 0x37, 0x3c, 0xea, 0x78, 0x35, 0xec, 0xc7, 0x2c, 0xef, 0xc7, 0x92,
	0xa5, 0x3c, 0x91, 0xe6, 0xdf, 0x08, 0xfc, 0x66, 0x98, 0xc5, 0x1d, 0x65, 0x94, 0xe9, 0x70, 0x53,
	0x02, 0x28, 0x3d, 0x6f, 0x6f, 0xfd, 0x51, 0x21, 0xeb, 0x17, 0x8f, 0xce, 0x9e, 0x1d, 0xfc, 0xf8,
	0x32, 0xd8, 0x20, 0xb5, 0xb1, 0x8e, 0x60, 0x5f, 0x5a, 0x69, 0x55, 0xda, 0xf5, 0xde, 0xfc, 0x58,
	0x47, 0x2f, 0x78, 0x22, 0x83, 0x75, 0x32, 0xcf, 0x3d, 0x33, 0x03, 0x4c, 0x95, 0x23, 0xb1, 0x41,
	0x6a, 0xa6, 0x60, 0x66, 0x71, 0x8d, 0xf1, 0x54, 0x9b, 0x2c, 0x9f, 0x3f, 0x2e, 0xbd, 0x04, 0x92,
	0x25, 0xc0, 0x5f, 0x39, 0x18, 0x94, 0x4b, 0x64, 0x86, 0x1b, 0x3a, 0x07, 0xdc, 0x0c, 0x37, 0x01,
	0x25, 0xf3, 0x5c, 0x08, 0x2d, 0x8d, 0xa1, 0x55, 0xdc, 0xd3, 0x9b, 0xc1, 0x2d, 0xd2, 0xc8, 0xb5,
	0x8c, 0xd4, 0x09, 0x8b, 0x65, 0x3a, 0xb0, 0x43, 0x3a, 0xdf, 0xaa, 0xb4, 0x1b, 0xbd, 0x45, 0x04,
	0x0f, 0x01, 0xdb, 0xfa, 0xb3, 0x4e, 0x82, 0x8b, 0x77, 0x0c, 0xd6, 0x48, 0x15, 0x65, 0xf4, 0x2e,
	0x5e, 0x01, 0xad, 0x8b, 0x7b, 0xde, 0xbb, 0xb8, 0xa7, 0x13, 0xe1, 0x65, 0xc6, 0x52, 0x1b, 0x95,
	0xa5, 0xf4, 0x3e, 0x8a, 0x00, 0x7c, 0x8d, 0x58, 0x70, 0x83, 0x2c, 0x94, 0xee, 0x56, 0x82, 0x7e,
	0x0a, 0x12, 0x52, 0x40, 0x5d, 0x81, 0x9f, 0xf2, 0x02, 0xf0, 0xc7, 0x2e, 0x9c, 0x64, 0xb1, 0x00,
	0xc1, 0x1b, 0x57, 0x48, 0x4d, 0xa5, 0xc6, 0xf2, 0x34, 0x94, 0xf4, 0x33, 0xe0, 0x4b, 0x3b, 0xd8,
	0x24, 0xf5, 0x30, 0x56, 0x32, 0xb5, 0x6e, 0xff, 0xcf, 0x61, 0xff, 0x1a, 0x02, 0x5d, 0x11, 0x5c,
	0x23, 0xc4, 0x3b, 0xfc, 0x34, 0x97, 0xf4, 0x0b, 0x60, 0xeb, 0xe8, 0xea, 0xd3, 0x1c, 0xf6, 0xcd,
	0xb5, 0xca, 0xb4, 0xb2, 0xa7, 0xf4, 0x01, 0x2e, 0x2d, 0x6c, 0x08, 0xe3, 0x58, 0xe0, 0xc2, 0x2f,
	0x81, 0x9b, 0x37, 0x63, 0x01, 0xcb, 0x56, 0xc9, 0x5c, 0x14, 0xf3, 0x81, 0xa1, 0x5f, 0x01, 0x8e,
	0x46, 0x70, 0x9b, 0x2c, 0xc9, 0x13, 0x2b, 0x53, 0x21, 0x05, 0x43, 0xfa, 0xeb, 0x56, 0xa5, 0x7d,
	0xa9, 0xd7, 0x28, 0xd0, 0xc7, 0x20, 0x5b, 0x26, 0xb3, 0x96, 0x0f, 0xe8, 0x37, 0xb0, 0xd4, 0xfd,
	0xe9, 0x4e, 0x21, 0x94, 0xbf, 0xdd, 0xb7, 0x78, 0x8a, 0xc2, 0x0e, 0x76, 0x48, 0x20, 0x94, 0x77,
	0x30, 0x2b, 0x55, 0xdf, 0x81, 0xaa, 0x59, 0x32, 0x8f, 0x0a, 0xf9, 0x1a, 0xa9, 0x26, 0xd2, 0x6a,
	0x15, 0xd2, 0x3d, 0x90, 0x78, 0x0b, 0xc2, 0xc0, 0xed, 0xd0, 0xb0, 0x30, 0x1b, 0xa5, 0x96, 0x3e,
	0xf4, 0x61, 0x70, 0xd0, 0xbe, 0x43, 0xdc, 0x77, 0xb8, 0xb5, 0x5a, 0xf5, 0x9d, 0xb3, 0x94, 0x90,
	0xa9, 0x75, 0x3e, 0xd9, 0xc7, 0xef, 0x94, 0x4c, 0xd7, 0x13, 0x2e, 0x6a, 0x56, 0xf3, 0x28, 0x52,
	0x21, 0x53, 0xa9, 0x90, 0x27, 0xf4, 0x11, 0xc6, 0xde, 0x83, 0x5d, 0x87, 0x05, 0xdb, 0x45, 0xb6,
	0xe7, 0x5a, 0x86, 0x52, 0x48, 0x77, 0xf2, 0x03, 0xd0, 0x5d, 0x06, 0xfc, 0xb8, 0x84, 0x5d, 0x10,
	0xdf, 0x66, 0x86, 0x0d, 0x74, 0x36, 0xca, 0xe9, 0x63, 0xf4, 0xc1, 0xdb, 0xcc, 0x3c, 0x71, 0xb6,
	0x8b, 0x44, 0x14, 0x67, 0xef, 0x98, 0x73, 0xdb, 0x13, 0x8c, 0x84, 0xb3, 0x5f, 0xf1, 0x81, 0x5b,
	0x17, 0xbd, 0x13, 0x2c, 0x8c, 0xb9, 0x31, 0xf4, 0x29, 0xae, 0x8b, 0xde, 0x89, 0x7d, 0x67, 0x3b,
	0x32, 0x57, 0xa1, 0xbf, 0x72, 0xd7, 0x87, 0x57, 0x85, 0x78, 0xe1, 0x35, 0x52, 0xe5, 0xa1, 0x55,
	0x63, 0x49, 0xbf, 0x6f, 0x55, 0xda, 0xb5, 0x9e, 0xb7, 0x82, 0xab, 0xa4, 0x5e, 0xba, 0x95, 0x3e,
	0x03, 0x6a, 0x02, 0x04, 0x9f, 0x90, 0xd5, 0x49, 0x38, 0x20, 0x45, 0x31, 0x69, 0x0f, 0x21, 0x29,
	0x27, 0xa1, 0x3a, 0x76, 0x14, 0xa4, 0xee, 0x26, 0xc1, 0x7c, 0x63, 0x7c, 0x20, 0xe9, 0x73, 0x3c,
	0x04, 0x00, 0x7b, 0x03, 0xe9, 0xc2, 0x82, 0x64, 0xcc, 0xfb, 0x32, 0xa6, 0x2f, 0x30, 0x2c, 0x00,
	0x1d, 0x3a, 0xc4, 0x3d, 0xfb, 0xe2, 0x2c, 0x47, 0x78, 0xf3, 0xf1, 0xe4, 0x61, 0xd9, 0x7e, 0x5c,
	0xbe, 0xbd, 0x63, 0x48, 0x35, 0x62, 0xfb, 0x71, 0xf1, 0xf2, 0x3e, 0x22, 0x4d, 0xdc, 0x3b, 0xc9,
	0x84, 0x8a, 0x4e, 0x99, 0x55, 0x89, 0xa4, 0x3f, 0x80, 0x0c, 0xdd, 0xff, 0x1c, 0xf0, 0x57, 0x2a,
	0x91, 0xc1, 0x6f, 0x95, 0xe2, 0x9d, 0xb8, 0x94, 0xa0, 0xbd, 0x56, 0xa5, 0xbd, 0x70, 0xf7, 0x97,
	0x4a, 0xe7, 0x3f, 0x5c, 0xb6, 0x3b, 0x67, 0xca, 0x99, 0x3b, 0xb6, 0x7f, 0xd8, 0xc7, 0xdc, 0x0e,
	0xb7, 0xfe, 0xae, 0x90, 0xe6, 0x05, 0x41, 0xf0, 0xfb, 0xfb, 0x50, 0x5a, 0x69, 0xcd, 0xb6, 0x17,
	0xee, 0xfe, 0xfa, 0x3f, 0xbb, 0x2d, 0x53, 0x56, 0x26, 0xbd, 0x25, 0x87, 0xf7, 0x54, 0xff, 0x40,
	0x24, 0x70, 0xef, 0xbf, 0x08, 0x59, 0x7b, 0xbf, 0x74, 0xba, 0x83, 0x54, 0xce, 0x76, 0x90, 0x1d,
	0x12, 0xa8, 0x34, 0xca, 0x74, 0xc2, 0xad, 0x4b, 0x6b, 0x93, 0x8d, 0x74, 0x58, 0x34, 0xb5, 0xe6,
	0x14, 0xf3, 0x12, 0x08, 0x57, 0x53, 0xc7, 0xbb, 0x2c, 0x95, 0x27, 0x76, 0x98, 0xe5, 0xbe, 0xc3,
	0xd5, 0xc7, 0xbb, 0x2f, 0x10, 0x70, 0x65, 0x50, 0xa5, 0x56, 0xea, 0x88, 0x87, 0x67, 0x3a, 0x5c,
	0xa3, 0x44, 0xe1, 0x5d, 0x4c, 0x2a, 0xd5, 0xdc, 0xf9, 0x4a, 0x15, 0x67, 0x5c, 0x30, 0x4f, 0x56,
	0xf1, 0x49, 0x38, 0xe8, 0x39, 0x0a, 0x28, 0x99, 0x87, 0xea, 0xba, 0x7b, 0x1f, 0x3a, 0xdd, 0xa5,
	0x5e, 0x61, 0x4e, 0xca, 0x72, 0x6d, 0xba, 0x2c, 0x43, 0x83, 0x51, 0x63, 0x6e, 0xa5, 0xaf, 0xca,
	0xf5, 0xa2, 0x97, 0x01, 0x88, 0x45, 0x79, 0x8d, 0x54, 0xe3, 0x2c, 0xcb, 0xa5, 0xa0, 0x04, 0xab,
	0x01, 0x5a, 0xc1, 0x36, 0x69, 0xba, 0x8b, 0xb2, 0x61, 0x96, 0xfb, 0xe8, 0x29, 0x41, 0x17, 0x60,
	0x83, 0x25, 0x47, 0x3c, 0xcd, 0x72, 0xe8, 0xd9, 0xdd, 0xb3, 0xd2, 0x72, 0x66, 0x58, 0xc4, 0xe6,
	0xee, 0xa5, 0xaf, 0xfd, 0xe8, 0xb0, 0x43, 0x56, 0xce, 0xed, 0x0a, 0xe2, 0x06, 0x88, 0x97, 0xa7,
	0xf7, 0x05, 0x79, 0x8b, 0x2c, 0x96, 0x72, 0x1e, 0x29, 0xba, 0x84, 0x3e, 0xf1, 0xba, 0xbd, 0x48,
	0x05, 0x5b, 0xa4, 0x51, 0x2a, 0x8c, 0x93, 0x5c, 0x06, 0xc9, 0x82, 0x97, 0xbc, 0xe4, 0x91, 0x3a,
	0x5f, 0x6b, 0x96, 0x2f, 0xd4, 0x9a, 0x4d, 0x52, 0xb7, 0xa3, 0x34, 0x95, 0xd0, 0xa8, 0x9b, 0x58,
	0xa9, 0x10, 0xe8, 0x0a, 0x98, 0x14, 0xb8, 0x1d, 0x2a, 0x41, 0x03, 0x0c, 0x17, 0x5a, 0xce, 0xbb,
	0x7d, 0x1e, 0xbe, 0x19, 0xe5, 0xcc, 0xd3, 0x2b, 0xe8, 0x5d, 0x04, 0x8f, 0x51, 0xb4, 0x4d, 0x9a,
	0x5a, 0x46, 0x2c, 0x4c, 0x2d, 0xcb, 0x22, 0x86, 0x14, 0x5d, 0x45, 0x2f, 0x6a, 0x19, 0xed, 0xa7,
	0xf6, 0x28, 0x7a, 0x08, 0x68, 0xb0, 0x4f, 0xae, 0xa7, 0xa3, 0xa4, 0x2f, 0xb5, 0x53, 0x96, 0xed,
	0x34, 0xcc, 0x92, 0x64, 0x94, 0x2a, 0xab, 0xa4, 0xa1, 0x1f, 0xc0, 0xba, 0x4d, 0x54, 0x1d, 0x45,
	0x07, 0x5e, 0xb3, 0x3f, 0x91, 0x04, 0x37, 0xc9, 0x62, 0x32, 0xce, 0x5d, 0x81, 0x96, 0x46, 0xa6,
	0x96, 0xae, 0x41, 0x4c, 0x17, 0x1c, 0x76, 0x8c, 0x90, 0xcb, 0x52, 0x77, 0x60, 0x6d, 0x4b, 0xd1,
	0x3a, 0x88, 0x1a, 0x88, 0x16, 0xb2, 0x3b, 0x64, 0x65, 0xac, 0x23, 0x95, 0xe4, 0x99, 0xb6, 0x53,
	0x5a, 0x0a, 0xda, 0x60, 0x8a, 0x2a, 0x16, 0xec, 0x90, 0x00, 0xdf, 0x0f, 0x37, 0x53, 0xfa, 0x0d,
	0xd0, 0x37, 0x27, 0x4c, 0x21, 0xdf, 0x26, 0xcb, 0x08, 0x6a, 0x51, 0x8a, 0xaf, 0x80, 0xf8, 0x72,
	0x81, 0x17, 0xd2, 0x07, 0x64, 0xc3, 0xc8, 0x41, 0x22, 0x53, 0x2b, 0x45, 0xf1, 0xfa, 0xca, 0x35,
	0x9b, 0xb0, 0x66, 0xbd, 0x14, 0xf8, 0xc7, 0x58, 0xac, 0xbd, 0x4e, 0x16, 0xca, 0xfc, 0x50, 0x82,
	0x5e, 0xc5, 0x39, 0xc8, 0x67, 0x47, 0x57, 0x04, 0x77, 0xc8, 0xea, 0x14, 0xcf, 0xb4, 0x8c, 0xb0,
	0x69, 0x5e, 0xc3, 0xfe, 0x5f, 0x0a, 0x7b, 0x9e, 0x70, 0x29, 0x99, 0x99, 0x3c, 0x62, 0x5c, 0x4b,
	0xee, 0x76, 0xbc, 0x0e, 0xa9, 0x4b, 0x1c, 0xb6, 0xa7, 0x25, 0xef, 0x8a, 0xe0, 0x63, 0x12, 0x68,
	0x99, 0x64, 0x56, 0xfa, 0x78, 0x33, 0x57, 0x6d, 0xe8, 0x8d, 0xd6, 0x6c, 0x7b, 0xb1, 0xb7, 0x8c,
	0x0c, 0x86, 0x7c, 0x4f, 0x08, 0xed, 0x22, 0x36, 0xe4, 0x06, 0x53, 0xd3, 0xd8, 0x37, 0xb4, 0x85,
	0x11, 0x1b, 0x72, 0x73, 0xe8, 0x21, 0x57, 0x76, 0xd2, 0x51, 0xe2, 0x25, 0xf4, 0xa6, 0xbf, 0xc2,
	0x28, 0x41, 0x81, 0x1b, 0xa2, 0xca, 0xd5, 0x5b, 0xad, 0x59, 0x97, 0xbc, 0x85, 0x0d, 0x49, 0xaa,
	0x52, 0xa1, 0xd2, 0x81, 0x4f, 0xfe, 0x5b, 0x3e, 0x49, 0x11, 0x2c, 0xd3, 0x3f, 0x1d, 0x2a, 0xc1,
	0x22, 0xa9, 0x04, 0xfd, 0x10, 0x2a, 0x4b, 0xcd, 0x01, 0x8f, 0xa5, 0x12, 0x8e, 0x4c, 0xf2, 0xd8,
	0x20, 0x79, 0x1b, 0x49, 0x07, 0x38, 0x72, 0xeb, 0xa7, 0x19, 0xb2, 0x52, 0xd4, 0xdb, 0x38, 0x0b,
	0x79, 0x8c, 0x5f, 0xb9, 0x38, 0xda, 0x56, 0xde, 0x33, 0xda, 0x9e, 0x19, 0x5f, 0x67, 0xce, 0x8d,
	0xaf, 0xab, 0x64, 0xce, 0x58, 0x1e, 0xe3, 0xef, 0x88, 0x46, 0x0f, 0x0d, 0x77, 0xd5, 0x44, 0x69,
	0x9d, 0x69, 0x29, 0xa0, 0xb6, 0x36, 0x7a, 0xa5, 0xed, 0xbe, 0x99, 0x48, 0x3d, 0x90, 0x6e, 0x56,
	0x74, 0x05, 0xc4, 0x57, 0xd7, 0x45, 0x00, 0x1f, 0x21, 0x06, 0xa5, 0x40, 0xba, 0x61, 0x92, 0x65,
	0x69, 0x7c, 0x5a, 0xd4, 0x58, 0x84, 0x8e, 0xd2, 0xf8, 0xd4, 0x7d, 0x17, 0x1d, 0x85, 0xbf, 0x25,
	0xd0, 0x38, 0x33, 0xa7, 0xd6, 0xce, 0xce, 0xa9, 0xfd, 0x2a, 0x5c, 0xea, 0xde, 0x3f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x60, 0x1c, 0xc4, 0x03, 0x05, 0x0e, 0x00, 0x00,
}
