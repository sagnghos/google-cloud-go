// Copyright 2024 Google LLC
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
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.3
// source: google/cloud/apihub/v1/provisioning_service.proto

package apihubpb

import (
	context "context"
	reflect "reflect"
	sync "sync"

	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The
// [CreateApiHubInstance][google.cloud.apihub.v1.Provisioning.CreateApiHubInstance]
// method's request.
type CreateApiHubInstanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The parent resource for the Api Hub instance resource.
	// Format: `projects/{project}/locations/{location}`
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Optional. Identifier to assign to the Api Hub instance. Must be unique
	// within scope of the parent resource. If the field is not provided, system
	// generated id will be used.
	//
	// This value should be 4-40 characters, and valid characters
	// are `/[a-z][A-Z][0-9]-_/`.
	ApiHubInstanceId string `protobuf:"bytes,2,opt,name=api_hub_instance_id,json=apiHubInstanceId,proto3" json:"api_hub_instance_id,omitempty"`
	// Required. The ApiHub instance.
	ApiHubInstance *ApiHubInstance `protobuf:"bytes,3,opt,name=api_hub_instance,json=apiHubInstance,proto3" json:"api_hub_instance,omitempty"`
}

func (x *CreateApiHubInstanceRequest) Reset() {
	*x = CreateApiHubInstanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_apihub_v1_provisioning_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateApiHubInstanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateApiHubInstanceRequest) ProtoMessage() {}

func (x *CreateApiHubInstanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_apihub_v1_provisioning_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateApiHubInstanceRequest.ProtoReflect.Descriptor instead.
func (*CreateApiHubInstanceRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_apihub_v1_provisioning_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateApiHubInstanceRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreateApiHubInstanceRequest) GetApiHubInstanceId() string {
	if x != nil {
		return x.ApiHubInstanceId
	}
	return ""
}

func (x *CreateApiHubInstanceRequest) GetApiHubInstance() *ApiHubInstance {
	if x != nil {
		return x.ApiHubInstance
	}
	return nil
}

// The
// [GetApiHubInstance][google.cloud.apihub.v1.Provisioning.GetApiHubInstance]
// method's request.
type GetApiHubInstanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the Api Hub instance to retrieve.
	// Format:
	// `projects/{project}/locations/{location}/apiHubInstances/{apiHubInstance}`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetApiHubInstanceRequest) Reset() {
	*x = GetApiHubInstanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_apihub_v1_provisioning_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetApiHubInstanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetApiHubInstanceRequest) ProtoMessage() {}

func (x *GetApiHubInstanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_apihub_v1_provisioning_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetApiHubInstanceRequest.ProtoReflect.Descriptor instead.
func (*GetApiHubInstanceRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_apihub_v1_provisioning_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetApiHubInstanceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The
// [LookupApiHubInstance][google.cloud.apihub.v1.Provisioning.LookupApiHubInstance]
// method's request.
type LookupApiHubInstanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. There will always be only one Api Hub instance for a GCP project
	// across all locations.
	// The parent resource for the Api Hub instance resource.
	// Format: `projects/{project}/locations/{location}`
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
}

func (x *LookupApiHubInstanceRequest) Reset() {
	*x = LookupApiHubInstanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_apihub_v1_provisioning_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LookupApiHubInstanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupApiHubInstanceRequest) ProtoMessage() {}

func (x *LookupApiHubInstanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_apihub_v1_provisioning_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupApiHubInstanceRequest.ProtoReflect.Descriptor instead.
func (*LookupApiHubInstanceRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_apihub_v1_provisioning_service_proto_rawDescGZIP(), []int{2}
}

func (x *LookupApiHubInstanceRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

// The
// [LookupApiHubInstance][google.cloud.apihub.v1.Provisioning.LookupApiHubInstance]
// method's response.`
type LookupApiHubInstanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// API Hub instance for a project if it exists, empty otherwise.
	ApiHubInstance *ApiHubInstance `protobuf:"bytes,1,opt,name=api_hub_instance,json=apiHubInstance,proto3" json:"api_hub_instance,omitempty"`
}

func (x *LookupApiHubInstanceResponse) Reset() {
	*x = LookupApiHubInstanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_apihub_v1_provisioning_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LookupApiHubInstanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupApiHubInstanceResponse) ProtoMessage() {}

func (x *LookupApiHubInstanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_apihub_v1_provisioning_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupApiHubInstanceResponse.ProtoReflect.Descriptor instead.
func (*LookupApiHubInstanceResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_apihub_v1_provisioning_service_proto_rawDescGZIP(), []int{3}
}

func (x *LookupApiHubInstanceResponse) GetApiHubInstance() *ApiHubInstance {
	if x != nil {
		return x.ApiHubInstance
	}
	return nil
}

var File_google_cloud_apihub_v1_provisioning_service_proto protoreflect.FileDescriptor

var file_google_cloud_apihub_v1_provisioning_service_proto_rawDesc = []byte{
	0x0a, 0x31, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x70, 0x69, 0x68, 0x75, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x70, 0x69, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69,
	0x68, 0x75, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xeb, 0x01, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x69, 0x48, 0x75, 0x62,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x41, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x29, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x12, 0x32, 0x0a, 0x13, 0x61, 0x70, 0x69, 0x5f, 0x68, 0x75, 0x62, 0x5f, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x01, 0x52, 0x10, 0x61, 0x70, 0x69, 0x48, 0x75, 0x62, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x55, 0x0a, 0x10, 0x61, 0x70, 0x69, 0x5f, 0x68, 0x75,
	0x62, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x70, 0x69, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x69, 0x48, 0x75, 0x62,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0e, 0x61,
	0x70, 0x69, 0x48, 0x75, 0x62, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x5c, 0x0a,
	0x18, 0x47, 0x65, 0x74, 0x41, 0x70, 0x69, 0x48, 0x75, 0x62, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2c, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x26, 0x0a,
	0x24, 0x61, 0x70, 0x69, 0x68, 0x75, 0x62, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x70, 0x69, 0x48, 0x75, 0x62, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x63, 0x0a, 0x1b, 0x4c,
	0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x41, 0x70, 0x69, 0x48, 0x75, 0x62, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x06, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2c, 0xe0, 0x41, 0x02, 0xfa,
	0x41, 0x26, 0x12, 0x24, 0x61, 0x70, 0x69, 0x68, 0x75, 0x62, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x70, 0x69, 0x48, 0x75, 0x62,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x22, 0x70, 0x0a, 0x1c, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x41, 0x70, 0x69, 0x48, 0x75, 0x62,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x50, 0x0a, 0x10, 0x61, 0x70, 0x69, 0x5f, 0x68, 0x75, 0x62, 0x5f, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x68, 0x75, 0x62,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x69, 0x48, 0x75, 0x62, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x0e, 0x61, 0x70, 0x69, 0x48, 0x75, 0x62, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x32, 0xef, 0x05, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x69, 0x6e, 0x67, 0x12, 0x8e, 0x02, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70,
	0x69, 0x48, 0x75, 0x62, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x33, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x68,
	0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x69, 0x48,
	0x75, 0x62, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6c, 0x6f, 0x6e, 0x67, 0x72,
	0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0xa1, 0x01, 0xca, 0x41, 0x23, 0x0a, 0x0e, 0x41, 0x70, 0x69, 0x48, 0x75, 0x62, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x11, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xda, 0x41, 0x2b, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x2c, 0x61, 0x70, 0x69, 0x5f, 0x68, 0x75, 0x62, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x2c, 0x61, 0x70, 0x69, 0x5f, 0x68, 0x75, 0x62, 0x5f, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x47, 0x3a, 0x10, 0x61,
	0x70, 0x69, 0x5f, 0x68, 0x75, 0x62, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22,
	0x33, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x61, 0x70, 0x69, 0x48, 0x75, 0x62, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x73, 0x12, 0xb1, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x70, 0x69, 0x48,
	0x75, 0x62, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x30, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x68, 0x75, 0x62,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x69, 0x48, 0x75, 0x62, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x68,
	0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x69, 0x48, 0x75, 0x62, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x22, 0x42, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x35, 0x12, 0x33, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x61, 0x70, 0x69, 0x48, 0x75, 0x62, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0xce, 0x01, 0x0a, 0x14, 0x4c, 0x6f, 0x6f,
	0x6b, 0x75, 0x70, 0x41, 0x70, 0x69, 0x48, 0x75, 0x62, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x12, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x70, 0x69, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75,
	0x70, 0x41, 0x70, 0x69, 0x48, 0x75, 0x62, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x41, 0x70, 0x69, 0x48, 0x75, 0x62, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4b, 0xda, 0x41,
	0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3c, 0x12, 0x3a, 0x2f,
	0x76, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x2a, 0x7d, 0x2f, 0x61, 0x70, 0x69, 0x48, 0x75, 0x62, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x73, 0x3a, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x1a, 0x49, 0xca, 0x41, 0x15, 0x61, 0x70,
	0x69, 0x68, 0x75, 0x62, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77,
	0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x42, 0xba, 0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x68, 0x75, 0x62,
	0x2e, 0x76, 0x31, 0x42, 0x18, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x68, 0x75, 0x62, 0x2f, 0x61, 0x70, 0x69, 0x76,
	0x31, 0x2f, 0x61, 0x70, 0x69, 0x68, 0x75, 0x62, 0x70, 0x62, 0x3b, 0x61, 0x70, 0x69, 0x68, 0x75,
	0x62, 0x70, 0x62, 0xaa, 0x02, 0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x41, 0x70, 0x69, 0x48, 0x75, 0x62, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x16, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x70, 0x69, 0x48,
	0x75, 0x62, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x48, 0x75, 0x62, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_apihub_v1_provisioning_service_proto_rawDescOnce sync.Once
	file_google_cloud_apihub_v1_provisioning_service_proto_rawDescData = file_google_cloud_apihub_v1_provisioning_service_proto_rawDesc
)

func file_google_cloud_apihub_v1_provisioning_service_proto_rawDescGZIP() []byte {
	file_google_cloud_apihub_v1_provisioning_service_proto_rawDescOnce.Do(func() {
		file_google_cloud_apihub_v1_provisioning_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_apihub_v1_provisioning_service_proto_rawDescData)
	})
	return file_google_cloud_apihub_v1_provisioning_service_proto_rawDescData
}

var file_google_cloud_apihub_v1_provisioning_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_cloud_apihub_v1_provisioning_service_proto_goTypes = []any{
	(*CreateApiHubInstanceRequest)(nil),  // 0: google.cloud.apihub.v1.CreateApiHubInstanceRequest
	(*GetApiHubInstanceRequest)(nil),     // 1: google.cloud.apihub.v1.GetApiHubInstanceRequest
	(*LookupApiHubInstanceRequest)(nil),  // 2: google.cloud.apihub.v1.LookupApiHubInstanceRequest
	(*LookupApiHubInstanceResponse)(nil), // 3: google.cloud.apihub.v1.LookupApiHubInstanceResponse
	(*ApiHubInstance)(nil),               // 4: google.cloud.apihub.v1.ApiHubInstance
	(*longrunningpb.Operation)(nil),      // 5: google.longrunning.Operation
}
var file_google_cloud_apihub_v1_provisioning_service_proto_depIdxs = []int32{
	4, // 0: google.cloud.apihub.v1.CreateApiHubInstanceRequest.api_hub_instance:type_name -> google.cloud.apihub.v1.ApiHubInstance
	4, // 1: google.cloud.apihub.v1.LookupApiHubInstanceResponse.api_hub_instance:type_name -> google.cloud.apihub.v1.ApiHubInstance
	0, // 2: google.cloud.apihub.v1.Provisioning.CreateApiHubInstance:input_type -> google.cloud.apihub.v1.CreateApiHubInstanceRequest
	1, // 3: google.cloud.apihub.v1.Provisioning.GetApiHubInstance:input_type -> google.cloud.apihub.v1.GetApiHubInstanceRequest
	2, // 4: google.cloud.apihub.v1.Provisioning.LookupApiHubInstance:input_type -> google.cloud.apihub.v1.LookupApiHubInstanceRequest
	5, // 5: google.cloud.apihub.v1.Provisioning.CreateApiHubInstance:output_type -> google.longrunning.Operation
	4, // 6: google.cloud.apihub.v1.Provisioning.GetApiHubInstance:output_type -> google.cloud.apihub.v1.ApiHubInstance
	3, // 7: google.cloud.apihub.v1.Provisioning.LookupApiHubInstance:output_type -> google.cloud.apihub.v1.LookupApiHubInstanceResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_cloud_apihub_v1_provisioning_service_proto_init() }
func file_google_cloud_apihub_v1_provisioning_service_proto_init() {
	if File_google_cloud_apihub_v1_provisioning_service_proto != nil {
		return
	}
	file_google_cloud_apihub_v1_common_fields_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_apihub_v1_provisioning_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateApiHubInstanceRequest); i {
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
		file_google_cloud_apihub_v1_provisioning_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetApiHubInstanceRequest); i {
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
		file_google_cloud_apihub_v1_provisioning_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*LookupApiHubInstanceRequest); i {
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
		file_google_cloud_apihub_v1_provisioning_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*LookupApiHubInstanceResponse); i {
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
			RawDescriptor: file_google_cloud_apihub_v1_provisioning_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_apihub_v1_provisioning_service_proto_goTypes,
		DependencyIndexes: file_google_cloud_apihub_v1_provisioning_service_proto_depIdxs,
		MessageInfos:      file_google_cloud_apihub_v1_provisioning_service_proto_msgTypes,
	}.Build()
	File_google_cloud_apihub_v1_provisioning_service_proto = out.File
	file_google_cloud_apihub_v1_provisioning_service_proto_rawDesc = nil
	file_google_cloud_apihub_v1_provisioning_service_proto_goTypes = nil
	file_google_cloud_apihub_v1_provisioning_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProvisioningClient is the client API for Provisioning service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProvisioningClient interface {
	// Provisions instance resources for the API Hub.
	CreateApiHubInstance(ctx context.Context, in *CreateApiHubInstanceRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Gets details of a single API Hub instance.
	GetApiHubInstance(ctx context.Context, in *GetApiHubInstanceRequest, opts ...grpc.CallOption) (*ApiHubInstance, error)
	// Looks up an Api Hub instance in a given GCP project. There will always be
	// only one Api Hub instance for a GCP project across all locations.
	LookupApiHubInstance(ctx context.Context, in *LookupApiHubInstanceRequest, opts ...grpc.CallOption) (*LookupApiHubInstanceResponse, error)
}

type provisioningClient struct {
	cc grpc.ClientConnInterface
}

func NewProvisioningClient(cc grpc.ClientConnInterface) ProvisioningClient {
	return &provisioningClient{cc}
}

func (c *provisioningClient) CreateApiHubInstance(ctx context.Context, in *CreateApiHubInstanceRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/google.cloud.apihub.v1.Provisioning/CreateApiHubInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *provisioningClient) GetApiHubInstance(ctx context.Context, in *GetApiHubInstanceRequest, opts ...grpc.CallOption) (*ApiHubInstance, error) {
	out := new(ApiHubInstance)
	err := c.cc.Invoke(ctx, "/google.cloud.apihub.v1.Provisioning/GetApiHubInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *provisioningClient) LookupApiHubInstance(ctx context.Context, in *LookupApiHubInstanceRequest, opts ...grpc.CallOption) (*LookupApiHubInstanceResponse, error) {
	out := new(LookupApiHubInstanceResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.apihub.v1.Provisioning/LookupApiHubInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProvisioningServer is the server API for Provisioning service.
type ProvisioningServer interface {
	// Provisions instance resources for the API Hub.
	CreateApiHubInstance(context.Context, *CreateApiHubInstanceRequest) (*longrunningpb.Operation, error)
	// Gets details of a single API Hub instance.
	GetApiHubInstance(context.Context, *GetApiHubInstanceRequest) (*ApiHubInstance, error)
	// Looks up an Api Hub instance in a given GCP project. There will always be
	// only one Api Hub instance for a GCP project across all locations.
	LookupApiHubInstance(context.Context, *LookupApiHubInstanceRequest) (*LookupApiHubInstanceResponse, error)
}

// UnimplementedProvisioningServer can be embedded to have forward compatible implementations.
type UnimplementedProvisioningServer struct {
}

func (*UnimplementedProvisioningServer) CreateApiHubInstance(context.Context, *CreateApiHubInstanceRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateApiHubInstance not implemented")
}
func (*UnimplementedProvisioningServer) GetApiHubInstance(context.Context, *GetApiHubInstanceRequest) (*ApiHubInstance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApiHubInstance not implemented")
}
func (*UnimplementedProvisioningServer) LookupApiHubInstance(context.Context, *LookupApiHubInstanceRequest) (*LookupApiHubInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookupApiHubInstance not implemented")
}

func RegisterProvisioningServer(s *grpc.Server, srv ProvisioningServer) {
	s.RegisterService(&_Provisioning_serviceDesc, srv)
}

func _Provisioning_CreateApiHubInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateApiHubInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvisioningServer).CreateApiHubInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.apihub.v1.Provisioning/CreateApiHubInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvisioningServer).CreateApiHubInstance(ctx, req.(*CreateApiHubInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provisioning_GetApiHubInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApiHubInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvisioningServer).GetApiHubInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.apihub.v1.Provisioning/GetApiHubInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvisioningServer).GetApiHubInstance(ctx, req.(*GetApiHubInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provisioning_LookupApiHubInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupApiHubInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvisioningServer).LookupApiHubInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.apihub.v1.Provisioning/LookupApiHubInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvisioningServer).LookupApiHubInstance(ctx, req.(*LookupApiHubInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Provisioning_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.apihub.v1.Provisioning",
	HandlerType: (*ProvisioningServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateApiHubInstance",
			Handler:    _Provisioning_CreateApiHubInstance_Handler,
		},
		{
			MethodName: "GetApiHubInstance",
			Handler:    _Provisioning_GetApiHubInstance_Handler,
		},
		{
			MethodName: "LookupApiHubInstance",
			Handler:    _Provisioning_LookupApiHubInstance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/apihub/v1/provisioning_service.proto",
}