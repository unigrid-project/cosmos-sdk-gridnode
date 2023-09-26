// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cosmossdkgridnode/gridnode/gridnode.proto

package types

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	GridnodeQuery_DelegatedAmount_FullMethodName  = "/cosmossdkgridnode.gridnode.GridnodeQuery/DelegatedAmount"
	GridnodeQuery_UndelegateTokens_FullMethodName = "/cosmossdkgridnode.gridnode.GridnodeQuery/UndelegateTokens"
)

// GridnodeQueryClient is the client API for GridnodeQuery service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GridnodeQueryClient interface {
	// DelegatedAmount queries the amount delegated by a specific account.
	DelegatedAmount(ctx context.Context, in *QueryDelegatedAmountRequest, opts ...grpc.CallOption) (*QueryDelegatedAmountResponse, error)
	UndelegateTokens(ctx context.Context, in *MsgGridnodeUndelegate, opts ...grpc.CallOption) (*MsgGridnodeUndelegateResponse, error)
}

type gridnodeQueryClient struct {
	cc grpc.ClientConnInterface
}

func NewGridnodeQueryClient(cc grpc.ClientConnInterface) GridnodeQueryClient {
	return &gridnodeQueryClient{cc}
}

func (c *gridnodeQueryClient) DelegatedAmount(ctx context.Context, in *QueryDelegatedAmountRequest, opts ...grpc.CallOption) (*QueryDelegatedAmountResponse, error) {
	out := new(QueryDelegatedAmountResponse)
	err := c.cc.Invoke(ctx, GridnodeQuery_DelegatedAmount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gridnodeQueryClient) UndelegateTokens(ctx context.Context, in *MsgGridnodeUndelegate, opts ...grpc.CallOption) (*MsgGridnodeUndelegateResponse, error) {
	out := new(MsgGridnodeUndelegateResponse)
	err := c.cc.Invoke(ctx, GridnodeQuery_UndelegateTokens_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GridnodeQueryServer is the server API for GridnodeQuery service.
// All implementations should embed UnimplementedGridnodeQueryServer
// for forward compatibility
type GridnodeQueryServer interface {
	// DelegatedAmount queries the amount delegated by a specific account.
	DelegatedAmount(context.Context, *QueryDelegatedAmountRequest) (*QueryDelegatedAmountResponse, error)
	UndelegateTokens(context.Context, *MsgGridnodeUndelegate) (*MsgGridnodeUndelegateResponse, error)
}

// UnimplementedGridnodeQueryServer should be embedded to have forward compatible implementations.
type UnimplementedGridnodeQueryServer struct {
}

func (UnimplementedGridnodeQueryServer) DelegatedAmount(context.Context, *QueryDelegatedAmountRequest) (*QueryDelegatedAmountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelegatedAmount not implemented")
}
func (UnimplementedGridnodeQueryServer) UndelegateTokens(context.Context, *MsgGridnodeUndelegate) (*MsgGridnodeUndelegateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UndelegateTokens not implemented")
}

// UnsafeGridnodeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GridnodeQueryServer will
// result in compilation errors.
type UnsafeGridnodeQueryServer interface {
	mustEmbedUnimplementedGridnodeQueryServer()
}

func RegisterGridnodeQueryServer(s grpc.ServiceRegistrar, srv GridnodeQueryServer) {
	s.RegisterService(&GridnodeQuery_ServiceDesc, srv)
}

func _GridnodeQuery_DelegatedAmount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDelegatedAmountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GridnodeQueryServer).DelegatedAmount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GridnodeQuery_DelegatedAmount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GridnodeQueryServer).DelegatedAmount(ctx, req.(*QueryDelegatedAmountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GridnodeQuery_UndelegateTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgGridnodeUndelegate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GridnodeQueryServer).UndelegateTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GridnodeQuery_UndelegateTokens_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GridnodeQueryServer).UndelegateTokens(ctx, req.(*MsgGridnodeUndelegate))
	}
	return interceptor(ctx, in, info, handler)
}

// GridnodeQuery_ServiceDesc is the grpc.ServiceDesc for GridnodeQuery service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GridnodeQuery_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmossdkgridnode.gridnode.GridnodeQuery",
	HandlerType: (*GridnodeQueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DelegatedAmount",
			Handler:    _GridnodeQuery_DelegatedAmount_Handler,
		},
		{
			MethodName: "UndelegateTokens",
			Handler:    _GridnodeQuery_UndelegateTokens_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmossdkgridnode/gridnode/gridnode.proto",
}

const (
	GridnodeMsg_DelegateTokens_FullMethodName = "/cosmossdkgridnode.gridnode.GridnodeMsg/DelegateTokens"
)

// GridnodeMsgClient is the client API for GridnodeMsg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GridnodeMsgClient interface {
	DelegateTokens(ctx context.Context, in *MsgGridnodeDelegate, opts ...grpc.CallOption) (*MsgGridnodeDelegateResponse, error)
}

type gridnodeMsgClient struct {
	cc grpc.ClientConnInterface
}

func NewGridnodeMsgClient(cc grpc.ClientConnInterface) GridnodeMsgClient {
	return &gridnodeMsgClient{cc}
}

func (c *gridnodeMsgClient) DelegateTokens(ctx context.Context, in *MsgGridnodeDelegate, opts ...grpc.CallOption) (*MsgGridnodeDelegateResponse, error) {
	out := new(MsgGridnodeDelegateResponse)
	err := c.cc.Invoke(ctx, GridnodeMsg_DelegateTokens_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GridnodeMsgServer is the server API for GridnodeMsg service.
// All implementations should embed UnimplementedGridnodeMsgServer
// for forward compatibility
type GridnodeMsgServer interface {
	DelegateTokens(context.Context, *MsgGridnodeDelegate) (*MsgGridnodeDelegateResponse, error)
}

// UnimplementedGridnodeMsgServer should be embedded to have forward compatible implementations.
type UnimplementedGridnodeMsgServer struct {
}

func (UnimplementedGridnodeMsgServer) DelegateTokens(context.Context, *MsgGridnodeDelegate) (*MsgGridnodeDelegateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelegateTokens not implemented")
}

// UnsafeGridnodeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GridnodeMsgServer will
// result in compilation errors.
type UnsafeGridnodeMsgServer interface {
	mustEmbedUnimplementedGridnodeMsgServer()
}

func RegisterGridnodeMsgServer(s grpc.ServiceRegistrar, srv GridnodeMsgServer) {
	s.RegisterService(&GridnodeMsg_ServiceDesc, srv)
}

func _GridnodeMsg_DelegateTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgGridnodeDelegate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GridnodeMsgServer).DelegateTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GridnodeMsg_DelegateTokens_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GridnodeMsgServer).DelegateTokens(ctx, req.(*MsgGridnodeDelegate))
	}
	return interceptor(ctx, in, info, handler)
}

// GridnodeMsg_ServiceDesc is the grpc.ServiceDesc for GridnodeMsg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GridnodeMsg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmossdkgridnode.gridnode.GridnodeMsg",
	HandlerType: (*GridnodeMsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DelegateTokens",
			Handler:    _GridnodeMsg_DelegateTokens_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmossdkgridnode/gridnode/gridnode.proto",
}
