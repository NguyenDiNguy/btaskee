// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: services/booking/proto/booking.proto

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
	Booking_CreateTask_FullMethodName    = "/booking.Booking/CreateTask"
	Booking_GetTask_FullMethodName       = "/booking.Booking/GetTask"
	Booking_AcceptTask_FullMethodName    = "/booking.Booking/AcceptTask"
	Booking_ConfirmTasker_FullMethodName = "/booking.Booking/ConfirmTasker"
)

// BookingClient is the client API for Booking service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookingClient interface {
	CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error)
	GetTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*GetTaskResponse, error)
	AcceptTask(ctx context.Context, in *AcceptTaskRequest, opts ...grpc.CallOption) (*AcceptTaskResponse, error)
	ConfirmTasker(ctx context.Context, in *ConfirmTaskerRequest, opts ...grpc.CallOption) (*ConfirmTaskerResponse, error)
}

type bookingClient struct {
	cc grpc.ClientConnInterface
}

func NewBookingClient(cc grpc.ClientConnInterface) BookingClient {
	return &bookingClient{cc}
}

func (c *bookingClient) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTaskResponse)
	err := c.cc.Invoke(ctx, Booking_CreateTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingClient) GetTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*GetTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTaskResponse)
	err := c.cc.Invoke(ctx, Booking_GetTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingClient) AcceptTask(ctx context.Context, in *AcceptTaskRequest, opts ...grpc.CallOption) (*AcceptTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AcceptTaskResponse)
	err := c.cc.Invoke(ctx, Booking_AcceptTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingClient) ConfirmTasker(ctx context.Context, in *ConfirmTaskerRequest, opts ...grpc.CallOption) (*ConfirmTaskerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ConfirmTaskerResponse)
	err := c.cc.Invoke(ctx, Booking_ConfirmTasker_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookingServer is the server API for Booking service.
// All implementations must embed UnimplementedBookingServer
// for forward compatibility.
type BookingServer interface {
	CreateTask(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error)
	GetTask(context.Context, *GetTaskRequest) (*GetTaskResponse, error)
	AcceptTask(context.Context, *AcceptTaskRequest) (*AcceptTaskResponse, error)
	ConfirmTasker(context.Context, *ConfirmTaskerRequest) (*ConfirmTaskerResponse, error)
	mustEmbedUnimplementedBookingServer()
}

// UnimplementedBookingServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBookingServer struct{}

func (UnimplementedBookingServer) CreateTask(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (UnimplementedBookingServer) GetTask(context.Context, *GetTaskRequest) (*GetTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTask not implemented")
}
func (UnimplementedBookingServer) AcceptTask(context.Context, *AcceptTaskRequest) (*AcceptTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptTask not implemented")
}
func (UnimplementedBookingServer) ConfirmTasker(context.Context, *ConfirmTaskerRequest) (*ConfirmTaskerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmTasker not implemented")
}
func (UnimplementedBookingServer) mustEmbedUnimplementedBookingServer() {}
func (UnimplementedBookingServer) testEmbeddedByValue()                 {}

// UnsafeBookingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookingServer will
// result in compilation errors.
type UnsafeBookingServer interface {
	mustEmbedUnimplementedBookingServer()
}

func RegisterBookingServer(s grpc.ServiceRegistrar, srv BookingServer) {
	// If the following call pancis, it indicates UnimplementedBookingServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Booking_ServiceDesc, srv)
}

func _Booking_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Booking_CreateTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServer).CreateTask(ctx, req.(*CreateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Booking_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Booking_GetTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServer).GetTask(ctx, req.(*GetTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Booking_AcceptTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServer).AcceptTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Booking_AcceptTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServer).AcceptTask(ctx, req.(*AcceptTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Booking_ConfirmTasker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmTaskerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServer).ConfirmTasker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Booking_ConfirmTasker_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServer).ConfirmTasker(ctx, req.(*ConfirmTaskerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Booking_ServiceDesc is the grpc.ServiceDesc for Booking service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Booking_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "booking.Booking",
	HandlerType: (*BookingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTask",
			Handler:    _Booking_CreateTask_Handler,
		},
		{
			MethodName: "GetTask",
			Handler:    _Booking_GetTask_Handler,
		},
		{
			MethodName: "AcceptTask",
			Handler:    _Booking_AcceptTask_Handler,
		},
		{
			MethodName: "ConfirmTasker",
			Handler:    _Booking_ConfirmTasker_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/booking/proto/booking.proto",
}
