// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ipv6_if_brief.proto

/*
Package cisco_ios_xr_ipv6_ma_oper_ipv6_network_nodes_node_interface_data_vrfs_vrf_briefs_brief is a generated protocol buffer package.

It is generated from these files:
	ipv6_if_brief.proto

It has these top-level messages:
	Ipv6IfBrief_KEYS
	Ipv6IfBrief
	Ipv6AddrNode
*/
package cisco_ios_xr_ipv6_ma_oper_ipv6_network_nodes_node_interface_data_vrfs_vrf_briefs_brief

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

// Brief Summary of IPv6 Interface
type Ipv6IfBrief_KEYS struct {
	NodeName      string `protobuf:"bytes,1,opt,name=node_name,json=nodeName" json:"node_name,omitempty"`
	VrfName       string `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
	InterfaceName string `protobuf:"bytes,3,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
}

func (m *Ipv6IfBrief_KEYS) Reset()                    { *m = Ipv6IfBrief_KEYS{} }
func (m *Ipv6IfBrief_KEYS) String() string            { return proto.CompactTextString(m) }
func (*Ipv6IfBrief_KEYS) ProtoMessage()               {}
func (*Ipv6IfBrief_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Ipv6IfBrief_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *Ipv6IfBrief_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ipv6IfBrief_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type Ipv6IfBrief struct {
	// State of Interface Line
	LineState string `protobuf:"bytes,50,opt,name=line_state,json=lineState" json:"line_state,omitempty"`
	// Address List
	AddressList []*Ipv6AddrNode `protobuf:"bytes,51,rep,name=address_list,json=addressList" json:"address_list,omitempty"`
	// Link Local Address
	LinkLocalAddress *Ipv6AddrNode `protobuf:"bytes,52,opt,name=link_local_address,json=linkLocalAddress" json:"link_local_address,omitempty"`
}

func (m *Ipv6IfBrief) Reset()                    { *m = Ipv6IfBrief{} }
func (m *Ipv6IfBrief) String() string            { return proto.CompactTextString(m) }
func (*Ipv6IfBrief) ProtoMessage()               {}
func (*Ipv6IfBrief) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Ipv6IfBrief) GetLineState() string {
	if m != nil {
		return m.LineState
	}
	return ""
}

func (m *Ipv6IfBrief) GetAddressList() []*Ipv6AddrNode {
	if m != nil {
		return m.AddressList
	}
	return nil
}

func (m *Ipv6IfBrief) GetLinkLocalAddress() *Ipv6AddrNode {
	if m != nil {
		return m.LinkLocalAddress
	}
	return nil
}

// List of IPv6 Addresses
type Ipv6AddrNode struct {
	// IPv6 Address
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	// Prefix Length of IPv6 Address
	PrefixLength uint32 `protobuf:"varint,2,opt,name=prefix_length,json=prefixLength" json:"prefix_length,omitempty"`
	// State of Address
	AddressState string `protobuf:"bytes,3,opt,name=address_state,json=addressState" json:"address_state,omitempty"`
	// Anycast address
	IsAnycast bool `protobuf:"varint,4,opt,name=is_anycast,json=isAnycast" json:"is_anycast,omitempty"`
	// Route-tag of the Address
	RouteTag uint32 `protobuf:"varint,5,opt,name=route_tag,json=routeTag" json:"route_tag,omitempty"`
}

func (m *Ipv6AddrNode) Reset()                    { *m = Ipv6AddrNode{} }
func (m *Ipv6AddrNode) String() string            { return proto.CompactTextString(m) }
func (*Ipv6AddrNode) ProtoMessage()               {}
func (*Ipv6AddrNode) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Ipv6AddrNode) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Ipv6AddrNode) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

func (m *Ipv6AddrNode) GetAddressState() string {
	if m != nil {
		return m.AddressState
	}
	return ""
}

func (m *Ipv6AddrNode) GetIsAnycast() bool {
	if m != nil {
		return m.IsAnycast
	}
	return false
}

func (m *Ipv6AddrNode) GetRouteTag() uint32 {
	if m != nil {
		return m.RouteTag
	}
	return 0
}

func init() {
	proto.RegisterType((*Ipv6IfBrief_KEYS)(nil), "cisco_ios_xr_ipv6_ma_oper.ipv6_network.nodes.node.interface_data.vrfs.vrf.briefs.brief.ipv6_if_brief_KEYS")
	proto.RegisterType((*Ipv6IfBrief)(nil), "cisco_ios_xr_ipv6_ma_oper.ipv6_network.nodes.node.interface_data.vrfs.vrf.briefs.brief.ipv6_if_brief")
	proto.RegisterType((*Ipv6AddrNode)(nil), "cisco_ios_xr_ipv6_ma_oper.ipv6_network.nodes.node.interface_data.vrfs.vrf.briefs.brief.ipv6_addr_node")
}

func init() { proto.RegisterFile("ipv6_if_brief.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0xc1, 0x4a, 0xc3, 0x30,
	0x18, 0xc7, 0xe9, 0xa6, 0xae, 0xcd, 0xd6, 0x21, 0xf1, 0x52, 0x91, 0x41, 0x99, 0x08, 0x3b, 0xf5,
	0xb0, 0x89, 0xf7, 0x1d, 0x3c, 0x39, 0x3c, 0x74, 0x22, 0x78, 0xfa, 0xc8, 0xda, 0x64, 0x86, 0x75,
	0x49, 0x49, 0x62, 0x9d, 0x8f, 0xe0, 0xdd, 0x07, 0xf1, 0x55, 0x7c, 0x23, 0x49, 0xb2, 0x2a, 0x7b,
	0x00, 0xbd, 0x24, 0xe4, 0xf7, 0xfd, 0xbf, 0xfc, 0x3f, 0xf2, 0x0f, 0x3a, 0xe3, 0x75, 0x73, 0x03,
	0x9c, 0xc1, 0x4a, 0x71, 0xca, 0xb2, 0x5a, 0x49, 0x23, 0xf1, 0x63, 0xc1, 0x75, 0x21, 0x81, 0x4b,
	0x0d, 0x3b, 0x05, 0x4e, 0xb1, 0x25, 0x20, 0x6b, 0xaa, 0x32, 0x77, 0x10, 0xd4, 0xbc, 0x4a, 0xb5,
	0xc9, 0x84, 0x2c, 0xa9, 0x76, 0x6b, 0xc6, 0x85, 0xa1, 0x8a, 0x91, 0x82, 0x42, 0x49, 0x0c, 0xc9,
	0x1a, 0xc5, 0xb4, 0x5d, 0x32, 0x77, 0xad, 0xf6, 0xdb, 0x58, 0x23, 0x7c, 0x60, 0x07, 0x77, 0xb7,
	0x4f, 0x4b, 0x7c, 0x81, 0x22, 0x7b, 0x05, 0x08, 0xb2, 0xa5, 0x49, 0x90, 0x06, 0x93, 0x28, 0x0f,
	0x2d, 0xb8, 0x27, 0x5b, 0x8a, 0xcf, 0x51, 0xd8, 0x28, 0xe6, 0x6b, 0x1d, 0x57, 0xeb, 0x35, 0x8a,
	0xb9, 0xd2, 0x15, 0x1a, 0xfe, 0xba, 0x3a, 0x41, 0xd7, 0x09, 0xe2, 0x1f, 0x6a, 0x65, 0xe3, 0xaf,
	0x0e, 0x8a, 0x0f, 0x5c, 0xf1, 0x08, 0xa1, 0x8a, 0x0b, 0x0a, 0xda, 0x10, 0x43, 0x93, 0xa9, 0x6b,
	0x8a, 0x2c, 0x59, 0x5a, 0x80, 0xdf, 0x03, 0x34, 0x20, 0x65, 0xa9, 0xa8, 0xd6, 0x50, 0x71, 0x6d,
	0x92, 0x59, 0xda, 0x9d, 0xf4, 0xa7, 0x2c, 0xfb, 0x9b, 0x57, 0xf1, 0xcd, 0xd6, 0x10, 0x6c, 0x4f,
	0xde, 0xdf, 0x7b, 0x2f, 0xb8, 0x36, 0xf8, 0x23, 0x40, 0xb8, 0xe2, 0x62, 0x03, 0x95, 0x2c, 0x48,
	0x05, 0xfb, 0x52, 0x72, 0x9d, 0x06, 0xff, 0x38, 0xd1, 0xa9, 0x9d, 0x60, 0x61, 0x07, 0x98, 0x7b,
	0xff, 0xf1, 0x67, 0x80, 0x86, 0x87, 0x22, 0x9c, 0xa0, 0x5e, 0x3b, 0x9d, 0xcf, 0xb0, 0x3d, 0xe2,
	0x4b, 0x14, 0xd7, 0x8a, 0x32, 0xbe, 0x83, 0x8a, 0x8a, 0xb5, 0x79, 0x76, 0x39, 0xc6, 0xf9, 0xc0,
	0xc3, 0x85, 0x63, 0x56, 0xd4, 0xbe, 0xb9, 0x8f, 0xc5, 0x67, 0xd9, 0x06, 0xe1, 0x93, 0x19, 0x21,
	0xc4, 0x35, 0x10, 0xf1, 0x56, 0x10, 0x6d, 0x92, 0xa3, 0x34, 0x98, 0x84, 0x79, 0xc4, 0xf5, 0xdc,
	0x03, 0xfb, 0x91, 0x94, 0x7c, 0x31, 0x14, 0x0c, 0x59, 0x27, 0xc7, 0xce, 0x24, 0x74, 0xe0, 0x81,
	0xac, 0x57, 0x27, 0xee, 0x6b, 0xcf, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0xc5, 0xbb, 0x7e, 0x25,
	0xf1, 0x02, 0x00, 0x00,
}
