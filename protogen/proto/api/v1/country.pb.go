// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: proto/api/v1/country.proto

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

type Country struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Country) Reset() {
	*x = Country{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_v1_country_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Country) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Country) ProtoMessage() {}

func (x *Country) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_v1_country_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Country.ProtoReflect.Descriptor instead.
func (*Country) Descriptor() ([]byte, []int) {
	return file_proto_api_v1_country_proto_rawDescGZIP(), []int{0}
}

func (x *Country) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Country) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CountryName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CountryName) Reset() {
	*x = CountryName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_v1_country_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountryName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountryName) ProtoMessage() {}

func (x *CountryName) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_v1_country_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountryName.ProtoReflect.Descriptor instead.
func (*CountryName) Descriptor() ([]byte, []int) {
	return file_proto_api_v1_country_proto_rawDescGZIP(), []int{1}
}

func (x *CountryName) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_proto_api_v1_country_proto protoreflect.FileDescriptor

var file_proto_api_v1_country_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x07,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x21, 0x0a, 0x0b, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x48,
	0x0a, 0x0e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x36, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_api_v1_country_proto_rawDescOnce sync.Once
	file_proto_api_v1_country_proto_rawDescData = file_proto_api_v1_country_proto_rawDesc
)

func file_proto_api_v1_country_proto_rawDescGZIP() []byte {
	file_proto_api_v1_country_proto_rawDescOnce.Do(func() {
		file_proto_api_v1_country_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_api_v1_country_proto_rawDescData)
	})
	return file_proto_api_v1_country_proto_rawDescData
}

var file_proto_api_v1_country_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_api_v1_country_proto_goTypes = []any{
	(*Country)(nil),     // 0: proto.Country
	(*CountryName)(nil), // 1: proto.CountryName
}
var file_proto_api_v1_country_proto_depIdxs = []int32{
	1, // 0: proto.CountryService.GetCountryByName:input_type -> proto.CountryName
	0, // 1: proto.CountryService.GetCountryByName:output_type -> proto.Country
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_api_v1_country_proto_init() }
func file_proto_api_v1_country_proto_init() {
	if File_proto_api_v1_country_proto != nil {
		return
	}
	file_proto_api_v1_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_api_v1_country_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Country); i {
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
		file_proto_api_v1_country_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CountryName); i {
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
			RawDescriptor: file_proto_api_v1_country_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_api_v1_country_proto_goTypes,
		DependencyIndexes: file_proto_api_v1_country_proto_depIdxs,
		MessageInfos:      file_proto_api_v1_country_proto_msgTypes,
	}.Build()
	File_proto_api_v1_country_proto = out.File
	file_proto_api_v1_country_proto_rawDesc = nil
	file_proto_api_v1_country_proto_goTypes = nil
	file_proto_api_v1_country_proto_depIdxs = nil
}