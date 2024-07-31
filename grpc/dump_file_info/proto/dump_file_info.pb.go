// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: dump_file_info/proto/dump_file_info.proto

// If you want to make any modifications, please synchronize to the yrdi project

package dump_file_info

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

type FileType int32

const (
	FileType_FileType_INVALID     FileType = 0
	FileType_FileType_DIRECTORY   FileType = 1
	FileType_FileType_REGULARFILE FileType = 2
	FileType_FileType_SYMLINK     FileType = 3
	FileType_FileType_BLOCKDEV    FileType = 4
	FileType_FileType_CHARDEV     FileType = 5
	FileType_FileType_FIFO        FileType = 6
	FileType_FileType_SOCKET      FileType = 7
)

// Enum value maps for FileType.
var (
	FileType_name = map[int32]string{
		0: "FileType_INVALID",
		1: "FileType_DIRECTORY",
		2: "FileType_REGULARFILE",
		3: "FileType_SYMLINK",
		4: "FileType_BLOCKDEV",
		5: "FileType_CHARDEV",
		6: "FileType_FIFO",
		7: "FileType_SOCKET",
	}
	FileType_value = map[string]int32{
		"FileType_INVALID":     0,
		"FileType_DIRECTORY":   1,
		"FileType_REGULARFILE": 2,
		"FileType_SYMLINK":     3,
		"FileType_BLOCKDEV":    4,
		"FileType_CHARDEV":     5,
		"FileType_FIFO":        6,
		"FileType_SOCKET":      7,
	}
)

func (x FileType) Enum() *FileType {
	p := new(FileType)
	*p = x
	return p
}

func (x FileType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FileType) Descriptor() protoreflect.EnumDescriptor {
	return file_dump_file_info_proto_dump_file_info_proto_enumTypes[0].Descriptor()
}

func (FileType) Type() protoreflect.EnumType {
	return &file_dump_file_info_proto_dump_file_info_proto_enumTypes[0]
}

func (x FileType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FileType.Descriptor instead.
func (FileType) EnumDescriptor() ([]byte, []int) {
	return file_dump_file_info_proto_dump_file_info_proto_rawDescGZIP(), []int{0}
}

type Layout int32

const (
	Layout_Layout_Invalid Layout = 0
	Layout_Layout_Raid0   Layout = 1
	Layout_Layout_Mirror  Layout = 2
	Layout_Layout_Standby Layout = 3
)

// Enum value maps for Layout.
var (
	Layout_name = map[int32]string{
		0: "Layout_Invalid",
		1: "Layout_Raid0",
		2: "Layout_Mirror",
		3: "Layout_Standby",
	}
	Layout_value = map[string]int32{
		"Layout_Invalid": 0,
		"Layout_Raid0":   1,
		"Layout_Mirror":  2,
		"Layout_Standby": 3,
	}
)

func (x Layout) Enum() *Layout {
	p := new(Layout)
	*p = x
	return p
}

func (x Layout) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Layout) Descriptor() protoreflect.EnumDescriptor {
	return file_dump_file_info_proto_dump_file_info_proto_enumTypes[1].Descriptor()
}

func (Layout) Type() protoreflect.EnumType {
	return &file_dump_file_info_proto_dump_file_info_proto_enumTypes[1]
}

func (x Layout) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Layout.Descriptor instead.
func (Layout) EnumDescriptor() ([]byte, []int) {
	return file_dump_file_info_proto_dump_file_info_proto_rawDescGZIP(), []int{1}
}

type S3Layer int32

const (
	S3Layer_S3Layer_Invalid S3Layer = 0
	S3Layer_S3Layer_Hot     S3Layer = 1
	S3Layer_S3Layer_Mixed   S3Layer = 2
	S3Layer_S3Layer_Cold    S3Layer = 3
)

// Enum value maps for S3Layer.
var (
	S3Layer_name = map[int32]string{
		0: "S3Layer_Invalid",
		1: "S3Layer_Hot",
		2: "S3Layer_Mixed",
		3: "S3Layer_Cold",
	}
	S3Layer_value = map[string]int32{
		"S3Layer_Invalid": 0,
		"S3Layer_Hot":     1,
		"S3Layer_Mixed":   2,
		"S3Layer_Cold":    3,
	}
)

func (x S3Layer) Enum() *S3Layer {
	p := new(S3Layer)
	*p = x
	return p
}

func (x S3Layer) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (S3Layer) Descriptor() protoreflect.EnumDescriptor {
	return file_dump_file_info_proto_dump_file_info_proto_enumTypes[2].Descriptor()
}

func (S3Layer) Type() protoreflect.EnumType {
	return &file_dump_file_info_proto_dump_file_info_proto_enumTypes[2]
}

func (x S3Layer) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use S3Layer.Descriptor instead.
func (S3Layer) EnumDescriptor() ([]byte, []int) {
	return file_dump_file_info_proto_dump_file_info_proto_rawDescGZIP(), []int{2}
}

type MetaInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Peid    string   `protobuf:"bytes,1,opt,name=peid,proto3" json:"peid,omitempty"`
	Eid     string   `protobuf:"bytes,2,opt,name=eid,proto3" json:"eid,omitempty"`
	Name    string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Type    FileType `protobuf:"varint,4,opt,name=type,proto3,enum=yrdi.FileType" json:"type,omitempty"`
	Layout  Layout   `protobuf:"varint,5,opt,name=layout,proto3,enum=yrdi.Layout" json:"layout,omitempty"`
	GroupId uint32   `protobuf:"varint,6,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	NodeId  string   `protobuf:"bytes,7,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
}

func (x *MetaInfo) Reset() {
	*x = MetaInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dump_file_info_proto_dump_file_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetaInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetaInfo) ProtoMessage() {}

func (x *MetaInfo) ProtoReflect() protoreflect.Message {
	mi := &file_dump_file_info_proto_dump_file_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetaInfo.ProtoReflect.Descriptor instead.
func (*MetaInfo) Descriptor() ([]byte, []int) {
	return file_dump_file_info_proto_dump_file_info_proto_rawDescGZIP(), []int{0}
}

func (x *MetaInfo) GetPeid() string {
	if x != nil {
		return x.Peid
	}
	return ""
}

func (x *MetaInfo) GetEid() string {
	if x != nil {
		return x.Eid
	}
	return ""
}

func (x *MetaInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MetaInfo) GetType() FileType {
	if x != nil {
		return x.Type
	}
	return FileType_FileType_INVALID
}

func (x *MetaInfo) GetLayout() Layout {
	if x != nil {
		return x.Layout
	}
	return Layout_Layout_Invalid
}

func (x *MetaInfo) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *MetaInfo) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

type DataInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	S3Layer      S3Layer  `protobuf:"varint,1,opt,name=s3_layer,json=s3Layer,proto3,enum=yrdi.S3Layer" json:"s3_layer,omitempty"`
	DataLayout   Layout   `protobuf:"varint,2,opt,name=data_layout,json=dataLayout,proto3,enum=yrdi.Layout" json:"data_layout,omitempty"`
	TieringId    uint32   `protobuf:"varint,3,opt,name=tiering_id,json=tieringId,proto3" json:"tiering_id,omitempty"`
	BucketLinkId uint32   `protobuf:"varint,4,opt,name=bucket_link_id,json=bucketLinkId,proto3" json:"bucket_link_id,omitempty"`
	ProjectId    uint32   `protobuf:"varint,5,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	StripeSize   uint64   `protobuf:"varint,6,opt,name=stripe_size,json=stripeSize,proto3" json:"stripe_size,omitempty"`
	StripeCount  uint32   `protobuf:"varint,7,opt,name=stripe_count,json=stripeCount,proto3" json:"stripe_count,omitempty"`
	Placement    []uint32 `protobuf:"varint,8,rep,packed,name=placement,proto3" json:"placement,omitempty"`
}

func (x *DataInfo) Reset() {
	*x = DataInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dump_file_info_proto_dump_file_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataInfo) ProtoMessage() {}

func (x *DataInfo) ProtoReflect() protoreflect.Message {
	mi := &file_dump_file_info_proto_dump_file_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataInfo.ProtoReflect.Descriptor instead.
func (*DataInfo) Descriptor() ([]byte, []int) {
	return file_dump_file_info_proto_dump_file_info_proto_rawDescGZIP(), []int{1}
}

func (x *DataInfo) GetS3Layer() S3Layer {
	if x != nil {
		return x.S3Layer
	}
	return S3Layer_S3Layer_Invalid
}

func (x *DataInfo) GetDataLayout() Layout {
	if x != nil {
		return x.DataLayout
	}
	return Layout_Layout_Invalid
}

func (x *DataInfo) GetTieringId() uint32 {
	if x != nil {
		return x.TieringId
	}
	return 0
}

func (x *DataInfo) GetBucketLinkId() uint32 {
	if x != nil {
		return x.BucketLinkId
	}
	return 0
}

func (x *DataInfo) GetProjectId() uint32 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

func (x *DataInfo) GetStripeSize() uint64 {
	if x != nil {
		return x.StripeSize
	}
	return 0
}

func (x *DataInfo) GetStripeCount() uint32 {
	if x != nil {
		return x.StripeCount
	}
	return 0
}

func (x *DataInfo) GetPlacement() []uint32 {
	if x != nil {
		return x.Placement
	}
	return nil
}

type StatInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode     int32  `protobuf:"varint,1,opt,name=mode,proto3" json:"mode,omitempty"`
	FileSize int64  `protobuf:"varint,2,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	NrBlocks uint64 `protobuf:"varint,3,opt,name=nr_blocks,json=nrBlocks,proto3" json:"nr_blocks,omitempty"` // Currently, this field has not been obtained
	Nlink    uint32 `protobuf:"varint,4,opt,name=nlink,proto3" json:"nlink,omitempty"`
	Uid      uint32 `protobuf:"varint,5,opt,name=uid,proto3" json:"uid,omitempty"`
	Gid      uint32 `protobuf:"varint,6,opt,name=gid,proto3" json:"gid,omitempty"`
	BtimeSec int64  `protobuf:"varint,7,opt,name=btime_sec,json=btimeSec,proto3" json:"btime_sec,omitempty"`
	BtimeNs  int64  `protobuf:"varint,8,opt,name=btime_ns,json=btimeNs,proto3" json:"btime_ns,omitempty"`
	CtimeSec int64  `protobuf:"varint,9,opt,name=ctime_sec,json=ctimeSec,proto3" json:"ctime_sec,omitempty"`
	CtimeNs  int64  `protobuf:"varint,10,opt,name=ctime_ns,json=ctimeNs,proto3" json:"ctime_ns,omitempty"`
	MtimeSec int64  `protobuf:"varint,11,opt,name=mtime_sec,json=mtimeSec,proto3" json:"mtime_sec,omitempty"`
	MtimeNs  int64  `protobuf:"varint,12,opt,name=mtime_ns,json=mtimeNs,proto3" json:"mtime_ns,omitempty"`
	AtimeSec int64  `protobuf:"varint,13,opt,name=atime_sec,json=atimeSec,proto3" json:"atime_sec,omitempty"`
	AtimeNs  int64  `protobuf:"varint,14,opt,name=atime_ns,json=atimeNs,proto3" json:"atime_ns,omitempty"`
}

func (x *StatInfo) Reset() {
	*x = StatInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dump_file_info_proto_dump_file_info_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatInfo) ProtoMessage() {}

func (x *StatInfo) ProtoReflect() protoreflect.Message {
	mi := &file_dump_file_info_proto_dump_file_info_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatInfo.ProtoReflect.Descriptor instead.
func (*StatInfo) Descriptor() ([]byte, []int) {
	return file_dump_file_info_proto_dump_file_info_proto_rawDescGZIP(), []int{2}
}

func (x *StatInfo) GetMode() int32 {
	if x != nil {
		return x.Mode
	}
	return 0
}

func (x *StatInfo) GetFileSize() int64 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *StatInfo) GetNrBlocks() uint64 {
	if x != nil {
		return x.NrBlocks
	}
	return 0
}

func (x *StatInfo) GetNlink() uint32 {
	if x != nil {
		return x.Nlink
	}
	return 0
}

func (x *StatInfo) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *StatInfo) GetGid() uint32 {
	if x != nil {
		return x.Gid
	}
	return 0
}

func (x *StatInfo) GetBtimeSec() int64 {
	if x != nil {
		return x.BtimeSec
	}
	return 0
}

func (x *StatInfo) GetBtimeNs() int64 {
	if x != nil {
		return x.BtimeNs
	}
	return 0
}

func (x *StatInfo) GetCtimeSec() int64 {
	if x != nil {
		return x.CtimeSec
	}
	return 0
}

func (x *StatInfo) GetCtimeNs() int64 {
	if x != nil {
		return x.CtimeNs
	}
	return 0
}

func (x *StatInfo) GetMtimeSec() int64 {
	if x != nil {
		return x.MtimeSec
	}
	return 0
}

func (x *StatInfo) GetMtimeNs() int64 {
	if x != nil {
		return x.MtimeNs
	}
	return 0
}

func (x *StatInfo) GetAtimeSec() int64 {
	if x != nil {
		return x.AtimeSec
	}
	return 0
}

func (x *StatInfo) GetAtimeNs() int64 {
	if x != nil {
		return x.AtimeNs
	}
	return 0
}

type DumpFileInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *DumpFileInfoRequest) Reset() {
	*x = DumpFileInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dump_file_info_proto_dump_file_info_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DumpFileInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DumpFileInfoRequest) ProtoMessage() {}

func (x *DumpFileInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dump_file_info_proto_dump_file_info_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DumpFileInfoRequest.ProtoReflect.Descriptor instead.
func (*DumpFileInfoRequest) Descriptor() ([]byte, []int) {
	return file_dump_file_info_proto_dump_file_info_proto_rawDescGZIP(), []int{3}
}

func (x *DumpFileInfoRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type DumpFileInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result   int32     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Msg      string    `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	MetaInfo *MetaInfo `protobuf:"bytes,3,opt,name=meta_info,json=metaInfo,proto3" json:"meta_info,omitempty"`
	DataInfo *DataInfo `protobuf:"bytes,4,opt,name=data_info,json=dataInfo,proto3" json:"data_info,omitempty"`
	StatInfo *StatInfo `protobuf:"bytes,5,opt,name=stat_info,json=statInfo,proto3" json:"stat_info,omitempty"`
}

func (x *DumpFileInfoResponse) Reset() {
	*x = DumpFileInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dump_file_info_proto_dump_file_info_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DumpFileInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DumpFileInfoResponse) ProtoMessage() {}

func (x *DumpFileInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dump_file_info_proto_dump_file_info_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DumpFileInfoResponse.ProtoReflect.Descriptor instead.
func (*DumpFileInfoResponse) Descriptor() ([]byte, []int) {
	return file_dump_file_info_proto_dump_file_info_proto_rawDescGZIP(), []int{4}
}

func (x *DumpFileInfoResponse) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

func (x *DumpFileInfoResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *DumpFileInfoResponse) GetMetaInfo() *MetaInfo {
	if x != nil {
		return x.MetaInfo
	}
	return nil
}

func (x *DumpFileInfoResponse) GetDataInfo() *DataInfo {
	if x != nil {
		return x.DataInfo
	}
	return nil
}

func (x *DumpFileInfoResponse) GetStatInfo() *StatInfo {
	if x != nil {
		return x.StatInfo
	}
	return nil
}

var File_dump_file_info_proto_dump_file_info_proto protoreflect.FileDescriptor

var file_dump_file_info_proto_dump_file_info_proto_rawDesc = []byte{
	0x0a, 0x29, 0x64, 0x75, 0x6d, 0x70, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x75, 0x6d, 0x70, 0x5f, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x79, 0x72, 0x64,
	0x69, 0x22, 0xc2, 0x01, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x65, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x65,
	0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x65, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x79, 0x72, 0x64, 0x69, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x06,
	0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x79,
	0x72, 0x64, 0x69, 0x2e, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x52, 0x06, 0x6c, 0x61, 0x79, 0x6f,
	0x75, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x22, 0xa9, 0x02, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x28, 0x0a, 0x08, 0x73, 0x33, 0x5f, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x79, 0x72, 0x64, 0x69, 0x2e, 0x53, 0x33, 0x4c,
	0x61, 0x79, 0x65, 0x72, 0x52, 0x07, 0x73, 0x33, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x2d, 0x0a,
	0x0b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x79, 0x72, 0x64, 0x69, 0x2e, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74,
	0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x74, 0x69, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x74, 0x69, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x62,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0c, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x22, 0xf2, 0x02, 0x0a, 0x08, 0x53, 0x74, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6d,
	0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x72, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x6e, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6e, 0x6c,
	0x69, 0x6e, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x03, 0x67, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x73, 0x65, 0x63, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x62, 0x74, 0x69, 0x6d,
	0x65, 0x53, 0x65, 0x63, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6e, 0x73,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62, 0x74, 0x69, 0x6d, 0x65, 0x4e, 0x73, 0x12,
	0x1b, 0x0a, 0x09, 0x63, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x63, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x63, 0x12, 0x19, 0x0a, 0x08,
	0x63, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6e, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x63, 0x74, 0x69, 0x6d, 0x65, 0x4e, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x73, 0x65, 0x63, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x74, 0x69, 0x6d,
	0x65, 0x53, 0x65, 0x63, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6e, 0x73,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6d, 0x74, 0x69, 0x6d, 0x65, 0x4e, 0x73, 0x12,
	0x1b, 0x0a, 0x09, 0x61, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x61, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x63, 0x12, 0x19, 0x0a, 0x08,
	0x61, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6e, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x61, 0x74, 0x69, 0x6d, 0x65, 0x4e, 0x73, 0x22, 0x29, 0x0a, 0x13, 0x44, 0x75, 0x6d, 0x70, 0x46,
	0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x22, 0xc7, 0x01, 0x0a, 0x14, 0x44, 0x75, 0x6d, 0x70, 0x46, 0x69, 0x6c, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x2b, 0x0a, 0x09, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x79, 0x72, 0x64, 0x69, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x2b, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x79, 0x72, 0x64, 0x69, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x2b, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x79, 0x72, 0x64, 0x69, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x08, 0x73, 0x74, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2a, 0xbd, 0x01, 0x0a,
	0x08, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x46, 0x69, 0x6c,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12,
	0x16, 0x0a, 0x12, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x44, 0x49, 0x52, 0x45,
	0x43, 0x54, 0x4f, 0x52, 0x59, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x46, 0x69, 0x6c, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x5f, 0x52, 0x45, 0x47, 0x55, 0x4c, 0x41, 0x52, 0x46, 0x49, 0x4c, 0x45, 0x10,
	0x02, 0x12, 0x14, 0x0a, 0x10, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x53, 0x59,
	0x4d, 0x4c, 0x49, 0x4e, 0x4b, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x46, 0x69, 0x6c, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x5f, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x44, 0x45, 0x56, 0x10, 0x04, 0x12, 0x14,
	0x0a, 0x10, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x43, 0x48, 0x41, 0x52, 0x44,
	0x45, 0x56, 0x10, 0x05, 0x12, 0x11, 0x0a, 0x0d, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x5f, 0x46, 0x49, 0x46, 0x4f, 0x10, 0x06, 0x12, 0x13, 0x0a, 0x0f, 0x46, 0x69, 0x6c, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x5f, 0x53, 0x4f, 0x43, 0x4b, 0x45, 0x54, 0x10, 0x07, 0x2a, 0x55, 0x0a, 0x06,
	0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x0e, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74,
	0x5f, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x4c, 0x61,
	0x79, 0x6f, 0x75, 0x74, 0x5f, 0x52, 0x61, 0x69, 0x64, 0x30, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d,
	0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x5f, 0x4d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x02, 0x12,
	0x12, 0x0a, 0x0e, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x5f, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x62,
	0x79, 0x10, 0x03, 0x2a, 0x54, 0x0a, 0x07, 0x53, 0x33, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x13,
	0x0a, 0x0f, 0x53, 0x33, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x33, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x48,
	0x6f, 0x74, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x33, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x5f,
	0x4d, 0x69, 0x78, 0x65, 0x64, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x33, 0x4c, 0x61, 0x79,
	0x65, 0x72, 0x5f, 0x43, 0x6f, 0x6c, 0x64, 0x10, 0x03, 0x32, 0x57, 0x0a, 0x0c, 0x44, 0x75, 0x6d,
	0x70, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x47, 0x0a, 0x0c, 0x44, 0x75, 0x6d,
	0x70, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x2e, 0x79, 0x72, 0x64, 0x69,
	0x2e, 0x44, 0x75, 0x6d, 0x70, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x79, 0x72, 0x64, 0x69, 0x2e, 0x44, 0x75, 0x6d, 0x70,
	0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x3a, 0x5a, 0x38, 0x79, 0x72, 0x63, 0x66, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x63, 0x61, 0x73, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x64, 0x75, 0x6d, 0x70, 0x5f, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x3b,
	0x64, 0x75, 0x6d, 0x70, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dump_file_info_proto_dump_file_info_proto_rawDescOnce sync.Once
	file_dump_file_info_proto_dump_file_info_proto_rawDescData = file_dump_file_info_proto_dump_file_info_proto_rawDesc
)

func file_dump_file_info_proto_dump_file_info_proto_rawDescGZIP() []byte {
	file_dump_file_info_proto_dump_file_info_proto_rawDescOnce.Do(func() {
		file_dump_file_info_proto_dump_file_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_dump_file_info_proto_dump_file_info_proto_rawDescData)
	})
	return file_dump_file_info_proto_dump_file_info_proto_rawDescData
}

var file_dump_file_info_proto_dump_file_info_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_dump_file_info_proto_dump_file_info_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_dump_file_info_proto_dump_file_info_proto_goTypes = []any{
	(FileType)(0),                // 0: yrdi.FileType
	(Layout)(0),                  // 1: yrdi.Layout
	(S3Layer)(0),                 // 2: yrdi.S3Layer
	(*MetaInfo)(nil),             // 3: yrdi.MetaInfo
	(*DataInfo)(nil),             // 4: yrdi.DataInfo
	(*StatInfo)(nil),             // 5: yrdi.StatInfo
	(*DumpFileInfoRequest)(nil),  // 6: yrdi.DumpFileInfoRequest
	(*DumpFileInfoResponse)(nil), // 7: yrdi.DumpFileInfoResponse
}
var file_dump_file_info_proto_dump_file_info_proto_depIdxs = []int32{
	0, // 0: yrdi.MetaInfo.type:type_name -> yrdi.FileType
	1, // 1: yrdi.MetaInfo.layout:type_name -> yrdi.Layout
	2, // 2: yrdi.DataInfo.s3_layer:type_name -> yrdi.S3Layer
	1, // 3: yrdi.DataInfo.data_layout:type_name -> yrdi.Layout
	3, // 4: yrdi.DumpFileInfoResponse.meta_info:type_name -> yrdi.MetaInfo
	4, // 5: yrdi.DumpFileInfoResponse.data_info:type_name -> yrdi.DataInfo
	5, // 6: yrdi.DumpFileInfoResponse.stat_info:type_name -> yrdi.StatInfo
	6, // 7: yrdi.DumpFileInfo.DumpFileInfo:input_type -> yrdi.DumpFileInfoRequest
	7, // 8: yrdi.DumpFileInfo.DumpFileInfo:output_type -> yrdi.DumpFileInfoResponse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_dump_file_info_proto_dump_file_info_proto_init() }
func file_dump_file_info_proto_dump_file_info_proto_init() {
	if File_dump_file_info_proto_dump_file_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dump_file_info_proto_dump_file_info_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*MetaInfo); i {
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
		file_dump_file_info_proto_dump_file_info_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*DataInfo); i {
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
		file_dump_file_info_proto_dump_file_info_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*StatInfo); i {
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
		file_dump_file_info_proto_dump_file_info_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*DumpFileInfoRequest); i {
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
		file_dump_file_info_proto_dump_file_info_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*DumpFileInfoResponse); i {
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
			RawDescriptor: file_dump_file_info_proto_dump_file_info_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dump_file_info_proto_dump_file_info_proto_goTypes,
		DependencyIndexes: file_dump_file_info_proto_dump_file_info_proto_depIdxs,
		EnumInfos:         file_dump_file_info_proto_dump_file_info_proto_enumTypes,
		MessageInfos:      file_dump_file_info_proto_dump_file_info_proto_msgTypes,
	}.Build()
	File_dump_file_info_proto_dump_file_info_proto = out.File
	file_dump_file_info_proto_dump_file_info_proto_rawDesc = nil
	file_dump_file_info_proto_dump_file_info_proto_goTypes = nil
	file_dump_file_info_proto_dump_file_info_proto_depIdxs = nil
}
