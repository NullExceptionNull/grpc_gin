// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// GobaseClient is the client API for Gobase service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GobaseClient interface {
	CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*UserInfo, error)
}

type gobaseClient struct {
	cc grpc.ClientConnInterface
}

func NewGobaseClient(cc grpc.ClientConnInterface) GobaseClient {
	return &gobaseClient{cc}
}

func (c *gobaseClient) CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*UserInfo, error) {
	out := new(UserInfo)
	err := c.cc.Invoke(ctx, "/gobase.v1.Gobase/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GobaseServer is the server API for Gobase service.
// All implementations must embed UnimplementedGobaseServer
// for forward compatibility
type GobaseServer interface {
	CreateUser(context.Context, *CreateUserReq) (*UserInfo, error)
	mustEmbedUnimplementedGobaseServer()
}

// UnimplementedGobaseServer must be embedded to have forward compatible implementations.
type UnimplementedGobaseServer struct {
}

func (UnimplementedGobaseServer) CreateUser(context.Context, *CreateUserReq) (*UserInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedGobaseServer) mustEmbedUnimplementedGobaseServer() {}

// UnsafeGobaseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GobaseServer will
// result in compilation errors.
type UnsafeGobaseServer interface {
	mustEmbedUnimplementedGobaseServer()
}

func RegisterGobaseServer(s grpc.ServiceRegistrar, srv GobaseServer) {
	s.RegisterService(&Gobase_ServiceDesc, srv)
}

func _Gobase_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GobaseServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gobase.v1.Gobase/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GobaseServer).CreateUser(ctx, req.(*CreateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Gobase_ServiceDesc is the grpc.ServiceDesc for Gobase service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gobase_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gobase.v1.Gobase",
	HandlerType: (*GobaseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _Gobase_CreateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gobase.proto",
}
