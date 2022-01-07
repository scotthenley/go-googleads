// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.18.1
// source: google/ads/googleads/v8/resources/account_link.proto

package resources

import (
	proto "github.com/golang/protobuf/proto"
	enums "github.com/scotthenley/go-googleads/pb/v8/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Represents the data sharing connection between a Google Ads account and
// another account
type AccountLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. Resource name of the account link.
	// AccountLink resource names have the form:
	// `customers/{customer_id}/accountLinks/{account_link_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the link.
	// This field is read only.
	AccountLinkId *int64 `protobuf:"varint,8,opt,name=account_link_id,json=accountLinkId,proto3,oneof" json:"account_link_id,omitempty"`
	// The status of the link.
	Status enums.AccountLinkStatusEnum_AccountLinkStatus `protobuf:"varint,3,opt,name=status,proto3,enum=google.ads.googleads.v8.enums.AccountLinkStatusEnum_AccountLinkStatus" json:"status,omitempty"`
	// Output only. The type of the linked account.
	Type enums.LinkedAccountTypeEnum_LinkedAccountType `protobuf:"varint,4,opt,name=type,proto3,enum=google.ads.googleads.v8.enums.LinkedAccountTypeEnum_LinkedAccountType" json:"type,omitempty"`
	// An account linked to this Google Ads account.
	//
	// Types that are assignable to LinkedAccount:
	//	*AccountLink_ThirdPartyAppAnalytics
	//	*AccountLink_DataPartner
	//	*AccountLink_GoogleAds
	LinkedAccount isAccountLink_LinkedAccount `protobuf_oneof:"linked_account"`
}

func (x *AccountLink) Reset() {
	*x = AccountLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v8_resources_account_link_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountLink) ProtoMessage() {}

func (x *AccountLink) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v8_resources_account_link_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountLink.ProtoReflect.Descriptor instead.
func (*AccountLink) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v8_resources_account_link_proto_rawDescGZIP(), []int{0}
}

func (x *AccountLink) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *AccountLink) GetAccountLinkId() int64 {
	if x != nil && x.AccountLinkId != nil {
		return *x.AccountLinkId
	}
	return 0
}

func (x *AccountLink) GetStatus() enums.AccountLinkStatusEnum_AccountLinkStatus {
	if x != nil {
		return x.Status
	}
	return enums.AccountLinkStatusEnum_UNSPECIFIED
}

func (x *AccountLink) GetType() enums.LinkedAccountTypeEnum_LinkedAccountType {
	if x != nil {
		return x.Type
	}
	return enums.LinkedAccountTypeEnum_UNSPECIFIED
}

func (m *AccountLink) GetLinkedAccount() isAccountLink_LinkedAccount {
	if m != nil {
		return m.LinkedAccount
	}
	return nil
}

func (x *AccountLink) GetThirdPartyAppAnalytics() *ThirdPartyAppAnalyticsLinkIdentifier {
	if x, ok := x.GetLinkedAccount().(*AccountLink_ThirdPartyAppAnalytics); ok {
		return x.ThirdPartyAppAnalytics
	}
	return nil
}

func (x *AccountLink) GetDataPartner() *DataPartnerLinkIdentifier {
	if x, ok := x.GetLinkedAccount().(*AccountLink_DataPartner); ok {
		return x.DataPartner
	}
	return nil
}

func (x *AccountLink) GetGoogleAds() *GoogleAdsLinkIdentifier {
	if x, ok := x.GetLinkedAccount().(*AccountLink_GoogleAds); ok {
		return x.GoogleAds
	}
	return nil
}

type isAccountLink_LinkedAccount interface {
	isAccountLink_LinkedAccount()
}

type AccountLink_ThirdPartyAppAnalytics struct {
	// Immutable. A third party app analytics link.
	ThirdPartyAppAnalytics *ThirdPartyAppAnalyticsLinkIdentifier `protobuf:"bytes,5,opt,name=third_party_app_analytics,json=thirdPartyAppAnalytics,proto3,oneof"`
}

type AccountLink_DataPartner struct {
	// Output only. Data partner link.
	DataPartner *DataPartnerLinkIdentifier `protobuf:"bytes,6,opt,name=data_partner,json=dataPartner,proto3,oneof"`
}

type AccountLink_GoogleAds struct {
	// Output only. Google Ads link.
	GoogleAds *GoogleAdsLinkIdentifier `protobuf:"bytes,7,opt,name=google_ads,json=googleAds,proto3,oneof"`
}

func (*AccountLink_ThirdPartyAppAnalytics) isAccountLink_LinkedAccount() {}

func (*AccountLink_DataPartner) isAccountLink_LinkedAccount() {}

func (*AccountLink_GoogleAds) isAccountLink_LinkedAccount() {}

// The identifiers of a Third Party App Analytics Link.
type ThirdPartyAppAnalyticsLinkIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The ID of the app analytics provider.
	// This field should not be empty when creating a new third
	// party app analytics link. It is unable to be modified after the creation of
	// the link.
	AppAnalyticsProviderId *int64 `protobuf:"varint,4,opt,name=app_analytics_provider_id,json=appAnalyticsProviderId,proto3,oneof" json:"app_analytics_provider_id,omitempty"`
	// Immutable. A string that uniquely identifies a mobile application from which the data
	// was collected to the Google Ads API. For iOS, the ID string is the 9 digit
	// string that appears at the end of an App Store URL (e.g., "422689480" for
	// "Gmail" whose App Store link is
	// https://apps.apple.com/us/app/gmail-email-by-google/id422689480). For
	// Android, the ID string is the application's package name (e.g.,
	// "com.google.android.gm" for "Gmail" given Google Play link
	// https://play.google.com/store/apps/details?id=com.google.android.gm)
	// This field should not be empty when creating a new third
	// party app analytics link. It is unable to be modified after the creation of
	// the link.
	AppId *string `protobuf:"bytes,5,opt,name=app_id,json=appId,proto3,oneof" json:"app_id,omitempty"`
	// Immutable. The vendor of the app.
	// This field should not be empty when creating a new third
	// party app analytics link. It is unable to be modified after the creation of
	// the link.
	AppVendor enums.MobileAppVendorEnum_MobileAppVendor `protobuf:"varint,3,opt,name=app_vendor,json=appVendor,proto3,enum=google.ads.googleads.v8.enums.MobileAppVendorEnum_MobileAppVendor" json:"app_vendor,omitempty"`
}

func (x *ThirdPartyAppAnalyticsLinkIdentifier) Reset() {
	*x = ThirdPartyAppAnalyticsLinkIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v8_resources_account_link_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThirdPartyAppAnalyticsLinkIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThirdPartyAppAnalyticsLinkIdentifier) ProtoMessage() {}

func (x *ThirdPartyAppAnalyticsLinkIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v8_resources_account_link_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThirdPartyAppAnalyticsLinkIdentifier.ProtoReflect.Descriptor instead.
func (*ThirdPartyAppAnalyticsLinkIdentifier) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v8_resources_account_link_proto_rawDescGZIP(), []int{1}
}

func (x *ThirdPartyAppAnalyticsLinkIdentifier) GetAppAnalyticsProviderId() int64 {
	if x != nil && x.AppAnalyticsProviderId != nil {
		return *x.AppAnalyticsProviderId
	}
	return 0
}

func (x *ThirdPartyAppAnalyticsLinkIdentifier) GetAppId() string {
	if x != nil && x.AppId != nil {
		return *x.AppId
	}
	return ""
}

func (x *ThirdPartyAppAnalyticsLinkIdentifier) GetAppVendor() enums.MobileAppVendorEnum_MobileAppVendor {
	if x != nil {
		return x.AppVendor
	}
	return enums.MobileAppVendorEnum_UNSPECIFIED
}

// The identifier for Data Partner account.
type DataPartnerLinkIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The customer ID of the Data partner account.
	// This field is required and should not be empty when creating a new
	// data partner link. It is unable to be modified after the creation of
	// the link.
	DataPartnerId *int64 `protobuf:"varint,1,opt,name=data_partner_id,json=dataPartnerId,proto3,oneof" json:"data_partner_id,omitempty"`
}

func (x *DataPartnerLinkIdentifier) Reset() {
	*x = DataPartnerLinkIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v8_resources_account_link_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataPartnerLinkIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataPartnerLinkIdentifier) ProtoMessage() {}

func (x *DataPartnerLinkIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v8_resources_account_link_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataPartnerLinkIdentifier.ProtoReflect.Descriptor instead.
func (*DataPartnerLinkIdentifier) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v8_resources_account_link_proto_rawDescGZIP(), []int{2}
}

func (x *DataPartnerLinkIdentifier) GetDataPartnerId() int64 {
	if x != nil && x.DataPartnerId != nil {
		return *x.DataPartnerId
	}
	return 0
}

// The identifier for Google Ads account.
type GoogleAdsLinkIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the Google Ads account.
	// This field is required and should not be empty when creating a new
	// Google Ads link. It is unable to be modified after the creation of
	// the link.
	Customer *string `protobuf:"bytes,3,opt,name=customer,proto3,oneof" json:"customer,omitempty"`
}

func (x *GoogleAdsLinkIdentifier) Reset() {
	*x = GoogleAdsLinkIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v8_resources_account_link_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoogleAdsLinkIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoogleAdsLinkIdentifier) ProtoMessage() {}

func (x *GoogleAdsLinkIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v8_resources_account_link_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoogleAdsLinkIdentifier.ProtoReflect.Descriptor instead.
func (*GoogleAdsLinkIdentifier) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v8_resources_account_link_proto_rawDescGZIP(), []int{3}
}

func (x *GoogleAdsLinkIdentifier) GetCustomer() string {
	if x != nil && x.Customer != nil {
		return *x.Customer
	}
	return ""
}

var File_google_ads_googleads_v8_resources_account_link_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v8_resources_account_link_proto_rawDesc = []byte{
	0x0a, 0x34, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x38, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f,
	0x76, 0x38, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x38, 0x2f, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2f, 0x76, 0x38, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x6d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x5f, 0x61, 0x70, 0x70, 0x5f, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x06, 0x0a,
	0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x51, 0x0a, 0x0d,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x2c, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x26, 0x0a, 0x24, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6e,
	0x6b, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x30, 0x0a, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f,
	0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x01, 0x52,
	0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x64, 0x88, 0x01,
	0x01, 0x12, 0x5e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x46, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c,
	0x69, 0x6e, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x5f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x46, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x89, 0x01, 0x0a, 0x19, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72,
	0x74, 0x79, 0x5f, 0x61, 0x70, 0x70, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x47, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x68, 0x69, 0x72, 0x64,
	0x50, 0x61, 0x72, 0x74, 0x79, 0x41, 0x70, 0x70, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63,
	0x73, 0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x42,
	0x03, 0xe0, 0x41, 0x05, 0x48, 0x00, 0x52, 0x16, 0x74, 0x68, 0x69, 0x72, 0x64, 0x50, 0x61, 0x72,
	0x74, 0x79, 0x41, 0x70, 0x70, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x12, 0x66,
	0x0a, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x50, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x50,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x12, 0x60, 0x0a, 0x0a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x5f, 0x61, 0x64, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x38, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x00, 0x52, 0x09, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x61, 0xea, 0x41, 0x5e, 0x0a, 0x24, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c,
	0x69, 0x6e, 0x6b, 0x12, 0x36, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x2f, 0x7b, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x10, 0x0a, 0x0e, 0x6c,
	0x69, 0x6e, 0x6b, 0x65, 0x64, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x12, 0x0a,
	0x10, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x69,
	0x64, 0x22, 0x9d, 0x02, 0x0a, 0x24, 0x54, 0x68, 0x69, 0x72, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79,
	0x41, 0x70, 0x70, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x4c, 0x69, 0x6e, 0x6b,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x43, 0x0a, 0x19, 0x61, 0x70,
	0x70, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0,
	0x41, 0x05, 0x48, 0x00, 0x52, 0x16, 0x61, 0x70, 0x70, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69,
	0x63, 0x73, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x1f, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x05, 0x48, 0x01, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x66, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x5f, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x42, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x41, 0x70, 0x70, 0x56, 0x65,
	0x6e, 0x64, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x41,
	0x70, 0x70, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x09, 0x61,
	0x70, 0x70, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x42, 0x1c, 0x0a, 0x1a, 0x5f, 0x61, 0x70, 0x70,
	0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x61, 0x70, 0x70, 0x5f, 0x69,
	0x64, 0x22, 0x61, 0x0a, 0x19, 0x44, 0x61, 0x74, 0x61, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72,
	0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x30,
	0x0a, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x48, 0x00, 0x52, 0x0d,
	0x64, 0x61, 0x74, 0x61, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01,
	0x42, 0x12, 0x0a, 0x10, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x22, 0x72, 0x0a, 0x17, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12,
	0x4a, 0x0a, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x29, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x48, 0x00, 0x52, 0x08,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0xfd, 0x01, 0x0a, 0x25, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x42, 0x10, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64,
	0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x38, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e,
	0x56, 0x38, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x21, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x5c, 0x56, 0x38, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0xea, 0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x38, 0x3a, 0x3a, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v8_resources_account_link_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v8_resources_account_link_proto_rawDescData = file_google_ads_googleads_v8_resources_account_link_proto_rawDesc
)

func file_google_ads_googleads_v8_resources_account_link_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v8_resources_account_link_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v8_resources_account_link_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v8_resources_account_link_proto_rawDescData)
	})
	return file_google_ads_googleads_v8_resources_account_link_proto_rawDescData
}

var file_google_ads_googleads_v8_resources_account_link_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_ads_googleads_v8_resources_account_link_proto_goTypes = []interface{}{
	(*AccountLink)(nil),                                // 0: google.ads.googleads.v8.resources.AccountLink
	(*ThirdPartyAppAnalyticsLinkIdentifier)(nil),       // 1: google.ads.googleads.v8.resources.ThirdPartyAppAnalyticsLinkIdentifier
	(*DataPartnerLinkIdentifier)(nil),                  // 2: google.ads.googleads.v8.resources.DataPartnerLinkIdentifier
	(*GoogleAdsLinkIdentifier)(nil),                    // 3: google.ads.googleads.v8.resources.GoogleAdsLinkIdentifier
	(enums.AccountLinkStatusEnum_AccountLinkStatus)(0), // 4: google.ads.googleads.v8.enums.AccountLinkStatusEnum.AccountLinkStatus
	(enums.LinkedAccountTypeEnum_LinkedAccountType)(0), // 5: google.ads.googleads.v8.enums.LinkedAccountTypeEnum.LinkedAccountType
	(enums.MobileAppVendorEnum_MobileAppVendor)(0),     // 6: google.ads.googleads.v8.enums.MobileAppVendorEnum.MobileAppVendor
}
var file_google_ads_googleads_v8_resources_account_link_proto_depIdxs = []int32{
	4, // 0: google.ads.googleads.v8.resources.AccountLink.status:type_name -> google.ads.googleads.v8.enums.AccountLinkStatusEnum.AccountLinkStatus
	5, // 1: google.ads.googleads.v8.resources.AccountLink.type:type_name -> google.ads.googleads.v8.enums.LinkedAccountTypeEnum.LinkedAccountType
	1, // 2: google.ads.googleads.v8.resources.AccountLink.third_party_app_analytics:type_name -> google.ads.googleads.v8.resources.ThirdPartyAppAnalyticsLinkIdentifier
	2, // 3: google.ads.googleads.v8.resources.AccountLink.data_partner:type_name -> google.ads.googleads.v8.resources.DataPartnerLinkIdentifier
	3, // 4: google.ads.googleads.v8.resources.AccountLink.google_ads:type_name -> google.ads.googleads.v8.resources.GoogleAdsLinkIdentifier
	6, // 5: google.ads.googleads.v8.resources.ThirdPartyAppAnalyticsLinkIdentifier.app_vendor:type_name -> google.ads.googleads.v8.enums.MobileAppVendorEnum.MobileAppVendor
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v8_resources_account_link_proto_init() }
func file_google_ads_googleads_v8_resources_account_link_proto_init() {
	if File_google_ads_googleads_v8_resources_account_link_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v8_resources_account_link_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountLink); i {
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
		file_google_ads_googleads_v8_resources_account_link_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThirdPartyAppAnalyticsLinkIdentifier); i {
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
		file_google_ads_googleads_v8_resources_account_link_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataPartnerLinkIdentifier); i {
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
		file_google_ads_googleads_v8_resources_account_link_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoogleAdsLinkIdentifier); i {
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
	file_google_ads_googleads_v8_resources_account_link_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*AccountLink_ThirdPartyAppAnalytics)(nil),
		(*AccountLink_DataPartner)(nil),
		(*AccountLink_GoogleAds)(nil),
	}
	file_google_ads_googleads_v8_resources_account_link_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_google_ads_googleads_v8_resources_account_link_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_google_ads_googleads_v8_resources_account_link_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v8_resources_account_link_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v8_resources_account_link_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v8_resources_account_link_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v8_resources_account_link_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v8_resources_account_link_proto = out.File
	file_google_ads_googleads_v8_resources_account_link_proto_rawDesc = nil
	file_google_ads_googleads_v8_resources_account_link_proto_goTypes = nil
	file_google_ads_googleads_v8_resources_account_link_proto_depIdxs = nil
}
