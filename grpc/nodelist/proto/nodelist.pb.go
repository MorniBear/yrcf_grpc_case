// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: nodelist/proto/nodelist.proto

package nodelist

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

type NodeInfo_NodeType int32

const (
	NodeInfo_MDS    NodeInfo_NodeType = 0
	NodeInfo_OSS    NodeInfo_NodeType = 1
	NodeInfo_CLIENT NodeInfo_NodeType = 2
	NodeInfo_MGR    NodeInfo_NodeType = 3
	NodeInfo_AGENT  NodeInfo_NodeType = 4
)

// Enum value maps for NodeInfo_NodeType.
var (
	NodeInfo_NodeType_name = map[int32]string{
		0: "MDS",
		1: "OSS",
		2: "CLIENT",
		3: "MGR",
		4: "AGENT",
	}
	NodeInfo_NodeType_value = map[string]int32{
		"MDS":    0,
		"OSS":    1,
		"CLIENT": 2,
		"MGR":    3,
		"AGENT":  4,
	}
)

func (x NodeInfo_NodeType) Enum() *NodeInfo_NodeType {
	p := new(NodeInfo_NodeType)
	*p = x
	return p
}

func (x NodeInfo_NodeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NodeInfo_NodeType) Descriptor() protoreflect.EnumDescriptor {
	return file_nodelist_proto_nodelist_proto_enumTypes[0].Descriptor()
}

func (NodeInfo_NodeType) Type() protoreflect.EnumType {
	return &file_nodelist_proto_nodelist_proto_enumTypes[0]
}

func (x NodeInfo_NodeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NodeInfo_NodeType.Descriptor instead.
func (NodeInfo_NodeType) EnumDescriptor() ([]byte, []int) {
	return file_nodelist_proto_nodelist_proto_rawDescGZIP(), []int{1, 0}
}

// grpc client: get node list
type NodeListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Client          bool `protobuf:"varint,1,opt,name=client,proto3" json:"client,omitempty"`
	HideInternalIps bool `protobuf:"varint,2,opt,name=hide_internal_ips,json=hideInternalIps,proto3" json:"hide_internal_ips,omitempty"`
	Agent           bool `protobuf:"varint,3,opt,name=agent,proto3" json:"agent,omitempty"`
}

func (x *NodeListRequest) Reset() {
	*x = NodeListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodelist_proto_nodelist_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeListRequest) ProtoMessage() {}

func (x *NodeListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nodelist_proto_nodelist_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeListRequest.ProtoReflect.Descriptor instead.
func (*NodeListRequest) Descriptor() ([]byte, []int) {
	return file_nodelist_proto_nodelist_proto_rawDescGZIP(), []int{0}
}

func (x *NodeListRequest) GetClient() bool {
	if x != nil {
		return x.Client
	}
	return false
}

func (x *NodeListRequest) GetHideInternalIps() bool {
	if x != nil {
		return x.HideInternalIps
	}
	return false
}

func (x *NodeListRequest) GetAgent() bool {
	if x != nil {
		return x.Agent
	}
	return false
}

// node information
type NodeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      NodeInfo_NodeType `protobuf:"varint,1,opt,name=type,proto3,enum=YrcfAgent.NodeInfo_NodeType" json:"type,omitempty"`
	NodeNumId uint32            `protobuf:"varint,2,opt,name=node_num_id,json=nodeNumId,proto3" json:"node_num_id,omitempty"`
	NodeName  string            `protobuf:"bytes,3,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
}

func (x *NodeInfo) Reset() {
	*x = NodeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodelist_proto_nodelist_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeInfo) ProtoMessage() {}

func (x *NodeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_nodelist_proto_nodelist_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeInfo.ProtoReflect.Descriptor instead.
func (*NodeInfo) Descriptor() ([]byte, []int) {
	return file_nodelist_proto_nodelist_proto_rawDescGZIP(), []int{1}
}

func (x *NodeInfo) GetType() NodeInfo_NodeType {
	if x != nil {
		return x.Type
	}
	return NodeInfo_MDS
}

func (x *NodeInfo) GetNodeNumId() uint32 {
	if x != nil {
		return x.NodeNumId
	}
	return 0
}

func (x *NodeInfo) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

// grpc server: return node list
type NodeListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result    *proto.ResMes `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	NodeLists []*NodeInfo   `protobuf:"bytes,2,rep,name=node_lists,json=nodeLists,proto3" json:"node_lists,omitempty"`
}

func (x *NodeListResponse) Reset() {
	*x = NodeListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodelist_proto_nodelist_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeListResponse) ProtoMessage() {}

func (x *NodeListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nodelist_proto_nodelist_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeListResponse.ProtoReflect.Descriptor instead.
func (*NodeListResponse) Descriptor() ([]byte, []int) {
	return file_nodelist_proto_nodelist_proto_rawDescGZIP(), []int{2}
}

func (x *NodeListResponse) GetResult() *proto.ResMes {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *NodeListResponse) GetNodeLists() []*NodeInfo {
	if x != nil {
		return x.NodeLists
	}
	return nil
}

var File_nodelist_proto_nodelist_proto protoreflect.FileDescriptor

var file_nodelist_proto_nodelist_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6e, 0x6f, 0x64, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x59, 0x72, 0x63, 0x66, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x1a, 0x0e, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x0f, 0x4e, 0x6f,
	0x64, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x68, 0x69, 0x64, 0x65, 0x5f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x70, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0f, 0x68, 0x69, 0x64, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x70,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x22, 0xb7, 0x01, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x30, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x59, 0x72, 0x63, 0x66, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x4e,
	0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0b, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6e,
	0x75, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6e, 0x6f, 0x64,
	0x65, 0x4e, 0x75, 0x6d, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x3c, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x07, 0x0a, 0x03, 0x4d, 0x44, 0x53, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4f, 0x53, 0x53, 0x10,
	0x01, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x07, 0x0a,
	0x03, 0x4d, 0x47, 0x52, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x10,
	0x04, 0x22, 0x67, 0x0a, 0x10, 0x4e, 0x6f, 0x64, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x4d, 0x65, 0x73, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x32, 0x0a, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x59, 0x72, 0x63,
	0x66, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x09, 0x6e, 0x6f, 0x64, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x32, 0x51, 0x0a, 0x08, 0x4e, 0x6f,
	0x64, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x1a, 0x2e, 0x59, 0x72, 0x63, 0x66, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x4e,
	0x6f, 0x64, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x59, 0x72, 0x63, 0x66, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2e, 0x5a,
	0x2c, 0x79, 0x72, 0x63, 0x66, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x3b, 0x6e, 0x6f, 0x64, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nodelist_proto_nodelist_proto_rawDescOnce sync.Once
	file_nodelist_proto_nodelist_proto_rawDescData = file_nodelist_proto_nodelist_proto_rawDesc
)

func file_nodelist_proto_nodelist_proto_rawDescGZIP() []byte {
	file_nodelist_proto_nodelist_proto_rawDescOnce.Do(func() {
		file_nodelist_proto_nodelist_proto_rawDescData = protoimpl.X.CompressGZIP(file_nodelist_proto_nodelist_proto_rawDescData)
	})
	return file_nodelist_proto_nodelist_proto_rawDescData
}

var file_nodelist_proto_nodelist_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_nodelist_proto_nodelist_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_nodelist_proto_nodelist_proto_goTypes = []any{
	(NodeInfo_NodeType)(0),   // 0: YrcfAgent.NodeInfo.NodeType
	(*NodeListRequest)(nil),  // 1: YrcfAgent.NodeListRequest
	(*NodeInfo)(nil),         // 2: YrcfAgent.NodeInfo
	(*NodeListResponse)(nil), // 3: YrcfAgent.NodeListResponse
	(*proto.ResMes)(nil),     // 4: ResMes
}
var file_nodelist_proto_nodelist_proto_depIdxs = []int32{
	0, // 0: YrcfAgent.NodeInfo.type:type_name -> YrcfAgent.NodeInfo.NodeType
	4, // 1: YrcfAgent.NodeListResponse.result:type_name -> ResMes
	2, // 2: YrcfAgent.NodeListResponse.node_lists:type_name -> YrcfAgent.NodeInfo
	1, // 3: YrcfAgent.NodeList.NodeList:input_type -> YrcfAgent.NodeListRequest
	3, // 4: YrcfAgent.NodeList.NodeList:output_type -> YrcfAgent.NodeListResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_nodelist_proto_nodelist_proto_init() }
func file_nodelist_proto_nodelist_proto_init() {
	if File_nodelist_proto_nodelist_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nodelist_proto_nodelist_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*NodeListRequest); i {
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
		file_nodelist_proto_nodelist_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*NodeInfo); i {
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
		file_nodelist_proto_nodelist_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*NodeListResponse); i {
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
			RawDescriptor: file_nodelist_proto_nodelist_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nodelist_proto_nodelist_proto_goTypes,
		DependencyIndexes: file_nodelist_proto_nodelist_proto_depIdxs,
		EnumInfos:         file_nodelist_proto_nodelist_proto_enumTypes,
		MessageInfos:      file_nodelist_proto_nodelist_proto_msgTypes,
	}.Build()
	File_nodelist_proto_nodelist_proto = out.File
	file_nodelist_proto_nodelist_proto_rawDesc = nil
	file_nodelist_proto_nodelist_proto_goTypes = nil
	file_nodelist_proto_nodelist_proto_depIdxs = nil
}
