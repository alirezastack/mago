// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package mago

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

// MagoServiceClient is the client API for MagoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MagoServiceClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
}

type magoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMagoServiceClient(cc grpc.ClientConnInterface) MagoServiceClient {
	return &magoServiceClient{cc}
}

func (c *magoServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/mago.MagoService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MagoServiceServer is the server API for MagoService service.
// All implementations must embed UnimplementedMagoServiceServer
// for forward compatibility
type MagoServiceServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	mustEmbedUnimplementedMagoServiceServer()
}

// UnimplementedMagoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMagoServiceServer struct {
}

func (UnimplementedMagoServiceServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedMagoServiceServer) mustEmbedUnimplementedMagoServiceServer() {}

// UnsafeMagoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MagoServiceServer will
// result in compilation errors.
type UnsafeMagoServiceServer interface {
	mustEmbedUnimplementedMagoServiceServer()
}

func RegisterMagoServiceServer(s grpc.ServiceRegistrar, srv MagoServiceServer) {
	s.RegisterService(&MagoService_ServiceDesc, srv)
}

func _MagoService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MagoServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mago.MagoService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MagoServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MagoService_ServiceDesc is the grpc.ServiceDesc for MagoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MagoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mago.MagoService",
	HandlerType: (*MagoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _MagoService_CreateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "magopb/mago.proto",
}
