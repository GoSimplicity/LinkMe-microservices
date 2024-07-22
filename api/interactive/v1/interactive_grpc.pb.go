// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: api/interactive/v1/interactive.proto

package v1

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
	Interactive_GetInteractive_FullMethodName  = "/api.interactive.v1.Interactive/GetInteractive"
	Interactive_ListInteractive_FullMethodName = "/api.interactive.v1.Interactive/ListInteractive"
	Interactive_AddReadCount_FullMethodName    = "/api.interactive.v1.Interactive/AddReadCount"
	Interactive_AddLikeCount_FullMethodName    = "/api.interactive.v1.Interactive/AddLikeCount"
	Interactive_AddCollectCount_FullMethodName = "/api.interactive.v1.Interactive/AddCollectCount"
)

// InteractiveClient is the client API for Interactive service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InteractiveClient interface {
	GetInteractive(ctx context.Context, in *GetInteractiveRequest, opts ...grpc.CallOption) (*GetInteractiveReply, error)
	ListInteractive(ctx context.Context, in *ListInteractiveRequest, opts ...grpc.CallOption) (*ListInteractiveReply, error)
	AddReadCount(ctx context.Context, in *AddCountRequest, opts ...grpc.CallOption) (*AddCountReply, error)
	AddLikeCount(ctx context.Context, in *AddCountRequest, opts ...grpc.CallOption) (*AddCountReply, error)
	AddCollectCount(ctx context.Context, in *AddCountRequest, opts ...grpc.CallOption) (*AddCountReply, error)
}

type interactiveClient struct {
	cc grpc.ClientConnInterface
}

func NewInteractiveClient(cc grpc.ClientConnInterface) InteractiveClient {
	return &interactiveClient{cc}
}

func (c *interactiveClient) GetInteractive(ctx context.Context, in *GetInteractiveRequest, opts ...grpc.CallOption) (*GetInteractiveReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetInteractiveReply)
	err := c.cc.Invoke(ctx, Interactive_GetInteractive_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactiveClient) ListInteractive(ctx context.Context, in *ListInteractiveRequest, opts ...grpc.CallOption) (*ListInteractiveReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListInteractiveReply)
	err := c.cc.Invoke(ctx, Interactive_ListInteractive_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactiveClient) AddReadCount(ctx context.Context, in *AddCountRequest, opts ...grpc.CallOption) (*AddCountReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddCountReply)
	err := c.cc.Invoke(ctx, Interactive_AddReadCount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactiveClient) AddLikeCount(ctx context.Context, in *AddCountRequest, opts ...grpc.CallOption) (*AddCountReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddCountReply)
	err := c.cc.Invoke(ctx, Interactive_AddLikeCount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactiveClient) AddCollectCount(ctx context.Context, in *AddCountRequest, opts ...grpc.CallOption) (*AddCountReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddCountReply)
	err := c.cc.Invoke(ctx, Interactive_AddCollectCount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InteractiveServer is the server API for Interactive service.
// All implementations must embed UnimplementedInteractiveServer
// for forward compatibility
type InteractiveServer interface {
	GetInteractive(context.Context, *GetInteractiveRequest) (*GetInteractiveReply, error)
	ListInteractive(context.Context, *ListInteractiveRequest) (*ListInteractiveReply, error)
	AddReadCount(context.Context, *AddCountRequest) (*AddCountReply, error)
	AddLikeCount(context.Context, *AddCountRequest) (*AddCountReply, error)
	AddCollectCount(context.Context, *AddCountRequest) (*AddCountReply, error)
	mustEmbedUnimplementedInteractiveServer()
}

// UnimplementedInteractiveServer must be embedded to have forward compatible implementations.
type UnimplementedInteractiveServer struct {
}

func (UnimplementedInteractiveServer) GetInteractive(context.Context, *GetInteractiveRequest) (*GetInteractiveReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInteractive not implemented")
}
func (UnimplementedInteractiveServer) ListInteractive(context.Context, *ListInteractiveRequest) (*ListInteractiveReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListInteractive not implemented")
}
func (UnimplementedInteractiveServer) AddReadCount(context.Context, *AddCountRequest) (*AddCountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddReadCount not implemented")
}
func (UnimplementedInteractiveServer) AddLikeCount(context.Context, *AddCountRequest) (*AddCountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddLikeCount not implemented")
}
func (UnimplementedInteractiveServer) AddCollectCount(context.Context, *AddCountRequest) (*AddCountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCollectCount not implemented")
}
func (UnimplementedInteractiveServer) mustEmbedUnimplementedInteractiveServer() {}

// UnsafeInteractiveServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InteractiveServer will
// result in compilation errors.
type UnsafeInteractiveServer interface {
	mustEmbedUnimplementedInteractiveServer()
}

func RegisterInteractiveServer(s grpc.ServiceRegistrar, srv InteractiveServer) {
	s.RegisterService(&Interactive_ServiceDesc, srv)
}

func _Interactive_GetInteractive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInteractiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractiveServer).GetInteractive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Interactive_GetInteractive_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractiveServer).GetInteractive(ctx, req.(*GetInteractiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Interactive_ListInteractive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInteractiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractiveServer).ListInteractive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Interactive_ListInteractive_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractiveServer).ListInteractive(ctx, req.(*ListInteractiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Interactive_AddReadCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractiveServer).AddReadCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Interactive_AddReadCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractiveServer).AddReadCount(ctx, req.(*AddCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Interactive_AddLikeCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractiveServer).AddLikeCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Interactive_AddLikeCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractiveServer).AddLikeCount(ctx, req.(*AddCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Interactive_AddCollectCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractiveServer).AddCollectCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Interactive_AddCollectCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractiveServer).AddCollectCount(ctx, req.(*AddCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Interactive_ServiceDesc is the grpc.ServiceDesc for Interactive service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Interactive_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.interactive.v1.Interactive",
	HandlerType: (*InteractiveServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInteractive",
			Handler:    _Interactive_GetInteractive_Handler,
		},
		{
			MethodName: "ListInteractive",
			Handler:    _Interactive_ListInteractive_Handler,
		},
		{
			MethodName: "AddReadCount",
			Handler:    _Interactive_AddReadCount_Handler,
		},
		{
			MethodName: "AddLikeCount",
			Handler:    _Interactive_AddLikeCount_Handler,
		},
		{
			MethodName: "AddCollectCount",
			Handler:    _Interactive_AddCollectCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/interactive/v1/interactive.proto",
}