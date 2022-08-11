// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: proto/user.proto

package user

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

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	CoreCreate(ctx context.Context, in *UserCoreCreateReq, opts ...grpc.CallOption) (*UserCoreCreateResp, error)
	InfoCreate(ctx context.Context, in *UserInfoCreateReq, opts ...grpc.CallOption) (*Response, error)
	GetUserCoreList(ctx context.Context, in *GetUserCoreListReq, opts ...grpc.CallOption) (*GetUserCoreListResp, error)
	InfoUpdate(ctx context.Context, in *UserInfoUpdateReq, opts ...grpc.CallOption) (*Response, error)
	Read(ctx context.Context, in *UserReadReq, opts ...grpc.CallOption) (*UserReadResp, error)
	InfoDelete(ctx context.Context, in *UserInfoDeleteReq, opts ...grpc.CallOption) (*Response, error)
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
	CheckToken(ctx context.Context, in *CheckTokenReq, opts ...grpc.CallOption) (*CheckTokenResp, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) CoreCreate(ctx context.Context, in *UserCoreCreateReq, opts ...grpc.CallOption) (*UserCoreCreateResp, error) {
	out := new(UserCoreCreateResp)
	err := c.cc.Invoke(ctx, "/user.User/coreCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) InfoCreate(ctx context.Context, in *UserInfoCreateReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/user.User/infoCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserCoreList(ctx context.Context, in *GetUserCoreListReq, opts ...grpc.CallOption) (*GetUserCoreListResp, error) {
	out := new(GetUserCoreListResp)
	err := c.cc.Invoke(ctx, "/user.User/getUserCoreList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) InfoUpdate(ctx context.Context, in *UserInfoUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/user.User/infoUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Read(ctx context.Context, in *UserReadReq, opts ...grpc.CallOption) (*UserReadResp, error) {
	out := new(UserReadResp)
	err := c.cc.Invoke(ctx, "/user.User/read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) InfoDelete(ctx context.Context, in *UserInfoDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/user.User/infoDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, "/user.User/login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) CheckToken(ctx context.Context, in *CheckTokenReq, opts ...grpc.CallOption) (*CheckTokenResp, error) {
	out := new(CheckTokenResp)
	err := c.cc.Invoke(ctx, "/user.User/checkToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	CoreCreate(context.Context, *UserCoreCreateReq) (*UserCoreCreateResp, error)
	InfoCreate(context.Context, *UserInfoCreateReq) (*Response, error)
	GetUserCoreList(context.Context, *GetUserCoreListReq) (*GetUserCoreListResp, error)
	InfoUpdate(context.Context, *UserInfoUpdateReq) (*Response, error)
	Read(context.Context, *UserReadReq) (*UserReadResp, error)
	InfoDelete(context.Context, *UserInfoDeleteReq) (*Response, error)
	Login(context.Context, *LoginReq) (*LoginResp, error)
	CheckToken(context.Context, *CheckTokenReq) (*CheckTokenResp, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) CoreCreate(context.Context, *UserCoreCreateReq) (*UserCoreCreateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CoreCreate not implemented")
}
func (UnimplementedUserServer) InfoCreate(context.Context, *UserInfoCreateReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InfoCreate not implemented")
}
func (UnimplementedUserServer) GetUserCoreList(context.Context, *GetUserCoreListReq) (*GetUserCoreListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserCoreList not implemented")
}
func (UnimplementedUserServer) InfoUpdate(context.Context, *UserInfoUpdateReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InfoUpdate not implemented")
}
func (UnimplementedUserServer) Read(context.Context, *UserReadReq) (*UserReadResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedUserServer) InfoDelete(context.Context, *UserInfoDeleteReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InfoDelete not implemented")
}
func (UnimplementedUserServer) Login(context.Context, *LoginReq) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServer) CheckToken(context.Context, *CheckTokenReq) (*CheckTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckToken not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_CoreCreate_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(UserCoreCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CoreCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/coreCreate",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(UserServer).CoreCreate(ctx, req.(*UserCoreCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_InfoCreate_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(UserInfoCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).InfoCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/infoCreate",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(UserServer).InfoCreate(ctx, req.(*UserInfoCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserCoreList_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(GetUserCoreListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserCoreList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/getUserCoreList",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(UserServer).GetUserCoreList(ctx, req.(*GetUserCoreListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_InfoUpdate_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(UserInfoUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).InfoUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/infoUpdate",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(UserServer).InfoUpdate(ctx, req.(*UserInfoUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Read_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(UserReadReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/read",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(UserServer).Read(ctx, req.(*UserReadReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_InfoDelete_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(UserInfoDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).InfoDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/infoDelete",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(UserServer).InfoDelete(ctx, req.(*UserInfoDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Login_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/login",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(UserServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_CheckToken_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(CheckTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CheckToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/checkToken",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(UserServer).CheckToken(ctx, req.(*CheckTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "coreCreate",
			Handler:    _User_CoreCreate_Handler,
		},
		{
			MethodName: "infoCreate",
			Handler:    _User_InfoCreate_Handler,
		},
		{
			MethodName: "getUserCoreList",
			Handler:    _User_GetUserCoreList_Handler,
		},
		{
			MethodName: "infoUpdate",
			Handler:    _User_InfoUpdate_Handler,
		},
		{
			MethodName: "read",
			Handler:    _User_Read_Handler,
		},
		{
			MethodName: "infoDelete",
			Handler:    _User_InfoDelete_Handler,
		},
		{
			MethodName: "login",
			Handler:    _User_Login_Handler,
		},
		{
			MethodName: "checkToken",
			Handler:    _User_CheckToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/user.proto",
}
