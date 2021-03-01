// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.3
// source: pkg/pb/deployment.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GithubDeploymentState int32

const (
	GithubDeploymentState_success     GithubDeploymentState = 0
	GithubDeploymentState_error       GithubDeploymentState = 1
	GithubDeploymentState_failure     GithubDeploymentState = 2
	GithubDeploymentState_inactive    GithubDeploymentState = 3
	GithubDeploymentState_in_progress GithubDeploymentState = 4
	GithubDeploymentState_queued      GithubDeploymentState = 5
	GithubDeploymentState_pending     GithubDeploymentState = 6
)

// Enum value maps for GithubDeploymentState.
var (
	GithubDeploymentState_name = map[int32]string{
		0: "success",
		1: "error",
		2: "failure",
		3: "inactive",
		4: "in_progress",
		5: "queued",
		6: "pending",
	}
	GithubDeploymentState_value = map[string]int32{
		"success":     0,
		"error":       1,
		"failure":     2,
		"inactive":    3,
		"in_progress": 4,
		"queued":      5,
		"pending":     6,
	}
)

func (x GithubDeploymentState) Enum() *GithubDeploymentState {
	p := new(GithubDeploymentState)
	*p = x
	return p
}

func (x GithubDeploymentState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GithubDeploymentState) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_pb_deployment_proto_enumTypes[0].Descriptor()
}

func (GithubDeploymentState) Type() protoreflect.EnumType {
	return &file_pkg_pb_deployment_proto_enumTypes[0]
}

func (x GithubDeploymentState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GithubDeploymentState.Descriptor instead.
func (GithubDeploymentState) EnumDescriptor() ([]byte, []int) {
	return file_pkg_pb_deployment_proto_rawDescGZIP(), []int{0}
}

type GithubRepository struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GithubRepository) Reset() {
	*x = GithubRepository{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_deployment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GithubRepository) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GithubRepository) ProtoMessage() {}

func (x *GithubRepository) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_deployment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GithubRepository.ProtoReflect.Descriptor instead.
func (*GithubRepository) Descriptor() ([]byte, []int) {
	return file_pkg_pb_deployment_proto_rawDescGZIP(), []int{0}
}

func (x *GithubRepository) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *GithubRepository) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeploymentSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Repository   *GithubRepository `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
	DeploymentID int64             `protobuf:"varint,2,opt,name=deploymentID,proto3" json:"deploymentID,omitempty"`
	Environment  string            `protobuf:"bytes,3,opt,name=environment,proto3" json:"environment,omitempty"`
	Ref          string            `protobuf:"bytes,4,opt,name=ref,proto3" json:"ref,omitempty"`
}

func (x *DeploymentSpec) Reset() {
	*x = DeploymentSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_deployment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeploymentSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploymentSpec) ProtoMessage() {}

func (x *DeploymentSpec) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_deployment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploymentSpec.ProtoReflect.Descriptor instead.
func (*DeploymentSpec) Descriptor() ([]byte, []int) {
	return file_pkg_pb_deployment_proto_rawDescGZIP(), []int{1}
}

func (x *DeploymentSpec) GetRepository() *GithubRepository {
	if x != nil {
		return x.Repository
	}
	return nil
}

func (x *DeploymentSpec) GetDeploymentID() int64 {
	if x != nil {
		return x.DeploymentID
	}
	return 0
}

func (x *DeploymentSpec) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

func (x *DeploymentSpec) GetRef() string {
	if x != nil {
		return x.Ref
	}
	return ""
}

type Kubernetes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resources []*structpb.Struct `protobuf:"bytes,1,rep,name=resources,proto3" json:"resources,omitempty"`
}

func (x *Kubernetes) Reset() {
	*x = Kubernetes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_deployment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Kubernetes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Kubernetes) ProtoMessage() {}

func (x *Kubernetes) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_deployment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Kubernetes.ProtoReflect.Descriptor instead.
func (*Kubernetes) Descriptor() ([]byte, []int) {
	return file_pkg_pb_deployment_proto_rawDescGZIP(), []int{2}
}

func (x *Kubernetes) GetResources() []*structpb.Struct {
	if x != nil {
		return x.Resources
	}
	return nil
}

type Payload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version    []int32     `protobuf:"varint,1,rep,packed,name=version,proto3" json:"version,omitempty"`
	Team       string      `protobuf:"bytes,2,opt,name=team,proto3" json:"team,omitempty"`
	Kubernetes *Kubernetes `protobuf:"bytes,3,opt,name=kubernetes,proto3" json:"kubernetes,omitempty"`
}

func (x *Payload) Reset() {
	*x = Payload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_deployment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payload) ProtoMessage() {}

func (x *Payload) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_deployment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payload.ProtoReflect.Descriptor instead.
func (*Payload) Descriptor() ([]byte, []int) {
	return file_pkg_pb_deployment_proto_rawDescGZIP(), []int{3}
}

func (x *Payload) GetVersion() []int32 {
	if x != nil {
		return x.Version
	}
	return nil
}

func (x *Payload) GetTeam() string {
	if x != nil {
		return x.Team
	}
	return ""
}

func (x *Payload) GetKubernetes() *Kubernetes {
	if x != nil {
		return x.Kubernetes
	}
	return nil
}

type DeploymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deployment  *DeploymentSpec        `protobuf:"bytes,1,opt,name=deployment,proto3" json:"deployment,omitempty"`
	Deadline    int64                  `protobuf:"varint,3,opt,name=deadline,proto3" json:"deadline,omitempty"`
	Cluster     string                 `protobuf:"bytes,5,opt,name=cluster,proto3" json:"cluster,omitempty"`
	DeliveryID  string                 `protobuf:"bytes,6,opt,name=deliveryID,proto3" json:"deliveryID,omitempty"`
	PayloadSpec *Payload               `protobuf:"bytes,7,opt,name=payloadSpec,proto3" json:"payloadSpec,omitempty"`
	Time        *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=time,proto3" json:"time,omitempty"`
	Status      *DeploymentStatus      `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
	GitRefSha   string                 `protobuf:"bytes,10,opt,name=gitRefSha,proto3" json:"gitRefSha,omitempty"`
}

func (x *DeploymentRequest) Reset() {
	*x = DeploymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_deployment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeploymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploymentRequest) ProtoMessage() {}

func (x *DeploymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_deployment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploymentRequest.ProtoReflect.Descriptor instead.
func (*DeploymentRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_deployment_proto_rawDescGZIP(), []int{4}
}

func (x *DeploymentRequest) GetDeployment() *DeploymentSpec {
	if x != nil {
		return x.Deployment
	}
	return nil
}

func (x *DeploymentRequest) GetDeadline() int64 {
	if x != nil {
		return x.Deadline
	}
	return 0
}

func (x *DeploymentRequest) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *DeploymentRequest) GetDeliveryID() string {
	if x != nil {
		return x.DeliveryID
	}
	return ""
}

func (x *DeploymentRequest) GetPayloadSpec() *Payload {
	if x != nil {
		return x.PayloadSpec
	}
	return nil
}

func (x *DeploymentRequest) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *DeploymentRequest) GetStatus() *DeploymentStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *DeploymentRequest) GetGitRefSha() string {
	if x != nil {
		return x.GitRefSha
	}
	return ""
}

type DeploymentStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deployment  *DeploymentSpec        `protobuf:"bytes,1,opt,name=deployment,proto3" json:"deployment,omitempty"`
	State       GithubDeploymentState  `protobuf:"varint,2,opt,name=state,proto3,enum=pb.GithubDeploymentState" json:"state,omitempty"`
	Description string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	DeliveryID  string                 `protobuf:"bytes,4,opt,name=deliveryID,proto3" json:"deliveryID,omitempty"`
	Team        string                 `protobuf:"bytes,5,opt,name=team,proto3" json:"team,omitempty"`
	Cluster     string                 `protobuf:"bytes,6,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Time        *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=time,proto3" json:"time,omitempty"`
	Id          string                 `protobuf:"bytes,9,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeploymentStatus) Reset() {
	*x = DeploymentStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_deployment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeploymentStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploymentStatus) ProtoMessage() {}

func (x *DeploymentStatus) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_deployment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploymentStatus.ProtoReflect.Descriptor instead.
func (*DeploymentStatus) Descriptor() ([]byte, []int) {
	return file_pkg_pb_deployment_proto_rawDescGZIP(), []int{5}
}

func (x *DeploymentStatus) GetDeployment() *DeploymentSpec {
	if x != nil {
		return x.Deployment
	}
	return nil
}

func (x *DeploymentStatus) GetState() GithubDeploymentState {
	if x != nil {
		return x.State
	}
	return GithubDeploymentState_success
}

func (x *DeploymentStatus) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *DeploymentStatus) GetDeliveryID() string {
	if x != nil {
		return x.DeliveryID
	}
	return ""
}

func (x *DeploymentStatus) GetTeam() string {
	if x != nil {
		return x.Team
	}
	return ""
}

func (x *DeploymentStatus) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *DeploymentStatus) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *DeploymentStatus) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SignedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message   []byte `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *SignedMessage) Reset() {
	*x = SignedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_deployment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedMessage) ProtoMessage() {}

func (x *SignedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_deployment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedMessage.ProtoReflect.Descriptor instead.
func (*SignedMessage) Descriptor() ([]byte, []int) {
	return file_pkg_pb_deployment_proto_rawDescGZIP(), []int{6}
}

func (x *SignedMessage) GetMessage() []byte {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *SignedMessage) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type GetDeploymentOpts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cluster string `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
}

func (x *GetDeploymentOpts) Reset() {
	*x = GetDeploymentOpts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_deployment_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeploymentOpts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeploymentOpts) ProtoMessage() {}

func (x *GetDeploymentOpts) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_deployment_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeploymentOpts.ProtoReflect.Descriptor instead.
func (*GetDeploymentOpts) Descriptor() ([]byte, []int) {
	return file_pkg_pb_deployment_proto_rawDescGZIP(), []int{7}
}

func (x *GetDeploymentOpts) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

type ReportStatusOpts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReportStatusOpts) Reset() {
	*x = ReportStatusOpts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_deployment_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportStatusOpts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportStatusOpts) ProtoMessage() {}

func (x *ReportStatusOpts) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_deployment_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportStatusOpts.ProtoReflect.Descriptor instead.
func (*ReportStatusOpts) Descriptor() ([]byte, []int) {
	return file_pkg_pb_deployment_proto_rawDescGZIP(), []int{8}
}

var File_pkg_pb_deployment_proto protoreflect.FileDescriptor

var file_pkg_pb_deployment_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3c, 0x0a, 0x10,
	0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x9e, 0x01, 0x0a, 0x0e, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x70, 0x65, 0x63, 0x12, 0x34, 0x0a,
	0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x52, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x6f, 0x72, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x66,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x65, 0x66, 0x22, 0x43, 0x0a, 0x0a, 0x4b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x12, 0x35, 0x0a, 0x09, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x22, 0x67, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x12, 0x2e, 0x0a, 0x0a, 0x6b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x70, 0x62, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x52, 0x0a, 0x6b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x22, 0xe8, 0x02, 0x0a, 0x11, 0x44, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x32, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x53, 0x70, 0x65, 0x63, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x49, 0x44, 0x12, 0x2d, 0x0a, 0x0b, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x53, 0x70, 0x65, 0x63, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x0b, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x53, 0x70, 0x65, 0x63, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x69, 0x74, 0x52, 0x65, 0x66,
	0x53, 0x68, 0x61, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x69, 0x74, 0x52, 0x65,
	0x66, 0x53, 0x68, 0x61, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x22, 0xb8, 0x02, 0x0a, 0x10, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x32, 0x0a, 0x0a, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x70, 0x62, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x70, 0x65,
	0x63, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2f, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x70,
	0x62, 0x2e, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x49, 0x44, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x49, 0x44,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x65, 0x61, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x2e,
	0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x4a, 0x04,
	0x08, 0x07, 0x10, 0x08, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22,
	0x47, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x2d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x70, 0x74, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x12, 0x0a, 0x10, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4f, 0x70, 0x74, 0x73, 0x2a, 0x74, 0x0a, 0x15, 0x47,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10,
	0x00, 0x12, 0x09, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07,
	0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x69, 0x6e, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x69, 0x6e, 0x5f, 0x70, 0x72,
	0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x71, 0x75, 0x65, 0x75,
	0x65, 0x64, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x10,
	0x06, 0x32, 0x87, 0x01, 0x0a, 0x06, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x12, 0x3f, 0x0a, 0x0b,
	0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x15, 0x2e, 0x70, 0x62,
	0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x70,
	0x74, 0x73, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3c, 0x0a,
	0x0c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x2e,
	0x70, 0x62, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x4f, 0x70, 0x74, 0x73, 0x22, 0x00, 0x42, 0x39, 0x0a, 0x18, 0x6e,
	0x6f, 0x2e, 0x6e, 0x61, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x61, 0x69, 0x73, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_pb_deployment_proto_rawDescOnce sync.Once
	file_pkg_pb_deployment_proto_rawDescData = file_pkg_pb_deployment_proto_rawDesc
)

func file_pkg_pb_deployment_proto_rawDescGZIP() []byte {
	file_pkg_pb_deployment_proto_rawDescOnce.Do(func() {
		file_pkg_pb_deployment_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pb_deployment_proto_rawDescData)
	})
	return file_pkg_pb_deployment_proto_rawDescData
}

var file_pkg_pb_deployment_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_pb_deployment_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pkg_pb_deployment_proto_goTypes = []interface{}{
	(GithubDeploymentState)(0),    // 0: pb.GithubDeploymentState
	(*GithubRepository)(nil),      // 1: pb.GithubRepository
	(*DeploymentSpec)(nil),        // 2: pb.DeploymentSpec
	(*Kubernetes)(nil),            // 3: pb.Kubernetes
	(*Payload)(nil),               // 4: pb.Payload
	(*DeploymentRequest)(nil),     // 5: pb.DeploymentRequest
	(*DeploymentStatus)(nil),      // 6: pb.DeploymentStatus
	(*SignedMessage)(nil),         // 7: pb.SignedMessage
	(*GetDeploymentOpts)(nil),     // 8: pb.GetDeploymentOpts
	(*ReportStatusOpts)(nil),      // 9: pb.ReportStatusOpts
	(*structpb.Struct)(nil),       // 10: google.protobuf.Struct
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
}
var file_pkg_pb_deployment_proto_depIdxs = []int32{
	1,  // 0: pb.DeploymentSpec.repository:type_name -> pb.GithubRepository
	10, // 1: pb.Kubernetes.resources:type_name -> google.protobuf.Struct
	3,  // 2: pb.Payload.kubernetes:type_name -> pb.Kubernetes
	2,  // 3: pb.DeploymentRequest.deployment:type_name -> pb.DeploymentSpec
	4,  // 4: pb.DeploymentRequest.payloadSpec:type_name -> pb.Payload
	11, // 5: pb.DeploymentRequest.time:type_name -> google.protobuf.Timestamp
	6,  // 6: pb.DeploymentRequest.status:type_name -> pb.DeploymentStatus
	2,  // 7: pb.DeploymentStatus.deployment:type_name -> pb.DeploymentSpec
	0,  // 8: pb.DeploymentStatus.state:type_name -> pb.GithubDeploymentState
	11, // 9: pb.DeploymentStatus.time:type_name -> google.protobuf.Timestamp
	8,  // 10: pb.Deploy.Deployments:input_type -> pb.GetDeploymentOpts
	6,  // 11: pb.Deploy.ReportStatus:input_type -> pb.DeploymentStatus
	5,  // 12: pb.Deploy.Deployments:output_type -> pb.DeploymentRequest
	9,  // 13: pb.Deploy.ReportStatus:output_type -> pb.ReportStatusOpts
	12, // [12:14] is the sub-list for method output_type
	10, // [10:12] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_pkg_pb_deployment_proto_init() }
func file_pkg_pb_deployment_proto_init() {
	if File_pkg_pb_deployment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_pb_deployment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GithubRepository); i {
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
		file_pkg_pb_deployment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeploymentSpec); i {
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
		file_pkg_pb_deployment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Kubernetes); i {
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
		file_pkg_pb_deployment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Payload); i {
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
		file_pkg_pb_deployment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeploymentRequest); i {
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
		file_pkg_pb_deployment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeploymentStatus); i {
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
		file_pkg_pb_deployment_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignedMessage); i {
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
		file_pkg_pb_deployment_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeploymentOpts); i {
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
		file_pkg_pb_deployment_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportStatusOpts); i {
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
			RawDescriptor: file_pkg_pb_deployment_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_pb_deployment_proto_goTypes,
		DependencyIndexes: file_pkg_pb_deployment_proto_depIdxs,
		EnumInfos:         file_pkg_pb_deployment_proto_enumTypes,
		MessageInfos:      file_pkg_pb_deployment_proto_msgTypes,
	}.Build()
	File_pkg_pb_deployment_proto = out.File
	file_pkg_pb_deployment_proto_rawDesc = nil
	file_pkg_pb_deployment_proto_goTypes = nil
	file_pkg_pb_deployment_proto_depIdxs = nil
}
