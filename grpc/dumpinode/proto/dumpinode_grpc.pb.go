// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: dumpinode/proto/dumpinode.proto

package dumpinode

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
	DumpInode_DumpInode_FullMethodName = "/YrcfAgent.DumpInode/DumpInode"
)

// DumpInodeClient is the client API for DumpInode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DumpInodeClient interface {
	DumpInode(ctx context.Context, in *DumpInodeRequest, opts ...grpc.CallOption) (*DumpInodeResponse, error)
}

type dumpInodeClient struct {
	cc grpc.ClientConnInterface
}

func NewDumpInodeClient(cc grpc.ClientConnInterface) DumpInodeClient {
	return &dumpInodeClient{cc}
}

func (c *dumpInodeClient) DumpInode(ctx context.Context, in *DumpInodeRequest, opts ...grpc.CallOption) (*DumpInodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DumpInodeResponse)
	err := c.cc.Invoke(ctx, DumpInode_DumpInode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DumpInodeServer is the server API for DumpInode service.
// All implementations must embed UnimplementedDumpInodeServer
// for forward compatibility
type DumpInodeServer interface {
	DumpInode(context.Context, *DumpInodeRequest) (*DumpInodeResponse, error)
	mustEmbedUnimplementedDumpInodeServer()
}

// UnimplementedDumpInodeServer must be embedded to have forward compatible implementations.
type UnimplementedDumpInodeServer struct {
}

func (UnimplementedDumpInodeServer) DumpInode(context.Context, *DumpInodeRequest) (*DumpInodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DumpInode not implemented")
}
func (UnimplementedDumpInodeServer) mustEmbedUnimplementedDumpInodeServer() {}

// UnsafeDumpInodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DumpInodeServer will
// result in compilation errors.
type UnsafeDumpInodeServer interface {
	mustEmbedUnimplementedDumpInodeServer()
}

func RegisterDumpInodeServer(s grpc.ServiceRegistrar, srv DumpInodeServer) {
	s.RegisterService(&DumpInode_ServiceDesc, srv)
}

func _DumpInode_DumpInode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DumpInodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DumpInodeServer).DumpInode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DumpInode_DumpInode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DumpInodeServer).DumpInode(ctx, req.(*DumpInodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DumpInode_ServiceDesc is the grpc.ServiceDesc for DumpInode service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DumpInode_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "YrcfAgent.DumpInode",
	HandlerType: (*DumpInodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DumpInode",
			Handler:    _DumpInode_DumpInode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dumpinode/proto/dumpinode.proto",
}
