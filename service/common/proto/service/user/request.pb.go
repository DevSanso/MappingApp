// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.3
// source: service/user/request.proto

package user

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

type Request_Method int32

const (
	Request_CREATE Request_Method = 0
	Request_UPDATE Request_Method = 1
	Request_DELETE Request_Method = 2
	Request_READ   Request_Method = 3
)

// Enum value maps for Request_Method.
var (
	Request_Method_name = map[int32]string{
		0: "CREATE",
		1: "UPDATE",
		2: "DELETE",
		3: "READ",
	}
	Request_Method_value = map[string]int32{
		"CREATE": 0,
		"UPDATE": 1,
		"DELETE": 2,
		"READ":   3,
	}
)

func (x Request_Method) Enum() *Request_Method {
	p := new(Request_Method)
	*p = x
	return p
}

func (x Request_Method) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Request_Method) Descriptor() protoreflect.EnumDescriptor {
	return file_service_user_request_proto_enumTypes[0].Descriptor()
}

func (Request_Method) Type() protoreflect.EnumType {
	return &file_service_user_request_proto_enumTypes[0]
}

func (x Request_Method) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Request_Method.Descriptor instead.
func (Request_Method) EnumDescriptor() ([]byte, []int) {
	return file_service_user_request_proto_rawDescGZIP(), []int{0, 0}
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method   Request_Method `protobuf:"varint,1,opt,name=method,proto3,enum=user.Request_Method" json:"method,omitempty"`
	UserUuid string         `protobuf:"bytes,2,opt,name=userUuid,proto3" json:"userUuid,omitempty"`
	// Types that are assignable to Body:
	//	*Request_ReadBody_
	//	*Request_OtherBody_
	Body isRequest_Body `protobuf_oneof:"body"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_service_user_request_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetMethod() Request_Method {
	if x != nil {
		return x.Method
	}
	return Request_CREATE
}

func (x *Request) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

func (m *Request) GetBody() isRequest_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (x *Request) GetReadBody() *Request_ReadBody {
	if x, ok := x.GetBody().(*Request_ReadBody_); ok {
		return x.ReadBody
	}
	return nil
}

func (x *Request) GetOtherBody() *Request_OtherBody {
	if x, ok := x.GetBody().(*Request_OtherBody_); ok {
		return x.OtherBody
	}
	return nil
}

type isRequest_Body interface {
	isRequest_Body()
}

type Request_ReadBody_ struct {
	ReadBody *Request_ReadBody `protobuf:"bytes,3,opt,name=read_body,json=readBody,proto3,oneof"`
}

type Request_OtherBody_ struct {
	OtherBody *Request_OtherBody `protobuf:"bytes,4,opt,name=other_body,json=otherBody,proto3,oneof"`
}

func (*Request_ReadBody_) isRequest_Body() {}

func (*Request_OtherBody_) isRequest_Body() {}

type Request_Privacy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Age     *int32   `protobuf:"varint,1,opt,name=age,proto3,oneof" json:"age,omitempty"`
	Stature *float32 `protobuf:"fixed32,2,opt,name=stature,proto3,oneof" json:"stature,omitempty"`
	Weight  *float32 `protobuf:"fixed32,3,opt,name=weight,proto3,oneof" json:"weight,omitempty"`
}

func (x *Request_Privacy) Reset() {
	*x = Request_Privacy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request_Privacy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request_Privacy) ProtoMessage() {}

func (x *Request_Privacy) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request_Privacy.ProtoReflect.Descriptor instead.
func (*Request_Privacy) Descriptor() ([]byte, []int) {
	return file_service_user_request_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Request_Privacy) GetAge() int32 {
	if x != nil && x.Age != nil {
		return *x.Age
	}
	return 0
}

func (x *Request_Privacy) GetStature() float32 {
	if x != nil && x.Stature != nil {
		return *x.Stature
	}
	return 0
}

func (x *Request_Privacy) GetWeight() float32 {
	if x != nil && x.Weight != nil {
		return *x.Weight
	}
	return 0
}

type Request_Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address1 string  `protobuf:"bytes,1,opt,name=address1,proto3" json:"address1,omitempty"`
	Address2 *string `protobuf:"bytes,2,opt,name=address2,proto3,oneof" json:"address2,omitempty"`
}

func (x *Request_Address) Reset() {
	*x = Request_Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request_Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request_Address) ProtoMessage() {}

func (x *Request_Address) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request_Address.ProtoReflect.Descriptor instead.
func (*Request_Address) Descriptor() ([]byte, []int) {
	return file_service_user_request_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Request_Address) GetAddress1() string {
	if x != nil {
		return x.Address1
	}
	return ""
}

func (x *Request_Address) GetAddress2() string {
	if x != nil && x.Address2 != nil {
		return *x.Address2
	}
	return ""
}

type Request_ReadBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReadPrivacy bool `protobuf:"varint,1,opt,name=read_privacy,json=readPrivacy,proto3" json:"read_privacy,omitempty"`
	ReadAddress bool `protobuf:"varint,2,opt,name=read_address,json=readAddress,proto3" json:"read_address,omitempty"`
}

func (x *Request_ReadBody) Reset() {
	*x = Request_ReadBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request_ReadBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request_ReadBody) ProtoMessage() {}

func (x *Request_ReadBody) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request_ReadBody.ProtoReflect.Descriptor instead.
func (*Request_ReadBody) Descriptor() ([]byte, []int) {
	return file_service_user_request_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Request_ReadBody) GetReadPrivacy() bool {
	if x != nil {
		return x.ReadPrivacy
	}
	return false
}

func (x *Request_ReadBody) GetReadAddress() bool {
	if x != nil {
		return x.ReadAddress
	}
	return false
}

type Request_OtherBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Privacy *Request_Privacy `protobuf:"bytes,1,opt,name=privacy,proto3,oneof" json:"privacy,omitempty"`
	Address *Request_Address `protobuf:"bytes,2,opt,name=address,proto3,oneof" json:"address,omitempty"`
}

func (x *Request_OtherBody) Reset() {
	*x = Request_OtherBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_request_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request_OtherBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request_OtherBody) ProtoMessage() {}

func (x *Request_OtherBody) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_request_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request_OtherBody.ProtoReflect.Descriptor instead.
func (*Request_OtherBody) Descriptor() ([]byte, []int) {
	return file_service_user_request_proto_rawDescGZIP(), []int{0, 3}
}

func (x *Request_OtherBody) GetPrivacy() *Request_Privacy {
	if x != nil {
		return x.Privacy
	}
	return nil
}

func (x *Request_OtherBody) GetAddress() *Request_Address {
	if x != nil {
		return x.Address
	}
	return nil
}

var File_service_user_request_proto protoreflect.FileDescriptor

var file_service_user_request_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x22, 0xba, 0x05, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c,
	0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x35, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x64,
	0x5f, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x42,
	0x6f, 0x64, 0x79, 0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x61, 0x64, 0x42, 0x6f, 0x64, 0x79, 0x12,
	0x38, 0x0a, 0x0a, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x42, 0x6f, 0x64, 0x79, 0x48, 0x00, 0x52, 0x09,
	0x6f, 0x74, 0x68, 0x65, 0x72, 0x42, 0x6f, 0x64, 0x79, 0x1a, 0x7b, 0x0a, 0x07, 0x50, 0x72, 0x69,
	0x76, 0x61, 0x63, 0x79, 0x12, 0x15, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x00, 0x52, 0x03, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x48, 0x01, 0x52, 0x07,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x72, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x77, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x48, 0x02, 0x52, 0x06, 0x77, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x61, 0x67, 0x65, 0x42,
	0x0a, 0x0a, 0x08, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f,
	0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x1a, 0x53, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x31, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x31, 0x12, 0x1f, 0x0a,
	0x08, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x08, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x32, 0x88, 0x01, 0x01, 0x42, 0x0b,
	0x0a, 0x09, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x32, 0x1a, 0x50, 0x0a, 0x08, 0x52,
	0x65, 0x61, 0x64, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x61, 0x64, 0x5f,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x72,
	0x65, 0x61, 0x64, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65,
	0x61, 0x64, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x72, 0x65, 0x61, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x1a, 0x8f, 0x01,
	0x0a, 0x09, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x34, 0x0a, 0x07, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x63, 0x79, 0x48, 0x00, 0x52, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x88, 0x01,
	0x01, 0x12, 0x34, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x48, 0x01, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x63, 0x79, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22,
	0x36, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x52, 0x45,
	0x41, 0x54, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10,
	0x01, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x02, 0x12, 0x08, 0x0a,
	0x04, 0x52, 0x45, 0x41, 0x44, 0x10, 0x03, 0x42, 0x06, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x42,
	0x14, 0x5a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_user_request_proto_rawDescOnce sync.Once
	file_service_user_request_proto_rawDescData = file_service_user_request_proto_rawDesc
)

func file_service_user_request_proto_rawDescGZIP() []byte {
	file_service_user_request_proto_rawDescOnce.Do(func() {
		file_service_user_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_user_request_proto_rawDescData)
	})
	return file_service_user_request_proto_rawDescData
}

var file_service_user_request_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_service_user_request_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_service_user_request_proto_goTypes = []interface{}{
	(Request_Method)(0),       // 0: user.Request.Method
	(*Request)(nil),           // 1: user.Request
	(*Request_Privacy)(nil),   // 2: user.Request.Privacy
	(*Request_Address)(nil),   // 3: user.Request.Address
	(*Request_ReadBody)(nil),  // 4: user.Request.ReadBody
	(*Request_OtherBody)(nil), // 5: user.Request.OtherBody
}
var file_service_user_request_proto_depIdxs = []int32{
	0, // 0: user.Request.method:type_name -> user.Request.Method
	4, // 1: user.Request.read_body:type_name -> user.Request.ReadBody
	5, // 2: user.Request.other_body:type_name -> user.Request.OtherBody
	2, // 3: user.Request.OtherBody.privacy:type_name -> user.Request.Privacy
	3, // 4: user.Request.OtherBody.address:type_name -> user.Request.Address
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_service_user_request_proto_init() }
func file_service_user_request_proto_init() {
	if File_service_user_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_user_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_service_user_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request_Privacy); i {
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
		file_service_user_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request_Address); i {
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
		file_service_user_request_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request_ReadBody); i {
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
		file_service_user_request_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request_OtherBody); i {
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
	file_service_user_request_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Request_ReadBody_)(nil),
		(*Request_OtherBody_)(nil),
	}
	file_service_user_request_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_service_user_request_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_service_user_request_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_user_request_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_service_user_request_proto_goTypes,
		DependencyIndexes: file_service_user_request_proto_depIdxs,
		EnumInfos:         file_service_user_request_proto_enumTypes,
		MessageInfos:      file_service_user_request_proto_msgTypes,
	}.Build()
	File_service_user_request_proto = out.File
	file_service_user_request_proto_rawDesc = nil
	file_service_user_request_proto_goTypes = nil
	file_service_user_request_proto_depIdxs = nil
}
