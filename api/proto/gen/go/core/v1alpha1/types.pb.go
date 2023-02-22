// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: core/v1alpha1/types.proto

package corev1alpha1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Primitive int32

const (
	Primitive_PRIMITIVE_UNSPECIFIED Primitive = 0
	Primitive_PRIMITIVE_STRING      Primitive = 1
	Primitive_PRIMITIVE_INTEGER     Primitive = 2
	Primitive_PRIMITIVE_FLOAT       Primitive = 3
	Primitive_PRIMITIVE_BOOL        Primitive = 4
	Primitive_PRIMITIVE_TIMESTAMP   Primitive = 5
	// 6-9 Reserved for future use.
	Primitive_PRIMITIVE_STRING_LIST    Primitive = 10
	Primitive_PRIMITIVE_INTEGER_LIST   Primitive = 11
	Primitive_PRIMITIVE_FLOAT_LIST     Primitive = 12
	Primitive_PRIMITIVE_BOOL_LIST      Primitive = 13
	Primitive_PRIMITIVE_TIMESTAMP_LIST Primitive = 14
)

// Enum value maps for Primitive.
var (
	Primitive_name = map[int32]string{
		0:  "PRIMITIVE_UNSPECIFIED",
		1:  "PRIMITIVE_STRING",
		2:  "PRIMITIVE_INTEGER",
		3:  "PRIMITIVE_FLOAT",
		4:  "PRIMITIVE_BOOL",
		5:  "PRIMITIVE_TIMESTAMP",
		10: "PRIMITIVE_STRING_LIST",
		11: "PRIMITIVE_INTEGER_LIST",
		12: "PRIMITIVE_FLOAT_LIST",
		13: "PRIMITIVE_BOOL_LIST",
		14: "PRIMITIVE_TIMESTAMP_LIST",
	}
	Primitive_value = map[string]int32{
		"PRIMITIVE_UNSPECIFIED":    0,
		"PRIMITIVE_STRING":         1,
		"PRIMITIVE_INTEGER":        2,
		"PRIMITIVE_FLOAT":          3,
		"PRIMITIVE_BOOL":           4,
		"PRIMITIVE_TIMESTAMP":      5,
		"PRIMITIVE_STRING_LIST":    10,
		"PRIMITIVE_INTEGER_LIST":   11,
		"PRIMITIVE_FLOAT_LIST":     12,
		"PRIMITIVE_BOOL_LIST":      13,
		"PRIMITIVE_TIMESTAMP_LIST": 14,
	}
)

func (x Primitive) Enum() *Primitive {
	p := new(Primitive)
	*p = x
	return p
}

func (x Primitive) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Primitive) Descriptor() protoreflect.EnumDescriptor {
	return file_core_v1alpha1_types_proto_enumTypes[0].Descriptor()
}

func (Primitive) Type() protoreflect.EnumType {
	return &file_core_v1alpha1_types_proto_enumTypes[0]
}

func (x Primitive) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Primitive.Descriptor instead.
func (Primitive) EnumDescriptor() ([]byte, []int) {
	return file_core_v1alpha1_types_proto_rawDescGZIP(), []int{0}
}

type AggrFn int32

const (
	AggrFn_AGGR_FN_UNSPECIFIED AggrFn = 0
	AggrFn_AGGR_FN_SUM         AggrFn = 1
	AggrFn_AGGR_FN_AVG         AggrFn = 2
	AggrFn_AGGR_FN_MAX         AggrFn = 3
	AggrFn_AGGR_FN_MIN         AggrFn = 4
	AggrFn_AGGR_FN_COUNT       AggrFn = 5
)

// Enum value maps for AggrFn.
var (
	AggrFn_name = map[int32]string{
		0: "AGGR_FN_UNSPECIFIED",
		1: "AGGR_FN_SUM",
		2: "AGGR_FN_AVG",
		3: "AGGR_FN_MAX",
		4: "AGGR_FN_MIN",
		5: "AGGR_FN_COUNT",
	}
	AggrFn_value = map[string]int32{
		"AGGR_FN_UNSPECIFIED": 0,
		"AGGR_FN_SUM":         1,
		"AGGR_FN_AVG":         2,
		"AGGR_FN_MAX":         3,
		"AGGR_FN_MIN":         4,
		"AGGR_FN_COUNT":       5,
	}
)

func (x AggrFn) Enum() *AggrFn {
	p := new(AggrFn)
	*p = x
	return p
}

func (x AggrFn) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AggrFn) Descriptor() protoreflect.EnumDescriptor {
	return file_core_v1alpha1_types_proto_enumTypes[1].Descriptor()
}

func (AggrFn) Type() protoreflect.EnumType {
	return &file_core_v1alpha1_types_proto_enumTypes[1]
}

func (x AggrFn) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AggrFn.Descriptor instead.
func (AggrFn) EnumDescriptor() ([]byte, []int) {
	return file_core_v1alpha1_types_proto_rawDescGZIP(), []int{1}
}

type Scalar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//
	//	*Scalar_StringValue
	//	*Scalar_IntValue
	//	*Scalar_FloatValue
	//	*Scalar_BoolValue
	//	*Scalar_TimestampValue
	Value isScalar_Value `protobuf_oneof:"value"`
}

func (x *Scalar) Reset() {
	*x = Scalar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1alpha1_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scalar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scalar) ProtoMessage() {}

func (x *Scalar) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1alpha1_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scalar.ProtoReflect.Descriptor instead.
func (*Scalar) Descriptor() ([]byte, []int) {
	return file_core_v1alpha1_types_proto_rawDescGZIP(), []int{0}
}

func (m *Scalar) GetValue() isScalar_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *Scalar) GetStringValue() string {
	if x, ok := x.GetValue().(*Scalar_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (x *Scalar) GetIntValue() int32 {
	if x, ok := x.GetValue().(*Scalar_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (x *Scalar) GetFloatValue() float64 {
	if x, ok := x.GetValue().(*Scalar_FloatValue); ok {
		return x.FloatValue
	}
	return 0
}

func (x *Scalar) GetBoolValue() bool {
	if x, ok := x.GetValue().(*Scalar_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (x *Scalar) GetTimestampValue() *timestamppb.Timestamp {
	if x, ok := x.GetValue().(*Scalar_TimestampValue); ok {
		return x.TimestampValue
	}
	return nil
}

type isScalar_Value interface {
	isScalar_Value()
}

type Scalar_StringValue struct {
	StringValue string `protobuf:"bytes,1,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type Scalar_IntValue struct {
	IntValue int32 `protobuf:"varint,2,opt,name=int_value,json=intValue,proto3,oneof"`
}

type Scalar_FloatValue struct {
	FloatValue float64 `protobuf:"fixed64,3,opt,name=float_value,json=floatValue,proto3,oneof"`
}

type Scalar_BoolValue struct {
	BoolValue bool `protobuf:"varint,4,opt,name=bool_value,json=boolValue,proto3,oneof"`
}

type Scalar_TimestampValue struct {
	TimestampValue *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=timestamp_value,json=timestampValue,proto3,oneof"`
}

func (*Scalar_StringValue) isScalar_Value() {}

func (*Scalar_IntValue) isScalar_Value() {}

func (*Scalar_FloatValue) isScalar_Value() {}

func (*Scalar_BoolValue) isScalar_Value() {}

func (*Scalar_TimestampValue) isScalar_Value() {}

type List struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []*Scalar `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *List) Reset() {
	*x = List{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1alpha1_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *List) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*List) ProtoMessage() {}

func (x *List) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1alpha1_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use List.ProtoReflect.Descriptor instead.
func (*List) Descriptor() ([]byte, []int) {
	return file_core_v1alpha1_types_proto_rawDescGZIP(), []int{1}
}

func (x *List) GetValues() []*Scalar {
	if x != nil {
		return x.Values
	}
	return nil
}

type Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//
	//	*Value_ScalarValue
	//	*Value_ListValue
	Value isValue_Value `protobuf_oneof:"value"`
}

func (x *Value) Reset() {
	*x = Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1alpha1_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1alpha1_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_core_v1alpha1_types_proto_rawDescGZIP(), []int{2}
}

func (m *Value) GetValue() isValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *Value) GetScalarValue() *Scalar {
	if x, ok := x.GetValue().(*Value_ScalarValue); ok {
		return x.ScalarValue
	}
	return nil
}

func (x *Value) GetListValue() *List {
	if x, ok := x.GetValue().(*Value_ListValue); ok {
		return x.ListValue
	}
	return nil
}

type isValue_Value interface {
	isValue_Value()
}

type Value_ScalarValue struct {
	ScalarValue *Scalar `protobuf:"bytes,1,opt,name=scalar_value,json=scalarValue,proto3,oneof"`
}

type Value_ListValue struct {
	ListValue *List `protobuf:"bytes,2,opt,name=list_value,json=listValue,proto3,oneof"`
}

func (*Value_ScalarValue) isValue_Value() {}

func (*Value_ListValue) isValue_Value() {}

type ObjectReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *ObjectReference) Reset() {
	*x = ObjectReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1alpha1_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectReference) ProtoMessage() {}

func (x *ObjectReference) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1alpha1_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectReference.ProtoReflect.Descriptor instead.
func (*ObjectReference) Descriptor() ([]byte, []int) {
	return file_core_v1alpha1_types_proto_rawDescGZIP(), []int{3}
}

func (x *ObjectReference) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ObjectReference) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type KeepPrevious struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Versions uint32               `protobuf:"varint,1,opt,name=versions,proto3" json:"versions,omitempty"`
	Over     *durationpb.Duration `protobuf:"bytes,2,opt,name=over,proto3" json:"over,omitempty"`
}

func (x *KeepPrevious) Reset() {
	*x = KeepPrevious{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1alpha1_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeepPrevious) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeepPrevious) ProtoMessage() {}

func (x *KeepPrevious) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1alpha1_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeepPrevious.ProtoReflect.Descriptor instead.
func (*KeepPrevious) Descriptor() ([]byte, []int) {
	return file_core_v1alpha1_types_proto_rawDescGZIP(), []int{4}
}

func (x *KeepPrevious) GetVersions() uint32 {
	if x != nil {
		return x.Versions
	}
	return 0
}

func (x *KeepPrevious) GetOver() *durationpb.Duration {
	if x != nil {
		return x.Over
	}
	return nil
}

type FeatureDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fqn          string               `protobuf:"bytes,1,opt,name=fqn,proto3" json:"fqn,omitempty"`
	Primitive    Primitive            `protobuf:"varint,2,opt,name=primitive,proto3,enum=core.v1alpha1.Primitive" json:"primitive,omitempty"`
	Aggr         []AggrFn             `protobuf:"varint,3,rep,packed,name=aggr,proto3,enum=core.v1alpha1.AggrFn" json:"aggr,omitempty"`
	Freshness    *durationpb.Duration `protobuf:"bytes,4,opt,name=freshness,proto3" json:"freshness,omitempty"`
	Staleness    *durationpb.Duration `protobuf:"bytes,5,opt,name=staleness,proto3" json:"staleness,omitempty"`
	Timeout      *durationpb.Duration `protobuf:"bytes,6,opt,name=timeout,proto3" json:"timeout,omitempty"`
	KeepPrevious *KeepPrevious        `protobuf:"bytes,7,opt,name=keep_previous,json=keepPrevious,proto3,oneof" json:"keep_previous,omitempty"`
	Keys         []string             `protobuf:"bytes,8,rep,name=keys,proto3" json:"keys,omitempty"`
	Builder      string               `protobuf:"bytes,15,opt,name=builder,proto3" json:"builder,omitempty"`
	DataSource   string               `protobuf:"bytes,16,opt,name=data_source,json=dataSource,proto3" json:"data_source,omitempty"`
	RuntimeEnv   string               `protobuf:"bytes,17,opt,name=runtime_env,json=runtimeEnv,proto3" json:"runtime_env,omitempty"`
}

func (x *FeatureDescriptor) Reset() {
	*x = FeatureDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1alpha1_types_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeatureDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureDescriptor) ProtoMessage() {}

func (x *FeatureDescriptor) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1alpha1_types_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureDescriptor.ProtoReflect.Descriptor instead.
func (*FeatureDescriptor) Descriptor() ([]byte, []int) {
	return file_core_v1alpha1_types_proto_rawDescGZIP(), []int{5}
}

func (x *FeatureDescriptor) GetFqn() string {
	if x != nil {
		return x.Fqn
	}
	return ""
}

func (x *FeatureDescriptor) GetPrimitive() Primitive {
	if x != nil {
		return x.Primitive
	}
	return Primitive_PRIMITIVE_UNSPECIFIED
}

func (x *FeatureDescriptor) GetAggr() []AggrFn {
	if x != nil {
		return x.Aggr
	}
	return nil
}

func (x *FeatureDescriptor) GetFreshness() *durationpb.Duration {
	if x != nil {
		return x.Freshness
	}
	return nil
}

func (x *FeatureDescriptor) GetStaleness() *durationpb.Duration {
	if x != nil {
		return x.Staleness
	}
	return nil
}

func (x *FeatureDescriptor) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *FeatureDescriptor) GetKeepPrevious() *KeepPrevious {
	if x != nil {
		return x.KeepPrevious
	}
	return nil
}

func (x *FeatureDescriptor) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *FeatureDescriptor) GetBuilder() string {
	if x != nil {
		return x.Builder
	}
	return ""
}

func (x *FeatureDescriptor) GetDataSource() string {
	if x != nil {
		return x.DataSource
	}
	return ""
}

func (x *FeatureDescriptor) GetRuntimeEnv() string {
	if x != nil {
		return x.RuntimeEnv
	}
	return ""
}

type FeatureValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fqn       string                 `protobuf:"bytes,1,opt,name=fqn,proto3" json:"fqn,omitempty"`
	Keys      map[string]string      `protobuf:"bytes,2,rep,name=keys,proto3" json:"keys,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Value     *Value                 `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Fresh     bool                   `protobuf:"varint,5,opt,name=fresh,proto3" json:"fresh,omitempty"`
}

func (x *FeatureValue) Reset() {
	*x = FeatureValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1alpha1_types_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeatureValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureValue) ProtoMessage() {}

func (x *FeatureValue) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1alpha1_types_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureValue.ProtoReflect.Descriptor instead.
func (*FeatureValue) Descriptor() ([]byte, []int) {
	return file_core_v1alpha1_types_proto_rawDescGZIP(), []int{6}
}

func (x *FeatureValue) GetFqn() string {
	if x != nil {
		return x.Fqn
	}
	return ""
}

func (x *FeatureValue) GetKeys() map[string]string {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *FeatureValue) GetValue() *Value {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *FeatureValue) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *FeatureValue) GetFresh() bool {
	if x != nil {
		return x.Fresh
	}
	return false
}

var File_core_v1alpha1_types_proto protoreflect.FileDescriptor

var file_core_v1alpha1_types_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe0, 0x01, 0x0a, 0x06, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x12,
	0x23, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x21, 0x0a, 0x0b, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x0a, 0x66, 0x6c, 0x6f, 0x61,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x09, 0x62, 0x6f,
	0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x45, 0x0a, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x0e,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x35, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x2d, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x82,
	0x01, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3a, 0x0a, 0x0c, 0x73, 0x63, 0x61, 0x6c,
	0x61, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53,
	0x63, 0x61, 0x6c, 0x61, 0x72, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x34, 0x0a, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52,
	0x09, 0x6c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x43, 0x0a, 0x0f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x59, 0x0a, 0x0c, 0x4b, 0x65, 0x65, 0x70,
	0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2d, 0x0a, 0x04, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x6f,
	0x76, 0x65, 0x72, 0x22, 0xc7, 0x04, 0x0a, 0x11, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x3e, 0x0a, 0x03, 0x66, 0x71, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2c, 0xfa, 0x42, 0x29, 0x72, 0x27, 0x32, 0x25, 0x28,
	0x69, 0x3f, 0x29, 0x5e, 0x28, 0x5b, 0x61, 0x30, 0x2d, 0x7a, 0x39, 0x5c, 0x2d, 0x5c, 0x2e, 0x5d,
	0x2a, 0x29, 0x28, 0x5c, 0x5b, 0x28, 0x5b, 0x61, 0x30, 0x2d, 0x7a, 0x39, 0x5d, 0x29, 0x2a, 0x5c,
	0x5d, 0x29, 0x3f, 0x24, 0x52, 0x03, 0x66, 0x71, 0x6e, 0x12, 0x40, 0x0a, 0x09, 0x70, 0x72, 0x69,
	0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x69,
	0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01,
	0x52, 0x09, 0x70, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x12, 0x3a, 0x0a, 0x04, 0x61,
	0x67, 0x67, 0x72, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x46, 0x6e,
	0x42, 0x0f, 0xfa, 0x42, 0x0c, 0x92, 0x01, 0x09, 0x18, 0x01, 0x22, 0x05, 0x82, 0x01, 0x02, 0x10,
	0x01, 0x52, 0x04, 0x61, 0x67, 0x67, 0x72, 0x12, 0x37, 0x0a, 0x09, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x6e, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x66, 0x72, 0x65, 0x73, 0x68, 0x6e, 0x65, 0x73, 0x73,
	0x12, 0x37, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09,
	0x73, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x45,
	0x0a, 0x0d, 0x6b, 0x65, 0x65, 0x70, 0x5f, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4b, 0x65, 0x65, 0x70, 0x50, 0x72, 0x65, 0x76, 0x69, 0x6f,
	0x75, 0x73, 0x48, 0x00, 0x52, 0x0c, 0x6b, 0x65, 0x65, 0x70, 0x50, 0x72, 0x65, 0x76, 0x69, 0x6f,
	0x75, 0x73, 0x88, 0x01, 0x01, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x65, 0x72, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x65, 0x6e, 0x76, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x45, 0x6e, 0x76, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x6b, 0x65, 0x65, 0x70, 0x5f, 0x70,
	0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x4a, 0x04, 0x08, 0x09, 0x10, 0x0f, 0x22, 0xbe, 0x02,
	0x0a, 0x0c, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3e,
	0x0a, 0x03, 0x66, 0x71, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2c, 0xfa, 0x42, 0x29,
	0x72, 0x27, 0x32, 0x25, 0x28, 0x69, 0x3f, 0x29, 0x5e, 0x28, 0x5b, 0x61, 0x30, 0x2d, 0x7a, 0x39,
	0x5c, 0x2d, 0x5c, 0x2e, 0x5d, 0x2a, 0x29, 0x28, 0x5c, 0x5b, 0x28, 0x5b, 0x61, 0x30, 0x2d, 0x7a,
	0x39, 0x5d, 0x29, 0x2a, 0x5c, 0x5d, 0x29, 0x3f, 0x24, 0x52, 0x03, 0x66, 0x71, 0x6e, 0x12, 0x39,
	0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x4b, 0x65, 0x79, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x14, 0x0a, 0x05, 0x66, 0x72, 0x65, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x1a, 0x37, 0x0a, 0x09, 0x4b, 0x65, 0x79, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x9d,
	0x02, 0x0a, 0x09, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x12, 0x19, 0x0a, 0x15,
	0x50, 0x52, 0x49, 0x4d, 0x49, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x52, 0x49, 0x4d, 0x49,
	0x54, 0x49, 0x56, 0x45, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x15, 0x0a,
	0x11, 0x50, 0x52, 0x49, 0x4d, 0x49, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x47,
	0x45, 0x52, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x52, 0x49, 0x4d, 0x49, 0x54, 0x49, 0x56,
	0x45, 0x5f, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x52, 0x49,
	0x4d, 0x49, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x42, 0x4f, 0x4f, 0x4c, 0x10, 0x04, 0x12, 0x17, 0x0a,
	0x13, 0x50, 0x52, 0x49, 0x4d, 0x49, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x53,
	0x54, 0x41, 0x4d, 0x50, 0x10, 0x05, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x52, 0x49, 0x4d, 0x49, 0x54,
	0x49, 0x56, 0x45, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x10,
	0x0a, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x52, 0x49, 0x4d, 0x49, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x49,
	0x4e, 0x54, 0x45, 0x47, 0x45, 0x52, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x0b, 0x12, 0x18, 0x0a,
	0x14, 0x50, 0x52, 0x49, 0x4d, 0x49, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x46, 0x4c, 0x4f, 0x41, 0x54,
	0x5f, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x0c, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x52, 0x49, 0x4d, 0x49,
	0x54, 0x49, 0x56, 0x45, 0x5f, 0x42, 0x4f, 0x4f, 0x4c, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x0d,
	0x12, 0x1c, 0x0a, 0x18, 0x50, 0x52, 0x49, 0x4d, 0x49, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x49,
	0x4d, 0x45, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x0e, 0x2a, 0x78,
	0x0a, 0x06, 0x41, 0x67, 0x67, 0x72, 0x46, 0x6e, 0x12, 0x17, 0x0a, 0x13, 0x41, 0x47, 0x47, 0x52,
	0x5f, 0x46, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x47, 0x47, 0x52, 0x5f, 0x46, 0x4e, 0x5f, 0x53, 0x55, 0x4d,
	0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x47, 0x47, 0x52, 0x5f, 0x46, 0x4e, 0x5f, 0x41, 0x56,
	0x47, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x47, 0x47, 0x52, 0x5f, 0x46, 0x4e, 0x5f, 0x4d,
	0x41, 0x58, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x47, 0x47, 0x52, 0x5f, 0x46, 0x4e, 0x5f,
	0x4d, 0x49, 0x4e, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x47, 0x47, 0x52, 0x5f, 0x46, 0x4e,
	0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x05, 0x42, 0xbd, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0a,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x47, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x2d,
	0x6d, 0x6c, 0x2f, 0x72, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x63, 0x6f, 0x72, 0x65, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa, 0x02, 0x0d, 0x43, 0x6f,
	0x72, 0x65, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x0d, 0x43, 0x6f,
	0x72, 0x65, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x19, 0x43, 0x6f,
	0x72, 0x65, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x43, 0x6f, 0x72, 0x65, 0x3a, 0x3a,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_core_v1alpha1_types_proto_rawDescOnce sync.Once
	file_core_v1alpha1_types_proto_rawDescData = file_core_v1alpha1_types_proto_rawDesc
)

func file_core_v1alpha1_types_proto_rawDescGZIP() []byte {
	file_core_v1alpha1_types_proto_rawDescOnce.Do(func() {
		file_core_v1alpha1_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_v1alpha1_types_proto_rawDescData)
	})
	return file_core_v1alpha1_types_proto_rawDescData
}

var file_core_v1alpha1_types_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_core_v1alpha1_types_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_core_v1alpha1_types_proto_goTypes = []interface{}{
	(Primitive)(0),                // 0: core.v1alpha1.Primitive
	(AggrFn)(0),                   // 1: core.v1alpha1.AggrFn
	(*Scalar)(nil),                // 2: core.v1alpha1.Scalar
	(*List)(nil),                  // 3: core.v1alpha1.List
	(*Value)(nil),                 // 4: core.v1alpha1.Value
	(*ObjectReference)(nil),       // 5: core.v1alpha1.ObjectReference
	(*KeepPrevious)(nil),          // 6: core.v1alpha1.KeepPrevious
	(*FeatureDescriptor)(nil),     // 7: core.v1alpha1.FeatureDescriptor
	(*FeatureValue)(nil),          // 8: core.v1alpha1.FeatureValue
	nil,                           // 9: core.v1alpha1.FeatureValue.KeysEntry
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),   // 11: google.protobuf.Duration
}
var file_core_v1alpha1_types_proto_depIdxs = []int32{
	10, // 0: core.v1alpha1.Scalar.timestamp_value:type_name -> google.protobuf.Timestamp
	2,  // 1: core.v1alpha1.List.values:type_name -> core.v1alpha1.Scalar
	2,  // 2: core.v1alpha1.Value.scalar_value:type_name -> core.v1alpha1.Scalar
	3,  // 3: core.v1alpha1.Value.list_value:type_name -> core.v1alpha1.List
	11, // 4: core.v1alpha1.KeepPrevious.over:type_name -> google.protobuf.Duration
	0,  // 5: core.v1alpha1.FeatureDescriptor.primitive:type_name -> core.v1alpha1.Primitive
	1,  // 6: core.v1alpha1.FeatureDescriptor.aggr:type_name -> core.v1alpha1.AggrFn
	11, // 7: core.v1alpha1.FeatureDescriptor.freshness:type_name -> google.protobuf.Duration
	11, // 8: core.v1alpha1.FeatureDescriptor.staleness:type_name -> google.protobuf.Duration
	11, // 9: core.v1alpha1.FeatureDescriptor.timeout:type_name -> google.protobuf.Duration
	6,  // 10: core.v1alpha1.FeatureDescriptor.keep_previous:type_name -> core.v1alpha1.KeepPrevious
	9,  // 11: core.v1alpha1.FeatureValue.keys:type_name -> core.v1alpha1.FeatureValue.KeysEntry
	4,  // 12: core.v1alpha1.FeatureValue.value:type_name -> core.v1alpha1.Value
	10, // 13: core.v1alpha1.FeatureValue.timestamp:type_name -> google.protobuf.Timestamp
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_core_v1alpha1_types_proto_init() }
func file_core_v1alpha1_types_proto_init() {
	if File_core_v1alpha1_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_core_v1alpha1_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scalar); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_core_v1alpha1_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*List); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_core_v1alpha1_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Value); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_core_v1alpha1_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectReference); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_core_v1alpha1_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeepPrevious); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_core_v1alpha1_types_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeatureDescriptor); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_core_v1alpha1_types_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeatureValue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_core_v1alpha1_types_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Scalar_StringValue)(nil),
		(*Scalar_IntValue)(nil),
		(*Scalar_FloatValue)(nil),
		(*Scalar_BoolValue)(nil),
		(*Scalar_TimestampValue)(nil),
	}
	file_core_v1alpha1_types_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Value_ScalarValue)(nil),
		(*Value_ListValue)(nil),
	}
	file_core_v1alpha1_types_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_core_v1alpha1_types_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_core_v1alpha1_types_proto_goTypes,
		DependencyIndexes: file_core_v1alpha1_types_proto_depIdxs,
		EnumInfos:         file_core_v1alpha1_types_proto_enumTypes,
		MessageInfos:      file_core_v1alpha1_types_proto_msgTypes,
	}.Build()
	File_core_v1alpha1_types_proto = out.File
	file_core_v1alpha1_types_proto_rawDesc = nil
	file_core_v1alpha1_types_proto_goTypes = nil
	file_core_v1alpha1_types_proto_depIdxs = nil
}
