// Copyright 2017 Google Inc.
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
// 	protoc-gen-go v1.22.0
// 	protoc        v3.12.2
// source: google/watcher/v1/watch.proto

package watcher

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	_ "github.com/golang/protobuf/ptypes/empty"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// A reported value can be in one of the following states:
type Change_State int32

const (
	// The element exists and its full value is included in data.
	Change_EXISTS Change_State = 0
	// The element does not exist.
	Change_DOES_NOT_EXIST Change_State = 1
	// Element may or may not exist. Used only for initial state delivery when
	// the client is not interested in fetching the initial state. See the
	// "Initial State" section above.
	Change_INITIAL_STATE_SKIPPED Change_State = 2
	// The element may exist, but some error has occurred. More information is
	// available in the data field - the value is a serialized Status
	// proto (from [google.rpc.Status][])
	Change_ERROR Change_State = 3
)

// Enum value maps for Change_State.
var (
	Change_State_name = map[int32]string{
		0: "EXISTS",
		1: "DOES_NOT_EXIST",
		2: "INITIAL_STATE_SKIPPED",
		3: "ERROR",
	}
	Change_State_value = map[string]int32{
		"EXISTS":                0,
		"DOES_NOT_EXIST":        1,
		"INITIAL_STATE_SKIPPED": 2,
		"ERROR":                 3,
	}
)

func (x Change_State) Enum() *Change_State {
	p := new(Change_State)
	*p = x
	return p
}

func (x Change_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Change_State) Descriptor() protoreflect.EnumDescriptor {
	return file_google_watcher_v1_watch_proto_enumTypes[0].Descriptor()
}

func (Change_State) Type() protoreflect.EnumType {
	return &file_google_watcher_v1_watch_proto_enumTypes[0]
}

func (x Change_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Change_State.Descriptor instead.
func (Change_State) EnumDescriptor() ([]byte, []int) {
	return file_google_watcher_v1_watch_proto_rawDescGZIP(), []int{2, 0}
}

// The message used by the client to register interest in an entity.
type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The `target` value **must** be a valid URL path pointing to an entity
	// to watch. Note that the service name **must** be
	// removed from the target field (e.g., the target field must say
	// "/foo/bar", not "myservice.googleapis.com/foo/bar"). A client is
	// also allowed to pass system-specific parameters in the URL that
	// are only obeyed by some implementations. Some parameters will be
	// implementation-specific. However, some have predefined meaning
	// and are listed here:
	//
	//  * recursive = true|false [default=false]
	//    If set to true, indicates that the client wants to watch all elements
	//    of entities in the subtree rooted at the entity's name in `target`. For
	//    descendants that are not the immediate children of the target, the
	//    `Change.element` will contain slashes.
	//
	//    Note that some namespaces and entities will not support recursive
	//    watching. When watching such an entity, a client must not set recursive
	//    to true. Otherwise, it will receive an `UNIMPLEMENTED` error.
	//
	// Normal URL encoding must be used inside `target`.  For example, if a query
	// parameter name or value, or the non-query parameter portion of `target`
	// contains a special character, it must be %-encoded.  We recommend that
	// clients and servers use their runtime's URL library to produce and consume
	// target values.
	Target string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	// The `resume_marker` specifies how much of the existing underlying state is
	// delivered to the client when the watch request is received by the
	// system. The client can set this marker in one of the following ways to get
	// different semantics:
	//
	// *   Parameter is not specified or has the value "".
	//     Semantics: Fetch initial state.
	//     The client wants the entity's initial state to be delivered. See the
	//     description in "Initial State".
	//
	// *   Parameter is set to the string "now" (UTF-8 encoding).
	//     Semantics: Fetch new changes only.
	//     The client just wants to get the changes received by the system after
	//     the watch point. The system may deliver changes from before the watch
	//     point as well.
	//
	// *   Parameter is set to a value received in an earlier
	//     `Change.resume_marker` field while watching the same entity.
	//     Semantics: Resume from a specific point.
	//     The client wants to receive the changes from a specific point; this
	//     value must correspond to a value received in the `Change.resume_marker`
	//     field. The system may deliver changes from before the `resume_marker`
	//     as well. If the system cannot resume the stream from this point (e.g.,
	//     if it is too far behind in the stream), it can raise the
	//     `FAILED_PRECONDITION` error.
	//
	// An implementation MUST support an unspecified parameter and the
	// empty string "" marker (initial state fetching) and the "now" marker.
	// It need not support resuming from a specific point.
	ResumeMarker []byte `protobuf:"bytes,2,opt,name=resume_marker,json=resumeMarker,proto3" json:"resume_marker,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_watcher_v1_watch_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_google_watcher_v1_watch_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_google_watcher_v1_watch_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *Request) GetResumeMarker() []byte {
	if x != nil {
		return x.ResumeMarker
	}
	return nil
}

// A batch of Change messages.
type ChangeBatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of Change messages.
	Changes []*Change `protobuf:"bytes,1,rep,name=changes,proto3" json:"changes,omitempty"`
}

func (x *ChangeBatch) Reset() {
	*x = ChangeBatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_watcher_v1_watch_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeBatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeBatch) ProtoMessage() {}

func (x *ChangeBatch) ProtoReflect() protoreflect.Message {
	mi := &file_google_watcher_v1_watch_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeBatch.ProtoReflect.Descriptor instead.
func (*ChangeBatch) Descriptor() ([]byte, []int) {
	return file_google_watcher_v1_watch_proto_rawDescGZIP(), []int{1}
}

func (x *ChangeBatch) GetChanges() []*Change {
	if x != nil {
		return x.Changes
	}
	return nil
}

// A Change indicates the most recent state of an element.
type Change struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the element, interpreted relative to the entity's actual
	// name. "" refers to the entity itself. The element name is a valid
	// UTF-8 string.
	Element string `protobuf:"bytes,1,opt,name=element,proto3" json:"element,omitempty"`
	// The state of the `element`.
	State Change_State `protobuf:"varint,2,opt,name=state,proto3,enum=google.watcher.v1.Change_State" json:"state,omitempty"`
	// The actual change data. This field is present only when `state() == EXISTS`
	// or `state() == ERROR`. Please see
	// [google.protobuf.Any][google.protobuf.Any] about how to use the Any type.
	Data *any.Any `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	// If present, provides a compact representation of all the messages that have
	// been received by the caller for the given entity, e.g., it could be a
	// sequence number or a multi-part timestamp/version vector. This marker can
	// be provided in the Request message, allowing the caller to resume the
	// stream watching at a specific point without fetching the initial state.
	ResumeMarker []byte `protobuf:"bytes,4,opt,name=resume_marker,json=resumeMarker,proto3" json:"resume_marker,omitempty"`
	// If true, this Change is followed by more Changes that are in the same group
	// as this Change.
	Continued bool `protobuf:"varint,5,opt,name=continued,proto3" json:"continued,omitempty"`
}

func (x *Change) Reset() {
	*x = Change{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_watcher_v1_watch_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Change) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Change) ProtoMessage() {}

func (x *Change) ProtoReflect() protoreflect.Message {
	mi := &file_google_watcher_v1_watch_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Change.ProtoReflect.Descriptor instead.
func (*Change) Descriptor() ([]byte, []int) {
	return file_google_watcher_v1_watch_proto_rawDescGZIP(), []int{2}
}

func (x *Change) GetElement() string {
	if x != nil {
		return x.Element
	}
	return ""
}

func (x *Change) GetState() Change_State {
	if x != nil {
		return x.State
	}
	return Change_EXISTS
}

func (x *Change) GetData() *any.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Change) GetResumeMarker() []byte {
	if x != nil {
		return x.ResumeMarker
	}
	return nil
}

func (x *Change) GetContinued() bool {
	if x != nil {
		return x.Continued
	}
	return false
}

var File_google_watcher_v1_watch_proto protoreflect.FileDescriptor

var file_google_watcher_v1_watch_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x11, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x72,
	0x65, 0x73, 0x75, 0x6d, 0x65, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x72,
	0x22, 0x42, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12,
	0x33, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x07, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x73, 0x22, 0x95, 0x02, 0x0a, 0x06, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x35, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65,
	0x73, 0x75, 0x6d, 0x65, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x65, 0x64, 0x22, 0x4d, 0x0a,
	0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53,
	0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x44, 0x4f, 0x45, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45,
	0x58, 0x49, 0x53, 0x54, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x49, 0x4e, 0x49, 0x54, 0x49, 0x41,
	0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x53, 0x4b, 0x49, 0x50, 0x50, 0x45, 0x44, 0x10,
	0x02, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x32, 0x63, 0x0a, 0x07,
	0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x58, 0x0a, 0x05, 0x57, 0x61, 0x74, 0x63, 0x68,
	0x12, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x22, 0x11, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x30,
	0x01, 0x42, 0x5f, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x57, 0x61, 0x74, 0x63,
	0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_watcher_v1_watch_proto_rawDescOnce sync.Once
	file_google_watcher_v1_watch_proto_rawDescData = file_google_watcher_v1_watch_proto_rawDesc
)

func file_google_watcher_v1_watch_proto_rawDescGZIP() []byte {
	file_google_watcher_v1_watch_proto_rawDescOnce.Do(func() {
		file_google_watcher_v1_watch_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_watcher_v1_watch_proto_rawDescData)
	})
	return file_google_watcher_v1_watch_proto_rawDescData
}

var file_google_watcher_v1_watch_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_watcher_v1_watch_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_watcher_v1_watch_proto_goTypes = []interface{}{
	(Change_State)(0),   // 0: google.watcher.v1.Change.State
	(*Request)(nil),     // 1: google.watcher.v1.Request
	(*ChangeBatch)(nil), // 2: google.watcher.v1.ChangeBatch
	(*Change)(nil),      // 3: google.watcher.v1.Change
	(*any.Any)(nil),     // 4: google.protobuf.Any
}
var file_google_watcher_v1_watch_proto_depIdxs = []int32{
	3, // 0: google.watcher.v1.ChangeBatch.changes:type_name -> google.watcher.v1.Change
	0, // 1: google.watcher.v1.Change.state:type_name -> google.watcher.v1.Change.State
	4, // 2: google.watcher.v1.Change.data:type_name -> google.protobuf.Any
	1, // 3: google.watcher.v1.Watcher.Watch:input_type -> google.watcher.v1.Request
	2, // 4: google.watcher.v1.Watcher.Watch:output_type -> google.watcher.v1.ChangeBatch
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_watcher_v1_watch_proto_init() }
func file_google_watcher_v1_watch_proto_init() {
	if File_google_watcher_v1_watch_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_watcher_v1_watch_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_google_watcher_v1_watch_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeBatch); i {
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
		file_google_watcher_v1_watch_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Change); i {
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
			RawDescriptor: file_google_watcher_v1_watch_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_watcher_v1_watch_proto_goTypes,
		DependencyIndexes: file_google_watcher_v1_watch_proto_depIdxs,
		EnumInfos:         file_google_watcher_v1_watch_proto_enumTypes,
		MessageInfos:      file_google_watcher_v1_watch_proto_msgTypes,
	}.Build()
	File_google_watcher_v1_watch_proto = out.File
	file_google_watcher_v1_watch_proto_rawDesc = nil
	file_google_watcher_v1_watch_proto_goTypes = nil
	file_google_watcher_v1_watch_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WatcherClient is the client API for Watcher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WatcherClient interface {
	// Start a streaming RPC to get watch information from the server.
	Watch(ctx context.Context, in *Request, opts ...grpc.CallOption) (Watcher_WatchClient, error)
}

type watcherClient struct {
	cc grpc.ClientConnInterface
}

func NewWatcherClient(cc grpc.ClientConnInterface) WatcherClient {
	return &watcherClient{cc}
}

func (c *watcherClient) Watch(ctx context.Context, in *Request, opts ...grpc.CallOption) (Watcher_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Watcher_serviceDesc.Streams[0], "/google.watcher.v1.Watcher/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &watcherWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Watcher_WatchClient interface {
	Recv() (*ChangeBatch, error)
	grpc.ClientStream
}

type watcherWatchClient struct {
	grpc.ClientStream
}

func (x *watcherWatchClient) Recv() (*ChangeBatch, error) {
	m := new(ChangeBatch)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WatcherServer is the server API for Watcher service.
type WatcherServer interface {
	// Start a streaming RPC to get watch information from the server.
	Watch(*Request, Watcher_WatchServer) error
}

// UnimplementedWatcherServer can be embedded to have forward compatible implementations.
type UnimplementedWatcherServer struct {
}

func (*UnimplementedWatcherServer) Watch(*Request, Watcher_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}

func RegisterWatcherServer(s *grpc.Server, srv WatcherServer) {
	s.RegisterService(&_Watcher_serviceDesc, srv)
}

func _Watcher_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WatcherServer).Watch(m, &watcherWatchServer{stream})
}

type Watcher_WatchServer interface {
	Send(*ChangeBatch) error
	grpc.ServerStream
}

type watcherWatchServer struct {
	grpc.ServerStream
}

func (x *watcherWatchServer) Send(m *ChangeBatch) error {
	return x.ServerStream.SendMsg(m)
}

var _Watcher_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.watcher.v1.Watcher",
	HandlerType: (*WatcherServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _Watcher_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "google/watcher/v1/watch.proto",
}