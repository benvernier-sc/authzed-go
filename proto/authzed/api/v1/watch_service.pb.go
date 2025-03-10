// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        (unknown)
// source: authzed/api/v1/watch_service.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// WatchRequest specifies the object definitions for which we want to start
// watching mutations, and an optional start snapshot for when to start
// watching.
type WatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// optional_object_types is a filter of resource object types to watch for changes.
	// If specified, only changes to the specified object types will be returned and
	// optional_relationship_filters cannot be used.
	OptionalObjectTypes []string `protobuf:"bytes,1,rep,name=optional_object_types,json=optionalObjectTypes,proto3" json:"optional_object_types,omitempty"`
	// optional_start_cursor is the ZedToken holding the point-in-time at
	// which to start watching for changes.
	// If not specified, the watch will begin at the current head revision
	// of the datastore, returning any updates that occur after the caller
	// makes the request.
	// Note that if this cursor references a point-in-time containing data
	// that has been garbage collected, an error will be returned.
	OptionalStartCursor *ZedToken `protobuf:"bytes,2,opt,name=optional_start_cursor,json=optionalStartCursor,proto3" json:"optional_start_cursor,omitempty"`
	// optional_relationship_filters, if specified, indicates the
	// filter(s) to apply to each relationship to be returned by watch.
	// The relationship will be returned as long as at least one filter matches,
	// this allows clients to match relationships on multiple filters on a single watch call.
	// If specified, optional_object_types cannot be used.
	OptionalRelationshipFilters []*RelationshipFilter `protobuf:"bytes,3,rep,name=optional_relationship_filters,json=optionalRelationshipFilters,proto3" json:"optional_relationship_filters,omitempty"`
}

func (x *WatchRequest) Reset() {
	*x = WatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authzed_api_v1_watch_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchRequest) ProtoMessage() {}

func (x *WatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authzed_api_v1_watch_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchRequest.ProtoReflect.Descriptor instead.
func (*WatchRequest) Descriptor() ([]byte, []int) {
	return file_authzed_api_v1_watch_service_proto_rawDescGZIP(), []int{0}
}

func (x *WatchRequest) GetOptionalObjectTypes() []string {
	if x != nil {
		return x.OptionalObjectTypes
	}
	return nil
}

func (x *WatchRequest) GetOptionalStartCursor() *ZedToken {
	if x != nil {
		return x.OptionalStartCursor
	}
	return nil
}

func (x *WatchRequest) GetOptionalRelationshipFilters() []*RelationshipFilter {
	if x != nil {
		return x.OptionalRelationshipFilters
	}
	return nil
}

// WatchResponse contains all tuple modification events in ascending
// timestamp order, from the requested start snapshot to a snapshot
// encoded in the watch response. The client can use the snapshot to resume
// watching where the previous watch response left off.
type WatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Updates        []*RelationshipUpdate `protobuf:"bytes,1,rep,name=updates,proto3" json:"updates,omitempty"`
	ChangesThrough *ZedToken             `protobuf:"bytes,2,opt,name=changes_through,json=changesThrough,proto3" json:"changes_through,omitempty"`
}

func (x *WatchResponse) Reset() {
	*x = WatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authzed_api_v1_watch_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchResponse) ProtoMessage() {}

func (x *WatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authzed_api_v1_watch_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchResponse.ProtoReflect.Descriptor instead.
func (*WatchResponse) Descriptor() ([]byte, []int) {
	return file_authzed_api_v1_watch_service_proto_rawDescGZIP(), []int{1}
}

func (x *WatchResponse) GetUpdates() []*RelationshipUpdate {
	if x != nil {
		return x.Updates
	}
	return nil
}

func (x *WatchResponse) GetChangesThrough() *ZedToken {
	if x != nil {
		return x.ChangesThrough
	}
	return nil
}

var File_authzed_api_v1_watch_service_proto protoreflect.FileDescriptor

var file_authzed_api_v1_watch_service_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x61, 0x75, 0x74,
	0x68, 0x7a, 0x65, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xca, 0x02, 0x0a, 0x0c, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x83, 0x01, 0x0a, 0x15, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x42, 0x4f, 0xfa, 0x42, 0x4c, 0x92, 0x01, 0x49, 0x08, 0x00, 0x22, 0x45, 0x72, 0x43, 0x28, 0x80,
	0x01, 0x32, 0x3e, 0x5e, 0x28, 0x5b, 0x61, 0x2d, 0x7a, 0x5d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d,
	0x39, 0x5f, 0x5d, 0x7b, 0x31, 0x2c, 0x36, 0x32, 0x7d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39,
	0x5d, 0x2f, 0x29, 0x2a, 0x5b, 0x61, 0x2d, 0x7a, 0x5d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39,
	0x5f, 0x5d, 0x7b, 0x31, 0x2c, 0x36, 0x32, 0x7d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d,
	0x24, 0x52, 0x13, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x4c, 0x0a, 0x15, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x5a, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x13, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x43, 0x75,
	0x72, 0x73, 0x6f, 0x72, 0x12, 0x66, 0x0a, 0x1d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52,
	0x1b, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x68, 0x69, 0x70, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x22, 0x90, 0x01, 0x0a,
	0x0d, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c,
	0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x12, 0x41, 0x0a, 0x0f,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x5f, 0x74, 0x68, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x5a, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x0e, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x54, 0x68, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x32,
	0x6c, 0x0a, 0x0c, 0x57, 0x61, 0x74, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x5c, 0x0a, 0x05, 0x57, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a,
	0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x01, 0x2a,
	0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x30, 0x01, 0x42, 0x4a, 0x0a,
	0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65,
	0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a,
	0x65, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_authzed_api_v1_watch_service_proto_rawDescOnce sync.Once
	file_authzed_api_v1_watch_service_proto_rawDescData = file_authzed_api_v1_watch_service_proto_rawDesc
)

func file_authzed_api_v1_watch_service_proto_rawDescGZIP() []byte {
	file_authzed_api_v1_watch_service_proto_rawDescOnce.Do(func() {
		file_authzed_api_v1_watch_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_authzed_api_v1_watch_service_proto_rawDescData)
	})
	return file_authzed_api_v1_watch_service_proto_rawDescData
}

var file_authzed_api_v1_watch_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_authzed_api_v1_watch_service_proto_goTypes = []interface{}{
	(*WatchRequest)(nil),       // 0: authzed.api.v1.WatchRequest
	(*WatchResponse)(nil),      // 1: authzed.api.v1.WatchResponse
	(*ZedToken)(nil),           // 2: authzed.api.v1.ZedToken
	(*RelationshipFilter)(nil), // 3: authzed.api.v1.RelationshipFilter
	(*RelationshipUpdate)(nil), // 4: authzed.api.v1.RelationshipUpdate
}
var file_authzed_api_v1_watch_service_proto_depIdxs = []int32{
	2, // 0: authzed.api.v1.WatchRequest.optional_start_cursor:type_name -> authzed.api.v1.ZedToken
	3, // 1: authzed.api.v1.WatchRequest.optional_relationship_filters:type_name -> authzed.api.v1.RelationshipFilter
	4, // 2: authzed.api.v1.WatchResponse.updates:type_name -> authzed.api.v1.RelationshipUpdate
	2, // 3: authzed.api.v1.WatchResponse.changes_through:type_name -> authzed.api.v1.ZedToken
	0, // 4: authzed.api.v1.WatchService.Watch:input_type -> authzed.api.v1.WatchRequest
	1, // 5: authzed.api.v1.WatchService.Watch:output_type -> authzed.api.v1.WatchResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_authzed_api_v1_watch_service_proto_init() }
func file_authzed_api_v1_watch_service_proto_init() {
	if File_authzed_api_v1_watch_service_proto != nil {
		return
	}
	file_authzed_api_v1_core_proto_init()
	file_authzed_api_v1_permission_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_authzed_api_v1_watch_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchRequest); i {
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
		file_authzed_api_v1_watch_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchResponse); i {
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
			RawDescriptor: file_authzed_api_v1_watch_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_authzed_api_v1_watch_service_proto_goTypes,
		DependencyIndexes: file_authzed_api_v1_watch_service_proto_depIdxs,
		MessageInfos:      file_authzed_api_v1_watch_service_proto_msgTypes,
	}.Build()
	File_authzed_api_v1_watch_service_proto = out.File
	file_authzed_api_v1_watch_service_proto_rawDesc = nil
	file_authzed_api_v1_watch_service_proto_goTypes = nil
	file_authzed_api_v1_watch_service_proto_depIdxs = nil
}
