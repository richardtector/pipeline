// Code generated by protoc-gen-go. DO NOT EDIT.
// source: isis_sh_chkpt_lsp.proto

/*
Package cisco_ios_xr_clns_isis_oper_isis_instances_instance_checkpoint_lsps_checkpoint_lsp is a generated protocol buffer package.

It is generated from these files:
	isis_sh_chkpt_lsp.proto

It has these top-level messages:
	IsisShChkptLsp_KEYS
	IsisShChkptLsp
*/
package cisco_ios_xr_clns_isis_oper_isis_instances_instance_checkpoint_lsps_checkpoint_lsp

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

// Checkpointed LSP
type IsisShChkptLsp_KEYS struct {
	InstanceName string `protobuf:"bytes,1,opt,name=instance_name,json=instanceName" json:"instance_name,omitempty"`
	Level        string `protobuf:"bytes,2,opt,name=level" json:"level,omitempty"`
	LspId        string `protobuf:"bytes,3,opt,name=lsp_id,json=lspId" json:"lsp_id,omitempty"`
}

func (m *IsisShChkptLsp_KEYS) Reset()                    { *m = IsisShChkptLsp_KEYS{} }
func (m *IsisShChkptLsp_KEYS) String() string            { return proto.CompactTextString(m) }
func (*IsisShChkptLsp_KEYS) ProtoMessage()               {}
func (*IsisShChkptLsp_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IsisShChkptLsp_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *IsisShChkptLsp_KEYS) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

func (m *IsisShChkptLsp_KEYS) GetLspId() string {
	if m != nil {
		return m.LspId
	}
	return ""
}

type IsisShChkptLsp struct {
	// LSP Checkpoint table object ID
	CheckpointLspObjectId uint32 `protobuf:"varint,50,opt,name=checkpoint_lsp_object_id,json=checkpointLspObjectId" json:"checkpoint_lsp_object_id,omitempty"`
	// LSP level
	CheckpointLspLevel uint32 `protobuf:"varint,51,opt,name=checkpoint_lsp_level,json=checkpointLspLevel" json:"checkpoint_lsp_level,omitempty"`
	// TRUE if this is a local LSP
	CheckpointLspLocalFlag bool `protobuf:"varint,52,opt,name=checkpoint_lsp_local_flag,json=checkpointLspLocalFlag" json:"checkpoint_lsp_local_flag,omitempty"`
	// LSP ID
	CheckpointLspId string `protobuf:"bytes,53,opt,name=checkpoint_lsp_id,json=checkpointLspId" json:"checkpoint_lsp_id,omitempty"`
}

func (m *IsisShChkptLsp) Reset()                    { *m = IsisShChkptLsp{} }
func (m *IsisShChkptLsp) String() string            { return proto.CompactTextString(m) }
func (*IsisShChkptLsp) ProtoMessage()               {}
func (*IsisShChkptLsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *IsisShChkptLsp) GetCheckpointLspObjectId() uint32 {
	if m != nil {
		return m.CheckpointLspObjectId
	}
	return 0
}

func (m *IsisShChkptLsp) GetCheckpointLspLevel() uint32 {
	if m != nil {
		return m.CheckpointLspLevel
	}
	return 0
}

func (m *IsisShChkptLsp) GetCheckpointLspLocalFlag() bool {
	if m != nil {
		return m.CheckpointLspLocalFlag
	}
	return false
}

func (m *IsisShChkptLsp) GetCheckpointLspId() string {
	if m != nil {
		return m.CheckpointLspId
	}
	return ""
}

func init() {
	proto.RegisterType((*IsisShChkptLsp_KEYS)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.checkpoint_lsps.checkpoint_lsp.isis_sh_chkpt_lsp_KEYS")
	proto.RegisterType((*IsisShChkptLsp)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.checkpoint_lsps.checkpoint_lsp.isis_sh_chkpt_lsp")
}

func init() { proto.RegisterFile("isis_sh_chkpt_lsp.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0x59, 0xc5, 0xa2, 0x83, 0x45, 0x1a, 0xda, 0x1a, 0x6f, 0xa5, 0x5e, 0x16, 0x0f, 0x8b,
	0x58, 0x45, 0x7c, 0x00, 0x85, 0xc5, 0xa2, 0xb0, 0x9e, 0x3c, 0x0d, 0x69, 0x12, 0xbb, 0x69, 0xd3,
	0x24, 0xec, 0x2c, 0xe2, 0x0b, 0xfb, 0x1e, 0xb2, 0xa9, 0xa5, 0xec, 0xf6, 0x36, 0x33, 0x7f, 0xbe,
	0x7f, 0xfe, 0x30, 0x70, 0x69, 0xc8, 0x10, 0x52, 0x89, 0xb2, 0x5c, 0x87, 0x1a, 0x2d, 0x85, 0x2c,
	0x54, 0xbe, 0xf6, 0xac, 0x90, 0x86, 0xa4, 0x47, 0xe3, 0x09, 0x7f, 0x2a, 0x94, 0xd6, 0x11, 0xc6,
	0xa7, 0x3e, 0xe8, 0x2a, 0x6b, 0xaa, 0xcc, 0x38, 0xaa, 0x85, 0x93, 0x7a, 0x5f, 0x65, 0xb2, 0xd4,
	0x72, 0x1d, 0xbc, 0x71, 0xd1, 0x89, 0x3a, 0xfd, 0x74, 0x05, 0xe3, 0x83, 0x75, 0xf8, 0xfa, 0xfc,
	0xf9, 0xc1, 0xae, 0xa1, 0xbf, 0x33, 0x41, 0x27, 0x36, 0x9a, 0x27, 0x93, 0x24, 0x3d, 0x2b, 0xce,
	0x77, 0xc3, 0x37, 0xb1, 0xd1, 0x6c, 0x08, 0x27, 0x56, 0x7f, 0x6b, 0xcb, 0x8f, 0xa2, 0xb8, 0x6d,
	0xd8, 0x08, 0x7a, 0x8d, 0x8d, 0x51, 0xfc, 0xf8, 0x7f, 0x4c, 0x21, 0x57, 0xd3, 0xdf, 0x04, 0x06,
	0x07, 0xcb, 0xd8, 0x23, 0xf0, 0x76, 0x26, 0xf4, 0x8b, 0x95, 0x96, 0x75, 0x83, 0xdf, 0x4d, 0x92,
	0xb4, 0x5f, 0x8c, 0xf6, 0xfa, 0x9c, 0xc2, 0x7b, 0x54, 0x73, 0xc5, 0x6e, 0x61, 0xd8, 0x01, 0xb7,
	0x51, 0x66, 0x11, 0x62, 0x2d, 0x68, 0x1e, 0x73, 0x3d, 0xc1, 0x55, 0x97, 0xf0, 0x52, 0x58, 0xfc,
	0xb2, 0x62, 0xc9, 0xef, 0x27, 0x49, 0x7a, 0x5a, 0x8c, 0xdb, 0x58, 0x23, 0xbf, 0x58, 0xb1, 0x64,
	0x37, 0x30, 0xe8, 0xa0, 0x46, 0xf1, 0x87, 0xf8, 0xbb, 0x8b, 0x16, 0x92, 0xab, 0x45, 0x2f, 0x9e,
	0x6b, 0xf6, 0x17, 0x00, 0x00, 0xff, 0xff, 0x59, 0x16, 0x62, 0x2e, 0xc9, 0x01, 0x00, 0x00,
}
