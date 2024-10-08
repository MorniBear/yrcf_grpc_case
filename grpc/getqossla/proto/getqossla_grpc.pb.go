// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: getqossla/proto/getqossla.proto

package getqossla

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
	GetQosSla_GetQosSla_FullMethodName = "/YrcfAgent.GetQosSla/GetQosSla"
)

// GetQosSlaClient is the client API for GetQosSla service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetQosSlaClient interface {
	GetQosSla(ctx context.Context, in *GetQosSlaRequest, opts ...grpc.CallOption) (*GetQosSlaResponse, error)
}

type getQosSlaClient struct {
	cc grpc.ClientConnInterface
}

func NewGetQosSlaClient(cc grpc.ClientConnInterface) GetQosSlaClient {
	return &getQosSlaClient{cc}
}

func (c *getQosSlaClient) GetQosSla(ctx context.Context, in *GetQosSlaRequest, opts ...grpc.CallOption) (*GetQosSlaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetQosSlaResponse)
	err := c.cc.Invoke(ctx, GetQosSla_GetQosSla_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetQosSlaServer is the server API for GetQosSla service.
// All implementations must embed UnimplementedGetQosSlaServer
// for forward compatibility
type GetQosSlaServer interface {
	GetQosSla(context.Context, *GetQosSlaRequest) (*GetQosSlaResponse, error)
	mustEmbedUnimplementedGetQosSlaServer()
}

// UnimplementedGetQosSlaServer must be embedded to have forward compatible implementations.
type UnimplementedGetQosSlaServer struct {
}

func (UnimplementedGetQosSlaServer) GetQosSla(context.Context, *GetQosSlaRequest) (*GetQosSlaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQosSla not implemented")
}
func (UnimplementedGetQosSlaServer) mustEmbedUnimplementedGetQosSlaServer() {}

// UnsafeGetQosSlaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetQosSlaServer will
// result in compilation errors.
type UnsafeGetQosSlaServer interface {
	mustEmbedUnimplementedGetQosSlaServer()
}

func RegisterGetQosSlaServer(s grpc.ServiceRegistrar, srv GetQosSlaServer) {
	s.RegisterService(&GetQosSla_ServiceDesc, srv)
}

func _GetQosSla_GetQosSla_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQosSlaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetQosSlaServer).GetQosSla(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GetQosSla_GetQosSla_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetQosSlaServer).GetQosSla(ctx, req.(*GetQosSlaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GetQosSla_ServiceDesc is the grpc.ServiceDesc for GetQosSla service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetQosSla_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "YrcfAgent.GetQosSla",
	HandlerType: (*GetQosSlaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetQosSla",
			Handler:    _GetQosSla_GetQosSla_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "getqossla/proto/getqossla.proto",
}
