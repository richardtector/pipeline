// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ipv4_rib_edm_route.proto

/*
Package cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_rib_table_ids_rib_table_id_rib_table_itf_hndls_rib_table_itf_hndl_itf_route is a generated protocol buffer package.

It is generated from these files:
	ipv4_rib_edm_route.proto

It has these top-level messages:
	Ipv4RibEdmRoute_KEYS
	Ipv4RibEdmRoute
	Ipv4RibEdmPath
	Ipv4RibEdmPathItem
	RibEdmLocalLabel
*/
package cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_rib_table_ids_rib_table_id_rib_table_itf_hndls_rib_table_itf_hndl_itf_route

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
	Tableid uint32 `protobuf:"varint,1,opt,name=tableid" json:"tableid,omitempty"`
	Handle  uint32 `protobuf:"varint,2,opt,name=handle" json:"handle,omitempty"`
	Address string `protobuf:"bytes,3,opt,name=address" json:"address,omitempty"`
}

func (m *Ipv4RibEdmRoute_KEYS) Reset()                    { *m = Ipv4RibEdmRoute_KEYS{} }
func (m *Ipv4RibEdmRoute_KEYS) String() string            { return proto.CompactTextString(m) }
func (*Ipv4RibEdmRoute_KEYS) ProtoMessage()               {}
func (*Ipv4RibEdmRoute_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Ipv4RibEdmRoute_KEYS) GetTableid() uint32 {
	if m != nil {
		return m.Tableid
	}
	return 0
}

func (m *Ipv4RibEdmRoute_KEYS) GetHandle() uint32 {
	if m != nil {
		return m.Handle
	}
	return 0
}

func (m *Ipv4RibEdmRoute_KEYS) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
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
	proto.RegisterType((*Ipv4RibEdmRoute_KEYS)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.rib_table_ids.rib_table_id.rib_table_itf_hndls.rib_table_itf_hndl.itf_route.ipv4_rib_edm_route_KEYS")
	proto.RegisterType((*Ipv4RibEdmRoute)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.rib_table_ids.rib_table_id.rib_table_itf_hndls.rib_table_itf_hndl.itf_route.ipv4_rib_edm_route")
	proto.RegisterType((*Ipv4RibEdmPath)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.rib_table_ids.rib_table_id.rib_table_itf_hndls.rib_table_itf_hndl.itf_route.ipv4_rib_edm_path")
	proto.RegisterType((*Ipv4RibEdmPathItem)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.rib_table_ids.rib_table_id.rib_table_itf_hndls.rib_table_itf_hndl.itf_route.ipv4_rib_edm_path_item")
	proto.RegisterType((*RibEdmLocalLabel)(nil), "cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.rib_table_ids.rib_table_id.rib_table_itf_hndls.rib_table_itf_hndl.itf_route.rib_edm_local_label")
}

func init() { proto.RegisterFile("ipv4_rib_edm_route.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0xdd, 0x72, 0x14, 0x37,
	0x16, 0xae, 0xc1, 0x60, 0x7b, 0x64, 0x8f, 0xf1, 0xb4, 0xbd, 0x46, 0x60, 0x7e, 0x06, 0xb3, 0x54,
	0xd9, 0x5b, 0x8b, 0xd9, 0x02, 0xd6, 0xbb, 0xcb, 0xfe, 0xc5, 0x18, 0x03, 0x13, 0x0c, 0x76, 0x06,
	0x17, 0x55, 0xb9, 0x52, 0x69, 0x5a, 0xea, 0x19, 0x15, 0xdd, 0xad, 0x46, 0xd2, 0x0c, 0xf6, 0x65,
	0x6e, 0x72, 0x93, 0xf7, 0xc8, 0x2b, 0xe4, 0x09, 0x72, 0x99, 0x97, 0xc9, 0x13, 0xa4, 0xce, 0x39,
	0xdd, 0x3d, 0xe3, 0x9f, 0xfb, 0xe4, 0xae, 0xcf, 0x77, 0xbe, 0x23, 0xe9, 0xfc, 0xe8, 0x6b, 0x31,
	0x6e, 0x8a, 0xf1, 0x33, 0xe1, 0x4c, 0x5f, 0x68, 0x95, 0x09, 0x67, 0x47, 0x41, 0x6f, 0x17, 0xce,
	0x06, 0x1b, 0xf9, 0xd8, 0xf8, 0xd8, 0x0a, 0x63, 0xbd, 0x38, 0x71, 0xc2, 0x14, 0x48, 0x42, 0xb6,
	0x2d, 0xb4, 0xdb, 0x06, 0xcb, 0x07, 0xd5, 0x3f, 0xc5, 0xaf, 0x20, 0xfb, 0xa9, 0x16, 0x46, 0xf9,
	0x33, 0xd6, 0xb4, 0x11, 0x12, 0x31, 0xcc, 0x55, 0xea, 0x2f, 0xc1, 0xb6, 0xe1, 0x03, 0xb7, 0xde,
	0xd0, 0xec, 0xc6, 0xc5, 0x03, 0x89, 0xb7, 0xfb, 0xdf, 0x7e, 0x88, 0x38, 0x9b, 0xc3, 0x30, 0xa3,
	0x78, 0xa3, 0xd3, 0xd8, 0x6c, 0xf5, 0x2a, 0x33, 0x5a, 0x63, 0xb3, 0x43, 0x99, 0xab, 0x54, 0xf3,
	0x2b, 0xe8, 0x28, 0x2d, 0x88, 0x90, 0x4a, 0x39, 0xed, 0x3d, 0x9f, 0xe9, 0x34, 0x36, 0x9b, 0xbd,
	0xca, 0xdc, 0xf8, 0xb9, 0xc9, 0xa2, 0x8b, 0xfb, 0xc0, 0x42, 0x85, 0xd3, 0x89, 0x39, 0xe1, 0x4f,
	0x90, 0x5f, 0x5a, 0xd1, 0x03, 0xd6, 0xa2, 0x2f, 0x91, 0xea, 0x7c, 0x10, 0x86, 0xfc, 0x29, 0xee,
	0xb3, 0x48, 0xe0, 0x01, 0x62, 0x40, 0xa2, 0xd3, 0x8e, 0xb5, 0xf3, 0xc6, 0xe6, 0xfc, 0x19, 0x91,
	0x10, 0xfc, 0x48, 0x58, 0x74, 0x8f, 0x2d, 0x60, 0x75, 0x63, 0x9b, 0x0a, 0xa3, 0xf8, 0xdf, 0x91,
	0xc2, 0x2a, 0xa8, 0xab, 0x68, 0xab, 0x92, 0x90, 0xcb, 0x4c, 0xf3, 0x1d, 0x3c, 0xc9, 0x62, 0x05,
	0xbe, 0x97, 0x99, 0x8e, 0x6e, 0xb1, 0x79, 0x93, 0xfb, 0x20, 0xf3, 0x58, 0xf3, 0x7f, 0xa0, 0xbf,
	0xb6, 0xa3, 0x75, 0xd6, 0x8c, 0x53, 0xa3, 0xf3, 0x00, 0xeb, 0xff, 0x13, 0xd7, 0x9f, 0x27, 0xa0,
	0xab, 0xa2, 0x3b, 0x8c, 0xd1, 0x19, 0xc3, 0x69, 0xa1, 0xf9, 0xbf, 0xd0, 0xdb, 0x44, 0xe4, 0xf8,
	0xb4, 0xc0, 0x75, 0x0b, 0x67, 0xac, 0x33, 0xe1, 0x94, 0x3f, 0xa7, 0xd0, 0xca, 0x8e, 0x6e, 0xb2,
	0x79, 0x3f, 0x56, 0x14, 0xf8, 0x6f, 0xaa, 0xbf, 0x1f, 0x2b, 0x0c, 0x5b, 0x65, 0xd7, 0x92, 0x54,
	0x0e, 0x3c, 0xff, 0x0f, 0xe2, 0x64, 0x44, 0x0f, 0xd9, 0x92, 0x3e, 0x09, 0x3a, 0x57, 0x5a, 0x09,
	0x72, 0xff, 0xb7, 0xd3, 0xd8, 0xbc, 0xda, 0x6b, 0x55, 0xe8, 0x2b, 0xa4, 0x2d, 0xb3, 0x99, 0x20,
	0x07, 0xfc, 0x7f, 0x18, 0x0a, 0x9f, 0x70, 0x0a, 0x65, 0xca, 0xec, 0xfe, 0x4f, 0xa7, 0xa8, 0xec,
	0xe8, 0x11, 0x8b, 0x94, 0x29, 0x0b, 0x2c, 0x6a, 0xd6, 0x57, 0xc8, 0x6a, 0xd7, 0x9e, 0x97, 0x15,
	0x7d, 0x8d, 0xcd, 0x66, 0x3a, 0x38, 0x13, 0xf3, 0x5d, 0x9a, 0x0c, 0xb2, 0xb0, 0x0d, 0x32, 0x0c,
	0xbd, 0x88, 0xed, 0x28, 0x0f, 0xfc, 0x45, 0xd9, 0x06, 0x80, 0xf6, 0x00, 0x81, 0x7d, 0x64, 0x08,
	0xce, 0xf4, 0xa1, 0x58, 0x46, 0xe9, 0x3c, 0x40, 0x4d, 0xf6, 0x68, 0x9f, 0xda, 0xd3, 0x2d, 0x1d,
	0xd0, 0xb5, 0xe0, 0x64, 0x92, 0x98, 0x58, 0x98, 0x5c, 0xe9, 0x13, 0xfe, 0x92, 0x7a, 0x5f, 0x82,
	0x5d, 0xc0, 0xa2, 0x2d, 0xb6, 0x4c, 0xc5, 0x2f, 0x9c, 0x8e, 0xb5, 0xd2, 0x70, 0xf2, 0x7d, 0xe4,
	0x5d, 0x47, 0xfc, 0xa8, 0x86, 0xa1, 0x89, 0x9f, 0xad, 0x17, 0x03, 0x67, 0x47, 0x05, 0x7f, 0x45,
	0x35, 0xf8, 0x6c, 0xfd, 0x6b, 0xb0, 0xa1, 0x13, 0x49, 0x6a, 0xbf, 0x08, 0x28, 0xdb, 0x6b, 0xea,
	0x04, 0xd8, 0xc7, 0x72, 0x00, 0x71, 0xc9, 0x17, 0x25, 0xe2, 0x54, 0x7a, 0xcf, 0xdf, 0x50, 0x5c,
	0xf2, 0x45, 0xed, 0x81, 0x0d, 0xce, 0xc2, 0xc4, 0x65, 0xca, 0xdd, 0xb2, 0xbd, 0x26, 0xa6, 0x84,
	0xd7, 0xd8, 0xac, 0x8c, 0x83, 0x19, 0x6b, 0xfe, 0x75, 0xa7, 0xb1, 0x39, 0xdf, 0x2b, 0xad, 0xe8,
	0x36, 0x6b, 0xd6, 0x65, 0xe5, 0x6f, 0xd1, 0x35, 0x01, 0xa2, 0xbf, 0xb1, 0xd5, 0x49, 0x3b, 0x70,
	0x44, 0x69, 0x68, 0x0f, 0x70, 0x28, 0x27, 0xad, 0x3a, 0x02, 0x17, 0x8e, 0xee, 0x3a, 0xa3, 0x79,
	0x13, 0x72, 0xa0, 0xf9, 0x3b, 0x3a, 0x04, 0x02, 0xbb, 0x03, 0x0d, 0x6d, 0x21, 0x67, 0x2a, 0xfb,
	0x3a, 0xe5, 0xef, 0xa9, 0x2d, 0x08, 0x1d, 0x00, 0x02, 0x37, 0xba, 0x3a, 0xcb, 0x21, 0x65, 0x3e,
	0x9e, 0x5c, 0xac, 0xd0, 0x4f, 0xeb, 0xbb, 0x77, 0x84, 0xa3, 0xc6, 0x42, 0x3f, 0xad, 0x6e, 0xde,
	0x5f, 0x58, 0x9b, 0xd6, 0xce, 0xac, 0x32, 0xc9, 0xa9, 0x08, 0x26, 0xd3, 0xfc, 0x1b, 0xa4, 0x51,
	0xf9, 0xdf, 0x21, 0x7e, 0x6c, 0x32, 0x1d, 0xfd, 0xd8, 0xa8, 0xee, 0x09, 0x8c, 0x04, 0xef, 0x75,
	0x1a, 0x9b, 0x0b, 0x4f, 0xbe, 0x6f, 0x6c, 0xff, 0x0e, 0x8a, 0xb8, 0x7d, 0x46, 0xa6, 0xe0, 0x38,
	0xe5, 0x85, 0x3d, 0x92, 0x61, 0xb8, 0xf1, 0x4b, 0x83, 0xb5, 0x2f, 0x10, 0xa2, 0x9f, 0x2e, 0x43,
	0x79, 0xa3, 0x33, 0xb3, 0xb9, 0xf0, 0xe4, 0x87, 0x3f, 0x48, 0x16, 0xc2, 0x04, 0x9d, 0xf5, 0x96,
	0x00, 0xef, 0x99, 0xfe, 0xbe, 0xca, 0x30, 0x9f, 0x5f, 0x19, 0x5b, 0xbb, 0x9c, 0x3a, 0x2d, 0xe6,
	0x8d, 0x33, 0x62, 0x0e, 0x77, 0xd5, 0xe4, 0x89, 0x75, 0x99, 0x0c, 0x30, 0x86, 0xde, 0x8e, 0x5c,
	0x4c, 0xbf, 0x82, 0x66, 0xaf, 0x3d, 0xe5, 0xf9, 0x80, 0x0e, 0xd0, 0xc0, 0xf1, 0x8e, 0xc8, 0xf5,
	0x49, 0x18, 0xda, 0xa2, 0xfc, 0x31, 0x34, 0xc7, 0x3b, 0xef, 0x09, 0x00, 0xd9, 0x32, 0x79, 0xd0,
	0x2e, 0x91, 0xb1, 0xa6, 0x61, 0xbe, 0x8a, 0x94, 0x56, 0x8d, 0xe2, 0x1c, 0x4f, 0x94, 0xe5, 0xda,
	0x79, 0x65, 0x49, 0xad, 0x54, 0xa2, 0x74, 0xce, 0xd2, 0x08, 0x03, 0xf4, 0x8e, 0x08, 0x9c, 0xcd,
	0xa1, 0x1a, 0xee, 0x3c, 0xe3, 0x73, 0x38, 0x7d, 0x95, 0x39, 0x91, 0xd1, 0xf9, 0x69, 0x19, 0xc5,
	0x1f, 0x82, 0x19, 0xcb, 0xa0, 0x4b, 0x15, 0x6d, 0x56, 0xff, 0x1e, 0x04, 0x49, 0x44, 0xd7, 0xd8,
	0x6c, 0x6a, 0x6d, 0xa1, 0x15, 0x67, 0x74, 0x7b, 0xc9, 0x8a, 0xb6, 0x58, 0x1b, 0x12, 0x15, 0x43,
	0x5b, 0xd4, 0xcd, 0xe3, 0x0b, 0xb8, 0xc0, 0x12, 0x38, 0xde, 0xd8, 0xe2, 0x18, 0xe0, 0xee, 0x59,
	0xea, 0xd8, 0x25, 0x94, 0xfa, 0x22, 0xa6, 0x5e, 0x51, 0x3f, 0xba, 0x04, 0x73, 0x7f, 0xc4, 0x56,
	0xce, 0xad, 0x8a, 0xe4, 0x16, 0x92, 0x97, 0xa7, 0xd7, 0x45, 0x7a, 0x87, 0x2d, 0xd6, 0x74, 0x99,
	0x18, 0xbe, 0x44, 0x35, 0x29, 0x79, 0xbb, 0x89, 0x89, 0x36, 0x58, 0xab, 0x66, 0x78, 0xa0, 0x5c,
	0x47, 0xca, 0x42, 0x49, 0xf9, 0x20, 0x13, 0x73, 0x5e, 0x1b, 0x96, 0x2f, 0x68, 0xc3, 0x3a, 0x6b,
	0x86, 0x51, 0x9e, 0x6b, 0xfc, 0xb1, 0xb6, 0x49, 0x59, 0x08, 0xe8, 0xe2, 0x13, 0x01, 0x46, 0xc9,
	0x28, 0x1e, 0x51, 0xbb, 0xc8, 0x82, 0xea, 0xf6, 0x65, 0xfc, 0x69, 0x54, 0x88, 0xd2, 0xbd, 0x42,
	0xd5, 0x25, 0xf0, 0x88, 0x48, 0x5b, 0xac, 0xed, 0x74, 0x22, 0xe2, 0x3c, 0x08, 0x9b, 0x08, 0x72,
	0xf1, 0x55, 0xaa, 0xa2, 0xd3, 0xc9, 0x5e, 0x1e, 0x0e, 0x93, 0x17, 0x88, 0x46, 0x7b, 0xec, 0x6e,
	0x3e, 0xca, 0xfa, 0xda, 0x01, 0xb3, 0xfe, 0xfd, 0xc5, 0x36, 0xcb, 0x46, 0xb9, 0x09, 0x46, 0x7b,
	0xfe, 0x27, 0x8c, 0x5b, 0x27, 0xd6, 0x61, 0xb2, 0x5f, 0x72, 0xf6, 0x26, 0x94, 0xe8, 0x3e, 0x5b,
	0xcc, 0xc6, 0x05, 0x08, 0xaa, 0xf6, 0x3a, 0x0f, 0x7c, 0x0d, 0x7b, 0xba, 0x00, 0xd8, 0x11, 0x41,
	0x30, 0xa5, 0x70, 0x60, 0x17, 0x6a, 0xd2, 0x0d, 0x24, 0xb5, 0x08, 0xad, 0x68, 0x8f, 0xd9, 0xca,
	0xd8, 0x25, 0x26, 0x2b, 0xac, 0x0b, 0x53, 0x5c, 0x8e, 0xdc, 0x68, 0xca, 0x55, 0x05, 0x3c, 0x62,
	0x11, 0xdd, 0x1f, 0xe9, 0xa7, 0xf8, 0x37, 0x91, 0xdf, 0x9e, 0x78, 0x2a, 0xfa, 0x16, 0x5b, 0x26,
	0xd0, 0xa9, 0x9a, 0x7c, 0x0b, 0xc9, 0xd7, 0x2b, 0xbc, 0xa2, 0x3e, 0x67, 0x37, 0xbd, 0x1e, 0x64,
	0x3a, 0x0f, 0x5a, 0x55, 0xb7, 0xaf, 0x8e, 0x59, 0xc7, 0x98, 0x1b, 0x35, 0xa1, 0xbc, 0x8c, 0x55,
	0xec, 0x5d, 0xb6, 0x50, 0xcf, 0x87, 0x51, 0xfc, 0x36, 0xbd, 0x5b, 0xca, 0xe9, 0xe8, 0xaa, 0xe8,
	0x31, 0x5b, 0x9d, 0xf2, 0x0b, 0xa7, 0x13, 0xfa, 0xc9, 0xdd, 0xa1, 0xff, 0x75, 0x4d, 0xec, 0x95,
	0x0e, 0x18, 0x49, 0xeb, 0x8b, 0x44, 0x48, 0xa7, 0x25, 0xac, 0x78, 0x17, 0x47, 0x97, 0x01, 0xb6,
	0xeb, 0xb4, 0xec, 0xaa, 0xe8, 0xaf, 0x2c, 0x72, 0x3a, 0xb3, 0x41, 0x97, 0xfd, 0x16, 0xa0, 0x36,
	0xfc, 0x5e, 0x67, 0x66, 0x73, 0xb1, 0xb7, 0x4c, 0x1e, 0x6a, 0xf9, 0xae, 0x52, 0x0e, 0x3a, 0x36,
	0x94, 0x9e, 0x46, 0xd3, 0x87, 0x4f, 0xbc, 0x43, 0x1d, 0x1b, 0x4a, 0x7f, 0x50, 0x42, 0x20, 0x3b,
	0xf9, 0x28, 0x2b, 0x29, 0xfc, 0x7e, 0x99, 0xc2, 0x28, 0x23, 0x02, 0x3c, 0x7a, 0xea, 0xe8, 0x8d,
	0xce, 0x0c, 0x0c, 0x6f, 0x65, 0xe3, 0x90, 0x9a, 0x5c, 0x99, 0x7c, 0x50, 0x0e, 0xff, 0x83, 0x72,
	0x48, 0x09, 0xac, 0xc7, 0x3f, 0x1f, 0x1a, 0x25, 0x12, 0x78, 0x20, 0xff, 0x19, 0x95, 0x65, 0x1e,
	0x80, 0x57, 0xf0, 0x42, 0x5e, 0x67, 0xcd, 0xac, 0x48, 0x3d, 0x39, 0x1f, 0x92, 0x13, 0x00, 0x70,
	0x6e, 0x7c, 0x77, 0x85, 0xad, 0x54, 0x7a, 0x9b, 0xda, 0x58, 0xa6, 0xb4, 0xcb, 0xc5, 0xa7, 0x68,
	0xe3, 0x92, 0xa7, 0xe8, 0x99, 0xe7, 0xe6, 0x95, 0x73, 0xcf, 0xcd, 0x55, 0x76, 0xcd, 0x07, 0x99,
	0x6a, 0x54, 0xd9, 0x56, 0x8f, 0x0c, 0x48, 0x35, 0x33, 0xce, 0x59, 0xa7, 0x15, 0x6a, 0x6b, 0xab,
	0x57, 0xdb, 0xb0, 0x67, 0xa6, 0xdd, 0x40, 0xc3, 0xdb, 0x0e, 0x04, 0xa4, 0x54, 0xd7, 0x45, 0x04,
	0x5f, 0x12, 0x86, 0x52, 0xa0, 0xe1, 0xf1, 0x27, 0x6c, 0x9e, 0x9e, 0x56, 0x1a, 0x4b, 0xd0, 0x61,
	0x9e, 0x9e, 0xc2, 0xbe, 0x54, 0xa8, 0x39, 0xda, 0x97, 0xf2, 0x99, 0x7e, 0x57, 0xce, 0x9f, 0x7d,
	0x57, 0xf6, 0x67, 0x31, 0xa9, 0xa7, 0xbf, 0x05, 0x00, 0x00, 0xff, 0xff, 0x1c, 0x3c, 0x6c, 0x1c,
	0x0f, 0x0d, 0x00, 0x00,
}
