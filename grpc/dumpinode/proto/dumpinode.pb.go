// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: dumpinode/proto/dumpinode.proto

package dumpinode

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	proto "yrcf_grpc_case/grpc/common/proto/"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DumpInodeRequest_NodeType int32

const (
	DumpInodeRequest_MGR DumpInodeRequest_NodeType = 0
	DumpInodeRequest_MDS DumpInodeRequest_NodeType = 1
	DumpInodeRequest_OSS DumpInodeRequest_NodeType = 2
)

// Enum value maps for DumpInodeRequest_NodeType.
var (
	DumpInodeRequest_NodeType_name = map[int32]string{
		0: "MGR",
		1: "MDS",
		2: "OSS",
	}
	DumpInodeRequest_NodeType_value = map[string]int32{
		"MGR": 0,
		"MDS": 1,
		"OSS": 2,
	}
)

func (x DumpInodeRequest_NodeType) Enum() *DumpInodeRequest_NodeType {
	p := new(DumpInodeRequest_NodeType)
	*p = x
	return p
}

func (x DumpInodeRequest_NodeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DumpInodeRequest_NodeType) Descriptor() protoreflect.EnumDescriptor {
	return file_dumpinode_proto_dumpinode_proto_enumTypes[0].Descriptor()
}

func (DumpInodeRequest_NodeType) Type() protoreflect.EnumType {
	return &file_dumpinode_proto_dumpinode_proto_enumTypes[0]
}

func (x DumpInodeRequest_NodeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DumpInodeRequest_NodeType.Descriptor instead.
func (DumpInodeRequest_NodeType) EnumDescriptor() ([]byte, []int) {
	return file_dumpinode_proto_dumpinode_proto_rawDescGZIP(), []int{0, 0}
}

type InodeInfo_EntryType int32

const (
	InodeInfo_DirEntryType_INVALID     InodeInfo_EntryType = 0
	InodeInfo_DirEntryType_DIRECTORY   InodeInfo_EntryType = 1
	InodeInfo_DirEntryType_REGULARFILE InodeInfo_EntryType = 2
	InodeInfo_DirEntryType_SYMLINK     InodeInfo_EntryType = 3
	InodeInfo_DirEntryType_BLOCKDEV    InodeInfo_EntryType = 4
	InodeInfo_DirEntryType_CHARDEV     InodeInfo_EntryType = 5
	InodeInfo_DirEntryType_FIFO        InodeInfo_EntryType = 6
	InodeInfo_DirEntryType_SOCKET      InodeInfo_EntryType = 7
)

// Enum value maps for InodeInfo_EntryType.
var (
	InodeInfo_EntryType_name = map[int32]string{
		0: "DirEntryType_INVALID",
		1: "DirEntryType_DIRECTORY",
		2: "DirEntryType_REGULARFILE",
		3: "DirEntryType_SYMLINK",
		4: "DirEntryType_BLOCKDEV",
		5: "DirEntryType_CHARDEV",
		6: "DirEntryType_FIFO",
		7: "DirEntryType_SOCKET",
	}
	InodeInfo_EntryType_value = map[string]int32{
		"DirEntryType_INVALID":     0,
		"DirEntryType_DIRECTORY":   1,
		"DirEntryType_REGULARFILE": 2,
		"DirEntryType_SYMLINK":     3,
		"DirEntryType_BLOCKDEV":    4,
		"DirEntryType_CHARDEV":     5,
		"DirEntryType_FIFO":        6,
		"DirEntryType_SOCKET":      7,
	}
)

func (x InodeInfo_EntryType) Enum() *InodeInfo_EntryType {
	p := new(InodeInfo_EntryType)
	*p = x
	return p
}

func (x InodeInfo_EntryType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InodeInfo_EntryType) Descriptor() protoreflect.EnumDescriptor {
	return file_dumpinode_proto_dumpinode_proto_enumTypes[1].Descriptor()
}

func (InodeInfo_EntryType) Type() protoreflect.EnumType {
	return &file_dumpinode_proto_dumpinode_proto_enumTypes[1]
}

func (x InodeInfo_EntryType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InodeInfo_EntryType.Descriptor instead.
func (InodeInfo_EntryType) EnumDescriptor() ([]byte, []int) {
	return file_dumpinode_proto_dumpinode_proto_rawDescGZIP(), []int{1, 0}
}

// message for dumpinode
type DumpInodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Use path for mounted.(default false)
	UseAbsolutePath bool `protobuf:"varint,1,opt,name=use_absolute_path,json=useAbsolutePath,proto3" json:"use_absolute_path,omitempty"`
	// Specify the directory.(default "")
	Path string                    `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Type DumpInodeRequest_NodeType `protobuf:"varint,3,opt,name=type,proto3,enum=YrcfAgent.DumpInodeRequest_NodeType" json:"type,omitempty"`
	// Specify the entryid.(default "")
	EntryId string `protobuf:"bytes,4,opt,name=entry_id,json=entryId,proto3" json:"entry_id,omitempty"`
}

func (x *DumpInodeRequest) Reset() {
	*x = DumpInodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dumpinode_proto_dumpinode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DumpInodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DumpInodeRequest) ProtoMessage() {}

func (x *DumpInodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dumpinode_proto_dumpinode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DumpInodeRequest.ProtoReflect.Descriptor instead.
func (*DumpInodeRequest) Descriptor() ([]byte, []int) {
	return file_dumpinode_proto_dumpinode_proto_rawDescGZIP(), []int{0}
}

func (x *DumpInodeRequest) GetUseAbsolutePath() bool {
	if x != nil {
		return x.UseAbsolutePath
	}
	return false
}

func (x *DumpInodeRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *DumpInodeRequest) GetType() DumpInodeRequest_NodeType {
	if x != nil {
		return x.Type
	}
	return DumpInodeRequest_MGR
}

func (x *DumpInodeRequest) GetEntryId() string {
	if x != nil {
		return x.EntryId
	}
	return ""
}

type InodeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntryType       InodeInfo_EntryType `protobuf:"varint,1,opt,name=entry_type,json=entryType,proto3,enum=YrcfAgent.InodeInfo_EntryType" json:"entry_type,omitempty"`
	ParentEntryId   string              `protobuf:"bytes,2,opt,name=parent_entry_id,json=parentEntryId,proto3" json:"parent_entry_id,omitempty"`
	ParentNodeId    string              `protobuf:"bytes,3,opt,name=parent_node_id,json=parentNodeId,proto3" json:"parent_node_id,omitempty"`
	OwnerNodeId     string              `protobuf:"bytes,4,opt,name=owner_node_id,json=ownerNodeId,proto3" json:"owner_node_id,omitempty"`
	Mode            uint32              `protobuf:"varint,5,opt,name=mode,proto3" json:"mode,omitempty"`
	Uid             uint32              `protobuf:"varint,6,opt,name=uid,proto3" json:"uid,omitempty"`
	Gid             uint32              `protobuf:"varint,7,opt,name=gid,proto3" json:"gid,omitempty"`
	Size            int64               `protobuf:"varint,8,opt,name=size,proto3" json:"size,omitempty"`
	NumLinks        int64               `protobuf:"varint,9,opt,name=num_links,json=numLinks,proto3" json:"num_links,omitempty"`
	Ctime           int64               `protobuf:"varint,10,opt,name=ctime,proto3" json:"ctime,omitempty"`
	Atime           int64               `protobuf:"varint,11,opt,name=atime,proto3" json:"atime,omitempty"`
	Mtime           int64               `protobuf:"varint,12,opt,name=mtime,proto3" json:"mtime,omitempty"`
	FeatureFlags    int64               `protobuf:"varint,13,opt,name=feature_flags,json=featureFlags,proto3" json:"feature_flags,omitempty"`
	FeatureFlagsStr string              `protobuf:"bytes,14,opt,name=feature_flags_str,json=featureFlagsStr,proto3" json:"feature_flags_str,omitempty"`
	NumSubDirs      uint32              `protobuf:"varint,15,opt,name=num_sub_dirs,json=numSubDirs,proto3" json:"num_sub_dirs,omitempty"`
	NumSubFiles     uint32              `protobuf:"varint,16,opt,name=num_sub_files,json=numSubFiles,proto3" json:"num_sub_files,omitempty"`
}

func (x *InodeInfo) Reset() {
	*x = InodeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dumpinode_proto_dumpinode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InodeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InodeInfo) ProtoMessage() {}

func (x *InodeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_dumpinode_proto_dumpinode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InodeInfo.ProtoReflect.Descriptor instead.
func (*InodeInfo) Descriptor() ([]byte, []int) {
	return file_dumpinode_proto_dumpinode_proto_rawDescGZIP(), []int{1}
}

func (x *InodeInfo) GetEntryType() InodeInfo_EntryType {
	if x != nil {
		return x.EntryType
	}
	return InodeInfo_DirEntryType_INVALID
}

func (x *InodeInfo) GetParentEntryId() string {
	if x != nil {
		return x.ParentEntryId
	}
	return ""
}

func (x *InodeInfo) GetParentNodeId() string {
	if x != nil {
		return x.ParentNodeId
	}
	return ""
}

func (x *InodeInfo) GetOwnerNodeId() string {
	if x != nil {
		return x.OwnerNodeId
	}
	return ""
}

func (x *InodeInfo) GetMode() uint32 {
	if x != nil {
		return x.Mode
	}
	return 0
}

func (x *InodeInfo) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *InodeInfo) GetGid() uint32 {
	if x != nil {
		return x.Gid
	}
	return 0
}

func (x *InodeInfo) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *InodeInfo) GetNumLinks() int64 {
	if x != nil {
		return x.NumLinks
	}
	return 0
}

func (x *InodeInfo) GetCtime() int64 {
	if x != nil {
		return x.Ctime
	}
	return 0
}

func (x *InodeInfo) GetAtime() int64 {
	if x != nil {
		return x.Atime
	}
	return 0
}

func (x *InodeInfo) GetMtime() int64 {
	if x != nil {
		return x.Mtime
	}
	return 0
}

func (x *InodeInfo) GetFeatureFlags() int64 {
	if x != nil {
		return x.FeatureFlags
	}
	return 0
}

func (x *InodeInfo) GetFeatureFlagsStr() string {
	if x != nil {
		return x.FeatureFlagsStr
	}
	return ""
}

func (x *InodeInfo) GetNumSubDirs() uint32 {
	if x != nil {
		return x.NumSubDirs
	}
	return 0
}

func (x *InodeInfo) GetNumSubFiles() uint32 {
	if x != nil {
		return x.NumSubFiles
	}
	return 0
}

type DumpInodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result       *proto.ResMes `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Command      string        `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`
	ResponseFrom string        `protobuf:"bytes,3,opt,name=response_from,json=responseFrom,proto3" json:"response_from,omitempty"`
	RespInfo     *InodeInfo    `protobuf:"bytes,4,opt,name=resp_info,json=respInfo,proto3" json:"resp_info,omitempty"`
}

func (x *DumpInodeResponse) Reset() {
	*x = DumpInodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dumpinode_proto_dumpinode_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DumpInodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DumpInodeResponse) ProtoMessage() {}

func (x *DumpInodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dumpinode_proto_dumpinode_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DumpInodeResponse.ProtoReflect.Descriptor instead.
func (*DumpInodeResponse) Descriptor() ([]byte, []int) {
	return file_dumpinode_proto_dumpinode_proto_rawDescGZIP(), []int{2}
}

func (x *DumpInodeResponse) GetResult() *proto.ResMes {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *DumpInodeResponse) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *DumpInodeResponse) GetResponseFrom() string {
	if x != nil {
		return x.ResponseFrom
	}
	return ""
}

func (x *DumpInodeResponse) GetRespInfo() *InodeInfo {
	if x != nil {
		return x.RespInfo
	}
	return nil
}

var File_dumpinode_proto_dumpinode_proto protoreflect.FileDescriptor

var file_dumpinode_proto_dumpinode_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x64, 0x75, 0x6d, 0x70, 0x69, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x64, 0x75, 0x6d, 0x70, 0x69, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x59, 0x72, 0x63, 0x66, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x1a, 0x0e, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xce, 0x01, 0x0a,
	0x10, 0x44, 0x75, 0x6d, 0x70, 0x49, 0x6e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2a, 0x0a, 0x11, 0x75, 0x73, 0x65, 0x5f, 0x61, 0x62, 0x73, 0x6f, 0x6c, 0x75, 0x74,
	0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x75, 0x73,
	0x65, 0x41, 0x62, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x12, 0x38, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x24, 0x2e, 0x59, 0x72, 0x63, 0x66, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x75, 0x6d, 0x70,
	0x49, 0x6e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4e, 0x6f, 0x64,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x22, 0x25, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x47, 0x52, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4d,
	0x44, 0x53, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x4f, 0x53, 0x53, 0x10, 0x02, 0x22, 0xdf, 0x05,
	0x0a, 0x09, 0x49, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3d, 0x0a, 0x0a, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1e, 0x2e, 0x59, 0x72, 0x63, 0x66, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x49, 0x6e, 0x6f, 0x64,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x09, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x6f, 0x64,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6d, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75,
	0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x03, 0x67, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x75, 0x6d, 0x5f,
	0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6e, 0x75, 0x6d,
	0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x61, 0x74, 0x69, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x6d, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x2a, 0x0a, 0x11,
	0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x5f, 0x73, 0x74,
	0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x46, 0x6c, 0x61, 0x67, 0x73, 0x53, 0x74, 0x72, 0x12, 0x20, 0x0a, 0x0c, 0x6e, 0x75, 0x6d, 0x5f,
	0x73, 0x75, 0x62, 0x5f, 0x64, 0x69, 0x72, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x6e, 0x75, 0x6d, 0x53, 0x75, 0x62, 0x44, 0x69, 0x72, 0x73, 0x12, 0x22, 0x0a, 0x0d, 0x6e, 0x75,
	0x6d, 0x5f, 0x73, 0x75, 0x62, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x6e, 0x75, 0x6d, 0x53, 0x75, 0x62, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x22, 0xde,
	0x01, 0x0a, 0x09, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x14,
	0x44, 0x69, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x44, 0x69, 0x72, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x4f, 0x52, 0x59,
	0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x44, 0x69, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x5f, 0x52, 0x45, 0x47, 0x55, 0x4c, 0x41, 0x52, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x02,
	0x12, 0x18, 0x0a, 0x14, 0x44, 0x69, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65,
	0x5f, 0x53, 0x59, 0x4d, 0x4c, 0x49, 0x4e, 0x4b, 0x10, 0x03, 0x12, 0x19, 0x0a, 0x15, 0x44, 0x69,
	0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x42, 0x4c, 0x4f, 0x43, 0x4b,
	0x44, 0x45, 0x56, 0x10, 0x04, 0x12, 0x18, 0x0a, 0x14, 0x44, 0x69, 0x72, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x43, 0x48, 0x41, 0x52, 0x44, 0x45, 0x56, 0x10, 0x05, 0x12,
	0x15, 0x0a, 0x11, 0x44, 0x69, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x5f,
	0x46, 0x49, 0x46, 0x4f, 0x10, 0x06, 0x12, 0x17, 0x0a, 0x13, 0x44, 0x69, 0x72, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x53, 0x4f, 0x43, 0x4b, 0x45, 0x54, 0x10, 0x07, 0x22,
	0xb0, 0x01, 0x0a, 0x11, 0x44, 0x75, 0x6d, 0x70, 0x49, 0x6e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x59, 0x72, 0x63, 0x66, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x2e, 0x52, 0x65, 0x73, 0x4d, 0x65, 0x73, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x12,
	0x31, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x70, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x59, 0x72, 0x63, 0x66, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x49,
	0x6e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x49, 0x6e,
	0x66, 0x6f, 0x32, 0x55, 0x0a, 0x09, 0x44, 0x75, 0x6d, 0x70, 0x49, 0x6e, 0x6f, 0x64, 0x65, 0x12,
	0x48, 0x0a, 0x09, 0x44, 0x75, 0x6d, 0x70, 0x49, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x2e, 0x59,
	0x72, 0x63, 0x66, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x75, 0x6d, 0x70, 0x49, 0x6e, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x59, 0x72, 0x63, 0x66,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x75, 0x6d, 0x70, 0x49, 0x6e, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x30, 0x5a, 0x2e, 0x79, 0x72, 0x63,
	0x66, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x64, 0x75, 0x6d, 0x70, 0x69, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x3b, 0x64, 0x75, 0x6d, 0x70, 0x69, 0x6e, 0x6f, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_dumpinode_proto_dumpinode_proto_rawDescOnce sync.Once
	file_dumpinode_proto_dumpinode_proto_rawDescData = file_dumpinode_proto_dumpinode_proto_rawDesc
)

func file_dumpinode_proto_dumpinode_proto_rawDescGZIP() []byte {
	file_dumpinode_proto_dumpinode_proto_rawDescOnce.Do(func() {
		file_dumpinode_proto_dumpinode_proto_rawDescData = protoimpl.X.CompressGZIP(file_dumpinode_proto_dumpinode_proto_rawDescData)
	})
	return file_dumpinode_proto_dumpinode_proto_rawDescData
}

var file_dumpinode_proto_dumpinode_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_dumpinode_proto_dumpinode_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_dumpinode_proto_dumpinode_proto_goTypes = []any{
	(DumpInodeRequest_NodeType)(0), // 0: YrcfAgent.DumpInodeRequest.NodeType
	(InodeInfo_EntryType)(0),       // 1: YrcfAgent.InodeInfo.EntryType
	(*DumpInodeRequest)(nil),       // 2: YrcfAgent.DumpInodeRequest
	(*InodeInfo)(nil),              // 3: YrcfAgent.InodeInfo
	(*DumpInodeResponse)(nil),      // 4: YrcfAgent.DumpInodeResponse
	(*proto.ResMes)(nil),           // 5: YrcfAgent.ResMes
}
var file_dumpinode_proto_dumpinode_proto_depIdxs = []int32{
	0, // 0: YrcfAgent.DumpInodeRequest.type:type_name -> YrcfAgent.DumpInodeRequest.NodeType
	1, // 1: YrcfAgent.InodeInfo.entry_type:type_name -> YrcfAgent.InodeInfo.EntryType
	5, // 2: YrcfAgent.DumpInodeResponse.result:type_name -> YrcfAgent.ResMes
	3, // 3: YrcfAgent.DumpInodeResponse.resp_info:type_name -> YrcfAgent.InodeInfo
	2, // 4: YrcfAgent.DumpInode.DumpInode:input_type -> YrcfAgent.DumpInodeRequest
	4, // 5: YrcfAgent.DumpInode.DumpInode:output_type -> YrcfAgent.DumpInodeResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_dumpinode_proto_dumpinode_proto_init() }
func file_dumpinode_proto_dumpinode_proto_init() {
	if File_dumpinode_proto_dumpinode_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dumpinode_proto_dumpinode_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*DumpInodeRequest); i {
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
		file_dumpinode_proto_dumpinode_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*InodeInfo); i {
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
		file_dumpinode_proto_dumpinode_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*DumpInodeResponse); i {
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
			RawDescriptor: file_dumpinode_proto_dumpinode_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dumpinode_proto_dumpinode_proto_goTypes,
		DependencyIndexes: file_dumpinode_proto_dumpinode_proto_depIdxs,
		EnumInfos:         file_dumpinode_proto_dumpinode_proto_enumTypes,
		MessageInfos:      file_dumpinode_proto_dumpinode_proto_msgTypes,
	}.Build()
	File_dumpinode_proto_dumpinode_proto = out.File
	file_dumpinode_proto_dumpinode_proto_rawDesc = nil
	file_dumpinode_proto_dumpinode_proto_goTypes = nil
	file_dumpinode_proto_dumpinode_proto_depIdxs = nil
}
