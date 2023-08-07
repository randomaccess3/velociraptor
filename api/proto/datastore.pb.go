// Code generated by protoc-gen-go. DO NOT EDIT.
// source: datastore.proto

package proto

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

type DSPathSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Components []string `protobuf:"bytes,1,rep,name=components,proto3" json:"components,omitempty"`
	PathType   int64    `protobuf:"varint,2,opt,name=path_type,json=pathType,proto3" json:"path_type,omitempty"`
	IsDir      bool     `protobuf:"varint,3,opt,name=is_dir,json=isDir,proto3" json:"is_dir,omitempty"`
	Tag        string   `protobuf:"bytes,4,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *DSPathSpec) Reset() {
	*x = DSPathSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datastore_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DSPathSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DSPathSpec) ProtoMessage() {}

func (x *DSPathSpec) ProtoReflect() protoreflect.Message {
	mi := &file_datastore_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DSPathSpec.ProtoReflect.Descriptor instead.
func (*DSPathSpec) Descriptor() ([]byte, []int) {
	return file_datastore_proto_rawDescGZIP(), []int{0}
}

func (x *DSPathSpec) GetComponents() []string {
	if x != nil {
		return x.Components
	}
	return nil
}

func (x *DSPathSpec) GetPathType() int64 {
	if x != nil {
		return x.PathType
	}
	return 0
}

func (x *DSPathSpec) GetIsDir() bool {
	if x != nil {
		return x.IsDir
	}
	return false
}

func (x *DSPathSpec) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

type DataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pathspec *DSPathSpec `protobuf:"bytes,1,opt,name=pathspec,proto3" json:"pathspec,omitempty"`
	Data     []byte      `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	// If set the request will block until the data is committed to
	// disk.
	Sync  bool   `protobuf:"varint,3,opt,name=sync,proto3" json:"sync,omitempty"`
	OrgId string `protobuf:"bytes,4,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
}

func (x *DataRequest) Reset() {
	*x = DataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datastore_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataRequest) ProtoMessage() {}

func (x *DataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_datastore_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataRequest.ProtoReflect.Descriptor instead.
func (*DataRequest) Descriptor() ([]byte, []int) {
	return file_datastore_proto_rawDescGZIP(), []int{1}
}

func (x *DataRequest) GetPathspec() *DSPathSpec {
	if x != nil {
		return x.Pathspec
	}
	return nil
}

func (x *DataRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *DataRequest) GetSync() bool {
	if x != nil {
		return x.Sync
	}
	return false
}

func (x *DataRequest) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

type DataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *DataResponse) Reset() {
	*x = DataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datastore_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataResponse) ProtoMessage() {}

func (x *DataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_datastore_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataResponse.ProtoReflect.Descriptor instead.
func (*DataResponse) Descriptor() ([]byte, []int) {
	return file_datastore_proto_rawDescGZIP(), []int{2}
}

func (x *DataResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type ListChildrenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Children []*DSPathSpec `protobuf:"bytes,1,rep,name=children,proto3" json:"children,omitempty"`
}

func (x *ListChildrenResponse) Reset() {
	*x = ListChildrenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datastore_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListChildrenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListChildrenResponse) ProtoMessage() {}

func (x *ListChildrenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_datastore_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListChildrenResponse.ProtoReflect.Descriptor instead.
func (*ListChildrenResponse) Descriptor() ([]byte, []int) {
	return file_datastore_proto_rawDescGZIP(), []int{3}
}

func (x *ListChildrenResponse) GetChildren() []*DSPathSpec {
	if x != nil {
		return x.Children
	}
	return nil
}

var File_datastore_proto protoreflect.FileDescriptor

var file_datastore_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x72, 0x0a, 0x0a, 0x44, 0x53, 0x50, 0x61,
	0x74, 0x68, 0x53, 0x70, 0x65, 0x63, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x74, 0x68, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x73, 0x5f, 0x64, 0x69, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x44, 0x69, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61,
	0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x22, 0x7b, 0x0a, 0x0b,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x08, 0x70,
	0x61, 0x74, 0x68, 0x73, 0x70, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x53, 0x50, 0x61, 0x74, 0x68, 0x53, 0x70, 0x65, 0x63,
	0x52, 0x08, 0x70, 0x61, 0x74, 0x68, 0x73, 0x70, 0x65, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x79, 0x6e, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x73, 0x79,
	0x6e, 0x63, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x22, 0x22, 0x0a, 0x0c, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x45, 0x0a,
	0x14, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x44, 0x53, 0x50, 0x61, 0x74, 0x68, 0x53, 0x70, 0x65, 0x63, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c,
	0x64, 0x72, 0x65, 0x6e, 0x42, 0x31, 0x5a, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x76, 0x65, 0x6c, 0x6f,
	0x63, 0x69, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2f, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x72, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_datastore_proto_rawDescOnce sync.Once
	file_datastore_proto_rawDescData = file_datastore_proto_rawDesc
)

func file_datastore_proto_rawDescGZIP() []byte {
	file_datastore_proto_rawDescOnce.Do(func() {
		file_datastore_proto_rawDescData = protoimpl.X.CompressGZIP(file_datastore_proto_rawDescData)
	})
	return file_datastore_proto_rawDescData
}

var file_datastore_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_datastore_proto_goTypes = []interface{}{
	(*DSPathSpec)(nil),           // 0: proto.DSPathSpec
	(*DataRequest)(nil),          // 1: proto.DataRequest
	(*DataResponse)(nil),         // 2: proto.DataResponse
	(*ListChildrenResponse)(nil), // 3: proto.ListChildrenResponse
}
var file_datastore_proto_depIdxs = []int32{
	0, // 0: proto.DataRequest.pathspec:type_name -> proto.DSPathSpec
	0, // 1: proto.ListChildrenResponse.children:type_name -> proto.DSPathSpec
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_datastore_proto_init() }
func file_datastore_proto_init() {
	if File_datastore_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_datastore_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DSPathSpec); i {
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
		file_datastore_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataRequest); i {
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
		file_datastore_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataResponse); i {
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
		file_datastore_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListChildrenResponse); i {
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
			RawDescriptor: file_datastore_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_datastore_proto_goTypes,
		DependencyIndexes: file_datastore_proto_depIdxs,
		MessageInfos:      file_datastore_proto_msgTypes,
	}.Build()
	File_datastore_proto = out.File
	file_datastore_proto_rawDesc = nil
	file_datastore_proto_goTypes = nil
	file_datastore_proto_depIdxs = nil
}
