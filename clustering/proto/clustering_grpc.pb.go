// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: clustering/proto/clustering.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ClusteringSericeClient is the client API for ClusteringSerice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClusteringSericeClient interface {
	Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error)
	Info(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ClusterInfo, error)
	PingIndirect(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type clusteringSericeClient struct {
	cc grpc.ClientConnInterface
}

func NewClusteringSericeClient(cc grpc.ClientConnInterface) ClusteringSericeClient {
	return &clusteringSericeClient{cc}
}

func (c *clusteringSericeClient) Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error) {
	out := new(JoinResponse)
	err := c.cc.Invoke(ctx, "/clustering.ClusteringSerice/Join", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusteringSericeClient) Info(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ClusterInfo, error) {
	out := new(ClusterInfo)
	err := c.cc.Invoke(ctx, "/clustering.ClusteringSerice/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusteringSericeClient) PingIndirect(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/clustering.ClusteringSerice/PingIndirect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusteringSericeClient) Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/clustering.ClusteringSerice/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusteringSericeServer is the server API for ClusteringSerice service.
// All implementations must embed UnimplementedClusteringSericeServer
// for forward compatibility
type ClusteringSericeServer interface {
	Join(context.Context, *JoinRequest) (*JoinResponse, error)
	Info(context.Context, *emptypb.Empty) (*ClusterInfo, error)
	PingIndirect(context.Context, *PingRequest) (*PingResponse, error)
	Ping(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	mustEmbedUnimplementedClusteringSericeServer()
}

// UnimplementedClusteringSericeServer must be embedded to have forward compatible implementations.
type UnimplementedClusteringSericeServer struct {
}

func (UnimplementedClusteringSericeServer) Join(context.Context, *JoinRequest) (*JoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Join not implemented")
}
func (UnimplementedClusteringSericeServer) Info(context.Context, *emptypb.Empty) (*ClusterInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (UnimplementedClusteringSericeServer) PingIndirect(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingIndirect not implemented")
}
func (UnimplementedClusteringSericeServer) Ping(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedClusteringSericeServer) mustEmbedUnimplementedClusteringSericeServer() {}

// UnsafeClusteringSericeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClusteringSericeServer will
// result in compilation errors.
type UnsafeClusteringSericeServer interface {
	mustEmbedUnimplementedClusteringSericeServer()
}

func RegisterClusteringSericeServer(s grpc.ServiceRegistrar, srv ClusteringSericeServer) {
	s.RegisterService(&ClusteringSerice_ServiceDesc, srv)
}

func _ClusteringSerice_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusteringSericeServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clustering.ClusteringSerice/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusteringSericeServer).Join(ctx, req.(*JoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusteringSerice_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusteringSericeServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clustering.ClusteringSerice/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusteringSericeServer).Info(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusteringSerice_PingIndirect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusteringSericeServer).PingIndirect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clustering.ClusteringSerice/PingIndirect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusteringSericeServer).PingIndirect(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusteringSerice_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusteringSericeServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clustering.ClusteringSerice/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusteringSericeServer).Ping(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ClusteringSerice_ServiceDesc is the grpc.ServiceDesc for ClusteringSerice service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClusteringSerice_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "clustering.ClusteringSerice",
	HandlerType: (*ClusteringSericeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Join",
			Handler:    _ClusteringSerice_Join_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _ClusteringSerice_Info_Handler,
		},
		{
			MethodName: "PingIndirect",
			Handler:    _ClusteringSerice_PingIndirect_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _ClusteringSerice_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "clustering/proto/clustering.proto",
}
