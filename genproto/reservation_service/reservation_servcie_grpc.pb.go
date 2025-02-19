// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.12
// source: reservation/reservation_servcie.proto

package reservation_service

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
	ReservationService_AddRestaurant_FullMethodName          = "/res_service.reservation_service/add_restaurant"
	ReservationService_GetRestaurant_FullMethodName          = "/res_service.reservation_service/get_restaurant"
	ReservationService_UpdateRestaurant_FullMethodName       = "/res_service.reservation_service/update_restaurant"
	ReservationService_DeleteRestaurant_FullMethodName       = "/res_service.reservation_service/delete_restaurant"
	ReservationService_GetRestaurants_FullMethodName         = "/res_service.reservation_service/get_restaurants"
	ReservationService_AddReservation_FullMethodName         = "/res_service.reservation_service/add_reservation"
	ReservationService_GetReservation_FullMethodName         = "/res_service.reservation_service/get_reservation"
	ReservationService_UpdateReservation_FullMethodName      = "/res_service.reservation_service/update_reservation"
	ReservationService_DeleteReservation_FullMethodName      = "/res_service.reservation_service/delete_reservation"
	ReservationService_GetReservations_FullMethodName        = "/res_service.reservation_service/get_reservations"
	ReservationService_AddReservationOrder_FullMethodName    = "/res_service.reservation_service/add_reservation_order"
	ReservationService_GetReservationOrder_FullMethodName    = "/res_service.reservation_service/get_reservation_order"
	ReservationService_UpdateReservationOrder_FullMethodName = "/res_service.reservation_service/update_reservation_order"
	ReservationService_DeleteReservationOrder_FullMethodName = "/res_service.reservation_service/delete_reservation_order"
	ReservationService_GetReservationOrders_FullMethodName   = "/res_service.reservation_service/get_reservation_orders"
	ReservationService_AddMenu_FullMethodName                = "/res_service.reservation_service/add_menu"
	ReservationService_GetMenu_FullMethodName                = "/res_service.reservation_service/get_menu"
	ReservationService_UpdateMenu_FullMethodName             = "/res_service.reservation_service/update_menu"
	ReservationService_DeleteMenu_FullMethodName             = "/res_service.reservation_service/delete_menu"
	ReservationService_GetMenus_FullMethodName               = "/res_service.reservation_service/get_menus"
)

// ReservationServiceClient is the client API for ReservationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReservationServiceClient interface {
	// Restaurant methods
	AddRestaurant(ctx context.Context, in *AddRestaurantRequest, opts ...grpc.CallOption) (*AddRestaurantResponse, error)
	GetRestaurant(ctx context.Context, in *GetRestaurantRequest, opts ...grpc.CallOption) (*GetRestaurantResponse, error)
	UpdateRestaurant(ctx context.Context, in *UpdateRestaurantRequest, opts ...grpc.CallOption) (*UpdateRestaurantResponse, error)
	DeleteRestaurant(ctx context.Context, in *DeleteRestaurantRequest, opts ...grpc.CallOption) (*DeleteRestaurantResponse, error)
	GetRestaurants(ctx context.Context, in *GetRestaurantsRequest, opts ...grpc.CallOption) (*GetRestaurantsResponse, error)
	// Reservation methods
	AddReservation(ctx context.Context, in *AddReservationRequest, opts ...grpc.CallOption) (*AddReservationResponse, error)
	GetReservation(ctx context.Context, in *GetReservationRequest, opts ...grpc.CallOption) (*GetReservationResponse, error)
	UpdateReservation(ctx context.Context, in *UpdateReservationRequest, opts ...grpc.CallOption) (*UpdateReservationResponse, error)
	DeleteReservation(ctx context.Context, in *DeleteReservationRequest, opts ...grpc.CallOption) (*DeleteReservationResponse, error)
	GetReservations(ctx context.Context, in *GetReservationsRequest, opts ...grpc.CallOption) (*GetReservationsResponse, error)
	// Reservation Order methods
	AddReservationOrder(ctx context.Context, in *AddReservationOrderRequest, opts ...grpc.CallOption) (*AddReservationOrderResponse, error)
	GetReservationOrder(ctx context.Context, in *GetReservationOrderRequest, opts ...grpc.CallOption) (*GetReservationOrderResponse, error)
	UpdateReservationOrder(ctx context.Context, in *UpdateReservationOrderRequest, opts ...grpc.CallOption) (*UpdateReservationOrderResponse, error)
	DeleteReservationOrder(ctx context.Context, in *DeleteReservationOrderRequest, opts ...grpc.CallOption) (*DeleteReservationOrderResponse, error)
	GetReservationOrders(ctx context.Context, in *GetReservationOrdersRequest, opts ...grpc.CallOption) (*GetReservationOrdersResponse, error)
	// Menu methods
	AddMenu(ctx context.Context, in *AddMenuRequest, opts ...grpc.CallOption) (*AddMenuResponse, error)
	GetMenu(ctx context.Context, in *GetMenuRequest, opts ...grpc.CallOption) (*GetMenuResponse, error)
	UpdateMenu(ctx context.Context, in *UpdateMenuRequest, opts ...grpc.CallOption) (*UpdateMenuResponse, error)
	DeleteMenu(ctx context.Context, in *DeleteMenuRequest, opts ...grpc.CallOption) (*DeleteMenuResponse, error)
	GetMenus(ctx context.Context, in *GetMenusRequest, opts ...grpc.CallOption) (*GetMenusResponse, error)
}

type reservationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReservationServiceClient(cc grpc.ClientConnInterface) ReservationServiceClient {
	return &reservationServiceClient{cc}
}

func (c *reservationServiceClient) AddRestaurant(ctx context.Context, in *AddRestaurantRequest, opts ...grpc.CallOption) (*AddRestaurantResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddRestaurantResponse)
	err := c.cc.Invoke(ctx, ReservationService_AddRestaurant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetRestaurant(ctx context.Context, in *GetRestaurantRequest, opts ...grpc.CallOption) (*GetRestaurantResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRestaurantResponse)
	err := c.cc.Invoke(ctx, ReservationService_GetRestaurant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) UpdateRestaurant(ctx context.Context, in *UpdateRestaurantRequest, opts ...grpc.CallOption) (*UpdateRestaurantResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateRestaurantResponse)
	err := c.cc.Invoke(ctx, ReservationService_UpdateRestaurant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) DeleteRestaurant(ctx context.Context, in *DeleteRestaurantRequest, opts ...grpc.CallOption) (*DeleteRestaurantResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteRestaurantResponse)
	err := c.cc.Invoke(ctx, ReservationService_DeleteRestaurant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetRestaurants(ctx context.Context, in *GetRestaurantsRequest, opts ...grpc.CallOption) (*GetRestaurantsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRestaurantsResponse)
	err := c.cc.Invoke(ctx, ReservationService_GetRestaurants_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) AddReservation(ctx context.Context, in *AddReservationRequest, opts ...grpc.CallOption) (*AddReservationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddReservationResponse)
	err := c.cc.Invoke(ctx, ReservationService_AddReservation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetReservation(ctx context.Context, in *GetReservationRequest, opts ...grpc.CallOption) (*GetReservationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetReservationResponse)
	err := c.cc.Invoke(ctx, ReservationService_GetReservation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) UpdateReservation(ctx context.Context, in *UpdateReservationRequest, opts ...grpc.CallOption) (*UpdateReservationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateReservationResponse)
	err := c.cc.Invoke(ctx, ReservationService_UpdateReservation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) DeleteReservation(ctx context.Context, in *DeleteReservationRequest, opts ...grpc.CallOption) (*DeleteReservationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteReservationResponse)
	err := c.cc.Invoke(ctx, ReservationService_DeleteReservation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetReservations(ctx context.Context, in *GetReservationsRequest, opts ...grpc.CallOption) (*GetReservationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetReservationsResponse)
	err := c.cc.Invoke(ctx, ReservationService_GetReservations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) AddReservationOrder(ctx context.Context, in *AddReservationOrderRequest, opts ...grpc.CallOption) (*AddReservationOrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddReservationOrderResponse)
	err := c.cc.Invoke(ctx, ReservationService_AddReservationOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetReservationOrder(ctx context.Context, in *GetReservationOrderRequest, opts ...grpc.CallOption) (*GetReservationOrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetReservationOrderResponse)
	err := c.cc.Invoke(ctx, ReservationService_GetReservationOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) UpdateReservationOrder(ctx context.Context, in *UpdateReservationOrderRequest, opts ...grpc.CallOption) (*UpdateReservationOrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateReservationOrderResponse)
	err := c.cc.Invoke(ctx, ReservationService_UpdateReservationOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) DeleteReservationOrder(ctx context.Context, in *DeleteReservationOrderRequest, opts ...grpc.CallOption) (*DeleteReservationOrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteReservationOrderResponse)
	err := c.cc.Invoke(ctx, ReservationService_DeleteReservationOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetReservationOrders(ctx context.Context, in *GetReservationOrdersRequest, opts ...grpc.CallOption) (*GetReservationOrdersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetReservationOrdersResponse)
	err := c.cc.Invoke(ctx, ReservationService_GetReservationOrders_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) AddMenu(ctx context.Context, in *AddMenuRequest, opts ...grpc.CallOption) (*AddMenuResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddMenuResponse)
	err := c.cc.Invoke(ctx, ReservationService_AddMenu_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetMenu(ctx context.Context, in *GetMenuRequest, opts ...grpc.CallOption) (*GetMenuResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMenuResponse)
	err := c.cc.Invoke(ctx, ReservationService_GetMenu_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) UpdateMenu(ctx context.Context, in *UpdateMenuRequest, opts ...grpc.CallOption) (*UpdateMenuResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateMenuResponse)
	err := c.cc.Invoke(ctx, ReservationService_UpdateMenu_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) DeleteMenu(ctx context.Context, in *DeleteMenuRequest, opts ...grpc.CallOption) (*DeleteMenuResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteMenuResponse)
	err := c.cc.Invoke(ctx, ReservationService_DeleteMenu_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetMenus(ctx context.Context, in *GetMenusRequest, opts ...grpc.CallOption) (*GetMenusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMenusResponse)
	err := c.cc.Invoke(ctx, ReservationService_GetMenus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReservationServiceServer is the server API for ReservationService service.
// All implementations must embed UnimplementedReservationServiceServer
// for forward compatibility
type ReservationServiceServer interface {
	// Restaurant methods
	AddRestaurant(context.Context, *AddRestaurantRequest) (*AddRestaurantResponse, error)
	GetRestaurant(context.Context, *GetRestaurantRequest) (*GetRestaurantResponse, error)
	UpdateRestaurant(context.Context, *UpdateRestaurantRequest) (*UpdateRestaurantResponse, error)
	DeleteRestaurant(context.Context, *DeleteRestaurantRequest) (*DeleteRestaurantResponse, error)
	GetRestaurants(context.Context, *GetRestaurantsRequest) (*GetRestaurantsResponse, error)
	// Reservation methods
	AddReservation(context.Context, *AddReservationRequest) (*AddReservationResponse, error)
	GetReservation(context.Context, *GetReservationRequest) (*GetReservationResponse, error)
	UpdateReservation(context.Context, *UpdateReservationRequest) (*UpdateReservationResponse, error)
	DeleteReservation(context.Context, *DeleteReservationRequest) (*DeleteReservationResponse, error)
	GetReservations(context.Context, *GetReservationsRequest) (*GetReservationsResponse, error)
	// Reservation Order methods
	AddReservationOrder(context.Context, *AddReservationOrderRequest) (*AddReservationOrderResponse, error)
	GetReservationOrder(context.Context, *GetReservationOrderRequest) (*GetReservationOrderResponse, error)
	UpdateReservationOrder(context.Context, *UpdateReservationOrderRequest) (*UpdateReservationOrderResponse, error)
	DeleteReservationOrder(context.Context, *DeleteReservationOrderRequest) (*DeleteReservationOrderResponse, error)
	GetReservationOrders(context.Context, *GetReservationOrdersRequest) (*GetReservationOrdersResponse, error)
	// Menu methods
	AddMenu(context.Context, *AddMenuRequest) (*AddMenuResponse, error)
	GetMenu(context.Context, *GetMenuRequest) (*GetMenuResponse, error)
	UpdateMenu(context.Context, *UpdateMenuRequest) (*UpdateMenuResponse, error)
	DeleteMenu(context.Context, *DeleteMenuRequest) (*DeleteMenuResponse, error)
	GetMenus(context.Context, *GetMenusRequest) (*GetMenusResponse, error)
	mustEmbedUnimplementedReservationServiceServer()
}

// UnimplementedReservationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReservationServiceServer struct {
}

func (UnimplementedReservationServiceServer) AddRestaurant(context.Context, *AddRestaurantRequest) (*AddRestaurantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRestaurant not implemented")
}
func (UnimplementedReservationServiceServer) GetRestaurant(context.Context, *GetRestaurantRequest) (*GetRestaurantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRestaurant not implemented")
}
func (UnimplementedReservationServiceServer) UpdateRestaurant(context.Context, *UpdateRestaurantRequest) (*UpdateRestaurantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRestaurant not implemented")
}
func (UnimplementedReservationServiceServer) DeleteRestaurant(context.Context, *DeleteRestaurantRequest) (*DeleteRestaurantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRestaurant not implemented")
}
func (UnimplementedReservationServiceServer) GetRestaurants(context.Context, *GetRestaurantsRequest) (*GetRestaurantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRestaurants not implemented")
}
func (UnimplementedReservationServiceServer) AddReservation(context.Context, *AddReservationRequest) (*AddReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddReservation not implemented")
}
func (UnimplementedReservationServiceServer) GetReservation(context.Context, *GetReservationRequest) (*GetReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReservation not implemented")
}
func (UnimplementedReservationServiceServer) UpdateReservation(context.Context, *UpdateReservationRequest) (*UpdateReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateReservation not implemented")
}
func (UnimplementedReservationServiceServer) DeleteReservation(context.Context, *DeleteReservationRequest) (*DeleteReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteReservation not implemented")
}
func (UnimplementedReservationServiceServer) GetReservations(context.Context, *GetReservationsRequest) (*GetReservationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReservations not implemented")
}
func (UnimplementedReservationServiceServer) AddReservationOrder(context.Context, *AddReservationOrderRequest) (*AddReservationOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddReservationOrder not implemented")
}
func (UnimplementedReservationServiceServer) GetReservationOrder(context.Context, *GetReservationOrderRequest) (*GetReservationOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReservationOrder not implemented")
}
func (UnimplementedReservationServiceServer) UpdateReservationOrder(context.Context, *UpdateReservationOrderRequest) (*UpdateReservationOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateReservationOrder not implemented")
}
func (UnimplementedReservationServiceServer) DeleteReservationOrder(context.Context, *DeleteReservationOrderRequest) (*DeleteReservationOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteReservationOrder not implemented")
}
func (UnimplementedReservationServiceServer) GetReservationOrders(context.Context, *GetReservationOrdersRequest) (*GetReservationOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReservationOrders not implemented")
}
func (UnimplementedReservationServiceServer) AddMenu(context.Context, *AddMenuRequest) (*AddMenuResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMenu not implemented")
}
func (UnimplementedReservationServiceServer) GetMenu(context.Context, *GetMenuRequest) (*GetMenuResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMenu not implemented")
}
func (UnimplementedReservationServiceServer) UpdateMenu(context.Context, *UpdateMenuRequest) (*UpdateMenuResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMenu not implemented")
}
func (UnimplementedReservationServiceServer) DeleteMenu(context.Context, *DeleteMenuRequest) (*DeleteMenuResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMenu not implemented")
}
func (UnimplementedReservationServiceServer) GetMenus(context.Context, *GetMenusRequest) (*GetMenusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMenus not implemented")
}
func (UnimplementedReservationServiceServer) mustEmbedUnimplementedReservationServiceServer() {}

// UnsafeReservationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReservationServiceServer will
// result in compilation errors.
type UnsafeReservationServiceServer interface {
	mustEmbedUnimplementedReservationServiceServer()
}

func RegisterReservationServiceServer(s grpc.ServiceRegistrar, srv ReservationServiceServer) {
	s.RegisterService(&ReservationService_ServiceDesc, srv)
}

func _ReservationService_AddRestaurant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRestaurantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).AddRestaurant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_AddRestaurant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).AddRestaurant(ctx, req.(*AddRestaurantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetRestaurant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRestaurantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetRestaurant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_GetRestaurant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetRestaurant(ctx, req.(*GetRestaurantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_UpdateRestaurant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRestaurantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).UpdateRestaurant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_UpdateRestaurant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).UpdateRestaurant(ctx, req.(*UpdateRestaurantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_DeleteRestaurant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRestaurantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).DeleteRestaurant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_DeleteRestaurant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).DeleteRestaurant(ctx, req.(*DeleteRestaurantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetRestaurants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRestaurantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetRestaurants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_GetRestaurants_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetRestaurants(ctx, req.(*GetRestaurantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_AddReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).AddReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_AddReservation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).AddReservation(ctx, req.(*AddReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_GetReservation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetReservation(ctx, req.(*GetReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_UpdateReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).UpdateReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_UpdateReservation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).UpdateReservation(ctx, req.(*UpdateReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_DeleteReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).DeleteReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_DeleteReservation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).DeleteReservation(ctx, req.(*DeleteReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetReservations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReservationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetReservations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_GetReservations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetReservations(ctx, req.(*GetReservationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_AddReservationOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddReservationOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).AddReservationOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_AddReservationOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).AddReservationOrder(ctx, req.(*AddReservationOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetReservationOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReservationOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetReservationOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_GetReservationOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetReservationOrder(ctx, req.(*GetReservationOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_UpdateReservationOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReservationOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).UpdateReservationOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_UpdateReservationOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).UpdateReservationOrder(ctx, req.(*UpdateReservationOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_DeleteReservationOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReservationOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).DeleteReservationOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_DeleteReservationOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).DeleteReservationOrder(ctx, req.(*DeleteReservationOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetReservationOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReservationOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetReservationOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_GetReservationOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetReservationOrders(ctx, req.(*GetReservationOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_AddMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMenuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).AddMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_AddMenu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).AddMenu(ctx, req.(*AddMenuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMenuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_GetMenu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetMenu(ctx, req.(*GetMenuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_UpdateMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMenuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).UpdateMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_UpdateMenu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).UpdateMenu(ctx, req.(*UpdateMenuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_DeleteMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMenuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).DeleteMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_DeleteMenu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).DeleteMenu(ctx, req.(*DeleteMenuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetMenus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMenusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetMenus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_GetMenus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetMenus(ctx, req.(*GetMenusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReservationService_ServiceDesc is the grpc.ServiceDesc for ReservationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReservationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "res_service.reservation_service",
	HandlerType: (*ReservationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "add_restaurant",
			Handler:    _ReservationService_AddRestaurant_Handler,
		},
		{
			MethodName: "get_restaurant",
			Handler:    _ReservationService_GetRestaurant_Handler,
		},
		{
			MethodName: "update_restaurant",
			Handler:    _ReservationService_UpdateRestaurant_Handler,
		},
		{
			MethodName: "delete_restaurant",
			Handler:    _ReservationService_DeleteRestaurant_Handler,
		},
		{
			MethodName: "get_restaurants",
			Handler:    _ReservationService_GetRestaurants_Handler,
		},
		{
			MethodName: "add_reservation",
			Handler:    _ReservationService_AddReservation_Handler,
		},
		{
			MethodName: "get_reservation",
			Handler:    _ReservationService_GetReservation_Handler,
		},
		{
			MethodName: "update_reservation",
			Handler:    _ReservationService_UpdateReservation_Handler,
		},
		{
			MethodName: "delete_reservation",
			Handler:    _ReservationService_DeleteReservation_Handler,
		},
		{
			MethodName: "get_reservations",
			Handler:    _ReservationService_GetReservations_Handler,
		},
		{
			MethodName: "add_reservation_order",
			Handler:    _ReservationService_AddReservationOrder_Handler,
		},
		{
			MethodName: "get_reservation_order",
			Handler:    _ReservationService_GetReservationOrder_Handler,
		},
		{
			MethodName: "update_reservation_order",
			Handler:    _ReservationService_UpdateReservationOrder_Handler,
		},
		{
			MethodName: "delete_reservation_order",
			Handler:    _ReservationService_DeleteReservationOrder_Handler,
		},
		{
			MethodName: "get_reservation_orders",
			Handler:    _ReservationService_GetReservationOrders_Handler,
		},
		{
			MethodName: "add_menu",
			Handler:    _ReservationService_AddMenu_Handler,
		},
		{
			MethodName: "get_menu",
			Handler:    _ReservationService_GetMenu_Handler,
		},
		{
			MethodName: "update_menu",
			Handler:    _ReservationService_UpdateMenu_Handler,
		},
		{
			MethodName: "delete_menu",
			Handler:    _ReservationService_DeleteMenu_Handler,
		},
		{
			MethodName: "get_menus",
			Handler:    _ReservationService_GetMenus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reservation/reservation_servcie.proto",
}
