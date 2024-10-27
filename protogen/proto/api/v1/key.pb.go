// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: proto/api/v1/key.proto

package protogen

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

type KeyType int32

const (
	KeyType_Text  KeyType = 0
	KeyType_File  KeyType = 1
	KeyType_Photo KeyType = 2
)

// Enum value maps for KeyType.
var (
	KeyType_name = map[int32]string{
		0: "Text",
		1: "File",
		2: "Photo",
	}
	KeyType_value = map[string]int32{
		"Text":  0,
		"File":  1,
		"Photo": 2,
	}
)

func (x KeyType) Enum() *KeyType {
	p := new(KeyType)
	*p = x
	return p
}

func (x KeyType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KeyType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_api_v1_key_proto_enumTypes[0].Descriptor()
}

func (KeyType) Type() protoreflect.EnumType {
	return &file_proto_api_v1_key_proto_enumTypes[0]
}

func (x KeyType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KeyType.Descriptor instead.
func (KeyType) EnumDescriptor() ([]byte, []int) {
	return file_proto_api_v1_key_proto_rawDescGZIP(), []int{0}
}

type Key struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Data           []byte  `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	KeyType        KeyType `protobuf:"varint,3,opt,name=key_type,json=keyType,proto3,enum=proto.KeyType" json:"key_type,omitempty"`
	SubscriptionId int64   `protobuf:"varint,4,opt,name=subscription_id,json=subscriptionId,proto3" json:"subscription_id,omitempty"`
}

func (x *Key) Reset() {
	*x = Key{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_v1_key_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Key) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Key) ProtoMessage() {}

func (x *Key) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_v1_key_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Key.ProtoReflect.Descriptor instead.
func (*Key) Descriptor() ([]byte, []int) {
	return file_proto_api_v1_key_proto_rawDescGZIP(), []int{0}
}

func (x *Key) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Key) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Key) GetKeyType() KeyType {
	if x != nil {
		return x.KeyType
	}
	return KeyType_Text
}

func (x *Key) GetSubscriptionId() int64 {
	if x != nil {
		return x.SubscriptionId
	}
	return 0
}

type KeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CountryId int64 `protobuf:"varint,2,opt,name=country_id,json=countryId,proto3" json:"country_id,omitempty"`
}

func (x *KeyRequest) Reset() {
	*x = KeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_v1_key_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyRequest) ProtoMessage() {}

func (x *KeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_v1_key_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyRequest.ProtoReflect.Descriptor instead.
func (*KeyRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_v1_key_proto_rawDescGZIP(), []int{1}
}

func (x *KeyRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *KeyRequest) GetCountryId() int64 {
	if x != nil {
		return x.CountryId
	}
	return 0
}

var File_proto_api_v1_key_proto protoreflect.FileDescriptor

var file_proto_api_v1_key_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6b,
	0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7d, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x29, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65,
	0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x27,
	0x0a, 0x0f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x44, 0x0a, 0x0a, 0x4b, 0x65, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x2a, 0x28, 0x0a,
	0x07, 0x4b, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74,
	0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05,
	0x50, 0x68, 0x6f, 0x74, 0x6f, 0x10, 0x02, 0x32, 0x35, 0x0a, 0x0a, 0x4b, 0x65, 0x79, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x12,
	0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x42, 0x0c,
	0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_api_v1_key_proto_rawDescOnce sync.Once
	file_proto_api_v1_key_proto_rawDescData = file_proto_api_v1_key_proto_rawDesc
)

func file_proto_api_v1_key_proto_rawDescGZIP() []byte {
	file_proto_api_v1_key_proto_rawDescOnce.Do(func() {
		file_proto_api_v1_key_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_api_v1_key_proto_rawDescData)
	})
	return file_proto_api_v1_key_proto_rawDescData
}

var file_proto_api_v1_key_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_api_v1_key_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_api_v1_key_proto_goTypes = []any{
	(KeyType)(0),       // 0: proto.KeyType
	(*Key)(nil),        // 1: proto.Key
	(*KeyRequest)(nil), // 2: proto.KeyRequest
}
var file_proto_api_v1_key_proto_depIdxs = []int32{
	0, // 0: proto.Key.key_type:type_name -> proto.KeyType
	2, // 1: proto.KeyService.GetKey:input_type -> proto.KeyRequest
	1, // 2: proto.KeyService.GetKey:output_type -> proto.Key
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_api_v1_key_proto_init() }
func file_proto_api_v1_key_proto_init() {
	if File_proto_api_v1_key_proto != nil {
		return
	}
	file_proto_api_v1_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_api_v1_key_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Key); i {
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
		file_proto_api_v1_key_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*KeyRequest); i {
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
			RawDescriptor: file_proto_api_v1_key_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_api_v1_key_proto_goTypes,
		DependencyIndexes: file_proto_api_v1_key_proto_depIdxs,
		EnumInfos:         file_proto_api_v1_key_proto_enumTypes,
		MessageInfos:      file_proto_api_v1_key_proto_msgTypes,
	}.Build()
	File_proto_api_v1_key_proto = out.File
	file_proto_api_v1_key_proto_rawDesc = nil
	file_proto_api_v1_key_proto_goTypes = nil
	file_proto_api_v1_key_proto_depIdxs = nil
}
