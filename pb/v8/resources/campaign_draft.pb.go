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
// source: google/ads/googleads/v8/resources/campaign_draft.proto

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

// A campaign draft.
type CampaignDraft struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the campaign draft.
	// Campaign draft resource names have the form:
	//
	// `customers/{customer_id}/campaignDrafts/{base_campaign_id}~{draft_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the draft.
	//
	// This field is read-only.
	DraftId *int64 `protobuf:"varint,9,opt,name=draft_id,json=draftId,proto3,oneof" json:"draft_id,omitempty"`
	// Immutable. The base campaign to which the draft belongs.
	BaseCampaign *string `protobuf:"bytes,10,opt,name=base_campaign,json=baseCampaign,proto3,oneof" json:"base_campaign,omitempty"`
	// The name of the campaign draft.
	//
	// This field is required and should not be empty when creating new
	// campaign drafts.
	//
	// It must not contain any null (code point 0x0), NL line feed
	// (code point 0xA) or carriage return (code point 0xD) characters.
	Name *string `protobuf:"bytes,11,opt,name=name,proto3,oneof" json:"name,omitempty"`
	// Output only. Resource name of the Campaign that results from overlaying the draft
	// changes onto the base campaign.
	//
	// This field is read-only.
	DraftCampaign *string `protobuf:"bytes,12,opt,name=draft_campaign,json=draftCampaign,proto3,oneof" json:"draft_campaign,omitempty"`
	// Output only. The status of the campaign draft. This field is read-only.
	//
	// When a new campaign draft is added, the status defaults to PROPOSED.
	Status enums.CampaignDraftStatusEnum_CampaignDraftStatus `protobuf:"varint,6,opt,name=status,proto3,enum=google.ads.googleads.v8.enums.CampaignDraftStatusEnum_CampaignDraftStatus" json:"status,omitempty"`
	// Output only. Whether there is an experiment based on this draft currently serving.
	HasExperimentRunning *bool `protobuf:"varint,13,opt,name=has_experiment_running,json=hasExperimentRunning,proto3,oneof" json:"has_experiment_running,omitempty"`
	// Output only. The resource name of the long-running operation that can be used to poll
	// for completion of draft promotion. This is only set if the draft promotion
	// is in progress or finished.
	LongRunningOperation *string `protobuf:"bytes,14,opt,name=long_running_operation,json=longRunningOperation,proto3,oneof" json:"long_running_operation,omitempty"`
}

func (x *CampaignDraft) Reset() {
	*x = CampaignDraft{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v8_resources_campaign_draft_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CampaignDraft) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CampaignDraft) ProtoMessage() {}

func (x *CampaignDraft) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v8_resources_campaign_draft_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CampaignDraft.ProtoReflect.Descriptor instead.
func (*CampaignDraft) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v8_resources_campaign_draft_proto_rawDescGZIP(), []int{0}
}

func (x *CampaignDraft) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *CampaignDraft) GetDraftId() int64 {
	if x != nil && x.DraftId != nil {
		return *x.DraftId
	}
	return 0
}

func (x *CampaignDraft) GetBaseCampaign() string {
	if x != nil && x.BaseCampaign != nil {
		return *x.BaseCampaign
	}
	return ""
}

func (x *CampaignDraft) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *CampaignDraft) GetDraftCampaign() string {
	if x != nil && x.DraftCampaign != nil {
		return *x.DraftCampaign
	}
	return ""
}

func (x *CampaignDraft) GetStatus() enums.CampaignDraftStatusEnum_CampaignDraftStatus {
	if x != nil {
		return x.Status
	}
	return enums.CampaignDraftStatusEnum_UNSPECIFIED
}

func (x *CampaignDraft) GetHasExperimentRunning() bool {
	if x != nil && x.HasExperimentRunning != nil {
		return *x.HasExperimentRunning
	}
	return false
}

func (x *CampaignDraft) GetLongRunningOperation() string {
	if x != nil && x.LongRunningOperation != nil {
		return *x.LongRunningOperation
	}
	return ""
}

var File_google_ads_googleads_v8_resources_campaign_draft_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v8_resources_campaign_draft_proto_rawDesc = []byte{
	0x0a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x38, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x64, 0x72, 0x61,
	0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x38, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x39, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2f, 0x76, 0x38, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x63, 0x61, 0x6d, 0x70, 0x61,
	0x69, 0x67, 0x6e, 0x5f, 0x64, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x9b, 0x06, 0x0a, 0x0d, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x44, 0x72, 0x61,
	0x66, 0x74, 0x12, 0x53, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2e, 0xe0, 0x41, 0x05, 0xfa, 0x41,
	0x28, 0x0a, 0x26, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x6d, 0x70,
	0x61, 0x69, 0x67, 0x6e, 0x44, 0x72, 0x61, 0x66, 0x74, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x08, 0x64, 0x72, 0x61, 0x66, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x00,
	0x52, 0x07, 0x64, 0x72, 0x61, 0x66, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x53, 0x0a, 0x0d,
	0x62, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x29, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x48, 0x01,
	0x52, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x88, 0x01,
	0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x55, 0x0a, 0x0e, 0x64, 0x72,
	0x61, 0x66, 0x74, 0x5f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x29, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x48, 0x03, 0x52,
	0x0d, 0x64, 0x72, 0x61, 0x66, 0x74, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x88, 0x01,
	0x01, 0x12, 0x67, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x4a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2e, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x44, 0x72, 0x61, 0x66, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69,
	0x67, 0x6e, 0x44, 0x72, 0x61, 0x66, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3e, 0x0a, 0x16, 0x68, 0x61,
	0x73, 0x5f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x75, 0x6e,
	0x6e, 0x69, 0x6e, 0x67, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48,
	0x04, 0x52, 0x14, 0x68, 0x61, 0x73, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x88, 0x01, 0x01, 0x12, 0x3e, 0x0a, 0x16, 0x6c, 0x6f,
	0x6e, 0x67, 0x5f, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48,
	0x05, 0x52, 0x14, 0x6c, 0x6f, 0x6e, 0x67, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x3a, 0x71, 0xea, 0x41, 0x6e, 0x0a,
	0x26, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69,
	0x67, 0x6e, 0x44, 0x72, 0x61, 0x66, 0x74, 0x12, 0x44, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d,
	0x2f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x44, 0x72, 0x61, 0x66, 0x74, 0x73, 0x2f,
	0x7b, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x69,
	0x64, 0x7d, 0x7e, 0x7b, 0x64, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x0b, 0x0a,
	0x09, 0x5f, 0x64, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x69, 0x64, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x62,
	0x61, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x42, 0x07, 0x0a, 0x05,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x64, 0x72, 0x61, 0x66, 0x74, 0x5f,
	0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x42, 0x19, 0x0a, 0x17, 0x5f, 0x68, 0x61, 0x73,
	0x5f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x75, 0x6e, 0x6e,
	0x69, 0x6e, 0x67, 0x42, 0x19, 0x0a, 0x17, 0x5f, 0x6c, 0x6f, 0x6e, 0x67, 0x5f, 0x72, 0x75, 0x6e,
	0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0xff,
	0x01, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x12, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69,
	0x67, 0x6e, 0x44, 0x72, 0x61, 0x66, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2f, 0x76, 0x38, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41,
	0xaa, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x38, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64,
	0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x38, 0x5c, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x3a, 0x3a, 0x56, 0x38, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v8_resources_campaign_draft_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v8_resources_campaign_draft_proto_rawDescData = file_google_ads_googleads_v8_resources_campaign_draft_proto_rawDesc
)

func file_google_ads_googleads_v8_resources_campaign_draft_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v8_resources_campaign_draft_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v8_resources_campaign_draft_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v8_resources_campaign_draft_proto_rawDescData)
	})
	return file_google_ads_googleads_v8_resources_campaign_draft_proto_rawDescData
}

var file_google_ads_googleads_v8_resources_campaign_draft_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v8_resources_campaign_draft_proto_goTypes = []interface{}{
	(*CampaignDraft)(nil), // 0: google.ads.googleads.v8.resources.CampaignDraft
	(enums.CampaignDraftStatusEnum_CampaignDraftStatus)(0), // 1: google.ads.googleads.v8.enums.CampaignDraftStatusEnum.CampaignDraftStatus
}
var file_google_ads_googleads_v8_resources_campaign_draft_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v8.resources.CampaignDraft.status:type_name -> google.ads.googleads.v8.enums.CampaignDraftStatusEnum.CampaignDraftStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v8_resources_campaign_draft_proto_init() }
func file_google_ads_googleads_v8_resources_campaign_draft_proto_init() {
	if File_google_ads_googleads_v8_resources_campaign_draft_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v8_resources_campaign_draft_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CampaignDraft); i {
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
	file_google_ads_googleads_v8_resources_campaign_draft_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v8_resources_campaign_draft_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v8_resources_campaign_draft_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v8_resources_campaign_draft_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v8_resources_campaign_draft_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v8_resources_campaign_draft_proto = out.File
	file_google_ads_googleads_v8_resources_campaign_draft_proto_rawDesc = nil
	file_google_ads_googleads_v8_resources_campaign_draft_proto_goTypes = nil
	file_google_ads_googleads_v8_resources_campaign_draft_proto_depIdxs = nil
}
