// Code generated by protoc-gen-go. DO NOT EDIT.
// source: isis_sh_traffic_global.proto

/*
Package cisco_ios_xr_clns_isis_oper_isis_instances_instance_statistics_global is a generated protocol buffer package.

It is generated from these files:
	isis_sh_traffic_global.proto

It has these top-level messages:
	IsisShTrafficGlobal_KEYS
	IsisShTrafficGlobal
	IsisTopoIdType
	IsisShTimestampType
	IsisAreaTopoStatsType
	IsisShTrafficAreaTopo
	IsisAreaStatsType
	IsisShTrafficArea
	IsisTrafficGlobalType
*/
package cisco_ios_xr_clns_isis_oper_isis_instances_instance_statistics_global

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

// IS-IS process traffic data
type IsisShTrafficGlobal_KEYS struct {
	InstanceName string `protobuf:"bytes,1,opt,name=instance_name,json=instanceName" json:"instance_name,omitempty"`
}

func (m *IsisShTrafficGlobal_KEYS) Reset()                    { *m = IsisShTrafficGlobal_KEYS{} }
func (m *IsisShTrafficGlobal_KEYS) String() string            { return proto.CompactTextString(m) }
func (*IsisShTrafficGlobal_KEYS) ProtoMessage()               {}
func (*IsisShTrafficGlobal_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IsisShTrafficGlobal_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

type IsisShTrafficGlobal struct {
	// Statistics
	Statistics *IsisTrafficGlobalType `protobuf:"bytes,50,opt,name=statistics" json:"statistics,omitempty"`
	// Per-area data
	PerAreaData []*IsisShTrafficArea `protobuf:"bytes,51,rep,name=per_area_data,json=perAreaData" json:"per_area_data,omitempty"`
}

func (m *IsisShTrafficGlobal) Reset()                    { *m = IsisShTrafficGlobal{} }
func (m *IsisShTrafficGlobal) String() string            { return proto.CompactTextString(m) }
func (*IsisShTrafficGlobal) ProtoMessage()               {}
func (*IsisShTrafficGlobal) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *IsisShTrafficGlobal) GetStatistics() *IsisTrafficGlobalType {
	if m != nil {
		return m.Statistics
	}
	return nil
}

func (m *IsisShTrafficGlobal) GetPerAreaData() []*IsisShTrafficArea {
	if m != nil {
		return m.PerAreaData
	}
	return nil
}

// Identification of an IS-IS topology
type IsisTopoIdType struct {
	// AF name
	AfName string `protobuf:"bytes,1,opt,name=af_name,json=afName" json:"af_name,omitempty"`
	// Sub-AF name
	SafName string `protobuf:"bytes,2,opt,name=saf_name,json=safName" json:"saf_name,omitempty"`
	// VRF Name
	VrfName string `protobuf:"bytes,3,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
	// Topology Name
	TopologyName string `protobuf:"bytes,4,opt,name=topology_name,json=topologyName" json:"topology_name,omitempty"`
}

func (m *IsisTopoIdType) Reset()                    { *m = IsisTopoIdType{} }
func (m *IsisTopoIdType) String() string            { return proto.CompactTextString(m) }
func (*IsisTopoIdType) ProtoMessage()               {}
func (*IsisTopoIdType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *IsisTopoIdType) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *IsisTopoIdType) GetSafName() string {
	if m != nil {
		return m.SafName
	}
	return ""
}

func (m *IsisTopoIdType) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *IsisTopoIdType) GetTopologyName() string {
	if m != nil {
		return m.TopologyName
	}
	return ""
}

// Timestamp for an event
type IsisShTimestampType struct {
	// Timestamp value (seconds)
	Seconds uint32 `protobuf:"varint,1,opt,name=seconds" json:"seconds,omitempty"`
	// Timestamp value (nanoseconds)
	NanoSeconds uint32 `protobuf:"varint,2,opt,name=nano_seconds,json=nanoSeconds" json:"nano_seconds,omitempty"`
}

func (m *IsisShTimestampType) Reset()                    { *m = IsisShTimestampType{} }
func (m *IsisShTimestampType) String() string            { return proto.CompactTextString(m) }
func (*IsisShTimestampType) ProtoMessage()               {}
func (*IsisShTimestampType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *IsisShTimestampType) GetSeconds() uint32 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *IsisShTimestampType) GetNanoSeconds() uint32 {
	if m != nil {
		return m.NanoSeconds
	}
	return 0
}

// Per-area, per-topology traffic statistics
type IsisAreaTopoStatsType struct {
	// Number of full SPF calculations run
	SpfRunCount uint32 `protobuf:"varint,1,opt,name=spf_run_count,json=spfRunCount" json:"spf_run_count,omitempty"`
	// Number of ISPF calculations run
	IspfRunCount uint32 `protobuf:"varint,2,opt,name=ispf_run_count,json=ispfRunCount" json:"ispf_run_count,omitempty"`
	// Number of Next Hop Calculations run
	NhcRunCount uint32 `protobuf:"varint,3,opt,name=nhc_run_count,json=nhcRunCount" json:"nhc_run_count,omitempty"`
	// Number of PRCs run
	PrcRunCount uint32 `protobuf:"varint,4,opt,name=prc_run_count,json=prcRunCount" json:"prc_run_count,omitempty"`
	// Number of periodic SPF calculations run
	PeriodicRunCount uint32 `protobuf:"varint,5,opt,name=periodic_run_count,json=periodicRunCount" json:"periodic_run_count,omitempty"`
}

func (m *IsisAreaTopoStatsType) Reset()                    { *m = IsisAreaTopoStatsType{} }
func (m *IsisAreaTopoStatsType) String() string            { return proto.CompactTextString(m) }
func (*IsisAreaTopoStatsType) ProtoMessage()               {}
func (*IsisAreaTopoStatsType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *IsisAreaTopoStatsType) GetSpfRunCount() uint32 {
	if m != nil {
		return m.SpfRunCount
	}
	return 0
}

func (m *IsisAreaTopoStatsType) GetIspfRunCount() uint32 {
	if m != nil {
		return m.IspfRunCount
	}
	return 0
}

func (m *IsisAreaTopoStatsType) GetNhcRunCount() uint32 {
	if m != nil {
		return m.NhcRunCount
	}
	return 0
}

func (m *IsisAreaTopoStatsType) GetPrcRunCount() uint32 {
	if m != nil {
		return m.PrcRunCount
	}
	return 0
}

func (m *IsisAreaTopoStatsType) GetPeriodicRunCount() uint32 {
	if m != nil {
		return m.PeriodicRunCount
	}
	return 0
}

// Per-area, per-topology traffic data
type IsisShTrafficAreaTopo struct {
	// Topology ID
	Id *IsisTopoIdType `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Statistics
	Statistics *IsisAreaTopoStatsType `protobuf:"bytes,2,opt,name=statistics" json:"statistics,omitempty"`
}

func (m *IsisShTrafficAreaTopo) Reset()                    { *m = IsisShTrafficAreaTopo{} }
func (m *IsisShTrafficAreaTopo) String() string            { return proto.CompactTextString(m) }
func (*IsisShTrafficAreaTopo) ProtoMessage()               {}
func (*IsisShTrafficAreaTopo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *IsisShTrafficAreaTopo) GetId() *IsisTopoIdType {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *IsisShTrafficAreaTopo) GetStatistics() *IsisAreaTopoStatsType {
	if m != nil {
		return m.Statistics
	}
	return nil
}

// Per-area traffic statistics
type IsisAreaStatsType struct {
	// Number of times system LSP rebuilt
	SystemLspBuildCount uint32 `protobuf:"varint,1,opt,name=system_lsp_build_count,json=systemLspBuildCount" json:"system_lsp_build_count,omitempty"`
	// Number of times system LSP refreshed
	SystemLspRefreshCount uint32 `protobuf:"varint,2,opt,name=system_lsp_refresh_count,json=systemLspRefreshCount" json:"system_lsp_refresh_count,omitempty"`
}

func (m *IsisAreaStatsType) Reset()                    { *m = IsisAreaStatsType{} }
func (m *IsisAreaStatsType) String() string            { return proto.CompactTextString(m) }
func (*IsisAreaStatsType) ProtoMessage()               {}
func (*IsisAreaStatsType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *IsisAreaStatsType) GetSystemLspBuildCount() uint32 {
	if m != nil {
		return m.SystemLspBuildCount
	}
	return 0
}

func (m *IsisAreaStatsType) GetSystemLspRefreshCount() uint32 {
	if m != nil {
		return m.SystemLspRefreshCount
	}
	return 0
}

// Per-area traffic data
type IsisShTrafficArea struct {
	// Level this data applies to
	Level string `protobuf:"bytes,1,opt,name=level" json:"level,omitempty"`
	// Statistics
	Statistics *IsisAreaStatsType `protobuf:"bytes,2,opt,name=statistics" json:"statistics,omitempty"`
	// Per-topoogy statistics
	PerTopologyData []*IsisShTrafficAreaTopo `protobuf:"bytes,3,rep,name=per_topology_data,json=perTopologyData" json:"per_topology_data,omitempty"`
}

func (m *IsisShTrafficArea) Reset()                    { *m = IsisShTrafficArea{} }
func (m *IsisShTrafficArea) String() string            { return proto.CompactTextString(m) }
func (*IsisShTrafficArea) ProtoMessage()               {}
func (*IsisShTrafficArea) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *IsisShTrafficArea) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

func (m *IsisShTrafficArea) GetStatistics() *IsisAreaStatsType {
	if m != nil {
		return m.Statistics
	}
	return nil
}

func (m *IsisShTrafficArea) GetPerTopologyData() []*IsisShTrafficAreaTopo {
	if m != nil {
		return m.PerTopologyData
	}
	return nil
}

// IS-IS process traffic statistics
type IsisTrafficGlobalType struct {
	// Fast-PSNP cache lookups
	FastPsnpLookupCount uint32 `protobuf:"varint,1,opt,name=fast_psnp_lookup_count,json=fastPsnpLookupCount" json:"fast_psnp_lookup_count,omitempty"`
	// Fast-PSNP cache hits
	FastPsnpLookupHitCount uint32 `protobuf:"varint,2,opt,name=fast_psnp_lookup_hit_count,json=fastPsnpLookupHitCount" json:"fast_psnp_lookup_hit_count,omitempty"`
	// Fast-CSNP cache lookups
	FastCsnpLookupCount uint32 `protobuf:"varint,3,opt,name=fast_csnp_lookup_count,json=fastCsnpLookupCount" json:"fast_csnp_lookup_count,omitempty"`
	// Fast-CSNP cache hits
	FastCsnpLookupHitCount uint32 `protobuf:"varint,4,opt,name=fast_csnp_lookup_hit_count,json=fastCsnpLookupHitCount" json:"fast_csnp_lookup_hit_count,omitempty"`
	// Fast-CSNP cache updates
	FastCsnpCacheUpdateCount uint32 `protobuf:"varint,5,opt,name=fast_csnp_cache_update_count,json=fastCsnpCacheUpdateCount" json:"fast_csnp_cache_update_count,omitempty"`
	// LSPs received with holdtime 0
	ZeroHoldtimeLspCount uint32 `protobuf:"varint,6,opt,name=zero_holdtime_lsp_count,json=zeroHoldtimeLspCount" json:"zero_holdtime_lsp_count,omitempty"`
	// LSPs received with invalid checksum
	InvalidChecksumLspCount uint32 `protobuf:"varint,7,opt,name=invalid_checksum_lsp_count,json=invalidChecksumLspCount" json:"invalid_checksum_lsp_count,omitempty"`
	// IIHs dropped - Not Used
	IihDroppedCount uint32 `protobuf:"varint,8,opt,name=iih_dropped_count,json=iihDroppedCount" json:"iih_dropped_count,omitempty"`
	// LSPs dropped
	LspDroppedCount uint32 `protobuf:"varint,9,opt,name=lsp_dropped_count,json=lspDroppedCount" json:"lsp_dropped_count,omitempty"`
	// SNPs dropped
	SnpDroppedCount uint32 `protobuf:"varint,10,opt,name=snp_dropped_count,json=snpDroppedCount" json:"snp_dropped_count,omitempty"`
	// Maximum IIH queue length - Not Used
	MaximumIihQueueLength uint32 `protobuf:"varint,11,opt,name=maximum_iih_queue_length,json=maximumIihQueueLength" json:"maximum_iih_queue_length,omitempty"`
	// Maximum update PDU queue length
	MaximumPduQueueLength uint32 `protobuf:"varint,12,opt,name=maximum_pdu_queue_length,json=maximumPduQueueLength" json:"maximum_pdu_queue_length,omitempty"`
	// Average hello process time
	AvgHelloProcessTime *IsisShTimestampType `protobuf:"bytes,13,opt,name=avg_hello_process_time,json=avgHelloProcessTime" json:"avg_hello_process_time,omitempty"`
	// Average hello receive rate in packets per second
	AvgHelloRecvRate uint32 `protobuf:"varint,14,opt,name=avg_hello_recv_rate,json=avgHelloRecvRate" json:"avg_hello_recv_rate,omitempty"`
	// Average CSNP process time
	AvgCsnpProcessTime *IsisShTimestampType `protobuf:"bytes,15,opt,name=avg_csnp_process_time,json=avgCsnpProcessTime" json:"avg_csnp_process_time,omitempty"`
	// Average csnp receive rate in packets per second
	AvgCsnpRecvRate uint32 `protobuf:"varint,16,opt,name=avg_csnp_recv_rate,json=avgCsnpRecvRate" json:"avg_csnp_recv_rate,omitempty"`
	// Average PSNP process time
	AvgPsnpProcessTime *IsisShTimestampType `protobuf:"bytes,17,opt,name=avg_psnp_process_time,json=avgPsnpProcessTime" json:"avg_psnp_process_time,omitempty"`
	// Average psnp receive rate in packets per second
	AvgPsnpRecvRate uint32 `protobuf:"varint,18,opt,name=avg_psnp_recv_rate,json=avgPsnpRecvRate" json:"avg_psnp_recv_rate,omitempty"`
	// Average LSP process time
	AvgLspProcessTime *IsisShTimestampType `protobuf:"bytes,19,opt,name=avg_lsp_process_time,json=avgLspProcessTime" json:"avg_lsp_process_time,omitempty"`
	// Average LSP receive rate in packets per second
	AvgLspRecvRate uint32 `protobuf:"varint,20,opt,name=avg_lsp_recv_rate,json=avgLspRecvRate" json:"avg_lsp_recv_rate,omitempty"`
	// Average hello transmit time
	AvgHelloTransmitTime *IsisShTimestampType `protobuf:"bytes,21,opt,name=avg_hello_transmit_time,json=avgHelloTransmitTime" json:"avg_hello_transmit_time,omitempty"`
	// Average hello send rate in packets per second
	AvgHelloSendRate uint32 `protobuf:"varint,22,opt,name=avg_hello_send_rate,json=avgHelloSendRate" json:"avg_hello_send_rate,omitempty"`
	// Average CSNP transmit time
	AvgCsnpTransmitTime *IsisShTimestampType `protobuf:"bytes,23,opt,name=avg_csnp_transmit_time,json=avgCsnpTransmitTime" json:"avg_csnp_transmit_time,omitempty"`
	// Average csnp send rate in packets per second
	AvgCsnpSendRate uint32 `protobuf:"varint,24,opt,name=avg_csnp_send_rate,json=avgCsnpSendRate" json:"avg_csnp_send_rate,omitempty"`
	// Average PSNP transmit time
	AvgPsnpTransmitTime *IsisShTimestampType `protobuf:"bytes,25,opt,name=avg_psnp_transmit_time,json=avgPsnpTransmitTime" json:"avg_psnp_transmit_time,omitempty"`
	// Average psnp send rate in packets per second
	AvgPsnpSendRate uint32 `protobuf:"varint,26,opt,name=avg_psnp_send_rate,json=avgPsnpSendRate" json:"avg_psnp_send_rate,omitempty"`
	// Average LSP transmit time
	AvgLspTransmitTime *IsisShTimestampType `protobuf:"bytes,27,opt,name=avg_lsp_transmit_time,json=avgLspTransmitTime" json:"avg_lsp_transmit_time,omitempty"`
	// Average LSP send rate in packets per second
	AvgLspSendRate uint32 `protobuf:"varint,28,opt,name=avg_lsp_send_rate,json=avgLspSendRate" json:"avg_lsp_send_rate,omitempty"`
}

func (m *IsisTrafficGlobalType) Reset()                    { *m = IsisTrafficGlobalType{} }
func (m *IsisTrafficGlobalType) String() string            { return proto.CompactTextString(m) }
func (*IsisTrafficGlobalType) ProtoMessage()               {}
func (*IsisTrafficGlobalType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *IsisTrafficGlobalType) GetFastPsnpLookupCount() uint32 {
	if m != nil {
		return m.FastPsnpLookupCount
	}
	return 0
}

func (m *IsisTrafficGlobalType) GetFastPsnpLookupHitCount() uint32 {
	if m != nil {
		return m.FastPsnpLookupHitCount
	}
	return 0
}

func (m *IsisTrafficGlobalType) GetFastCsnpLookupCount() uint32 {
	if m != nil {
		return m.FastCsnpLookupCount
	}
	return 0
}

func (m *IsisTrafficGlobalType) GetFastCsnpLookupHitCount() uint32 {
	if m != nil {
		return m.FastCsnpLookupHitCount
	}
	return 0
}

func (m *IsisTrafficGlobalType) GetFastCsnpCacheUpdateCount() uint32 {
	if m != nil {
		return m.FastCsnpCacheUpdateCount
	}
	return 0
}

func (m *IsisTrafficGlobalType) GetZeroHoldtimeLspCount() uint32 {
	if m != nil {
		return m.ZeroHoldtimeLspCount
	}
	return 0
}

func (m *IsisTrafficGlobalType) GetInvalidChecksumLspCount() uint32 {
	if m != nil {
		return m.InvalidChecksumLspCount
	}
	return 0
}

func (m *IsisTrafficGlobalType) GetIihDroppedCount() uint32 {
	if m != nil {
		return m.IihDroppedCount
	}
	return 0
}

func (m *IsisTrafficGlobalType) GetLspDroppedCount() uint32 {
	if m != nil {
		return m.LspDroppedCount
	}
	return 0
}

func (m *IsisTrafficGlobalType) GetSnpDroppedCount() uint32 {
	if m != nil {
		return m.SnpDroppedCount
	}
	return 0
}

func (m *IsisTrafficGlobalType) GetMaximumIihQueueLength() uint32 {
	if m != nil {
		return m.MaximumIihQueueLength
	}
	return 0
}

func (m *IsisTrafficGlobalType) GetMaximumPduQueueLength() uint32 {
	if m != nil {
		return m.MaximumPduQueueLength
	}
	return 0
}

func (m *IsisTrafficGlobalType) GetAvgHelloProcessTime() *IsisShTimestampType {
	if m != nil {
		return m.AvgHelloProcessTime
	}
	return nil
}

func (m *IsisTrafficGlobalType) GetAvgHelloRecvRate() uint32 {
	if m != nil {
		return m.AvgHelloRecvRate
	}
	return 0
}

func (m *IsisTrafficGlobalType) GetAvgCsnpProcessTime() *IsisShTimestampType {
	if m != nil {
		return m.AvgCsnpProcessTime
	}
	return nil
}

func (m *IsisTrafficGlobalType) GetAvgCsnpRecvRate() uint32 {
	if m != nil {
		return m.AvgCsnpRecvRate
	}
	return 0
}

func (m *IsisTrafficGlobalType) GetAvgPsnpProcessTime() *IsisShTimestampType {
	if m != nil {
		return m.AvgPsnpProcessTime
	}
	return nil
}

func (m *IsisTrafficGlobalType) GetAvgPsnpRecvRate() uint32 {
	if m != nil {
		return m.AvgPsnpRecvRate
	}
	return 0
}

func (m *IsisTrafficGlobalType) GetAvgLspProcessTime() *IsisShTimestampType {
	if m != nil {
		return m.AvgLspProcessTime
	}
	return nil
}

func (m *IsisTrafficGlobalType) GetAvgLspRecvRate() uint32 {
	if m != nil {
		return m.AvgLspRecvRate
	}
	return 0
}

func (m *IsisTrafficGlobalType) GetAvgHelloTransmitTime() *IsisShTimestampType {
	if m != nil {
		return m.AvgHelloTransmitTime
	}
	return nil
}

func (m *IsisTrafficGlobalType) GetAvgHelloSendRate() uint32 {
	if m != nil {
		return m.AvgHelloSendRate
	}
	return 0
}

func (m *IsisTrafficGlobalType) GetAvgCsnpTransmitTime() *IsisShTimestampType {
	if m != nil {
		return m.AvgCsnpTransmitTime
	}
	return nil
}

func (m *IsisTrafficGlobalType) GetAvgCsnpSendRate() uint32 {
	if m != nil {
		return m.AvgCsnpSendRate
	}
	return 0
}

func (m *IsisTrafficGlobalType) GetAvgPsnpTransmitTime() *IsisShTimestampType {
	if m != nil {
		return m.AvgPsnpTransmitTime
	}
	return nil
}

func (m *IsisTrafficGlobalType) GetAvgPsnpSendRate() uint32 {
	if m != nil {
		return m.AvgPsnpSendRate
	}
	return 0
}

func (m *IsisTrafficGlobalType) GetAvgLspTransmitTime() *IsisShTimestampType {
	if m != nil {
		return m.AvgLspTransmitTime
	}
	return nil
}

func (m *IsisTrafficGlobalType) GetAvgLspSendRate() uint32 {
	if m != nil {
		return m.AvgLspSendRate
	}
	return 0
}

func init() {
	proto.RegisterType((*IsisShTrafficGlobal_KEYS)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.statistics_global.isis_sh_traffic_global_KEYS")
	proto.RegisterType((*IsisShTrafficGlobal)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.statistics_global.isis_sh_traffic_global")
	proto.RegisterType((*IsisTopoIdType)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.statistics_global.isis_topo_id_type")
	proto.RegisterType((*IsisShTimestampType)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.statistics_global.isis_sh_timestamp_type")
	proto.RegisterType((*IsisAreaTopoStatsType)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.statistics_global.isis_area_topo_stats_type")
	proto.RegisterType((*IsisShTrafficAreaTopo)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.statistics_global.isis_sh_traffic_area_topo")
	proto.RegisterType((*IsisAreaStatsType)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.statistics_global.isis_area_stats_type")
	proto.RegisterType((*IsisShTrafficArea)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.statistics_global.isis_sh_traffic_area")
	proto.RegisterType((*IsisTrafficGlobalType)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.statistics_global.isis_traffic_global_type")
}

func init() { proto.RegisterFile("isis_sh_traffic_global.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1084 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x57, 0x4f, 0x6f, 0xdb, 0xb6,
	0x1b, 0x86, 0x9c, 0x34, 0x69, 0xe9, 0x38, 0xa9, 0x15, 0x27, 0x51, 0xd2, 0x1c, 0xf2, 0x73, 0x7f,
	0x87, 0xec, 0x9f, 0x0f, 0x09, 0x86, 0x02, 0x1b, 0x30, 0x60, 0x75, 0x0b, 0x64, 0x98, 0x31, 0x78,
	0x4a, 0x0a, 0x6c, 0x28, 0x06, 0x96, 0x91, 0x68, 0x8b, 0xa8, 0x24, 0x72, 0x24, 0x25, 0x34, 0xdb,
	0xa1, 0x3b, 0x0c, 0xd8, 0x0e, 0xc3, 0x30, 0x0c, 0x3b, 0xed, 0x43, 0xec, 0xbc, 0xef, 0xb2, 0x2f,
	0x33, 0x90, 0x94, 0x2c, 0x32, 0x76, 0x6e, 0x9e, 0x6f, 0x31, 0xdf, 0xe7, 0xe5, 0xf3, 0x3c, 0x7c,
	0xf9, 0x50, 0x08, 0x38, 0x26, 0x82, 0x08, 0x28, 0x12, 0x28, 0x39, 0x9a, 0x4c, 0x48, 0x04, 0xa7,
	0x29, 0xbd, 0x46, 0xe9, 0x80, 0x71, 0x2a, 0xa9, 0xff, 0x3c, 0x22, 0x22, 0xa2, 0x90, 0x50, 0x01,
	0xdf, 0x70, 0x18, 0xa5, 0xb9, 0x80, 0x1a, 0x4f, 0x19, 0xe6, 0x03, 0xf5, 0xd7, 0x80, 0xe4, 0x42,
	0xa2, 0x3c, 0xc2, 0xcd, 0x5f, 0x03, 0x21, 0x91, 0x24, 0x42, 0x92, 0x48, 0x54, 0x9b, 0xf5, 0x9f,
	0x82, 0x47, 0x8b, 0x69, 0xe0, 0xe7, 0xcf, 0xbf, 0xbe, 0xf4, 0x1f, 0x83, 0x4e, 0xdd, 0x0c, 0x73,
	0x94, 0xe1, 0xc0, 0x3b, 0xf1, 0x4e, 0x1f, 0x84, 0x5b, 0xf5, 0xe2, 0x17, 0x28, 0xc3, 0xfd, 0x3f,
	0x5b, 0x60, 0x7f, 0xf1, 0x26, 0xfe, 0x5b, 0x00, 0x1a, 0xce, 0xe0, 0xec, 0xc4, 0x3b, 0x6d, 0x9f,
	0xc1, 0xc1, 0x52, 0xa4, 0x6b, 0xe4, 0x6d, 0xd1, 0xf2, 0x86, 0xe1, 0xd0, 0xa2, 0xf4, 0xdf, 0x82,
	0x0e, 0xc3, 0x1c, 0x22, 0x8e, 0x11, 0x8c, 0x91, 0x44, 0xc1, 0xf9, 0xc9, 0xda, 0x69, 0xfb, 0xec,
	0xe5, 0x32, 0x35, 0x58, 0xb6, 0x15, 0x4f, 0xd8, 0x66, 0x98, 0x7f, 0xca, 0x31, 0x7a, 0x86, 0x24,
	0xea, 0xff, 0xe4, 0x81, 0xae, 0x51, 0x4a, 0x19, 0x85, 0x24, 0xd6, 0x12, 0xfd, 0x03, 0xb0, 0x89,
	0x26, 0xf6, 0x89, 0x6e, 0xa0, 0x89, 0x3a, 0x4b, 0xff, 0x10, 0xdc, 0x17, 0x75, 0xa5, 0xa5, 0x2b,
	0x9b, 0xa2, 0x29, 0x95, 0xbc, 0x2a, 0xad, 0x99, 0x52, 0xc9, 0x4d, 0xe9, 0x31, 0xe8, 0xa8, 0xed,
	0x53, 0x3a, 0xbd, 0x31, 0xf5, 0x75, 0x33, 0xa6, 0x7a, 0x51, 0x8f, 0xe9, 0x85, 0x35, 0x25, 0x92,
	0x61, 0x21, 0x51, 0xc6, 0x8c, 0x9a, 0x00, 0x6c, 0x0a, 0x1c, 0xd1, 0x3c, 0x16, 0x5a, 0x4d, 0x27,
	0xac, 0x7f, 0xfa, 0xff, 0x03, 0x5b, 0x39, 0xca, 0x29, 0xac, 0xcb, 0x2d, 0x5d, 0x6e, 0xab, 0xb5,
	0x4b, 0xb3, 0xd4, 0xff, 0xc7, 0x03, 0x87, 0x7a, 0x5f, 0x7d, 0xc6, 0xda, 0xa5, 0x3a, 0x27, 0x61,
	0xb6, 0xee, 0x83, 0x8e, 0x60, 0x13, 0xc8, 0x8b, 0x1c, 0x46, 0xb4, 0xc8, 0x65, 0x45, 0xd0, 0x16,
	0x6c, 0x12, 0x16, 0xf9, 0x50, 0x2d, 0xf9, 0xff, 0x07, 0xdb, 0xc4, 0x05, 0x19, 0x9a, 0x2d, 0x62,
	0xa3, 0xfa, 0xa0, 0x93, 0x27, 0x91, 0x05, 0x5a, 0xab, 0xb4, 0x24, 0x91, 0x8d, 0x61, 0xdc, 0xc6,
	0xac, 0x1b, 0x0c, 0xe3, 0x0d, 0xe6, 0x7d, 0xe0, 0x33, 0xcc, 0x09, 0x8d, 0x89, 0x0d, 0xbc, 0xa7,
	0x81, 0x0f, 0xeb, 0x4a, 0x8d, 0xee, 0xff, 0xdc, 0xaa, 0xdc, 0xdd, 0x1a, 0xb2, 0x36, 0xea, 0x27,
	0xa0, 0x45, 0x62, 0x6d, 0xa9, 0x7d, 0xf6, 0xd5, 0x52, 0xaf, 0xb5, 0x75, 0x59, 0xc2, 0x16, 0x89,
	0xfd, 0x1f, 0x3c, 0x27, 0x49, 0x2d, 0x4d, 0xf9, 0x6a, 0x99, 0x94, 0x8b, 0xc6, 0x67, 0x47, 0xa9,
	0xff, 0xa3, 0x07, 0x7a, 0x0d, 0xd2, 0x9a, 0xf1, 0x39, 0xd8, 0x17, 0x37, 0x42, 0xe2, 0x0c, 0xa6,
	0x82, 0xc1, 0xeb, 0x82, 0xa4, 0xb1, 0x33, 0xec, 0x5d, 0x53, 0x1d, 0x09, 0xf6, 0x54, 0xd5, 0xcc,
	0x18, 0x9e, 0x80, 0xc0, 0x6a, 0xe2, 0x78, 0xc2, 0xb1, 0x48, 0x9c, 0xf1, 0xef, 0xcd, 0xda, 0x42,
	0x53, 0x35, 0x13, 0xf9, 0xbb, 0x55, 0xc9, 0xb8, 0x35, 0x11, 0xbf, 0x07, 0xee, 0xa5, 0xb8, 0xc4,
	0x69, 0x95, 0x28, 0xf3, 0xc3, 0xff, 0x7e, 0xc1, 0xb9, 0xbd, 0x5c, 0xfa, 0xb9, 0x2d, 0x3e, 0x32,
	0xff, 0x17, 0x0f, 0x74, 0xd5, 0xf3, 0x33, 0x0b, 0xa7, 0x7e, 0x82, 0xd6, 0xf4, 0x13, 0xf4, 0xea,
	0x3f, 0x7c, 0x82, 0x34, 0x61, 0xb8, 0xc3, 0x30, 0xbf, 0xaa, 0x98, 0xf5, 0x5b, 0xf4, 0x57, 0x17,
	0x04, 0x77, 0xbd, 0x9a, 0x6a, 0x8a, 0x13, 0x24, 0x24, 0x64, 0x22, 0x67, 0x30, 0xa5, 0xf4, 0x75,
	0xc1, 0xdc, 0x29, 0xaa, 0xea, 0x58, 0xe4, 0x6c, 0xa4, 0x6b, 0x66, 0x8a, 0x1f, 0x81, 0xa3, 0xb9,
	0xa6, 0x84, 0x48, 0x67, 0x8e, 0xfb, 0x6e, 0xe3, 0x05, 0x91, 0xa6, 0xb7, 0x26, 0x8c, 0xe6, 0x08,
	0xd7, 0x1a, 0xc2, 0xe1, 0x1d, 0x84, 0xd1, 0x42, 0xc2, 0xf5, 0x86, 0x70, 0x38, 0x4f, 0xf8, 0x09,
	0x38, 0x6e, 0x7a, 0x23, 0x14, 0x25, 0x18, 0x16, 0x2c, 0x46, 0x12, 0x3b, 0x6f, 0x40, 0x50, 0x77,
	0x0f, 0x15, 0xe2, 0x85, 0x06, 0x98, 0xfe, 0x0f, 0xc1, 0xc1, 0x77, 0x98, 0x53, 0x98, 0xd0, 0x34,
	0x56, 0x2f, 0xa8, 0xbe, 0xb9, 0xa6, 0x75, 0x43, 0xb7, 0xf6, 0x54, 0xf9, 0xa2, 0xaa, 0x8e, 0x44,
	0x25, 0xf9, 0x63, 0x70, 0x44, 0xf2, 0x12, 0xa5, 0x24, 0x86, 0x51, 0x82, 0xa3, 0xd7, 0xa2, 0xc8,
	0xac, 0xce, 0x4d, 0xdd, 0x79, 0x50, 0x21, 0x86, 0x15, 0x60, 0xd6, 0xfc, 0x2e, 0xe8, 0x12, 0x92,
	0xc0, 0x98, 0x53, 0xc6, 0x70, 0x1d, 0xab, 0xfb, 0xba, 0x67, 0x87, 0x90, 0xe4, 0x99, 0x59, 0x9f,
	0x61, 0xd5, 0xbe, 0x2e, 0xf6, 0x81, 0xc1, 0xa6, 0x82, 0xdd, 0xc6, 0xaa, 0x53, 0x70, 0xb1, 0xc0,
	0x60, 0x45, 0xee, 0x62, 0x9f, 0x80, 0x20, 0x43, 0x6f, 0x48, 0x56, 0x64, 0x50, 0x69, 0xf9, 0xb6,
	0xc0, 0x05, 0x86, 0x29, 0xce, 0xa7, 0x32, 0x09, 0xda, 0x26, 0xaa, 0x55, 0xfd, 0x33, 0x92, 0x7c,
	0xa9, 0xaa, 0x23, 0x5d, 0xb4, 0x1b, 0x59, 0x5c, 0xb8, 0x8d, 0x5b, 0x4e, 0xe3, 0x38, 0x2e, 0xec,
	0xc6, 0xdf, 0x3d, 0xb0, 0x8f, 0xca, 0x29, 0x4c, 0x70, 0x9a, 0x52, 0xc8, 0x38, 0x8d, 0xb0, 0x10,
	0xfa, 0xab, 0x15, 0x74, 0x74, 0x82, 0xbf, 0x59, 0x76, 0x78, 0x9c, 0x0f, 0x62, 0xb8, 0x8b, 0xca,
	0xe9, 0x85, 0xe2, 0x1e, 0x1b, 0xea, 0x2b, 0x92, 0x61, 0xff, 0x03, 0xb0, 0xdb, 0x68, 0xe2, 0x38,
	0x2a, 0x21, 0x47, 0x12, 0x07, 0xdb, 0xe6, 0xcb, 0x51, 0x77, 0x84, 0x38, 0x2a, 0x43, 0x24, 0xb1,
	0xff, 0x9b, 0x07, 0xf6, 0x14, 0x5e, 0xdf, 0x36, 0xc7, 0xc2, 0xce, 0x2a, 0x2c, 0xf8, 0xa8, 0x9c,
	0xaa, 0x5b, 0x6c, 0x3b, 0x78, 0x0f, 0xf8, 0x33, 0x45, 0x8d, 0x81, 0x87, 0x66, 0xea, 0x15, 0x7e,
	0x4e, 0x3f, 0x9b, 0xd3, 0xdf, 0x5d, 0x95, 0xfe, 0xf1, 0x62, 0xfd, 0xcc, 0xd5, 0xef, 0xcf, 0xf4,
	0x8f, 0x6d, 0xfd, 0xbf, 0x7a, 0xa0, 0xa7, 0xd0, 0x2a, 0x12, 0x8e, 0xfc, 0xdd, 0x55, 0xc8, 0xef,
	0xa2, 0x72, 0x3a, 0x12, 0x8e, 0xfa, 0x77, 0x40, 0xb7, 0xd6, 0xd3, 0x88, 0xef, 0x69, 0xf1, 0xdb,
	0x06, 0x3d, 0xd3, 0xfe, 0x87, 0x07, 0x0e, 0x9a, 0xbb, 0x26, 0x39, 0xca, 0x45, 0x46, 0xa4, 0x91,
	0xbf, 0xb7, 0x0a, 0xf9, 0xbd, 0xfa, 0x3a, 0x5f, 0x55, 0xdc, 0xf3, 0x09, 0x10, 0x38, 0x8f, 0x8d,
	0x87, 0x7d, 0x37, 0x01, 0x97, 0x38, 0x8f, 0xb5, 0x8b, 0x3a, 0xc5, 0xfa, 0xbe, 0xb9, 0x26, 0x0e,
	0x56, 0x95, 0x62, 0x75, 0xa5, 0x1d, 0x0f, 0x76, 0x06, 0x1a, 0x0b, 0x81, 0x93, 0x81, 0x39, 0x07,
	0x6c, 0xde, 0xc1, 0xe1, 0xaa, 0x1c, 0x8c, 0xef, 0x70, 0xc0, 0x5c, 0x07, 0x47, 0x4e, 0x0a, 0x66,
	0x0e, 0xea, 0x14, 0xab, 0x5b, 0xe7, 0x1a, 0x78, 0xb4, 0xaa, 0x14, 0x8f, 0x84, 0xab, 0xdf, 0xca,
	0x41, 0x23, 0xff, 0xd8, 0xce, 0x41, 0xad, 0xfe, 0x7a, 0x43, 0xff, 0xaf, 0x7b, 0xfe, 0x6f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x80, 0xca, 0x94, 0xfd, 0x0b, 0x0f, 0x00, 0x00,
}
