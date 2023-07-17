//
//Copyright 2023 Koor Technologies, Inc. All rights reserved.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: api/v1/version_service.proto

package apiv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// ProductVersions describes a map of products with version strings
type ProductVersions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kube         string `protobuf:"bytes,1,opt,name=kube,proto3" json:"kube,omitempty"`
	KoorOperator string `protobuf:"bytes,2,opt,name=koor_operator,json=koorOperator,proto3" json:"koor_operator,omitempty"`
	Ksd          string `protobuf:"bytes,3,opt,name=ksd,proto3" json:"ksd,omitempty"`
	Ceph         string `protobuf:"bytes,4,opt,name=ceph,proto3" json:"ceph,omitempty"`
}

func (x *ProductVersions) Reset() {
	*x = ProductVersions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_version_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductVersions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductVersions) ProtoMessage() {}

func (x *ProductVersions) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_version_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductVersions.ProtoReflect.Descriptor instead.
func (*ProductVersions) Descriptor() ([]byte, []int) {
	return file_api_v1_version_service_proto_rawDescGZIP(), []int{0}
}

func (x *ProductVersions) GetKube() string {
	if x != nil {
		return x.Kube
	}
	return ""
}

func (x *ProductVersions) GetKoorOperator() string {
	if x != nil {
		return x.KoorOperator
	}
	return ""
}

func (x *ProductVersions) GetKsd() string {
	if x != nil {
		return x.Ksd
	}
	return ""
}

func (x *ProductVersions) GetCeph() string {
	if x != nil {
		return x.Ceph
	}
	return ""
}

type OperatorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Versions *ProductVersions `protobuf:"bytes,1,opt,name=versions,proto3" json:"versions,omitempty"`
}

func (x *OperatorRequest) Reset() {
	*x = OperatorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_version_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperatorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperatorRequest) ProtoMessage() {}

func (x *OperatorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_version_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperatorRequest.ProtoReflect.Descriptor instead.
func (*OperatorRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_version_service_proto_rawDescGZIP(), []int{1}
}

func (x *OperatorRequest) GetVersions() *ProductVersions {
	if x != nil {
		return x.Versions
	}
	return nil
}

// DetailedVersion defines a version of a product with a container image or helm chart
type DetailedVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version        string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	ImageUri       string `protobuf:"bytes,2,opt,name=image_uri,json=imageUri,proto3" json:"image_uri,omitempty"`
	ImageHash      string `protobuf:"bytes,3,opt,name=image_hash,json=imageHash,proto3" json:"image_hash,omitempty"`
	HelmRepository string `protobuf:"bytes,4,opt,name=helm_repository,json=helmRepository,proto3" json:"helm_repository,omitempty"`
	HelmChart      string `protobuf:"bytes,5,opt,name=helm_chart,json=helmChart,proto3" json:"helm_chart,omitempty"`
}

func (x *DetailedVersion) Reset() {
	*x = DetailedVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_version_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailedVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailedVersion) ProtoMessage() {}

func (x *DetailedVersion) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_version_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailedVersion.ProtoReflect.Descriptor instead.
func (*DetailedVersion) Descriptor() ([]byte, []int) {
	return file_api_v1_version_service_proto_rawDescGZIP(), []int{2}
}

func (x *DetailedVersion) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *DetailedVersion) GetImageUri() string {
	if x != nil {
		return x.ImageUri
	}
	return ""
}

func (x *DetailedVersion) GetImageHash() string {
	if x != nil {
		return x.ImageHash
	}
	return ""
}

func (x *DetailedVersion) GetHelmRepository() string {
	if x != nil {
		return x.HelmRepository
	}
	return ""
}

func (x *DetailedVersion) GetHelmChart() string {
	if x != nil {
		return x.HelmChart
	}
	return ""
}

// DetailedProductVersions describes a map of products with images or helm charts
type DetailedProductVersions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KoorOperator *DetailedVersion `protobuf:"bytes,1,opt,name=koor_operator,json=koorOperator,proto3" json:"koor_operator,omitempty"`
	Ksd          *DetailedVersion `protobuf:"bytes,2,opt,name=ksd,proto3" json:"ksd,omitempty"`
	Ceph         *DetailedVersion `protobuf:"bytes,3,opt,name=ceph,proto3" json:"ceph,omitempty"`
}

func (x *DetailedProductVersions) Reset() {
	*x = DetailedProductVersions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_version_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailedProductVersions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailedProductVersions) ProtoMessage() {}

func (x *DetailedProductVersions) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_version_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailedProductVersions.ProtoReflect.Descriptor instead.
func (*DetailedProductVersions) Descriptor() ([]byte, []int) {
	return file_api_v1_version_service_proto_rawDescGZIP(), []int{3}
}

func (x *DetailedProductVersions) GetKoorOperator() *DetailedVersion {
	if x != nil {
		return x.KoorOperator
	}
	return nil
}

func (x *DetailedProductVersions) GetKsd() *DetailedVersion {
	if x != nil {
		return x.Ksd
	}
	return nil
}

func (x *DetailedProductVersions) GetCeph() *DetailedVersion {
	if x != nil {
		return x.Ceph
	}
	return nil
}

type OperatorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Versions *DetailedProductVersions `protobuf:"bytes,1,opt,name=versions,proto3" json:"versions,omitempty"`
}

func (x *OperatorResponse) Reset() {
	*x = OperatorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_version_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperatorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperatorResponse) ProtoMessage() {}

func (x *OperatorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_version_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperatorResponse.ProtoReflect.Descriptor instead.
func (*OperatorResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_version_service_proto_rawDescGZIP(), []int{4}
}

func (x *OperatorResponse) GetVersions() *DetailedProductVersions {
	if x != nil {
		return x.Versions
	}
	return nil
}

// VersionMatrix describes a map of products with all possible images or helm chart versions
type VersionMatrix struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KoorOperator map[string]*DetailedVersion `protobuf:"bytes,1,rep,name=koor_operator,json=koorOperator,proto3" json:"koor_operator,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Ksd          map[string]*DetailedVersion `protobuf:"bytes,2,rep,name=ksd,proto3" json:"ksd,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Ceph         map[string]*DetailedVersion `protobuf:"bytes,3,rep,name=ceph,proto3" json:"ceph,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *VersionMatrix) Reset() {
	*x = VersionMatrix{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_version_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionMatrix) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionMatrix) ProtoMessage() {}

func (x *VersionMatrix) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_version_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionMatrix.ProtoReflect.Descriptor instead.
func (*VersionMatrix) Descriptor() ([]byte, []int) {
	return file_api_v1_version_service_proto_rawDescGZIP(), []int{5}
}

func (x *VersionMatrix) GetKoorOperator() map[string]*DetailedVersion {
	if x != nil {
		return x.KoorOperator
	}
	return nil
}

func (x *VersionMatrix) GetKsd() map[string]*DetailedVersion {
	if x != nil {
		return x.Ksd
	}
	return nil
}

func (x *VersionMatrix) GetCeph() map[string]*DetailedVersion {
	if x != nil {
		return x.Ceph
	}
	return nil
}

var File_api_v1_version_service_proto protoreflect.FileDescriptor

var file_api_v1_version_service_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x90, 0x02, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x3a, 0x0a, 0x04, 0x6b, 0x75, 0x62, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x26, 0xfa, 0x42, 0x23, 0x72, 0x21, 0x32, 0x1f, 0x5e, 0x76, 0x3f, 0x5c, 0x64, 0x2b,
	0x28, 0x3f, 0x3a, 0x5c, 0x2e, 0x5c, 0x64, 0x2b, 0x29, 0x7b, 0x30, 0x2c, 0x32, 0x7d, 0x28, 0x3f,
	0x3a, 0x5c, 0x2d, 0x5c, 0x64, 0x2b, 0x29, 0x3f, 0x24, 0x52, 0x04, 0x6b, 0x75, 0x62, 0x65, 0x12,
	0x4b, 0x0a, 0x0d, 0x6b, 0x6f, 0x6f, 0x72, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x26, 0xfa, 0x42, 0x23, 0x72, 0x21, 0x32, 0x1f, 0x5e,
	0x76, 0x3f, 0x5c, 0x64, 0x2b, 0x28, 0x3f, 0x3a, 0x5c, 0x2e, 0x5c, 0x64, 0x2b, 0x29, 0x7b, 0x30,
	0x2c, 0x32, 0x7d, 0x28, 0x3f, 0x3a, 0x5c, 0x2d, 0x5c, 0x64, 0x2b, 0x29, 0x3f, 0x24, 0x52, 0x0c,
	0x6b, 0x6f, 0x6f, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x38, 0x0a, 0x03,
	0x6b, 0x73, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x26, 0xfa, 0x42, 0x23, 0x72, 0x21,
	0x32, 0x1f, 0x5e, 0x76, 0x3f, 0x5c, 0x64, 0x2b, 0x28, 0x3f, 0x3a, 0x5c, 0x2e, 0x5c, 0x64, 0x2b,
	0x29, 0x7b, 0x30, 0x2c, 0x32, 0x7d, 0x28, 0x3f, 0x3a, 0x5c, 0x2d, 0x5c, 0x64, 0x2b, 0x29, 0x3f,
	0x24, 0x52, 0x03, 0x6b, 0x73, 0x64, 0x12, 0x3a, 0x0a, 0x04, 0x63, 0x65, 0x70, 0x68, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x26, 0xfa, 0x42, 0x23, 0x72, 0x21, 0x32, 0x1f, 0x5e, 0x76, 0x3f,
	0x5c, 0x64, 0x2b, 0x28, 0x3f, 0x3a, 0x5c, 0x2e, 0x5c, 0x64, 0x2b, 0x29, 0x7b, 0x30, 0x2c, 0x32,
	0x7d, 0x28, 0x3f, 0x3a, 0x5c, 0x2d, 0x5c, 0x64, 0x2b, 0x29, 0x3f, 0x24, 0x52, 0x04, 0x63, 0x65,
	0x70, 0x68, 0x22, 0x46, 0x0a, 0x0f, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xeb, 0x01, 0x0a, 0x0f, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x40,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x26, 0xfa, 0x42, 0x23, 0x72, 0x21, 0x32, 0x1f, 0x5e, 0x76, 0x3f, 0x5c, 0x64, 0x2b, 0x28, 0x3f,
	0x3a, 0x5c, 0x2e, 0x5c, 0x64, 0x2b, 0x29, 0x7b, 0x30, 0x2c, 0x32, 0x7d, 0x28, 0x3f, 0x3a, 0x5c,
	0x2d, 0x5c, 0x64, 0x2b, 0x29, 0x3f, 0x24, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x25, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x88, 0x01, 0x01, 0x52, 0x08, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x69, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x31, 0x0a, 0x0f, 0x68, 0x65, 0x6c, 0x6d, 0x5f, 0x72,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x88, 0x01, 0x01, 0x52, 0x0e, 0x68, 0x65, 0x6c, 0x6d, 0x52,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x65, 0x6c,
	0x6d, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x68,
	0x65, 0x6c, 0x6d, 0x43, 0x68, 0x61, 0x72, 0x74, 0x22, 0xaf, 0x01, 0x0a, 0x17, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3c, 0x0a, 0x0d, 0x6b, 0x6f, 0x6f, 0x72, 0x5f, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x6b, 0x6f, 0x6f, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x12, 0x29, 0x0a, 0x03, 0x6b, 0x73, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x65,
	0x64, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x6b, 0x73, 0x64, 0x12, 0x2b, 0x0a,
	0x04, 0x63, 0x65, 0x70, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x63, 0x65, 0x70, 0x68, 0x22, 0x4f, 0x0a, 0x10, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b,
	0x0a, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x65, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xc1, 0x03, 0x0a, 0x0d,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x12, 0x4c, 0x0a,
	0x0d, 0x6b, 0x6f, 0x6f, 0x72, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x2e, 0x4b, 0x6f, 0x6f, 0x72,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x6b,
	0x6f, 0x6f, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x30, 0x0a, 0x03, 0x6b,
	0x73, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x2e,
	0x4b, 0x73, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x6b, 0x73, 0x64, 0x12, 0x33, 0x0a,
	0x04, 0x63, 0x65, 0x70, 0x68, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x74, 0x72,
	0x69, 0x78, 0x2e, 0x43, 0x65, 0x70, 0x68, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x63, 0x65,
	0x70, 0x68, 0x1a, 0x58, 0x0a, 0x11, 0x4b, 0x6f, 0x6f, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4f, 0x0a, 0x08,
	0x4b, 0x73, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x50, 0x0a,
	0x09, 0x43, 0x65, 0x70, 0x68, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32,
	0x51, 0x0a, 0x0e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3f, 0x0a, 0x08, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x17, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x8d, 0x01, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x42, 0x13, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x6f, 0x72, 0x2d, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x69, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x58,
	0x58, 0xaa, 0x02, 0x06, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x06, 0x41, 0x70, 0x69,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x12, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x07, 0x41, 0x70, 0x69, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_version_service_proto_rawDescOnce sync.Once
	file_api_v1_version_service_proto_rawDescData = file_api_v1_version_service_proto_rawDesc
)

func file_api_v1_version_service_proto_rawDescGZIP() []byte {
	file_api_v1_version_service_proto_rawDescOnce.Do(func() {
		file_api_v1_version_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_version_service_proto_rawDescData)
	})
	return file_api_v1_version_service_proto_rawDescData
}

var file_api_v1_version_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_v1_version_service_proto_goTypes = []interface{}{
	(*ProductVersions)(nil),         // 0: api.v1.ProductVersions
	(*OperatorRequest)(nil),         // 1: api.v1.OperatorRequest
	(*DetailedVersion)(nil),         // 2: api.v1.DetailedVersion
	(*DetailedProductVersions)(nil), // 3: api.v1.DetailedProductVersions
	(*OperatorResponse)(nil),        // 4: api.v1.OperatorResponse
	(*VersionMatrix)(nil),           // 5: api.v1.VersionMatrix
	nil,                             // 6: api.v1.VersionMatrix.KoorOperatorEntry
	nil,                             // 7: api.v1.VersionMatrix.KsdEntry
	nil,                             // 8: api.v1.VersionMatrix.CephEntry
}
var file_api_v1_version_service_proto_depIdxs = []int32{
	0,  // 0: api.v1.OperatorRequest.versions:type_name -> api.v1.ProductVersions
	2,  // 1: api.v1.DetailedProductVersions.koor_operator:type_name -> api.v1.DetailedVersion
	2,  // 2: api.v1.DetailedProductVersions.ksd:type_name -> api.v1.DetailedVersion
	2,  // 3: api.v1.DetailedProductVersions.ceph:type_name -> api.v1.DetailedVersion
	3,  // 4: api.v1.OperatorResponse.versions:type_name -> api.v1.DetailedProductVersions
	6,  // 5: api.v1.VersionMatrix.koor_operator:type_name -> api.v1.VersionMatrix.KoorOperatorEntry
	7,  // 6: api.v1.VersionMatrix.ksd:type_name -> api.v1.VersionMatrix.KsdEntry
	8,  // 7: api.v1.VersionMatrix.ceph:type_name -> api.v1.VersionMatrix.CephEntry
	2,  // 8: api.v1.VersionMatrix.KoorOperatorEntry.value:type_name -> api.v1.DetailedVersion
	2,  // 9: api.v1.VersionMatrix.KsdEntry.value:type_name -> api.v1.DetailedVersion
	2,  // 10: api.v1.VersionMatrix.CephEntry.value:type_name -> api.v1.DetailedVersion
	1,  // 11: api.v1.VersionService.Operator:input_type -> api.v1.OperatorRequest
	4,  // 12: api.v1.VersionService.Operator:output_type -> api.v1.OperatorResponse
	12, // [12:13] is the sub-list for method output_type
	11, // [11:12] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_api_v1_version_service_proto_init() }
func file_api_v1_version_service_proto_init() {
	if File_api_v1_version_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_version_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductVersions); i {
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
		file_api_v1_version_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperatorRequest); i {
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
		file_api_v1_version_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailedVersion); i {
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
		file_api_v1_version_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailedProductVersions); i {
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
		file_api_v1_version_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperatorResponse); i {
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
		file_api_v1_version_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionMatrix); i {
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
			RawDescriptor: file_api_v1_version_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_version_service_proto_goTypes,
		DependencyIndexes: file_api_v1_version_service_proto_depIdxs,
		MessageInfos:      file_api_v1_version_service_proto_msgTypes,
	}.Build()
	File_api_v1_version_service_proto = out.File
	file_api_v1_version_service_proto_rawDesc = nil
	file_api_v1_version_service_proto_goTypes = nil
	file_api_v1_version_service_proto_depIdxs = nil
}
