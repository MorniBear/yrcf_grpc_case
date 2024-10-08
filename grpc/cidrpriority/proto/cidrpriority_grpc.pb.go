// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: cidrpriority/proto/cidrpriority.proto

package cidrpriority

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
	CidrPriority_GetCidrPriority_FullMethodName = "/YrcfAgent.CidrPriority/GetCidrPriority"
	CidrPriority_SetCidrPriority_FullMethodName = "/YrcfAgent.CidrPriority/SetCidrPriority"
)

// CidrPriorityClient is the client API for CidrPriority service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CidrPriorityClient interface {
	GetCidrPriority(ctx context.Context, in *GetCidrPriRequest, opts ...grpc.CallOption) (*GetCidrPriResponse, error)
	SetCidrPriority(ctx context.Context, in *SetCidrPriRequest, opts ...grpc.CallOption) (*SetCidrPriResponse, error)
}

type cidrPriorityClient struct {
	cc grpc.ClientConnInterface
}

func NewCidrPriorityClient(cc grpc.ClientConnInterface) CidrPriorityClient {
	return &cidrPriorityClient{cc}
}

func (c *cidrPriorityClient) GetCidrPriority(ctx context.Context, in *GetCidrPriRequest, opts ...grpc.CallOption) (*GetCidrPriResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCidrPriResponse)
	err := c.cc.Invoke(ctx, CidrPriority_GetCidrPriority_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cidrPriorityClient) SetCidrPriority(ctx context.Context, in *SetCidrPriRequest, opts ...grpc.CallOption) (*SetCidrPriResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetCidrPriResponse)
	err := c.cc.Invoke(ctx, CidrPriority_SetCidrPriority_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CidrPriorityServer is the server API for CidrPriority service.
// All implementations must embed UnimplementedCidrPriorityServer
// for forward compatibility
type CidrPriorityServer interface {
	GetCidrPriority(context.Context, *GetCidrPriRequest) (*GetCidrPriResponse, error)
	SetCidrPriority(context.Context, *SetCidrPriRequest) (*SetCidrPriResponse, error)
	mustEmbedUnimplementedCidrPriorityServer()
}

// UnimplementedCidrPriorityServer must be embedded to have forward compatible implementations.
type UnimplementedCidrPriorityServer struct {
}

func (UnimplementedCidrPriorityServer) GetCidrPriority(context.Context, *GetCidrPriRequest) (*GetCidrPriResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCidrPriority not implemented")
}
func (UnimplementedCidrPriorityServer) SetCidrPriority(context.Context, *SetCidrPriRequest) (*SetCidrPriResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCidrPriority not implemented")
}
func (UnimplementedCidrPriorityServer) mustEmbedUnimplementedCidrPriorityServer() {}

// UnsafeCidrPriorityServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CidrPriorityServer will
// result in compilation errors.
type UnsafeCidrPriorityServer interface {
	mustEmbedUnimplementedCidrPriorityServer()
}

func RegisterCidrPriorityServer(s grpc.ServiceRegistrar, srv CidrPriorityServer) {
	s.RegisterService(&CidrPriority_ServiceDesc, srv)
}

func _CidrPriority_GetCidrPriority_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCidrPriRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CidrPriorityServer).GetCidrPriority(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CidrPriority_GetCidrPriority_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CidrPriorityServer).GetCidrPriority(ctx, req.(*GetCidrPriRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CidrPriority_SetCidrPriority_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetCidrPriRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CidrPriorityServer).SetCidrPriority(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CidrPriority_SetCidrPriority_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CidrPriorityServer).SetCidrPriority(ctx, req.(*SetCidrPriRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CidrPriority_ServiceDesc is the grpc.ServiceDesc for CidrPriority service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CidrPriority_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "YrcfAgent.CidrPriority",
	HandlerType: (*CidrPriorityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCidrPriority",
			Handler:    _CidrPriority_GetCidrPriority_Handler,
		},
		{
			MethodName: "SetCidrPriority",
			Handler:    _CidrPriority_SetCidrPriority_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cidrpriority/proto/cidrpriority.proto",
}
