// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: spirits/v1/spirit.proto

package spiritsv1

import (
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

type CreateSpiritRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Spirit to create. The ID will be filled in on response.
	Spirit *Spirit `protobuf:"bytes,1,opt,name=spirit,proto3" json:"spirit,omitempty"`
}

func (x *CreateSpiritRequest) Reset() {
	*x = CreateSpiritRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spirits_v1_spirit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSpiritRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSpiritRequest) ProtoMessage() {}

func (x *CreateSpiritRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spirits_v1_spirit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSpiritRequest.ProtoReflect.Descriptor instead.
func (*CreateSpiritRequest) Descriptor() ([]byte, []int) {
	return file_spirits_v1_spirit_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSpiritRequest) GetSpirit() *Spirit {
	if x != nil {
		return x.Spirit
	}
	return nil
}

type CreateSpiritResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The newly created Spirit.
	Spirit *Spirit `protobuf:"bytes,1,opt,name=spirit,proto3" json:"spirit,omitempty"`
}

func (x *CreateSpiritResponse) Reset() {
	*x = CreateSpiritResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spirits_v1_spirit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSpiritResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSpiritResponse) ProtoMessage() {}

func (x *CreateSpiritResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spirits_v1_spirit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSpiritResponse.ProtoReflect.Descriptor instead.
func (*CreateSpiritResponse) Descriptor() ([]byte, []int) {
	return file_spirits_v1_spirit_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSpiritResponse) GetSpirit() *Spirit {
	if x != nil {
		return x.Spirit
	}
	return nil
}

type GetSpiritRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the Spirit to get.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetSpiritRequest) Reset() {
	*x = GetSpiritRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spirits_v1_spirit_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSpiritRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSpiritRequest) ProtoMessage() {}

func (x *GetSpiritRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spirits_v1_spirit_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSpiritRequest.ProtoReflect.Descriptor instead.
func (*GetSpiritRequest) Descriptor() ([]byte, []int) {
	return file_spirits_v1_spirit_proto_rawDescGZIP(), []int{2}
}

func (x *GetSpiritRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetSpiritResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The retrieved Spirit.
	Spirit *Spirit `protobuf:"bytes,1,opt,name=spirit,proto3" json:"spirit,omitempty"`
}

func (x *GetSpiritResponse) Reset() {
	*x = GetSpiritResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spirits_v1_spirit_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSpiritResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSpiritResponse) ProtoMessage() {}

func (x *GetSpiritResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spirits_v1_spirit_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSpiritResponse.ProtoReflect.Descriptor instead.
func (*GetSpiritResponse) Descriptor() ([]byte, []int) {
	return file_spirits_v1_spirit_proto_rawDescGZIP(), []int{3}
}

func (x *GetSpiritResponse) GetSpirit() *Spirit {
	if x != nil {
		return x.Spirit
	}
	return nil
}

type ListSpiritsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An optional name to use as a filter.
	Name *string `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
}

func (x *ListSpiritsRequest) Reset() {
	*x = ListSpiritsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spirits_v1_spirit_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSpiritsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSpiritsRequest) ProtoMessage() {}

func (x *ListSpiritsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spirits_v1_spirit_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSpiritsRequest.ProtoReflect.Descriptor instead.
func (*ListSpiritsRequest) Descriptor() ([]byte, []int) {
	return file_spirits_v1_spirit_proto_rawDescGZIP(), []int{4}
}

func (x *ListSpiritsRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

type ListSpiritsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Spirits []*Spirit `protobuf:"bytes,1,rep,name=spirits,proto3" json:"spirits,omitempty"`
}

func (x *ListSpiritsResponse) Reset() {
	*x = ListSpiritsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spirits_v1_spirit_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSpiritsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSpiritsResponse) ProtoMessage() {}

func (x *ListSpiritsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spirits_v1_spirit_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSpiritsResponse.ProtoReflect.Descriptor instead.
func (*ListSpiritsResponse) Descriptor() ([]byte, []int) {
	return file_spirits_v1_spirit_proto_rawDescGZIP(), []int{5}
}

func (x *ListSpiritsResponse) GetSpirits() []*Spirit {
	if x != nil {
		return x.Spirits
	}
	return nil
}

type UpdateSpiritRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The new version of the Spirit.
	Spirit *Spirit `protobuf:"bytes,1,opt,name=spirit,proto3" json:"spirit,omitempty"`
}

func (x *UpdateSpiritRequest) Reset() {
	*x = UpdateSpiritRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spirits_v1_spirit_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSpiritRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSpiritRequest) ProtoMessage() {}

func (x *UpdateSpiritRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spirits_v1_spirit_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSpiritRequest.ProtoReflect.Descriptor instead.
func (*UpdateSpiritRequest) Descriptor() ([]byte, []int) {
	return file_spirits_v1_spirit_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateSpiritRequest) GetSpirit() *Spirit {
	if x != nil {
		return x.Spirit
	}
	return nil
}

type UpdateSpiritResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The newly updated Spirit.
	Spirit *Spirit `protobuf:"bytes,1,opt,name=spirit,proto3" json:"spirit,omitempty"`
}

func (x *UpdateSpiritResponse) Reset() {
	*x = UpdateSpiritResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spirits_v1_spirit_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSpiritResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSpiritResponse) ProtoMessage() {}

func (x *UpdateSpiritResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spirits_v1_spirit_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSpiritResponse.ProtoReflect.Descriptor instead.
func (*UpdateSpiritResponse) Descriptor() ([]byte, []int) {
	return file_spirits_v1_spirit_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateSpiritResponse) GetSpirit() *Spirit {
	if x != nil {
		return x.Spirit
	}
	return nil
}

type DeleteSpiritRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the Spirit to delete.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteSpiritRequest) Reset() {
	*x = DeleteSpiritRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spirits_v1_spirit_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSpiritRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSpiritRequest) ProtoMessage() {}

func (x *DeleteSpiritRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spirits_v1_spirit_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSpiritRequest.ProtoReflect.Descriptor instead.
func (*DeleteSpiritRequest) Descriptor() ([]byte, []int) {
	return file_spirits_v1_spirit_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteSpiritRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteSpiritResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The newly deleted Spirit.
	Spirit *Spirit `protobuf:"bytes,1,opt,name=spirit,proto3" json:"spirit,omitempty"`
}

func (x *DeleteSpiritResponse) Reset() {
	*x = DeleteSpiritResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spirits_v1_spirit_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSpiritResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSpiritResponse) ProtoMessage() {}

func (x *DeleteSpiritResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spirits_v1_spirit_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSpiritResponse.ProtoReflect.Descriptor instead.
func (*DeleteSpiritResponse) Descriptor() ([]byte, []int) {
	return file_spirits_v1_spirit_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteSpiritResponse) GetSpirit() *Spirit {
	if x != nil {
		return x.Spirit
	}
	return nil
}

// Spirit describes a single actor in a Battle.
type Spirit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Metadata about this Spirit.
	Meta *Meta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// Name of this Spirit - colloquial name used to identify the Spirit.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The quantitative description of this Spirit's abilities.
	Stats *SpiritStats `protobuf:"bytes,3,opt,name=stats,proto3" json:"stats,omitempty"`
	// The SpiritAction's that this Spirit can take on their turn in a Battle.
	Actions []*SpiritAction `protobuf:"bytes,4,rep,name=actions,proto3" json:"actions,omitempty"`
}

func (x *Spirit) Reset() {
	*x = Spirit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spirits_v1_spirit_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Spirit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Spirit) ProtoMessage() {}

func (x *Spirit) ProtoReflect() protoreflect.Message {
	mi := &file_spirits_v1_spirit_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Spirit.ProtoReflect.Descriptor instead.
func (*Spirit) Descriptor() ([]byte, []int) {
	return file_spirits_v1_spirit_proto_rawDescGZIP(), []int{10}
}

func (x *Spirit) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Spirit) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Spirit) GetStats() *SpiritStats {
	if x != nil {
		return x.Stats
	}
	return nil
}

func (x *Spirit) GetActions() []*SpiritAction {
	if x != nil {
		return x.Actions
	}
	return nil
}

// SpiritStats are a quantitative description of a Spirit's abilities.
type SpiritStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Health               int64 `protobuf:"varint,1,opt,name=health,proto3" json:"health,omitempty"`
	PhysicalPower        int64 `protobuf:"varint,2,opt,name=physical_power,json=physicalPower,proto3" json:"physical_power,omitempty"`
	PhysicalConstitution int64 `protobuf:"varint,3,opt,name=physical_constitution,json=physicalConstitution,proto3" json:"physical_constitution,omitempty"`
	MentalPower          int64 `protobuf:"varint,4,opt,name=mental_power,json=mentalPower,proto3" json:"mental_power,omitempty"`
	MentalConstitution   int64 `protobuf:"varint,5,opt,name=mental_constitution,json=mentalConstitution,proto3" json:"mental_constitution,omitempty"`
	Agility              int64 `protobuf:"varint,6,opt,name=agility,proto3" json:"agility,omitempty"`
}

func (x *SpiritStats) Reset() {
	*x = SpiritStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spirits_v1_spirit_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpiritStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpiritStats) ProtoMessage() {}

func (x *SpiritStats) ProtoReflect() protoreflect.Message {
	mi := &file_spirits_v1_spirit_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpiritStats.ProtoReflect.Descriptor instead.
func (*SpiritStats) Descriptor() ([]byte, []int) {
	return file_spirits_v1_spirit_proto_rawDescGZIP(), []int{11}
}

func (x *SpiritStats) GetHealth() int64 {
	if x != nil {
		return x.Health
	}
	return 0
}

func (x *SpiritStats) GetPhysicalPower() int64 {
	if x != nil {
		return x.PhysicalPower
	}
	return 0
}

func (x *SpiritStats) GetPhysicalConstitution() int64 {
	if x != nil {
		return x.PhysicalConstitution
	}
	return 0
}

func (x *SpiritStats) GetMentalPower() int64 {
	if x != nil {
		return x.MentalPower
	}
	return 0
}

func (x *SpiritStats) GetMentalConstitution() int64 {
	if x != nil {
		return x.MentalConstitution
	}
	return 0
}

func (x *SpiritStats) GetAgility() int64 {
	if x != nil {
		return x.Agility
	}
	return 0
}

// SpiritAction is a reference to an Action that a Spirit can invoke during
// Battle.
type SpiritAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of this SpiritAction.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are assignable to Definition:
	//
	//	*SpiritAction_ActionId
	//	*SpiritAction_Inline
	Definition isSpiritAction_Definition `protobuf_oneof:"definition"`
}

func (x *SpiritAction) Reset() {
	*x = SpiritAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spirits_v1_spirit_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpiritAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpiritAction) ProtoMessage() {}

func (x *SpiritAction) ProtoReflect() protoreflect.Message {
	mi := &file_spirits_v1_spirit_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpiritAction.ProtoReflect.Descriptor instead.
func (*SpiritAction) Descriptor() ([]byte, []int) {
	return file_spirits_v1_spirit_proto_rawDescGZIP(), []int{12}
}

func (x *SpiritAction) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (m *SpiritAction) GetDefinition() isSpiritAction_Definition {
	if m != nil {
		return m.Definition
	}
	return nil
}

func (x *SpiritAction) GetActionId() string {
	if x, ok := x.GetDefinition().(*SpiritAction_ActionId); ok {
		return x.ActionId
	}
	return ""
}

func (x *SpiritAction) GetInline() *Action {
	if x, ok := x.GetDefinition().(*SpiritAction_Inline); ok {
		return x.Inline
	}
	return nil
}

type isSpiritAction_Definition interface {
	isSpiritAction_Definition()
}

type SpiritAction_ActionId struct {
	// The ID of the Action.
	ActionId string `protobuf:"bytes,2,opt,name=action_id,json=actionId,proto3,oneof"`
}

type SpiritAction_Inline struct {
	// An inline definition of the Action.
	Inline *Action `protobuf:"bytes,3,opt,name=inline,proto3,oneof"`
}

func (*SpiritAction_ActionId) isSpiritAction_Definition() {}

func (*SpiritAction_Inline) isSpiritAction_Definition() {}

var File_spirits_v1_spirit_proto protoreflect.FileDescriptor

var file_spirits_v1_spirit_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x70, 0x69, 0x72, 0x69, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x70, 0x69,
	0x72, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73, 0x70, 0x69, 0x72, 0x69,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x73, 0x70, 0x69, 0x72, 0x69, 0x74, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15,
	0x73, 0x70, 0x69, 0x72, 0x69, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x70, 0x69,
	0x72, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x70,
	0x69, 0x72, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x70, 0x69,
	0x72, 0x69, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x69, 0x72, 0x69, 0x74, 0x52, 0x06,
	0x73, 0x70, 0x69, 0x72, 0x69, 0x74, 0x22, 0x42, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x70, 0x69, 0x72, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a,
	0x0a, 0x06, 0x73, 0x70, 0x69, 0x72, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x73, 0x70, 0x69, 0x72, 0x69, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x69, 0x72,
	0x69, 0x74, 0x52, 0x06, 0x73, 0x70, 0x69, 0x72, 0x69, 0x74, 0x22, 0x22, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x53, 0x70, 0x69, 0x72, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3f,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x70, 0x69, 0x72, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x70, 0x69, 0x72, 0x69, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x69, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x70, 0x69, 0x72, 0x69, 0x74, 0x52, 0x06, 0x73, 0x70, 0x69, 0x72, 0x69, 0x74, 0x22,
	0x36, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x70, 0x69, 0x72, 0x69, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x07,
	0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x43, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x70, 0x69, 0x72, 0x69, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c,
	0x0a, 0x07, 0x73, 0x70, 0x69, 0x72, 0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x69, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x69,
	0x72, 0x69, 0x74, 0x52, 0x07, 0x73, 0x70, 0x69, 0x72, 0x69, 0x74, 0x73, 0x22, 0x41, 0x0a, 0x13,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x70, 0x69, 0x72, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x70, 0x69, 0x72, 0x69, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x69, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x70, 0x69, 0x72, 0x69, 0x74, 0x52, 0x06, 0x73, 0x70, 0x69, 0x72, 0x69, 0x74, 0x22,
	0x42, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x70, 0x69, 0x72, 0x69, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x70, 0x69, 0x72, 0x69,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x69, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x69, 0x72, 0x69, 0x74, 0x52, 0x06, 0x73, 0x70, 0x69,
	0x72, 0x69, 0x74, 0x22, 0x25, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x70, 0x69,
	0x72, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x42, 0x0a, 0x14, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x53, 0x70, 0x69, 0x72, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x70, 0x69, 0x72, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x69, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x70, 0x69, 0x72, 0x69, 0x74, 0x52, 0x06, 0x73, 0x70, 0x69, 0x72, 0x69, 0x74, 0x22, 0xa5,
	0x01, 0x0a, 0x06, 0x53, 0x70, 0x69, 0x72, 0x69, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x6d, 0x65, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x69, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x69, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x70, 0x69, 0x72, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x73, 0x12, 0x32, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x69, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x70, 0x69, 0x72, 0x69, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xef, 0x01, 0x0a, 0x0b, 0x53, 0x70, 0x69, 0x72, 0x69,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x25,
	0x0a, 0x0e, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x70, 0x6f, 0x77, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c,
	0x50, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x33, 0x0a, 0x15, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61,
	0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x43, 0x6f,
	0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65,
	0x6e, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x2f, 0x0a,
	0x13, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x67, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x61, 0x67, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x22, 0x7d, 0x0a, 0x0c, 0x53, 0x70, 0x69, 0x72,
	0x69, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x09,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x08, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x06, 0x69,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x70,
	0x69, 0x72, 0x69, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48,
	0x00, 0x52, 0x06, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xdb, 0x04, 0x0a, 0x0d, 0x53, 0x70, 0x69, 0x72,
	0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x71, 0x0a, 0x0c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x70, 0x69, 0x72, 0x69, 0x74, 0x12, 0x1f, 0x2e, 0x73, 0x70, 0x69, 0x72,
	0x69, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x70, 0x69,
	0x72, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x70, 0x69,
	0x72, 0x69, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x70,
	0x69, 0x72, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x18, 0x22, 0x13, 0x2f, 0x73, 0x70, 0x69, 0x72, 0x69, 0x74, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x70, 0x69, 0x72, 0x69, 0x74, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x6c, 0x0a, 0x09,
	0x47, 0x65, 0x74, 0x53, 0x70, 0x69, 0x72, 0x69, 0x74, 0x12, 0x1c, 0x2e, 0x73, 0x70, 0x69, 0x72,
	0x69, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x70, 0x69, 0x72, 0x69, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x69, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x70, 0x69, 0x72, 0x69, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a,
	0x2f, 0x73, 0x70, 0x69, 0x72, 0x69, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x69, 0x64, 0x3d,
	0x73, 0x70, 0x69, 0x72, 0x69, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0x6b, 0x0a, 0x0b, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x70, 0x69, 0x72, 0x69, 0x74, 0x73, 0x12, 0x1e, 0x2e, 0x73, 0x70, 0x69, 0x72,
	0x69, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x70, 0x69, 0x72, 0x69,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x70, 0x69, 0x72,
	0x69, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x70, 0x69, 0x72, 0x69,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x15, 0x12, 0x13, 0x2f, 0x73, 0x70, 0x69, 0x72, 0x69, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x70, 0x69, 0x72, 0x69, 0x74, 0x73, 0x12, 0x84, 0x01, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x70, 0x69, 0x72, 0x69, 0x74, 0x12, 0x1f, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x69,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x70, 0x69, 0x72,
	0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x70, 0x69, 0x72,
	0x69, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x70, 0x69,
	0x72, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x31, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x2b, 0x1a, 0x26, 0x2f, 0x73, 0x70, 0x69, 0x72, 0x69, 0x74, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x7b, 0x73, 0x70, 0x69, 0x72, 0x69, 0x74, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x69, 0x64,
	0x3d, 0x73, 0x70, 0x69, 0x72, 0x69, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x75,
	0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x70, 0x69, 0x72, 0x69, 0x74, 0x12, 0x1f,
	0x2e, 0x73, 0x70, 0x69, 0x72, 0x69, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x53, 0x70, 0x69, 0x72, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x20, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x69, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x70, 0x69, 0x72, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x2a, 0x1a, 0x2f, 0x73, 0x70, 0x69, 0x72,
	0x69, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x69, 0x64, 0x3d, 0x73, 0x70, 0x69, 0x72, 0x69,
	0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x6b, 0x65, 0x65, 0x73, 0x6c, 0x65, 0x72, 0x2f, 0x73, 0x70,
	0x69, 0x72, 0x69, 0x74, 0x73, 0x3b, 0x73, 0x70, 0x69, 0x72, 0x69, 0x74, 0x73, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spirits_v1_spirit_proto_rawDescOnce sync.Once
	file_spirits_v1_spirit_proto_rawDescData = file_spirits_v1_spirit_proto_rawDesc
)

func file_spirits_v1_spirit_proto_rawDescGZIP() []byte {
	file_spirits_v1_spirit_proto_rawDescOnce.Do(func() {
		file_spirits_v1_spirit_proto_rawDescData = protoimpl.X.CompressGZIP(file_spirits_v1_spirit_proto_rawDescData)
	})
	return file_spirits_v1_spirit_proto_rawDescData
}

var file_spirits_v1_spirit_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_spirits_v1_spirit_proto_goTypes = []interface{}{
	(*CreateSpiritRequest)(nil),  // 0: spirits.v1.CreateSpiritRequest
	(*CreateSpiritResponse)(nil), // 1: spirits.v1.CreateSpiritResponse
	(*GetSpiritRequest)(nil),     // 2: spirits.v1.GetSpiritRequest
	(*GetSpiritResponse)(nil),    // 3: spirits.v1.GetSpiritResponse
	(*ListSpiritsRequest)(nil),   // 4: spirits.v1.ListSpiritsRequest
	(*ListSpiritsResponse)(nil),  // 5: spirits.v1.ListSpiritsResponse
	(*UpdateSpiritRequest)(nil),  // 6: spirits.v1.UpdateSpiritRequest
	(*UpdateSpiritResponse)(nil), // 7: spirits.v1.UpdateSpiritResponse
	(*DeleteSpiritRequest)(nil),  // 8: spirits.v1.DeleteSpiritRequest
	(*DeleteSpiritResponse)(nil), // 9: spirits.v1.DeleteSpiritResponse
	(*Spirit)(nil),               // 10: spirits.v1.Spirit
	(*SpiritStats)(nil),          // 11: spirits.v1.SpiritStats
	(*SpiritAction)(nil),         // 12: spirits.v1.SpiritAction
	(*Meta)(nil),                 // 13: spirits.v1.Meta
	(*Action)(nil),               // 14: spirits.v1.Action
}
var file_spirits_v1_spirit_proto_depIdxs = []int32{
	10, // 0: spirits.v1.CreateSpiritRequest.spirit:type_name -> spirits.v1.Spirit
	10, // 1: spirits.v1.CreateSpiritResponse.spirit:type_name -> spirits.v1.Spirit
	10, // 2: spirits.v1.GetSpiritResponse.spirit:type_name -> spirits.v1.Spirit
	10, // 3: spirits.v1.ListSpiritsResponse.spirits:type_name -> spirits.v1.Spirit
	10, // 4: spirits.v1.UpdateSpiritRequest.spirit:type_name -> spirits.v1.Spirit
	10, // 5: spirits.v1.UpdateSpiritResponse.spirit:type_name -> spirits.v1.Spirit
	10, // 6: spirits.v1.DeleteSpiritResponse.spirit:type_name -> spirits.v1.Spirit
	13, // 7: spirits.v1.Spirit.meta:type_name -> spirits.v1.Meta
	11, // 8: spirits.v1.Spirit.stats:type_name -> spirits.v1.SpiritStats
	12, // 9: spirits.v1.Spirit.actions:type_name -> spirits.v1.SpiritAction
	14, // 10: spirits.v1.SpiritAction.inline:type_name -> spirits.v1.Action
	0,  // 11: spirits.v1.SpiritService.CreateSpirit:input_type -> spirits.v1.CreateSpiritRequest
	2,  // 12: spirits.v1.SpiritService.GetSpirit:input_type -> spirits.v1.GetSpiritRequest
	4,  // 13: spirits.v1.SpiritService.ListSpirits:input_type -> spirits.v1.ListSpiritsRequest
	6,  // 14: spirits.v1.SpiritService.UpdateSpirit:input_type -> spirits.v1.UpdateSpiritRequest
	8,  // 15: spirits.v1.SpiritService.DeleteSpirit:input_type -> spirits.v1.DeleteSpiritRequest
	1,  // 16: spirits.v1.SpiritService.CreateSpirit:output_type -> spirits.v1.CreateSpiritResponse
	3,  // 17: spirits.v1.SpiritService.GetSpirit:output_type -> spirits.v1.GetSpiritResponse
	5,  // 18: spirits.v1.SpiritService.ListSpirits:output_type -> spirits.v1.ListSpiritsResponse
	7,  // 19: spirits.v1.SpiritService.UpdateSpirit:output_type -> spirits.v1.UpdateSpiritResponse
	9,  // 20: spirits.v1.SpiritService.DeleteSpirit:output_type -> spirits.v1.DeleteSpiritResponse
	16, // [16:21] is the sub-list for method output_type
	11, // [11:16] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_spirits_v1_spirit_proto_init() }
func file_spirits_v1_spirit_proto_init() {
	if File_spirits_v1_spirit_proto != nil {
		return
	}
	file_spirits_v1_action_proto_init()
	file_spirits_v1_meta_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_spirits_v1_spirit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSpiritRequest); i {
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
		file_spirits_v1_spirit_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSpiritResponse); i {
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
		file_spirits_v1_spirit_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSpiritRequest); i {
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
		file_spirits_v1_spirit_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSpiritResponse); i {
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
		file_spirits_v1_spirit_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSpiritsRequest); i {
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
		file_spirits_v1_spirit_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSpiritsResponse); i {
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
		file_spirits_v1_spirit_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSpiritRequest); i {
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
		file_spirits_v1_spirit_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSpiritResponse); i {
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
		file_spirits_v1_spirit_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSpiritRequest); i {
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
		file_spirits_v1_spirit_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSpiritResponse); i {
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
		file_spirits_v1_spirit_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Spirit); i {
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
		file_spirits_v1_spirit_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpiritStats); i {
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
		file_spirits_v1_spirit_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpiritAction); i {
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
	file_spirits_v1_spirit_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_spirits_v1_spirit_proto_msgTypes[12].OneofWrappers = []interface{}{
		(*SpiritAction_ActionId)(nil),
		(*SpiritAction_Inline)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spirits_v1_spirit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spirits_v1_spirit_proto_goTypes,
		DependencyIndexes: file_spirits_v1_spirit_proto_depIdxs,
		MessageInfos:      file_spirits_v1_spirit_proto_msgTypes,
	}.Build()
	File_spirits_v1_spirit_proto = out.File
	file_spirits_v1_spirit_proto_rawDesc = nil
	file_spirits_v1_spirit_proto_goTypes = nil
	file_spirits_v1_spirit_proto_depIdxs = nil
}
