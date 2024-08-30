// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0
// source: api/server.proto

package serverpb

import (
	_ "github.com/asjard/asjard/pkg/protobuf/httppb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HelloReq_Kind int32

const (
	HelloReq_K_A HelloReq_Kind = 0
	HelloReq_K_B HelloReq_Kind = 1
)

// Enum value maps for HelloReq_Kind.
var (
	HelloReq_Kind_name = map[int32]string{
		0: "K_A",
		1: "K_B",
	}
	HelloReq_Kind_value = map[string]int32{
		"K_A": 0,
		"K_B": 1,
	}
)

func (x HelloReq_Kind) Enum() *HelloReq_Kind {
	p := new(HelloReq_Kind)
	*p = x
	return p
}

func (x HelloReq_Kind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HelloReq_Kind) Descriptor() protoreflect.EnumDescriptor {
	return file_api_server_proto_enumTypes[0].Descriptor()
}

func (HelloReq_Kind) Type() protoreflect.EnumType {
	return &file_api_server_proto_enumTypes[0]
}

func (x HelloReq_Kind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HelloReq_Kind.Descriptor instead.
func (HelloReq_Kind) EnumDescriptor() ([]byte, []int) {
	return file_api_server_proto_rawDescGZIP(), []int{0, 0}
}

type HelloReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 区域ID
	RegionId string `protobuf:"bytes,1,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	// 项目ID
	ProjectId string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// 用户ID
	UserId int64 `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// 字符串列表
	StrList []string `protobuf:"bytes,4,rep,name=str_list,json=strList,proto3" json:"str_list,omitempty"`
	// 数字列表
	IntList []int64 `protobuf:"varint,5,rep,packed,name=int_list,json=intList,proto3" json:"int_list,omitempty"`
	// 对象
	Obj *HelloReq_Obj `protobuf:"bytes,6,opt,name=obj,proto3" json:"obj,omitempty"`
	// 对象列表
	Objs []*HelloReq_Obj `protobuf:"bytes,7,rep,name=objs,proto3" json:"objs,omitempty"`
	// 配置
	Configs *HelloReq_Configs `protobuf:"bytes,8,opt,name=configs,proto3" json:"configs,omitempty"`
	// 分页
	Page int32 `protobuf:"varint,9,opt,name=page,proto3" json:"page,omitempty"`
	// 每页大小
	Size int32 `protobuf:"varint,10,opt,name=size,proto3" json:"size,omitempty"`
	// 排序
	Sort string `protobuf:"bytes,11,opt,name=sort,proto3" json:"sort,omitempty"`
	// 布尔类型
	Ok *bool `protobuf:"varint,12,opt,name=ok,proto3,oneof" json:"ok,omitempty"`
	// 可选整形参数
	IntOptionalValue *int32 `protobuf:"varint,13,opt,name=int_optional_value,json=intOptionalValue,proto3,oneof" json:"int_optional_value,omitempty"`
	// 可选字符串参数
	StringOptionalValue *string `protobuf:"bytes,14,opt,name=string_optional_value,json=stringOptionalValue,proto3,oneof" json:"string_optional_value,omitempty"`
	// 可选枚举参数
	Kind *HelloReq_Kind `protobuf:"varint,15,opt,name=kind,proto3,enum=api.v1.server.HelloReq_Kind,oneof" json:"kind,omitempty"`
	// 枚举列表
	Kinds      []HelloReq_Kind `protobuf:"varint,16,rep,packed,name=kinds,proto3,enum=api.v1.server.HelloReq_Kind" json:"kinds,omitempty"`
	BytesValue []byte          `protobuf:"bytes,17,opt,name=bytes_value,json=bytesValue,proto3" json:"bytes_value,omitempty"`
	// openapi 会把这个字段解析为字符串
	Uint64Value      uint64                 `protobuf:"varint,18,opt,name=uint64_value,json=uint64Value,proto3" json:"uint64_value,omitempty"`
	GoogleInt64Value *wrapperspb.Int64Value `protobuf:"bytes,19,opt,name=google_int64_value,json=googleInt64Value,proto3" json:"google_int64_value,omitempty"`
	App              string                 `protobuf:"bytes,20,opt,name=app,proto3" json:"app,omitempty"`
	Region           string                 `protobuf:"bytes,21,opt,name=region,proto3" json:"region,omitempty"`
	Az               string                 `protobuf:"bytes,22,opt,name=az,proto3" json:"az,omitempty"`
	Instance         *HelloReq_Instance     `protobuf:"bytes,23,opt,name=instance,proto3" json:"instance,omitempty"`
}

func (x *HelloReq) Reset() {
	*x = HelloReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReq) ProtoMessage() {}

func (x *HelloReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReq.ProtoReflect.Descriptor instead.
func (*HelloReq) Descriptor() ([]byte, []int) {
	return file_api_server_proto_rawDescGZIP(), []int{0}
}

func (x *HelloReq) GetRegionId() string {
	if x != nil {
		return x.RegionId
	}
	return ""
}

func (x *HelloReq) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *HelloReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *HelloReq) GetStrList() []string {
	if x != nil {
		return x.StrList
	}
	return nil
}

func (x *HelloReq) GetIntList() []int64 {
	if x != nil {
		return x.IntList
	}
	return nil
}

func (x *HelloReq) GetObj() *HelloReq_Obj {
	if x != nil {
		return x.Obj
	}
	return nil
}

func (x *HelloReq) GetObjs() []*HelloReq_Obj {
	if x != nil {
		return x.Objs
	}
	return nil
}

func (x *HelloReq) GetConfigs() *HelloReq_Configs {
	if x != nil {
		return x.Configs
	}
	return nil
}

func (x *HelloReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *HelloReq) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *HelloReq) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

func (x *HelloReq) GetOk() bool {
	if x != nil && x.Ok != nil {
		return *x.Ok
	}
	return false
}

func (x *HelloReq) GetIntOptionalValue() int32 {
	if x != nil && x.IntOptionalValue != nil {
		return *x.IntOptionalValue
	}
	return 0
}

func (x *HelloReq) GetStringOptionalValue() string {
	if x != nil && x.StringOptionalValue != nil {
		return *x.StringOptionalValue
	}
	return ""
}

func (x *HelloReq) GetKind() HelloReq_Kind {
	if x != nil && x.Kind != nil {
		return *x.Kind
	}
	return HelloReq_K_A
}

func (x *HelloReq) GetKinds() []HelloReq_Kind {
	if x != nil {
		return x.Kinds
	}
	return nil
}

func (x *HelloReq) GetBytesValue() []byte {
	if x != nil {
		return x.BytesValue
	}
	return nil
}

func (x *HelloReq) GetUint64Value() uint64 {
	if x != nil {
		return x.Uint64Value
	}
	return 0
}

func (x *HelloReq) GetGoogleInt64Value() *wrapperspb.Int64Value {
	if x != nil {
		return x.GoogleInt64Value
	}
	return nil
}

func (x *HelloReq) GetApp() string {
	if x != nil {
		return x.App
	}
	return ""
}

func (x *HelloReq) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *HelloReq) GetAz() string {
	if x != nil {
		return x.Az
	}
	return ""
}

func (x *HelloReq) GetInstance() *HelloReq_Instance {
	if x != nil {
		return x.Instance
	}
	return nil
}

type HelloReq_Obj struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FieldInt int32  `protobuf:"varint,1,opt,name=field_int,json=fieldInt,proto3" json:"field_int,omitempty"`
	FieldStr string `protobuf:"bytes,2,opt,name=field_str,json=fieldStr,proto3" json:"field_str,omitempty"`
}

func (x *HelloReq_Obj) Reset() {
	*x = HelloReq_Obj{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReq_Obj) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReq_Obj) ProtoMessage() {}

func (x *HelloReq_Obj) ProtoReflect() protoreflect.Message {
	mi := &file_api_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReq_Obj.ProtoReflect.Descriptor instead.
func (*HelloReq_Obj) Descriptor() ([]byte, []int) {
	return file_api_server_proto_rawDescGZIP(), []int{0, 0}
}

func (x *HelloReq_Obj) GetFieldInt() int32 {
	if x != nil {
		return x.FieldInt
	}
	return 0
}

func (x *HelloReq_Obj) GetFieldStr() string {
	if x != nil {
		return x.FieldStr
	}
	return ""
}

type HelloReq_Configs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timeout                                     string `protobuf:"bytes,1,opt,name=timeout,proto3" json:"timeout,omitempty"`
	FieldInDifferentFileUnderSameSection        string `protobuf:"bytes,2,opt,name=field_in_different_file_under_same_section,json=fieldInDifferentFileUnderSameSection,proto3" json:"field_in_different_file_under_same_section,omitempty"`
	AnotherFieldInDifferentFileUnderSameSection string `protobuf:"bytes,3,opt,name=another_field_in_different_file_under_same_section,json=anotherFieldInDifferentFileUnderSameSection,proto3" json:"another_field_in_different_file_under_same_section,omitempty"`
	KeyInDifferentSourcer                       string `protobuf:"bytes,4,opt,name=key_in_different_sourcer,json=keyInDifferentSourcer,proto3" json:"key_in_different_sourcer,omitempty"`
}

func (x *HelloReq_Configs) Reset() {
	*x = HelloReq_Configs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReq_Configs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReq_Configs) ProtoMessage() {}

func (x *HelloReq_Configs) ProtoReflect() protoreflect.Message {
	mi := &file_api_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReq_Configs.ProtoReflect.Descriptor instead.
func (*HelloReq_Configs) Descriptor() ([]byte, []int) {
	return file_api_server_proto_rawDescGZIP(), []int{0, 1}
}

func (x *HelloReq_Configs) GetTimeout() string {
	if x != nil {
		return x.Timeout
	}
	return ""
}

func (x *HelloReq_Configs) GetFieldInDifferentFileUnderSameSection() string {
	if x != nil {
		return x.FieldInDifferentFileUnderSameSection
	}
	return ""
}

func (x *HelloReq_Configs) GetAnotherFieldInDifferentFileUnderSameSection() string {
	if x != nil {
		return x.AnotherFieldInDifferentFileUnderSameSection
	}
	return ""
}

func (x *HelloReq_Configs) GetKeyInDifferentSourcer() string {
	if x != nil {
		return x.KeyInDifferentSourcer
	}
	return ""
}

type HelloReq_Instance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	SystemCode string            `protobuf:"bytes,3,opt,name=system_code,json=systemCode,proto3" json:"system_code,omitempty"`
	Version    string            `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	Metadata   map[string]string `protobuf:"bytes,5,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *HelloReq_Instance) Reset() {
	*x = HelloReq_Instance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReq_Instance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReq_Instance) ProtoMessage() {}

func (x *HelloReq_Instance) ProtoReflect() protoreflect.Message {
	mi := &file_api_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReq_Instance.ProtoReflect.Descriptor instead.
func (*HelloReq_Instance) Descriptor() ([]byte, []int) {
	return file_api_server_proto_rawDescGZIP(), []int{0, 2}
}

func (x *HelloReq_Instance) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *HelloReq_Instance) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HelloReq_Instance) GetSystemCode() string {
	if x != nil {
		return x.SystemCode
	}
	return ""
}

func (x *HelloReq_Instance) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *HelloReq_Instance) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_api_server_proto protoreflect.FileDescriptor

var file_api_server_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x1a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73,
	0x6a, 0x61, 0x72, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x68, 0x74,
	0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x0c, 0x0a, 0x08, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52,
	0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x72, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x03, 0x52, 0x07, 0x69, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2d, 0x0a,
	0x03, 0x6f, 0x62, 0x6a, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x71, 0x2e, 0x4f, 0x62, 0x6a, 0x52, 0x03, 0x6f, 0x62, 0x6a, 0x12, 0x2f, 0x0a, 0x04,
	0x6f, 0x62, 0x6a, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x71, 0x2e, 0x4f, 0x62, 0x6a, 0x52, 0x04, 0x6f, 0x62, 0x6a, 0x73, 0x12, 0x39, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x73, 0x6f, 0x72, 0x74, 0x12, 0x13, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08,
	0x48, 0x00, 0x52, 0x02, 0x6f, 0x6b, 0x88, 0x01, 0x01, 0x12, 0x31, 0x0a, 0x12, 0x69, 0x6e, 0x74,
	0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x10, 0x69, 0x6e, 0x74, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x88, 0x01, 0x01, 0x12, 0x37, 0x0a, 0x15,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x13, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x88, 0x01, 0x01, 0x12, 0x35, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x2e, 0x4b, 0x69, 0x6e,
	0x64, 0x48, 0x03, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x88, 0x01, 0x01, 0x12, 0x32, 0x0a, 0x05,
	0x6b, 0x69, 0x6e, 0x64, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x52, 0x65, 0x71, 0x2e, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x05, 0x6b, 0x69, 0x6e, 0x64, 0x73,
	0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x11, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x62, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x49, 0x0a, 0x12, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x69,
	0x6e, 0x74, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x10, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x61, 0x70, 0x70, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70,
	0x70, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x15, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x61, 0x7a, 0x18,
	0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x61, 0x7a, 0x12, 0x3c, 0x0a, 0x08, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x52, 0x65, 0x71, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x08, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x1a, 0x3f, 0x0a, 0x03, 0x4f, 0x62, 0x6a, 0x12, 0x1b,
	0x0a, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x49, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x73, 0x74, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x53, 0x74, 0x72, 0x1a, 0x9f, 0x02, 0x0a, 0x07, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x58,
	0x0a, 0x2a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x69, 0x6e, 0x5f, 0x64, 0x69, 0x66, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x5f,
	0x73, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x24, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x49, 0x6e, 0x44, 0x69, 0x66, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x6e, 0x64, 0x65, 0x72, 0x53, 0x61, 0x6d,
	0x65, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x67, 0x0a, 0x32, 0x61, 0x6e, 0x6f, 0x74,
	0x68, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x69, 0x6e, 0x5f, 0x64, 0x69, 0x66,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x75, 0x6e, 0x64, 0x65,
	0x72, 0x5f, 0x73, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x2b, 0x61, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x49, 0x6e, 0x44, 0x69, 0x66, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x6c,
	0x65, 0x55, 0x6e, 0x64, 0x65, 0x72, 0x53, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x37, 0x0a, 0x18, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x6e, 0x5f, 0x64, 0x69, 0x66, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x15, 0x6b, 0x65, 0x79, 0x49, 0x6e, 0x44, 0x69, 0x66, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x72, 0x1a, 0xf2, 0x01, 0x0a, 0x08, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x4a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65,
	0x71, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x18, 0x0a, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x07, 0x0a, 0x03, 0x4b, 0x5f, 0x41, 0x10, 0x00,
	0x12, 0x07, 0x0a, 0x03, 0x4b, 0x5f, 0x42, 0x10, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x6f, 0x6b,
	0x42, 0x15, 0x0a, 0x13, 0x5f, 0x69, 0x6e, 0x74, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x18, 0x0a, 0x16, 0x5f, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x32, 0x8e, 0x04, 0x0a, 0x06, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x8f, 0x02, 0x0a, 0x03, 0x53, 0x61, 0x79, 0x12, 0x17, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x22,
	0xd5, 0x01, 0x82, 0xb5, 0x18, 0x47, 0x6a, 0x0c, 0xe6, 0x8e, 0xa5, 0xe5, 0x8f, 0xa3, 0xe6, 0x8f,
	0x8f, 0xe8, 0xbf, 0xb0, 0x22, 0x37, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x82, 0xb5, 0x18,
	0x39, 0x12, 0x37, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x7b,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x82, 0xb5, 0x18, 0x49, 0x62, 0x0e,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x5f, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x2a, 0x37,
	0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x7d, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x7b, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x5b, 0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65,
	0x71, 0x22, 0x21, 0x82, 0xb5, 0x18, 0x08, 0x12, 0x06, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x82,
	0xb5, 0x18, 0x11, 0x4a, 0x01, 0x2f, 0x52, 0x01, 0x2f, 0x5a, 0x01, 0x2f, 0x12, 0x06, 0x2f, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x41, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x0a, 0x82, 0xb5, 0x18,
	0x06, 0x12, 0x04, 0x2f, 0x6c, 0x6f, 0x67, 0x12, 0x3a, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x6c, 0x12,
	0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65,
	0x71, 0x22, 0x00, 0x1a, 0x16, 0x82, 0xa6, 0x1d, 0x12, 0x5a, 0x10, 0x2f, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x42, 0x32, 0x5a, 0x30, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x6a, 0x61, 0x72, 0x64,
	0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_server_proto_rawDescOnce sync.Once
	file_api_server_proto_rawDescData = file_api_server_proto_rawDesc
)

func file_api_server_proto_rawDescGZIP() []byte {
	file_api_server_proto_rawDescOnce.Do(func() {
		file_api_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_server_proto_rawDescData)
	})
	return file_api_server_proto_rawDescData
}

var file_api_server_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_server_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_server_proto_goTypes = []interface{}{
	(HelloReq_Kind)(0),            // 0: api.v1.server.HelloReq.Kind
	(*HelloReq)(nil),              // 1: api.v1.server.HelloReq
	(*HelloReq_Obj)(nil),          // 2: api.v1.server.HelloReq.Obj
	(*HelloReq_Configs)(nil),      // 3: api.v1.server.HelloReq.Configs
	(*HelloReq_Instance)(nil),     // 4: api.v1.server.HelloReq.Instance
	nil,                           // 5: api.v1.server.HelloReq.Instance.MetadataEntry
	(*wrapperspb.Int64Value)(nil), // 6: google.protobuf.Int64Value
	(*emptypb.Empty)(nil),         // 7: google.protobuf.Empty
}
var file_api_server_proto_depIdxs = []int32{
	2,  // 0: api.v1.server.HelloReq.obj:type_name -> api.v1.server.HelloReq.Obj
	2,  // 1: api.v1.server.HelloReq.objs:type_name -> api.v1.server.HelloReq.Obj
	3,  // 2: api.v1.server.HelloReq.configs:type_name -> api.v1.server.HelloReq.Configs
	0,  // 3: api.v1.server.HelloReq.kind:type_name -> api.v1.server.HelloReq.Kind
	0,  // 4: api.v1.server.HelloReq.kinds:type_name -> api.v1.server.HelloReq.Kind
	6,  // 5: api.v1.server.HelloReq.google_int64_value:type_name -> google.protobuf.Int64Value
	4,  // 6: api.v1.server.HelloReq.instance:type_name -> api.v1.server.HelloReq.Instance
	5,  // 7: api.v1.server.HelloReq.Instance.metadata:type_name -> api.v1.server.HelloReq.Instance.MetadataEntry
	1,  // 8: api.v1.server.Server.Say:input_type -> api.v1.server.HelloReq
	7,  // 9: api.v1.server.Server.Hello:input_type -> google.protobuf.Empty
	7,  // 10: api.v1.server.Server.Log:input_type -> google.protobuf.Empty
	1,  // 11: api.v1.server.Server.Call:input_type -> api.v1.server.HelloReq
	1,  // 12: api.v1.server.Server.Say:output_type -> api.v1.server.HelloReq
	1,  // 13: api.v1.server.Server.Hello:output_type -> api.v1.server.HelloReq
	7,  // 14: api.v1.server.Server.Log:output_type -> google.protobuf.Empty
	1,  // 15: api.v1.server.Server.Call:output_type -> api.v1.server.HelloReq
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_api_server_proto_init() }
func file_api_server_proto_init() {
	if File_api_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReq); i {
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
		file_api_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReq_Obj); i {
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
		file_api_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReq_Configs); i {
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
		file_api_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReq_Instance); i {
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
	file_api_server_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_server_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_server_proto_goTypes,
		DependencyIndexes: file_api_server_proto_depIdxs,
		EnumInfos:         file_api_server_proto_enumTypes,
		MessageInfos:      file_api_server_proto_msgTypes,
	}.Build()
	File_api_server_proto = out.File
	file_api_server_proto_rawDesc = nil
	file_api_server_proto_goTypes = nil
	file_api_server_proto_depIdxs = nil
}
