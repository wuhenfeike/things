// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: proto/timedjob.proto

package timedjob

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

// TimedJobClient is the client API for TimedJob service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TimedJobClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type timedJobClient struct {
	cc grpc.ClientConnInterface
}

func NewTimedJobClient(cc grpc.ClientConnInterface) TimedJobClient {
	return &timedJobClient{cc}
}

func (c *timedJobClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/timedjob.timedJob/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TimedJobServer is the server API for TimedJob service.
// All implementations must embed UnimplementedTimedJobServer
// for forward compatibility
type TimedJobServer interface {
	Ping(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedTimedJobServer()
}

// UnimplementedTimedJobServer must be embedded to have forward compatible implementations.
type UnimplementedTimedJobServer struct {
}

func (UnimplementedTimedJobServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedTimedJobServer) mustEmbedUnimplementedTimedJobServer() {}

// UnsafeTimedJobServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TimedJobServer will
// result in compilation errors.
type UnsafeTimedJobServer interface {
	mustEmbedUnimplementedTimedJobServer()
}

func RegisterTimedJobServer(s grpc.ServiceRegistrar, srv TimedJobServer) {
	s.RegisterService(&TimedJob_ServiceDesc, srv)
}

func _TimedJob_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimedJobServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/timedjob.timedJob/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimedJobServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// TimedJob_ServiceDesc is the grpc.ServiceDesc for TimedJob service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TimedJob_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "timedjob.timedJob",
	HandlerType: (*TimedJobServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _TimedJob_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/timedjob.proto",
}