// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: setattr/proto/setattr.proto

package setattr

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	proto "yrcf_grpc_case/grpc/common/proto"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// message for rpc SetAttr
type SetAttrRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UseAbsolutePath bool   `protobuf:"varint,1,opt,name=use_absolute_path,json=useAbsolutePath,proto3" json:"use_absolute_path,omitempty"`
	Path            string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	// no action with UINT32_MAX
	AccessMode string `protobuf:"bytes,3,opt,name=access_mode,json=accessMode,proto3" json:"access_mode,omitempty"`
	Uid        uint32 `protobuf:"varint,4,opt,name=uid,proto3" json:"uid,omitempty"`
	Gid        uint32 `protobuf:"varint,5,opt,name=gid,proto3" json:"gid,omitempty"`
	// UNIX timestamps, for each pair, no action with any one assigned -1
	MtimeSecs *int64 `protobuf:"varint,6,opt,name=mtime_secs,json=mtimeSecs,proto3,oneof" json:"mtime_secs,omitempty"`
	MtimeNano *int64 `protobuf:"varint,7,opt,name=mtime_nano,json=mtimeNano,proto3,oneof" json:"mtime_nano,omitempty"`
	AtimeSecs *int64 `protobuf:"varint,8,opt,name=atime_secs,json=atimeSecs,proto3,oneof" json:"atime_secs,omitempty"`
	AtimeNano *int64 `protobuf:"varint,9,opt,name=atime_nano,json=atimeNano,proto3,oneof" json:"atime_nano,omitempty"`
}

func (x *SetAttrRequest) Reset() {
	*x = SetAttrRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_setattr_proto_setattr_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetAttrRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAttrRequest) ProtoMessage() {}

func (x *SetAttrRequest) ProtoReflect() protoreflect.Message {
	mi := &file_setattr_proto_setattr_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAttrRequest.ProtoReflect.Descriptor instead.
func (*SetAttrRequest) Descriptor() ([]byte, []int) {
	return file_setattr_proto_setattr_proto_rawDescGZIP(), []int{0}
}

func (x *SetAttrRequest) GetUseAbsolutePath() bool {
	if x != nil {
		return x.UseAbsolutePath
	}
	return false
}

func (x *SetAttrRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *SetAttrRequest) GetAccessMode() string {
	if x != nil {
		return x.AccessMode
	}
	return ""
}

func (x *SetAttrRequest) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *SetAttrRequest) GetGid() uint32 {
	if x != nil {
		return x.Gid
	}
	return 0
}

func (x *SetAttrRequest) GetMtimeSecs() int64 {
	if x != nil && x.MtimeSecs != nil {
		return *x.MtimeSecs
	}
	return 0
}

func (x *SetAttrRequest) GetMtimeNano() int64 {
	if x != nil && x.MtimeNano != nil {
		return *x.MtimeNano
	}
	return 0
}

func (x *SetAttrRequest) GetAtimeSecs() int64 {
	if x != nil && x.AtimeSecs != nil {
		return *x.AtimeSecs
	}
	return 0
}

func (x *SetAttrRequest) GetAtimeNano() int64 {
	if x != nil && x.AtimeNano != nil {
		return *x.AtimeNano
	}
	return 0
}

type SetAttrResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *proto.ResMes `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *SetAttrResponse) Reset() {
	*x = SetAttrResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_setattr_proto_setattr_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetAttrResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAttrResponse) ProtoMessage() {}

func (x *SetAttrResponse) ProtoReflect() protoreflect.Message {
	mi := &file_setattr_proto_setattr_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAttrResponse.ProtoReflect.Descriptor instead.
func (*SetAttrResponse) Descriptor() ([]byte, []int) {
	return file_setattr_proto_setattr_proto_rawDescGZIP(), []int{1}
}

func (x *SetAttrResponse) GetResult() *proto.ResMes {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_setattr_proto_setattr_proto protoreflect.FileDescriptor

var file_setattr_proto_setattr_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x65, 0x74, 0x61, 0x74, 0x74, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x73, 0x65, 0x74, 0x61, 0x74, 0x74, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x59,
	0x72, 0x63, 0x66, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x1a, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe1, 0x02, 0x0a, 0x0e, 0x53, 0x65, 0x74,
	0x41, 0x74, 0x74, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x75,
	0x73, 0x65, 0x5f, 0x61, 0x62, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x75, 0x73, 0x65, 0x41, 0x62, 0x73, 0x6f, 0x6c,
	0x75, 0x74, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x67, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x67, 0x69, 0x64,
	0x12, 0x22, 0x0a, 0x0a, 0x6d, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x09, 0x6d, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x63,
	0x73, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x6d, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6e, 0x61,
	0x6e, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x09, 0x6d, 0x74, 0x69, 0x6d,
	0x65, 0x4e, 0x61, 0x6e, 0x6f, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x61, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x73, 0x65, 0x63, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x48, 0x02, 0x52, 0x09,
	0x61, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x63, 0x73, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a,
	0x61, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6e, 0x61, 0x6e, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03,
	0x48, 0x03, 0x52, 0x09, 0x61, 0x74, 0x69, 0x6d, 0x65, 0x4e, 0x61, 0x6e, 0x6f, 0x88, 0x01, 0x01,
	0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6d, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x73, 0x42,
	0x0d, 0x0a, 0x0b, 0x5f, 0x6d, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6e, 0x61, 0x6e, 0x6f, 0x42, 0x0d,
	0x0a, 0x0b, 0x5f, 0x61, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x73, 0x42, 0x0d, 0x0a,
	0x0b, 0x5f, 0x61, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6e, 0x61, 0x6e, 0x6f, 0x22, 0x32, 0x0a, 0x0f,
	0x53, 0x65, 0x74, 0x41, 0x74, 0x74, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1f, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x07, 0x2e, 0x52, 0x65, 0x73, 0x4d, 0x65, 0x73, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x32, 0x4d, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x41, 0x74, 0x74, 0x72, 0x12, 0x42, 0x0a, 0x07, 0x53,
	0x65, 0x74, 0x41, 0x74, 0x74, 0x72, 0x12, 0x19, 0x2e, 0x59, 0x72, 0x63, 0x66, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x2e, 0x53, 0x65, 0x74, 0x41, 0x74, 0x74, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x59, 0x72, 0x63, 0x66, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x65,
	0x74, 0x41, 0x74, 0x74, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x2c, 0x5a, 0x2a, 0x79, 0x72, 0x63, 0x66, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x61, 0x73,
	0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65, 0x74, 0x61, 0x74, 0x74, 0x72, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x3b, 0x73, 0x65, 0x74, 0x61, 0x74, 0x74, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_setattr_proto_setattr_proto_rawDescOnce sync.Once
	file_setattr_proto_setattr_proto_rawDescData = file_setattr_proto_setattr_proto_rawDesc
)

func file_setattr_proto_setattr_proto_rawDescGZIP() []byte {
	file_setattr_proto_setattr_proto_rawDescOnce.Do(func() {
		file_setattr_proto_setattr_proto_rawDescData = protoimpl.X.CompressGZIP(file_setattr_proto_setattr_proto_rawDescData)
	})
	return file_setattr_proto_setattr_proto_rawDescData
}

var file_setattr_proto_setattr_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_setattr_proto_setattr_proto_goTypes = []any{
	(*SetAttrRequest)(nil),  // 0: YrcfAgent.SetAttrRequest
	(*SetAttrResponse)(nil), // 1: YrcfAgent.SetAttrResponse
	(*proto.ResMes)(nil),    // 2: ResMes
}
var file_setattr_proto_setattr_proto_depIdxs = []int32{
	2, // 0: YrcfAgent.SetAttrResponse.result:type_name -> ResMes
	0, // 1: YrcfAgent.SetAttr.SetAttr:input_type -> YrcfAgent.SetAttrRequest
	1, // 2: YrcfAgent.SetAttr.SetAttr:output_type -> YrcfAgent.SetAttrResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_setattr_proto_setattr_proto_init() }
func file_setattr_proto_setattr_proto_init() {
	if File_setattr_proto_setattr_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_setattr_proto_setattr_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SetAttrRequest); i {
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
		file_setattr_proto_setattr_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SetAttrResponse); i {
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
	file_setattr_proto_setattr_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_setattr_proto_setattr_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_setattr_proto_setattr_proto_goTypes,
		DependencyIndexes: file_setattr_proto_setattr_proto_depIdxs,
		MessageInfos:      file_setattr_proto_setattr_proto_msgTypes,
	}.Build()
	File_setattr_proto_setattr_proto = out.File
	file_setattr_proto_setattr_proto_rawDesc = nil
	file_setattr_proto_setattr_proto_goTypes = nil
	file_setattr_proto_setattr_proto_depIdxs = nil
}
