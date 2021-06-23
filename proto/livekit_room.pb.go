// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: livekit_room.proto

package livekit

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

type CreateRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// number of seconds the room should cleanup after being empty
	EmptyTimeout    uint32 `protobuf:"varint,2,opt,name=empty_timeout,json=emptyTimeout,proto3" json:"empty_timeout,omitempty"`
	MaxParticipants uint32 `protobuf:"varint,3,opt,name=max_participants,json=maxParticipants,proto3" json:"max_participants,omitempty"`
	// override the node room is allocated to, for debugging
	NodeId string `protobuf:"bytes,4,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
}

func (x *CreateRoomRequest) Reset() {
	*x = CreateRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_room_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoomRequest) ProtoMessage() {}

func (x *CreateRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_room_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoomRequest.ProtoReflect.Descriptor instead.
func (*CreateRoomRequest) Descriptor() ([]byte, []int) {
	return file_livekit_room_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRoomRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateRoomRequest) GetEmptyTimeout() uint32 {
	if x != nil {
		return x.EmptyTimeout
	}
	return 0
}

func (x *CreateRoomRequest) GetMaxParticipants() uint32 {
	if x != nil {
		return x.MaxParticipants
	}
	return 0
}

func (x *CreateRoomRequest) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

type ListRoomsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListRoomsRequest) Reset() {
	*x = ListRoomsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_room_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRoomsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRoomsRequest) ProtoMessage() {}

func (x *ListRoomsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_room_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRoomsRequest.ProtoReflect.Descriptor instead.
func (*ListRoomsRequest) Descriptor() ([]byte, []int) {
	return file_livekit_room_proto_rawDescGZIP(), []int{1}
}

type ListRoomsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rooms []*Room `protobuf:"bytes,1,rep,name=rooms,proto3" json:"rooms,omitempty"`
}

func (x *ListRoomsResponse) Reset() {
	*x = ListRoomsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_room_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRoomsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRoomsResponse) ProtoMessage() {}

func (x *ListRoomsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_room_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRoomsResponse.ProtoReflect.Descriptor instead.
func (*ListRoomsResponse) Descriptor() ([]byte, []int) {
	return file_livekit_room_proto_rawDescGZIP(), []int{2}
}

func (x *ListRoomsResponse) GetRooms() []*Room {
	if x != nil {
		return x.Rooms
	}
	return nil
}

type DeleteRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Room string `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
}

func (x *DeleteRoomRequest) Reset() {
	*x = DeleteRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_room_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRoomRequest) ProtoMessage() {}

func (x *DeleteRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_room_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRoomRequest.ProtoReflect.Descriptor instead.
func (*DeleteRoomRequest) Descriptor() ([]byte, []int) {
	return file_livekit_room_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteRoomRequest) GetRoom() string {
	if x != nil {
		return x.Room
	}
	return ""
}

type DeleteRoomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteRoomResponse) Reset() {
	*x = DeleteRoomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_room_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRoomResponse) ProtoMessage() {}

func (x *DeleteRoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_room_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRoomResponse.ProtoReflect.Descriptor instead.
func (*DeleteRoomResponse) Descriptor() ([]byte, []int) {
	return file_livekit_room_proto_rawDescGZIP(), []int{4}
}

type ListParticipantsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Room string `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
}

func (x *ListParticipantsRequest) Reset() {
	*x = ListParticipantsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_room_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListParticipantsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListParticipantsRequest) ProtoMessage() {}

func (x *ListParticipantsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_room_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListParticipantsRequest.ProtoReflect.Descriptor instead.
func (*ListParticipantsRequest) Descriptor() ([]byte, []int) {
	return file_livekit_room_proto_rawDescGZIP(), []int{5}
}

func (x *ListParticipantsRequest) GetRoom() string {
	if x != nil {
		return x.Room
	}
	return ""
}

type ListParticipantsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Participants []*ParticipantInfo `protobuf:"bytes,1,rep,name=participants,proto3" json:"participants,omitempty"`
}

func (x *ListParticipantsResponse) Reset() {
	*x = ListParticipantsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_room_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListParticipantsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListParticipantsResponse) ProtoMessage() {}

func (x *ListParticipantsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_room_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListParticipantsResponse.ProtoReflect.Descriptor instead.
func (*ListParticipantsResponse) Descriptor() ([]byte, []int) {
	return file_livekit_room_proto_rawDescGZIP(), []int{6}
}

func (x *ListParticipantsResponse) GetParticipants() []*ParticipantInfo {
	if x != nil {
		return x.Participants
	}
	return nil
}

type RoomParticipantIdentity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Room     string `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
	Identity string `protobuf:"bytes,2,opt,name=identity,proto3" json:"identity,omitempty"`
}

func (x *RoomParticipantIdentity) Reset() {
	*x = RoomParticipantIdentity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_room_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomParticipantIdentity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomParticipantIdentity) ProtoMessage() {}

func (x *RoomParticipantIdentity) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_room_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomParticipantIdentity.ProtoReflect.Descriptor instead.
func (*RoomParticipantIdentity) Descriptor() ([]byte, []int) {
	return file_livekit_room_proto_rawDescGZIP(), []int{7}
}

func (x *RoomParticipantIdentity) GetRoom() string {
	if x != nil {
		return x.Room
	}
	return ""
}

func (x *RoomParticipantIdentity) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

type RemoveParticipantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveParticipantResponse) Reset() {
	*x = RemoveParticipantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_room_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveParticipantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveParticipantResponse) ProtoMessage() {}

func (x *RemoveParticipantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_room_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveParticipantResponse.ProtoReflect.Descriptor instead.
func (*RemoveParticipantResponse) Descriptor() ([]byte, []int) {
	return file_livekit_room_proto_rawDescGZIP(), []int{8}
}

type MuteRoomTrackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Room     string `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
	Identity string `protobuf:"bytes,2,opt,name=identity,proto3" json:"identity,omitempty"`
	TrackSid string `protobuf:"bytes,3,opt,name=track_sid,json=trackSid,proto3" json:"track_sid,omitempty"`
	Muted    bool   `protobuf:"varint,4,opt,name=muted,proto3" json:"muted,omitempty"`
}

func (x *MuteRoomTrackRequest) Reset() {
	*x = MuteRoomTrackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_room_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MuteRoomTrackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MuteRoomTrackRequest) ProtoMessage() {}

func (x *MuteRoomTrackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_room_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MuteRoomTrackRequest.ProtoReflect.Descriptor instead.
func (*MuteRoomTrackRequest) Descriptor() ([]byte, []int) {
	return file_livekit_room_proto_rawDescGZIP(), []int{9}
}

func (x *MuteRoomTrackRequest) GetRoom() string {
	if x != nil {
		return x.Room
	}
	return ""
}

func (x *MuteRoomTrackRequest) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *MuteRoomTrackRequest) GetTrackSid() string {
	if x != nil {
		return x.TrackSid
	}
	return ""
}

func (x *MuteRoomTrackRequest) GetMuted() bool {
	if x != nil {
		return x.Muted
	}
	return false
}

type MuteRoomTrackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Track *TrackInfo `protobuf:"bytes,1,opt,name=track,proto3" json:"track,omitempty"`
}

func (x *MuteRoomTrackResponse) Reset() {
	*x = MuteRoomTrackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_room_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MuteRoomTrackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MuteRoomTrackResponse) ProtoMessage() {}

func (x *MuteRoomTrackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_room_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MuteRoomTrackResponse.ProtoReflect.Descriptor instead.
func (*MuteRoomTrackResponse) Descriptor() ([]byte, []int) {
	return file_livekit_room_proto_rawDescGZIP(), []int{10}
}

func (x *MuteRoomTrackResponse) GetTrack() *TrackInfo {
	if x != nil {
		return x.Track
	}
	return nil
}

type ParticipantPermission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CanSubscribe bool `protobuf:"varint,1,opt,name=can_subscribe,json=canSubscribe,proto3" json:"can_subscribe,omitempty"`
	CanPublish   bool `protobuf:"varint,2,opt,name=can_publish,json=canPublish,proto3" json:"can_publish,omitempty"`
}

func (x *ParticipantPermission) Reset() {
	*x = ParticipantPermission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_room_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParticipantPermission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParticipantPermission) ProtoMessage() {}

func (x *ParticipantPermission) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_room_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParticipantPermission.ProtoReflect.Descriptor instead.
func (*ParticipantPermission) Descriptor() ([]byte, []int) {
	return file_livekit_room_proto_rawDescGZIP(), []int{11}
}

func (x *ParticipantPermission) GetCanSubscribe() bool {
	if x != nil {
		return x.CanSubscribe
	}
	return false
}

func (x *ParticipantPermission) GetCanPublish() bool {
	if x != nil {
		return x.CanPublish
	}
	return false
}

type UpdateParticipantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Room       string                 `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
	Identity   string                 `protobuf:"bytes,2,opt,name=identity,proto3" json:"identity,omitempty"`
	Metadata   string                 `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Permission *ParticipantPermission `protobuf:"bytes,4,opt,name=permission,proto3" json:"permission,omitempty"`
}

func (x *UpdateParticipantRequest) Reset() {
	*x = UpdateParticipantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_room_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateParticipantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateParticipantRequest) ProtoMessage() {}

func (x *UpdateParticipantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_room_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateParticipantRequest.ProtoReflect.Descriptor instead.
func (*UpdateParticipantRequest) Descriptor() ([]byte, []int) {
	return file_livekit_room_proto_rawDescGZIP(), []int{12}
}

func (x *UpdateParticipantRequest) GetRoom() string {
	if x != nil {
		return x.Room
	}
	return ""
}

func (x *UpdateParticipantRequest) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *UpdateParticipantRequest) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

func (x *UpdateParticipantRequest) GetPermission() *ParticipantPermission {
	if x != nil {
		return x.Permission
	}
	return nil
}

var File_livekit_room_proto protoreflect.FileDescriptor

var file_livekit_room_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x1a, 0x14, 0x6c,
	0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x90, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f,
	0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a,
	0x0d, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x6d, 0x61, 0x78, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x6d, 0x61,
	0x78, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x17, 0x0a,
	0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f,
	0x6f, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x38, 0x0a, 0x11, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x23, 0x0a, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x05, 0x72,
	0x6f, 0x6f, 0x6d, 0x73, 0x22, 0x27, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f,
	0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6f,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x22, 0x14, 0x0a,
	0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x2d, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f,
	0x6f, 0x6d, 0x22, 0x58, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c,
	0x0a, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x50,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x22, 0x49, 0x0a, 0x17,
	0x52, 0x6f, 0x6f, 0x6d, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x1b, 0x0a, 0x19, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x79, 0x0a, 0x14, 0x4d, 0x75, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d,
	0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1b, 0x0a, 0x09,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x53, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x75, 0x74,
	0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x6d, 0x75, 0x74, 0x65, 0x64, 0x22,
	0x41, 0x0a, 0x15, 0x4d, 0x75, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x72, 0x61, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69,
	0x74, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x22, 0x5d, 0x0a, 0x15, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e,
	0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x63,
	0x61, 0x6e, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0c, 0x63, 0x61, 0x6e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x6e, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x63, 0x61, 0x6e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x22, 0xa6, 0x01, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f,
	0x6f, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1a,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3e, 0x0a, 0x0a, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69,
	0x70, 0x61, 0x6e, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0a,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0xfa, 0x04, 0x0a, 0x0b, 0x52,
	0x6f, 0x6f, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x1a, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b,
	0x69, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x52,
	0x6f, 0x6f, 0x6d, 0x12, 0x42, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x73,
	0x12, 0x19, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x6f, 0x6f, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6c, 0x69,
	0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x1a, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57,
	0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e,
	0x74, 0x73, 0x12, 0x20, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x12, 0x20, 0x2e, 0x6c, 0x69, 0x76, 0x65,
	0x6b, 0x69, 0x74, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
	0x61, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x6c, 0x69,
	0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x59, 0x0a, 0x11, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x12, 0x20, 0x2e, 0x6c, 0x69, 0x76,
	0x65, 0x6b, 0x69, 0x74, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69,
	0x70, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x22, 0x2e, 0x6c,
	0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x53, 0x0a, 0x12, 0x4d, 0x75, 0x74, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65,
	0x64, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x12, 0x1d, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74,
	0x2e, 0x4d, 0x75, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e,
	0x4d, 0x75, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x12, 0x21, 0x2e, 0x6c, 0x69, 0x76,
	0x65, 0x6b, 0x69, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
	0x61, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2f, 0x6c, 0x69,
	0x76, 0x65, 0x6b, 0x69, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_livekit_room_proto_rawDescOnce sync.Once
	file_livekit_room_proto_rawDescData = file_livekit_room_proto_rawDesc
)

func file_livekit_room_proto_rawDescGZIP() []byte {
	file_livekit_room_proto_rawDescOnce.Do(func() {
		file_livekit_room_proto_rawDescData = protoimpl.X.CompressGZIP(file_livekit_room_proto_rawDescData)
	})
	return file_livekit_room_proto_rawDescData
}

var file_livekit_room_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_livekit_room_proto_goTypes = []interface{}{
	(*CreateRoomRequest)(nil),         // 0: livekit.CreateRoomRequest
	(*ListRoomsRequest)(nil),          // 1: livekit.ListRoomsRequest
	(*ListRoomsResponse)(nil),         // 2: livekit.ListRoomsResponse
	(*DeleteRoomRequest)(nil),         // 3: livekit.DeleteRoomRequest
	(*DeleteRoomResponse)(nil),        // 4: livekit.DeleteRoomResponse
	(*ListParticipantsRequest)(nil),   // 5: livekit.ListParticipantsRequest
	(*ListParticipantsResponse)(nil),  // 6: livekit.ListParticipantsResponse
	(*RoomParticipantIdentity)(nil),   // 7: livekit.RoomParticipantIdentity
	(*RemoveParticipantResponse)(nil), // 8: livekit.RemoveParticipantResponse
	(*MuteRoomTrackRequest)(nil),      // 9: livekit.MuteRoomTrackRequest
	(*MuteRoomTrackResponse)(nil),     // 10: livekit.MuteRoomTrackResponse
	(*ParticipantPermission)(nil),     // 11: livekit.ParticipantPermission
	(*UpdateParticipantRequest)(nil),  // 12: livekit.UpdateParticipantRequest
	(*Room)(nil),                      // 13: livekit.Room
	(*ParticipantInfo)(nil),           // 14: livekit.ParticipantInfo
	(*TrackInfo)(nil),                 // 15: livekit.TrackInfo
}
var file_livekit_room_proto_depIdxs = []int32{
	13, // 0: livekit.ListRoomsResponse.rooms:type_name -> livekit.Room
	14, // 1: livekit.ListParticipantsResponse.participants:type_name -> livekit.ParticipantInfo
	15, // 2: livekit.MuteRoomTrackResponse.track:type_name -> livekit.TrackInfo
	11, // 3: livekit.UpdateParticipantRequest.permission:type_name -> livekit.ParticipantPermission
	0,  // 4: livekit.RoomService.CreateRoom:input_type -> livekit.CreateRoomRequest
	1,  // 5: livekit.RoomService.ListRooms:input_type -> livekit.ListRoomsRequest
	3,  // 6: livekit.RoomService.DeleteRoom:input_type -> livekit.DeleteRoomRequest
	5,  // 7: livekit.RoomService.ListParticipants:input_type -> livekit.ListParticipantsRequest
	7,  // 8: livekit.RoomService.GetParticipant:input_type -> livekit.RoomParticipantIdentity
	7,  // 9: livekit.RoomService.RemoveParticipant:input_type -> livekit.RoomParticipantIdentity
	9,  // 10: livekit.RoomService.MutePublishedTrack:input_type -> livekit.MuteRoomTrackRequest
	12, // 11: livekit.RoomService.UpdateParticipant:input_type -> livekit.UpdateParticipantRequest
	13, // 12: livekit.RoomService.CreateRoom:output_type -> livekit.Room
	2,  // 13: livekit.RoomService.ListRooms:output_type -> livekit.ListRoomsResponse
	4,  // 14: livekit.RoomService.DeleteRoom:output_type -> livekit.DeleteRoomResponse
	6,  // 15: livekit.RoomService.ListParticipants:output_type -> livekit.ListParticipantsResponse
	14, // 16: livekit.RoomService.GetParticipant:output_type -> livekit.ParticipantInfo
	8,  // 17: livekit.RoomService.RemoveParticipant:output_type -> livekit.RemoveParticipantResponse
	10, // 18: livekit.RoomService.MutePublishedTrack:output_type -> livekit.MuteRoomTrackResponse
	14, // 19: livekit.RoomService.UpdateParticipant:output_type -> livekit.ParticipantInfo
	12, // [12:20] is the sub-list for method output_type
	4,  // [4:12] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_livekit_room_proto_init() }
func file_livekit_room_proto_init() {
	if File_livekit_room_proto != nil {
		return
	}
	file_livekit_models_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_livekit_room_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoomRequest); i {
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
		file_livekit_room_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRoomsRequest); i {
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
		file_livekit_room_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRoomsResponse); i {
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
		file_livekit_room_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRoomRequest); i {
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
		file_livekit_room_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRoomResponse); i {
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
		file_livekit_room_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListParticipantsRequest); i {
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
		file_livekit_room_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListParticipantsResponse); i {
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
		file_livekit_room_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomParticipantIdentity); i {
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
		file_livekit_room_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveParticipantResponse); i {
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
		file_livekit_room_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MuteRoomTrackRequest); i {
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
		file_livekit_room_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MuteRoomTrackResponse); i {
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
		file_livekit_room_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParticipantPermission); i {
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
		file_livekit_room_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateParticipantRequest); i {
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
			RawDescriptor: file_livekit_room_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_livekit_room_proto_goTypes,
		DependencyIndexes: file_livekit_room_proto_depIdxs,
		MessageInfos:      file_livekit_room_proto_msgTypes,
	}.Build()
	File_livekit_room_proto = out.File
	file_livekit_room_proto_rawDesc = nil
	file_livekit_room_proto_goTypes = nil
	file_livekit_room_proto_depIdxs = nil
}
