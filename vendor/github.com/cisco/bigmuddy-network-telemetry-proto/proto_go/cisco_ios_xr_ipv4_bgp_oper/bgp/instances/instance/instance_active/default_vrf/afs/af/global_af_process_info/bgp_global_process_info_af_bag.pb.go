// Code generated by protoc-gen-go.
// source: bgp_global_process_info_af_bag.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_active_default_vrf_afs_af_global_af_process_info is a generated protocol buffer package.

It is generated from these files:
	bgp_global_process_info_af_bag.proto

It has these top-level messages:
	BgpGlobalProcessInfoAfBag_KEYS
	BgpGlobalProcessInfoAfBag
	BgpTimespec
	BgpRibInstallTimeInfo
	BgpGlobalProcessInfoAfGbl_
	BgpGlobalProcessInfoAfVrf_
*/
package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_active_default_vrf_afs_af_global_af_process_info

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

// BGP address family specific information common to all BGP processes
type BgpGlobalProcessInfoAfBag_KEYS struct {
	InstanceName string `protobuf:"bytes,1,opt,name=instance_name,json=instanceName" json:"instance_name,omitempty"`
	AfName       string `protobuf:"bytes,2,opt,name=af_name,json=afName" json:"af_name,omitempty"`
}

func (m *BgpGlobalProcessInfoAfBag_KEYS) Reset()                    { *m = BgpGlobalProcessInfoAfBag_KEYS{} }
func (m *BgpGlobalProcessInfoAfBag_KEYS) String() string            { return proto.CompactTextString(m) }
func (*BgpGlobalProcessInfoAfBag_KEYS) ProtoMessage()               {}
func (*BgpGlobalProcessInfoAfBag_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BgpGlobalProcessInfoAfBag_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *BgpGlobalProcessInfoAfBag_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

type BgpGlobalProcessInfoAfBag struct {
	// Name of the VRF
	VrfName string `protobuf:"bytes,50,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
	// Address family id
	AfName string `protobuf:"bytes,51,opt,name=af_name,json=afName" json:"af_name,omitempty"`
	// Non-stop routing enabled
	IsNsrEnabled bool `protobuf:"varint,52,opt,name=is_nsr_enabled,json=isNsrEnabled" json:"is_nsr_enabled,omitempty"`
	// Global information
	Global *BgpGlobalProcessInfoAfGbl_ `protobuf:"bytes,53,opt,name=global" json:"global,omitempty"`
	// VRF information
	Vrf *BgpGlobalProcessInfoAfVrf_ `protobuf:"bytes,54,opt,name=vrf" json:"vrf,omitempty"`
}

func (m *BgpGlobalProcessInfoAfBag) Reset()                    { *m = BgpGlobalProcessInfoAfBag{} }
func (m *BgpGlobalProcessInfoAfBag) String() string            { return proto.CompactTextString(m) }
func (*BgpGlobalProcessInfoAfBag) ProtoMessage()               {}
func (*BgpGlobalProcessInfoAfBag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BgpGlobalProcessInfoAfBag) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *BgpGlobalProcessInfoAfBag) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *BgpGlobalProcessInfoAfBag) GetIsNsrEnabled() bool {
	if m != nil {
		return m.IsNsrEnabled
	}
	return false
}

func (m *BgpGlobalProcessInfoAfBag) GetGlobal() *BgpGlobalProcessInfoAfGbl_ {
	if m != nil {
		return m.Global
	}
	return nil
}

func (m *BgpGlobalProcessInfoAfBag) GetVrf() *BgpGlobalProcessInfoAfVrf_ {
	if m != nil {
		return m.Vrf
	}
	return nil
}

type BgpTimespec struct {
	// Seconds part of time value
	Seconds uint32 `protobuf:"varint,1,opt,name=seconds" json:"seconds,omitempty"`
	// Nanoseconds part of time value
	Nanoseconds uint32 `protobuf:"varint,2,opt,name=nanoseconds" json:"nanoseconds,omitempty"`
}

func (m *BgpTimespec) Reset()                    { *m = BgpTimespec{} }
func (m *BgpTimespec) String() string            { return proto.CompactTextString(m) }
func (*BgpTimespec) ProtoMessage()               {}
func (*BgpTimespec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *BgpTimespec) GetSeconds() uint32 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *BgpTimespec) GetNanoseconds() uint32 {
	if m != nil {
		return m.Nanoseconds
	}
	return 0
}

type BgpRibInstallTimeInfo struct {
	// RIB update time
	UpdateTime *BgpTimespec `protobuf:"bytes,1,opt,name=update_time,json=updateTime" json:"update_time,omitempty"`
	// RIB install time
	InstallTime *BgpTimespec `protobuf:"bytes,2,opt,name=install_time,json=installTime" json:"install_time,omitempty"`
	// Installed routes
	InstalledCount uint32 `protobuf:"varint,3,opt,name=installed_count,json=installedCount" json:"installed_count,omitempty"`
	// Modified routes
	ModifiedCount uint32 `protobuf:"varint,4,opt,name=modified_count,json=modifiedCount" json:"modified_count,omitempty"`
	// Withdrawn routes
	WithdrawnCount uint32 `protobuf:"varint,5,opt,name=withdrawn_count,json=withdrawnCount" json:"withdrawn_count,omitempty"`
	// Start version
	StartVersion uint32 `protobuf:"varint,6,opt,name=start_version,json=startVersion" json:"start_version,omitempty"`
	// Target version
	TargetVersion uint32 `protobuf:"varint,7,opt,name=target_version,json=targetVersion" json:"target_version,omitempty"`
}

func (m *BgpRibInstallTimeInfo) Reset()                    { *m = BgpRibInstallTimeInfo{} }
func (m *BgpRibInstallTimeInfo) String() string            { return proto.CompactTextString(m) }
func (*BgpRibInstallTimeInfo) ProtoMessage()               {}
func (*BgpRibInstallTimeInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *BgpRibInstallTimeInfo) GetUpdateTime() *BgpTimespec {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *BgpRibInstallTimeInfo) GetInstallTime() *BgpTimespec {
	if m != nil {
		return m.InstallTime
	}
	return nil
}

func (m *BgpRibInstallTimeInfo) GetInstalledCount() uint32 {
	if m != nil {
		return m.InstalledCount
	}
	return 0
}

func (m *BgpRibInstallTimeInfo) GetModifiedCount() uint32 {
	if m != nil {
		return m.ModifiedCount
	}
	return 0
}

func (m *BgpRibInstallTimeInfo) GetWithdrawnCount() uint32 {
	if m != nil {
		return m.WithdrawnCount
	}
	return 0
}

func (m *BgpRibInstallTimeInfo) GetStartVersion() uint32 {
	if m != nil {
		return m.StartVersion
	}
	return 0
}

func (m *BgpRibInstallTimeInfo) GetTargetVersion() uint32 {
	if m != nil {
		return m.TargetVersion
	}
	return 0
}

type BgpGlobalProcessInfoAfGbl_ struct {
	// Period (in seconds) of address-family scanner runs
	ScannerPeriod uint32 `protobuf:"varint,1,opt,name=scanner_period,json=scannerPeriod" json:"scanner_period,omitempty"`
	// NH Tunnel Version
	NhTunnelVersion uint32 `protobuf:"varint,2,opt,name=nh_tunnel_version,json=nhTunnelVersion" json:"nh_tunnel_version,omitempty"`
	// Sync-group Version
	SyncgrpVersion []uint32 `protobuf:"varint,3,rep,packed,name=syncgrp_version,json=syncgrpVersion" json:"syncgrp_version,omitempty"`
	// Total prefixes of address-family scanned
	ScanPrefixes uint32 `protobuf:"varint,4,opt,name=scan_prefixes,json=scanPrefixes" json:"scan_prefixes,omitempty"`
	// Number of prefixes scanned in a segment of address-family
	ScanSegmentPrefixes uint32 `protobuf:"varint,5,opt,name=scan_segment_prefixes,json=scanSegmentPrefixes" json:"scan_segment_prefixes,omitempty"`
	// Number of segments to scan all prefixes of address-family
	ScanSegments uint32 `protobuf:"varint,6,opt,name=scan_segments,json=scanSegments" json:"scan_segments,omitempty"`
	// Is inter-AS install to pim on ASBR enabled
	InterAsInstallEnabled bool `protobuf:"varint,7,opt,name=inter_as_install_enabled,json=interAsInstallEnabled" json:"inter_as_install_enabled,omitempty"`
	// Is global table mcast enabled
	GlobalMcastEnabled bool `protobuf:"varint,8,opt,name=global_mcast_enabled,json=globalMcastEnabled" json:"global_mcast_enabled,omitempty"`
	// Is segmented mcast enabled
	SegmentedMcastEnabled bool `protobuf:"varint,9,opt,name=segmented_mcast_enabled,json=segmentedMcastEnabled" json:"segmented_mcast_enabled,omitempty"`
	// Prefix validation disabled
	GblafrpkiDisable bool `protobuf:"varint,10,opt,name=gblafrpki_disable,json=gblafrpkiDisable" json:"gblafrpki_disable,omitempty"`
	// Prefix v. use validity
	GblafrpkiUseValidity bool `protobuf:"varint,11,opt,name=gblafrpki_use_validity,json=gblafrpkiUseValidity" json:"gblafrpki_use_validity,omitempty"`
	// Prefix v. allow invalid
	GblafrpkiAllowInvalid bool `protobuf:"varint,12,opt,name=gblafrpki_allow_invalid,json=gblafrpkiAllowInvalid" json:"gblafrpki_allow_invalid,omitempty"`
	// Prefix v. signal ibgp
	GblafrpkiSignalIbgp bool `protobuf:"varint,13,opt,name=gblafrpki_signal_ibgp,json=gblafrpkiSignalIbgp" json:"gblafrpki_signal_ibgp,omitempty"`
	// Update wait-install enabled
	UpdateWaitInstallEnabled bool `protobuf:"varint,14,opt,name=update_wait_install_enabled,json=updateWaitInstallEnabled" json:"update_wait_install_enabled,omitempty"`
	// Counter for RIB ack requests
	RibAckRequests uint32 `protobuf:"varint,15,opt,name=rib_ack_requests,json=ribAckRequests" json:"rib_ack_requests,omitempty"`
	// Counter for RIB ack received
	RibAcksReceived uint32 `protobuf:"varint,16,opt,name=rib_acks_received,json=ribAcksReceived" json:"rib_acks_received,omitempty"`
	// Counter for slow RIB acks
	RibSlowAcks uint32 `protobuf:"varint,17,opt,name=rib_slow_acks,json=ribSlowAcks" json:"rib_slow_acks,omitempty"`
	// RIB install info
	RibInstall *BgpRibInstallTimeInfo `protobuf:"bytes,18,opt,name=rib_install,json=ribInstall" json:"rib_install,omitempty"`
	// Permanent Network Enabled
	IsPermNetCfg bool `protobuf:"varint,19,opt,name=is_perm_net_cfg,json=isPermNetCfg" json:"is_perm_net_cfg,omitempty"`
	// Count of removed perm paths
	PermNetDelCount uint32 `protobuf:"varint,20,opt,name=perm_net_del_count,json=permNetDelCount" json:"perm_net_del_count,omitempty"`
	// Count of stale perm paths
	PermNetStaleDelCount uint32 `protobuf:"varint,21,opt,name=perm_net_stale_del_count,json=permNetStaleDelCount" json:"perm_net_stale_del_count,omitempty"`
	// Count of stale marked perm paths
	PermNetStaleMarkCount uint32 `protobuf:"varint,22,opt,name=perm_net_stale_mark_count,json=permNetStaleMarkCount" json:"perm_net_stale_mark_count,omitempty"`
	// Count of inserted perm paths
	PermNetInsertCount uint32 `protobuf:"varint,23,opt,name=perm_net_insert_count,json=permNetInsertCount" json:"perm_net_insert_count,omitempty"`
	// Count of existing perm paths
	PermNetExistingCount uint32 `protobuf:"varint,24,opt,name=perm_net_existing_count,json=permNetExistingCount" json:"perm_net_existing_count,omitempty"`
	// Count of perm nets given by RPL
	PermNetRplQueryCount uint32 `protobuf:"varint,25,opt,name=perm_net_rpl_query_count,json=permNetRplQueryCount" json:"perm_net_rpl_query_count,omitempty"`
	// Count of perm nets processed in RPL
	PermNetRplProcessCount uint32 `protobuf:"varint,26,opt,name=perm_net_rpl_process_count,json=permNetRplProcessCount" json:"perm_net_rpl_process_count,omitempty"`
	// Count of neighbors configured with perm net
	PermNbrCount              uint32 `protobuf:"varint,27,opt,name=perm_nbr_count,json=permNbrCount" json:"perm_nbr_count,omitempty"`
	RibPermPelemNotFoundCount uint32 `protobuf:"varint,28,opt,name=rib_perm_pelem_not_found_count,json=ribPermPelemNotFoundCount" json:"rib_perm_pelem_not_found_count,omitempty"`
	RibPermPathNotFoundCount  uint32 `protobuf:"varint,29,opt,name=rib_perm_path_not_found_count,json=ribPermPathNotFoundCount" json:"rib_perm_path_not_found_count,omitempty"`
	RibPermPelemFoundCount    uint32 `protobuf:"varint,30,opt,name=rib_perm_pelem_found_count,json=ribPermPelemFoundCount" json:"rib_perm_pelem_found_count,omitempty"`
	RibRegPathFoundCount      uint32 `protobuf:"varint,31,opt,name=rib_reg_path_found_count,json=ribRegPathFoundCount" json:"rib_reg_path_found_count,omitempty"`
	RibPermPathFoundCount     uint32 `protobuf:"varint,32,opt,name=rib_perm_path_found_count,json=ribPermPathFoundCount" json:"rib_perm_path_found_count,omitempty"`
	// Count of freed perm pelems
	PermPelemFreeCount uint32 `protobuf:"varint,33,opt,name=perm_pelem_free_count,json=permPelemFreeCount" json:"perm_pelem_free_count,omitempty"`
	// Count of perm paths refreshed
	PermPathRefreshCount uint32 `protobuf:"varint,34,opt,name=perm_path_refresh_count,json=permPathRefreshCount" json:"perm_path_refresh_count,omitempty"`
	// Count of bumped perm pelems
	PermPelemBumpCount uint32 `protobuf:"varint,35,opt,name=perm_pelem_bump_count,json=permPelemBumpCount" json:"perm_pelem_bump_count,omitempty"`
	// Count of bumped perm pelems
	PermPelemAllBumpCount uint32 `protobuf:"varint,36,opt,name=perm_pelem_all_bump_count,json=permPelemAllBumpCount" json:"perm_pelem_all_bump_count,omitempty"`
}

func (m *BgpGlobalProcessInfoAfGbl_) Reset()                    { *m = BgpGlobalProcessInfoAfGbl_{} }
func (m *BgpGlobalProcessInfoAfGbl_) String() string            { return proto.CompactTextString(m) }
func (*BgpGlobalProcessInfoAfGbl_) ProtoMessage()               {}
func (*BgpGlobalProcessInfoAfGbl_) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *BgpGlobalProcessInfoAfGbl_) GetScannerPeriod() uint32 {
	if m != nil {
		return m.ScannerPeriod
	}
	return 0
}

func (m *BgpGlobalProcessInfoAfGbl_) GetNhTunnelVersion() uint32 {
	if m != nil {
		return m.NhTunnelVersion
	}
	return 0
}

func (m *BgpGlobalProcessInfoAfGbl_) GetSyncgrpVersion() []uint32 {
	if m != nil {
		return m.SyncgrpVersion
	}
	return nil
}

func (m *BgpGlobalProcessInfoAfGbl_) GetScanPrefixes() uint32 {
	if m != nil {
		return m.ScanPrefixes
	}
	return 0
}

func (m *BgpGlobalProcessInfoAfGbl_) GetScanSegmentPrefixes() uint32 {
	if m != nil {
		return m.ScanSegmentPrefixes
	}
	return 0
}

func (m *BgpGlobalProcessInfoAfGbl_) GetScanSegments() uint32 {
	if m != nil {
		return m.ScanSegments
	}
	return 0
}

func (m *BgpGlobalProcessInfoAfGbl_) GetInterAsInstallEnabled() bool {
	if m != nil {
		return m.InterAsInstallEnabled
	}
	return false
}

func (m *BgpGlobalProcessInfoAfGbl_) GetGlobalMcastEnabled() bool {
	if m != nil {
		return m.GlobalMcastEnabled
	}
	return false
}

func (m *BgpGlobalProcessInfoAfGbl_) GetSegmentedMcastEnabled() bool {
	if m != nil {
		return m.SegmentedMcastEnabled
	}
	return false
}

func (m *BgpGlobalProcessInfoAfGbl_) GetGblafrpkiDisable() bool {
	if m != nil {
		return m.GblafrpkiDisable
	}
	return false
}

func (m *BgpGlobalProcessInfoAfGbl_) GetGblafrpkiUseValidity() bool {
	if m != nil {
		return m.GblafrpkiUseValidity
	}
	return false
}

func (m *BgpGlobalProcessInfoAfGbl_) GetGblafrpkiAllowInvalid() bool {
	if m != nil {
		return m.GblafrpkiAllowInvalid
	}
	return false
}

func (m *BgpGlobalProcessInfoAfGbl_) GetGblafrpkiSignalIbgp() bool {
	if m != nil {
		return m.GblafrpkiSignalIbgp
	}
	return false
}

func (m *BgpGlobalProcessInfoAfGbl_) GetUpdateWaitInstallEnabled() bool {
	if m != nil {
		return m.UpdateWaitInstallEnabled
	}
	return false
}

func (m *BgpGlobalProcessInfoAfGbl_) GetRibAckRequests() uint32 {
	if m != nil {
		return m.RibAckRequests
	}
	return 0
}

func (m *BgpGlobalProcessInfoAfGbl_) GetRibAcksReceived() uint32 {
	if m != nil {
		return m.RibAcksReceived
	}
	return 0
}

func (m *BgpGlobalProcessInfoAfGbl_) GetRibSlowAcks() uint32 {
	if m != nil {
		return m.RibSlowAcks
	}
	return 0
}

func (m *BgpGlobalProcessInfoAfGbl_) GetRibInstall() *BgpRibInstallTimeInfo {
	if m != nil {
		return m.RibInstall
	}
	return nil
}

func (m *BgpGlobalProcessInfoAfGbl_) GetIsPermNetCfg() bool {
	if m != nil {
		return m.IsPermNetCfg
	}
	return false
}

func (m *BgpGlobalProcessInfoAfGbl_) GetPermNetDelCount() uint32 {
	if m != nil {
		return m.PermNetDelCount
	}
	return 0
}

func (m *BgpGlobalProcessInfoAfGbl_) GetPermNetStaleDelCount() uint32 {
	if m != nil {
		return m.PermNetStaleDelCount
	}
	return 0
}

func (m *BgpGlobalProcessInfoAfGbl_) GetPermNetStaleMarkCount() uint32 {
	if m != nil {
		return m.PermNetStaleMarkCount
	}
	return 0
}

func (m *BgpGlobalProcessInfoAfGbl_) GetPermNetInsertCount() uint32 {
	if m != nil {
		return m.PermNetInsertCount
	}
	return 0
}

func (m *BgpGlobalProcessInfoAfGbl_) GetPermNetExistingCount() uint32 {
	if m != nil {
		return m.PermNetExistingCount
	}
	return 0
}

func (m *BgpGlobalProcessInfoAfGbl_) GetPermNetRplQueryCount() uint32 {
	if m != nil {
		return m.PermNetRplQueryCount
	}
	return 0
}

func (m *BgpGlobalProcessInfoAfGbl_) GetPermNetRplProcessCount() uint32 {
	if m != nil {
		return m.PermNetRplProcessCount
	}
	return 0
}

func (m *BgpGlobalProcessInfoAfGbl_) GetPermNbrCount() uint32 {
	if m != nil {
		return m.PermNbrCount
	}
	return 0
}

func (m *BgpGlobalProcessInfoAfGbl_) GetRibPermPelemNotFoundCount() uint32 {
	if m != nil {
		return m.RibPermPelemNotFoundCount
	}
	return 0
}

func (m *BgpGlobalProcessInfoAfGbl_) GetRibPermPathNotFoundCount() uint32 {
	if m != nil {
		return m.RibPermPathNotFoundCount
	}
	return 0
}

func (m *BgpGlobalProcessInfoAfGbl_) GetRibPermPelemFoundCount() uint32 {
	if m != nil {
		return m.RibPermPelemFoundCount
	}
	return 0
}

func (m *BgpGlobalProcessInfoAfGbl_) GetRibRegPathFoundCount() uint32 {
	if m != nil {
		return m.RibRegPathFoundCount
	}
	return 0
}

func (m *BgpGlobalProcessInfoAfGbl_) GetRibPermPathFoundCount() uint32 {
	if m != nil {
		return m.RibPermPathFoundCount
	}
	return 0
}

func (m *BgpGlobalProcessInfoAfGbl_) GetPermPelemFreeCount() uint32 {
	if m != nil {
		return m.PermPelemFreeCount
	}
	return 0
}

func (m *BgpGlobalProcessInfoAfGbl_) GetPermPathRefreshCount() uint32 {
	if m != nil {
		return m.PermPathRefreshCount
	}
	return 0
}

func (m *BgpGlobalProcessInfoAfGbl_) GetPermPelemBumpCount() uint32 {
	if m != nil {
		return m.PermPelemBumpCount
	}
	return 0
}

func (m *BgpGlobalProcessInfoAfGbl_) GetPermPelemAllBumpCount() uint32 {
	if m != nil {
		return m.PermPelemAllBumpCount
	}
	return 0
}

type BgpGlobalProcessInfoAfVrf_ struct {
	// Table state is active
	TableIsActive bool `protobuf:"varint,1,opt,name=table_is_active,json=tableIsActive" json:"table_is_active,omitempty"`
	// Table ID
	TableId uint32 `protobuf:"varint,2,opt,name=table_id,json=tableId" json:"table_id,omitempty"`
	// Table version
	TableVersion uint32 `protobuf:"varint,3,opt,name=table_version,json=tableVersion" json:"table_version,omitempty"`
	// RD version
	RdVersion uint32 `protobuf:"varint,4,opt,name=rd_version,json=rdVersion" json:"rd_version,omitempty"`
	// Table version synced to RIB
	RibVersion uint32 `protobuf:"varint,5,opt,name=rib_version,json=ribVersion" json:"rib_version,omitempty"`
	// NSR conv version
	NsrConvVersion uint32 `protobuf:"varint,6,opt,name=nsr_conv_version,json=nsrConvVersion" json:"nsr_conv_version,omitempty"`
	// NSR is converged
	NsrIsConv bool `protobuf:"varint,7,opt,name=nsr_is_conv,json=nsrIsConv" json:"nsr_is_conv,omitempty"`
	// Client reflection enabled
	ClientReflectionEnabled bool `protobuf:"varint,8,opt,name=client_reflection_enabled,json=clientReflectionEnabled" json:"client_reflection_enabled,omitempty"`
	// Dampening enabled
	DampeningEnabled bool `protobuf:"varint,9,opt,name=dampening_enabled,json=dampeningEnabled" json:"dampening_enabled,omitempty"`
	// Distance for eBGP routes (external)
	EbgpDistance uint32 `protobuf:"varint,10,opt,name=ebgp_distance,json=ebgpDistance" json:"ebgp_distance,omitempty"`
	// Distance for iBGP routes (internal)
	IbgpDistance uint32 `protobuf:"varint,11,opt,name=ibgp_distance,json=ibgpDistance" json:"ibgp_distance,omitempty"`
	// Dist for aggregate routes (local)
	AggregatesDistance uint32 `protobuf:"varint,12,opt,name=aggregates_distance,json=aggregatesDistance" json:"aggregates_distance,omitempty"`
	// Update generation enabled for MED change
	DynamicMedEnabled bool `protobuf:"varint,13,opt,name=dynamic_med_enabled,json=dynamicMedEnabled" json:"dynamic_med_enabled,omitempty"`
	// Delay in update generation after a MED change (in minutes)
	DynamicMedInterval uint32 `protobuf:"varint,14,opt,name=dynamic_med_interval,json=dynamicMedInterval" json:"dynamic_med_interval,omitempty"`
	// Dynamic MED timer running
	DynamicMedTimerRunning bool `protobuf:"varint,15,opt,name=dynamic_med_timer_running,json=dynamicMedTimerRunning" json:"dynamic_med_timer_running,omitempty"`
	// Dynamic MED timer value (in seconds) left on the timer
	DynamicMedTimerValue uint32 `protobuf:"varint,16,opt,name=dynamic_med_timer_value,json=dynamicMedTimerValue" json:"dynamic_med_timer_value,omitempty"`
	// Dynamic MED periodic timer running
	DynamicMedPeriodicTimerRunning bool `protobuf:"varint,17,opt,name=dynamic_med_periodic_timer_running,json=dynamicMedPeriodicTimerRunning" json:"dynamic_med_periodic_timer_running,omitempty"`
	// Dynamic MED periodic timer value (in seconds) left on the timer
	DynamicMedPeriodicTimerValue uint32 `protobuf:"varint,18,opt,name=dynamic_med_periodic_timer_value,json=dynamicMedPeriodicTimerValue" json:"dynamic_med_periodic_timer_value,omitempty"`
	// Received convergence notification from RIB
	RibHasConverged bool `protobuf:"varint,19,opt,name=rib_has_converged,json=ribHasConverged" json:"rib_has_converged,omitempty"`
	// Last convergence version received from RIB
	RibConvergenceVersion uint32 `protobuf:"varint,20,opt,name=rib_convergence_version,json=ribConvergenceVersion" json:"rib_convergence_version,omitempty"`
	// Indicates if RIB table is in prefix-limit state
	IsRibTableFull bool `protobuf:"varint,21,opt,name=is_rib_table_full,json=isRibTableFull" json:"is_rib_table_full,omitempty"`
	// Version when RIB table became full (from non-full)
	RibTableFullVersion uint32 `protobuf:"varint,22,opt,name=rib_table_full_version,json=ribTableFullVersion" json:"rib_table_full_version,omitempty"`
	// Is nexthop resoultion minimum prefix-length configured ?
	NexthopResolutionMinimumPrefixLengthConfigured bool `protobuf:"varint,23,opt,name=nexthop_resolution_minimum_prefix_length_configured,json=nexthopResolutionMinimumPrefixLengthConfigured" json:"nexthop_resolution_minimum_prefix_length_configured,omitempty"`
	// Nexthop resoultion minimum prefix-length
	NexthopResolutionMinimumPrefixLength uint32 `protobuf:"varint,24,opt,name=nexthop_resolution_minimum_prefix_length,json=nexthopResolutionMinimumPrefixLength" json:"nexthop_resolution_minimum_prefix_length,omitempty"`
	// Selective eBGP multipath isenabled
	SelectiveEbgpMultipathEnabled bool `protobuf:"varint,25,opt,name=selective_ebgp_multipath_enabled,json=selectiveEbgpMultipathEnabled" json:"selective_ebgp_multipath_enabled,omitempty"`
	// Selective iBGP multipath isenabled
	SelectiveIbgpMultipathEnabled bool `protobuf:"varint,26,opt,name=selective_ibgp_multipath_enabled,json=selectiveIbgpMultipathEnabled" json:"selective_ibgp_multipath_enabled,omitempty"`
	// Selective eiBGP multipath isenabled
	SelectiveEibgpMultipathEnabled bool `protobuf:"varint,27,opt,name=selective_eibgp_multipath_enabled,json=selectiveEibgpMultipathEnabled" json:"selective_eibgp_multipath_enabled,omitempty"`
	// Table version acked by RIB
	RibAckedTableVersion uint32 `protobuf:"varint,28,opt,name=rib_acked_table_version,json=ribAckedTableVersion" json:"rib_acked_table_version,omitempty"`
}

func (m *BgpGlobalProcessInfoAfVrf_) Reset()                    { *m = BgpGlobalProcessInfoAfVrf_{} }
func (m *BgpGlobalProcessInfoAfVrf_) String() string            { return proto.CompactTextString(m) }
func (*BgpGlobalProcessInfoAfVrf_) ProtoMessage()               {}
func (*BgpGlobalProcessInfoAfVrf_) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *BgpGlobalProcessInfoAfVrf_) GetTableIsActive() bool {
	if m != nil {
		return m.TableIsActive
	}
	return false
}

func (m *BgpGlobalProcessInfoAfVrf_) GetTableId() uint32 {
	if m != nil {
		return m.TableId
	}
	return 0
}

func (m *BgpGlobalProcessInfoAfVrf_) GetTableVersion() uint32 {
	if m != nil {
		return m.TableVersion
	}
	return 0
}

func (m *BgpGlobalProcessInfoAfVrf_) GetRdVersion() uint32 {
	if m != nil {
		return m.RdVersion
	}
	return 0
}

func (m *BgpGlobalProcessInfoAfVrf_) GetRibVersion() uint32 {
	if m != nil {
		return m.RibVersion
	}
	return 0
}

func (m *BgpGlobalProcessInfoAfVrf_) GetNsrConvVersion() uint32 {
	if m != nil {
		return m.NsrConvVersion
	}
	return 0
}

func (m *BgpGlobalProcessInfoAfVrf_) GetNsrIsConv() bool {
	if m != nil {
		return m.NsrIsConv
	}
	return false
}

func (m *BgpGlobalProcessInfoAfVrf_) GetClientReflectionEnabled() bool {
	if m != nil {
		return m.ClientReflectionEnabled
	}
	return false
}

func (m *BgpGlobalProcessInfoAfVrf_) GetDampeningEnabled() bool {
	if m != nil {
		return m.DampeningEnabled
	}
	return false
}

func (m *BgpGlobalProcessInfoAfVrf_) GetEbgpDistance() uint32 {
	if m != nil {
		return m.EbgpDistance
	}
	return 0
}

func (m *BgpGlobalProcessInfoAfVrf_) GetIbgpDistance() uint32 {
	if m != nil {
		return m.IbgpDistance
	}
	return 0
}

func (m *BgpGlobalProcessInfoAfVrf_) GetAggregatesDistance() uint32 {
	if m != nil {
		return m.AggregatesDistance
	}
	return 0
}

func (m *BgpGlobalProcessInfoAfVrf_) GetDynamicMedEnabled() bool {
	if m != nil {
		return m.DynamicMedEnabled
	}
	return false
}

func (m *BgpGlobalProcessInfoAfVrf_) GetDynamicMedInterval() uint32 {
	if m != nil {
		return m.DynamicMedInterval
	}
	return 0
}

func (m *BgpGlobalProcessInfoAfVrf_) GetDynamicMedTimerRunning() bool {
	if m != nil {
		return m.DynamicMedTimerRunning
	}
	return false
}

func (m *BgpGlobalProcessInfoAfVrf_) GetDynamicMedTimerValue() uint32 {
	if m != nil {
		return m.DynamicMedTimerValue
	}
	return 0
}

func (m *BgpGlobalProcessInfoAfVrf_) GetDynamicMedPeriodicTimerRunning() bool {
	if m != nil {
		return m.DynamicMedPeriodicTimerRunning
	}
	return false
}

func (m *BgpGlobalProcessInfoAfVrf_) GetDynamicMedPeriodicTimerValue() uint32 {
	if m != nil {
		return m.DynamicMedPeriodicTimerValue
	}
	return 0
}

func (m *BgpGlobalProcessInfoAfVrf_) GetRibHasConverged() bool {
	if m != nil {
		return m.RibHasConverged
	}
	return false
}

func (m *BgpGlobalProcessInfoAfVrf_) GetRibConvergenceVersion() uint32 {
	if m != nil {
		return m.RibConvergenceVersion
	}
	return 0
}

func (m *BgpGlobalProcessInfoAfVrf_) GetIsRibTableFull() bool {
	if m != nil {
		return m.IsRibTableFull
	}
	return false
}

func (m *BgpGlobalProcessInfoAfVrf_) GetRibTableFullVersion() uint32 {
	if m != nil {
		return m.RibTableFullVersion
	}
	return 0
}

func (m *BgpGlobalProcessInfoAfVrf_) GetNexthopResolutionMinimumPrefixLengthConfigured() bool {
	if m != nil {
		return m.NexthopResolutionMinimumPrefixLengthConfigured
	}
	return false
}

func (m *BgpGlobalProcessInfoAfVrf_) GetNexthopResolutionMinimumPrefixLength() uint32 {
	if m != nil {
		return m.NexthopResolutionMinimumPrefixLength
	}
	return 0
}

func (m *BgpGlobalProcessInfoAfVrf_) GetSelectiveEbgpMultipathEnabled() bool {
	if m != nil {
		return m.SelectiveEbgpMultipathEnabled
	}
	return false
}

func (m *BgpGlobalProcessInfoAfVrf_) GetSelectiveIbgpMultipathEnabled() bool {
	if m != nil {
		return m.SelectiveIbgpMultipathEnabled
	}
	return false
}

func (m *BgpGlobalProcessInfoAfVrf_) GetSelectiveEibgpMultipathEnabled() bool {
	if m != nil {
		return m.SelectiveEibgpMultipathEnabled
	}
	return false
}

func (m *BgpGlobalProcessInfoAfVrf_) GetRibAckedTableVersion() uint32 {
	if m != nil {
		return m.RibAckedTableVersion
	}
	return 0
}

func init() {
	proto.RegisterType((*BgpGlobalProcessInfoAfBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.afs.af.global_af_process_info.bgp_global_process_info_af_bag_KEYS")
	proto.RegisterType((*BgpGlobalProcessInfoAfBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.afs.af.global_af_process_info.bgp_global_process_info_af_bag")
	proto.RegisterType((*BgpTimespec)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.afs.af.global_af_process_info.bgp_timespec")
	proto.RegisterType((*BgpRibInstallTimeInfo)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.afs.af.global_af_process_info.bgp_rib_install_time_info")
	proto.RegisterType((*BgpGlobalProcessInfoAfGbl_)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.afs.af.global_af_process_info.bgp_global_process_info_af_gbl_")
	proto.RegisterType((*BgpGlobalProcessInfoAfVrf_)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.afs.af.global_af_process_info.bgp_global_process_info_af_vrf_")
}

func init() { proto.RegisterFile("bgp_global_process_info_af_bag.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1738 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x58, 0x5b, 0x6f, 0x1c, 0x49,
	0x15, 0xd6, 0x6c, 0x96, 0xd8, 0xa9, 0x99, 0xf1, 0xa5, 0x7c, 0xeb, 0xc9, 0xc5, 0xf1, 0x8e, 0xbd,
	0xc4, 0xb0, 0xd2, 0x00, 0xb9, 0xb1, 0x44, 0x42, 0xc8, 0x38, 0x09, 0xcc, 0x82, 0x23, 0xd3, 0x0e,
	0x41, 0x3c, 0x95, 0xaa, 0xbb, 0x4f, 0xb7, 0x4b, 0xee, 0xae, 0xee, 0xad, 0xaa, 0x1e, 0xc7, 0x7f,
	0x80, 0x27, 0x7e, 0x01, 0x5a, 0x89, 0x27, 0x24, 0x7e, 0xc8, 0xfe, 0x2c, 0x1e, 0x50, 0x5d, 0xba,
	0xba, 0x67, 0xe2, 0x35, 0xfb, 0x14, 0xf6, 0x2d, 0x3e, 0xdf, 0xf7, 0x9d, 0x73, 0xaa, 0xfa, 0x5c,
	0x6a, 0x82, 0x0e, 0xa2, 0xac, 0x22, 0x59, 0x5e, 0x46, 0x34, 0x27, 0x95, 0x28, 0x63, 0x90, 0x92,
	0x30, 0x9e, 0x96, 0x84, 0xa6, 0x24, 0xa2, 0xd9, 0xa4, 0x12, 0xa5, 0x2a, 0xf1, 0x45, 0xcc, 0x64,
	0x5c, 0x12, 0x56, 0x4a, 0xf2, 0x5e, 0x10, 0x56, 0xcd, 0x9e, 0x12, 0xad, 0x2b, 0x2b, 0x10, 0x93,
	0x28, 0xab, 0x26, 0x8c, 0x4b, 0x45, 0x79, 0x0c, 0xd2, 0xff, 0xcb, 0xff, 0x83, 0xd0, 0x58, 0xb1,
	0x19, 0x4c, 0x12, 0x48, 0x69, 0x9d, 0x2b, 0x32, 0x13, 0xe9, 0x84, 0xa6, 0x72, 0x42, 0xd3, 0x89,
	0x0b, 0x4b, 0xd3, 0xb9, 0xc8, 0xe3, 0x18, 0xed, 0xdf, 0x9c, 0x14, 0xf9, 0xc3, 0xab, 0xbf, 0x9e,
	0xe1, 0x7d, 0x34, 0xf4, 0x31, 0x38, 0x2d, 0x20, 0xe8, 0xed, 0xf5, 0x0e, 0xef, 0x84, 0x83, 0xc6,
	0xf8, 0x86, 0x16, 0x80, 0x77, 0xd0, 0x12, 0x4d, 0x2d, 0xfc, 0x89, 0x81, 0x6f, 0xd3, 0x54, 0x03,
	0xe3, 0x6f, 0x6f, 0xa1, 0xdd, 0x9b, 0xa3, 0xe0, 0x11, 0x5a, 0x9e, 0x09, 0x27, 0x7e, 0x6c, 0xc4,
	0x4b, 0x33, 0x91, 0x2e, 0xba, 0x7d, 0xd2, 0x75, 0x8b, 0x0f, 0xd0, 0x0a, 0x93, 0x84, 0x4b, 0x41,
	0x80, 0xd3, 0x28, 0x87, 0x24, 0x78, 0xba, 0xd7, 0x3b, 0x5c, 0x0e, 0x07, 0x4c, 0xbe, 0x91, 0xe2,
	0x95, 0xb5, 0xe1, 0x7f, 0xf5, 0xd0, 0x6d, 0x1b, 0x38, 0x78, 0xb6, 0xd7, 0x3b, 0xec, 0x3f, 0xfe,
	0x7b, 0x6f, 0xf2, 0x11, 0x6f, 0x78, 0x72, 0xc3, 0xc1, 0xb3, 0x28, 0x27, 0xa1, 0x4b, 0x0e, 0xff,
	0xb3, 0x87, 0x6e, 0xcd, 0x44, 0x1a, 0x3c, 0xff, 0x81, 0x25, 0xa9, 0x3f, 0x4b, 0xa8, 0x33, 0x1b,
	0x7f, 0x85, 0x06, 0x9a, 0xa7, 0x58, 0x01, 0xb2, 0x82, 0x18, 0x07, 0x68, 0x49, 0x42, 0x5c, 0xf2,
	0x44, 0x9a, 0x72, 0x18, 0x86, 0xcd, 0x9f, 0x78, 0x0f, 0xf5, 0x39, 0xe5, 0x65, 0x83, 0x7e, 0x62,
	0xd0, 0xae, 0x69, 0xfc, 0xcd, 0xa7, 0x68, 0xa4, 0x9d, 0x09, 0x16, 0x11, 0x93, 0x76, 0x9e, 0x1b,
	0xc7, 0x26, 0x2c, 0xfe, 0x47, 0x0f, 0xf5, 0xeb, 0x2a, 0xa1, 0x0a, 0x8c, 0xd1, 0xb8, 0xef, 0x3f,
	0xbe, 0xfa, 0xe8, 0x57, 0xd2, 0x1c, 0x35, 0x44, 0x36, 0x9b, 0xb7, 0xac, 0x00, 0xfc, 0x4d, 0x0f,
	0x0d, 0xba, 0x29, 0x9b, 0xe3, 0xfd, 0x5f, 0xb3, 0xeb, 0xbb, 0x74, 0x4c, 0x7a, 0x8f, 0xd0, 0xaa,
	0xfb, 0x13, 0x12, 0x12, 0x97, 0x35, 0x57, 0xc1, 0x2d, 0x73, 0xff, 0x2b, 0xde, 0x7c, 0xac, 0xad,
	0xf8, 0x73, 0xb4, 0x52, 0x94, 0x09, 0x4b, 0x99, 0xe7, 0x7d, 0x6a, 0x78, 0xc3, 0xc6, 0x6a, 0x69,
	0x8f, 0xd0, 0xea, 0x25, 0x53, 0xe7, 0x89, 0xa0, 0x97, 0xdc, 0xf1, 0x7e, 0x64, 0xfd, 0x79, 0xb3,
	0x25, 0xee, 0xa3, 0xa1, 0x54, 0x54, 0x28, 0x32, 0x03, 0x21, 0x59, 0xc9, 0x83, 0xdb, 0x86, 0x36,
	0x30, 0xc6, 0x77, 0xd6, 0xa6, 0x83, 0x2a, 0x2a, 0x32, 0x68, 0x59, 0x4b, 0x36, 0xa8, 0xb5, 0x3a,
	0xda, 0xf8, 0xdb, 0x55, 0xf4, 0xf0, 0x7f, 0x34, 0x8e, 0x76, 0x25, 0x63, 0xca, 0x39, 0x08, 0x52,
	0x81, 0x60, 0x65, 0xe2, 0xaa, 0x70, 0xe8, 0xac, 0xa7, 0xc6, 0x88, 0x7f, 0x8a, 0xd6, 0xf9, 0x39,
	0x51, 0x35, 0xe7, 0x90, 0xfb, 0xa0, 0xb6, 0x22, 0x57, 0xf9, 0xf9, 0x5b, 0x63, 0x6f, 0xb2, 0x7b,
	0x84, 0x56, 0xe5, 0x15, 0x8f, 0x33, 0x51, 0x79, 0xe6, 0xad, 0xbd, 0x5b, 0xfa, 0xac, 0xce, 0xdc,
	0x10, 0xf5, 0x59, 0x63, 0xca, 0x49, 0x25, 0x20, 0x65, 0xef, 0x41, 0xba, 0xab, 0x1b, 0x68, 0xe3,
	0xa9, 0xb3, 0xe1, 0xc7, 0x68, 0xcb, 0x90, 0x24, 0x64, 0x05, 0x70, 0xd5, 0x92, 0xed, 0xfd, 0x6d,
	0x68, 0xf0, 0xcc, 0x62, 0x5e, 0xd3, 0x38, 0x76, 0x1a, 0xe9, 0x2f, 0xb1, 0xe5, 0x4a, 0xfc, 0x4b,
	0x14, 0x30, 0xae, 0x40, 0x10, 0x2a, 0x7d, 0xf3, 0x34, 0x23, 0x70, 0xc9, 0x8c, 0xc0, 0x2d, 0x83,
	0x1f, 0xc9, 0xa9, 0x45, 0x9b, 0x59, 0xf8, 0x73, 0xb4, 0xe9, 0x6e, 0xb4, 0x88, 0xa9, 0x54, 0x5e,
	0xb4, 0x6c, 0x44, 0xd8, 0x62, 0x27, 0x1a, 0x6a, 0x14, 0xcf, 0xd1, 0x8e, 0x4b, 0x05, 0x92, 0x05,
	0xd1, 0x1d, 0x1b, 0xc9, 0xc3, 0x73, 0xba, 0x2f, 0xd0, 0x7a, 0x16, 0xe5, 0x34, 0x15, 0xd5, 0x05,
	0x23, 0x09, 0x93, 0xda, 0x1a, 0x20, 0xa3, 0x58, 0xf3, 0xc0, 0x4b, 0x6b, 0xc7, 0x4f, 0xd1, 0x76,
	0x4b, 0xae, 0x25, 0x90, 0x19, 0xcd, 0x59, 0xc2, 0xd4, 0x55, 0xd0, 0x37, 0x8a, 0x4d, 0x8f, 0xfe,
	0x59, 0xc2, 0x3b, 0x87, 0xe9, 0xd4, 0x5a, 0x15, 0xcd, 0xf3, 0xf2, 0x92, 0x30, 0x6e, 0x94, 0xc1,
	0xc0, 0xa6, 0xe6, 0xe1, 0x23, 0x8d, 0x4e, 0x2d, 0xa8, 0x3f, 0x4b, 0xab, 0x93, 0x2c, 0xe3, 0x34,
	0x27, 0x2c, 0xca, 0xaa, 0x60, 0x68, 0x54, 0x1b, 0x1e, 0x3c, 0x33, 0xd8, 0x34, 0xca, 0x2a, 0xfc,
	0x6b, 0x74, 0xcf, 0xcd, 0xa3, 0x4b, 0xca, 0xd4, 0x07, 0x97, 0xbe, 0x62, 0x94, 0x81, 0xa5, 0xfc,
	0x85, 0x32, 0xb5, 0x70, 0xef, 0x87, 0x68, 0x4d, 0x0f, 0x3a, 0x1a, 0x5f, 0x10, 0x01, 0x5f, 0xd7,
	0x20, 0x95, 0x0c, 0x56, 0x6d, 0x13, 0x09, 0x16, 0x1d, 0xc5, 0x17, 0xa1, 0xb3, 0xea, 0x6a, 0x75,
	0x4c, 0x49, 0x04, 0xc4, 0xc0, 0x66, 0x90, 0x04, 0x6b, 0xb6, 0x5a, 0x2d, 0x55, 0x86, 0xce, 0x8c,
	0xc7, 0x68, 0xa8, 0xb9, 0x52, 0x9f, 0x5c, 0x0b, 0x82, 0x75, 0x3b, 0x67, 0x05, 0x8b, 0xce, 0xf2,
	0xf2, 0x52, 0x73, 0xf1, 0xbf, 0x7b, 0xa8, 0xdf, 0x99, 0xb1, 0x01, 0x36, 0xb3, 0xea, 0x6f, 0x1f,
	0x7f, 0xbb, 0x5c, 0x3b, 0xe8, 0x43, 0x24, 0x58, 0xe4, 0xee, 0x0a, 0x7f, 0x8e, 0x56, 0x99, 0xd4,
	0xad, 0x5c, 0x10, 0x0e, 0x8a, 0xc4, 0x69, 0x16, 0x6c, 0x34, 0xfb, 0xfc, 0x14, 0x44, 0xf1, 0x06,
	0xd4, 0x71, 0x9a, 0xe1, 0x2f, 0x10, 0xf6, 0x9c, 0x04, 0x72, 0x37, 0x92, 0x36, 0xed, 0x15, 0x55,
	0x96, 0xf7, 0x12, 0x72, 0x3b, 0x93, 0x9e, 0xa3, 0xc0, 0x93, 0x75, 0x14, 0xe8, 0x48, 0xb6, 0x8c,
	0x64, 0xd3, 0x49, 0xce, 0x34, 0xea, 0x75, 0x5f, 0xa2, 0xd1, 0x82, 0xae, 0xa0, 0xe2, 0xc2, 0x09,
	0xb7, 0x8d, 0x70, 0xab, 0x2b, 0x3c, 0xa1, 0xe2, 0xc2, 0x2a, 0x7f, 0x81, 0xb6, 0xbc, 0x92, 0x71,
	0x09, 0x42, 0x39, 0xd5, 0x8e, 0x51, 0x61, 0xa7, 0x9a, 0x1a, 0xc8, 0x4a, 0x9e, 0xa1, 0x1d, 0x2f,
	0x81, 0xf7, 0x4c, 0x2a, 0xc6, 0x33, 0x27, 0x0a, 0xe6, 0x72, 0x7c, 0xe5, 0xc0, 0x0f, 0xcf, 0x26,
	0xaa, 0x9c, 0x7c, 0x5d, 0x83, 0xb8, 0x72, 0xba, 0xd1, 0x9c, 0x2e, 0xac, 0xf2, 0x3f, 0x69, 0xd0,
	0xea, 0x5e, 0xa0, 0xbb, 0x73, 0xba, 0xe6, 0x53, 0x59, 0xe5, 0x5d, 0xa3, 0xdc, 0x6e, 0x95, 0xa7,
	0x16, 0xb6, 0xda, 0x03, 0xb4, 0x62, 0xb5, 0x91, 0x70, 0xfc, 0x7b, 0x76, 0x3e, 0x19, 0x7e, 0x24,
	0x2c, 0xeb, 0x08, 0xed, 0xea, 0xcf, 0x6d, 0x98, 0x15, 0xe4, 0x50, 0x10, 0x5e, 0x2a, 0x92, 0x96,
	0x35, 0x6f, 0x36, 0xcd, 0x7d, 0xa3, 0x1a, 0x09, 0x16, 0xe9, 0x2f, 0x7b, 0xaa, 0x39, 0x6f, 0x4a,
	0xf5, 0x5a, 0x33, 0xac, 0x8b, 0xdf, 0xa0, 0x07, 0xad, 0x0b, 0xaa, 0xce, 0x3f, 0xf0, 0xf0, 0xc0,
	0x78, 0x08, 0x1a, 0x0f, 0x54, 0x9d, 0xcf, 0x3b, 0x78, 0x81, 0xee, 0x2e, 0xe4, 0xd0, 0x55, 0xef,
	0xda, 0x53, 0x76, 0xe3, 0x77, 0xb4, 0xcf, 0x91, 0xf6, 0x4b, 0x04, 0x64, 0x36, 0x76, 0x57, 0xf9,
	0xd0, 0xde, 0xac, 0x60, 0x51, 0x08, 0x99, 0x0e, 0xdb, 0xd1, 0x7d, 0x89, 0x46, 0xf3, 0x49, 0x77,
	0x85, 0x7b, 0xb6, 0x6a, 0x3a, 0x09, 0x77, 0x94, 0x4d, 0xd5, 0xb8, 0x4c, 0x05, 0x80, 0x53, 0x7d,
	0xd6, 0x56, 0x8d, 0xcd, 0x52, 0x00, 0xcc, 0x57, 0x8d, 0x09, 0x24, 0x20, 0x15, 0x20, 0xcf, 0x9d,
	0x68, 0xdc, 0x7e, 0x7d, 0x1d, 0x27, 0xb4, 0xe0, 0x75, 0x91, 0xa2, 0xba, 0xa8, 0x9c, 0x68, 0x7f,
	0x21, 0xd2, 0x6f, 0xeb, 0xa2, 0x9a, 0x6f, 0x06, 0x2b, 0xd1, 0x0d, 0xdc, 0x91, 0x1d, 0xb4, 0xcd,
	0x60, 0x64, 0x47, 0x79, 0xee, 0x95, 0xe3, 0xff, 0xf4, 0x6f, 0x5c, 0xe3, 0xfa, 0x69, 0x89, 0x7f,
	0x8c, 0x56, 0x95, 0x9e, 0x92, 0x84, 0x49, 0x37, 0x59, 0xcc, 0x1e, 0x5f, 0xd6, 0x4f, 0x82, 0x28,
	0x87, 0xa9, 0x3c, 0x32, 0x46, 0xfd, 0x0b, 0xc1, 0xf1, 0x12, 0xb7, 0xbe, 0x97, 0x2c, 0x21, 0xd1,
	0x4b, 0xd3, 0x42, 0xed, 0xd2, 0x36, 0x45, 0x69, 0x8c, 0xcd, 0xca, 0x7e, 0x80, 0x90, 0x48, 0x3c,
	0xc3, 0xee, 0xeb, 0x3b, 0x22, 0x69, 0xe0, 0x87, 0x76, 0x4e, 0x36, 0xb8, 0x5d, 0xd1, 0x7a, 0x3c,
	0x35, 0x84, 0x43, 0xb4, 0xa6, 0x7f, 0x6a, 0xc4, 0x25, 0x9f, 0x2d, 0xbc, 0x70, 0x56, 0xb8, 0x14,
	0xc7, 0x25, 0x9f, 0x35, 0xcc, 0x5d, 0xd4, 0xd7, 0x4c, 0x26, 0x0d, 0xd9, 0x6d, 0xe4, 0x3b, 0x5c,
	0x8a, 0xa9, 0xd4, 0x34, 0xfc, 0x02, 0x8d, 0xe2, 0x9c, 0xe9, 0x17, 0x81, 0x80, 0x34, 0x87, 0x58,
	0xb1, 0x92, 0x2f, 0xac, 0xe2, 0x1d, 0x4b, 0x08, 0x3d, 0xde, 0xd9, 0xab, 0x09, 0x2d, 0x2a, 0xe0,
	0x7a, 0x46, 0xcc, 0x6f, 0xe2, 0x35, 0x0f, 0x34, 0xe4, 0x7d, 0x34, 0x04, 0x7d, 0xfd, 0x09, 0xb3,
	0x93, 0xdb, 0x2c, 0xe0, 0x61, 0x38, 0xd0, 0xc6, 0x97, 0xce, 0x66, 0x7e, 0xda, 0xcd, 0x91, 0xfa,
	0x96, 0xc4, 0xba, 0xa4, 0x9f, 0xa1, 0x0d, 0x9a, 0x65, 0x02, 0x32, 0xaa, 0x40, 0xb6, 0xd4, 0x81,
	0xad, 0x99, 0x16, 0xf2, 0x82, 0x09, 0xda, 0x48, 0xae, 0x38, 0x2d, 0x58, 0x4c, 0x0a, 0x48, 0x7c,
	0xa6, 0x76, 0xc5, 0xae, 0x3b, 0xe8, 0x04, 0x92, 0xce, 0xcb, 0xa4, 0xcb, 0x37, 0xcf, 0x97, 0x19,
	0xcd, 0xcd, 0x66, 0x1d, 0x86, 0xb8, 0x15, 0x4c, 0x1d, 0x82, 0x7f, 0x85, 0x46, 0x5d, 0x85, 0xde,
	0x29, 0x82, 0x88, 0x9a, 0xeb, 0x0b, 0x30, 0xcb, 0x75, 0x39, 0xdc, 0x6e, 0x65, 0xfa, 0x69, 0x2c,
	0x42, 0x8b, 0xea, 0xd6, 0xf9, 0x50, 0x3a, 0xa3, 0x79, 0x0d, 0x6e, 0xd5, 0x6e, 0x2e, 0x08, 0xdf,
	0x69, 0x0c, 0x7f, 0x85, 0xc6, 0x5d, 0x99, 0x7d, 0x74, 0xb2, 0x78, 0x21, 0xf4, 0xba, 0x09, 0xbd,
	0xdb, 0x7a, 0x38, 0x75, 0xbc, 0xb9, 0x14, 0x5e, 0xa3, 0xbd, 0x1b, 0x7c, 0xd9, 0x5c, 0xb0, 0xc9,
	0xe5, 0xfe, 0x77, 0x78, 0xb2, 0x39, 0xb9, 0xf7, 0xc2, 0x39, 0xb5, 0xc5, 0x06, 0x22, 0x83, 0xc4,
	0xad, 0x4d, 0xfd, 0x5e, 0xf8, 0x3d, 0x35, 0x25, 0x67, 0xcc, 0xfa, 0xc1, 0xa4, 0xb9, 0x0d, 0x4f,
	0xaf, 0xf2, 0xa6, 0x90, 0x37, 0xfd, 0x70, 0x3a, 0x6e, 0xd1, 0xa6, 0x9e, 0x7f, 0x82, 0xd6, 0x99,
	0x34, 0x0b, 0xdc, 0x76, 0x59, 0x5a, 0xe7, 0xb9, 0xd9, 0x9e, 0xcb, 0xe1, 0x0a, 0x93, 0x21, 0x8b,
	0xde, 0x6a, 0xf3, 0xeb, 0x3a, 0xcf, 0xf1, 0x13, 0xb4, 0x3d, 0xcf, 0xf3, 0x11, 0xec, 0xd2, 0xdc,
	0x10, 0x1d, 0x76, 0xe3, 0xff, 0x02, 0x3d, 0xe1, 0xf0, 0x5e, 0x9d, 0x97, 0x15, 0x11, 0x20, 0xcb,
	0xbc, 0x36, 0x0d, 0x51, 0x30, 0xce, 0x8a, 0xba, 0x70, 0xaf, 0x66, 0x92, 0x03, 0xcf, 0x94, 0x1e,
	0x70, 0x3c, 0x65, 0x59, 0x2d, 0x20, 0x31, 0x0b, 0x75, 0x39, 0x9c, 0x38, 0x69, 0xe8, 0x95, 0x27,
	0x56, 0x68, 0x9f, 0xd4, 0x7f, 0x34, 0xb2, 0x63, 0xaf, 0xc2, 0xef, 0xd0, 0xe1, 0xf7, 0x0d, 0xe6,
	0xb6, 0xef, 0xc1, 0xf7, 0x89, 0x80, 0x7f, 0x87, 0xf6, 0x24, 0x98, 0x66, 0x9d, 0x01, 0x31, 0x5d,
	0x57, 0xd4, 0xb9, 0x62, 0x66, 0x3a, 0x37, 0xd5, 0x3f, 0x32, 0x19, 0x3f, 0xf0, 0xbc, 0x57, 0x51,
	0x56, 0x9d, 0x34, 0xac, 0xa6, 0x13, 0xe6, 0x1c, 0xb1, 0xeb, 0x1d, 0xdd, 0x5d, 0x70, 0x34, 0xbd,
	0xce, 0xd1, 0x14, 0x7d, 0xd6, 0xc9, 0xe8, 0x3b, 0x3c, 0xdd, 0xb3, 0xd5, 0xda, 0xa6, 0xc4, 0xae,
	0x73, 0xf5, 0xcc, 0x56, 0x0e, 0x8d, 0x2f, 0x74, 0xbb, 0xcc, 0x8d, 0xda, 0xfb, 0x7e, 0x1f, 0x1e,
	0x69, 0xf4, 0x6d, 0x67, 0xe4, 0x46, 0xb7, 0xcd, 0x7f, 0x68, 0x3d, 0xf9, 0x6f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x03, 0x32, 0xce, 0x12, 0xf8, 0x12, 0x00, 0x00,
}