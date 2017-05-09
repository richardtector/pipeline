// Code generated by protoc-gen-go.
// source: inv_xml_entity_basic_info.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_plat_chas_invmgr_oper_platform_inventory_racks_rack_attributes_basic_info is a generated protocol buffer package.

It is generated from these files:
	inv_xml_entity_basic_info.proto

It has these top-level messages:
	InvXmlEntityBasicInfo_KEYS
	InvXmlEntityBasicInfo
*/
package cisco_ios_xr_plat_chas_invmgr_oper_platform_inventory_racks_rack_attributes_basic_info

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

// Bag contains all the basic inventory information for each entity
type InvXmlEntityBasicInfo_KEYS struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *InvXmlEntityBasicInfo_KEYS) Reset()                    { *m = InvXmlEntityBasicInfo_KEYS{} }
func (m *InvXmlEntityBasicInfo_KEYS) String() string            { return proto.CompactTextString(m) }
func (*InvXmlEntityBasicInfo_KEYS) ProtoMessage()               {}
func (*InvXmlEntityBasicInfo_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *InvXmlEntityBasicInfo_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type InvXmlEntityBasicInfo struct {
	// name string for the entity
	Name string `protobuf:"bytes,50,opt,name=name" json:"name,omitempty"`
	// describes in user-readable terms what the entity in question does
	Description string `protobuf:"bytes,51,opt,name=description" json:"description,omitempty"`
	// model name
	ModelName string `protobuf:"bytes,52,opt,name=model_name,json=modelName" json:"model_name,omitempty"`
	// hw revision string
	HardwareRevision string `protobuf:"bytes,53,opt,name=hardware_revision,json=hardwareRevision" json:"hardware_revision,omitempty"`
	// serial number
	SerialNumber string `protobuf:"bytes,54,opt,name=serial_number,json=serialNumber" json:"serial_number,omitempty"`
	// firmware revision string
	FirmwareRevision string `protobuf:"bytes,55,opt,name=firmware_revision,json=firmwareRevision" json:"firmware_revision,omitempty"`
	// software revision string
	SoftwareRevision string `protobuf:"bytes,56,opt,name=software_revision,json=softwareRevision" json:"software_revision,omitempty"`
	// maps to the vendor OID string
	VendorType string `protobuf:"bytes,57,opt,name=vendor_type,json=vendorType" json:"vendor_type,omitempty"`
	// 1 if Field Replaceable Unit 0, if not
	IsFieldReplaceableUnit bool `protobuf:"varint,58,opt,name=is_field_replaceable_unit,json=isFieldReplaceableUnit" json:"is_field_replaceable_unit,omitempty"`
}

func (m *InvXmlEntityBasicInfo) Reset()                    { *m = InvXmlEntityBasicInfo{} }
func (m *InvXmlEntityBasicInfo) String() string            { return proto.CompactTextString(m) }
func (*InvXmlEntityBasicInfo) ProtoMessage()               {}
func (*InvXmlEntityBasicInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *InvXmlEntityBasicInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InvXmlEntityBasicInfo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *InvXmlEntityBasicInfo) GetModelName() string {
	if m != nil {
		return m.ModelName
	}
	return ""
}

func (m *InvXmlEntityBasicInfo) GetHardwareRevision() string {
	if m != nil {
		return m.HardwareRevision
	}
	return ""
}

func (m *InvXmlEntityBasicInfo) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

func (m *InvXmlEntityBasicInfo) GetFirmwareRevision() string {
	if m != nil {
		return m.FirmwareRevision
	}
	return ""
}

func (m *InvXmlEntityBasicInfo) GetSoftwareRevision() string {
	if m != nil {
		return m.SoftwareRevision
	}
	return ""
}

func (m *InvXmlEntityBasicInfo) GetVendorType() string {
	if m != nil {
		return m.VendorType
	}
	return ""
}

func (m *InvXmlEntityBasicInfo) GetIsFieldReplaceableUnit() bool {
	if m != nil {
		return m.IsFieldReplaceableUnit
	}
	return false
}

func init() {
	proto.RegisterType((*InvXmlEntityBasicInfo_KEYS)(nil), "cisco_ios_xr_plat_chas_invmgr_oper.platform_inventory.racks.rack.attributes.basic_info.inv_xml_entity_basic_info_KEYS")
	proto.RegisterType((*InvXmlEntityBasicInfo)(nil), "cisco_ios_xr_plat_chas_invmgr_oper.platform_inventory.racks.rack.attributes.basic_info.inv_xml_entity_basic_info")
}

func init() { proto.RegisterFile("inv_xml_entity_basic_info.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x3b, 0x4f, 0xc3, 0x30,
	0x14, 0x85, 0x55, 0x84, 0x10, 0x75, 0x41, 0xa2, 0x19, 0x50, 0x3a, 0x40, 0xab, 0xb2, 0x54, 0xaa,
	0x94, 0x81, 0x96, 0x47, 0xd9, 0x61, 0x41, 0xea, 0x10, 0x1e, 0x12, 0xd3, 0x95, 0x93, 0xdc, 0xd0,
	0x2b, 0x12, 0x3b, 0xba, 0x76, 0x43, 0xf3, 0xcf, 0x19, 0x51, 0x1c, 0xda, 0xd2, 0xa1, 0x8b, 0x65,
	0x7d, 0xe7, 0x9c, 0xcf, 0x8b, 0x45, 0x9f, 0x54, 0x09, 0xab, 0x3c, 0x03, 0x54, 0x96, 0x6c, 0x05,
	0x91, 0x34, 0x14, 0x03, 0xa9, 0x54, 0x07, 0x05, 0x6b, 0xab, 0xbd, 0xf7, 0x98, 0x4c, 0xac, 0x81,
	0xb4, 0x81, 0x15, 0x43, 0x91, 0x49, 0x0b, 0xf1, 0x42, 0x1a, 0x20, 0x55, 0xe6, 0x9f, 0x0c, 0xba,
	0x40, 0x0e, 0x6a, 0x9a, 0x6a, 0xce, 0x6b, 0x88, 0xca, 0x6a, 0xae, 0x02, 0x96, 0xf1, 0x97, 0x71,
	0x67, 0x20, 0xad, 0x65, 0x8a, 0x96, 0x16, 0x4d, 0xb0, 0xb5, 0x0f, 0xa7, 0xe2, 0x72, 0xef, 0xd3,
	0xf0, 0xfc, 0xf8, 0xf1, 0xe2, 0x79, 0xe2, 0x50, 0xc9, 0x1c, 0xfd, 0xd6, 0xa0, 0x35, 0x6a, 0x87,
	0xee, 0x3e, 0xfc, 0x39, 0x10, 0xbd, 0xbd, 0xb3, 0xcd, 0xe2, 0x7a, 0xbb, 0xf0, 0x06, 0xa2, 0x93,
	0xa0, 0x89, 0x99, 0x0a, 0x4b, 0x5a, 0xf9, 0x13, 0x17, 0xfd, 0x47, 0xde, 0x85, 0x10, 0xb9, 0x4e,
	0x30, 0x03, 0xb7, 0x9d, 0xba, 0x42, 0xdb, 0x91, 0x79, 0x2d, 0x18, 0x8b, 0xee, 0x42, 0x72, 0xf2,
	0x2d, 0x19, 0x81, 0xb1, 0x24, 0x53, 0x6b, 0x6e, 0x5c, 0xeb, 0x6c, 0x1d, 0x84, 0x7f, 0xdc, 0xbb,
	0x12, 0xa7, 0x06, 0x99, 0x64, 0x06, 0x6a, 0x99, 0x47, 0xc8, 0xfe, 0xad, 0x2b, 0x9e, 0x34, 0x70,
	0xee, 0x58, 0x6d, 0x4c, 0x89, 0xf3, 0x5d, 0xe3, 0x5d, 0x63, 0x5c, 0x07, 0x1b, 0xe3, 0x58, 0x74,
	0x8d, 0x4e, 0xed, 0x6e, 0xf9, 0xbe, 0x29, 0xaf, 0x83, 0x4d, 0xb9, 0x2f, 0x3a, 0x25, 0xaa, 0x44,
	0x33, 0xd8, 0xaa, 0x40, 0x7f, 0xe6, 0x6a, 0xa2, 0x41, 0xaf, 0x55, 0x81, 0xde, 0x4c, 0xf4, 0xc8,
	0x40, 0x4a, 0x98, 0x25, 0xc0, 0x58, 0x64, 0x32, 0x46, 0x19, 0x65, 0x08, 0x4b, 0x45, 0xd6, 0x7f,
	0x18, 0xb4, 0x46, 0xc7, 0xe1, 0x39, 0x99, 0xa7, 0x3a, 0x0f, 0xb7, 0xf1, 0x9b, 0x22, 0x1b, 0x1d,
	0xb9, 0xff, 0x30, 0xf9, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x24, 0xb8, 0xfb, 0xa8, 0x32, 0x02, 0x00,
	0x00,
}