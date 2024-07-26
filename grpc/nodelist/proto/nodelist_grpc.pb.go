// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: nodelist/proto/nodelist.proto

package nodelist

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	NodeList_NodeList_FullMethodName = "/YrcfAgent.NodeList/NodeList"
)

// NodeListClient is the client API for NodeList service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeListClient interface {
	NodeList(ctx context.Context, in *NodeListRequest, opts ...grpc.CallOption) (*NodeListResponse, error)
}

type nodeListClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeListClient(cc grpc.ClientConnInterface) NodeListClient {
	return &nodeListClient{cc}
}

func (c *nodeListClient) NodeList(ctx context.Context, in *NodeListRequest, opts ...grpc.CallOption) (*NodeListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NodeListResponse)
	err := c.cc.Invoke(ctx, NodeList_NodeList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeListServer is the server API for NodeList service.
// All implementations must embed UnimplementedNodeListServer
// for forward compatibility
type NodeListServer interface {
	NodeList(context.Context, *NodeListRequest) (*NodeListResponse, error)
	mustEmbedUnimplementedNodeListServer()
}

// UnimplementedNodeListServer must be embedded to have forward compatible implementations.
type UnimplementedNodeListServer struct {
}

func (UnimplementedNodeListServer) NodeList(context.Context, *NodeListRequest) (*NodeListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeList not implemented")
}
func (UnimplementedNodeListServer) mustEmbedUnimplementedNodeListServer() {}

// UnsafeNodeListServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeListServer will
// result in compilation errors.
type UnsafeNodeListServer interface {
	mustEmbedUnimplementedNodeListServer()
}

func RegisterNodeListServer(s grpc.ServiceRegistrar, srv NodeListServer) {
	s.RegisterService(&NodeList_ServiceDesc, srv)
}

func _NodeList_NodeList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeListServer).NodeList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeList_NodeList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeListServer).NodeList(ctx, req.(*NodeListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NodeList_ServiceDesc is the grpc.ServiceDesc for NodeList service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodeList_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "YrcfAgent.NodeList",
	HandlerType: (*NodeListServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NodeList",
			Handler:    _NodeList_NodeList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nodelist/proto/nodelist.proto",
}
