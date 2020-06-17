// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helm.proto

package pb

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Release struct {
	ReleaseName          *wrappers.StringValue `protobuf:"bytes,1,opt,name=release_name,json=releaseName,proto3" json:"release_name,omitempty"`
	Version              *wrappers.Int32Value  `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Namespace            *wrappers.StringValue `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Status               *wrappers.StringValue `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Description          *wrappers.StringValue `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	FirstDeployedTime    *timestamp.Timestamp  `protobuf:"bytes,6,opt,name=first_deployed_time,json=firstDeployedTime,proto3" json:"first_deployed_time,omitempty"`
	LastDeployedTime     *timestamp.Timestamp  `protobuf:"bytes,7,opt,name=last_deployed_time,json=lastDeployedTime,proto3" json:"last_deployed_time,omitempty"`
	DeletedTime          *timestamp.Timestamp  `protobuf:"bytes,8,opt,name=deleted_time,json=deletedTime,proto3" json:"deleted_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Release) Reset()         { *m = Release{} }
func (m *Release) String() string { return proto.CompactTextString(m) }
func (*Release) ProtoMessage()    {}
func (*Release) Descriptor() ([]byte, []int) {
	return fileDescriptor_141ba45979af1a8e, []int{0}
}

func (m *Release) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Release.Unmarshal(m, b)
}
func (m *Release) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Release.Marshal(b, m, deterministic)
}
func (m *Release) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Release.Merge(m, src)
}
func (m *Release) XXX_Size() int {
	return xxx_messageInfo_Release.Size(m)
}
func (m *Release) XXX_DiscardUnknown() {
	xxx_messageInfo_Release.DiscardUnknown(m)
}

var xxx_messageInfo_Release proto.InternalMessageInfo

func (m *Release) GetReleaseName() *wrappers.StringValue {
	if m != nil {
		return m.ReleaseName
	}
	return nil
}

func (m *Release) GetVersion() *wrappers.Int32Value {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *Release) GetNamespace() *wrappers.StringValue {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *Release) GetStatus() *wrappers.StringValue {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Release) GetDescription() *wrappers.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *Release) GetFirstDeployedTime() *timestamp.Timestamp {
	if m != nil {
		return m.FirstDeployedTime
	}
	return nil
}

func (m *Release) GetLastDeployedTime() *timestamp.Timestamp {
	if m != nil {
		return m.LastDeployedTime
	}
	return nil
}

func (m *Release) GetDeletedTime() *timestamp.Timestamp {
	if m != nil {
		return m.DeletedTime
	}
	return nil
}

type ListReleasesRequest struct {
	RuntimeId            *wrappers.StringValue `protobuf:"bytes,1,opt,name=runtime_id,json=runtimeId,proto3" json:"runtime_id,omitempty"`
	ReleaseName          *wrappers.StringValue `protobuf:"bytes,2,opt,name=release_name,json=releaseName,proto3" json:"release_name,omitempty"`
	Namespace            *wrappers.StringValue `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Status               *wrappers.StringValue `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ListReleasesRequest) Reset()         { *m = ListReleasesRequest{} }
func (m *ListReleasesRequest) String() string { return proto.CompactTextString(m) }
func (*ListReleasesRequest) ProtoMessage()    {}
func (*ListReleasesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_141ba45979af1a8e, []int{1}
}

func (m *ListReleasesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListReleasesRequest.Unmarshal(m, b)
}
func (m *ListReleasesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListReleasesRequest.Marshal(b, m, deterministic)
}
func (m *ListReleasesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListReleasesRequest.Merge(m, src)
}
func (m *ListReleasesRequest) XXX_Size() int {
	return xxx_messageInfo_ListReleasesRequest.Size(m)
}
func (m *ListReleasesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListReleasesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListReleasesRequest proto.InternalMessageInfo

func (m *ListReleasesRequest) GetRuntimeId() *wrappers.StringValue {
	if m != nil {
		return m.RuntimeId
	}
	return nil
}

func (m *ListReleasesRequest) GetReleaseName() *wrappers.StringValue {
	if m != nil {
		return m.ReleaseName
	}
	return nil
}

func (m *ListReleasesRequest) GetNamespace() *wrappers.StringValue {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *ListReleasesRequest) GetStatus() *wrappers.StringValue {
	if m != nil {
		return m.Status
	}
	return nil
}

type ListReleaseResponse struct {
	ReleaseSet           []*Release `protobuf:"bytes,1,rep,name=release_set,json=releaseSet,proto3" json:"release_set,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListReleaseResponse) Reset()         { *m = ListReleaseResponse{} }
func (m *ListReleaseResponse) String() string { return proto.CompactTextString(m) }
func (*ListReleaseResponse) ProtoMessage()    {}
func (*ListReleaseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_141ba45979af1a8e, []int{2}
}

func (m *ListReleaseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListReleaseResponse.Unmarshal(m, b)
}
func (m *ListReleaseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListReleaseResponse.Marshal(b, m, deterministic)
}
func (m *ListReleaseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListReleaseResponse.Merge(m, src)
}
func (m *ListReleaseResponse) XXX_Size() int {
	return xxx_messageInfo_ListReleaseResponse.Size(m)
}
func (m *ListReleaseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListReleaseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListReleaseResponse proto.InternalMessageInfo

func (m *ListReleaseResponse) GetReleaseSet() []*Release {
	if m != nil {
		return m.ReleaseSet
	}
	return nil
}

type CreateReleaseRequest struct {
	// required, id of app to run in cluster
	AppId *wrappers.StringValue `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	// required, id of app version
	VersionId *wrappers.StringValue `protobuf:"bytes,2,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
	// required, id of runtime
	RuntimeId *wrappers.StringValue `protobuf:"bytes,3,opt,name=runtime_id,json=runtimeId,proto3" json:"runtime_id,omitempty"`
	// release name
	ReleaseName *wrappers.StringValue `protobuf:"bytes,4,opt,name=release_name,json=releaseName,proto3" json:"release_name,omitempty"`
	// namespace
	Namespace            *wrappers.StringValue `protobuf:"bytes,5,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CreateReleaseRequest) Reset()         { *m = CreateReleaseRequest{} }
func (m *CreateReleaseRequest) String() string { return proto.CompactTextString(m) }
func (*CreateReleaseRequest) ProtoMessage()    {}
func (*CreateReleaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_141ba45979af1a8e, []int{3}
}

func (m *CreateReleaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateReleaseRequest.Unmarshal(m, b)
}
func (m *CreateReleaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateReleaseRequest.Marshal(b, m, deterministic)
}
func (m *CreateReleaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateReleaseRequest.Merge(m, src)
}
func (m *CreateReleaseRequest) XXX_Size() int {
	return xxx_messageInfo_CreateReleaseRequest.Size(m)
}
func (m *CreateReleaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateReleaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateReleaseRequest proto.InternalMessageInfo

func (m *CreateReleaseRequest) GetAppId() *wrappers.StringValue {
	if m != nil {
		return m.AppId
	}
	return nil
}

func (m *CreateReleaseRequest) GetVersionId() *wrappers.StringValue {
	if m != nil {
		return m.VersionId
	}
	return nil
}

func (m *CreateReleaseRequest) GetRuntimeId() *wrappers.StringValue {
	if m != nil {
		return m.RuntimeId
	}
	return nil
}

func (m *CreateReleaseRequest) GetReleaseName() *wrappers.StringValue {
	if m != nil {
		return m.ReleaseName
	}
	return nil
}

func (m *CreateReleaseRequest) GetNamespace() *wrappers.StringValue {
	if m != nil {
		return m.Namespace
	}
	return nil
}

type CreateReleaseResponse struct {
	ReleaseName          *wrappers.StringValue `protobuf:"bytes,1,opt,name=release_name,json=releaseName,proto3" json:"release_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CreateReleaseResponse) Reset()         { *m = CreateReleaseResponse{} }
func (m *CreateReleaseResponse) String() string { return proto.CompactTextString(m) }
func (*CreateReleaseResponse) ProtoMessage()    {}
func (*CreateReleaseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_141ba45979af1a8e, []int{4}
}

func (m *CreateReleaseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateReleaseResponse.Unmarshal(m, b)
}
func (m *CreateReleaseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateReleaseResponse.Marshal(b, m, deterministic)
}
func (m *CreateReleaseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateReleaseResponse.Merge(m, src)
}
func (m *CreateReleaseResponse) XXX_Size() int {
	return xxx_messageInfo_CreateReleaseResponse.Size(m)
}
func (m *CreateReleaseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateReleaseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateReleaseResponse proto.InternalMessageInfo

func (m *CreateReleaseResponse) GetReleaseName() *wrappers.StringValue {
	if m != nil {
		return m.ReleaseName
	}
	return nil
}

type UpgradeReleaseRequest struct {
	RuntimeId            *wrappers.StringValue `protobuf:"bytes,1,opt,name=runtime_id,json=runtimeId,proto3" json:"runtime_id,omitempty"`
	ReleaseName          *wrappers.StringValue `protobuf:"bytes,2,opt,name=release_name,json=releaseName,proto3" json:"release_name,omitempty"`
	VersionId            *wrappers.StringValue `protobuf:"bytes,3,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpgradeReleaseRequest) Reset()         { *m = UpgradeReleaseRequest{} }
func (m *UpgradeReleaseRequest) String() string { return proto.CompactTextString(m) }
func (*UpgradeReleaseRequest) ProtoMessage()    {}
func (*UpgradeReleaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_141ba45979af1a8e, []int{5}
}

func (m *UpgradeReleaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpgradeReleaseRequest.Unmarshal(m, b)
}
func (m *UpgradeReleaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpgradeReleaseRequest.Marshal(b, m, deterministic)
}
func (m *UpgradeReleaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpgradeReleaseRequest.Merge(m, src)
}
func (m *UpgradeReleaseRequest) XXX_Size() int {
	return xxx_messageInfo_UpgradeReleaseRequest.Size(m)
}
func (m *UpgradeReleaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpgradeReleaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpgradeReleaseRequest proto.InternalMessageInfo

func (m *UpgradeReleaseRequest) GetRuntimeId() *wrappers.StringValue {
	if m != nil {
		return m.RuntimeId
	}
	return nil
}

func (m *UpgradeReleaseRequest) GetReleaseName() *wrappers.StringValue {
	if m != nil {
		return m.ReleaseName
	}
	return nil
}

func (m *UpgradeReleaseRequest) GetVersionId() *wrappers.StringValue {
	if m != nil {
		return m.VersionId
	}
	return nil
}

type UpgradeReleaseResponse struct {
	ReleaseName          *wrappers.StringValue `protobuf:"bytes,1,opt,name=release_name,json=releaseName,proto3" json:"release_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpgradeReleaseResponse) Reset()         { *m = UpgradeReleaseResponse{} }
func (m *UpgradeReleaseResponse) String() string { return proto.CompactTextString(m) }
func (*UpgradeReleaseResponse) ProtoMessage()    {}
func (*UpgradeReleaseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_141ba45979af1a8e, []int{6}
}

func (m *UpgradeReleaseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpgradeReleaseResponse.Unmarshal(m, b)
}
func (m *UpgradeReleaseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpgradeReleaseResponse.Marshal(b, m, deterministic)
}
func (m *UpgradeReleaseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpgradeReleaseResponse.Merge(m, src)
}
func (m *UpgradeReleaseResponse) XXX_Size() int {
	return xxx_messageInfo_UpgradeReleaseResponse.Size(m)
}
func (m *UpgradeReleaseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpgradeReleaseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpgradeReleaseResponse proto.InternalMessageInfo

func (m *UpgradeReleaseResponse) GetReleaseName() *wrappers.StringValue {
	if m != nil {
		return m.ReleaseName
	}
	return nil
}

type RollbackReleaseRequest struct {
	RuntimeId            *wrappers.StringValue `protobuf:"bytes,1,opt,name=runtime_id,json=runtimeId,proto3" json:"runtime_id,omitempty"`
	ReleaseName          *wrappers.StringValue `protobuf:"bytes,2,opt,name=release_name,json=releaseName,proto3" json:"release_name,omitempty"`
	Version              *wrappers.Int32Value  `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *RollbackReleaseRequest) Reset()         { *m = RollbackReleaseRequest{} }
func (m *RollbackReleaseRequest) String() string { return proto.CompactTextString(m) }
func (*RollbackReleaseRequest) ProtoMessage()    {}
func (*RollbackReleaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_141ba45979af1a8e, []int{7}
}

func (m *RollbackReleaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RollbackReleaseRequest.Unmarshal(m, b)
}
func (m *RollbackReleaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RollbackReleaseRequest.Marshal(b, m, deterministic)
}
func (m *RollbackReleaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RollbackReleaseRequest.Merge(m, src)
}
func (m *RollbackReleaseRequest) XXX_Size() int {
	return xxx_messageInfo_RollbackReleaseRequest.Size(m)
}
func (m *RollbackReleaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RollbackReleaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RollbackReleaseRequest proto.InternalMessageInfo

func (m *RollbackReleaseRequest) GetRuntimeId() *wrappers.StringValue {
	if m != nil {
		return m.RuntimeId
	}
	return nil
}

func (m *RollbackReleaseRequest) GetReleaseName() *wrappers.StringValue {
	if m != nil {
		return m.ReleaseName
	}
	return nil
}

func (m *RollbackReleaseRequest) GetVersion() *wrappers.Int32Value {
	if m != nil {
		return m.Version
	}
	return nil
}

type RollbackReleaseResponse struct {
	ReleaseName          *wrappers.StringValue `protobuf:"bytes,1,opt,name=release_name,json=releaseName,proto3" json:"release_name,omitempty"`
	Version              *wrappers.Int32Value  `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *RollbackReleaseResponse) Reset()         { *m = RollbackReleaseResponse{} }
func (m *RollbackReleaseResponse) String() string { return proto.CompactTextString(m) }
func (*RollbackReleaseResponse) ProtoMessage()    {}
func (*RollbackReleaseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_141ba45979af1a8e, []int{8}
}

func (m *RollbackReleaseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RollbackReleaseResponse.Unmarshal(m, b)
}
func (m *RollbackReleaseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RollbackReleaseResponse.Marshal(b, m, deterministic)
}
func (m *RollbackReleaseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RollbackReleaseResponse.Merge(m, src)
}
func (m *RollbackReleaseResponse) XXX_Size() int {
	return xxx_messageInfo_RollbackReleaseResponse.Size(m)
}
func (m *RollbackReleaseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RollbackReleaseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RollbackReleaseResponse proto.InternalMessageInfo

func (m *RollbackReleaseResponse) GetReleaseName() *wrappers.StringValue {
	if m != nil {
		return m.ReleaseName
	}
	return nil
}

func (m *RollbackReleaseResponse) GetVersion() *wrappers.Int32Value {
	if m != nil {
		return m.Version
	}
	return nil
}

type DeleteReleaseRequest struct {
	RuntimeId            *wrappers.StringValue `protobuf:"bytes,1,opt,name=runtime_id,json=runtimeId,proto3" json:"runtime_id,omitempty"`
	ReleaseName          *wrappers.StringValue `protobuf:"bytes,2,opt,name=release_name,json=releaseName,proto3" json:"release_name,omitempty"`
	Purge                *wrappers.BoolValue   `protobuf:"bytes,3,opt,name=purge,proto3" json:"purge,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *DeleteReleaseRequest) Reset()         { *m = DeleteReleaseRequest{} }
func (m *DeleteReleaseRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteReleaseRequest) ProtoMessage()    {}
func (*DeleteReleaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_141ba45979af1a8e, []int{9}
}

func (m *DeleteReleaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteReleaseRequest.Unmarshal(m, b)
}
func (m *DeleteReleaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteReleaseRequest.Marshal(b, m, deterministic)
}
func (m *DeleteReleaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteReleaseRequest.Merge(m, src)
}
func (m *DeleteReleaseRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteReleaseRequest.Size(m)
}
func (m *DeleteReleaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteReleaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteReleaseRequest proto.InternalMessageInfo

func (m *DeleteReleaseRequest) GetRuntimeId() *wrappers.StringValue {
	if m != nil {
		return m.RuntimeId
	}
	return nil
}

func (m *DeleteReleaseRequest) GetReleaseName() *wrappers.StringValue {
	if m != nil {
		return m.ReleaseName
	}
	return nil
}

func (m *DeleteReleaseRequest) GetPurge() *wrappers.BoolValue {
	if m != nil {
		return m.Purge
	}
	return nil
}

type DeleteReleaseResponse struct {
	ReleaseName          *wrappers.StringValue `protobuf:"bytes,1,opt,name=release_name,json=releaseName,proto3" json:"release_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *DeleteReleaseResponse) Reset()         { *m = DeleteReleaseResponse{} }
func (m *DeleteReleaseResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteReleaseResponse) ProtoMessage()    {}
func (*DeleteReleaseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_141ba45979af1a8e, []int{10}
}

func (m *DeleteReleaseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteReleaseResponse.Unmarshal(m, b)
}
func (m *DeleteReleaseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteReleaseResponse.Marshal(b, m, deterministic)
}
func (m *DeleteReleaseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteReleaseResponse.Merge(m, src)
}
func (m *DeleteReleaseResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteReleaseResponse.Size(m)
}
func (m *DeleteReleaseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteReleaseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteReleaseResponse proto.InternalMessageInfo

func (m *DeleteReleaseResponse) GetReleaseName() *wrappers.StringValue {
	if m != nil {
		return m.ReleaseName
	}
	return nil
}

func init() {
	proto.RegisterType((*Release)(nil), "openpitrix.Release")
	proto.RegisterType((*ListReleasesRequest)(nil), "openpitrix.ListReleasesRequest")
	proto.RegisterType((*ListReleaseResponse)(nil), "openpitrix.ListReleaseResponse")
	proto.RegisterType((*CreateReleaseRequest)(nil), "openpitrix.CreateReleaseRequest")
	proto.RegisterType((*CreateReleaseResponse)(nil), "openpitrix.CreateReleaseResponse")
	proto.RegisterType((*UpgradeReleaseRequest)(nil), "openpitrix.UpgradeReleaseRequest")
	proto.RegisterType((*UpgradeReleaseResponse)(nil), "openpitrix.UpgradeReleaseResponse")
	proto.RegisterType((*RollbackReleaseRequest)(nil), "openpitrix.RollbackReleaseRequest")
	proto.RegisterType((*RollbackReleaseResponse)(nil), "openpitrix.RollbackReleaseResponse")
	proto.RegisterType((*DeleteReleaseRequest)(nil), "openpitrix.DeleteReleaseRequest")
	proto.RegisterType((*DeleteReleaseResponse)(nil), "openpitrix.DeleteReleaseResponse")
}

func init() { proto.RegisterFile("helm.proto", fileDescriptor_141ba45979af1a8e) }

var fileDescriptor_141ba45979af1a8e = []byte{
	// 773 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0xf3, 0x57, 0x3a, 0x09, 0x69, 0xbb, 0xfd, 0x21, 0x72, 0x2b, 0x1a, 0x8c, 0x90, 0xaa,
	0xd2, 0x26, 0x90, 0xc2, 0xa5, 0x08, 0x50, 0x4a, 0x0f, 0x94, 0xbf, 0x43, 0x0a, 0x08, 0xb8, 0x44,
	0x9b, 0x78, 0x6b, 0x2c, 0x1c, 0x7b, 0xd9, 0xdd, 0xb4, 0xf4, 0x86, 0x10, 0x48, 0xbd, 0x12, 0x9e,
	0x80, 0x57, 0xe1, 0xc0, 0x8d, 0x0b, 0xbc, 0x02, 0x0f, 0x82, 0x6c, 0xaf, 0x69, 0xd6, 0x0e, 0x6d,
	0x4a, 0x2b, 0xe8, 0x29, 0x91, 0xe7, 0xfb, 0xbe, 0x99, 0x9d, 0x6f, 0x76, 0x16, 0xe0, 0x25, 0x71,
	0x3a, 0x15, 0xca, 0x3c, 0xe1, 0x21, 0xf0, 0x28, 0x71, 0xa9, 0x2d, 0x98, 0xfd, 0x46, 0x9f, 0xb5,
	0x3c, 0xcf, 0x72, 0x48, 0x35, 0x88, 0xb4, 0xba, 0x5b, 0x55, 0xd2, 0xa1, 0x62, 0x37, 0x04, 0xea,
	0xe7, 0xe3, 0xc1, 0x1d, 0x86, 0x29, 0x25, 0x8c, 0xcb, 0xf8, 0x7c, 0x3c, 0x2e, 0xec, 0x0e, 0xe1,
	0x02, 0x77, 0xa8, 0x04, 0xcc, 0x49, 0x00, 0xa6, 0x76, 0x15, 0xbb, 0xae, 0x27, 0xb0, 0xb0, 0x3d,
	0x37, 0xa2, 0x2f, 0x05, 0x3f, 0xed, 0x65, 0x8b, 0xb8, 0xcb, 0x7c, 0x07, 0x5b, 0x16, 0x61, 0x55,
	0x8f, 0x06, 0x88, 0x24, 0xda, 0xd8, 0xcb, 0xc0, 0x48, 0x83, 0x38, 0x04, 0x73, 0x82, 0x6e, 0x43,
	0x81, 0x85, 0x7f, 0x9b, 0x2e, 0xee, 0x90, 0x92, 0x56, 0xd6, 0x16, 0xf2, 0xb5, 0xb9, 0x4a, 0x98,
	0xae, 0x12, 0xd5, 0x53, 0xd9, 0x14, 0xcc, 0x76, 0xad, 0xa7, 0xd8, 0xe9, 0x92, 0x46, 0x5e, 0x32,
	0x1e, 0xe1, 0x0e, 0x41, 0xd7, 0x61, 0x64, 0x9b, 0x30, 0x6e, 0x7b, 0x6e, 0x29, 0x15, 0x70, 0x67,
	0x13, 0xdc, 0x0d, 0x57, 0xac, 0xd4, 0x42, 0x6a, 0x84, 0x45, 0xab, 0x30, 0xea, 0xe7, 0xe3, 0x14,
	0xb7, 0x49, 0x29, 0x3d, 0x44, 0xd2, 0x7d, 0x38, 0xba, 0x06, 0x39, 0x2e, 0xb0, 0xe8, 0xf2, 0x52,
	0x66, 0x08, 0xa2, 0xc4, 0xa2, 0x5b, 0x90, 0x37, 0x09, 0x6f, 0x33, 0x3b, 0xe8, 0x4b, 0x29, 0x3b,
	0xcc, 0x41, 0xfb, 0x08, 0xe8, 0x1e, 0x4c, 0x6e, 0xd9, 0x8c, 0x8b, 0xa6, 0x49, 0xa8, 0xe3, 0xed,
	0x12, 0xb3, 0xe9, 0x7b, 0x54, 0xca, 0x05, 0x3a, 0x7a, 0x42, 0xe7, 0x71, 0x64, 0x60, 0x63, 0x22,
	0xa0, 0xad, 0x4b, 0x96, 0xff, 0x1d, 0xdd, 0x05, 0xe4, 0xe0, 0x84, 0xd4, 0xc8, 0xa1, 0x52, 0xe3,
	0x3e, 0x4b, 0x51, 0xba, 0x09, 0x05, 0x93, 0x38, 0x44, 0x44, 0x1a, 0x67, 0x0e, 0xd5, 0xc8, 0x4b,
	0xbc, 0xff, 0xc5, 0xf8, 0x90, 0x82, 0xc9, 0x07, 0x36, 0x17, 0x72, 0x1c, 0x78, 0x83, 0xbc, 0xee,
	0x12, 0x2e, 0xd0, 0x0d, 0x00, 0xd6, 0x75, 0x7d, 0xc5, 0xa6, 0x6d, 0x0e, 0x35, 0x14, 0xa3, 0x12,
	0xbf, 0x61, 0x26, 0x66, 0x2a, 0x75, 0xd4, 0x99, 0xfa, 0xe7, 0xc3, 0x61, 0xdc, 0x57, 0xda, 0xd0,
	0x20, 0x9c, 0x7a, 0x2e, 0xf7, 0xc5, 0xa2, 0xba, 0x9a, 0x9c, 0x88, 0x92, 0x56, 0x4e, 0x2f, 0xe4,
	0x6b, 0x93, 0x95, 0xfd, 0x5b, 0x5f, 0x89, 0x18, 0x20, 0x71, 0x9b, 0x44, 0x18, 0x5f, 0x52, 0x30,
	0x75, 0x87, 0x11, 0x2c, 0xc8, 0x6f, 0xbd, 0xb0, 0xab, 0x2b, 0x90, 0xc3, 0x94, 0x0e, 0xdb, 0xd1,
	0x2c, 0xa6, 0x74, 0xc3, 0xf4, 0xad, 0x90, 0x97, 0xc6, 0x27, 0x0e, 0xd3, 0xcb, 0x51, 0x89, 0x0f,
	0xc9, 0x7d, 0x3e, 0xa6, 0x8f, 0xe7, 0x63, 0xe6, 0x58, 0x3e, 0x66, 0x8f, 0xe4, 0xa3, 0xf1, 0x0c,
	0xa6, 0x63, 0x3d, 0x94, 0x9e, 0x1c, 0x77, 0x63, 0x19, 0xdf, 0x35, 0x98, 0x7e, 0x42, 0x2d, 0x86,
	0xcd, 0xb8, 0x3f, 0xff, 0x77, 0xea, 0x55, 0xa3, 0xd3, 0x47, 0x32, 0xda, 0x78, 0x0e, 0x33, 0xf1,
	0x33, 0x9d, 0x54, 0xbf, 0xbe, 0x69, 0x30, 0xd3, 0xf0, 0x1c, 0xa7, 0x85, 0xdb, 0xaf, 0x4e, 0x55,
	0xc3, 0xfa, 0x9e, 0x9e, 0xf4, 0xf0, 0x4f, 0x8f, 0xf1, 0x51, 0x83, 0x73, 0x89, 0xf3, 0x9c, 0x50,
	0xb3, 0xfe, 0xf2, 0x39, 0x34, 0xbe, 0x6a, 0x30, 0xb5, 0x1e, 0xec, 0xe5, 0x53, 0xd5, 0xe1, 0x2b,
	0x90, 0xa5, 0x5d, 0x66, 0x45, 0x4b, 0x38, 0xf9, 0xac, 0xac, 0x79, 0x9e, 0x23, 0xb7, 0x55, 0x00,
	0xf4, 0xaf, 0x6d, 0xec, 0x1c, 0x27, 0xd4, 0xd9, 0xda, 0xe7, 0x2c, 0x14, 0xa5, 0xe8, 0x43, 0xec,
	0x62, 0x8b, 0x30, 0xb4, 0x03, 0x85, 0xfe, 0xc7, 0x0b, 0xcd, 0xf7, 0x6f, 0xe6, 0x01, 0xcf, 0x9a,
	0xfe, 0x27, 0x40, 0x54, 0xa5, 0x71, 0xa9, 0x57, 0x2f, 0xa2, 0x40, 0xb4, 0x2c, 0x43, 0xef, 0x7e,
	0xfc, 0xfc, 0x94, 0x2a, 0xa2, 0x42, 0x75, 0xfb, 0x6a, 0x95, 0x45, 0x89, 0xde, 0x6a, 0x70, 0x56,
	0xd9, 0x4e, 0xa8, 0xdc, 0xaf, 0x3c, 0x68, 0xf9, 0xeb, 0x17, 0x0e, 0x40, 0xc8, 0xec, 0x8b, 0xbd,
	0xfa, 0x38, 0x2a, 0x86, 0x31, 0x25, 0xff, 0x84, 0xae, 0xe4, 0x5f, 0xd5, 0x16, 0xd1, 0x7b, 0x0d,
	0x8a, 0xea, 0x8d, 0x47, 0x4a, 0x86, 0x81, 0x1b, 0x4e, 0x37, 0x0e, 0x82, 0xc8, 0x2a, 0x2e, 0xf7,
	0xea, 0x13, 0x68, 0x4c, 0x06, 0xd5, 0x32, 0x6a, 0x89, 0x32, 0xf6, 0x34, 0x18, 0x8b, 0x5d, 0x26,
	0xa4, 0x24, 0x19, 0xbc, 0x39, 0xf4, 0x8b, 0x07, 0x62, 0x64, 0x25, 0x4b, 0xbd, 0x3a, 0x42, 0xe3,
	0x51, 0x54, 0x2d, 0xc5, 0x48, 0x94, 0xe2, 0x9b, 0xa2, 0xcc, 0x9e, 0x6a, 0xca, 0xa0, 0xeb, 0xa5,
	0x9a, 0x32, 0x70, 0x70, 0xa5, 0x29, 0x61, 0x4c, 0x2d, 0x61, 0x31, 0x5e, 0xc2, 0x5a, 0xe6, 0x45,
	0x8a, 0xb6, 0x5a, 0xb9, 0x60, 0x98, 0x57, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x0a, 0xf9, 0x10,
	0xa6, 0x2a, 0x0c, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ReleaseManagerClient is the client API for ReleaseManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReleaseManagerClient interface {
	ListReleases(ctx context.Context, in *ListReleasesRequest, opts ...grpc.CallOption) (*ListReleaseResponse, error)
	CreateRelease(ctx context.Context, in *CreateReleaseRequest, opts ...grpc.CallOption) (*CreateReleaseResponse, error)
	UpgradeRelease(ctx context.Context, in *UpgradeReleaseRequest, opts ...grpc.CallOption) (*UpgradeReleaseResponse, error)
	RollbackRelease(ctx context.Context, in *RollbackReleaseRequest, opts ...grpc.CallOption) (*RollbackReleaseResponse, error)
	DeleteRelease(ctx context.Context, in *DeleteReleaseRequest, opts ...grpc.CallOption) (*DeleteReleaseResponse, error)
}

type releaseManagerClient struct {
	cc *grpc.ClientConn
}

func NewReleaseManagerClient(cc *grpc.ClientConn) ReleaseManagerClient {
	return &releaseManagerClient{cc}
}

func (c *releaseManagerClient) ListReleases(ctx context.Context, in *ListReleasesRequest, opts ...grpc.CallOption) (*ListReleaseResponse, error) {
	out := new(ListReleaseResponse)
	err := c.cc.Invoke(ctx, "/openpitrix.ReleaseManager/ListReleases", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseManagerClient) CreateRelease(ctx context.Context, in *CreateReleaseRequest, opts ...grpc.CallOption) (*CreateReleaseResponse, error) {
	out := new(CreateReleaseResponse)
	err := c.cc.Invoke(ctx, "/openpitrix.ReleaseManager/CreateRelease", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseManagerClient) UpgradeRelease(ctx context.Context, in *UpgradeReleaseRequest, opts ...grpc.CallOption) (*UpgradeReleaseResponse, error) {
	out := new(UpgradeReleaseResponse)
	err := c.cc.Invoke(ctx, "/openpitrix.ReleaseManager/UpgradeRelease", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseManagerClient) RollbackRelease(ctx context.Context, in *RollbackReleaseRequest, opts ...grpc.CallOption) (*RollbackReleaseResponse, error) {
	out := new(RollbackReleaseResponse)
	err := c.cc.Invoke(ctx, "/openpitrix.ReleaseManager/RollbackRelease", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseManagerClient) DeleteRelease(ctx context.Context, in *DeleteReleaseRequest, opts ...grpc.CallOption) (*DeleteReleaseResponse, error) {
	out := new(DeleteReleaseResponse)
	err := c.cc.Invoke(ctx, "/openpitrix.ReleaseManager/DeleteRelease", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReleaseManagerServer is the server API for ReleaseManager service.
type ReleaseManagerServer interface {
	ListReleases(context.Context, *ListReleasesRequest) (*ListReleaseResponse, error)
	CreateRelease(context.Context, *CreateReleaseRequest) (*CreateReleaseResponse, error)
	UpgradeRelease(context.Context, *UpgradeReleaseRequest) (*UpgradeReleaseResponse, error)
	RollbackRelease(context.Context, *RollbackReleaseRequest) (*RollbackReleaseResponse, error)
	DeleteRelease(context.Context, *DeleteReleaseRequest) (*DeleteReleaseResponse, error)
}

// UnimplementedReleaseManagerServer can be embedded to have forward compatible implementations.
type UnimplementedReleaseManagerServer struct {
}

func (*UnimplementedReleaseManagerServer) ListReleases(ctx context.Context, req *ListReleasesRequest) (*ListReleaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListReleases not implemented")
}
func (*UnimplementedReleaseManagerServer) CreateRelease(ctx context.Context, req *CreateReleaseRequest) (*CreateReleaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRelease not implemented")
}
func (*UnimplementedReleaseManagerServer) UpgradeRelease(ctx context.Context, req *UpgradeReleaseRequest) (*UpgradeReleaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpgradeRelease not implemented")
}
func (*UnimplementedReleaseManagerServer) RollbackRelease(ctx context.Context, req *RollbackReleaseRequest) (*RollbackReleaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RollbackRelease not implemented")
}
func (*UnimplementedReleaseManagerServer) DeleteRelease(ctx context.Context, req *DeleteReleaseRequest) (*DeleteReleaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRelease not implemented")
}

func RegisterReleaseManagerServer(s *grpc.Server, srv ReleaseManagerServer) {
	s.RegisterService(&_ReleaseManager_serviceDesc, srv)
}

func _ReleaseManager_ListReleases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReleasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseManagerServer).ListReleases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.ReleaseManager/ListReleases",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseManagerServer).ListReleases(ctx, req.(*ListReleasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReleaseManager_CreateRelease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseManagerServer).CreateRelease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.ReleaseManager/CreateRelease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseManagerServer).CreateRelease(ctx, req.(*CreateReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReleaseManager_UpgradeRelease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpgradeReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseManagerServer).UpgradeRelease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.ReleaseManager/UpgradeRelease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseManagerServer).UpgradeRelease(ctx, req.(*UpgradeReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReleaseManager_RollbackRelease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RollbackReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseManagerServer).RollbackRelease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.ReleaseManager/RollbackRelease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseManagerServer).RollbackRelease(ctx, req.(*RollbackReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReleaseManager_DeleteRelease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseManagerServer).DeleteRelease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.ReleaseManager/DeleteRelease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseManagerServer).DeleteRelease(ctx, req.(*DeleteReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReleaseManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openpitrix.ReleaseManager",
	HandlerType: (*ReleaseManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListReleases",
			Handler:    _ReleaseManager_ListReleases_Handler,
		},
		{
			MethodName: "CreateRelease",
			Handler:    _ReleaseManager_CreateRelease_Handler,
		},
		{
			MethodName: "UpgradeRelease",
			Handler:    _ReleaseManager_UpgradeRelease_Handler,
		},
		{
			MethodName: "RollbackRelease",
			Handler:    _ReleaseManager_RollbackRelease_Handler,
		},
		{
			MethodName: "DeleteRelease",
			Handler:    _ReleaseManager_DeleteRelease_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helm.proto",
}