// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ospfv3_edm_request.proto

/*
Package cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_request_list_process_table_request is a generated protocol buffer package.

It is generated from these files:
	ospfv3_edm_request.proto

It has these top-level messages:
	Ospfv3EdmRequest_KEYS
	Ospfv3EdmRequest
	Ospfv3EdmLsaSum
*/
package cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_request_list_process_table_request

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

// OSPFv3 request list information
type Ospfv3EdmRequest_KEYS struct {
	ProcessName     string `protobuf:"bytes,1,opt,name=process_name,json=processName" json:"process_name,omitempty"`
	InterfaceName   string `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
	NeighborAddress string `protobuf:"bytes,3,opt,name=neighbor_address,json=neighborAddress" json:"neighbor_address,omitempty"`
}

func (m *Ospfv3EdmRequest_KEYS) Reset()                    { *m = Ospfv3EdmRequest_KEYS{} }
func (m *Ospfv3EdmRequest_KEYS) String() string            { return proto.CompactTextString(m) }
func (*Ospfv3EdmRequest_KEYS) ProtoMessage()               {}
func (*Ospfv3EdmRequest_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Ospfv3EdmRequest_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *Ospfv3EdmRequest_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *Ospfv3EdmRequest_KEYS) GetNeighborAddress() string {
	if m != nil {
		return m.NeighborAddress
	}
	return ""
}

type Ospfv3EdmRequest struct {
	// Neighbor IP address
	RequestNeighborAddress string `protobuf:"bytes,50,opt,name=request_neighbor_address,json=requestNeighborAddress" json:"request_neighbor_address,omitempty"`
	// If true, virtual link is requested
	IsRequestVirtualLink bool `protobuf:"varint,51,opt,name=is_request_virtual_link,json=isRequestVirtualLink" json:"is_request_virtual_link,omitempty"`
	// Request virtual link ID
	RequestVirtualLinkId uint32 `protobuf:"varint,52,opt,name=request_virtual_link_id,json=requestVirtualLinkId" json:"request_virtual_link_id,omitempty"`
	// If true, sham link is requested
	IsRequestShamLink bool `protobuf:"varint,53,opt,name=is_request_sham_link,json=isRequestShamLink" json:"is_request_sham_link,omitempty"`
	// Request sham link ID
	RequestShamLinkId uint32 `protobuf:"varint,54,opt,name=request_sham_link_id,json=requestShamLinkId" json:"request_sham_link_id,omitempty"`
	// List of request list entries
	RequestList []*Ospfv3EdmLsaSum `protobuf:"bytes,55,rep,name=request_list,json=requestList" json:"request_list,omitempty"`
}

func (m *Ospfv3EdmRequest) Reset()                    { *m = Ospfv3EdmRequest{} }
func (m *Ospfv3EdmRequest) String() string            { return proto.CompactTextString(m) }
func (*Ospfv3EdmRequest) ProtoMessage()               {}
func (*Ospfv3EdmRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Ospfv3EdmRequest) GetRequestNeighborAddress() string {
	if m != nil {
		return m.RequestNeighborAddress
	}
	return ""
}

func (m *Ospfv3EdmRequest) GetIsRequestVirtualLink() bool {
	if m != nil {
		return m.IsRequestVirtualLink
	}
	return false
}

func (m *Ospfv3EdmRequest) GetRequestVirtualLinkId() uint32 {
	if m != nil {
		return m.RequestVirtualLinkId
	}
	return 0
}

func (m *Ospfv3EdmRequest) GetIsRequestShamLink() bool {
	if m != nil {
		return m.IsRequestShamLink
	}
	return false
}

func (m *Ospfv3EdmRequest) GetRequestShamLinkId() uint32 {
	if m != nil {
		return m.RequestShamLinkId
	}
	return 0
}

func (m *Ospfv3EdmRequest) GetRequestList() []*Ospfv3EdmLsaSum {
	if m != nil {
		return m.RequestList
	}
	return nil
}

// LSA summary entry
type Ospfv3EdmLsaSum struct {
	// LSA type
	HeaderLsaType string `protobuf:"bytes,1,opt,name=header_lsa_type,json=headerLsaType" json:"header_lsa_type,omitempty"`
	// Age of the LSA (seconds)
	HeaderLsaAge uint32 `protobuf:"varint,2,opt,name=header_lsa_age,json=headerLsaAge" json:"header_lsa_age,omitempty"`
	// LSA ID
	HeaderLsaId string `protobuf:"bytes,3,opt,name=header_lsa_id,json=headerLsaId" json:"header_lsa_id,omitempty"`
	// Router ID of the advertising router
	HeaderAdvertisingRouter string `protobuf:"bytes,4,opt,name=header_advertising_router,json=headerAdvertisingRouter" json:"header_advertising_router,omitempty"`
	// Current LSA sequence number
	HeaderSequenceNumber int32 `protobuf:"zigzag32,5,opt,name=header_sequence_number,json=headerSequenceNumber" json:"header_sequence_number,omitempty"`
}

func (m *Ospfv3EdmLsaSum) Reset()                    { *m = Ospfv3EdmLsaSum{} }
func (m *Ospfv3EdmLsaSum) String() string            { return proto.CompactTextString(m) }
func (*Ospfv3EdmLsaSum) ProtoMessage()               {}
func (*Ospfv3EdmLsaSum) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Ospfv3EdmLsaSum) GetHeaderLsaType() string {
	if m != nil {
		return m.HeaderLsaType
	}
	return ""
}

func (m *Ospfv3EdmLsaSum) GetHeaderLsaAge() uint32 {
	if m != nil {
		return m.HeaderLsaAge
	}
	return 0
}

func (m *Ospfv3EdmLsaSum) GetHeaderLsaId() string {
	if m != nil {
		return m.HeaderLsaId
	}
	return ""
}

func (m *Ospfv3EdmLsaSum) GetHeaderAdvertisingRouter() string {
	if m != nil {
		return m.HeaderAdvertisingRouter
	}
	return ""
}

func (m *Ospfv3EdmLsaSum) GetHeaderSequenceNumber() int32 {
	if m != nil {
		return m.HeaderSequenceNumber
	}
	return 0
}

func init() {
	proto.RegisterType((*Ospfv3EdmRequest_KEYS)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.request_list_process_table.request.ospfv3_edm_request_KEYS")
	proto.RegisterType((*Ospfv3EdmRequest)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.request_list_process_table.request.ospfv3_edm_request")
	proto.RegisterType((*Ospfv3EdmLsaSum)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.request_list_process_table.request.ospfv3_edm_lsa_sum")
}

func init() { proto.RegisterFile("ospfv3_edm_request.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 472 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0xc6, 0x15, 0x0a, 0x08, 0xdc, 0x66, 0x97, 0x5a, 0xd1, 0x36, 0xdc, 0x42, 0x05, 0x28, 0x5c,
	0xb2, 0xd2, 0x76, 0x77, 0x41, 0xdc, 0x7a, 0xe0, 0x50, 0x51, 0xf5, 0x90, 0x22, 0x24, 0x4e, 0x96,
	0x5b, 0x4f, 0x5b, 0xab, 0xf9, 0x87, 0xed, 0x44, 0xf4, 0x25, 0x78, 0x02, 0x5e, 0x85, 0xf7, 0xe2,
	0x88, 0x62, 0x3b, 0x21, 0x55, 0x38, 0xef, 0x2d, 0xfa, 0xbe, 0xdf, 0xcc, 0x37, 0xd6, 0x4c, 0x90,
	0x9f, 0xcb, 0x62, 0x57, 0xcd, 0x08, 0xb0, 0x94, 0x08, 0xf8, 0x5e, 0x82, 0x54, 0x51, 0x21, 0x72,
	0x95, 0x63, 0xd8, 0x72, 0xb9, 0xcd, 0x09, 0xcf, 0x25, 0xf9, 0x21, 0x08, 0x2f, 0xaa, 0x7b, 0x62,
	0xd9, 0xbc, 0x00, 0x11, 0x99, 0xef, 0x9a, 0xdd, 0x82, 0x94, 0x20, 0x9b, 0xaf, 0x88, 0xc1, 0x8e,
	0x96, 0x89, 0x22, 0x95, 0xd8, 0x45, 0xb6, 0x23, 0x49, 0xb8, 0x54, 0xc4, 0x02, 0x44, 0xd1, 0x4d,
	0x02, 0x8d, 0x35, 0xfd, 0xe9, 0xa0, 0x49, 0x7f, 0x06, 0xf2, 0xf9, 0xd3, 0xb7, 0x35, 0x7e, 0x85,
	0x46, 0x4d, 0x51, 0x46, 0x53, 0xf0, 0x9d, 0xc0, 0x09, 0x9f, 0xc7, 0x43, 0xab, 0xad, 0x68, 0x0a,
	0xf8, 0x0d, 0xba, 0xe0, 0x99, 0x02, 0xb1, 0xa3, 0x5b, 0x30, 0xd0, 0x23, 0x0d, 0xb9, 0xad, 0xaa,
	0xb1, 0x77, 0xe8, 0x45, 0x06, 0x7c, 0x7f, 0xd8, 0xe4, 0x82, 0x50, 0xc6, 0x04, 0x48, 0xe9, 0x0f,
	0x34, 0x78, 0xd9, 0xe8, 0x73, 0x23, 0x4f, 0x7f, 0x0f, 0x10, 0xee, 0x0f, 0x84, 0x3f, 0x20, 0xbf,
	0x99, 0xad, 0xd7, 0xe9, 0x46, 0x77, 0xba, 0xb2, 0xfe, 0xea, 0xbc, 0x21, 0xbe, 0x43, 0x13, 0x2e,
	0xdb, 0x87, 0x55, 0x5c, 0xa8, 0x92, 0x26, 0x24, 0xe1, 0xd9, 0xd1, 0x9f, 0x05, 0x4e, 0xf8, 0x2c,
	0xf6, 0xb8, 0x8c, 0x8d, 0xfb, 0xd5, 0x98, 0x4b, 0x9e, 0x1d, 0xeb, 0xb2, 0xff, 0xd5, 0x10, 0xce,
	0xfc, 0xdb, 0xc0, 0x09, 0xdd, 0xd8, 0x13, 0xbd, 0xa2, 0x05, 0xc3, 0xd7, 0xc8, 0xeb, 0xa4, 0xc9,
	0x03, 0x4d, 0x4d, 0xd4, 0x9d, 0x8e, 0x1a, 0xb7, 0x51, 0xeb, 0x03, 0x4d, 0x75, 0xce, 0x35, 0xf2,
	0x7a, 0x74, 0x1d, 0x72, 0xaf, 0x43, 0xc6, 0xe2, 0x1c, 0x5f, 0x30, 0xfc, 0xcb, 0x41, 0xa3, 0xee,
	0x62, 0xfd, 0xf7, 0xc1, 0x20, 0x1c, 0xde, 0x9c, 0xa2, 0x07, 0x39, 0x98, 0xa8, 0xb3, 0x9b, 0x44,
	0x52, 0x22, 0xcb, 0x34, 0x1e, 0x5a, 0x6f, 0xc9, 0xa5, 0x9a, 0xfe, 0x71, 0xce, 0xf6, 0x67, 0x19,
	0xfc, 0x16, 0x5d, 0x1e, 0x80, 0x32, 0x10, 0x5a, 0x51, 0xa7, 0xa2, 0x39, 0x27, 0xd7, 0xc8, 0x4b,
	0x49, 0xbf, 0x9c, 0x0a, 0xc0, 0xaf, 0xd1, 0x45, 0x87, 0xa3, 0x7b, 0x73, 0x50, 0x6e, 0x3c, 0x6a,
	0xb1, 0xf9, 0x1e, 0xf0, 0x14, 0xb9, 0x1d, 0x8a, 0x33, 0x7b, 0x4c, 0xc3, 0x16, 0x5a, 0x30, 0xfc,
	0x11, 0xbd, 0xb4, 0x0c, 0x65, 0x15, 0x08, 0xc5, 0x25, 0xcf, 0xf6, 0x44, 0xe4, 0xa5, 0x02, 0xe1,
	0x3f, 0xd6, 0xfc, 0xc4, 0x00, 0xf3, 0x7f, 0x7e, 0xac, 0x6d, 0x7c, 0x8b, 0xae, 0x6c, 0xad, 0xac,
	0x9f, 0x96, 0xd5, 0xc7, 0x5d, 0xa6, 0x1b, 0x10, 0xfe, 0x93, 0xc0, 0x09, 0xc7, 0xb1, 0x67, 0xdc,
	0xb5, 0x35, 0x57, 0xda, 0xdb, 0x3c, 0xd5, 0x7f, 0xee, 0xec, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x62, 0xe0, 0x91, 0x94, 0xd5, 0x03, 0x00, 0x00,
}
