// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: common.proto

package common

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AttributeId_AttributeKind int32

const (
	AttributeId_INT64   AttributeId_AttributeKind = 0
	AttributeId_FLOAT64 AttributeId_AttributeKind = 1
	AttributeId_STRING  AttributeId_AttributeKind = 2
	AttributeId_BOOL    AttributeId_AttributeKind = 3
)

// Enum value maps for AttributeId_AttributeKind.
var (
	AttributeId_AttributeKind_name = map[int32]string{
		0: "INT64",
		1: "FLOAT64",
		2: "STRING",
		3: "BOOL",
	}
	AttributeId_AttributeKind_value = map[string]int32{
		"INT64":   0,
		"FLOAT64": 1,
		"STRING":  2,
		"BOOL":    3,
	}
)

func (x AttributeId_AttributeKind) Enum() *AttributeId_AttributeKind {
	p := new(AttributeId_AttributeKind)
	*p = x
	return p
}

func (x AttributeId_AttributeKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AttributeId_AttributeKind) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[0].Descriptor()
}

func (AttributeId_AttributeKind) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[0]
}

func (x AttributeId_AttributeKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AttributeId_AttributeKind.Descriptor instead.
func (AttributeId_AttributeKind) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0, 0}
}

type Permission_PermissionKind int32

const (
	Permission_ALLOW Permission_PermissionKind = 0
	Permission_DENY  Permission_PermissionKind = 1
)

// Enum value maps for Permission_PermissionKind.
var (
	Permission_PermissionKind_name = map[int32]string{
		0: "ALLOW",
		1: "DENY",
	}
	Permission_PermissionKind_value = map[string]int32{
		"ALLOW": 0,
		"DENY":  1,
	}
)

func (x Permission_PermissionKind) Enum() *Permission_PermissionKind {
	p := new(Permission_PermissionKind)
	*p = x
	return p
}

func (x Permission_PermissionKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Permission_PermissionKind) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[1].Descriptor()
}

func (Permission_PermissionKind) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[1]
}

func (x Permission_PermissionKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Permission_PermissionKind.Descriptor instead.
func (Permission_PermissionKind) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{7, 0}
}

type AttributeId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string                    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Kind AttributeId_AttributeKind `protobuf:"varint,2,opt,name=kind,proto3,enum=common.AttributeId_AttributeKind" json:"kind,omitempty"`
}

func (x *AttributeId) Reset() {
	*x = AttributeId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttributeId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttributeId) ProtoMessage() {}

func (x *AttributeId) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttributeId.ProtoReflect.Descriptor instead.
func (*AttributeId) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

func (x *AttributeId) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AttributeId) GetKind() AttributeId_AttributeKind {
	if x != nil {
		return x.Kind
	}
	return AttributeId_INT64
}

type Attribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    *AttributeId `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Value []byte       `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Attribute) Reset() {
	*x = Attribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attribute) ProtoMessage() {}

func (x *Attribute) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attribute.ProtoReflect.Descriptor instead.
func (*Attribute) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

func (x *Attribute) GetId() *AttributeId {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Attribute) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type Int64Attribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Int64Attribute) Reset() {
	*x = Int64Attribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Int64Attribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Int64Attribute) ProtoMessage() {}

func (x *Int64Attribute) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Int64Attribute.ProtoReflect.Descriptor instead.
func (*Int64Attribute) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2}
}

func (x *Int64Attribute) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type Float64Attribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value float64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Float64Attribute) Reset() {
	*x = Float64Attribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Float64Attribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Float64Attribute) ProtoMessage() {}

func (x *Float64Attribute) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Float64Attribute.ProtoReflect.Descriptor instead.
func (*Float64Attribute) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{3}
}

func (x *Float64Attribute) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type StringAttribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *StringAttribute) Reset() {
	*x = StringAttribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringAttribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringAttribute) ProtoMessage() {}

func (x *StringAttribute) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringAttribute.ProtoReflect.Descriptor instead.
func (*StringAttribute) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{4}
}

func (x *StringAttribute) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type BoolAttribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value bool `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *BoolAttribute) Reset() {
	*x = BoolAttribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoolAttribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoolAttribute) ProtoMessage() {}

func (x *BoolAttribute) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoolAttribute.ProtoReflect.Descriptor instead.
func (*BoolAttribute) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{5}
}

func (x *BoolAttribute) GetValue() bool {
	if x != nil {
		return x.Value
	}
	return false
}

type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{6}
}

func (x *Resource) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Resource) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

type Permission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string                    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Kind      Permission_PermissionKind `protobuf:"varint,2,opt,name=kind,proto3,enum=common.Permission_PermissionKind" json:"kind,omitempty"`
	Condition *Condition                `protobuf:"bytes,3,opt,name=condition,proto3" json:"condition,omitempty"`
}

func (x *Permission) Reset() {
	*x = Permission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Permission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Permission) ProtoMessage() {}

func (x *Permission) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Permission.ProtoReflect.Descriptor instead.
func (*Permission) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{7}
}

func (x *Permission) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Permission) GetKind() Permission_PermissionKind {
	if x != nil {
		return x.Kind
	}
	return Permission_ALLOW
}

func (x *Permission) GetCondition() *Condition {
	if x != nil {
		return x.Condition
	}
	return nil
}

type Condition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Expression string `protobuf:"bytes,1,opt,name=expression,proto3" json:"expression,omitempty"`
}

func (x *Condition) Reset() {
	*x = Condition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Condition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Condition) ProtoMessage() {}

func (x *Condition) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Condition.ProtoReflect.Descriptor instead.
func (*Condition) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{8}
}

func (x *Condition) GetExpression() string {
	if x != nil {
		return x.Expression
	}
	return ""
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x97, 0x01, 0x0a, 0x0b, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x04, 0x6b, 0x69,
	0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x49, 0x64, 0x2e, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e,
	0x64, 0x22, 0x3d, 0x0a, 0x0d, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x4b, 0x69,
	0x6e, 0x64, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x36, 0x34, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54,
	0x52, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x4f, 0x4f, 0x4c, 0x10, 0x03,
	0x22, 0x46, 0x0a, 0x09, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x23, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x49, 0x64, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x26, 0x0a, 0x0e, 0x49, 0x6e, 0x74, 0x36,
	0x34, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x28, 0x0a, 0x10, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x36, 0x34, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x27, 0x0a, 0x0f, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x25, 0x0a, 0x0d, 0x42, 0x6f, 0x6f, 0x6c, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x2e, 0x0a, 0x08, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x22, 0xaf, 0x01, 0x0a, 0x0a, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x35, 0x0a,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04,
	0x6b, 0x69, 0x6e, 0x64, 0x12, 0x2f, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x25, 0x0a, 0x0e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x4c, 0x4c, 0x4f, 0x57,
	0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x4e, 0x59, 0x10, 0x01, 0x22, 0x2b, 0x0a, 0x09,
	0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65,
	0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x31, 0x32, 0x73, 0x2f, 0x6f, 0x6f, 0x72,
	0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_common_proto_goTypes = []interface{}{
	(AttributeId_AttributeKind)(0), // 0: common.AttributeId.AttributeKind
	(Permission_PermissionKind)(0), // 1: common.Permission.PermissionKind
	(*AttributeId)(nil),            // 2: common.AttributeId
	(*Attribute)(nil),              // 3: common.Attribute
	(*Int64Attribute)(nil),         // 4: common.Int64Attribute
	(*Float64Attribute)(nil),       // 5: common.Float64Attribute
	(*StringAttribute)(nil),        // 6: common.StringAttribute
	(*BoolAttribute)(nil),          // 7: common.BoolAttribute
	(*Resource)(nil),               // 8: common.Resource
	(*Permission)(nil),             // 9: common.Permission
	(*Condition)(nil),              // 10: common.Condition
}
var file_common_proto_depIdxs = []int32{
	0,  // 0: common.AttributeId.kind:type_name -> common.AttributeId.AttributeKind
	2,  // 1: common.Attribute.id:type_name -> common.AttributeId
	1,  // 2: common.Permission.kind:type_name -> common.Permission.PermissionKind
	10, // 3: common.Permission.condition:type_name -> common.Condition
	4,  // [4:4] is the sub-list for method output_type
	4,  // [4:4] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttributeId); i {
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
		file_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attribute); i {
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
		file_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Int64Attribute); i {
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
		file_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Float64Attribute); i {
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
		file_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringAttribute); i {
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
		file_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoolAttribute); i {
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
		file_common_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
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
		file_common_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Permission); i {
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
		file_common_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Condition); i {
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
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		EnumInfos:         file_common_proto_enumTypes,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}