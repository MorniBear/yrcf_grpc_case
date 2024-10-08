// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: projectquota/proto/projectquota.proto

package projectquota

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
	ProjectQuota_AddOrUpdateProjectQuota_FullMethodName = "/YrcfAgent.ProjectQuota/AddOrUpdateProjectQuota"
	ProjectQuota_ListProjectQuota_FullMethodName        = "/YrcfAgent.ProjectQuota/ListProjectQuota"
	ProjectQuota_DeleteProjectQuota_FullMethodName      = "/YrcfAgent.ProjectQuota/DeleteProjectQuota"
)

// ProjectQuotaClient is the client API for ProjectQuota service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProjectQuotaClient interface {
	AddOrUpdateProjectQuota(ctx context.Context, in *AddOrUpdateProjectQuotaRequest, opts ...grpc.CallOption) (*AddOrUpdateProjectQuotaResponse, error)
	ListProjectQuota(ctx context.Context, in *ListProjectQuotaRequest, opts ...grpc.CallOption) (*ListProjectQuotaResponse, error)
	DeleteProjectQuota(ctx context.Context, in *DeleteProjectQuotaRequest, opts ...grpc.CallOption) (*DeleteProjectQuotaResponse, error)
}

type projectQuotaClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectQuotaClient(cc grpc.ClientConnInterface) ProjectQuotaClient {
	return &projectQuotaClient{cc}
}

func (c *projectQuotaClient) AddOrUpdateProjectQuota(ctx context.Context, in *AddOrUpdateProjectQuotaRequest, opts ...grpc.CallOption) (*AddOrUpdateProjectQuotaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddOrUpdateProjectQuotaResponse)
	err := c.cc.Invoke(ctx, ProjectQuota_AddOrUpdateProjectQuota_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectQuotaClient) ListProjectQuota(ctx context.Context, in *ListProjectQuotaRequest, opts ...grpc.CallOption) (*ListProjectQuotaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListProjectQuotaResponse)
	err := c.cc.Invoke(ctx, ProjectQuota_ListProjectQuota_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectQuotaClient) DeleteProjectQuota(ctx context.Context, in *DeleteProjectQuotaRequest, opts ...grpc.CallOption) (*DeleteProjectQuotaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteProjectQuotaResponse)
	err := c.cc.Invoke(ctx, ProjectQuota_DeleteProjectQuota_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectQuotaServer is the server API for ProjectQuota service.
// All implementations must embed UnimplementedProjectQuotaServer
// for forward compatibility
type ProjectQuotaServer interface {
	AddOrUpdateProjectQuota(context.Context, *AddOrUpdateProjectQuotaRequest) (*AddOrUpdateProjectQuotaResponse, error)
	ListProjectQuota(context.Context, *ListProjectQuotaRequest) (*ListProjectQuotaResponse, error)
	DeleteProjectQuota(context.Context, *DeleteProjectQuotaRequest) (*DeleteProjectQuotaResponse, error)
	mustEmbedUnimplementedProjectQuotaServer()
}

// UnimplementedProjectQuotaServer must be embedded to have forward compatible implementations.
type UnimplementedProjectQuotaServer struct {
}

func (UnimplementedProjectQuotaServer) AddOrUpdateProjectQuota(context.Context, *AddOrUpdateProjectQuotaRequest) (*AddOrUpdateProjectQuotaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddOrUpdateProjectQuota not implemented")
}
func (UnimplementedProjectQuotaServer) ListProjectQuota(context.Context, *ListProjectQuotaRequest) (*ListProjectQuotaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProjectQuota not implemented")
}
func (UnimplementedProjectQuotaServer) DeleteProjectQuota(context.Context, *DeleteProjectQuotaRequest) (*DeleteProjectQuotaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProjectQuota not implemented")
}
func (UnimplementedProjectQuotaServer) mustEmbedUnimplementedProjectQuotaServer() {}

// UnsafeProjectQuotaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProjectQuotaServer will
// result in compilation errors.
type UnsafeProjectQuotaServer interface {
	mustEmbedUnimplementedProjectQuotaServer()
}

func RegisterProjectQuotaServer(s grpc.ServiceRegistrar, srv ProjectQuotaServer) {
	s.RegisterService(&ProjectQuota_ServiceDesc, srv)
}

func _ProjectQuota_AddOrUpdateProjectQuota_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddOrUpdateProjectQuotaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectQuotaServer).AddOrUpdateProjectQuota(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectQuota_AddOrUpdateProjectQuota_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectQuotaServer).AddOrUpdateProjectQuota(ctx, req.(*AddOrUpdateProjectQuotaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectQuota_ListProjectQuota_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProjectQuotaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectQuotaServer).ListProjectQuota(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectQuota_ListProjectQuota_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectQuotaServer).ListProjectQuota(ctx, req.(*ListProjectQuotaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectQuota_DeleteProjectQuota_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProjectQuotaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectQuotaServer).DeleteProjectQuota(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectQuota_DeleteProjectQuota_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectQuotaServer).DeleteProjectQuota(ctx, req.(*DeleteProjectQuotaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProjectQuota_ServiceDesc is the grpc.ServiceDesc for ProjectQuota service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProjectQuota_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "YrcfAgent.ProjectQuota",
	HandlerType: (*ProjectQuotaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddOrUpdateProjectQuota",
			Handler:    _ProjectQuota_AddOrUpdateProjectQuota_Handler,
		},
		{
			MethodName: "ListProjectQuota",
			Handler:    _ProjectQuota_ListProjectQuota_Handler,
		},
		{
			MethodName: "DeleteProjectQuota",
			Handler:    _ProjectQuota_DeleteProjectQuota_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "projectquota/proto/projectquota.proto",
}
