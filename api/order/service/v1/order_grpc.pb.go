// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: order/service/v1/order.proto

package v1

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

// OrderClient is the client API for Order service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderClient interface {
	ListOrder(ctx context.Context, in *ListOrderReq, opts ...grpc.CallOption) (*ListOrderReply, error)
	CreateOrder(ctx context.Context, in *CreateOrderReq, opts ...grpc.CallOption) (*CreateOrderReply, error)
	GetOrder(ctx context.Context, in *GetOrderReq, opts ...grpc.CallOption) (*GetOrderReply, error)
	UpdateOrder(ctx context.Context, in *UpdateOrderReq, opts ...grpc.CallOption) (*UpdateOrderReply, error)
}

type orderClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderClient(cc grpc.ClientConnInterface) OrderClient {
	return &orderClient{cc}
}

func (c *orderClient) ListOrder(ctx context.Context, in *ListOrderReq, opts ...grpc.CallOption) (*ListOrderReply, error) {
	out := new(ListOrderReply)
	err := c.cc.Invoke(ctx, "/api.order.service.v1.Order/ListOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) CreateOrder(ctx context.Context, in *CreateOrderReq, opts ...grpc.CallOption) (*CreateOrderReply, error) {
	out := new(CreateOrderReply)
	err := c.cc.Invoke(ctx, "/api.order.service.v1.Order/CreateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) GetOrder(ctx context.Context, in *GetOrderReq, opts ...grpc.CallOption) (*GetOrderReply, error) {
	out := new(GetOrderReply)
	err := c.cc.Invoke(ctx, "/api.order.service.v1.Order/GetOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) UpdateOrder(ctx context.Context, in *UpdateOrderReq, opts ...grpc.CallOption) (*UpdateOrderReply, error) {
	out := new(UpdateOrderReply)
	err := c.cc.Invoke(ctx, "/api.order.service.v1.Order/UpdateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServer is the server API for Order service.
// All implementations must embed UnimplementedOrderServer
// for forward compatibility
type OrderServer interface {
	ListOrder(context.Context, *ListOrderReq) (*ListOrderReply, error)
	CreateOrder(context.Context, *CreateOrderReq) (*CreateOrderReply, error)
	GetOrder(context.Context, *GetOrderReq) (*GetOrderReply, error)
	UpdateOrder(context.Context, *UpdateOrderReq) (*UpdateOrderReply, error)
	mustEmbedUnimplementedOrderServer()
}

// UnimplementedOrderServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServer struct {
}

func (UnimplementedOrderServer) ListOrder(context.Context, *ListOrderReq) (*ListOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrder not implemented")
}
func (UnimplementedOrderServer) CreateOrder(context.Context, *CreateOrderReq) (*CreateOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedOrderServer) GetOrder(context.Context, *GetOrderReq) (*GetOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedOrderServer) UpdateOrder(context.Context, *UpdateOrderReq) (*UpdateOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrder not implemented")
}
func (UnimplementedOrderServer) mustEmbedUnimplementedOrderServer() {}

// UnsafeOrderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServer will
// result in compilation errors.
type UnsafeOrderServer interface {
	mustEmbedUnimplementedOrderServer()
}

func RegisterOrderServer(s grpc.ServiceRegistrar, srv OrderServer) {
	s.RegisterService(&Order_ServiceDesc, srv)
}

func _Order_ListOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).ListOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.order.service.v1.Order/ListOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).ListOrder(ctx, req.(*ListOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.order.service.v1.Order/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).CreateOrder(ctx, req.(*CreateOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.order.service.v1.Order/GetOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).GetOrder(ctx, req.(*GetOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_UpdateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).UpdateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.order.service.v1.Order/UpdateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).UpdateOrder(ctx, req.(*UpdateOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Order_ServiceDesc is the grpc.ServiceDesc for Order service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Order_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.order.service.v1.Order",
	HandlerType: (*OrderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListOrder",
			Handler:    _Order_ListOrder_Handler,
		},
		{
			MethodName: "CreateOrder",
			Handler:    _Order_CreateOrder_Handler,
		},
		{
			MethodName: "GetOrder",
			Handler:    _Order_GetOrder_Handler,
		},
		{
			MethodName: "UpdateOrder",
			Handler:    _Order_UpdateOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order/service/v1/order.proto",
}
