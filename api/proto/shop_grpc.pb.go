// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: api/proto/shop.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ShopService_ListBicycles_FullMethodName           = "/shop.ShopService/ListBicycles"
	ShopService_GetBicycle_FullMethodName             = "/shop.ShopService/GetBicycle"
	ShopService_SearchBicycles_FullMethodName         = "/shop.ShopService/SearchBicycles"
	ShopService_CreateOrder_FullMethodName            = "/shop.ShopService/CreateOrder"
	ShopService_GetOrder_FullMethodName               = "/shop.ShopService/GetOrder"
	ShopService_CancelOrder_FullMethodName            = "/shop.ShopService/CancelOrder"
	ShopService_ListOrders_FullMethodName             = "/shop.ShopService/ListOrders"
	ShopService_CreateRental_FullMethodName           = "/shop.ShopService/CreateRental"
	ShopService_GetRental_FullMethodName              = "/shop.ShopService/GetRental"
	ShopService_EndRental_FullMethodName              = "/shop.ShopService/EndRental"
	ShopService_ListRentals_FullMethodName            = "/shop.ShopService/ListRentals"
	ShopService_SendOrderConfirmation_FullMethodName  = "/shop.ShopService/SendOrderConfirmation"
	ShopService_SendRentalConfirmation_FullMethodName = "/shop.ShopService/SendRentalConfirmation"
)

// ShopServiceClient is the client API for ShopService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShopServiceClient interface {
	// Catalog
	ListBicycles(ctx context.Context, in *ListBicyclesRequest, opts ...grpc.CallOption) (*ListBicyclesResponse, error)
	GetBicycle(ctx context.Context, in *GetBicycleRequest, opts ...grpc.CallOption) (*Bicycle, error)
	SearchBicycles(ctx context.Context, in *SearchBicyclesRequest, opts ...grpc.CallOption) (*ListBicyclesResponse, error)
	// Orders
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*Order, error)
	GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*Order, error)
	CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*Empty, error)
	ListOrders(ctx context.Context, in *ListOrdersRequest, opts ...grpc.CallOption) (*ListOrdersResponse, error)
	// Rentals
	CreateRental(ctx context.Context, in *CreateRentalRequest, opts ...grpc.CallOption) (*Rental, error)
	GetRental(ctx context.Context, in *GetRentalRequest, opts ...grpc.CallOption) (*Rental, error)
	EndRental(ctx context.Context, in *EndRentalRequest, opts ...grpc.CallOption) (*Empty, error)
	ListRentals(ctx context.Context, in *ListRentalsRequest, opts ...grpc.CallOption) (*ListRentalsResponse, error)
	// Notifications
	SendOrderConfirmation(ctx context.Context, in *SendOrderConfirmationRequest, opts ...grpc.CallOption) (*Empty, error)
	SendRentalConfirmation(ctx context.Context, in *SendRentalConfirmationRequest, opts ...grpc.CallOption) (*Empty, error)
}

type shopServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShopServiceClient(cc grpc.ClientConnInterface) ShopServiceClient {
	return &shopServiceClient{cc}
}

func (c *shopServiceClient) ListBicycles(ctx context.Context, in *ListBicyclesRequest, opts ...grpc.CallOption) (*ListBicyclesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListBicyclesResponse)
	err := c.cc.Invoke(ctx, ShopService_ListBicycles_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) GetBicycle(ctx context.Context, in *GetBicycleRequest, opts ...grpc.CallOption) (*Bicycle, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Bicycle)
	err := c.cc.Invoke(ctx, ShopService_GetBicycle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) SearchBicycles(ctx context.Context, in *SearchBicyclesRequest, opts ...grpc.CallOption) (*ListBicyclesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListBicyclesResponse)
	err := c.cc.Invoke(ctx, ShopService_SearchBicycles_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*Order, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Order)
	err := c.cc.Invoke(ctx, ShopService_CreateOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*Order, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Order)
	err := c.cc.Invoke(ctx, ShopService_GetOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, ShopService_CancelOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) ListOrders(ctx context.Context, in *ListOrdersRequest, opts ...grpc.CallOption) (*ListOrdersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListOrdersResponse)
	err := c.cc.Invoke(ctx, ShopService_ListOrders_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) CreateRental(ctx context.Context, in *CreateRentalRequest, opts ...grpc.CallOption) (*Rental, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Rental)
	err := c.cc.Invoke(ctx, ShopService_CreateRental_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) GetRental(ctx context.Context, in *GetRentalRequest, opts ...grpc.CallOption) (*Rental, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Rental)
	err := c.cc.Invoke(ctx, ShopService_GetRental_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) EndRental(ctx context.Context, in *EndRentalRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, ShopService_EndRental_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) ListRentals(ctx context.Context, in *ListRentalsRequest, opts ...grpc.CallOption) (*ListRentalsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListRentalsResponse)
	err := c.cc.Invoke(ctx, ShopService_ListRentals_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) SendOrderConfirmation(ctx context.Context, in *SendOrderConfirmationRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, ShopService_SendOrderConfirmation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) SendRentalConfirmation(ctx context.Context, in *SendRentalConfirmationRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, ShopService_SendRentalConfirmation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShopServiceServer is the server API for ShopService service.
// All implementations must embed UnimplementedShopServiceServer
// for forward compatibility.
type ShopServiceServer interface {
	// Catalog
	ListBicycles(context.Context, *ListBicyclesRequest) (*ListBicyclesResponse, error)
	GetBicycle(context.Context, *GetBicycleRequest) (*Bicycle, error)
	SearchBicycles(context.Context, *SearchBicyclesRequest) (*ListBicyclesResponse, error)
	// Orders
	CreateOrder(context.Context, *CreateOrderRequest) (*Order, error)
	GetOrder(context.Context, *GetOrderRequest) (*Order, error)
	CancelOrder(context.Context, *CancelOrderRequest) (*Empty, error)
	ListOrders(context.Context, *ListOrdersRequest) (*ListOrdersResponse, error)
	// Rentals
	CreateRental(context.Context, *CreateRentalRequest) (*Rental, error)
	GetRental(context.Context, *GetRentalRequest) (*Rental, error)
	EndRental(context.Context, *EndRentalRequest) (*Empty, error)
	ListRentals(context.Context, *ListRentalsRequest) (*ListRentalsResponse, error)
	// Notifications
	SendOrderConfirmation(context.Context, *SendOrderConfirmationRequest) (*Empty, error)
	SendRentalConfirmation(context.Context, *SendRentalConfirmationRequest) (*Empty, error)
	mustEmbedUnimplementedShopServiceServer()
}

// UnimplementedShopServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedShopServiceServer struct{}

func (UnimplementedShopServiceServer) ListBicycles(context.Context, *ListBicyclesRequest) (*ListBicyclesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBicycles not implemented")
}
func (UnimplementedShopServiceServer) GetBicycle(context.Context, *GetBicycleRequest) (*Bicycle, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBicycle not implemented")
}
func (UnimplementedShopServiceServer) SearchBicycles(context.Context, *SearchBicyclesRequest) (*ListBicyclesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchBicycles not implemented")
}
func (UnimplementedShopServiceServer) CreateOrder(context.Context, *CreateOrderRequest) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedShopServiceServer) GetOrder(context.Context, *GetOrderRequest) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedShopServiceServer) CancelOrder(context.Context, *CancelOrderRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOrder not implemented")
}
func (UnimplementedShopServiceServer) ListOrders(context.Context, *ListOrdersRequest) (*ListOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrders not implemented")
}
func (UnimplementedShopServiceServer) CreateRental(context.Context, *CreateRentalRequest) (*Rental, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRental not implemented")
}
func (UnimplementedShopServiceServer) GetRental(context.Context, *GetRentalRequest) (*Rental, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRental not implemented")
}
func (UnimplementedShopServiceServer) EndRental(context.Context, *EndRentalRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EndRental not implemented")
}
func (UnimplementedShopServiceServer) ListRentals(context.Context, *ListRentalsRequest) (*ListRentalsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRentals not implemented")
}
func (UnimplementedShopServiceServer) SendOrderConfirmation(context.Context, *SendOrderConfirmationRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendOrderConfirmation not implemented")
}
func (UnimplementedShopServiceServer) SendRentalConfirmation(context.Context, *SendRentalConfirmationRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendRentalConfirmation not implemented")
}
func (UnimplementedShopServiceServer) mustEmbedUnimplementedShopServiceServer() {}
func (UnimplementedShopServiceServer) testEmbeddedByValue()                     {}

// UnsafeShopServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShopServiceServer will
// result in compilation errors.
type UnsafeShopServiceServer interface {
	mustEmbedUnimplementedShopServiceServer()
}

func RegisterShopServiceServer(s grpc.ServiceRegistrar, srv ShopServiceServer) {
	// If the following call pancis, it indicates UnimplementedShopServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ShopService_ServiceDesc, srv)
}

func _ShopService_ListBicycles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBicyclesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).ListBicycles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShopService_ListBicycles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).ListBicycles(ctx, req.(*ListBicyclesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_GetBicycle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBicycleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).GetBicycle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShopService_GetBicycle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).GetBicycle(ctx, req.(*GetBicycleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_SearchBicycles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchBicyclesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).SearchBicycles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShopService_SearchBicycles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).SearchBicycles(ctx, req.(*SearchBicyclesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShopService_CreateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShopService_GetOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).GetOrder(ctx, req.(*GetOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShopService_CancelOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).CancelOrder(ctx, req.(*CancelOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_ListOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).ListOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShopService_ListOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).ListOrders(ctx, req.(*ListOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_CreateRental_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRentalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).CreateRental(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShopService_CreateRental_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).CreateRental(ctx, req.(*CreateRentalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_GetRental_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRentalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).GetRental(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShopService_GetRental_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).GetRental(ctx, req.(*GetRentalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_EndRental_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndRentalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).EndRental(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShopService_EndRental_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).EndRental(ctx, req.(*EndRentalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_ListRentals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRentalsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).ListRentals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShopService_ListRentals_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).ListRentals(ctx, req.(*ListRentalsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_SendOrderConfirmation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendOrderConfirmationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).SendOrderConfirmation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShopService_SendOrderConfirmation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).SendOrderConfirmation(ctx, req.(*SendOrderConfirmationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_SendRentalConfirmation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRentalConfirmationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).SendRentalConfirmation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShopService_SendRentalConfirmation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).SendRentalConfirmation(ctx, req.(*SendRentalConfirmationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShopService_ServiceDesc is the grpc.ServiceDesc for ShopService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShopService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shop.ShopService",
	HandlerType: (*ShopServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListBicycles",
			Handler:    _ShopService_ListBicycles_Handler,
		},
		{
			MethodName: "GetBicycle",
			Handler:    _ShopService_GetBicycle_Handler,
		},
		{
			MethodName: "SearchBicycles",
			Handler:    _ShopService_SearchBicycles_Handler,
		},
		{
			MethodName: "CreateOrder",
			Handler:    _ShopService_CreateOrder_Handler,
		},
		{
			MethodName: "GetOrder",
			Handler:    _ShopService_GetOrder_Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _ShopService_CancelOrder_Handler,
		},
		{
			MethodName: "ListOrders",
			Handler:    _ShopService_ListOrders_Handler,
		},
		{
			MethodName: "CreateRental",
			Handler:    _ShopService_CreateRental_Handler,
		},
		{
			MethodName: "GetRental",
			Handler:    _ShopService_GetRental_Handler,
		},
		{
			MethodName: "EndRental",
			Handler:    _ShopService_EndRental_Handler,
		},
		{
			MethodName: "ListRentals",
			Handler:    _ShopService_ListRentals_Handler,
		},
		{
			MethodName: "SendOrderConfirmation",
			Handler:    _ShopService_SendOrderConfirmation_Handler,
		},
		{
			MethodName: "SendRentalConfirmation",
			Handler:    _ShopService_SendRentalConfirmation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/shop.proto",
}
