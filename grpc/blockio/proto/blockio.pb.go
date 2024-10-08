// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: blockio/proto/blockio.proto

package blockio

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

type BlockLinkInfo_BlockType int32

const (
	BlockLinkInfo_BLOCK_IMPORT BlockLinkInfo_BlockType = 0
	BlockLinkInfo_BLOCK_EXPORT BlockLinkInfo_BlockType = 1
)

// Enum value maps for BlockLinkInfo_BlockType.
var (
	BlockLinkInfo_BlockType_name = map[int32]string{
		0: "BLOCK_IMPORT",
		1: "BLOCK_EXPORT",
	}
	BlockLinkInfo_BlockType_value = map[string]int32{
		"BLOCK_IMPORT": 0,
		"BLOCK_EXPORT": 1,
	}
)

func (x BlockLinkInfo_BlockType) Enum() *BlockLinkInfo_BlockType {
	p := new(BlockLinkInfo_BlockType)
	*p = x
	return p
}

func (x BlockLinkInfo_BlockType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BlockLinkInfo_BlockType) Descriptor() protoreflect.EnumDescriptor {
	return file_blockio_proto_blockio_proto_enumTypes[0].Descriptor()
}

func (BlockLinkInfo_BlockType) Type() protoreflect.EnumType {
	return &file_blockio_proto_blockio_proto_enumTypes[0]
}

func (x BlockLinkInfo_BlockType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BlockLinkInfo_BlockType.Descriptor instead.
func (BlockLinkInfo_BlockType) EnumDescriptor() ([]byte, []int) {
	return file_blockio_proto_blockio_proto_rawDescGZIP(), []int{0, 0}
}

type BlockLinkInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LinkId    uint32                  `protobuf:"varint,1,opt,name=link_id,json=linkId,proto3" json:"link_id,omitempty"`
	BlockType BlockLinkInfo_BlockType `protobuf:"varint,2,opt,name=block_type,json=blockType,proto3,enum=YrcfAgent.BlockLinkInfo_BlockType" json:"block_type,omitempty"`
}

func (x *BlockLinkInfo) Reset() {
	*x = BlockLinkInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockio_proto_blockio_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockLinkInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockLinkInfo) ProtoMessage() {}

func (x *BlockLinkInfo) ProtoReflect() protoreflect.Message {
	mi := &file_blockio_proto_blockio_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockLinkInfo.ProtoReflect.Descriptor instead.
func (*BlockLinkInfo) Descriptor() ([]byte, []int) {
	return file_blockio_proto_blockio_proto_rawDescGZIP(), []int{0}
}

func (x *BlockLinkInfo) GetLinkId() uint32 {
	if x != nil {
		return x.LinkId
	}
	return 0
}

func (x *BlockLinkInfo) GetBlockType() BlockLinkInfo_BlockType {
	if x != nil {
		return x.BlockType
	}
	return BlockLinkInfo_BLOCK_IMPORT
}

// add block link ID
type AddBlockLinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockLinkInfo *BlockLinkInfo `protobuf:"bytes,1,opt,name=block_link_info,json=blockLinkInfo,proto3" json:"block_link_info,omitempty"`
}

func (x *AddBlockLinkRequest) Reset() {
	*x = AddBlockLinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockio_proto_blockio_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBlockLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBlockLinkRequest) ProtoMessage() {}

func (x *AddBlockLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blockio_proto_blockio_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBlockLinkRequest.ProtoReflect.Descriptor instead.
func (*AddBlockLinkRequest) Descriptor() ([]byte, []int) {
	return file_blockio_proto_blockio_proto_rawDescGZIP(), []int{1}
}

func (x *AddBlockLinkRequest) GetBlockLinkInfo() *BlockLinkInfo {
	if x != nil {
		return x.BlockLinkInfo
	}
	return nil
}

type AddBlockLinkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *proto.ResMes `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *AddBlockLinkResponse) Reset() {
	*x = AddBlockLinkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockio_proto_blockio_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBlockLinkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBlockLinkResponse) ProtoMessage() {}

func (x *AddBlockLinkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blockio_proto_blockio_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBlockLinkResponse.ProtoReflect.Descriptor instead.
func (*AddBlockLinkResponse) Descriptor() ([]byte, []int) {
	return file_blockio_proto_blockio_proto_rawDescGZIP(), []int{2}
}

func (x *AddBlockLinkResponse) GetResult() *proto.ResMes {
	if x != nil {
		return x.Result
	}
	return nil
}

// delete block link ID
type DelBlockLinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockLinkInfo *BlockLinkInfo `protobuf:"bytes,1,opt,name=block_link_info,json=blockLinkInfo,proto3" json:"block_link_info,omitempty"`
}

func (x *DelBlockLinkRequest) Reset() {
	*x = DelBlockLinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockio_proto_blockio_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelBlockLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelBlockLinkRequest) ProtoMessage() {}

func (x *DelBlockLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blockio_proto_blockio_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelBlockLinkRequest.ProtoReflect.Descriptor instead.
func (*DelBlockLinkRequest) Descriptor() ([]byte, []int) {
	return file_blockio_proto_blockio_proto_rawDescGZIP(), []int{3}
}

func (x *DelBlockLinkRequest) GetBlockLinkInfo() *BlockLinkInfo {
	if x != nil {
		return x.BlockLinkInfo
	}
	return nil
}

type DelBlockLinkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *proto.ResMes `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *DelBlockLinkResponse) Reset() {
	*x = DelBlockLinkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockio_proto_blockio_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelBlockLinkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelBlockLinkResponse) ProtoMessage() {}

func (x *DelBlockLinkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blockio_proto_blockio_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelBlockLinkResponse.ProtoReflect.Descriptor instead.
func (*DelBlockLinkResponse) Descriptor() ([]byte, []int) {
	return file_blockio_proto_blockio_proto_rawDescGZIP(), []int{4}
}

func (x *DelBlockLinkResponse) GetResult() *proto.ResMes {
	if x != nil {
		return x.Result
	}
	return nil
}

// list block links ID
type ListBlockLinksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListBlockLinksRequest) Reset() {
	*x = ListBlockLinksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockio_proto_blockio_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBlockLinksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBlockLinksRequest) ProtoMessage() {}

func (x *ListBlockLinksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blockio_proto_blockio_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBlockLinksRequest.ProtoReflect.Descriptor instead.
func (*ListBlockLinksRequest) Descriptor() ([]byte, []int) {
	return file_blockio_proto_blockio_proto_rawDescGZIP(), []int{5}
}

type ListBlockLinksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result    *proto.ResMes    `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	BlockInfo []*BlockLinkInfo `protobuf:"bytes,2,rep,name=block_info,json=blockInfo,proto3" json:"block_info,omitempty"`
}

func (x *ListBlockLinksResponse) Reset() {
	*x = ListBlockLinksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockio_proto_blockio_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBlockLinksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBlockLinksResponse) ProtoMessage() {}

func (x *ListBlockLinksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blockio_proto_blockio_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBlockLinksResponse.ProtoReflect.Descriptor instead.
func (*ListBlockLinksResponse) Descriptor() ([]byte, []int) {
	return file_blockio_proto_blockio_proto_rawDescGZIP(), []int{6}
}

func (x *ListBlockLinksResponse) GetResult() *proto.ResMes {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *ListBlockLinksResponse) GetBlockInfo() []*BlockLinkInfo {
	if x != nil {
		return x.BlockInfo
	}
	return nil
}

var File_blockio_proto_blockio_proto protoreflect.FileDescriptor

var file_blockio_proto_blockio_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x69, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x59,
	0x72, 0x63, 0x66, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x1a, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x01, 0x0a, 0x0d, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x69,
	0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6c, 0x69, 0x6e,
	0x6b, 0x49, 0x64, 0x12, 0x41, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x59, 0x72, 0x63, 0x66, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x22, 0x2f, 0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x49, 0x4d, 0x50,
	0x4f, 0x52, 0x54, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x45,
	0x58, 0x50, 0x4f, 0x52, 0x54, 0x10, 0x01, 0x22, 0x57, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40,
	0x0a, 0x0f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x59, 0x72, 0x63, 0x66, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0d, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0x37, 0x0a, 0x14, 0x41, 0x64, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x6e, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x4d, 0x65,
	0x73, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x57, 0x0a, 0x13, 0x44, 0x65, 0x6c,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x40, 0x0a, 0x0f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x59, 0x72, 0x63, 0x66,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x6e, 0x6b, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x6e,
	0x66, 0x6f, 0x22, 0x37, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4c, 0x69,
	0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x52, 0x65, 0x73,
	0x4d, 0x65, 0x73, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x17, 0x0a, 0x15, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x72, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07,
	0x2e, 0x52, 0x65, 0x73, 0x4d, 0x65, 0x73, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x37, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x59, 0x72, 0x63, 0x66, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0x82, 0x02, 0x0a, 0x07, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x49, 0x4f, 0x12, 0x55, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x20, 0x2e, 0x59, 0x72, 0x63, 0x66, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x6e, 0x6b,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x59, 0x72, 0x63, 0x66, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4c, 0x69,
	0x6e, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0c, 0x41,
	0x64, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x1e, 0x2e, 0x59, 0x72,
	0x63, 0x66, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x59, 0x72,
	0x63, 0x66, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0c,
	0x44, 0x65, 0x6c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x1e, 0x2e, 0x59,
	0x72, 0x63, 0x66, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x59,
	0x72, 0x63, 0x66, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2c, 0x5a,
	0x2a, 0x79, 0x72, 0x63, 0x66, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x69, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x3b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x69, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_blockio_proto_blockio_proto_rawDescOnce sync.Once
	file_blockio_proto_blockio_proto_rawDescData = file_blockio_proto_blockio_proto_rawDesc
)

func file_blockio_proto_blockio_proto_rawDescGZIP() []byte {
	file_blockio_proto_blockio_proto_rawDescOnce.Do(func() {
		file_blockio_proto_blockio_proto_rawDescData = protoimpl.X.CompressGZIP(file_blockio_proto_blockio_proto_rawDescData)
	})
	return file_blockio_proto_blockio_proto_rawDescData
}

var file_blockio_proto_blockio_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_blockio_proto_blockio_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_blockio_proto_blockio_proto_goTypes = []any{
	(BlockLinkInfo_BlockType)(0),   // 0: YrcfAgent.BlockLinkInfo.BlockType
	(*BlockLinkInfo)(nil),          // 1: YrcfAgent.BlockLinkInfo
	(*AddBlockLinkRequest)(nil),    // 2: YrcfAgent.AddBlockLinkRequest
	(*AddBlockLinkResponse)(nil),   // 3: YrcfAgent.AddBlockLinkResponse
	(*DelBlockLinkRequest)(nil),    // 4: YrcfAgent.DelBlockLinkRequest
	(*DelBlockLinkResponse)(nil),   // 5: YrcfAgent.DelBlockLinkResponse
	(*ListBlockLinksRequest)(nil),  // 6: YrcfAgent.ListBlockLinksRequest
	(*ListBlockLinksResponse)(nil), // 7: YrcfAgent.ListBlockLinksResponse
	(*proto.ResMes)(nil),           // 8: ResMes
}
var file_blockio_proto_blockio_proto_depIdxs = []int32{
	0,  // 0: YrcfAgent.BlockLinkInfo.block_type:type_name -> YrcfAgent.BlockLinkInfo.BlockType
	1,  // 1: YrcfAgent.AddBlockLinkRequest.block_link_info:type_name -> YrcfAgent.BlockLinkInfo
	8,  // 2: YrcfAgent.AddBlockLinkResponse.result:type_name -> ResMes
	1,  // 3: YrcfAgent.DelBlockLinkRequest.block_link_info:type_name -> YrcfAgent.BlockLinkInfo
	8,  // 4: YrcfAgent.DelBlockLinkResponse.result:type_name -> ResMes
	8,  // 5: YrcfAgent.ListBlockLinksResponse.result:type_name -> ResMes
	1,  // 6: YrcfAgent.ListBlockLinksResponse.block_info:type_name -> YrcfAgent.BlockLinkInfo
	6,  // 7: YrcfAgent.BlockIO.ListBlockLinks:input_type -> YrcfAgent.ListBlockLinksRequest
	2,  // 8: YrcfAgent.BlockIO.AddBlockLink:input_type -> YrcfAgent.AddBlockLinkRequest
	4,  // 9: YrcfAgent.BlockIO.DelBlockLink:input_type -> YrcfAgent.DelBlockLinkRequest
	7,  // 10: YrcfAgent.BlockIO.ListBlockLinks:output_type -> YrcfAgent.ListBlockLinksResponse
	3,  // 11: YrcfAgent.BlockIO.AddBlockLink:output_type -> YrcfAgent.AddBlockLinkResponse
	5,  // 12: YrcfAgent.BlockIO.DelBlockLink:output_type -> YrcfAgent.DelBlockLinkResponse
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_blockio_proto_blockio_proto_init() }
func file_blockio_proto_blockio_proto_init() {
	if File_blockio_proto_blockio_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blockio_proto_blockio_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*BlockLinkInfo); i {
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
		file_blockio_proto_blockio_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AddBlockLinkRequest); i {
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
		file_blockio_proto_blockio_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*AddBlockLinkResponse); i {
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
		file_blockio_proto_blockio_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*DelBlockLinkRequest); i {
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
		file_blockio_proto_blockio_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*DelBlockLinkResponse); i {
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
		file_blockio_proto_blockio_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ListBlockLinksRequest); i {
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
		file_blockio_proto_blockio_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ListBlockLinksResponse); i {
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
			RawDescriptor: file_blockio_proto_blockio_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_blockio_proto_blockio_proto_goTypes,
		DependencyIndexes: file_blockio_proto_blockio_proto_depIdxs,
		EnumInfos:         file_blockio_proto_blockio_proto_enumTypes,
		MessageInfos:      file_blockio_proto_blockio_proto_msgTypes,
	}.Build()
	File_blockio_proto_blockio_proto = out.File
	file_blockio_proto_blockio_proto_rawDesc = nil
	file_blockio_proto_blockio_proto_goTypes = nil
	file_blockio_proto_blockio_proto_depIdxs = nil
}
