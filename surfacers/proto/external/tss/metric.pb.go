// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.2
// source: github.com/cloudprober/cloudprober/surfacers/proto/external/tss/metric.proto

package tss

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MetricLabel_Scope int32

const (
	MetricLabel_INVALID  MetricLabel_Scope = 0
	MetricLabel_ORG      MetricLabel_Scope = 1
	MetricLabel_RESOURCE MetricLabel_Scope = 2
)

// Enum value maps for MetricLabel_Scope.
var (
	MetricLabel_Scope_name = map[int32]string{
		0: "INVALID",
		1: "ORG",
		2: "RESOURCE",
	}
	MetricLabel_Scope_value = map[string]int32{
		"INVALID":  0,
		"ORG":      1,
		"RESOURCE": 2,
	}
)

func (x MetricLabel_Scope) Enum() *MetricLabel_Scope {
	p := new(MetricLabel_Scope)
	*p = x
	return p
}

func (x MetricLabel_Scope) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MetricLabel_Scope) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto_enumTypes[0].Descriptor()
}

func (MetricLabel_Scope) Type() protoreflect.EnumType {
	return &file_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto_enumTypes[0]
}

func (x MetricLabel_Scope) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MetricLabel_Scope.Descriptor instead.
func (MetricLabel_Scope) EnumDescriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto_rawDescGZIP(), []int{0, 0}
}

type ResourceMetric_Type int32

const (
	ResourceMetric_INVALID ResourceMetric_Type = 0
	ResourceMetric_HTTP    ResourceMetric_Type = 1
)

// Enum value maps for ResourceMetric_Type.
var (
	ResourceMetric_Type_name = map[int32]string{
		0: "INVALID",
		1: "HTTP",
	}
	ResourceMetric_Type_value = map[string]int32{
		"INVALID": 0,
		"HTTP":    1,
	}
)

func (x ResourceMetric_Type) Enum() *ResourceMetric_Type {
	p := new(ResourceMetric_Type)
	*p = x
	return p
}

func (x ResourceMetric_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResourceMetric_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto_enumTypes[1].Descriptor()
}

func (ResourceMetric_Type) Type() protoreflect.EnumType {
	return &file_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto_enumTypes[1]
}

func (x ResourceMetric_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResourceMetric_Type.Descriptor instead.
func (ResourceMetric_Type) EnumDescriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto_rawDescGZIP(), []int{1, 0}
}

type ResourceMetric_Kind int32

const (
	ResourceMetric_UNDEFINED ResourceMetric_Kind = 0
	ResourceMetric_COUNTER   ResourceMetric_Kind = 1
	ResourceMetric_GUAGE     ResourceMetric_Kind = 2
)

// Enum value maps for ResourceMetric_Kind.
var (
	ResourceMetric_Kind_name = map[int32]string{
		0: "UNDEFINED",
		1: "COUNTER",
		2: "GUAGE",
	}
	ResourceMetric_Kind_value = map[string]int32{
		"UNDEFINED": 0,
		"COUNTER":   1,
		"GUAGE":     2,
	}
)

func (x ResourceMetric_Kind) Enum() *ResourceMetric_Kind {
	p := new(ResourceMetric_Kind)
	*p = x
	return p
}

func (x ResourceMetric_Kind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResourceMetric_Kind) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto_enumTypes[2].Descriptor()
}

func (ResourceMetric_Kind) Type() protoreflect.EnumType {
	return &file_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto_enumTypes[2]
}

func (x ResourceMetric_Kind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResourceMetric_Kind.Descriptor instead.
func (ResourceMetric_Kind) EnumDescriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto_rawDescGZIP(), []int{1, 1}
}

type MetricLabel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string            `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string            `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Scope MetricLabel_Scope `protobuf:"varint,3,opt,name=scope,proto3,enum=cloudprober.surfacer.external.tss.MetricLabel_Scope" json:"scope,omitempty"`
}

func (x *MetricLabel) Reset() {
	*x = MetricLabel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricLabel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricLabel) ProtoMessage() {}

func (x *MetricLabel) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricLabel.ProtoReflect.Descriptor instead.
func (*MetricLabel) Descriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto_rawDescGZIP(), []int{0}
}

func (x *MetricLabel) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *MetricLabel) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *MetricLabel) GetScope() MetricLabel_Scope {
	if x != nil {
		return x.Scope
	}
	return MetricLabel_INVALID
}

type ResourceMetric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type ResourceMetric_Type `protobuf:"varint,1,opt,name=type,proto3,enum=cloudprober.surfacer.external.tss.ResourceMetric_Type" json:"type,omitempty"`
	Kind ResourceMetric_Kind `protobuf:"varint,2,opt,name=kind,proto3,enum=cloudprober.surfacer.external.tss.ResourceMetric_Kind" json:"kind,omitempty"`
	Foo  string              `protobuf:"bytes,10,opt,name=foo,proto3" json:"foo,omitempty"`
}

func (x *ResourceMetric) Reset() {
	*x = ResourceMetric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceMetric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceMetric) ProtoMessage() {}

func (x *ResourceMetric) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceMetric.ProtoReflect.Descriptor instead.
func (*ResourceMetric) Descriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto_rawDescGZIP(), []int{1}
}

func (x *ResourceMetric) GetType() ResourceMetric_Type {
	if x != nil {
		return x.Type
	}
	return ResourceMetric_INVALID
}

func (x *ResourceMetric) GetKind() ResourceMetric_Kind {
	if x != nil {
		return x.Kind
	}
	return ResourceMetric_UNDEFINED
}

func (x *ResourceMetric) GetFoo() string {
	if x != nil {
		return x.Foo
	}
	return ""
}

type ResourceMetricList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metrics []*ResourceMetric `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
}

func (x *ResourceMetricList) Reset() {
	*x = ResourceMetricList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceMetricList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceMetricList) ProtoMessage() {}

func (x *ResourceMetricList) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceMetricList.ProtoReflect.Descriptor instead.
func (*ResourceMetricList) Descriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto_rawDescGZIP(), []int{2}
}

func (x *ResourceMetricList) GetMetrics() []*ResourceMetric {
	if x != nil {
		return x.Metrics
	}
	return nil
}

var File_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto protoreflect.FileDescriptor

var file_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto_rawDesc = []byte{
	0x0a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x62, 0x65, 0x72, 0x2f, 0x73, 0x75, 0x72, 0x66, 0x61, 0x63, 0x65, 0x72, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x73,
	0x73, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x73, 0x75, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x72, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x74, 0x73,
	0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xae, 0x01, 0x0a, 0x0b, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x4a, 0x0a, 0x05, 0x73, 0x63,
	0x6f, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x73, 0x75, 0x72, 0x66, 0x61, 0x63, 0x65, 0x72,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x74, 0x73, 0x73, 0x2e, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52,
	0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x22, 0x2b, 0x0a, 0x05, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03,
	0x4f, 0x52, 0x47, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43,
	0x45, 0x10, 0x02, 0x22, 0x88, 0x02, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x4a, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x36, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62,
	0x65, 0x72, 0x2e, 0x73, 0x75, 0x72, 0x66, 0x61, 0x63, 0x65, 0x72, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x74, 0x73, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x4a, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x36, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x73,
	0x75, 0x72, 0x66, 0x61, 0x63, 0x65, 0x72, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2e, 0x74, 0x73, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x2e, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x66, 0x6f, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x66, 0x6f, 0x6f,
	0x22, 0x1d, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x54, 0x54, 0x50, 0x10, 0x01, 0x22,
	0x2d, 0x0a, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x44, 0x45, 0x46,
	0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x45,
	0x52, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x55, 0x41, 0x47, 0x45, 0x10, 0x02, 0x22, 0x61,
	0x0a, 0x12, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x4b, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f,
	0x62, 0x65, 0x72, 0x2e, 0x73, 0x75, 0x72, 0x66, 0x61, 0x63, 0x65, 0x72, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x74, 0x73, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x73, 0x75, 0x72, 0x66, 0x61, 0x63, 0x65, 0x72,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x74, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto_rawDescOnce sync.Once
	file_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto_rawDescData = file_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto_rawDesc
)

func file_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto_rawDescGZIP() []byte {
	file_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto_rawDescOnce.Do(func() {
		file_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto_rawDescData)
	})
	return file_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto_rawDescData
}

var file_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto_goTypes = []any{
	(MetricLabel_Scope)(0),     // 0: cloudprober.surfacer.external.tss.MetricLabel.Scope
	(ResourceMetric_Type)(0),   // 1: cloudprober.surfacer.external.tss.ResourceMetric.Type
	(ResourceMetric_Kind)(0),   // 2: cloudprober.surfacer.external.tss.ResourceMetric.Kind
	(*MetricLabel)(nil),        // 3: cloudprober.surfacer.external.tss.MetricLabel
	(*ResourceMetric)(nil),     // 4: cloudprober.surfacer.external.tss.ResourceMetric
	(*ResourceMetricList)(nil), // 5: cloudprober.surfacer.external.tss.ResourceMetricList
}
var file_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto_depIdxs = []int32{
	0, // 0: cloudprober.surfacer.external.tss.MetricLabel.scope:type_name -> cloudprober.surfacer.external.tss.MetricLabel.Scope
	1, // 1: cloudprober.surfacer.external.tss.ResourceMetric.type:type_name -> cloudprober.surfacer.external.tss.ResourceMetric.Type
	2, // 2: cloudprober.surfacer.external.tss.ResourceMetric.kind:type_name -> cloudprober.surfacer.external.tss.ResourceMetric.Kind
	4, // 3: cloudprober.surfacer.external.tss.ResourceMetricList.metrics:type_name -> cloudprober.surfacer.external.tss.ResourceMetric
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto_init() }
func file_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto_init() {
	if File_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*MetricLabel); i {
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
		file_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ResourceMetric); i {
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
		file_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ResourceMetricList); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto_goTypes,
		DependencyIndexes: file_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto_depIdxs,
		EnumInfos:         file_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto_enumTypes,
		MessageInfos:      file_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto_msgTypes,
	}.Build()
	File_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto = out.File
	file_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto_rawDesc = nil
	file_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto_goTypes = nil
	file_github_com_cloudprober_cloudprober_surfacers_proto_external_tss_metric_proto_depIdxs = nil
}
