// Code generated by protoc-gen-go. DO NOT EDIT.
// source: telemetry.proto

/*
Package telemetry is a generated protocol buffer package.

It is generated from these files:
	telemetry.proto

It has these top-level messages:
	Telemetry
	TelemetryField
	TelemetryGPBTable
	TelemetryRowGPB
*/
package telemetry

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

// Telemetry message is the outermost payload message used to stream
// telemetry in a Model Driven Telemetry (MDT) system. MDT provides a
// mechanism for an external entity to subscribe to a data set defined in
// a Yang model and receive periodic or event-based updates of the data
// set from an MDT-capable device.
type Telemetry struct {
	//
	// node_id_str is a string encoded unique node ID of the MDT-capable
	// device producing the message. (node_id_uuid alternative is not currently
	// produced in IOS-XR)
	//
	// Types that are valid to be assigned to NodeId:
	//	*Telemetry_NodeIdStr
	NodeId isTelemetry_NodeId `protobuf_oneof:"node_id"`
	//
	// subscription_id_str is the name of the subscription against which
	// this content is being produced. (subscription_id alternative is not
	//  currently produced in IOS-XR)
	//
	// Types that are valid to be assigned to Subscription:
	//	*Telemetry_SubscriptionIdStr
	Subscription isTelemetry_Subscription `protobuf_oneof:"subscription"`
	//
	// sensor_path is not currently produced in IOS-XR
	// string   sensor_path = 5;
	//
	// encoding_path is the Yang path leading to the content in this message.
	// The Yang tree encoded in the content section of this message is rooted
	// at the point described by the encoding_path.
	EncodingPath string `protobuf:"bytes,6,opt,name=encoding_path,json=encodingPath" json:"encoding_path,omitempty"`
	//
	// model_version is not currently produced in IOS-XR
	// string   model_version = 7;
	//
	// collection_id identifies messages belonging to a collection round.
	// Multiple message may be generated from a collection round.
	CollectionId uint64 `protobuf:"varint,8,opt,name=collection_id,json=collectionId" json:"collection_id,omitempty"`
	//
	// collection_start_time is the time when the collection identified by
	// the collection_id begins - encoded as milliseconds since the epoch.
	// If a single collection is spread over multiple Telemetry Messages,
	// collection_start_time may be encoded in the first Telemetry Message
	// for the collection only.
	CollectionStartTime uint64 `protobuf:"varint,9,opt,name=collection_start_time,json=collectionStartTime" json:"collection_start_time,omitempty"`
	//
	// msg_timestamp is the time when the data encoded in the Telemetry
	// message is generated - encoded as milliseconds since the epoch.
	MsgTimestamp uint64 `protobuf:"varint,10,opt,name=msg_timestamp,json=msgTimestamp" json:"msg_timestamp,omitempty"`
	//
	// data_gpbkv contains the payload data if data is being encoded in the
	// self-describing GPB-KV format.
	DataGpbkv []*TelemetryField `protobuf:"bytes,11,rep,name=data_gpbkv,json=dataGpbkv" json:"data_gpbkv,omitempty"`
	//
	// data_gpb contains the payload data if data is being encoded as
	// serialised GPB messages.
	DataGpb *TelemetryGPBTable `protobuf:"bytes,12,opt,name=data_gpb,json=dataGpb" json:"data_gpb,omitempty"`
	//
	// collection_end_time is the timestamp when the last Telemetry message
	// for a collection has been encoded - encoded as milliseconds since the
	// epoch. If a single collection is spread over multiple Telemetry
	// messages, collection_end_time is encoded in the last Telemetry Message
	// for the collection only.
	CollectionEndTime uint64 `protobuf:"varint,13,opt,name=collection_end_time,json=collectionEndTime" json:"collection_end_time,omitempty"`
}

func (m *Telemetry) Reset()                    { *m = Telemetry{} }
func (m *Telemetry) String() string            { return proto.CompactTextString(m) }
func (*Telemetry) ProtoMessage()               {}
func (*Telemetry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isTelemetry_NodeId interface {
	isTelemetry_NodeId()
}
type isTelemetry_Subscription interface {
	isTelemetry_Subscription()
}

type Telemetry_NodeIdStr struct {
	NodeIdStr string `protobuf:"bytes,1,opt,name=node_id_str,json=nodeIdStr,oneof"`
}
type Telemetry_SubscriptionIdStr struct {
	SubscriptionIdStr string `protobuf:"bytes,3,opt,name=subscription_id_str,json=subscriptionIdStr,oneof"`
}

func (*Telemetry_NodeIdStr) isTelemetry_NodeId()               {}
func (*Telemetry_SubscriptionIdStr) isTelemetry_Subscription() {}

func (m *Telemetry) GetNodeId() isTelemetry_NodeId {
	if m != nil {
		return m.NodeId
	}
	return nil
}
func (m *Telemetry) GetSubscription() isTelemetry_Subscription {
	if m != nil {
		return m.Subscription
	}
	return nil
}

func (m *Telemetry) GetNodeIdStr() string {
	if x, ok := m.GetNodeId().(*Telemetry_NodeIdStr); ok {
		return x.NodeIdStr
	}
	return ""
}

func (m *Telemetry) GetSubscriptionIdStr() string {
	if x, ok := m.GetSubscription().(*Telemetry_SubscriptionIdStr); ok {
		return x.SubscriptionIdStr
	}
	return ""
}

func (m *Telemetry) GetEncodingPath() string {
	if m != nil {
		return m.EncodingPath
	}
	return ""
}

func (m *Telemetry) GetCollectionId() uint64 {
	if m != nil {
		return m.CollectionId
	}
	return 0
}

func (m *Telemetry) GetCollectionStartTime() uint64 {
	if m != nil {
		return m.CollectionStartTime
	}
	return 0
}

func (m *Telemetry) GetMsgTimestamp() uint64 {
	if m != nil {
		return m.MsgTimestamp
	}
	return 0
}

func (m *Telemetry) GetDataGpbkv() []*TelemetryField {
	if m != nil {
		return m.DataGpbkv
	}
	return nil
}

func (m *Telemetry) GetDataGpb() *TelemetryGPBTable {
	if m != nil {
		return m.DataGpb
	}
	return nil
}

func (m *Telemetry) GetCollectionEndTime() uint64 {
	if m != nil {
		return m.CollectionEndTime
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Telemetry) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Telemetry_OneofMarshaler, _Telemetry_OneofUnmarshaler, _Telemetry_OneofSizer, []interface{}{
		(*Telemetry_NodeIdStr)(nil),
		(*Telemetry_SubscriptionIdStr)(nil),
	}
}

func _Telemetry_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Telemetry)
	// node_id
	switch x := m.NodeId.(type) {
	case *Telemetry_NodeIdStr:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.NodeIdStr)
	case nil:
	default:
		return fmt.Errorf("Telemetry.NodeId has unexpected type %T", x)
	}
	// subscription
	switch x := m.Subscription.(type) {
	case *Telemetry_SubscriptionIdStr:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.SubscriptionIdStr)
	case nil:
	default:
		return fmt.Errorf("Telemetry.Subscription has unexpected type %T", x)
	}
	return nil
}

func _Telemetry_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Telemetry)
	switch tag {
	case 1: // node_id.node_id_str
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.NodeId = &Telemetry_NodeIdStr{x}
		return true, err
	case 3: // subscription.subscription_id_str
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Subscription = &Telemetry_SubscriptionIdStr{x}
		return true, err
	default:
		return false, nil
	}
}

func _Telemetry_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Telemetry)
	// node_id
	switch x := m.NodeId.(type) {
	case *Telemetry_NodeIdStr:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.NodeIdStr)))
		n += len(x.NodeIdStr)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	// subscription
	switch x := m.Subscription.(type) {
	case *Telemetry_SubscriptionIdStr:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.SubscriptionIdStr)))
		n += len(x.SubscriptionIdStr)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

//
// TelemetryField messages are used to export content in the self
// describing GPB KV form. The TelemetryField message is sufficient to
// decode telemetry messages for all models. KV-GPB encoding is very
// similar in concept, to JSON encoding
type TelemetryField struct {
	//
	// timestamp represents the starting time of the generation of data
	// starting from this key, value pair in this message - encoded as
	// milliseconds since the epoch. It is encoded when different from the
	// msg_timestamp in the containing Telemetry Message. This field can be
	// omitted if the value is the same as a TelemetryField message up the
	// hierarchy within the same Telemetry Message as well.
	Timestamp uint64 `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	//
	// name: string encoding of the name in the key, value pair. It is
	// the corresponding YANG element name.
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	//
	// value_by_type, if present, for the corresponding YANG element
	// represented by the name field in the same TelemetryField message. The
	// value is encoded to the matching type as defined in the YANG model.
	// YANG models often define new types (derived types) using one or more
	// base types.  The types included in the oneof grouping is sufficient to
	// represent such derived types. Derived types represented as a Yang
	// container are encoded using the nesting primitive defined in this
	// encoding proposal.
	//
	// Types that are valid to be assigned to ValueByType:
	//	*TelemetryField_BytesValue
	//	*TelemetryField_StringValue
	//	*TelemetryField_BoolValue
	//	*TelemetryField_Uint32Value
	//	*TelemetryField_Uint64Value
	//	*TelemetryField_Sint32Value
	//	*TelemetryField_Sint64Value
	//	*TelemetryField_DoubleValue
	//	*TelemetryField_FloatValue
	ValueByType isTelemetryField_ValueByType `protobuf_oneof:"value_by_type"`
	//
	// The Yang model may include nesting (e.g hierarchy of containers). The
	// next level of nesting, if present, is encoded, starting from fields.
	Fields []*TelemetryField `protobuf:"bytes,15,rep,name=fields" json:"fields,omitempty"`
}

func (m *TelemetryField) Reset()                    { *m = TelemetryField{} }
func (m *TelemetryField) String() string            { return proto.CompactTextString(m) }
func (*TelemetryField) ProtoMessage()               {}
func (*TelemetryField) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isTelemetryField_ValueByType interface {
	isTelemetryField_ValueByType()
}

type TelemetryField_BytesValue struct {
	BytesValue []byte `protobuf:"bytes,4,opt,name=bytes_value,json=bytesValue,proto3,oneof"`
}
type TelemetryField_StringValue struct {
	StringValue string `protobuf:"bytes,5,opt,name=string_value,json=stringValue,oneof"`
}
type TelemetryField_BoolValue struct {
	BoolValue bool `protobuf:"varint,6,opt,name=bool_value,json=boolValue,oneof"`
}
type TelemetryField_Uint32Value struct {
	Uint32Value uint32 `protobuf:"varint,7,opt,name=uint32_value,json=uint32Value,oneof"`
}
type TelemetryField_Uint64Value struct {
	Uint64Value uint64 `protobuf:"varint,8,opt,name=uint64_value,json=uint64Value,oneof"`
}
type TelemetryField_Sint32Value struct {
	Sint32Value int32 `protobuf:"zigzag32,9,opt,name=sint32_value,json=sint32Value,oneof"`
}
type TelemetryField_Sint64Value struct {
	Sint64Value int64 `protobuf:"zigzag64,10,opt,name=sint64_value,json=sint64Value,oneof"`
}
type TelemetryField_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,11,opt,name=double_value,json=doubleValue,oneof"`
}
type TelemetryField_FloatValue struct {
	FloatValue float32 `protobuf:"fixed32,12,opt,name=float_value,json=floatValue,oneof"`
}

func (*TelemetryField_BytesValue) isTelemetryField_ValueByType()  {}
func (*TelemetryField_StringValue) isTelemetryField_ValueByType() {}
func (*TelemetryField_BoolValue) isTelemetryField_ValueByType()   {}
func (*TelemetryField_Uint32Value) isTelemetryField_ValueByType() {}
func (*TelemetryField_Uint64Value) isTelemetryField_ValueByType() {}
func (*TelemetryField_Sint32Value) isTelemetryField_ValueByType() {}
func (*TelemetryField_Sint64Value) isTelemetryField_ValueByType() {}
func (*TelemetryField_DoubleValue) isTelemetryField_ValueByType() {}
func (*TelemetryField_FloatValue) isTelemetryField_ValueByType()  {}

func (m *TelemetryField) GetValueByType() isTelemetryField_ValueByType {
	if m != nil {
		return m.ValueByType
	}
	return nil
}

func (m *TelemetryField) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *TelemetryField) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TelemetryField) GetBytesValue() []byte {
	if x, ok := m.GetValueByType().(*TelemetryField_BytesValue); ok {
		return x.BytesValue
	}
	return nil
}

func (m *TelemetryField) GetStringValue() string {
	if x, ok := m.GetValueByType().(*TelemetryField_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *TelemetryField) GetBoolValue() bool {
	if x, ok := m.GetValueByType().(*TelemetryField_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (m *TelemetryField) GetUint32Value() uint32 {
	if x, ok := m.GetValueByType().(*TelemetryField_Uint32Value); ok {
		return x.Uint32Value
	}
	return 0
}

func (m *TelemetryField) GetUint64Value() uint64 {
	if x, ok := m.GetValueByType().(*TelemetryField_Uint64Value); ok {
		return x.Uint64Value
	}
	return 0
}

func (m *TelemetryField) GetSint32Value() int32 {
	if x, ok := m.GetValueByType().(*TelemetryField_Sint32Value); ok {
		return x.Sint32Value
	}
	return 0
}

func (m *TelemetryField) GetSint64Value() int64 {
	if x, ok := m.GetValueByType().(*TelemetryField_Sint64Value); ok {
		return x.Sint64Value
	}
	return 0
}

func (m *TelemetryField) GetDoubleValue() float64 {
	if x, ok := m.GetValueByType().(*TelemetryField_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (m *TelemetryField) GetFloatValue() float32 {
	if x, ok := m.GetValueByType().(*TelemetryField_FloatValue); ok {
		return x.FloatValue
	}
	return 0
}

func (m *TelemetryField) GetFields() []*TelemetryField {
	if m != nil {
		return m.Fields
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TelemetryField) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TelemetryField_OneofMarshaler, _TelemetryField_OneofUnmarshaler, _TelemetryField_OneofSizer, []interface{}{
		(*TelemetryField_BytesValue)(nil),
		(*TelemetryField_StringValue)(nil),
		(*TelemetryField_BoolValue)(nil),
		(*TelemetryField_Uint32Value)(nil),
		(*TelemetryField_Uint64Value)(nil),
		(*TelemetryField_Sint32Value)(nil),
		(*TelemetryField_Sint64Value)(nil),
		(*TelemetryField_DoubleValue)(nil),
		(*TelemetryField_FloatValue)(nil),
	}
}

func _TelemetryField_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TelemetryField)
	// value_by_type
	switch x := m.ValueByType.(type) {
	case *TelemetryField_BytesValue:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.BytesValue)
	case *TelemetryField_StringValue:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.StringValue)
	case *TelemetryField_BoolValue:
		t := uint64(0)
		if x.BoolValue {
			t = 1
		}
		b.EncodeVarint(6<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *TelemetryField_Uint32Value:
		b.EncodeVarint(7<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Uint32Value))
	case *TelemetryField_Uint64Value:
		b.EncodeVarint(8<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Uint64Value))
	case *TelemetryField_Sint32Value:
		b.EncodeVarint(9<<3 | proto.WireVarint)
		b.EncodeZigzag32(uint64(x.Sint32Value))
	case *TelemetryField_Sint64Value:
		b.EncodeVarint(10<<3 | proto.WireVarint)
		b.EncodeZigzag64(uint64(x.Sint64Value))
	case *TelemetryField_DoubleValue:
		b.EncodeVarint(11<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.DoubleValue))
	case *TelemetryField_FloatValue:
		b.EncodeVarint(12<<3 | proto.WireFixed32)
		b.EncodeFixed32(uint64(math.Float32bits(x.FloatValue)))
	case nil:
	default:
		return fmt.Errorf("TelemetryField.ValueByType has unexpected type %T", x)
	}
	return nil
}

func _TelemetryField_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TelemetryField)
	switch tag {
	case 4: // value_by_type.bytes_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.ValueByType = &TelemetryField_BytesValue{x}
		return true, err
	case 5: // value_by_type.string_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.ValueByType = &TelemetryField_StringValue{x}
		return true, err
	case 6: // value_by_type.bool_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ValueByType = &TelemetryField_BoolValue{x != 0}
		return true, err
	case 7: // value_by_type.uint32_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ValueByType = &TelemetryField_Uint32Value{uint32(x)}
		return true, err
	case 8: // value_by_type.uint64_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ValueByType = &TelemetryField_Uint64Value{x}
		return true, err
	case 9: // value_by_type.sint32_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeZigzag32()
		m.ValueByType = &TelemetryField_Sint32Value{int32(x)}
		return true, err
	case 10: // value_by_type.sint64_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeZigzag64()
		m.ValueByType = &TelemetryField_Sint64Value{int64(x)}
		return true, err
	case 11: // value_by_type.double_value
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.ValueByType = &TelemetryField_DoubleValue{math.Float64frombits(x)}
		return true, err
	case 12: // value_by_type.float_value
		if wire != proto.WireFixed32 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed32()
		m.ValueByType = &TelemetryField_FloatValue{math.Float32frombits(uint32(x))}
		return true, err
	default:
		return false, nil
	}
}

func _TelemetryField_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TelemetryField)
	// value_by_type
	switch x := m.ValueByType.(type) {
	case *TelemetryField_BytesValue:
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.BytesValue)))
		n += len(x.BytesValue)
	case *TelemetryField_StringValue:
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.StringValue)))
		n += len(x.StringValue)
	case *TelemetryField_BoolValue:
		n += proto.SizeVarint(6<<3 | proto.WireVarint)
		n += 1
	case *TelemetryField_Uint32Value:
		n += proto.SizeVarint(7<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Uint32Value))
	case *TelemetryField_Uint64Value:
		n += proto.SizeVarint(8<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Uint64Value))
	case *TelemetryField_Sint32Value:
		n += proto.SizeVarint(9<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64((uint32(x.Sint32Value) << 1) ^ uint32((int32(x.Sint32Value) >> 31))))
	case *TelemetryField_Sint64Value:
		n += proto.SizeVarint(10<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(uint64(x.Sint64Value<<1) ^ uint64((int64(x.Sint64Value) >> 63))))
	case *TelemetryField_DoubleValue:
		n += proto.SizeVarint(11<<3 | proto.WireFixed64)
		n += 8
	case *TelemetryField_FloatValue:
		n += proto.SizeVarint(12<<3 | proto.WireFixed32)
		n += 4
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// TelemetryGPBTable contains a repeated number of TelemetryRowGPB,
// each of which represents content from a subtree instance in the
// the YANG model. For example; a TelemetryGPBTable might contain
// the interface statistics of a collection of interfaces.
type TelemetryGPBTable struct {
	Row []*TelemetryRowGPB `protobuf:"bytes,1,rep,name=row" json:"row,omitempty"`
}

func (m *TelemetryGPBTable) Reset()                    { *m = TelemetryGPBTable{} }
func (m *TelemetryGPBTable) String() string            { return proto.CompactTextString(m) }
func (*TelemetryGPBTable) ProtoMessage()               {}
func (*TelemetryGPBTable) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TelemetryGPBTable) GetRow() []*TelemetryRowGPB {
	if m != nil {
		return m.Row
	}
	return nil
}

//
// TelemetryRowGPB, in conjunction with the Telemetry encoding_path and
// model_version, unambiguously represents the root of a subtree in
// the YANG model, and content from that subtree encoded in serialised
// GPB messages. For example; a TelemetryRowGPB might contain the
// interface statistics of one interface. Per encoding-path .proto
// messages are required to decode keys/content pairs below.
type TelemetryRowGPB struct {
	//
	// timestamp at which the data for this instance of the TelemetryRowGPB
	// message was generated by an MDT-capable device - encoded as
	// milliseconds since the epoch.  When included, this is typically
	// different from the msg_timestamp in the containing Telemetry message.
	Timestamp uint64 `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	//
	// keys: if the encoding-path includes one or more list elements, and/or
	// ends in a list element, the keys field is a GPB encoded message that
	// contains the sequence of key values for each such list element in the
	// encoding-path traversed starting from the root.  The set of keys
	// unambiguously identifies the instance of data encoded in the
	// TelemetryRowGPB message. Corresponding protobuf message definition will
	// be required to decode the byte stream. The encoding_path field in
	// Telemetry message, together with model_version field should be
	// sufficient to identify the corresponding protobuf message.
	Keys []byte `protobuf:"bytes,10,opt,name=keys,proto3" json:"keys,omitempty"`
	//
	// content: the content field is a GPB encoded message that contains the
	// data for the corresponding encoding-path. A separate decoding pass
	// would be performed by consumer with the content field as a GPB message
	// and the matching .proto used to decode the message. Corresponding
	// protobuf message definition will be required to decode the byte
	// stream. The encoding_path field in Telemetry message, together with
	// model_version field should be sufficient to identify the corresponding
	// protobuf message. The decoded combination of keys (when present) and
	// content, unambiguously represents an instance of the data set, as
	// defined in the Yang model, identified by the encoding-path in the
	// containing Telemetry message.
	Content []byte `protobuf:"bytes,11,opt,name=content,proto3" json:"content,omitempty"`
}

func (m *TelemetryRowGPB) Reset()                    { *m = TelemetryRowGPB{} }
func (m *TelemetryRowGPB) String() string            { return proto.CompactTextString(m) }
func (*TelemetryRowGPB) ProtoMessage()               {}
func (*TelemetryRowGPB) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TelemetryRowGPB) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *TelemetryRowGPB) GetKeys() []byte {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *TelemetryRowGPB) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func init() {
	proto.RegisterType((*Telemetry)(nil), "telemetry.Telemetry")
	proto.RegisterType((*TelemetryField)(nil), "telemetry.TelemetryField")
	proto.RegisterType((*TelemetryGPBTable)(nil), "telemetry.TelemetryGPBTable")
	proto.RegisterType((*TelemetryRowGPB)(nil), "telemetry.TelemetryRowGPB")
}

func init() { proto.RegisterFile("telemetry.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 546 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0x41, 0x6f, 0x9b, 0x40,
	0x10, 0x85, 0xb3, 0xb1, 0x6b, 0x9b, 0x01, 0xc7, 0xf2, 0x46, 0x95, 0xb6, 0x55, 0xa4, 0x12, 0xe7,
	0xc2, 0xa1, 0xb2, 0x5a, 0x27, 0x4a, 0x7b, 0xad, 0xa5, 0xd6, 0xc9, 0xcd, 0xda, 0x58, 0xbd, 0x55,
	0x08, 0xcc, 0xc6, 0x41, 0x01, 0x16, 0xb1, 0xeb, 0x44, 0xfe, 0x77, 0xbd, 0xf7, 0x4f, 0x55, 0xb3,
	0x80, 0xc1, 0x6d, 0xa4, 0xdc, 0xe0, 0xcd, 0x37, 0x4f, 0xc3, 0xcc, 0x13, 0x30, 0xd2, 0x22, 0x11,
	0xa9, 0xd0, 0xc5, 0x6e, 0x9a, 0x17, 0x52, 0x4b, 0x6a, 0xed, 0x85, 0xc9, 0xef, 0x0e, 0x58, 0xab,
	0xfa, 0x8d, 0xba, 0x60, 0x67, 0x32, 0x12, 0x7e, 0x1c, 0xf9, 0x4a, 0x17, 0x8c, 0xb8, 0xc4, 0xb3,
	0x6e, 0x8e, 0xb8, 0x85, 0xe2, 0x6d, 0x74, 0xa7, 0x0b, 0xfa, 0x09, 0x4e, 0xd5, 0x36, 0x54, 0xeb,
	0x22, 0xce, 0x75, 0x2c, 0xb3, 0x9a, 0xec, 0x18, 0x92, 0xf0, 0x71, 0xbb, 0x58, 0x76, 0x5c, 0xc0,
	0x50, 0x64, 0x6b, 0x19, 0xc5, 0xd9, 0xc6, 0xcf, 0x03, 0xfd, 0xc0, 0x7a, 0xc8, 0x72, 0xa7, 0x16,
	0x97, 0x81, 0x7e, 0x40, 0x68, 0x2d, 0x93, 0x44, 0xac, 0x2b, 0x53, 0x36, 0x70, 0x89, 0xd7, 0xe5,
	0x4e, 0x23, 0xde, 0x46, 0x74, 0x06, 0x6f, 0x5b, 0x90, 0xd2, 0x41, 0xa1, 0x7d, 0x1d, 0xa7, 0x82,
	0x59, 0x06, 0x3e, 0x6d, 0x8a, 0x77, 0x58, 0x5b, 0xc5, 0xa9, 0x40, 0xe3, 0x54, 0x6d, 0x0c, 0xa6,
	0x74, 0x90, 0xe6, 0x0c, 0x4a, 0xe3, 0x54, 0x6d, 0x56, 0xb5, 0x46, 0xbf, 0x02, 0x44, 0x81, 0x0e,
	0xfc, 0x4d, 0x1e, 0x3e, 0x3e, 0x31, 0xdb, 0xed, 0x78, 0xf6, 0xec, 0xdd, 0xb4, 0xd9, 0xda, 0x7e,
	0x41, 0x3f, 0x62, 0x91, 0x44, 0xdc, 0x42, 0x78, 0x81, 0x2c, 0xfd, 0x02, 0x83, 0xba, 0x93, 0x39,
	0x2e, 0xf1, 0xec, 0xd9, 0xd9, 0x4b, 0x7d, 0x8b, 0xe5, 0x7c, 0x15, 0x84, 0x89, 0xe0, 0xfd, 0xaa,
	0x95, 0x4e, 0xa1, 0x35, 0xae, 0x2f, 0xb2, 0xa8, 0xfc, 0x92, 0xa1, 0x99, 0x6e, 0xdc, 0x94, 0xbe,
	0x67, 0x11, 0xce, 0x39, 0xb7, 0xa0, 0x5f, 0x5d, 0x66, 0x7e, 0x02, 0x4e, 0x7b, 0xcb, 0x93, 0x3f,
	0x1d, 0x38, 0x39, 0x9c, 0x90, 0x9e, 0x81, 0xd5, 0x7c, 0x31, 0x31, 0x9e, 0x8d, 0x40, 0x29, 0x74,
	0xb3, 0x20, 0x15, 0xec, 0xd8, 0x1c, 0xc2, 0x3c, 0xd3, 0x73, 0xb0, 0xc3, 0x9d, 0x16, 0xca, 0x7f,
	0x0a, 0x92, 0xad, 0x60, 0x5d, 0x97, 0x78, 0xce, 0xcd, 0x11, 0x07, 0x23, 0xfe, 0x44, 0x8d, 0x5e,
	0x80, 0xa3, 0x74, 0x81, 0x67, 0x2c, 0x99, 0x37, 0x55, 0x3a, 0xec, 0x52, 0x2d, 0xa1, 0x0f, 0x00,
	0xa1, 0x94, 0x49, 0x85, 0xe0, 0xa9, 0x07, 0x18, 0x20, 0xd4, 0xf6, 0x2e, 0xdb, 0x38, 0xd3, 0x97,
	0xb3, 0x0a, 0xe9, 0xbb, 0xc4, 0x1b, 0xa2, 0x4b, 0xa9, 0x1e, 0x40, 0xd7, 0x57, 0x15, 0x64, 0xd2,
	0x50, 0x43, 0xd7, 0x57, 0xcd, 0x3c, 0x6d, 0x27, 0x4c, 0xc1, 0xd8, 0xcc, 0x73, 0xe8, 0xa4, 0xda,
	0x4e, 0x78, 0x7e, 0x5a, 0x43, 0x2d, 0xa7, 0x48, 0x6e, 0xc3, 0x44, 0x54, 0x90, 0xed, 0x12, 0x8f,
	0x20, 0x54, 0xaa, 0x25, 0x74, 0x0e, 0xf6, 0x7d, 0x22, 0x03, 0x5d, 0x31, 0x78, 0xed, 0x63, 0xdc,
	0x90, 0x11, 0x4b, 0xe4, 0x33, 0xf4, 0xee, 0x71, 0xff, 0x8a, 0x8d, 0x5e, 0xcb, 0x50, 0x05, 0xce,
	0x47, 0x30, 0x34, 0x7e, 0x7e, 0xb8, 0xf3, 0xf5, 0x2e, 0x17, 0x93, 0x6f, 0x30, 0xfe, 0x2f, 0x36,
	0xf4, 0x23, 0x74, 0x0a, 0xf9, 0xcc, 0x88, 0x71, 0x7d, 0xff, 0x92, 0x2b, 0x97, 0xcf, 0x8b, 0xe5,
	0x9c, 0x23, 0x36, 0xf9, 0x05, 0xa3, 0x7f, 0xf4, 0xd7, 0x03, 0xf1, 0x28, 0x76, 0xca, 0x2c, 0xc7,
	0xe1, 0xe6, 0x99, 0x32, 0xe8, 0xaf, 0x65, 0xa6, 0x45, 0xa6, 0xcd, 0x3a, 0x1c, 0x5e, 0xbf, 0x86,
	0x3d, 0xf3, 0x13, 0xb9, 0xfc, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x43, 0x6f, 0x37, 0x57, 0x04,
	0x00, 0x00,
}
