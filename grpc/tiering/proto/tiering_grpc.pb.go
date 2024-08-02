// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: tiering/proto/tiering.proto

package tiering

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
	Tiering_AddTiering_FullMethodName         = "/YrcfAgent.Tiering/AddTiering"
	Tiering_DeleteTiering_FullMethodName      = "/YrcfAgent.Tiering/DeleteTiering"
	Tiering_UpdateTiering_FullMethodName      = "/YrcfAgent.Tiering/UpdateTiering"
	Tiering_FlushTiering_FullMethodName       = "/YrcfAgent.Tiering/FlushTiering"
	Tiering_ListTiering_FullMethodName        = "/YrcfAgent.Tiering/ListTiering"
	Tiering_StatTiering_FullMethodName        = "/YrcfAgent.Tiering/StatTiering"
	Tiering_RecoverTiering_FullMethodName     = "/YrcfAgent.Tiering/RecoverTiering"
	Tiering_RecoverTieringStat_FullMethodName = "/YrcfAgent.Tiering/RecoverTieringStat"
	Tiering_ScanTiering_FullMethodName        = "/YrcfAgent.Tiering/ScanTiering"
	Tiering_InfoTiering_FullMethodName        = "/YrcfAgent.Tiering/InfoTiering"
	Tiering_AbortTiering_FullMethodName       = "/YrcfAgent.Tiering/AbortTiering"
	Tiering_PauseTiering_FullMethodName       = "/YrcfAgent.Tiering/PauseTiering"
)

// TieringClient is the client API for Tiering service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TieringClient interface {
	AddTiering(ctx context.Context, in *AddTieringRequest, opts ...grpc.CallOption) (*AddTieringResponse, error)
	DeleteTiering(ctx context.Context, in *DeleteTieringRequest, opts ...grpc.CallOption) (*DeleteTieringResponse, error)
	UpdateTiering(ctx context.Context, in *UpdateTieringRequest, opts ...grpc.CallOption) (*UpdateTieringResponse, error)
	FlushTiering(ctx context.Context, in *FlushTieringRequest, opts ...grpc.CallOption) (*FlushTieringResponse, error)
	ListTiering(ctx context.Context, in *ListTieringRequest, opts ...grpc.CallOption) (*ListTieringResponse, error)
	StatTiering(ctx context.Context, in *StatTieringRequest, opts ...grpc.CallOption) (*StatTieringResponse, error)
	RecoverTiering(ctx context.Context, in *RecoverTieringRequest, opts ...grpc.CallOption) (*RecoverTieringResponse, error)
	RecoverTieringStat(ctx context.Context, in *RecoverTieringStatRequest, opts ...grpc.CallOption) (*RecoverTieringStatResponse, error)
	ScanTiering(ctx context.Context, in *ScanTieringRequest, opts ...grpc.CallOption) (*ScanTieringResponse, error)
	InfoTiering(ctx context.Context, in *InfoTieringRequest, opts ...grpc.CallOption) (*InfoTieringResponse, error)
	AbortTiering(ctx context.Context, in *AbortTieringRequest, opts ...grpc.CallOption) (*AbortTieringResponse, error)
	PauseTiering(ctx context.Context, in *PauseTieringRequest, opts ...grpc.CallOption) (*PauseTieringResponse, error)
}

type tieringClient struct {
	cc grpc.ClientConnInterface
}

func NewTieringClient(cc grpc.ClientConnInterface) TieringClient {
	return &tieringClient{cc}
}

func (c *tieringClient) AddTiering(ctx context.Context, in *AddTieringRequest, opts ...grpc.CallOption) (*AddTieringResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddTieringResponse)
	err := c.cc.Invoke(ctx, Tiering_AddTiering_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tieringClient) DeleteTiering(ctx context.Context, in *DeleteTieringRequest, opts ...grpc.CallOption) (*DeleteTieringResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteTieringResponse)
	err := c.cc.Invoke(ctx, Tiering_DeleteTiering_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tieringClient) UpdateTiering(ctx context.Context, in *UpdateTieringRequest, opts ...grpc.CallOption) (*UpdateTieringResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateTieringResponse)
	err := c.cc.Invoke(ctx, Tiering_UpdateTiering_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tieringClient) FlushTiering(ctx context.Context, in *FlushTieringRequest, opts ...grpc.CallOption) (*FlushTieringResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FlushTieringResponse)
	err := c.cc.Invoke(ctx, Tiering_FlushTiering_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tieringClient) ListTiering(ctx context.Context, in *ListTieringRequest, opts ...grpc.CallOption) (*ListTieringResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListTieringResponse)
	err := c.cc.Invoke(ctx, Tiering_ListTiering_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tieringClient) StatTiering(ctx context.Context, in *StatTieringRequest, opts ...grpc.CallOption) (*StatTieringResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StatTieringResponse)
	err := c.cc.Invoke(ctx, Tiering_StatTiering_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tieringClient) RecoverTiering(ctx context.Context, in *RecoverTieringRequest, opts ...grpc.CallOption) (*RecoverTieringResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RecoverTieringResponse)
	err := c.cc.Invoke(ctx, Tiering_RecoverTiering_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tieringClient) RecoverTieringStat(ctx context.Context, in *RecoverTieringStatRequest, opts ...grpc.CallOption) (*RecoverTieringStatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RecoverTieringStatResponse)
	err := c.cc.Invoke(ctx, Tiering_RecoverTieringStat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tieringClient) ScanTiering(ctx context.Context, in *ScanTieringRequest, opts ...grpc.CallOption) (*ScanTieringResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ScanTieringResponse)
	err := c.cc.Invoke(ctx, Tiering_ScanTiering_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tieringClient) InfoTiering(ctx context.Context, in *InfoTieringRequest, opts ...grpc.CallOption) (*InfoTieringResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InfoTieringResponse)
	err := c.cc.Invoke(ctx, Tiering_InfoTiering_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tieringClient) AbortTiering(ctx context.Context, in *AbortTieringRequest, opts ...grpc.CallOption) (*AbortTieringResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AbortTieringResponse)
	err := c.cc.Invoke(ctx, Tiering_AbortTiering_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tieringClient) PauseTiering(ctx context.Context, in *PauseTieringRequest, opts ...grpc.CallOption) (*PauseTieringResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PauseTieringResponse)
	err := c.cc.Invoke(ctx, Tiering_PauseTiering_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TieringServer is the server API for Tiering service.
// All implementations must embed UnimplementedTieringServer
// for forward compatibility
type TieringServer interface {
	AddTiering(context.Context, *AddTieringRequest) (*AddTieringResponse, error)
	DeleteTiering(context.Context, *DeleteTieringRequest) (*DeleteTieringResponse, error)
	UpdateTiering(context.Context, *UpdateTieringRequest) (*UpdateTieringResponse, error)
	FlushTiering(context.Context, *FlushTieringRequest) (*FlushTieringResponse, error)
	ListTiering(context.Context, *ListTieringRequest) (*ListTieringResponse, error)
	StatTiering(context.Context, *StatTieringRequest) (*StatTieringResponse, error)
	RecoverTiering(context.Context, *RecoverTieringRequest) (*RecoverTieringResponse, error)
	RecoverTieringStat(context.Context, *RecoverTieringStatRequest) (*RecoverTieringStatResponse, error)
	ScanTiering(context.Context, *ScanTieringRequest) (*ScanTieringResponse, error)
	InfoTiering(context.Context, *InfoTieringRequest) (*InfoTieringResponse, error)
	AbortTiering(context.Context, *AbortTieringRequest) (*AbortTieringResponse, error)
	PauseTiering(context.Context, *PauseTieringRequest) (*PauseTieringResponse, error)
	mustEmbedUnimplementedTieringServer()
}

// UnimplementedTieringServer must be embedded to have forward compatible implementations.
type UnimplementedTieringServer struct {
}

func (UnimplementedTieringServer) AddTiering(context.Context, *AddTieringRequest) (*AddTieringResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTiering not implemented")
}
func (UnimplementedTieringServer) DeleteTiering(context.Context, *DeleteTieringRequest) (*DeleteTieringResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTiering not implemented")
}
func (UnimplementedTieringServer) UpdateTiering(context.Context, *UpdateTieringRequest) (*UpdateTieringResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTiering not implemented")
}
func (UnimplementedTieringServer) FlushTiering(context.Context, *FlushTieringRequest) (*FlushTieringResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FlushTiering not implemented")
}
func (UnimplementedTieringServer) ListTiering(context.Context, *ListTieringRequest) (*ListTieringResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTiering not implemented")
}
func (UnimplementedTieringServer) StatTiering(context.Context, *StatTieringRequest) (*StatTieringResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StatTiering not implemented")
}
func (UnimplementedTieringServer) RecoverTiering(context.Context, *RecoverTieringRequest) (*RecoverTieringResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecoverTiering not implemented")
}
func (UnimplementedTieringServer) RecoverTieringStat(context.Context, *RecoverTieringStatRequest) (*RecoverTieringStatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecoverTieringStat not implemented")
}
func (UnimplementedTieringServer) ScanTiering(context.Context, *ScanTieringRequest) (*ScanTieringResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScanTiering not implemented")
}
func (UnimplementedTieringServer) InfoTiering(context.Context, *InfoTieringRequest) (*InfoTieringResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InfoTiering not implemented")
}
func (UnimplementedTieringServer) AbortTiering(context.Context, *AbortTieringRequest) (*AbortTieringResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AbortTiering not implemented")
}
func (UnimplementedTieringServer) PauseTiering(context.Context, *PauseTieringRequest) (*PauseTieringResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PauseTiering not implemented")
}
func (UnimplementedTieringServer) mustEmbedUnimplementedTieringServer() {}

// UnsafeTieringServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TieringServer will
// result in compilation errors.
type UnsafeTieringServer interface {
	mustEmbedUnimplementedTieringServer()
}

func RegisterTieringServer(s grpc.ServiceRegistrar, srv TieringServer) {
	s.RegisterService(&Tiering_ServiceDesc, srv)
}

func _Tiering_AddTiering_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTieringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TieringServer).AddTiering(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tiering_AddTiering_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TieringServer).AddTiering(ctx, req.(*AddTieringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tiering_DeleteTiering_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTieringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TieringServer).DeleteTiering(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tiering_DeleteTiering_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TieringServer).DeleteTiering(ctx, req.(*DeleteTieringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tiering_UpdateTiering_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTieringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TieringServer).UpdateTiering(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tiering_UpdateTiering_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TieringServer).UpdateTiering(ctx, req.(*UpdateTieringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tiering_FlushTiering_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlushTieringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TieringServer).FlushTiering(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tiering_FlushTiering_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TieringServer).FlushTiering(ctx, req.(*FlushTieringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tiering_ListTiering_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTieringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TieringServer).ListTiering(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tiering_ListTiering_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TieringServer).ListTiering(ctx, req.(*ListTieringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tiering_StatTiering_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatTieringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TieringServer).StatTiering(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tiering_StatTiering_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TieringServer).StatTiering(ctx, req.(*StatTieringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tiering_RecoverTiering_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecoverTieringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TieringServer).RecoverTiering(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tiering_RecoverTiering_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TieringServer).RecoverTiering(ctx, req.(*RecoverTieringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tiering_RecoverTieringStat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecoverTieringStatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TieringServer).RecoverTieringStat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tiering_RecoverTieringStat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TieringServer).RecoverTieringStat(ctx, req.(*RecoverTieringStatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tiering_ScanTiering_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScanTieringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TieringServer).ScanTiering(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tiering_ScanTiering_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TieringServer).ScanTiering(ctx, req.(*ScanTieringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tiering_InfoTiering_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoTieringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TieringServer).InfoTiering(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tiering_InfoTiering_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TieringServer).InfoTiering(ctx, req.(*InfoTieringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tiering_AbortTiering_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AbortTieringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TieringServer).AbortTiering(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tiering_AbortTiering_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TieringServer).AbortTiering(ctx, req.(*AbortTieringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tiering_PauseTiering_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PauseTieringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TieringServer).PauseTiering(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tiering_PauseTiering_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TieringServer).PauseTiering(ctx, req.(*PauseTieringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Tiering_ServiceDesc is the grpc.ServiceDesc for Tiering service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Tiering_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "YrcfAgent.Tiering",
	HandlerType: (*TieringServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTiering",
			Handler:    _Tiering_AddTiering_Handler,
		},
		{
			MethodName: "DeleteTiering",
			Handler:    _Tiering_DeleteTiering_Handler,
		},
		{
			MethodName: "UpdateTiering",
			Handler:    _Tiering_UpdateTiering_Handler,
		},
		{
			MethodName: "FlushTiering",
			Handler:    _Tiering_FlushTiering_Handler,
		},
		{
			MethodName: "ListTiering",
			Handler:    _Tiering_ListTiering_Handler,
		},
		{
			MethodName: "StatTiering",
			Handler:    _Tiering_StatTiering_Handler,
		},
		{
			MethodName: "RecoverTiering",
			Handler:    _Tiering_RecoverTiering_Handler,
		},
		{
			MethodName: "RecoverTieringStat",
			Handler:    _Tiering_RecoverTieringStat_Handler,
		},
		{
			MethodName: "ScanTiering",
			Handler:    _Tiering_ScanTiering_Handler,
		},
		{
			MethodName: "InfoTiering",
			Handler:    _Tiering_InfoTiering_Handler,
		},
		{
			MethodName: "AbortTiering",
			Handler:    _Tiering_AbortTiering_Handler,
		},
		{
			MethodName: "PauseTiering",
			Handler:    _Tiering_PauseTiering_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tiering/proto/tiering.proto",
}
