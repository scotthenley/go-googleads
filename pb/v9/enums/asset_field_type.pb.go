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
// source: google/ads/googleads/v9/enums/asset_field_type.proto

package enums

import (
	proto "github.com/golang/protobuf/proto"
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

// Enum describing the possible placements of an asset.
type AssetFieldTypeEnum_AssetFieldType int32

const (
	// Not specified.
	AssetFieldTypeEnum_UNSPECIFIED AssetFieldTypeEnum_AssetFieldType = 0
	// Used for return value only. Represents value unknown in this version.
	AssetFieldTypeEnum_UNKNOWN AssetFieldTypeEnum_AssetFieldType = 1
	// The asset is linked for use as a headline.
	AssetFieldTypeEnum_HEADLINE AssetFieldTypeEnum_AssetFieldType = 2
	// The asset is linked for use as a description.
	AssetFieldTypeEnum_DESCRIPTION AssetFieldTypeEnum_AssetFieldType = 3
	// The asset is linked for use as mandatory ad text.
	AssetFieldTypeEnum_MANDATORY_AD_TEXT AssetFieldTypeEnum_AssetFieldType = 4
	// The asset is linked for use as a marketing image.
	AssetFieldTypeEnum_MARKETING_IMAGE AssetFieldTypeEnum_AssetFieldType = 5
	// The asset is linked for use as a media bundle.
	AssetFieldTypeEnum_MEDIA_BUNDLE AssetFieldTypeEnum_AssetFieldType = 6
	// The asset is linked for use as a YouTube video.
	AssetFieldTypeEnum_YOUTUBE_VIDEO AssetFieldTypeEnum_AssetFieldType = 7
	// The asset is linked to indicate that a hotels campaign is "Book on
	// Google" enabled.
	AssetFieldTypeEnum_BOOK_ON_GOOGLE AssetFieldTypeEnum_AssetFieldType = 8
	// The asset is linked for use as a Lead Form extension.
	AssetFieldTypeEnum_LEAD_FORM AssetFieldTypeEnum_AssetFieldType = 9
	// The asset is linked for use as a Promotion extension.
	AssetFieldTypeEnum_PROMOTION AssetFieldTypeEnum_AssetFieldType = 10
	// The asset is linked for use as a Callout extension.
	AssetFieldTypeEnum_CALLOUT AssetFieldTypeEnum_AssetFieldType = 11
	// The asset is linked for use as a Structured Snippet extension.
	AssetFieldTypeEnum_STRUCTURED_SNIPPET AssetFieldTypeEnum_AssetFieldType = 12
	// The asset is linked for use as a Sitelink extension.
	AssetFieldTypeEnum_SITELINK AssetFieldTypeEnum_AssetFieldType = 13
	// The asset is linked for use as a Mobile App extension.
	AssetFieldTypeEnum_MOBILE_APP AssetFieldTypeEnum_AssetFieldType = 14
	// The asset is linked for use as a Hotel Callout extension.
	AssetFieldTypeEnum_HOTEL_CALLOUT AssetFieldTypeEnum_AssetFieldType = 15
	// The asset is linked for use as a Call extension.
	AssetFieldTypeEnum_CALL AssetFieldTypeEnum_AssetFieldType = 16
	// The asset is linked for use as a Price extension.
	AssetFieldTypeEnum_PRICE AssetFieldTypeEnum_AssetFieldType = 24
	// The asset is linked for use as a long headline.
	AssetFieldTypeEnum_LONG_HEADLINE AssetFieldTypeEnum_AssetFieldType = 17
	// The asset is linked for use as a business name.
	AssetFieldTypeEnum_BUSINESS_NAME AssetFieldTypeEnum_AssetFieldType = 18
	// The asset is linked for use as a square marketing image.
	AssetFieldTypeEnum_SQUARE_MARKETING_IMAGE AssetFieldTypeEnum_AssetFieldType = 19
	// The asset is linked for use as a portrait marketing image.
	AssetFieldTypeEnum_PORTRAIT_MARKETING_IMAGE AssetFieldTypeEnum_AssetFieldType = 20
	// The asset is linked for use as a logo.
	AssetFieldTypeEnum_LOGO AssetFieldTypeEnum_AssetFieldType = 21
	// The asset is linked for use as a landscape logo.
	AssetFieldTypeEnum_LANDSCAPE_LOGO AssetFieldTypeEnum_AssetFieldType = 22
	// The asset is linked for use as a non YouTube logo.
	AssetFieldTypeEnum_VIDEO AssetFieldTypeEnum_AssetFieldType = 23
	// The asset is linked for use to select a call-to-action.
	AssetFieldTypeEnum_CALL_TO_ACTION_SELECTION AssetFieldTypeEnum_AssetFieldType = 25
)

// Enum value maps for AssetFieldTypeEnum_AssetFieldType.
var (
	AssetFieldTypeEnum_AssetFieldType_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "HEADLINE",
		3:  "DESCRIPTION",
		4:  "MANDATORY_AD_TEXT",
		5:  "MARKETING_IMAGE",
		6:  "MEDIA_BUNDLE",
		7:  "YOUTUBE_VIDEO",
		8:  "BOOK_ON_GOOGLE",
		9:  "LEAD_FORM",
		10: "PROMOTION",
		11: "CALLOUT",
		12: "STRUCTURED_SNIPPET",
		13: "SITELINK",
		14: "MOBILE_APP",
		15: "HOTEL_CALLOUT",
		16: "CALL",
		24: "PRICE",
		17: "LONG_HEADLINE",
		18: "BUSINESS_NAME",
		19: "SQUARE_MARKETING_IMAGE",
		20: "PORTRAIT_MARKETING_IMAGE",
		21: "LOGO",
		22: "LANDSCAPE_LOGO",
		23: "VIDEO",
		25: "CALL_TO_ACTION_SELECTION",
	}
	AssetFieldTypeEnum_AssetFieldType_value = map[string]int32{
		"UNSPECIFIED":              0,
		"UNKNOWN":                  1,
		"HEADLINE":                 2,
		"DESCRIPTION":              3,
		"MANDATORY_AD_TEXT":        4,
		"MARKETING_IMAGE":          5,
		"MEDIA_BUNDLE":             6,
		"YOUTUBE_VIDEO":            7,
		"BOOK_ON_GOOGLE":           8,
		"LEAD_FORM":                9,
		"PROMOTION":                10,
		"CALLOUT":                  11,
		"STRUCTURED_SNIPPET":       12,
		"SITELINK":                 13,
		"MOBILE_APP":               14,
		"HOTEL_CALLOUT":            15,
		"CALL":                     16,
		"PRICE":                    24,
		"LONG_HEADLINE":            17,
		"BUSINESS_NAME":            18,
		"SQUARE_MARKETING_IMAGE":   19,
		"PORTRAIT_MARKETING_IMAGE": 20,
		"LOGO":                     21,
		"LANDSCAPE_LOGO":           22,
		"VIDEO":                    23,
		"CALL_TO_ACTION_SELECTION": 25,
	}
)

func (x AssetFieldTypeEnum_AssetFieldType) Enum() *AssetFieldTypeEnum_AssetFieldType {
	p := new(AssetFieldTypeEnum_AssetFieldType)
	*p = x
	return p
}

func (x AssetFieldTypeEnum_AssetFieldType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AssetFieldTypeEnum_AssetFieldType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v9_enums_asset_field_type_proto_enumTypes[0].Descriptor()
}

func (AssetFieldTypeEnum_AssetFieldType) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v9_enums_asset_field_type_proto_enumTypes[0]
}

func (x AssetFieldTypeEnum_AssetFieldType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AssetFieldTypeEnum_AssetFieldType.Descriptor instead.
func (AssetFieldTypeEnum_AssetFieldType) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v9_enums_asset_field_type_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing the possible placements of an asset.
type AssetFieldTypeEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AssetFieldTypeEnum) Reset() {
	*x = AssetFieldTypeEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v9_enums_asset_field_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetFieldTypeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetFieldTypeEnum) ProtoMessage() {}

func (x *AssetFieldTypeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v9_enums_asset_field_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetFieldTypeEnum.ProtoReflect.Descriptor instead.
func (*AssetFieldTypeEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v9_enums_asset_field_type_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v9_enums_asset_field_type_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v9_enums_asset_field_type_proto_rawDesc = []byte{
	0x0a, 0x34, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xf9, 0x03, 0x0a, 0x12, 0x41, 0x73, 0x73, 0x65, 0x74, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xe2, 0x03, 0x0a, 0x0e, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a,
	0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x48,
	0x45, 0x41, 0x44, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x44, 0x45, 0x53,
	0x43, 0x52, 0x49, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x4d, 0x41,
	0x4e, 0x44, 0x41, 0x54, 0x4f, 0x52, 0x59, 0x5f, 0x41, 0x44, 0x5f, 0x54, 0x45, 0x58, 0x54, 0x10,
	0x04, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x41, 0x52, 0x4b, 0x45, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x49,
	0x4d, 0x41, 0x47, 0x45, 0x10, 0x05, 0x12, 0x10, 0x0a, 0x0c, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x5f,
	0x42, 0x55, 0x4e, 0x44, 0x4c, 0x45, 0x10, 0x06, 0x12, 0x11, 0x0a, 0x0d, 0x59, 0x4f, 0x55, 0x54,
	0x55, 0x42, 0x45, 0x5f, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x10, 0x07, 0x12, 0x12, 0x0a, 0x0e, 0x42,
	0x4f, 0x4f, 0x4b, 0x5f, 0x4f, 0x4e, 0x5f, 0x47, 0x4f, 0x4f, 0x47, 0x4c, 0x45, 0x10, 0x08, 0x12,
	0x0d, 0x0a, 0x09, 0x4c, 0x45, 0x41, 0x44, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x10, 0x09, 0x12, 0x0d,
	0x0a, 0x09, 0x50, 0x52, 0x4f, 0x4d, 0x4f, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0a, 0x12, 0x0b, 0x0a,
	0x07, 0x43, 0x41, 0x4c, 0x4c, 0x4f, 0x55, 0x54, 0x10, 0x0b, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54,
	0x52, 0x55, 0x43, 0x54, 0x55, 0x52, 0x45, 0x44, 0x5f, 0x53, 0x4e, 0x49, 0x50, 0x50, 0x45, 0x54,
	0x10, 0x0c, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x49, 0x54, 0x45, 0x4c, 0x49, 0x4e, 0x4b, 0x10, 0x0d,
	0x12, 0x0e, 0x0a, 0x0a, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x5f, 0x41, 0x50, 0x50, 0x10, 0x0e,
	0x12, 0x11, 0x0a, 0x0d, 0x48, 0x4f, 0x54, 0x45, 0x4c, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x4f, 0x55,
	0x54, 0x10, 0x0f, 0x12, 0x08, 0x0a, 0x04, 0x43, 0x41, 0x4c, 0x4c, 0x10, 0x10, 0x12, 0x09, 0x0a,
	0x05, 0x50, 0x52, 0x49, 0x43, 0x45, 0x10, 0x18, 0x12, 0x11, 0x0a, 0x0d, 0x4c, 0x4f, 0x4e, 0x47,
	0x5f, 0x48, 0x45, 0x41, 0x44, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x11, 0x12, 0x11, 0x0a, 0x0d, 0x42,
	0x55, 0x53, 0x49, 0x4e, 0x45, 0x53, 0x53, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x12, 0x12, 0x1a,
	0x0a, 0x16, 0x53, 0x51, 0x55, 0x41, 0x52, 0x45, 0x5f, 0x4d, 0x41, 0x52, 0x4b, 0x45, 0x54, 0x49,
	0x4e, 0x47, 0x5f, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x10, 0x13, 0x12, 0x1c, 0x0a, 0x18, 0x50, 0x4f,
	0x52, 0x54, 0x52, 0x41, 0x49, 0x54, 0x5f, 0x4d, 0x41, 0x52, 0x4b, 0x45, 0x54, 0x49, 0x4e, 0x47,
	0x5f, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x10, 0x14, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x4f, 0x47, 0x4f,
	0x10, 0x15, 0x12, 0x12, 0x0a, 0x0e, 0x4c, 0x41, 0x4e, 0x44, 0x53, 0x43, 0x41, 0x50, 0x45, 0x5f,
	0x4c, 0x4f, 0x47, 0x4f, 0x10, 0x16, 0x12, 0x09, 0x0a, 0x05, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x10,
	0x17, 0x12, 0x1c, 0x0a, 0x18, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x54, 0x4f, 0x5f, 0x41, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x19, 0x42,
	0xe8, 0x01, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x13, 0x41, 0x73, 0x73, 0x65, 0x74, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2f, 0x76, 0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x39,
	0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c,
	0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x39,
	0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a,
	0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a,
	0x3a, 0x56, 0x39, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_google_ads_googleads_v9_enums_asset_field_type_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v9_enums_asset_field_type_proto_rawDescData = file_google_ads_googleads_v9_enums_asset_field_type_proto_rawDesc
)

func file_google_ads_googleads_v9_enums_asset_field_type_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v9_enums_asset_field_type_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v9_enums_asset_field_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v9_enums_asset_field_type_proto_rawDescData)
	})
	return file_google_ads_googleads_v9_enums_asset_field_type_proto_rawDescData
}

var file_google_ads_googleads_v9_enums_asset_field_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v9_enums_asset_field_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v9_enums_asset_field_type_proto_goTypes = []interface{}{
	(AssetFieldTypeEnum_AssetFieldType)(0), // 0: google.ads.googleads.v9.enums.AssetFieldTypeEnum.AssetFieldType
	(*AssetFieldTypeEnum)(nil),             // 1: google.ads.googleads.v9.enums.AssetFieldTypeEnum
}
var file_google_ads_googleads_v9_enums_asset_field_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v9_enums_asset_field_type_proto_init() }
func file_google_ads_googleads_v9_enums_asset_field_type_proto_init() {
	if File_google_ads_googleads_v9_enums_asset_field_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v9_enums_asset_field_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetFieldTypeEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v9_enums_asset_field_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v9_enums_asset_field_type_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v9_enums_asset_field_type_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v9_enums_asset_field_type_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v9_enums_asset_field_type_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v9_enums_asset_field_type_proto = out.File
	file_google_ads_googleads_v9_enums_asset_field_type_proto_rawDesc = nil
	file_google_ads_googleads_v9_enums_asset_field_type_proto_goTypes = nil
	file_google_ads_googleads_v9_enums_asset_field_type_proto_depIdxs = nil
}
