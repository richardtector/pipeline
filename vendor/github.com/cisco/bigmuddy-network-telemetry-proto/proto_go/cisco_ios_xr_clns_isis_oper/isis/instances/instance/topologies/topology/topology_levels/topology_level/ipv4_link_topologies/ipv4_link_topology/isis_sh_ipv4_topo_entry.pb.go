// Code generated by protoc-gen-go. DO NOT EDIT.
// source: isis_sh_ipv4_topo_entry.proto

/*
Package cisco_ios_xr_clns_isis_oper_isis_instances_instance_topologies_topology_topology_levels_topology_level_ipv4_link_topologies_ipv4_link_topology is a generated protocol buffer package.

It is generated from these files:
	isis_sh_ipv4_topo_entry.proto

It has these top-level messages:
	IsisShIpv4TopoEntry_KEYS
	IsisShIpv4TopoEntry
	IsisNodeIdType
	IsisSnpaType
	IsisPerPriorityCounts
	IsisShRepEl
	IsisShIpv4FrrBackup
	IsisShIpv4Path
	IsisShTopoNeighbor
	IsisShIpv4TopoReachableDetails
	IsisShIpv4TopoReachableStatus
*/
package cisco_ios_xr_clns_isis_oper_isis_instances_instance_topologies_topology_topology_levels_topology_level_ipv4_link_topologies_ipv4_link_topology

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

// IPv4 IS Link Topology Entry
type IsisShIpv4TopoEntry_KEYS struct {
	InstanceName string `protobuf:"bytes,1,opt,name=instance_name,json=instanceName" json:"instance_name,omitempty"`
	AfName       string `protobuf:"bytes,2,opt,name=af_name,json=afName" json:"af_name,omitempty"`
	SafName      string `protobuf:"bytes,3,opt,name=saf_name,json=safName" json:"saf_name,omitempty"`
	TopologyName string `protobuf:"bytes,4,opt,name=topology_name,json=topologyName" json:"topology_name,omitempty"`
	Level        string `protobuf:"bytes,5,opt,name=level" json:"level,omitempty"`
	SystemId     string `protobuf:"bytes,6,opt,name=system_id,json=systemId" json:"system_id,omitempty"`
}

func (m *IsisShIpv4TopoEntry_KEYS) Reset()                    { *m = IsisShIpv4TopoEntry_KEYS{} }
func (m *IsisShIpv4TopoEntry_KEYS) String() string            { return proto.CompactTextString(m) }
func (*IsisShIpv4TopoEntry_KEYS) ProtoMessage()               {}
func (*IsisShIpv4TopoEntry_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IsisShIpv4TopoEntry_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *IsisShIpv4TopoEntry_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *IsisShIpv4TopoEntry_KEYS) GetSafName() string {
	if m != nil {
		return m.SafName
	}
	return ""
}

func (m *IsisShIpv4TopoEntry_KEYS) GetTopologyName() string {
	if m != nil {
		return m.TopologyName
	}
	return ""
}

func (m *IsisShIpv4TopoEntry_KEYS) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

func (m *IsisShIpv4TopoEntry_KEYS) GetSystemId() string {
	if m != nil {
		return m.SystemId
	}
	return ""
}

type IsisShIpv4TopoEntry struct {
	// Source Address
	SourceAddress string `protobuf:"bytes,50,opt,name=source_address,json=sourceAddress" json:"source_address,omitempty"`
	// Does the IS participate in the topology?
	IsParticipant bool `protobuf:"varint,51,opt,name=is_participant,json=isParticipant" json:"is_participant,omitempty"`
	// Is the IS overloaded?
	IsOverloaded bool `protobuf:"varint,52,opt,name=is_overloaded,json=isOverloaded" json:"is_overloaded,omitempty"`
	// Is the IS attached?
	IsAttached bool `protobuf:"varint,53,opt,name=is_attached,json=isAttached" json:"is_attached,omitempty"`
	// Is the IS reachable, and, if so, its status within the SPT
	ReachabilityStatus *IsisShIpv4TopoReachableStatus `protobuf:"bytes,54,opt,name=reachability_status,json=reachabilityStatus" json:"reachability_status,omitempty"`
	// Per-priority counts of prefix items advertised by the IS
	AdvertisedPrefixItemCounts *IsisPerPriorityCounts `protobuf:"bytes,55,opt,name=advertised_prefix_item_counts,json=advertisedPrefixItemCounts" json:"advertised_prefix_item_counts,omitempty"`
}

func (m *IsisShIpv4TopoEntry) Reset()                    { *m = IsisShIpv4TopoEntry{} }
func (m *IsisShIpv4TopoEntry) String() string            { return proto.CompactTextString(m) }
func (*IsisShIpv4TopoEntry) ProtoMessage()               {}
func (*IsisShIpv4TopoEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *IsisShIpv4TopoEntry) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *IsisShIpv4TopoEntry) GetIsParticipant() bool {
	if m != nil {
		return m.IsParticipant
	}
	return false
}

func (m *IsisShIpv4TopoEntry) GetIsOverloaded() bool {
	if m != nil {
		return m.IsOverloaded
	}
	return false
}

func (m *IsisShIpv4TopoEntry) GetIsAttached() bool {
	if m != nil {
		return m.IsAttached
	}
	return false
}

func (m *IsisShIpv4TopoEntry) GetReachabilityStatus() *IsisShIpv4TopoReachableStatus {
	if m != nil {
		return m.ReachabilityStatus
	}
	return nil
}

func (m *IsisShIpv4TopoEntry) GetAdvertisedPrefixItemCounts() *IsisPerPriorityCounts {
	if m != nil {
		return m.AdvertisedPrefixItemCounts
	}
	return nil
}

type IsisNodeIdType struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *IsisNodeIdType) Reset()                    { *m = IsisNodeIdType{} }
func (m *IsisNodeIdType) String() string            { return proto.CompactTextString(m) }
func (*IsisNodeIdType) ProtoMessage()               {}
func (*IsisNodeIdType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *IsisNodeIdType) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type IsisSnpaType struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *IsisSnpaType) Reset()                    { *m = IsisSnpaType{} }
func (m *IsisSnpaType) String() string            { return proto.CompactTextString(m) }
func (*IsisSnpaType) ProtoMessage()               {}
func (*IsisSnpaType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *IsisSnpaType) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Per-priority counts
type IsisPerPriorityCounts struct {
	// Critical priority
	Critical uint32 `protobuf:"varint,1,opt,name=critical" json:"critical,omitempty"`
	// High priority
	High uint32 `protobuf:"varint,2,opt,name=high" json:"high,omitempty"`
	// Medium priority
	Medium uint32 `protobuf:"varint,3,opt,name=medium" json:"medium,omitempty"`
	// Low priority
	Low uint32 `protobuf:"varint,4,opt,name=low" json:"low,omitempty"`
}

func (m *IsisPerPriorityCounts) Reset()                    { *m = IsisPerPriorityCounts{} }
func (m *IsisPerPriorityCounts) String() string            { return proto.CompactTextString(m) }
func (*IsisPerPriorityCounts) ProtoMessage()               {}
func (*IsisPerPriorityCounts) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *IsisPerPriorityCounts) GetCritical() uint32 {
	if m != nil {
		return m.Critical
	}
	return 0
}

func (m *IsisPerPriorityCounts) GetHigh() uint32 {
	if m != nil {
		return m.High
	}
	return 0
}

func (m *IsisPerPriorityCounts) GetMedium() uint32 {
	if m != nil {
		return m.Medium
	}
	return 0
}

func (m *IsisPerPriorityCounts) GetLow() uint32 {
	if m != nil {
		return m.Low
	}
	return 0
}

// OSPF Repair Element
type IsisShRepEl struct {
	// RepairElementNodeID
	RepairElementNodeId string `protobuf:"bytes,1,opt,name=repair_element_node_id,json=repairElementNodeId" json:"repair_element_node_id,omitempty"`
	// RepairIPv4Addr
	RepairIpv4Addr string `protobuf:"bytes,2,opt,name=repair_ipv4_addr,json=repairIpv4Addr" json:"repair_ipv4_addr,omitempty"`
	// RepairIPv6Addr
	RepairIpv6Addr string `protobuf:"bytes,3,opt,name=repair_ipv6_addr,json=repairIpv6Addr" json:"repair_ipv6_addr,omitempty"`
	// Repair Label
	RepairLabel uint32 `protobuf:"varint,4,opt,name=repair_label,json=repairLabel" json:"repair_label,omitempty"`
	// Repair Element Type
	RepairElementType uint32 `protobuf:"varint,5,opt,name=repair_element_type,json=repairElementType" json:"repair_element_type,omitempty"`
}

func (m *IsisShRepEl) Reset()                    { *m = IsisShRepEl{} }
func (m *IsisShRepEl) String() string            { return proto.CompactTextString(m) }
func (*IsisShRepEl) ProtoMessage()               {}
func (*IsisShRepEl) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *IsisShRepEl) GetRepairElementNodeId() string {
	if m != nil {
		return m.RepairElementNodeId
	}
	return ""
}

func (m *IsisShRepEl) GetRepairIpv4Addr() string {
	if m != nil {
		return m.RepairIpv4Addr
	}
	return ""
}

func (m *IsisShRepEl) GetRepairIpv6Addr() string {
	if m != nil {
		return m.RepairIpv6Addr
	}
	return ""
}

func (m *IsisShRepEl) GetRepairLabel() uint32 {
	if m != nil {
		return m.RepairLabel
	}
	return 0
}

func (m *IsisShRepEl) GetRepairElementType() uint32 {
	if m != nil {
		return m.RepairElementType
	}
	return 0
}

// FRR backup path
type IsisShIpv4FrrBackup struct {
	// Next hop neighbor ID
	NeighborId string `protobuf:"bytes,1,opt,name=neighbor_id,json=neighborId" json:"neighbor_id,omitempty"`
	// Interface to send the packet out of
	EgressInterface string `protobuf:"bytes,2,opt,name=egress_interface,json=egressInterface" json:"egress_interface,omitempty"`
	// Next hop neighbor's forwarding address
	NeighborAddress string `protobuf:"bytes,3,opt,name=neighbor_address,json=neighborAddress" json:"neighbor_address,omitempty"`
	// Tunnel Interface to send the packet out of
	TunnelEgressInterface string `protobuf:"bytes,4,opt,name=tunnel_egress_interface,json=tunnelEgressInterface" json:"tunnel_egress_interface,omitempty"`
	// Next hop neighbor's SNPA
	NeighborSnpa *IsisSnpaType `protobuf:"bytes,5,opt,name=neighbor_snpa,json=neighborSnpa" json:"neighbor_snpa,omitempty"`
	// Remote LFA PQ Node's ID
	RemoteLfaSystemId string `protobuf:"bytes,6,opt,name=remote_lfa_system_id,json=remoteLfaSystemId" json:"remote_lfa_system_id,omitempty"`
	// Remote LFA Router ID
	RemoteLfaRouterId string `protobuf:"bytes,7,opt,name=remote_lfa_router_id,json=remoteLfaRouterId" json:"remote_lfa_router_id,omitempty"`
	// Remote LFA PQ Node's ID
	RemoteLfaSystemPid string `protobuf:"bytes,8,opt,name=remote_lfa_system_pid,json=remoteLfaSystemPid" json:"remote_lfa_system_pid,omitempty"`
	// Remote LFA Router ID
	RemoteLfaRouterPid string `protobuf:"bytes,9,opt,name=remote_lfa_router_pid,json=remoteLfaRouterPid" json:"remote_lfa_router_pid,omitempty"`
	// Distance to the network via this backup path
	TotalBackupDistance uint32 `protobuf:"varint,10,opt,name=total_backup_distance,json=totalBackupDistance" json:"total_backup_distance,omitempty"`
	// Segment routing sid value received from first hop
	SegmentRoutingSidValue uint32 `protobuf:"varint,11,opt,name=segment_routing_sid_value,json=segmentRoutingSidValue" json:"segment_routing_sid_value,omitempty"`
	// Number of SIDs in TI-LFA/rLFA
	NumSid uint32 `protobuf:"varint,12,opt,name=num_sid,json=numSid" json:"num_sid,omitempty"`
	// Segment routing sid values for TI-LFA/rLFA
	SegmentRoutingSidValues []uint32 `protobuf:"varint,13,rep,packed,name=segment_routing_sid_values,json=segmentRoutingSidValues" json:"segment_routing_sid_values,omitempty"`
	// Backup Repair List Size
	BackupRepairListSize uint32 `protobuf:"varint,14,opt,name=backup_repair_list_size,json=backupRepairListSize" json:"backup_repair_list_size,omitempty"`
	// Ti LFA computation which provided backup path
	TilfaComputation string `protobuf:"bytes,15,opt,name=tilfa_computation,json=tilfaComputation" json:"tilfa_computation,omitempty"`
	// Backup Repair List
	BackupRepairList []*IsisShRepEl `protobuf:"bytes,16,rep,name=backup_repair_list,json=backupRepairList" json:"backup_repair_list,omitempty"`
	// Is the backup path via downstream node?
	IsDownstream bool `protobuf:"varint,17,opt,name=is_downstream,json=isDownstream" json:"is_downstream,omitempty"`
	// Is the backup path line card disjoint with primary?
	IsLcDisjoint bool `protobuf:"varint,18,opt,name=is_lc_disjoint,json=isLcDisjoint" json:"is_lc_disjoint,omitempty"`
	// Is the backup path node protecting?
	IsNodeProtecting bool `protobuf:"varint,19,opt,name=is_node_protecting,json=isNodeProtecting" json:"is_node_protecting,omitempty"`
	// Is the backup path an ECMP to the network?
	IsPrimaryPath bool `protobuf:"varint,20,opt,name=is_primary_path,json=isPrimaryPath" json:"is_primary_path,omitempty"`
	// Is the backup path SRLG disjoint with primary?
	IsSrlgDisjoint bool `protobuf:"varint,21,opt,name=is_srlg_disjoint,json=isSrlgDisjoint" json:"is_srlg_disjoint,omitempty"`
	// Is the backup path via a Remote LFA?
	IsRemoteLfa bool `protobuf:"varint,22,opt,name=is_remote_lfa,json=isRemoteLfa" json:"is_remote_lfa,omitempty"`
	// Is the backup path via a TI-LFA?
	IsEpcfrrLfa bool `protobuf:"varint,23,opt,name=is_epcfrr_lfa,json=isEpcfrrLfa" json:"is_epcfrr_lfa,omitempty"`
	// Is SR TE tunnel requested
	IsTunnelRequested bool `protobuf:"varint,24,opt,name=is_tunnel_requested,json=isTunnelRequested" json:"is_tunnel_requested,omitempty"`
	// Weight configured on the interface
	Weight uint32 `protobuf:"varint,25,opt,name=weight" json:"weight,omitempty"`
}

func (m *IsisShIpv4FrrBackup) Reset()                    { *m = IsisShIpv4FrrBackup{} }
func (m *IsisShIpv4FrrBackup) String() string            { return proto.CompactTextString(m) }
func (*IsisShIpv4FrrBackup) ProtoMessage()               {}
func (*IsisShIpv4FrrBackup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *IsisShIpv4FrrBackup) GetNeighborId() string {
	if m != nil {
		return m.NeighborId
	}
	return ""
}

func (m *IsisShIpv4FrrBackup) GetEgressInterface() string {
	if m != nil {
		return m.EgressInterface
	}
	return ""
}

func (m *IsisShIpv4FrrBackup) GetNeighborAddress() string {
	if m != nil {
		return m.NeighborAddress
	}
	return ""
}

func (m *IsisShIpv4FrrBackup) GetTunnelEgressInterface() string {
	if m != nil {
		return m.TunnelEgressInterface
	}
	return ""
}

func (m *IsisShIpv4FrrBackup) GetNeighborSnpa() *IsisSnpaType {
	if m != nil {
		return m.NeighborSnpa
	}
	return nil
}

func (m *IsisShIpv4FrrBackup) GetRemoteLfaSystemId() string {
	if m != nil {
		return m.RemoteLfaSystemId
	}
	return ""
}

func (m *IsisShIpv4FrrBackup) GetRemoteLfaRouterId() string {
	if m != nil {
		return m.RemoteLfaRouterId
	}
	return ""
}

func (m *IsisShIpv4FrrBackup) GetRemoteLfaSystemPid() string {
	if m != nil {
		return m.RemoteLfaSystemPid
	}
	return ""
}

func (m *IsisShIpv4FrrBackup) GetRemoteLfaRouterPid() string {
	if m != nil {
		return m.RemoteLfaRouterPid
	}
	return ""
}

func (m *IsisShIpv4FrrBackup) GetTotalBackupDistance() uint32 {
	if m != nil {
		return m.TotalBackupDistance
	}
	return 0
}

func (m *IsisShIpv4FrrBackup) GetSegmentRoutingSidValue() uint32 {
	if m != nil {
		return m.SegmentRoutingSidValue
	}
	return 0
}

func (m *IsisShIpv4FrrBackup) GetNumSid() uint32 {
	if m != nil {
		return m.NumSid
	}
	return 0
}

func (m *IsisShIpv4FrrBackup) GetSegmentRoutingSidValues() []uint32 {
	if m != nil {
		return m.SegmentRoutingSidValues
	}
	return nil
}

func (m *IsisShIpv4FrrBackup) GetBackupRepairListSize() uint32 {
	if m != nil {
		return m.BackupRepairListSize
	}
	return 0
}

func (m *IsisShIpv4FrrBackup) GetTilfaComputation() string {
	if m != nil {
		return m.TilfaComputation
	}
	return ""
}

func (m *IsisShIpv4FrrBackup) GetBackupRepairList() []*IsisShRepEl {
	if m != nil {
		return m.BackupRepairList
	}
	return nil
}

func (m *IsisShIpv4FrrBackup) GetIsDownstream() bool {
	if m != nil {
		return m.IsDownstream
	}
	return false
}

func (m *IsisShIpv4FrrBackup) GetIsLcDisjoint() bool {
	if m != nil {
		return m.IsLcDisjoint
	}
	return false
}

func (m *IsisShIpv4FrrBackup) GetIsNodeProtecting() bool {
	if m != nil {
		return m.IsNodeProtecting
	}
	return false
}

func (m *IsisShIpv4FrrBackup) GetIsPrimaryPath() bool {
	if m != nil {
		return m.IsPrimaryPath
	}
	return false
}

func (m *IsisShIpv4FrrBackup) GetIsSrlgDisjoint() bool {
	if m != nil {
		return m.IsSrlgDisjoint
	}
	return false
}

func (m *IsisShIpv4FrrBackup) GetIsRemoteLfa() bool {
	if m != nil {
		return m.IsRemoteLfa
	}
	return false
}

func (m *IsisShIpv4FrrBackup) GetIsEpcfrrLfa() bool {
	if m != nil {
		return m.IsEpcfrrLfa
	}
	return false
}

func (m *IsisShIpv4FrrBackup) GetIsTunnelRequested() bool {
	if m != nil {
		return m.IsTunnelRequested
	}
	return false
}

func (m *IsisShIpv4FrrBackup) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

// IPv4 path to a destination
type IsisShIpv4Path struct {
	// Next hop neighbor ID
	NeighborId string `protobuf:"bytes,1,opt,name=neighbor_id,json=neighborId" json:"neighbor_id,omitempty"`
	// Interface to send the packet out of
	EgressInterface string `protobuf:"bytes,2,opt,name=egress_interface,json=egressInterface" json:"egress_interface,omitempty"`
	// Next hop neighbor's forwarding address
	NeighborAddress string `protobuf:"bytes,3,opt,name=neighbor_address,json=neighborAddress" json:"neighbor_address,omitempty"`
	// Next hop neighbor's SNPA
	NeighborSnpa *IsisSnpaType `protobuf:"bytes,4,opt,name=neighbor_snpa,json=neighborSnpa" json:"neighbor_snpa,omitempty"`
	// Tag associated with the path
	Tag uint32 `protobuf:"varint,5,opt,name=tag" json:"tag,omitempty"`
	// FRR backup for this path
	FrrBackup *IsisShIpv4FrrBackup `protobuf:"bytes,6,opt,name=frr_backup,json=frrBackup" json:"frr_backup,omitempty"`
	// Uloop Explicit List
	UloopExplicitList []*IsisShRepEl `protobuf:"bytes,7,rep,name=uloop_explicit_list,json=uloopExplicitList" json:"uloop_explicit_list,omitempty"`
	// Explicit path tunnel interface
	TunnelInterface string `protobuf:"bytes,8,opt,name=tunnel_interface,json=tunnelInterface" json:"tunnel_interface,omitempty"`
	// Segment routing sid value received from first hop
	SegmentRoutingSidValue uint32 `protobuf:"varint,9,opt,name=segment_routing_sid_value,json=segmentRoutingSidValue" json:"segment_routing_sid_value,omitempty"`
	// Weight configured on the interface
	Weight uint32 `protobuf:"varint,10,opt,name=weight" json:"weight,omitempty"`
}

func (m *IsisShIpv4Path) Reset()                    { *m = IsisShIpv4Path{} }
func (m *IsisShIpv4Path) String() string            { return proto.CompactTextString(m) }
func (*IsisShIpv4Path) ProtoMessage()               {}
func (*IsisShIpv4Path) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *IsisShIpv4Path) GetNeighborId() string {
	if m != nil {
		return m.NeighborId
	}
	return ""
}

func (m *IsisShIpv4Path) GetEgressInterface() string {
	if m != nil {
		return m.EgressInterface
	}
	return ""
}

func (m *IsisShIpv4Path) GetNeighborAddress() string {
	if m != nil {
		return m.NeighborAddress
	}
	return ""
}

func (m *IsisShIpv4Path) GetNeighborSnpa() *IsisSnpaType {
	if m != nil {
		return m.NeighborSnpa
	}
	return nil
}

func (m *IsisShIpv4Path) GetTag() uint32 {
	if m != nil {
		return m.Tag
	}
	return 0
}

func (m *IsisShIpv4Path) GetFrrBackup() *IsisShIpv4FrrBackup {
	if m != nil {
		return m.FrrBackup
	}
	return nil
}

func (m *IsisShIpv4Path) GetUloopExplicitList() []*IsisShRepEl {
	if m != nil {
		return m.UloopExplicitList
	}
	return nil
}

func (m *IsisShIpv4Path) GetTunnelInterface() string {
	if m != nil {
		return m.TunnelInterface
	}
	return ""
}

func (m *IsisShIpv4Path) GetSegmentRoutingSidValue() uint32 {
	if m != nil {
		return m.SegmentRoutingSidValue
	}
	return 0
}

func (m *IsisShIpv4Path) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

// SPT Neighbor
type IsisShTopoNeighbor struct {
	// Neighbor ID
	NeighborId string `protobuf:"bytes,1,opt,name=neighbor_id,json=neighborId" json:"neighbor_id,omitempty"`
	// Pseudonode between system and its neighbor
	IntermediatePseudonode *IsisNodeIdType `protobuf:"bytes,2,opt,name=intermediate_pseudonode,json=intermediatePseudonode" json:"intermediate_pseudonode,omitempty"`
}

func (m *IsisShTopoNeighbor) Reset()                    { *m = IsisShTopoNeighbor{} }
func (m *IsisShTopoNeighbor) String() string            { return proto.CompactTextString(m) }
func (*IsisShTopoNeighbor) ProtoMessage()               {}
func (*IsisShTopoNeighbor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *IsisShTopoNeighbor) GetNeighborId() string {
	if m != nil {
		return m.NeighborId
	}
	return ""
}

func (m *IsisShTopoNeighbor) GetIntermediatePseudonode() *IsisNodeIdType {
	if m != nil {
		return m.IntermediatePseudonode
	}
	return nil
}

// Status of a reachable IPv4 IS
type IsisShIpv4TopoReachableDetails struct {
	// Distance to the IS
	RootDistance uint32 `protobuf:"varint,1,opt,name=root_distance,json=rootDistance" json:"root_distance,omitempty"`
	// Distance to the IS
	MulticastRootDistance uint32 `protobuf:"varint,2,opt,name=multicast_root_distance,json=multicastRootDistance" json:"multicast_root_distance,omitempty"`
	// First hops towards the IS
	Paths []*IsisShIpv4Path `protobuf:"bytes,3,rep,name=paths" json:"paths,omitempty"`
	// Multicast intact first hops towards the IS
	MulticastPaths []*IsisShIpv4Path `protobuf:"bytes,4,rep,name=multicast_paths,json=multicastPaths" json:"multicast_paths,omitempty"`
	// Parents of the IS within the SPT
	Parents []*IsisShTopoNeighbor `protobuf:"bytes,5,rep,name=parents" json:"parents,omitempty"`
	// Children of the IS within the SPT
	Children []*IsisShTopoNeighbor `protobuf:"bytes,6,rep,name=children" json:"children,omitempty"`
}

func (m *IsisShIpv4TopoReachableDetails) Reset()                    { *m = IsisShIpv4TopoReachableDetails{} }
func (m *IsisShIpv4TopoReachableDetails) String() string            { return proto.CompactTextString(m) }
func (*IsisShIpv4TopoReachableDetails) ProtoMessage()               {}
func (*IsisShIpv4TopoReachableDetails) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *IsisShIpv4TopoReachableDetails) GetRootDistance() uint32 {
	if m != nil {
		return m.RootDistance
	}
	return 0
}

func (m *IsisShIpv4TopoReachableDetails) GetMulticastRootDistance() uint32 {
	if m != nil {
		return m.MulticastRootDistance
	}
	return 0
}

func (m *IsisShIpv4TopoReachableDetails) GetPaths() []*IsisShIpv4Path {
	if m != nil {
		return m.Paths
	}
	return nil
}

func (m *IsisShIpv4TopoReachableDetails) GetMulticastPaths() []*IsisShIpv4Path {
	if m != nil {
		return m.MulticastPaths
	}
	return nil
}

func (m *IsisShIpv4TopoReachableDetails) GetParents() []*IsisShTopoNeighbor {
	if m != nil {
		return m.Parents
	}
	return nil
}

func (m *IsisShIpv4TopoReachableDetails) GetChildren() []*IsisShTopoNeighbor {
	if m != nil {
		return m.Children
	}
	return nil
}

// Reachability status of an IPv4 IS
type IsisShIpv4TopoReachableStatus struct {
	ReachableStatus string `protobuf:"bytes,1,opt,name=reachable_status,json=reachableStatus" json:"reachable_status,omitempty"`
	// Status of the IS within the SPT
	ReachableDetails *IsisShIpv4TopoReachableDetails `protobuf:"bytes,2,opt,name=reachable_details,json=reachableDetails" json:"reachable_details,omitempty"`
}

func (m *IsisShIpv4TopoReachableStatus) Reset()                    { *m = IsisShIpv4TopoReachableStatus{} }
func (m *IsisShIpv4TopoReachableStatus) String() string            { return proto.CompactTextString(m) }
func (*IsisShIpv4TopoReachableStatus) ProtoMessage()               {}
func (*IsisShIpv4TopoReachableStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *IsisShIpv4TopoReachableStatus) GetReachableStatus() string {
	if m != nil {
		return m.ReachableStatus
	}
	return ""
}

func (m *IsisShIpv4TopoReachableStatus) GetReachableDetails() *IsisShIpv4TopoReachableDetails {
	if m != nil {
		return m.ReachableDetails
	}
	return nil
}

func init() {
	proto.RegisterType((*IsisShIpv4TopoEntry_KEYS)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.ipv4_link_topologies.ipv4_link_topology.isis_sh_ipv4_topo_entry_KEYS")
	proto.RegisterType((*IsisShIpv4TopoEntry)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.ipv4_link_topologies.ipv4_link_topology.isis_sh_ipv4_topo_entry")
	proto.RegisterType((*IsisNodeIdType)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.ipv4_link_topologies.ipv4_link_topology.isis_node_id_type")
	proto.RegisterType((*IsisSnpaType)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.ipv4_link_topologies.ipv4_link_topology.isis_snpa_type")
	proto.RegisterType((*IsisPerPriorityCounts)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.ipv4_link_topologies.ipv4_link_topology.isis_per_priority_counts")
	proto.RegisterType((*IsisShRepEl)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.ipv4_link_topologies.ipv4_link_topology.isis_sh_rep_el")
	proto.RegisterType((*IsisShIpv4FrrBackup)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.ipv4_link_topologies.ipv4_link_topology.isis_sh_ipv4_frr_backup")
	proto.RegisterType((*IsisShIpv4Path)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.ipv4_link_topologies.ipv4_link_topology.isis_sh_ipv4_path")
	proto.RegisterType((*IsisShTopoNeighbor)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.ipv4_link_topologies.ipv4_link_topology.isis_sh_topo_neighbor")
	proto.RegisterType((*IsisShIpv4TopoReachableDetails)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.ipv4_link_topologies.ipv4_link_topology.isis_sh_ipv4_topo_reachable_details")
	proto.RegisterType((*IsisShIpv4TopoReachableStatus)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.ipv4_link_topologies.ipv4_link_topology.isis_sh_ipv4_topo_reachable_status")
}

func init() { proto.RegisterFile("isis_sh_ipv4_topo_entry.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1459 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x58, 0xcd, 0x6f, 0x1c, 0xc5,
	0x12, 0xd7, 0xc4, 0xdf, 0xb5, 0x5e, 0x7b, 0xb7, 0xfd, 0xb1, 0x93, 0xbc, 0x17, 0x3d, 0xbf, 0xcd,
	0x7b, 0x91, 0x23, 0xd0, 0x22, 0x9c, 0xc4, 0x08, 0x71, 0x0a, 0xb1, 0x0f, 0x16, 0x51, 0xb0, 0x66,
	0x23, 0x24, 0x4e, 0xad, 0xf6, 0x4c, 0xef, 0x6e, 0x93, 0x99, 0xe9, 0xa1, 0xbb, 0xc7, 0xc9, 0xe6,
	0xca, 0x9d, 0x0b, 0x12, 0x47, 0x24, 0x0e, 0x08, 0x24, 0x6e, 0xe4, 0xc2, 0x11, 0x21, 0x24, 0x0e,
	0x08, 0xe5, 0x0f, 0xe0, 0x84, 0xc4, 0x31, 0xff, 0x02, 0x12, 0xea, 0xea, 0xf9, 0xb0, 0xd7, 0xd8,
	0x39, 0xb2, 0xbe, 0xf5, 0xfc, 0x7e, 0x55, 0xd3, 0x55, 0x3d, 0x55, 0xbf, 0xae, 0x5d, 0xb8, 0x2e,
	0xb4, 0xd0, 0x54, 0x8f, 0xa8, 0xc8, 0x8e, 0xef, 0x50, 0x23, 0x33, 0x49, 0x79, 0x6a, 0xd4, 0xb8,
	0x97, 0x29, 0x69, 0x24, 0xf9, 0xd4, 0x0b, 0x85, 0x0e, 0x25, 0x15, 0x52, 0xd3, 0xa7, 0x8a, 0x86,
	0x71, 0xaa, 0x29, 0x7a, 0xc8, 0x8c, 0xab, 0x9e, 0x5d, 0xf5, 0x44, 0xaa, 0x0d, 0x4b, 0x43, 0x5e,
	0xaf, 0x7a, 0xf6, 0x35, 0xb1, 0x1c, 0x0a, 0xae, 0xcb, 0xe5, 0xb8, 0x5a, 0xd0, 0x98, 0x1f, 0xf3,
	0x58, 0x4f, 0x3c, 0xf7, 0x70, 0xfb, 0x58, 0xa4, 0x8f, 0xe9, 0x09, 0xe7, 0x33, 0xe0, 0xb8, 0xfb,
	0xc2, 0x83, 0x7f, 0x9f, 0x13, 0x32, 0x7d, 0x6f, 0xff, 0xc3, 0x3e, 0xb9, 0x01, 0xcd, 0x32, 0x0e,
	0x9a, 0xb2, 0x84, 0xfb, 0xde, 0x96, 0xb7, 0xbd, 0x14, 0x2c, 0x97, 0xe0, 0x43, 0x96, 0x70, 0xd2,
	0x81, 0x05, 0x36, 0x70, 0xf4, 0x15, 0xa4, 0xe7, 0xd9, 0x00, 0x89, 0xab, 0xb0, 0xa8, 0x4b, 0x66,
	0x06, 0x99, 0x05, 0x5d, 0x50, 0x37, 0xa0, 0x59, 0xc5, 0x8c, 0xfc, 0xac, 0x7b, 0x71, 0x09, 0xa2,
	0xd1, 0x3a, 0xcc, 0x61, 0x3e, 0xfe, 0x1c, 0x92, 0xee, 0x81, 0xfc, 0x0b, 0x96, 0xf4, 0x58, 0x1b,
	0x9e, 0x50, 0x11, 0xf9, 0xf3, 0xc8, 0x2c, 0x3a, 0xe0, 0x20, 0xea, 0x7e, 0x32, 0x07, 0x9d, 0x73,
	0x32, 0x22, 0xff, 0x87, 0x15, 0x2d, 0x73, 0x15, 0x72, 0xca, 0xa2, 0x48, 0x71, 0xad, 0xfd, 0x1d,
	0xf4, 0x6e, 0x3a, 0xf4, 0x9e, 0x03, 0xad, 0x99, 0xd0, 0x34, 0x63, 0xca, 0x88, 0x50, 0x64, 0x2c,
	0x35, 0xfe, 0xed, 0x2d, 0x6f, 0x7b, 0x31, 0x68, 0x0a, 0x7d, 0x58, 0x83, 0x78, 0x34, 0x9a, 0xca,
	0x63, 0xae, 0x62, 0xc9, 0x22, 0x1e, 0xf9, 0x77, 0xd0, 0x6a, 0x59, 0xe8, 0xf7, 0x2b, 0x8c, 0xfc,
	0x07, 0x1a, 0x42, 0x53, 0x66, 0x0c, 0x0b, 0x47, 0x3c, 0xf2, 0xef, 0xa2, 0x09, 0x08, 0x7d, 0xaf,
	0x40, 0xc8, 0x1f, 0x1e, 0xac, 0x29, 0xce, 0xc2, 0x11, 0x3b, 0x12, 0xb1, 0x30, 0x63, 0xaa, 0x0d,
	0x33, 0xb9, 0xf6, 0x77, 0xb7, 0xbc, 0xed, 0xc6, 0xce, 0x77, 0x5e, 0x6f, 0xba, 0x2a, 0xa6, 0x77,
	0xf6, 0x6c, 0x8b, 0xe8, 0x63, 0x5e, 0x84, 0x1e, 0x90, 0x93, 0xf9, 0xf4, 0x11, 0x23, 0x2f, 0x3d,
	0xb8, 0xce, 0xa2, 0x63, 0xae, 0x8c, 0xd0, 0x3c, 0xa2, 0x99, 0xe2, 0x03, 0xf1, 0x94, 0x0a, 0xfb,
	0x11, 0x43, 0x99, 0xa7, 0x46, 0xfb, 0x6f, 0x61, 0xc2, 0xdf, 0x4c, 0x67, 0xc2, 0x19, 0x57, 0x34,
	0x53, 0x42, 0x2a, 0xfb, 0x85, 0x5c, 0xc0, 0xc1, 0xb5, 0x3a, 0x9f, 0x43, 0x4c, 0xe7, 0xc0, 0xf0,
	0xe4, 0x3e, 0x72, 0xdd, 0x5b, 0xd0, 0x46, 0xbf, 0x54, 0x46, 0x9c, 0x8a, 0x88, 0x9a, 0x71, 0x86,
	0xd5, 0x7c, 0xcc, 0xe2, 0xbc, 0xec, 0x21, 0xf7, 0xd0, 0xbd, 0x69, 0xab, 0xcd, 0x9e, 0x69, 0x9a,
	0xb1, 0x8b, 0xec, 0x0c, 0xf8, 0xe7, 0x85, 0x42, 0xae, 0xc1, 0x62, 0xa8, 0x84, 0x11, 0x21, 0x8b,
	0xd1, 0xa9, 0x19, 0x54, 0xcf, 0x84, 0xc0, 0xec, 0x48, 0x0c, 0x47, 0xd8, 0x99, 0xcd, 0x00, 0xd7,
	0x64, 0x13, 0xe6, 0x13, 0x1e, 0x89, 0x3c, 0xc1, 0xae, 0x6c, 0x06, 0xc5, 0x13, 0x69, 0xc1, 0x4c,
	0x2c, 0x9f, 0x60, 0x2b, 0x36, 0x03, 0xbb, 0xec, 0xbe, 0xf4, 0xca, 0xf0, 0x46, 0x54, 0xf1, 0x8c,
	0xf2, 0x98, 0xdc, 0x86, 0x4d, 0xc5, 0x33, 0x26, 0x14, 0xe5, 0x31, 0x4f, 0x78, 0x6a, 0xca, 0x2c,
	0x8b, 0x78, 0xd7, 0x1c, 0xbb, 0xef, 0xc8, 0x87, 0x32, 0xe2, 0x07, 0x11, 0xd9, 0x86, 0x56, 0xe1,
	0x84, 0x47, 0x6c, 0xfb, 0xaf, 0xd0, 0x8a, 0x15, 0x87, 0x1f, 0x64, 0xc7, 0x77, 0x6c, 0x03, 0x9e,
	0xb6, 0xdc, 0x75, 0x96, 0x33, 0x13, 0x96, 0xbb, 0x68, 0xf9, 0x5f, 0x58, 0x2e, 0x2c, 0x63, 0x76,
	0xc4, 0xe3, 0x22, 0xec, 0x86, 0xc3, 0x1e, 0x58, 0x88, 0xf4, 0x60, 0x6d, 0x22, 0x56, 0x7b, 0xc2,
	0x28, 0x27, 0xcd, 0xa0, 0x7d, 0x2a, 0xd0, 0x47, 0xe3, 0x8c, 0x77, 0xbf, 0x6d, 0x4c, 0xa8, 0xc7,
	0x40, 0x29, 0x7a, 0xc4, 0xc2, 0xc7, 0x79, 0x66, 0x5b, 0x39, 0xe5, 0x62, 0x38, 0x3a, 0x92, 0xaa,
	0x4e, 0x16, 0x4a, 0xe8, 0x20, 0x22, 0xb7, 0xa0, 0xc5, 0x87, 0x56, 0x41, 0xa8, 0x48, 0x0d, 0x57,
	0x03, 0x16, 0x96, 0x7a, 0xb8, 0xea, 0xf0, 0x83, 0x12, 0xb6, 0xa6, 0xd5, 0xbb, 0x4a, 0x2d, 0x72,
	0x49, 0xae, 0x96, 0x78, 0xa9, 0x46, 0xbb, 0xd0, 0x31, 0x79, 0x9a, 0xf2, 0x98, 0x9e, 0x79, 0xb9,
	0x93, 0xcc, 0x0d, 0x47, 0xef, 0x4f, 0x6c, 0xf1, 0x93, 0x07, 0xcd, 0x6a, 0x0f, 0x5b, 0x5c, 0x98,
	0x75, 0x63, 0xe7, 0x8b, 0x29, 0x95, 0x94, 0xb2, 0xfc, 0x83, 0xe5, 0x32, 0xea, 0x7e, 0x9a, 0x31,
	0xf2, 0x06, 0xac, 0x2b, 0x9e, 0x48, 0xc3, 0x69, 0x3c, 0x60, 0x74, 0x52, 0xf7, 0xdb, 0x8e, 0x7b,
	0x30, 0x60, 0xfd, 0xe2, 0x02, 0x98, 0x70, 0x50, 0x32, 0x37, 0x1c, 0xbf, 0xd7, 0xc2, 0x84, 0x43,
	0x80, 0xcc, 0x41, 0x44, 0xde, 0x84, 0x8d, 0xb3, 0x3b, 0x64, 0x22, 0xf2, 0x17, 0xd1, 0x83, 0x4c,
	0x6c, 0x71, 0x28, 0x26, 0x5d, 0x8a, 0x3d, 0xac, 0xcb, 0xd2, 0x84, 0x8b, 0xdb, 0xc4, 0xba, 0xec,
	0xc0, 0x86, 0x91, 0x86, 0xc5, 0x45, 0x35, 0xd1, 0x48, 0xb8, 0x83, 0xf5, 0x01, 0x6b, 0x71, 0x0d,
	0xc9, 0x77, 0x91, 0xdb, 0x2b, 0x28, 0xf2, 0x36, 0x5c, 0xd5, 0x7c, 0x88, 0x65, 0x6b, 0xf7, 0x10,
	0xe9, 0x90, 0x6a, 0x11, 0x51, 0x27, 0x0e, 0x0d, 0xf4, 0xdb, 0x2c, 0x0c, 0x02, 0xc7, 0xf7, 0x45,
	0xf4, 0x81, 0x65, 0xed, 0x95, 0x9c, 0xe6, 0x89, 0x35, 0xf7, 0x97, 0x5d, 0x8b, 0xa7, 0x79, 0xd2,
	0x17, 0x11, 0x79, 0x07, 0xae, 0x9d, 0xfb, 0x4e, 0xed, 0x37, 0xb7, 0x66, 0xb6, 0x9b, 0x41, 0xe7,
	0xef, 0x5f, 0xaa, 0xc9, 0x5d, 0xe8, 0x14, 0xe1, 0x97, 0x8d, 0x27, 0xb4, 0xa1, 0x5a, 0x3c, 0xe3,
	0xfe, 0x0a, 0xee, 0xb2, 0xee, 0xe8, 0xc0, 0xb5, 0xa0, 0xd0, 0xa6, 0x2f, 0x9e, 0x71, 0xf2, 0x1a,
	0xb4, 0x8d, 0xb0, 0x27, 0x15, 0xca, 0x24, 0xcb, 0x0d, 0x33, 0x42, 0xa6, 0xfe, 0x2a, 0x1e, 0x55,
	0x0b, 0x89, 0xfb, 0x35, 0x4e, 0x7e, 0xf5, 0x80, 0x9c, 0xdd, 0xc4, 0x6f, 0x6d, 0xcd, 0x4c, 0x71,
	0xf1, 0x96, 0xe2, 0x18, 0xb4, 0x26, 0x0f, 0xa0, 0x18, 0x13, 0x22, 0xf9, 0x24, 0xd5, 0x46, 0x71,
	0x96, 0xf8, 0xed, 0x72, 0x4c, 0xd8, 0xab, 0x30, 0xf2, 0x3f, 0x1c, 0x39, 0xe2, 0xd0, 0x96, 0xc5,
	0x47, 0x52, 0xa4, 0xc6, 0x27, 0xa5, 0xd5, 0x83, 0x70, 0xaf, 0xc0, 0xc8, 0xeb, 0x40, 0xca, 0x3b,
	0xc5, 0xce, 0x93, 0x3c, 0xb4, 0x5f, 0xc7, 0x5f, 0x43, 0xcb, 0x96, 0xd0, 0x56, 0x6a, 0x0f, 0x2b,
	0x9c, 0xdc, 0x84, 0x55, 0x7b, 0x5d, 0x28, 0x91, 0x30, 0x35, 0xa6, 0x19, 0x33, 0x23, 0x7f, 0xbd,
	0x9a, 0x63, 0x1c, 0x7a, 0xc8, 0xcc, 0xc8, 0x0a, 0xae, 0x4d, 0x41, 0xc5, 0xc3, 0x7a, 0xf7, 0x0d,
	0x34, 0x5c, 0x11, 0xba, 0xaf, 0xe2, 0x61, 0xb5, 0x7f, 0x17, 0x53, 0xa9, 0x2b, 0xdf, 0xdf, 0x44,
	0xb3, 0x86, 0xd0, 0x41, 0x59, 0xf0, 0x85, 0x0d, 0xcf, 0x42, 0xab, 0x9c, 0xd6, 0xa6, 0x53, 0xda,
	0xec, 0x23, 0x66, 0x6d, 0x7a, 0xb0, 0x26, 0x34, 0x2d, 0x54, 0x4d, 0xf1, 0x8f, 0x73, 0xae, 0x0d,
	0x8f, 0x7c, 0x1f, 0x2d, 0xdb, 0x42, 0x3f, 0x42, 0x26, 0x28, 0x09, 0x7b, 0x5d, 0x3d, 0xb1, 0x9a,
	0x60, 0xfc, 0xab, 0xae, 0x96, 0xdd, 0x53, 0xf7, 0xcf, 0xf9, 0xe2, 0x9a, 0x2d, 0xd5, 0xda, 0x26,
	0xf9, 0x4f, 0xe9, 0xf4, 0x59, 0xbd, 0x9d, 0xbd, 0x8c, 0x7a, 0xdb, 0x82, 0x19, 0xc3, 0x86, 0xc5,
	0x0d, 0x69, 0x97, 0xe4, 0x67, 0x0f, 0xa0, 0xbe, 0x06, 0x51, 0x78, 0x1b, 0x3b, 0x5f, 0x4f, 0xf7,
	0x60, 0x5a, 0xc7, 0x1b, 0x2c, 0x0d, 0x94, 0x72, 0xba, 0x4a, 0x5e, 0x78, 0xb0, 0x96, 0xc7, 0x52,
	0x66, 0x94, 0x3f, 0xcd, 0x62, 0x11, 0x0a, 0xe3, 0xb4, 0x65, 0xe1, 0x72, 0x68, 0x4b, 0x1b, 0x63,
	0xdf, 0x2f, 0x42, 0x47, 0x71, 0xb9, 0x05, 0xad, 0xa2, 0x8d, 0xea, 0x52, 0x76, 0xd7, 0xd6, 0xaa,
	0xc3, 0xeb, 0x52, 0xbe, 0xf0, 0x32, 0x59, 0xba, 0xf0, 0x32, 0xa9, 0xfb, 0x0f, 0x4e, 0xf5, 0xdf,
	0xe7, 0x57, 0x60, 0xa3, 0x8c, 0x11, 0x7f, 0x0a, 0x94, 0x95, 0xf4, 0xea, 0x1e, 0xfc, 0xcd, 0x83,
	0x0e, 0x86, 0x6c, 0x27, 0x4f, 0x66, 0x38, 0xcd, 0x34, 0xcf, 0x23, 0x69, 0xb5, 0x0d, 0x7b, 0xb1,
	0xb1, 0xf3, 0xe5, 0x74, 0x7e, 0x8e, 0x93, 0x13, 0x7d, 0xb0, 0x79, 0x32, 0x85, 0xc3, 0x2a, 0x83,
	0xee, 0x67, 0x0b, 0x70, 0xe3, 0xa2, 0x1f, 0x4a, 0x11, 0x37, 0x4c, 0xc4, 0xda, 0xde, 0x0d, 0x4a,
	0x4a, 0x53, 0x0f, 0x03, 0x6e, 0x78, 0x5f, 0xb6, 0x60, 0x35, 0x05, 0xec, 0x42, 0x27, 0xc9, 0x63,
	0x3b, 0xcc, 0x6b, 0xfb, 0xe9, 0x4e, 0x9a, 0xbb, 0x99, 0x7e, 0xa3, 0xa2, 0x83, 0x93, 0x7e, 0xdf,
	0x7b, 0x30, 0x67, 0x05, 0xd1, 0x2a, 0xd6, 0xcc, 0xf4, 0x1e, 0xe8, 0x49, 0xed, 0x0e, 0x5c, 0xc0,
	0xe4, 0x17, 0x0f, 0x56, 0xeb, 0x9c, 0x5d, 0x12, 0xb3, 0x97, 0x26, 0x89, 0x95, 0x2a, 0xf4, 0x43,
	0xcc, 0xe6, 0x07, 0x0f, 0x16, 0x32, 0xa6, 0xb8, 0xfd, 0x95, 0x3b, 0x87, 0x59, 0x7c, 0x35, 0xb5,
	0x59, 0x9c, 0x6a, 0xe3, 0xa0, 0x0c, 0x9b, 0xfc, 0xe8, 0xc1, 0x62, 0x38, 0x12, 0x71, 0xa4, 0x78,
	0xea, 0xcf, 0x5f, 0xaa, 0x1c, 0xaa, 0xb8, 0xbb, 0xcf, 0xaf, 0x40, 0xf7, 0xd5, 0x7f, 0x5f, 0x58,
	0x4d, 0x9d, 0xc4, 0x0a, 0x01, 0x5b, 0xad, 0xf0, 0xe2, 0x5f, 0x8d, 0xdf, 0x3d, 0x68, 0x9f, 0xe9,
	0xea, 0x42, 0xbf, 0x9e, 0x5f, 0xaa, 0xbf, 0x6e, 0x8a, 0xd8, 0x83, 0x3a, 0xf5, 0x3d, 0x87, 0x1c,
	0xcd, 0xe3, 0x5f, 0x97, 0xb7, 0xff, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x3a, 0x37, 0x37, 0xa1, 0xdb,
	0x14, 0x00, 0x00,
}
