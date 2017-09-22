// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mpls_sh_leaf.proto

/*
Package cisco_ios_xr_fib_common_oper_mpls_forwarding_nodes_node_label_fib_forwarding_details_forwarding_detail is a generated protocol buffer package.

It is generated from these files:
	mpls_sh_leaf.proto

It has these top-level messages:
	MplsShLeaf_KEYS
	MplsShLeaf
	MplsFwdInfo
	MplsAdjInfo
	MplsLdiInfo
	MplsMcastInfo
*/
package cisco_ios_xr_fib_common_oper_mpls_forwarding_nodes_node_label_fib_forwarding_details_forwarding_detail

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

// Information about label leaf
type MplsShLeaf_KEYS struct {
	NodeName   string `protobuf:"bytes,1,opt,name=node_name,json=nodeName" json:"node_name,omitempty"`
	LabelValue uint32 `protobuf:"varint,2,opt,name=label_value,json=labelValue" json:"label_value,omitempty"`
	Eos        string `protobuf:"bytes,3,opt,name=eos" json:"eos,omitempty"`
}

func (m *MplsShLeaf_KEYS) Reset()                    { *m = MplsShLeaf_KEYS{} }
func (m *MplsShLeaf_KEYS) String() string            { return proto.CompactTextString(m) }
func (*MplsShLeaf_KEYS) ProtoMessage()               {}
func (*MplsShLeaf_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MplsShLeaf_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *MplsShLeaf_KEYS) GetLabelValue() uint32 {
	if m != nil {
		return m.LabelValue
	}
	return 0
}

func (m *MplsShLeaf_KEYS) GetEos() string {
	if m != nil {
		return m.Eos
	}
	return ""
}

type MplsShLeaf struct {
	// Local label
	LeafLocalLabel uint32 `protobuf:"varint,50,opt,name=leaf_local_label,json=leafLocalLabel" json:"leaf_local_label,omitempty"`
	// EOS bit
	EosBit uint32 `protobuf:"varint,51,opt,name=eos_bit,json=eosBit" json:"eos_bit,omitempty"`
	// Label-infos in FIB leaf
	LabelInformation []*MplsAdjInfo `protobuf:"bytes,52,rep,name=label_information,json=labelInformation" json:"label_information,omitempty"`
	// LDI-info in FIB leaf
	LdiInformation *MplsLdiInfo `protobuf:"bytes,53,opt,name=ldi_information,json=ldiInformation" json:"ldi_information,omitempty"`
	// Hardware info
	HardwareInformation []byte `protobuf:"bytes,54,opt,name=hardware_information,json=hardwareInformation,proto3" json:"hardware_information,omitempty"`
	// Number of references to the leaf
	LeafReferanceCount uint32 `protobuf:"varint,55,opt,name=leaf_referance_count,json=leafReferanceCount" json:"leaf_referance_count,omitempty"`
	// The leaf flags
	LeafFlags uint32 `protobuf:"varint,56,opt,name=leaf_flags,json=leafFlags" json:"leaf_flags,omitempty"`
	// Number of references to the pathlist
	PathListReferanceCount uint32 `protobuf:"varint,57,opt,name=path_list_referance_count,json=pathListReferanceCount" json:"path_list_referance_count,omitempty"`
	// The pathlist flags
	PathListFlags uint32 `protobuf:"varint,58,opt,name=path_list_flags,json=pathListFlags" json:"path_list_flags,omitempty"`
	// Number of references to the LDI
	LdiReferanceCount uint32 `protobuf:"varint,59,opt,name=ldi_referance_count,json=ldiReferanceCount" json:"ldi_referance_count,omitempty"`
	// The LDI flags
	LdiFlags uint32 `protobuf:"varint,60,opt,name=ldi_flags,json=ldiFlags" json:"ldi_flags,omitempty"`
	// The LDI type
	LdiType uint32 `protobuf:"varint,61,opt,name=ldi_type,json=ldiType" json:"ldi_type,omitempty"`
	// The pointer to the LDI
	LdiPointer uint32 `protobuf:"varint,62,opt,name=ldi_pointer,json=ldiPointer" json:"ldi_pointer,omitempty"`
	// The LW-LDI type
	LwLdiType uint32 `protobuf:"varint,63,opt,name=lw_ldi_type,json=lwLdiType" json:"lw_ldi_type,omitempty"`
	// The pointer to the LW-LDI
	LwLdiPointer uint32 `protobuf:"varint,64,opt,name=lw_ldi_pointer,json=lwLdiPointer" json:"lw_ldi_pointer,omitempty"`
	// The LW-LDI refcounter
	LwLdiRefernaceCount uint32 `protobuf:"varint,65,opt,name=lw_ldi_refernace_count,json=lwLdiRefernaceCount" json:"lw_ldi_refernace_count,omitempty"`
	// The pointer to the shared LDI in LW-LDI
	LwSharedLdiPointer uint32 `protobuf:"varint,66,opt,name=lw_shared_ldi_pointer,json=lwSharedLdiPointer" json:"lw_shared_ldi_pointer,omitempty"`
	// The LSPA flags
	LspaFlags uint32 `protobuf:"varint,67,opt,name=lspa_flags,json=lspaFlags" json:"lspa_flags,omitempty"`
	// The AFI table ID
	AfiTableId uint32 `protobuf:"varint,68,opt,name=afi_table_id,json=afiTableId" json:"afi_table_id,omitempty"`
	// The unicast or multicast label
	MulticastLabel bool `protobuf:"varint,69,opt,name=multicast_label,json=multicastLabel" json:"multicast_label,omitempty"`
	// The multicast info
	MulticastInformation *MplsMcastInfo `protobuf:"bytes,70,opt,name=multicast_information,json=multicastInformation" json:"multicast_information,omitempty"`
	// The time of last update in msec
	LeafTimeInMilliSeconds uint64 `protobuf:"varint,71,opt,name=leaf_time_in_milli_seconds,json=leafTimeInMilliSeconds" json:"leaf_time_in_milli_seconds,omitempty"`
}

func (m *MplsShLeaf) Reset()                    { *m = MplsShLeaf{} }
func (m *MplsShLeaf) String() string            { return proto.CompactTextString(m) }
func (*MplsShLeaf) ProtoMessage()               {}
func (*MplsShLeaf) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MplsShLeaf) GetLeafLocalLabel() uint32 {
	if m != nil {
		return m.LeafLocalLabel
	}
	return 0
}

func (m *MplsShLeaf) GetEosBit() uint32 {
	if m != nil {
		return m.EosBit
	}
	return 0
}

func (m *MplsShLeaf) GetLabelInformation() []*MplsAdjInfo {
	if m != nil {
		return m.LabelInformation
	}
	return nil
}

func (m *MplsShLeaf) GetLdiInformation() *MplsLdiInfo {
	if m != nil {
		return m.LdiInformation
	}
	return nil
}

func (m *MplsShLeaf) GetHardwareInformation() []byte {
	if m != nil {
		return m.HardwareInformation
	}
	return nil
}

func (m *MplsShLeaf) GetLeafReferanceCount() uint32 {
	if m != nil {
		return m.LeafReferanceCount
	}
	return 0
}

func (m *MplsShLeaf) GetLeafFlags() uint32 {
	if m != nil {
		return m.LeafFlags
	}
	return 0
}

func (m *MplsShLeaf) GetPathListReferanceCount() uint32 {
	if m != nil {
		return m.PathListReferanceCount
	}
	return 0
}

func (m *MplsShLeaf) GetPathListFlags() uint32 {
	if m != nil {
		return m.PathListFlags
	}
	return 0
}

func (m *MplsShLeaf) GetLdiReferanceCount() uint32 {
	if m != nil {
		return m.LdiReferanceCount
	}
	return 0
}

func (m *MplsShLeaf) GetLdiFlags() uint32 {
	if m != nil {
		return m.LdiFlags
	}
	return 0
}

func (m *MplsShLeaf) GetLdiType() uint32 {
	if m != nil {
		return m.LdiType
	}
	return 0
}

func (m *MplsShLeaf) GetLdiPointer() uint32 {
	if m != nil {
		return m.LdiPointer
	}
	return 0
}

func (m *MplsShLeaf) GetLwLdiType() uint32 {
	if m != nil {
		return m.LwLdiType
	}
	return 0
}

func (m *MplsShLeaf) GetLwLdiPointer() uint32 {
	if m != nil {
		return m.LwLdiPointer
	}
	return 0
}

func (m *MplsShLeaf) GetLwLdiRefernaceCount() uint32 {
	if m != nil {
		return m.LwLdiRefernaceCount
	}
	return 0
}

func (m *MplsShLeaf) GetLwSharedLdiPointer() uint32 {
	if m != nil {
		return m.LwSharedLdiPointer
	}
	return 0
}

func (m *MplsShLeaf) GetLspaFlags() uint32 {
	if m != nil {
		return m.LspaFlags
	}
	return 0
}

func (m *MplsShLeaf) GetAfiTableId() uint32 {
	if m != nil {
		return m.AfiTableId
	}
	return 0
}

func (m *MplsShLeaf) GetMulticastLabel() bool {
	if m != nil {
		return m.MulticastLabel
	}
	return false
}

func (m *MplsShLeaf) GetMulticastInformation() *MplsMcastInfo {
	if m != nil {
		return m.MulticastInformation
	}
	return nil
}

func (m *MplsShLeaf) GetLeafTimeInMilliSeconds() uint64 {
	if m != nil {
		return m.LeafTimeInMilliSeconds
	}
	return 0
}

type MplsFwdInfo struct {
	// L3 MTU
	L3Mtu uint32 `protobuf:"varint,1,opt,name=l3_mtu,json=l3Mtu" json:"l3_mtu,omitempty"`
	// Total encapsulation size: L2 + MPLS
	TotalEncapsulationSize uint32 `protobuf:"varint,2,opt,name=total_encapsulation_size,json=totalEncapsulationSize" json:"total_encapsulation_size,omitempty"`
	// Length of L2 encapsulation
	MacSize uint32 `protobuf:"varint,3,opt,name=mac_size,json=macSize" json:"mac_size,omitempty"`
	// Label stack
	LabelStack []uint32 `protobuf:"varint,4,rep,packed,name=label_stack,json=labelStack" json:"label_stack,omitempty"`
	// Number of packets switched
	TransmitNumberOfPacketsSwitched uint64 `protobuf:"varint,5,opt,name=transmit_number_of_packets_switched,json=transmitNumberOfPacketsSwitched" json:"transmit_number_of_packets_switched,omitempty"`
	// Number of Bytes switched
	TransmitNumberOfBytesSwitched uint64 `protobuf:"varint,6,opt,name=transmit_number_of_bytes_switched,json=transmitNumberOfBytesSwitched" json:"transmit_number_of_bytes_switched,omitempty"`
	// Status
	Status int32 `protobuf:"zigzag32,7,opt,name=status" json:"status,omitempty"`
	// Next hop interface
	NextHopInterface string `protobuf:"bytes,8,opt,name=next_hop_interface,json=nextHopInterface" json:"next_hop_interface,omitempty"`
	// The address family (V4/V6)
	NextHopProtocol string `protobuf:"bytes,9,opt,name=next_hop_protocol,json=nextHopProtocol" json:"next_hop_protocol,omitempty"`
	// Next hop address in string format
	NextHopString string `protobuf:"bytes,10,opt,name=next_hop_string,json=nextHopString" json:"next_hop_string,omitempty"`
}

func (m *MplsFwdInfo) Reset()                    { *m = MplsFwdInfo{} }
func (m *MplsFwdInfo) String() string            { return proto.CompactTextString(m) }
func (*MplsFwdInfo) ProtoMessage()               {}
func (*MplsFwdInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MplsFwdInfo) GetL3Mtu() uint32 {
	if m != nil {
		return m.L3Mtu
	}
	return 0
}

func (m *MplsFwdInfo) GetTotalEncapsulationSize() uint32 {
	if m != nil {
		return m.TotalEncapsulationSize
	}
	return 0
}

func (m *MplsFwdInfo) GetMacSize() uint32 {
	if m != nil {
		return m.MacSize
	}
	return 0
}

func (m *MplsFwdInfo) GetLabelStack() []uint32 {
	if m != nil {
		return m.LabelStack
	}
	return nil
}

func (m *MplsFwdInfo) GetTransmitNumberOfPacketsSwitched() uint64 {
	if m != nil {
		return m.TransmitNumberOfPacketsSwitched
	}
	return 0
}

func (m *MplsFwdInfo) GetTransmitNumberOfBytesSwitched() uint64 {
	if m != nil {
		return m.TransmitNumberOfBytesSwitched
	}
	return 0
}

func (m *MplsFwdInfo) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *MplsFwdInfo) GetNextHopInterface() string {
	if m != nil {
		return m.NextHopInterface
	}
	return ""
}

func (m *MplsFwdInfo) GetNextHopProtocol() string {
	if m != nil {
		return m.NextHopProtocol
	}
	return ""
}

func (m *MplsFwdInfo) GetNextHopString() string {
	if m != nil {
		return m.NextHopString
	}
	return ""
}

type MplsAdjInfo struct {
	// Label-Info type
	LabelInformationType uint32 `protobuf:"varint,1,opt,name=label_information_type,json=labelInformationType" json:"label_information_type,omitempty"`
	// Local label
	LocalLabel uint32 `protobuf:"varint,2,opt,name=local_label,json=localLabel" json:"local_label,omitempty"`
	// Outgoing label
	OutgoingLabel uint32 `protobuf:"varint,3,opt,name=outgoing_label,json=outgoingLabel" json:"outgoing_label,omitempty"`
	// MPLS Adjacency flags
	MplsAdjacencyFlags uint32 `protobuf:"varint,4,opt,name=mpls_adjacency_flags,json=mplsAdjacencyFlags" json:"mpls_adjacency_flags,omitempty"`
	// Tunnel id present?
	TunnelIdPresent bool `protobuf:"varint,5,opt,name=tunnel_id_present,json=tunnelIdPresent" json:"tunnel_id_present,omitempty"`
	// Outgoing interface
	OutgoingInterface string `protobuf:"bytes,6,opt,name=outgoing_interface,json=outgoingInterface" json:"outgoing_interface,omitempty"`
	// Outgoing Physical Interface
	OutgoingPhysicalInterface string `protobuf:"bytes,7,opt,name=outgoing_physical_interface,json=outgoingPhysicalInterface" json:"outgoing_physical_interface,omitempty"`
	// Tunnel Interface
	TunnelInterface string `protobuf:"bytes,8,opt,name=tunnel_interface,json=tunnelInterface" json:"tunnel_interface,omitempty"`
	// Detail label info
	LabelInformationDetail    *MplsFwdInfo `protobuf:"bytes,9,opt,name=label_information_detail,json=labelInformationDetail" json:"label_information_detail,omitempty"`
	LabelInformationPathIndex uint32       `protobuf:"varint,10,opt,name=label_information_path_index,json=labelInformationPathIndex" json:"label_information_path_index,omitempty"`
	// NHinfo Type
	LabelInformationNextHopType string `protobuf:"bytes,11,opt,name=label_information_next_hop_type,json=labelInformationNextHopType" json:"label_information_next_hop_type,omitempty"`
	// The address family (v4/v6)
	LabelInformationNextHopProtocol string `protobuf:"bytes,12,opt,name=label_information_next_hop_protocol,json=labelInformationNextHopProtocol" json:"label_information_next_hop_protocol,omitempty"`
	// Bytes transmitted per LSP
	TxBytes uint64 `protobuf:"varint,13,opt,name=tx_bytes,json=txBytes" json:"tx_bytes,omitempty"`
	// Packets transmitted per LSP
	TxPackets uint64 `protobuf:"varint,14,opt,name=tx_packets,json=txPackets" json:"tx_packets,omitempty"`
	// Output Interface in string format
	OutgoingInterfaceString string `protobuf:"bytes,15,opt,name=outgoing_interface_string,json=outgoingInterfaceString" json:"outgoing_interface_string,omitempty"`
	// Output Label in string format
	OutgoingLabelString string `protobuf:"bytes,16,opt,name=outgoing_label_string,json=outgoingLabelString" json:"outgoing_label_string,omitempty"`
	// Prefix Or ID
	PrefixOrId string `protobuf:"bytes,17,opt,name=prefix_or_id,json=prefixOrId" json:"prefix_or_id,omitempty"`
	// Next hop address in string format
	LabelInformationNextHopString string `protobuf:"bytes,18,opt,name=label_information_next_hop_string,json=labelInformationNextHopString" json:"label_information_next_hop_string,omitempty"`
	// The version of the route
	LabelInformationRouteVersion uint64 `protobuf:"varint,19,opt,name=label_information_route_version,json=labelInformationRouteVersion" json:"label_information_route_version,omitempty"`
	// The time of last update in msec
	LabelInformationTimeInMilliSeconds uint64 `protobuf:"varint,20,opt,name=label_information_time_in_milli_seconds,json=labelInformationTimeInMilliSeconds" json:"label_information_time_in_milli_seconds,omitempty"`
}

func (m *MplsAdjInfo) Reset()                    { *m = MplsAdjInfo{} }
func (m *MplsAdjInfo) String() string            { return proto.CompactTextString(m) }
func (*MplsAdjInfo) ProtoMessage()               {}
func (*MplsAdjInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *MplsAdjInfo) GetLabelInformationType() uint32 {
	if m != nil {
		return m.LabelInformationType
	}
	return 0
}

func (m *MplsAdjInfo) GetLocalLabel() uint32 {
	if m != nil {
		return m.LocalLabel
	}
	return 0
}

func (m *MplsAdjInfo) GetOutgoingLabel() uint32 {
	if m != nil {
		return m.OutgoingLabel
	}
	return 0
}

func (m *MplsAdjInfo) GetMplsAdjacencyFlags() uint32 {
	if m != nil {
		return m.MplsAdjacencyFlags
	}
	return 0
}

func (m *MplsAdjInfo) GetTunnelIdPresent() bool {
	if m != nil {
		return m.TunnelIdPresent
	}
	return false
}

func (m *MplsAdjInfo) GetOutgoingInterface() string {
	if m != nil {
		return m.OutgoingInterface
	}
	return ""
}

func (m *MplsAdjInfo) GetOutgoingPhysicalInterface() string {
	if m != nil {
		return m.OutgoingPhysicalInterface
	}
	return ""
}

func (m *MplsAdjInfo) GetTunnelInterface() string {
	if m != nil {
		return m.TunnelInterface
	}
	return ""
}

func (m *MplsAdjInfo) GetLabelInformationDetail() *MplsFwdInfo {
	if m != nil {
		return m.LabelInformationDetail
	}
	return nil
}

func (m *MplsAdjInfo) GetLabelInformationPathIndex() uint32 {
	if m != nil {
		return m.LabelInformationPathIndex
	}
	return 0
}

func (m *MplsAdjInfo) GetLabelInformationNextHopType() string {
	if m != nil {
		return m.LabelInformationNextHopType
	}
	return ""
}

func (m *MplsAdjInfo) GetLabelInformationNextHopProtocol() string {
	if m != nil {
		return m.LabelInformationNextHopProtocol
	}
	return ""
}

func (m *MplsAdjInfo) GetTxBytes() uint64 {
	if m != nil {
		return m.TxBytes
	}
	return 0
}

func (m *MplsAdjInfo) GetTxPackets() uint64 {
	if m != nil {
		return m.TxPackets
	}
	return 0
}

func (m *MplsAdjInfo) GetOutgoingInterfaceString() string {
	if m != nil {
		return m.OutgoingInterfaceString
	}
	return ""
}

func (m *MplsAdjInfo) GetOutgoingLabelString() string {
	if m != nil {
		return m.OutgoingLabelString
	}
	return ""
}

func (m *MplsAdjInfo) GetPrefixOrId() string {
	if m != nil {
		return m.PrefixOrId
	}
	return ""
}

func (m *MplsAdjInfo) GetLabelInformationNextHopString() string {
	if m != nil {
		return m.LabelInformationNextHopString
	}
	return ""
}

func (m *MplsAdjInfo) GetLabelInformationRouteVersion() uint64 {
	if m != nil {
		return m.LabelInformationRouteVersion
	}
	return 0
}

func (m *MplsAdjInfo) GetLabelInformationTimeInMilliSeconds() uint64 {
	if m != nil {
		return m.LabelInformationTimeInMilliSeconds
	}
	return 0
}

// Detailed load sharing information for mpls table entries
type MplsLdiInfo struct {
	// Hardware info
	LdiHardwareInformation []byte `protobuf:"bytes,1,opt,name=ldi_hardware_information,json=ldiHardwareInformation,proto3" json:"ldi_hardware_information,omitempty"`
}

func (m *MplsLdiInfo) Reset()                    { *m = MplsLdiInfo{} }
func (m *MplsLdiInfo) String() string            { return proto.CompactTextString(m) }
func (*MplsLdiInfo) ProtoMessage()               {}
func (*MplsLdiInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *MplsLdiInfo) GetLdiHardwareInformation() []byte {
	if m != nil {
		return m.LdiHardwareInformation
	}
	return nil
}

// Information for mpls multicast entries
type MplsMcastInfo struct {
	// MOL base flags
	MulticastMolBaseFlags uint32 `protobuf:"varint,1,opt,name=multicast_mol_base_flags,json=multicastMolBaseFlags" json:"multicast_mol_base_flags,omitempty"`
	// MOL flags
	MulticastMolFlags uint32 `protobuf:"varint,2,opt,name=multicast_mol_flags,json=multicastMolFlags" json:"multicast_mol_flags,omitempty"`
	// MOL refcount
	MulticastMolReferanceCount uint32 `protobuf:"varint,3,opt,name=multicast_mol_referance_count,json=multicastMolReferanceCount" json:"multicast_mol_referance_count,omitempty"`
	// multicast mpls tunnel
	MulticastTunnelInterfaceHandler string `protobuf:"bytes,4,opt,name=multicast_tunnel_interface_handler,json=multicastTunnelInterfaceHandler" json:"multicast_tunnel_interface_handler,omitempty"`
	// multicast mpls P2MP-TE tunnel id or MLDP Tunnel LSMID on all nodes
	MulticastTunnelId uint32 `protobuf:"varint,5,opt,name=multicast_tunnel_id,json=multicastTunnelId" json:"multicast_tunnel_id,omitempty"`
	// multicast nhinfo for p2mp TE Head
	MulticastTunnelNextHopInformation uint32 `protobuf:"varint,6,opt,name=multicast_tunnel_next_hop_information,json=multicastTunnelNextHopInformation" json:"multicast_tunnel_next_hop_information,omitempty"`
	// multicast LSPVIF for MLDP Tunnels
	MulticastTunnelLspvif uint32 `protobuf:"varint,7,opt,name=multicast_tunnel_lspvif,json=multicastTunnelLspvif" json:"multicast_tunnel_lspvif,omitempty"`
	// num multicast mpls output paths
	MulticastMplsOutputPaths uint32 `protobuf:"varint,8,opt,name=multicast_mpls_output_paths,json=multicastMplsOutputPaths" json:"multicast_mpls_output_paths,omitempty"`
	// num multicast mpls prot output paths
	MulticastMplsProtocolOutputPaths uint32 `protobuf:"varint,9,opt,name=multicast_mpls_protocol_output_paths,json=multicastMplsProtocolOutputPaths" json:"multicast_mpls_protocol_output_paths,omitempty"`
	// num multicast mpls local output paths
	MulticastMplsLocalOutputPaths uint32 `protobuf:"varint,10,opt,name=multicast_mpls_local_output_paths,json=multicastMplsLocalOutputPaths" json:"multicast_mpls_local_output_paths,omitempty"`
	// The multicast RPF-ID
	MulticastRpfId uint32 `protobuf:"varint,11,opt,name=multicast_rpf_id,json=multicastRpfId" json:"multicast_rpf_id,omitempty"`
	// The multicast ENCAP-ID
	MulticastEncapId uint32 `protobuf:"varint,12,opt,name=multicast_encap_id,json=multicastEncapId" json:"multicast_encap_id,omitempty"`
	// The multicast platform data len
	MulticastPlatformDataLength uint32 `protobuf:"varint,13,opt,name=multicast_platform_data_length,json=multicastPlatformDataLength" json:"multicast_platform_data_length,omitempty"`
	// The multicast platform data
	MulticastPlatformData []byte `protobuf:"bytes,14,opt,name=multicast_platform_data,json=multicastPlatformData,proto3" json:"multicast_platform_data,omitempty"`
}

func (m *MplsMcastInfo) Reset()                    { *m = MplsMcastInfo{} }
func (m *MplsMcastInfo) String() string            { return proto.CompactTextString(m) }
func (*MplsMcastInfo) ProtoMessage()               {}
func (*MplsMcastInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *MplsMcastInfo) GetMulticastMolBaseFlags() uint32 {
	if m != nil {
		return m.MulticastMolBaseFlags
	}
	return 0
}

func (m *MplsMcastInfo) GetMulticastMolFlags() uint32 {
	if m != nil {
		return m.MulticastMolFlags
	}
	return 0
}

func (m *MplsMcastInfo) GetMulticastMolReferanceCount() uint32 {
	if m != nil {
		return m.MulticastMolReferanceCount
	}
	return 0
}

func (m *MplsMcastInfo) GetMulticastTunnelInterfaceHandler() string {
	if m != nil {
		return m.MulticastTunnelInterfaceHandler
	}
	return ""
}

func (m *MplsMcastInfo) GetMulticastTunnelId() uint32 {
	if m != nil {
		return m.MulticastTunnelId
	}
	return 0
}

func (m *MplsMcastInfo) GetMulticastTunnelNextHopInformation() uint32 {
	if m != nil {
		return m.MulticastTunnelNextHopInformation
	}
	return 0
}

func (m *MplsMcastInfo) GetMulticastTunnelLspvif() uint32 {
	if m != nil {
		return m.MulticastTunnelLspvif
	}
	return 0
}

func (m *MplsMcastInfo) GetMulticastMplsOutputPaths() uint32 {
	if m != nil {
		return m.MulticastMplsOutputPaths
	}
	return 0
}

func (m *MplsMcastInfo) GetMulticastMplsProtocolOutputPaths() uint32 {
	if m != nil {
		return m.MulticastMplsProtocolOutputPaths
	}
	return 0
}

func (m *MplsMcastInfo) GetMulticastMplsLocalOutputPaths() uint32 {
	if m != nil {
		return m.MulticastMplsLocalOutputPaths
	}
	return 0
}

func (m *MplsMcastInfo) GetMulticastRpfId() uint32 {
	if m != nil {
		return m.MulticastRpfId
	}
	return 0
}

func (m *MplsMcastInfo) GetMulticastEncapId() uint32 {
	if m != nil {
		return m.MulticastEncapId
	}
	return 0
}

func (m *MplsMcastInfo) GetMulticastPlatformDataLength() uint32 {
	if m != nil {
		return m.MulticastPlatformDataLength
	}
	return 0
}

func (m *MplsMcastInfo) GetMulticastPlatformData() []byte {
	if m != nil {
		return m.MulticastPlatformData
	}
	return nil
}

func init() {
	proto.RegisterType((*MplsShLeaf_KEYS)(nil), "cisco_ios_xr_fib_common_oper.mpls_forwarding.nodes.node.label_fib.forwarding_details.forwarding_detail.mpls_sh_leaf_KEYS")
	proto.RegisterType((*MplsShLeaf)(nil), "cisco_ios_xr_fib_common_oper.mpls_forwarding.nodes.node.label_fib.forwarding_details.forwarding_detail.mpls_sh_leaf")
	proto.RegisterType((*MplsFwdInfo)(nil), "cisco_ios_xr_fib_common_oper.mpls_forwarding.nodes.node.label_fib.forwarding_details.forwarding_detail.mpls_fwd_info")
	proto.RegisterType((*MplsAdjInfo)(nil), "cisco_ios_xr_fib_common_oper.mpls_forwarding.nodes.node.label_fib.forwarding_details.forwarding_detail.mpls_adj_info")
	proto.RegisterType((*MplsLdiInfo)(nil), "cisco_ios_xr_fib_common_oper.mpls_forwarding.nodes.node.label_fib.forwarding_details.forwarding_detail.mpls_ldi_info")
	proto.RegisterType((*MplsMcastInfo)(nil), "cisco_ios_xr_fib_common_oper.mpls_forwarding.nodes.node.label_fib.forwarding_details.forwarding_detail.mpls_mcast_info")
}

func init() { proto.RegisterFile("mpls_sh_leaf.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1560 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x57, 0xdb, 0x6e, 0x23, 0xc7,
	0x11, 0x05, 0xa3, 0x35, 0x25, 0x96, 0x48, 0x49, 0x6c, 0x69, 0xb9, 0x23, 0xcb, 0xb2, 0xb8, 0xb2,
	0x1d, 0x33, 0x46, 0x22, 0xd8, 0x2b, 0xc7, 0x97, 0x4d, 0x6c, 0x67, 0x6f, 0x8e, 0x08, 0x53, 0x5a,
	0x62, 0x28, 0x18, 0xc8, 0x53, 0xa3, 0x39, 0xd3, 0x23, 0x76, 0xdc, 0x73, 0xc1, 0x74, 0x73, 0x49,
	0xf9, 0x13, 0xf2, 0x11, 0x79, 0xc8, 0x73, 0x80, 0xfc, 0x41, 0x3e, 0x21, 0x1f, 0x93, 0x2f, 0x30,
	0xba, 0x7a, 0x6e, 0x1c, 0xca, 0xfb, 0xa8, 0x97, 0x05, 0xb7, 0xea, 0xd4, 0xe9, 0x9a, 0xaa, 0xae,
	0x53, 0x2d, 0x20, 0x61, 0x22, 0x15, 0x55, 0x33, 0x2a, 0x39, 0x0b, 0xce, 0x92, 0x34, 0xd6, 0x31,
	0x09, 0x3c, 0xa1, 0xbc, 0x98, 0x8a, 0x58, 0xd1, 0x65, 0x4a, 0x03, 0x31, 0xa5, 0x5e, 0x1c, 0x86,
	0x71, 0x44, 0xe3, 0x84, 0xa7, 0x67, 0x18, 0x10, 0xc4, 0xe9, 0x82, 0xa5, 0xbe, 0x88, 0x6e, 0xce,
	0xa2, 0xd8, 0xe7, 0x0a, 0xff, 0x3d, 0x93, 0x6c, 0xca, 0xa5, 0x09, 0x38, 0x2b, 0xfd, 0xd4, 0xe7,
	0x9a, 0x09, 0xa9, 0xd6, 0x4d, 0xa7, 0x1e, 0x74, 0xab, 0xa7, 0xd3, 0x1f, 0x5e, 0xfd, 0x6d, 0x42,
	0x8e, 0xa0, 0x65, 0xb8, 0x68, 0xc4, 0x42, 0xee, 0x34, 0xfa, 0x8d, 0x41, 0xcb, 0xdd, 0x32, 0x86,
	0x2b, 0x16, 0x72, 0x72, 0x02, 0xdb, 0xf6, 0x8c, 0x37, 0x4c, 0xce, 0xb9, 0xf3, 0x9b, 0x7e, 0x63,
	0xd0, 0x71, 0x01, 0x4d, 0x3f, 0x1a, 0x0b, 0xd9, 0x83, 0x0d, 0x1e, 0x2b, 0x67, 0x03, 0xe3, 0xcc,
	0xcf, 0xd3, 0x7f, 0x00, 0xb4, 0xab, 0xa7, 0x90, 0x01, 0xec, 0xe1, 0x69, 0x32, 0xf6, 0x98, 0xa4,
	0x18, 0xeb, 0x3c, 0x41, 0xa2, 0x1d, 0x63, 0x1f, 0x19, 0xf3, 0xc8, 0x58, 0xc9, 0x23, 0xd8, 0xe4,
	0xb1, 0xa2, 0x53, 0xa1, 0x9d, 0x73, 0x04, 0x34, 0x79, 0xac, 0x9e, 0x0b, 0x4d, 0xfe, 0xd5, 0x80,
	0xae, 0xcd, 0x43, 0x44, 0x41, 0x9c, 0x86, 0x4c, 0x8b, 0x38, 0x72, 0x3e, 0xef, 0x6f, 0x0c, 0xb6,
	0x9f, 0xcc, 0xcf, 0xee, 0xa7, 0x7a, 0x96, 0x89, 0xf9, 0x7f, 0xc7, 0x1c, 0xdc, 0x3d, 0x8c, 0x1e,
	0x96, 0xe9, 0x90, 0x7f, 0x36, 0x60, 0x57, 0xfa, 0x62, 0x25, 0xc5, 0x3f, 0xf6, 0x1b, 0xf7, 0x9e,
	0x62, 0x9e, 0x83, 0xbb, 0x23, 0x7d, 0x51, 0x4d, 0xf0, 0x33, 0x38, 0x98, 0xb1, 0xd4, 0x5f, 0xb0,
	0x94, 0xaf, 0x24, 0xf9, 0x45, 0xbf, 0x31, 0x68, 0xbb, 0xfb, 0xb9, 0xaf, 0x1a, 0xf2, 0x29, 0x1c,
	0x60, 0xef, 0x52, 0x1e, 0xf0, 0x94, 0x45, 0x1e, 0xa7, 0x5e, 0x3c, 0x8f, 0xb4, 0xf3, 0x25, 0xb6,
	0x87, 0x18, 0x9f, 0x9b, 0xbb, 0x5e, 0x18, 0x0f, 0x39, 0x06, 0xc0, 0x88, 0x40, 0xb2, 0x1b, 0xe5,
	0x7c, 0x85, 0xb8, 0x96, 0xb1, 0x7c, 0x6f, 0x0c, 0xe4, 0x6b, 0x38, 0x4c, 0x98, 0x9e, 0x51, 0x29,
	0x94, 0x5e, 0x63, 0xfd, 0x1a, 0xd1, 0x3d, 0x03, 0x18, 0x09, 0xa5, 0x6b, 0xcc, 0xbf, 0x85, 0xdd,
	0x32, 0xd4, 0xd2, 0x3f, 0xc5, 0x80, 0x4e, 0x1e, 0x60, 0x8f, 0x38, 0x83, 0x7d, 0x53, 0x82, 0x3a,
	0xf9, 0x9f, 0x10, 0xdb, 0x95, 0xbe, 0xa8, 0xf1, 0x1e, 0x41, 0xcb, 0xe0, 0x2d, 0xe3, 0x9f, 0x11,
	0xb5, 0x25, 0x7d, 0x61, 0xc9, 0x0e, 0xc1, 0xfc, 0xa6, 0xfa, 0x36, 0xe1, 0xce, 0x37, 0xe8, 0xdb,
	0x94, 0xbe, 0xb8, 0xbe, 0x4d, 0xec, 0x6c, 0xf8, 0x82, 0x26, 0xb1, 0x88, 0x34, 0x4f, 0x9d, 0x6f,
	0xb3, 0xd9, 0xf0, 0xc5, 0xd8, 0x5a, 0xc8, 0xfb, 0xb0, 0x2d, 0x17, 0xb4, 0x08, 0xff, 0x2e, 0xab,
	0xc5, 0x62, 0x94, 0x11, 0x7c, 0x08, 0x3b, 0x99, 0x3f, 0xe7, 0xf8, 0x0b, 0x42, 0xda, 0x08, 0xc9,
	0x59, 0xce, 0xa1, 0x97, 0xa1, 0xf0, 0x8b, 0x22, 0x56, 0x7c, 0xd1, 0x33, 0x44, 0xef, 0x23, 0xda,
	0xcd, 0x7d, 0xf6, 0x9b, 0x3e, 0x83, 0x87, 0x72, 0x41, 0xd5, 0x8c, 0xa5, 0xdc, 0x5f, 0x39, 0xe1,
	0x79, 0xd6, 0xb8, 0xc5, 0x04, 0x7d, 0x95, 0x73, 0x4c, 0xe3, 0x54, 0xc2, 0xb2, 0x3a, 0xbc, 0xc8,
	0x92, 0x55, 0x09, 0xb3, 0x85, 0xe8, 0x43, 0x9b, 0x05, 0x82, 0x6a, 0x36, 0x95, 0x9c, 0x0a, 0xdf,
	0x79, 0x69, 0x3f, 0x97, 0x05, 0xe2, 0xda, 0x98, 0x86, 0x3e, 0xf9, 0x18, 0x76, 0xc3, 0xb9, 0xd4,
	0xc2, 0x63, 0x4a, 0x67, 0x63, 0xfe, 0xaa, 0xdf, 0x18, 0x6c, 0xb9, 0x3b, 0x85, 0xd9, 0x8e, 0xf9,
	0xbf, 0x1b, 0xf0, 0xb0, 0x44, 0x56, 0x6f, 0xe2, 0xf7, 0x38, 0x2e, 0x8b, 0x7b, 0x1d, 0x97, 0xb0,
	0xc8, 0xc2, 0x3d, 0x28, 0xb2, 0xaa, 0xce, 0xc0, 0x53, 0x78, 0x17, 0x6f, 0xb4, 0x16, 0xa1, 0x99,
	0x1b, 0x1a, 0x0a, 0x29, 0x05, 0x55, 0xdc, 0x8b, 0x23, 0x5f, 0x39, 0x7f, 0xed, 0x37, 0x06, 0x0f,
	0xdc, 0x9e, 0x41, 0x5c, 0x8b, 0x90, 0x0f, 0xa3, 0x4b, 0xe3, 0x9e, 0x58, 0xef, 0xe9, 0xff, 0x36,
	0xa0, 0x63, 0xf3, 0x5d, 0xf8, 0x78, 0x06, 0x79, 0x08, 0x4d, 0x79, 0x4e, 0x43, 0x3d, 0x47, 0xad,
	0xed, 0xb8, 0xef, 0xc8, 0xf3, 0x4b, 0x3d, 0x27, 0x5f, 0x81, 0xa3, 0x63, 0xcd, 0x24, 0xe5, 0x91,
	0xc7, 0x12, 0x35, 0x97, 0x78, 0x36, 0x55, 0xe2, 0xe7, 0x5c, 0x75, 0x7b, 0xe8, 0x7f, 0x55, 0x75,
	0x4f, 0xc4, 0xcf, 0xdc, 0xdc, 0xd0, 0x90, 0x79, 0x16, 0xb9, 0x61, 0x6f, 0x68, 0xc8, 0x3c, 0x74,
	0x15, 0xea, 0xad, 0x34, 0xf3, 0x7e, 0x72, 0x1e, 0xf4, 0x37, 0x0a, 0xf5, 0x9e, 0x18, 0x0b, 0x19,
	0xc1, 0x07, 0x3a, 0x65, 0x91, 0x0a, 0x85, 0xa6, 0xd1, 0x3c, 0x9c, 0xf2, 0x94, 0xc6, 0x01, 0x4d,
	0x98, 0xf7, 0x13, 0xd7, 0x8a, 0xaa, 0x85, 0xd0, 0xde, 0x8c, 0xfb, 0xce, 0x3b, 0xf8, 0x8d, 0x27,
	0x39, 0xf4, 0x0a, 0x91, 0xaf, 0x83, 0xb1, 0xc5, 0x4d, 0x32, 0x18, 0xb9, 0x80, 0xc7, 0x77, 0xb0,
	0x4d, 0x6f, 0x35, 0xaf, 0x70, 0x35, 0x91, 0xeb, 0xb8, 0xce, 0xf5, 0xdc, 0xa0, 0x0a, 0xa6, 0x1e,
	0x34, 0x95, 0x66, 0x7a, 0xae, 0x9c, 0xcd, 0x7e, 0x63, 0xd0, 0x75, 0xb3, 0xff, 0x91, 0xdf, 0x03,
	0x89, 0xf8, 0x52, 0xd3, 0x59, 0x9c, 0x50, 0xbc, 0xb5, 0x01, 0xf3, 0xb8, 0xb3, 0x85, 0xcb, 0x67,
	0xcf, 0x78, 0x2e, 0xe2, 0x64, 0x98, 0xdb, 0xc9, 0x27, 0xd0, 0x2d, 0xd0, 0xb8, 0x68, 0xbd, 0x58,
	0x3a, 0x2d, 0x04, 0xef, 0x66, 0xe0, 0x71, 0x66, 0x36, 0xe2, 0x52, 0x60, 0x95, 0x4e, 0x45, 0x74,
	0xe3, 0x00, 0x22, 0x3b, 0x19, 0x72, 0x82, 0xc6, 0xd3, 0xff, 0xb6, 0xb2, 0x86, 0xe6, 0x8b, 0x80,
	0x7c, 0x0e, 0xbd, 0xb5, 0xd5, 0x64, 0x07, 0xde, 0x36, 0xf8, 0xa0, 0xbe, 0x28, 0x0a, 0xf1, 0xa8,
	0xec, 0xc3, 0x7c, 0xb1, 0x96, 0xbb, 0xf0, 0x23, 0xd8, 0x89, 0xe7, 0xfa, 0x26, 0x36, 0xd7, 0xd5,
	0x62, 0x6c, 0x73, 0x3b, 0xb9, 0xd5, 0xc2, 0x3e, 0x85, 0x83, 0x3c, 0x1d, 0xe6, 0xf1, 0xc8, 0xbb,
	0xcd, 0xe6, 0xf7, 0x81, 0x9d, 0x73, 0xe3, 0x7b, 0x96, 0xbb, 0xec, 0x20, 0x7f, 0x02, 0x5d, 0x3d,
	0x8f, 0x22, 0x93, 0xb0, 0x4f, 0x93, 0x94, 0x2b, 0x1e, 0x69, 0xec, 0xf0, 0x96, 0xbb, 0x6b, 0x1d,
	0x43, 0x7f, 0x6c, 0xcd, 0xe4, 0x0f, 0x40, 0x8a, 0x24, 0xca, 0x7a, 0x37, 0xb1, 0x30, 0xdd, 0xdc,
	0x53, 0x16, 0xfc, 0x5b, 0x38, 0x2a, 0xe0, 0xc9, 0xec, 0x56, 0x09, 0xf3, 0x81, 0x65, 0xdc, 0x26,
	0xc6, 0x1d, 0xe6, 0x90, 0x71, 0x86, 0x28, 0xe3, 0x7f, 0x07, 0x7b, 0x79, 0x6a, 0xb5, 0xe6, 0xe6,
	0x99, 0x15, 0xd0, 0xff, 0x34, 0xc0, 0x59, 0x2f, 0xbb, 0x9d, 0x6b, 0xec, 0xf1, 0x7d, 0x6f, 0xdd,
	0x7c, 0xc0, 0xdd, 0x5e, 0xbd, 0xdf, 0x2f, 0x11, 0x45, 0xbe, 0x83, 0xf7, 0xd6, 0x13, 0xc6, 0x85,
	0x26, 0x22, 0x9f, 0x2f, 0xf1, 0xba, 0x75, 0xdc, 0xc3, 0x7a, 0xf4, 0x98, 0xe9, 0xd9, 0xd0, 0x00,
	0xc8, 0x4b, 0x38, 0x59, 0x27, 0x28, 0x2e, 0x2d, 0xde, 0xb8, 0x6d, 0x2c, 0xd6, 0x51, 0x9d, 0xe3,
	0xca, 0x5e, 0x61, 0xbc, 0x78, 0x23, 0xf8, 0xe0, 0x2d, 0x2c, 0xc5, 0x98, 0xb4, 0x91, 0xe9, 0xe4,
	0x57, 0x98, 0x8a, 0xb1, 0x39, 0x84, 0x2d, 0xbd, 0xb4, 0x23, 0xee, 0x74, 0x70, 0xb2, 0x37, 0xf5,
	0x12, 0x67, 0xd9, 0xec, 0x13, 0xbd, 0xcc, 0xb5, 0xc4, 0xd9, 0x41, 0x67, 0x4b, 0x2f, 0x33, 0xd1,
	0x20, 0x4f, 0xe1, 0x70, 0xfd, 0x6a, 0xe5, 0xa3, 0xb7, 0x8b, 0xa7, 0x3f, 0x5a, 0xbb, 0x61, 0x76,
	0x08, 0xc9, 0x13, 0x78, 0xb8, 0x3a, 0x1b, 0x79, 0xdc, 0x1e, 0xc6, 0xed, 0xaf, 0x8c, 0x48, 0x16,
	0xd3, 0x87, 0x76, 0x92, 0xf2, 0x40, 0x2c, 0x69, 0x9c, 0x9a, 0xfd, 0xd5, 0x45, 0x28, 0x58, 0xdb,
	0xeb, 0x74, 0x88, 0xf2, 0xf5, 0x96, 0xca, 0x64, 0x27, 0x10, 0x0c, 0x3b, 0xfe, 0x95, 0xba, 0x64,
	0x67, 0xbd, 0xba, 0xab, 0x53, 0x69, 0x3c, 0xd7, 0x9c, 0xbe, 0xe1, 0xa9, 0x32, 0x9b, 0x6e, 0x1f,
	0xeb, 0xf1, 0x5e, 0x9d, 0xc7, 0x35, 0xa0, 0x1f, 0x2d, 0x86, 0x4c, 0xe0, 0xe3, 0x3b, 0x94, 0xe5,
	0xce, 0x2d, 0x74, 0x80, 0x74, 0xa7, 0x6b, 0x52, 0xb3, 0xbe, 0x91, 0x86, 0x99, 0x7e, 0xe5, 0xaf,
	0x44, 0xb3, 0x79, 0xcc, 0xef, 0x3b, 0x5f, 0x86, 0x0d, 0x7c, 0x19, 0xf6, 0xa4, 0x2f, 0x2e, 0xd6,
	0x1f, 0x87, 0xa7, 0xff, 0x6f, 0xc2, 0x6e, 0x6d, 0x85, 0x92, 0x2f, 0xc1, 0x29, 0x57, 0x7b, 0x18,
	0x4b, 0x3a, 0x65, 0x8a, 0x67, 0x9a, 0x64, 0xf5, 0xb0, 0x5c, 0xfd, 0x97, 0xb1, 0x7c, 0xce, 0x14,
	0x2f, 0x5e, 0x6d, 0xab, 0x81, 0x36, 0xc6, 0x0a, 0x63, 0xb7, 0x1a, 0x63, 0xf1, 0xcf, 0xe0, 0x78,
	0x15, 0x5f, 0x7f, 0xef, 0x59, 0xb9, 0x7c, 0xb7, 0x1a, 0x59, 0x7b, 0xf8, 0xfd, 0x00, 0xa7, 0x25,
	0x45, 0x5d, 0x78, 0xe8, 0x8c, 0x45, 0xbe, 0xe4, 0x29, 0x2a, 0x69, 0xcb, 0x3d, 0x29, 0x90, 0xd7,
	0xab, 0x4a, 0x74, 0x61, 0x61, 0xab, 0xf9, 0x17, 0x02, 0x8b, 0xc2, 0x5a, 0xcd, 0x3f, 0x8b, 0xf6,
	0xc9, 0x18, 0x3e, 0x5a, 0xc3, 0x57, 0x76, 0x5b, 0xd9, 0x83, 0x26, 0x32, 0x3c, 0xae, 0x31, 0x5c,
	0xe5, 0xcb, 0xae, 0x7c, 0xa7, 0x7c, 0x01, 0x8f, 0xd6, 0x18, 0xa5, 0x4a, 0xde, 0x88, 0x00, 0x95,
	0xb7, 0x5a, 0x79, 0xcb, 0x31, 0x42, 0x27, 0xf9, 0x06, 0x8e, 0x2a, 0x95, 0x34, 0xfd, 0x8c, 0xe7,
	0x3a, 0x99, 0x6b, 0x14, 0x27, 0x85, 0x02, 0xdc, 0x71, 0xcb, 0xae, 0x5e, 0x26, 0x52, 0xbd, 0x46,
	0x80, 0x91, 0x26, 0x45, 0xae, 0xe0, 0xc3, 0x5a, 0x78, 0x2e, 0x22, 0xab, 0x3c, 0x2d, 0xe4, 0xe9,
	0xaf, 0xf0, 0xe4, 0x3a, 0x52, 0xe5, 0xbb, 0x80, 0xc7, 0x35, 0x3e, 0xbb, 0x28, 0x57, 0xc8, 0xac,
	0x58, 0x1e, 0xaf, 0x90, 0xe1, 0x1f, 0x92, 0x55, 0xa6, 0x01, 0xec, 0x95, 0x4c, 0x69, 0x12, 0x98,
	0x7e, 0x6c, 0xdb, 0x3f, 0x3c, 0x0b, 0xbb, 0x9b, 0x04, 0x43, 0xdf, 0xbc, 0x2b, 0x4a, 0x24, 0xbe,
	0xc0, 0x0c, 0xb6, 0x8d, 0xd8, 0x92, 0x03, 0xdf, 0x5e, 0x43, 0x9f, 0xbc, 0x80, 0xf7, 0x4b, 0x74,
	0x22, 0x99, 0x36, 0x3d, 0xa0, 0x3e, 0xd3, 0x8c, 0x4a, 0x1e, 0xdd, 0xe8, 0x19, 0x4a, 0x61, 0xc7,
	0x2d, 0xcb, 0x3a, 0xce, 0x40, 0x2f, 0x99, 0x66, 0x23, 0x84, 0xac, 0x76, 0x6b, 0x85, 0x04, 0xb5,
	0xb2, 0x5d, 0xe9, 0x56, 0x35, 0x7a, 0xda, 0xc4, 0xea, 0x9e, 0xff, 0x12, 0x00, 0x00, 0xff, 0xff,
	0xdf, 0x71, 0x6e, 0x08, 0x48, 0x10, 0x00, 0x00,
}
