// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.1
// source: proto/api/v1/subscription.proto

package protogen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type Subscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId             int64                  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CountryId          int64                  `protobuf:"varint,3,opt,name=country_id,json=countryId,proto3" json:"country_id,omitempty"`
	Trial              bool                   `protobuf:"varint,4,opt,name=trial,proto3" json:"trial,omitempty"`
	ExpirationDatetime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=expiration_datetime,json=expirationDatetime,proto3" json:"expiration_datetime,omitempty"`
}

func (x *Subscription) Reset() {
	*x = Subscription{}
	mi := &file_proto_api_v1_subscription_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Subscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscription) ProtoMessage() {}

func (x *Subscription) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_v1_subscription_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subscription.ProtoReflect.Descriptor instead.
func (*Subscription) Descriptor() ([]byte, []int) {
	return file_proto_api_v1_subscription_proto_rawDescGZIP(), []int{0}
}

func (x *Subscription) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Subscription) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Subscription) GetCountryId() int64 {
	if x != nil {
		return x.CountryId
	}
	return 0
}

func (x *Subscription) GetTrial() bool {
	if x != nil {
		return x.Trial
	}
	return false
}

func (x *Subscription) GetExpirationDatetime() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpirationDatetime
	}
	return nil
}

type Subscriptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subscriptions []*Subscription `protobuf:"bytes,1,rep,name=subscriptions,proto3" json:"subscriptions,omitempty"`
}

func (x *Subscriptions) Reset() {
	*x = Subscriptions{}
	mi := &file_proto_api_v1_subscription_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Subscriptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscriptions) ProtoMessage() {}

func (x *Subscriptions) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_v1_subscription_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subscriptions.ProtoReflect.Descriptor instead.
func (*Subscriptions) Descriptor() ([]byte, []int) {
	return file_proto_api_v1_subscription_proto_rawDescGZIP(), []int{1}
}

func (x *Subscriptions) GetSubscriptions() []*Subscription {
	if x != nil {
		return x.Subscriptions
	}
	return nil
}

type GetSubscriptionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CountryId int64 `protobuf:"varint,2,opt,name=country_id,json=countryId,proto3" json:"country_id,omitempty"`
	Active    bool  `protobuf:"varint,3,opt,name=active,proto3" json:"active,omitempty"`
}

func (x *GetSubscriptionsRequest) Reset() {
	*x = GetSubscriptionsRequest{}
	mi := &file_proto_api_v1_subscription_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSubscriptionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubscriptionsRequest) ProtoMessage() {}

func (x *GetSubscriptionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_v1_subscription_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubscriptionsRequest.ProtoReflect.Descriptor instead.
func (*GetSubscriptionsRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_v1_subscription_proto_rawDescGZIP(), []int{2}
}

func (x *GetSubscriptionsRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetSubscriptionsRequest) GetCountryId() int64 {
	if x != nil {
		return x.CountryId
	}
	return 0
}

func (x *GetSubscriptionsRequest) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

type CreateSubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId             int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CountryId          int64                  `protobuf:"varint,2,opt,name=country_id,json=countryId,proto3" json:"country_id,omitempty"`
	Trial              bool                   `protobuf:"varint,3,opt,name=trial,proto3" json:"trial,omitempty"`
	ExpirationDatetime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=expiration_datetime,json=expirationDatetime,proto3" json:"expiration_datetime,omitempty"`
}

func (x *CreateSubscriptionRequest) Reset() {
	*x = CreateSubscriptionRequest{}
	mi := &file_proto_api_v1_subscription_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSubscriptionRequest) ProtoMessage() {}

func (x *CreateSubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_v1_subscription_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSubscriptionRequest.ProtoReflect.Descriptor instead.
func (*CreateSubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_v1_subscription_proto_rawDescGZIP(), []int{3}
}

func (x *CreateSubscriptionRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateSubscriptionRequest) GetCountryId() int64 {
	if x != nil {
		return x.CountryId
	}
	return 0
}

func (x *CreateSubscriptionRequest) GetTrial() bool {
	if x != nil {
		return x.Trial
	}
	return false
}

func (x *CreateSubscriptionRequest) GetExpirationDatetime() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpirationDatetime
	}
	return nil
}

type DeactivateSubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubscriptionId int64 `protobuf:"varint,1,opt,name=subscription_id,json=subscriptionId,proto3" json:"subscription_id,omitempty"`
}

func (x *DeactivateSubscriptionRequest) Reset() {
	*x = DeactivateSubscriptionRequest{}
	mi := &file_proto_api_v1_subscription_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeactivateSubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeactivateSubscriptionRequest) ProtoMessage() {}

func (x *DeactivateSubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_v1_subscription_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeactivateSubscriptionRequest.ProtoReflect.Descriptor instead.
func (*DeactivateSubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_v1_subscription_proto_rawDescGZIP(), []int{4}
}

func (x *DeactivateSubscriptionRequest) GetSubscriptionId() int64 {
	if x != nil {
		return x.SubscriptionId
	}
	return 0
}

type HasActiveSubscriptionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HasActiveSubscription bool `protobuf:"varint,1,opt,name=has_active_subscription,json=hasActiveSubscription,proto3" json:"has_active_subscription,omitempty"`
}

func (x *HasActiveSubscriptionResponse) Reset() {
	*x = HasActiveSubscriptionResponse{}
	mi := &file_proto_api_v1_subscription_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HasActiveSubscriptionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HasActiveSubscriptionResponse) ProtoMessage() {}

func (x *HasActiveSubscriptionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_v1_subscription_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HasActiveSubscriptionResponse.ProtoReflect.Descriptor instead.
func (*HasActiveSubscriptionResponse) Descriptor() ([]byte, []int) {
	return file_proto_api_v1_subscription_proto_rawDescGZIP(), []int{5}
}

func (x *HasActiveSubscriptionResponse) GetHasActiveSubscription() bool {
	if x != nil {
		return x.HasActiveSubscription
	}
	return false
}

var File_proto_api_v1_subscription_proto protoreflect.FileDescriptor

var file_proto_api_v1_subscription_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xb9, 0x01, 0x0a, 0x0c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x72, 0x69, 0x61,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x4b,
	0x0a, 0x13, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x12, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x4a, 0x0a, 0x0d, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x39, 0x0a, 0x0d,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x69, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x22, 0xb6, 0x01, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x72, 0x69, 0x61,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x4b,
	0x0a, 0x13, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x12, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x48, 0x0a, 0x1d, 0x44,
	0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x57, 0x0a, 0x1d, 0x48, 0x61, 0x73, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x17, 0x68, 0x61, 0x73, 0x5f, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x68, 0x61, 0x73, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xd4,
	0x02, 0x0a, 0x13, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x4d, 0x0a, 0x14, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x4c, 0x0a, 0x15, 0x48, 0x61, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x48, 0x61, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a,
	0x16, 0x44, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x44, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_api_v1_subscription_proto_rawDescOnce sync.Once
	file_proto_api_v1_subscription_proto_rawDescData = file_proto_api_v1_subscription_proto_rawDesc
)

func file_proto_api_v1_subscription_proto_rawDescGZIP() []byte {
	file_proto_api_v1_subscription_proto_rawDescOnce.Do(func() {
		file_proto_api_v1_subscription_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_api_v1_subscription_proto_rawDescData)
	})
	return file_proto_api_v1_subscription_proto_rawDescData
}

var file_proto_api_v1_subscription_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_api_v1_subscription_proto_goTypes = []any{
	(*Subscription)(nil),                  // 0: proto.Subscription
	(*Subscriptions)(nil),                 // 1: proto.Subscriptions
	(*GetSubscriptionsRequest)(nil),       // 2: proto.GetSubscriptionsRequest
	(*CreateSubscriptionRequest)(nil),     // 3: proto.CreateSubscriptionRequest
	(*DeactivateSubscriptionRequest)(nil), // 4: proto.DeactivateSubscriptionRequest
	(*HasActiveSubscriptionResponse)(nil), // 5: proto.HasActiveSubscriptionResponse
	(*timestamppb.Timestamp)(nil),         // 6: google.protobuf.Timestamp
	(*UserId)(nil),                        // 7: proto.UserId
	(*emptypb.Empty)(nil),                 // 8: google.protobuf.Empty
}
var file_proto_api_v1_subscription_proto_depIdxs = []int32{
	6, // 0: proto.Subscription.expiration_datetime:type_name -> google.protobuf.Timestamp
	0, // 1: proto.Subscriptions.subscriptions:type_name -> proto.Subscription
	6, // 2: proto.CreateSubscriptionRequest.expiration_datetime:type_name -> google.protobuf.Timestamp
	2, // 3: proto.SubscriptionService.GetSubscriptions:input_type -> proto.GetSubscriptionsRequest
	3, // 4: proto.SubscriptionService.ActivateSubscription:input_type -> proto.CreateSubscriptionRequest
	7, // 5: proto.SubscriptionService.HasActiveSubscription:input_type -> proto.UserId
	4, // 6: proto.SubscriptionService.DeactivateSubscription:input_type -> proto.DeactivateSubscriptionRequest
	1, // 7: proto.SubscriptionService.GetSubscriptions:output_type -> proto.Subscriptions
	0, // 8: proto.SubscriptionService.ActivateSubscription:output_type -> proto.Subscription
	5, // 9: proto.SubscriptionService.HasActiveSubscription:output_type -> proto.HasActiveSubscriptionResponse
	8, // 10: proto.SubscriptionService.DeactivateSubscription:output_type -> google.protobuf.Empty
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_api_v1_subscription_proto_init() }
func file_proto_api_v1_subscription_proto_init() {
	if File_proto_api_v1_subscription_proto != nil {
		return
	}
	file_proto_api_v1_user_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_api_v1_subscription_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_api_v1_subscription_proto_goTypes,
		DependencyIndexes: file_proto_api_v1_subscription_proto_depIdxs,
		MessageInfos:      file_proto_api_v1_subscription_proto_msgTypes,
	}.Build()
	File_proto_api_v1_subscription_proto = out.File
	file_proto_api_v1_subscription_proto_rawDesc = nil
	file_proto_api_v1_subscription_proto_goTypes = nil
	file_proto_api_v1_subscription_proto_depIdxs = nil
}
