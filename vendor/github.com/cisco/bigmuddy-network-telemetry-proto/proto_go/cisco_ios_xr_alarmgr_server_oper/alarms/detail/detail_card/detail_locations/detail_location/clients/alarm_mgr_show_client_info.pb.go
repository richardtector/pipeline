// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alarm_mgr_show_client_info.proto

/*
Package cisco_ios_xr_alarmgr_server_oper_alarms_detail_detail_card_detail_locations_detail_location_clients is a generated protocol buffer package.

It is generated from these files:
	alarm_mgr_show_client_info.proto

It has these top-level messages:
	AlarmMgrShowClientInfo_KEYS
	AlarmMgrShowClientInfo
	AlarmMgrShowClientData
*/
package cisco_ios_xr_alarmgr_server_oper_alarms_detail_detail_card_detail_locations_detail_location_clients

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

// Alarm mgr show alarm info
type AlarmMgrShowClientInfo_KEYS struct {
	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
}

func (m *AlarmMgrShowClientInfo_KEYS) Reset()                    { *m = AlarmMgrShowClientInfo_KEYS{} }
func (m *AlarmMgrShowClientInfo_KEYS) String() string            { return proto.CompactTextString(m) }
func (*AlarmMgrShowClientInfo_KEYS) ProtoMessage()               {}
func (*AlarmMgrShowClientInfo_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AlarmMgrShowClientInfo_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type AlarmMgrShowClientInfo struct {
	// Client List
	ClientInfo []*AlarmMgrShowClientData `protobuf:"bytes,50,rep,name=client_info,json=clientInfo" json:"client_info,omitempty"`
}

func (m *AlarmMgrShowClientInfo) Reset()                    { *m = AlarmMgrShowClientInfo{} }
func (m *AlarmMgrShowClientInfo) String() string            { return proto.CompactTextString(m) }
func (*AlarmMgrShowClientInfo) ProtoMessage()               {}
func (*AlarmMgrShowClientInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AlarmMgrShowClientInfo) GetClientInfo() []*AlarmMgrShowClientData {
	if m != nil {
		return m.ClientInfo
	}
	return nil
}

// Alarm mgr show client data
type AlarmMgrShowClientData struct {
	// Alarm client
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Alarms agent id of the client
	Id uint32 `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	// The location of this client
	Location string `protobuf:"bytes,3,opt,name=location" json:"location,omitempty"`
	// The client handle through which interface
	Handle string `protobuf:"bytes,4,opt,name=handle" json:"handle,omitempty"`
	// The current state of the client
	State string `protobuf:"bytes,5,opt,name=state" json:"state,omitempty"`
	// The type of the client
	Type string `protobuf:"bytes,6,opt,name=type" json:"type,omitempty"`
	// The current subscription status of the client
	FilterDisp bool `protobuf:"varint,7,opt,name=filter_disp,json=filterDisp" json:"filter_disp,omitempty"`
	// Alarms agent subscriber id of the client
	SubscriberId uint32 `protobuf:"varint,8,opt,name=subscriber_id,json=subscriberId" json:"subscriber_id,omitempty"`
	// The filter used for alarm severity
	FilterSeverity string `protobuf:"bytes,9,opt,name=filter_severity,json=filterSeverity" json:"filter_severity,omitempty"`
	// The filter used for alarm bi-state state+
	FilterState string `protobuf:"bytes,10,opt,name=filter_state,json=filterState" json:"filter_state,omitempty"`
	// The filter used for alarm group
	FilterGroup string `protobuf:"bytes,11,opt,name=filter_group,json=filterGroup" json:"filter_group,omitempty"`
	// Number of times the agent connected to the alarm mgr
	ConnectCount uint32 `protobuf:"varint,12,opt,name=connect_count,json=connectCount" json:"connect_count,omitempty"`
	// Agent connect timestamp
	ConnectTimestamp string `protobuf:"bytes,13,opt,name=connect_timestamp,json=connectTimestamp" json:"connect_timestamp,omitempty"`
	// Number of times the agent queried for alarms
	GetCount uint32 `protobuf:"varint,14,opt,name=get_count,json=getCount" json:"get_count,omitempty"`
	// Number of times the agent subscribed for alarms
	SubscribeCount uint32 `protobuf:"varint,15,opt,name=subscribe_count,json=subscribeCount" json:"subscribe_count,omitempty"`
	// Number of times the agent reported alarms
	ReportCount uint32 `protobuf:"varint,16,opt,name=report_count,json=reportCount" json:"report_count,omitempty"`
}

func (m *AlarmMgrShowClientData) Reset()                    { *m = AlarmMgrShowClientData{} }
func (m *AlarmMgrShowClientData) String() string            { return proto.CompactTextString(m) }
func (*AlarmMgrShowClientData) ProtoMessage()               {}
func (*AlarmMgrShowClientData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AlarmMgrShowClientData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AlarmMgrShowClientData) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AlarmMgrShowClientData) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *AlarmMgrShowClientData) GetHandle() string {
	if m != nil {
		return m.Handle
	}
	return ""
}

func (m *AlarmMgrShowClientData) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *AlarmMgrShowClientData) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *AlarmMgrShowClientData) GetFilterDisp() bool {
	if m != nil {
		return m.FilterDisp
	}
	return false
}

func (m *AlarmMgrShowClientData) GetSubscriberId() uint32 {
	if m != nil {
		return m.SubscriberId
	}
	return 0
}

func (m *AlarmMgrShowClientData) GetFilterSeverity() string {
	if m != nil {
		return m.FilterSeverity
	}
	return ""
}

func (m *AlarmMgrShowClientData) GetFilterState() string {
	if m != nil {
		return m.FilterState
	}
	return ""
}

func (m *AlarmMgrShowClientData) GetFilterGroup() string {
	if m != nil {
		return m.FilterGroup
	}
	return ""
}

func (m *AlarmMgrShowClientData) GetConnectCount() uint32 {
	if m != nil {
		return m.ConnectCount
	}
	return 0
}

func (m *AlarmMgrShowClientData) GetConnectTimestamp() string {
	if m != nil {
		return m.ConnectTimestamp
	}
	return ""
}

func (m *AlarmMgrShowClientData) GetGetCount() uint32 {
	if m != nil {
		return m.GetCount
	}
	return 0
}

func (m *AlarmMgrShowClientData) GetSubscribeCount() uint32 {
	if m != nil {
		return m.SubscribeCount
	}
	return 0
}

func (m *AlarmMgrShowClientData) GetReportCount() uint32 {
	if m != nil {
		return m.ReportCount
	}
	return 0
}

func init() {
	proto.RegisterType((*AlarmMgrShowClientInfo_KEYS)(nil), "cisco_ios_xr_alarmgr_server_oper.alarms.detail.detail_card.detail_locations.detail_location.clients.alarm_mgr_show_client_info_KEYS")
	proto.RegisterType((*AlarmMgrShowClientInfo)(nil), "cisco_ios_xr_alarmgr_server_oper.alarms.detail.detail_card.detail_locations.detail_location.clients.alarm_mgr_show_client_info")
	proto.RegisterType((*AlarmMgrShowClientData)(nil), "cisco_ios_xr_alarmgr_server_oper.alarms.detail.detail_card.detail_locations.detail_location.clients.alarm_mgr_show_client_data")
}

func init() { proto.RegisterFile("alarm_mgr_show_client_info.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 451 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0x87, 0x95, 0x6e, 0xb7, 0xdb, 0x4e, 0xfa, 0x67, 0xb1, 0x10, 0x58, 0xcb, 0x61, 0xb3, 0xe5,
	0xb0, 0x95, 0x90, 0x72, 0x58, 0x6e, 0x5c, 0x01, 0xa1, 0x8a, 0x5b, 0xcb, 0x85, 0x93, 0xe5, 0xda,
	0xd3, 0xae, 0xa5, 0xc4, 0x8e, 0x6c, 0x77, 0x61, 0x4f, 0xbc, 0x01, 0xcf, 0xc0, 0x9b, 0xf0, 0x6a,
	0x28, 0xb6, 0xd3, 0xad, 0x10, 0xbd, 0x72, 0x8a, 0xe7, 0xcb, 0x97, 0x99, 0xdf, 0x24, 0x0a, 0x14,
	0xbc, 0xe2, 0xb6, 0x66, 0xf5, 0xce, 0x32, 0x77, 0x6f, 0xbe, 0x31, 0x51, 0x29, 0xd4, 0x9e, 0x29,
	0xbd, 0x35, 0x65, 0x63, 0x8d, 0x37, 0x44, 0x08, 0xe5, 0x84, 0x61, 0xca, 0x38, 0xf6, 0xdd, 0xb2,
	0xa0, 0xb7, 0x32, 0xda, 0x07, 0xb4, 0xcc, 0x34, 0x68, 0xcb, 0xc0, 0x5c, 0x29, 0xd1, 0x73, 0x55,
	0xa5, 0x0b, 0x13, 0xdc, 0xca, 0xee, 0x5c, 0x19, 0xc1, 0xbd, 0x32, 0xda, 0xfd, 0x0d, 0xca, 0x38,
	0xce, 0xcd, 0xdf, 0xc1, 0xf5, 0xe9, 0x20, 0xec, 0xf3, 0xc7, 0xaf, 0x6b, 0xf2, 0x12, 0x2e, 0xb4,
	0x91, 0xc8, 0x94, 0xa4, 0x59, 0x91, 0x2d, 0x46, 0xab, 0x41, 0x5b, 0x2e, 0xe5, 0xfc, 0x77, 0x06,
	0x57, 0xa7, 0x1f, 0x26, 0xbf, 0x32, 0xc8, 0x8f, 0x6a, 0x7a, 0x57, 0x9c, 0x2d, 0xf2, 0xbb, 0x1f,
	0xe5, 0x7f, 0x58, 0xab, 0xfc, 0x77, 0x2c, 0xc9, 0x3d, 0x5f, 0x41, 0x2c, 0x96, 0x7a, 0x6b, 0xe6,
	0x3f, 0xfb, 0xa7, 0x36, 0x68, 0x55, 0x42, 0xa0, 0xaf, 0x79, 0x8d, 0x69, 0xed, 0x70, 0x26, 0x53,
	0xe8, 0x29, 0x49, 0x7b, 0x45, 0xb6, 0x98, 0xac, 0x7a, 0x4a, 0x92, 0x2b, 0x18, 0x76, 0xd3, 0xe9,
	0x59, 0xf0, 0x0e, 0x35, 0x79, 0x01, 0x83, 0x7b, 0xae, 0x65, 0x85, 0xb4, 0x1f, 0x5f, 0x5c, 0xac,
	0xc8, 0x73, 0x38, 0x77, 0x9e, 0x7b, 0xa4, 0xe7, 0x01, 0xc7, 0xa2, 0x9d, 0xe6, 0x1f, 0x1b, 0xa4,
	0x83, 0x38, 0xad, 0x3d, 0x93, 0x6b, 0xc8, 0xb7, 0xaa, 0xf2, 0x68, 0x99, 0x54, 0xae, 0xa1, 0x17,
	0x45, 0xb6, 0x18, 0xae, 0x20, 0xa2, 0x0f, 0xca, 0x35, 0xe4, 0x35, 0x4c, 0xdc, 0x7e, 0xe3, 0x84,
	0x55, 0x1b, 0xb4, 0xed, 0x27, 0x1a, 0x86, 0x64, 0xe3, 0x27, 0xb8, 0x94, 0xe4, 0x16, 0x66, 0xa9,
	0x8b, 0xc3, 0x07, 0xb4, 0xca, 0x3f, 0xd2, 0x51, 0x18, 0x32, 0x8d, 0x78, 0x9d, 0x28, 0xb9, 0x81,
	0x71, 0x27, 0x86, 0x7c, 0x10, 0xac, 0x14, 0x61, 0x1d, 0x52, 0x3e, 0x29, 0x3b, 0x6b, 0xf6, 0x0d,
	0xcd, 0x8f, 0x95, 0x4f, 0x2d, 0x6a, 0x33, 0x09, 0xa3, 0x35, 0x0a, 0xcf, 0x84, 0xd9, 0x6b, 0x4f,
	0xc7, 0x31, 0x53, 0x82, 0xef, 0x5b, 0x46, 0xde, 0xc0, 0xb3, 0x4e, 0xf2, 0xaa, 0x46, 0xe7, 0x79,
	0xdd, 0xd0, 0x49, 0x68, 0x76, 0x99, 0x6e, 0x7c, 0xe9, 0x38, 0x79, 0x05, 0xa3, 0x1d, 0x76, 0xdd,
	0xa6, 0xa1, 0xdb, 0x70, 0x87, 0xa9, 0xd3, 0x2d, 0xcc, 0x0e, 0xdb, 0x26, 0x65, 0x16, 0x94, 0xe9,
	0x01, 0x47, 0xf1, 0x06, 0xc6, 0x16, 0x1b, 0x63, 0xbb, 0x46, 0x97, 0xc1, 0xca, 0x23, 0x0b, 0xca,
	0x66, 0x10, 0x7e, 0xbd, 0xb7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x63, 0x6e, 0x2c, 0x9e,
	0x03, 0x00, 0x00,
}
