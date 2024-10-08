// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: acl/proto/acl.proto

package acl

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
	Acl_AddAcl_FullMethodName    = "/YrcfAgent.Acl/AddAcl"
	Acl_ListAcl_FullMethodName   = "/YrcfAgent.Acl/ListAcl"
	Acl_DeleteAcl_FullMethodName = "/YrcfAgent.Acl/DeleteAcl"
)

// AclClient is the client API for Acl service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AclClient interface {
	AddAcl(ctx context.Context, in *AddAclRequest, opts ...grpc.CallOption) (*AddAclResponse, error)
	ListAcl(ctx context.Context, in *ListAclRequest, opts ...grpc.CallOption) (*ListAclResponse, error)
	DeleteAcl(ctx context.Context, in *DeleteAclRequest, opts ...grpc.CallOption) (*DeleteAclResponse, error)
}

type aclClient struct {
	cc grpc.ClientConnInterface
}

func NewAclClient(cc grpc.ClientConnInterface) AclClient {
	return &aclClient{cc}
}

func (c *aclClient) AddAcl(ctx context.Context, in *AddAclRequest, opts ...grpc.CallOption) (*AddAclResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddAclResponse)
	err := c.cc.Invoke(ctx, Acl_AddAcl_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aclClient) ListAcl(ctx context.Context, in *ListAclRequest, opts ...grpc.CallOption) (*ListAclResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListAclResponse)
	err := c.cc.Invoke(ctx, Acl_ListAcl_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aclClient) DeleteAcl(ctx context.Context, in *DeleteAclRequest, opts ...grpc.CallOption) (*DeleteAclResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteAclResponse)
	err := c.cc.Invoke(ctx, Acl_DeleteAcl_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AclServer is the server API for Acl service.
// All implementations must embed UnimplementedAclServer
// for forward compatibility
type AclServer interface {
	AddAcl(context.Context, *AddAclRequest) (*AddAclResponse, error)
	ListAcl(context.Context, *ListAclRequest) (*ListAclResponse, error)
	DeleteAcl(context.Context, *DeleteAclRequest) (*DeleteAclResponse, error)
	mustEmbedUnimplementedAclServer()
}

// UnimplementedAclServer must be embedded to have forward compatible implementations.
type UnimplementedAclServer struct {
}

func (UnimplementedAclServer) AddAcl(context.Context, *AddAclRequest) (*AddAclResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAcl not implemented")
}
func (UnimplementedAclServer) ListAcl(context.Context, *ListAclRequest) (*ListAclResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAcl not implemented")
}
func (UnimplementedAclServer) DeleteAcl(context.Context, *DeleteAclRequest) (*DeleteAclResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAcl not implemented")
}
func (UnimplementedAclServer) mustEmbedUnimplementedAclServer() {}

// UnsafeAclServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AclServer will
// result in compilation errors.
type UnsafeAclServer interface {
	mustEmbedUnimplementedAclServer()
}

func RegisterAclServer(s grpc.ServiceRegistrar, srv AclServer) {
	s.RegisterService(&Acl_ServiceDesc, srv)
}

func _Acl_AddAcl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAclRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AclServer).AddAcl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Acl_AddAcl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AclServer).AddAcl(ctx, req.(*AddAclRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Acl_ListAcl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAclRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AclServer).ListAcl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Acl_ListAcl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AclServer).ListAcl(ctx, req.(*ListAclRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Acl_DeleteAcl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAclRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AclServer).DeleteAcl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Acl_DeleteAcl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AclServer).DeleteAcl(ctx, req.(*DeleteAclRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Acl_ServiceDesc is the grpc.ServiceDesc for Acl service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Acl_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "YrcfAgent.Acl",
	HandlerType: (*AclServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddAcl",
			Handler:    _Acl_AddAcl_Handler,
		},
		{
			MethodName: "ListAcl",
			Handler:    _Acl_ListAcl_Handler,
		},
		{
			MethodName: "DeleteAcl",
			Handler:    _Acl_DeleteAcl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "acl/proto/acl.proto",
}
