// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: proto/vid.proto

package vid

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
	VidmgrInfoManage_VidmgrInfoCreate_FullMethodName = "/vid.VidmgrInfoManage/vidmgrInfoCreate"
	VidmgrInfoManage_VidmgrInfoUpdate_FullMethodName = "/vid.VidmgrInfoManage/vidmgrInfoUpdate"
	VidmgrInfoManage_VidmgrInfoDelete_FullMethodName = "/vid.VidmgrInfoManage/vidmgrInfoDelete"
	VidmgrInfoManage_VidmgrInfoIndex_FullMethodName  = "/vid.VidmgrInfoManage/vidmgrInfoIndex"
	VidmgrInfoManage_VidmgrInfoRead_FullMethodName   = "/vid.VidmgrInfoManage/vidmgrInfoRead"
	VidmgrInfoManage_VidmgrInfoCount_FullMethodName  = "/vid.VidmgrInfoManage/vidmgrInfoCount"
)

// VidmgrInfoManageClient is the client API for VidmgrInfoManage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VidmgrInfoManageClient interface {
	// 新建服务
	VidmgrInfoCreate(ctx context.Context, in *VidmgrInfo, opts ...grpc.CallOption) (*Response, error)
	// 更新服务
	VidmgrInfoUpdate(ctx context.Context, in *VidmgrInfo, opts ...grpc.CallOption) (*Response, error)
	// 删除服务
	VidmgrInfoDelete(ctx context.Context, in *VidmgrInfoDeleteReq, opts ...grpc.CallOption) (*Response, error)
	// 获取服务列表
	VidmgrInfoIndex(ctx context.Context, in *VidmgrInfoIndexReq, opts ...grpc.CallOption) (*VidmgrInfoIndexResp, error)
	// 获取服务信息详情
	VidmgrInfoRead(ctx context.Context, in *VidmgrInfoReadReq, opts ...grpc.CallOption) (*VidmgrInfo, error)
	// 获取服务统计  在线，离线，未激活
	VidmgrInfoCount(ctx context.Context, in *VidmgrInfoCountReq, opts ...grpc.CallOption) (*VidmgrInfoCountResp, error)
}

type vidmgrInfoManageClient struct {
	cc grpc.ClientConnInterface
}

func NewVidmgrInfoManageClient(cc grpc.ClientConnInterface) VidmgrInfoManageClient {
	return &vidmgrInfoManageClient{cc}
}

func (c *vidmgrInfoManageClient) VidmgrInfoCreate(ctx context.Context, in *VidmgrInfo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, VidmgrInfoManage_VidmgrInfoCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vidmgrInfoManageClient) VidmgrInfoUpdate(ctx context.Context, in *VidmgrInfo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, VidmgrInfoManage_VidmgrInfoUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vidmgrInfoManageClient) VidmgrInfoDelete(ctx context.Context, in *VidmgrInfoDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, VidmgrInfoManage_VidmgrInfoDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vidmgrInfoManageClient) VidmgrInfoIndex(ctx context.Context, in *VidmgrInfoIndexReq, opts ...grpc.CallOption) (*VidmgrInfoIndexResp, error) {
	out := new(VidmgrInfoIndexResp)
	err := c.cc.Invoke(ctx, VidmgrInfoManage_VidmgrInfoIndex_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vidmgrInfoManageClient) VidmgrInfoRead(ctx context.Context, in *VidmgrInfoReadReq, opts ...grpc.CallOption) (*VidmgrInfo, error) {
	out := new(VidmgrInfo)
	err := c.cc.Invoke(ctx, VidmgrInfoManage_VidmgrInfoRead_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vidmgrInfoManageClient) VidmgrInfoCount(ctx context.Context, in *VidmgrInfoCountReq, opts ...grpc.CallOption) (*VidmgrInfoCountResp, error) {
	out := new(VidmgrInfoCountResp)
	err := c.cc.Invoke(ctx, VidmgrInfoManage_VidmgrInfoCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VidmgrInfoManageServer is the server API for VidmgrInfoManage service.
// All implementations must embed UnimplementedVidmgrInfoManageServer
// for forward compatibility
type VidmgrInfoManageServer interface {
	// 新建服务
	VidmgrInfoCreate(context.Context, *VidmgrInfo) (*Response, error)
	// 更新服务
	VidmgrInfoUpdate(context.Context, *VidmgrInfo) (*Response, error)
	// 删除服务
	VidmgrInfoDelete(context.Context, *VidmgrInfoDeleteReq) (*Response, error)
	// 获取服务列表
	VidmgrInfoIndex(context.Context, *VidmgrInfoIndexReq) (*VidmgrInfoIndexResp, error)
	// 获取服务信息详情
	VidmgrInfoRead(context.Context, *VidmgrInfoReadReq) (*VidmgrInfo, error)
	// 获取服务统计  在线，离线，未激活
	VidmgrInfoCount(context.Context, *VidmgrInfoCountReq) (*VidmgrInfoCountResp, error)
	mustEmbedUnimplementedVidmgrInfoManageServer()
}

// UnimplementedVidmgrInfoManageServer must be embedded to have forward compatible implementations.
type UnimplementedVidmgrInfoManageServer struct {
}

func (UnimplementedVidmgrInfoManageServer) VidmgrInfoCreate(context.Context, *VidmgrInfo) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VidmgrInfoCreate not implemented")
}
func (UnimplementedVidmgrInfoManageServer) VidmgrInfoUpdate(context.Context, *VidmgrInfo) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VidmgrInfoUpdate not implemented")
}
func (UnimplementedVidmgrInfoManageServer) VidmgrInfoDelete(context.Context, *VidmgrInfoDeleteReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VidmgrInfoDelete not implemented")
}
func (UnimplementedVidmgrInfoManageServer) VidmgrInfoIndex(context.Context, *VidmgrInfoIndexReq) (*VidmgrInfoIndexResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VidmgrInfoIndex not implemented")
}
func (UnimplementedVidmgrInfoManageServer) VidmgrInfoRead(context.Context, *VidmgrInfoReadReq) (*VidmgrInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VidmgrInfoRead not implemented")
}
func (UnimplementedVidmgrInfoManageServer) VidmgrInfoCount(context.Context, *VidmgrInfoCountReq) (*VidmgrInfoCountResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VidmgrInfoCount not implemented")
}
func (UnimplementedVidmgrInfoManageServer) mustEmbedUnimplementedVidmgrInfoManageServer() {}

// UnsafeVidmgrInfoManageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VidmgrInfoManageServer will
// result in compilation errors.
type UnsafeVidmgrInfoManageServer interface {
	mustEmbedUnimplementedVidmgrInfoManageServer()
}

func RegisterVidmgrInfoManageServer(s grpc.ServiceRegistrar, srv VidmgrInfoManageServer) {
	s.RegisterService(&VidmgrInfoManage_ServiceDesc, srv)
}

func _VidmgrInfoManage_VidmgrInfoCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VidmgrInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VidmgrInfoManageServer).VidmgrInfoCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VidmgrInfoManage_VidmgrInfoCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VidmgrInfoManageServer).VidmgrInfoCreate(ctx, req.(*VidmgrInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _VidmgrInfoManage_VidmgrInfoUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VidmgrInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VidmgrInfoManageServer).VidmgrInfoUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VidmgrInfoManage_VidmgrInfoUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VidmgrInfoManageServer).VidmgrInfoUpdate(ctx, req.(*VidmgrInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _VidmgrInfoManage_VidmgrInfoDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VidmgrInfoDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VidmgrInfoManageServer).VidmgrInfoDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VidmgrInfoManage_VidmgrInfoDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VidmgrInfoManageServer).VidmgrInfoDelete(ctx, req.(*VidmgrInfoDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VidmgrInfoManage_VidmgrInfoIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VidmgrInfoIndexReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VidmgrInfoManageServer).VidmgrInfoIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VidmgrInfoManage_VidmgrInfoIndex_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VidmgrInfoManageServer).VidmgrInfoIndex(ctx, req.(*VidmgrInfoIndexReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VidmgrInfoManage_VidmgrInfoRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VidmgrInfoReadReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VidmgrInfoManageServer).VidmgrInfoRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VidmgrInfoManage_VidmgrInfoRead_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VidmgrInfoManageServer).VidmgrInfoRead(ctx, req.(*VidmgrInfoReadReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VidmgrInfoManage_VidmgrInfoCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VidmgrInfoCountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VidmgrInfoManageServer).VidmgrInfoCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VidmgrInfoManage_VidmgrInfoCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VidmgrInfoManageServer).VidmgrInfoCount(ctx, req.(*VidmgrInfoCountReq))
	}
	return interceptor(ctx, in, info, handler)
}

// VidmgrInfoManage_ServiceDesc is the grpc.ServiceDesc for VidmgrInfoManage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VidmgrInfoManage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vid.VidmgrInfoManage",
	HandlerType: (*VidmgrInfoManageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "vidmgrInfoCreate",
			Handler:    _VidmgrInfoManage_VidmgrInfoCreate_Handler,
		},
		{
			MethodName: "vidmgrInfoUpdate",
			Handler:    _VidmgrInfoManage_VidmgrInfoUpdate_Handler,
		},
		{
			MethodName: "vidmgrInfoDelete",
			Handler:    _VidmgrInfoManage_VidmgrInfoDelete_Handler,
		},
		{
			MethodName: "vidmgrInfoIndex",
			Handler:    _VidmgrInfoManage_VidmgrInfoIndex_Handler,
		},
		{
			MethodName: "vidmgrInfoRead",
			Handler:    _VidmgrInfoManage_VidmgrInfoRead_Handler,
		},
		{
			MethodName: "vidmgrInfoCount",
			Handler:    _VidmgrInfoManage_VidmgrInfoCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/vid.proto",
}

const (
	VidmgrStreamManage_VidmgrStreamCreate_FullMethodName = "/vid.VidmgrStreamManage/vidmgrStreamCreate"
	VidmgrStreamManage_VidmgrStreamUpdate_FullMethodName = "/vid.VidmgrStreamManage/vidmgrStreamUpdate"
	VidmgrStreamManage_VidmgrStreamDelete_FullMethodName = "/vid.VidmgrStreamManage/vidmgrStreamDelete"
	VidmgrStreamManage_VidmgrStreamIndex_FullMethodName  = "/vid.VidmgrStreamManage/vidmgrStreamIndex"
	VidmgrStreamManage_VidmgrStreamRead_FullMethodName   = "/vid.VidmgrStreamManage/vidmgrStreamRead"
	VidmgrStreamManage_VidmgrStreamCount_FullMethodName  = "/vid.VidmgrStreamManage/vidmgrStreamCount"
)

// VidmgrStreamManageClient is the client API for VidmgrStreamManage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VidmgrStreamManageClient interface {
	// 流添加
	VidmgrStreamCreate(ctx context.Context, in *VidmgrStream, opts ...grpc.CallOption) (*Response, error)
	// 流更新
	VidmgrStreamUpdate(ctx context.Context, in *VidmgrStream, opts ...grpc.CallOption) (*Response, error)
	// 删除流
	VidmgrStreamDelete(ctx context.Context, in *VidmgrStreamDeleteReq, opts ...grpc.CallOption) (*Response, error)
	// 获取流列表
	VidmgrStreamIndex(ctx context.Context, in *VidmgrStreamIndexReq, opts ...grpc.CallOption) (*VidmgrStreamIndexResp, error)
	// 获取流信息详情
	VidmgrStreamRead(ctx context.Context, in *VidmgrStreamReadReq, opts ...grpc.CallOption) (*VidmgrStream, error)
	// 统计流 在线，离线，未激活
	VidmgrStreamCount(ctx context.Context, in *VidmgrStreamCountReq, opts ...grpc.CallOption) (*VidmgrStreamCountResp, error)
}

type vidmgrStreamManageClient struct {
	cc grpc.ClientConnInterface
}

func NewVidmgrStreamManageClient(cc grpc.ClientConnInterface) VidmgrStreamManageClient {
	return &vidmgrStreamManageClient{cc}
}

func (c *vidmgrStreamManageClient) VidmgrStreamCreate(ctx context.Context, in *VidmgrStream, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, VidmgrStreamManage_VidmgrStreamCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vidmgrStreamManageClient) VidmgrStreamUpdate(ctx context.Context, in *VidmgrStream, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, VidmgrStreamManage_VidmgrStreamUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vidmgrStreamManageClient) VidmgrStreamDelete(ctx context.Context, in *VidmgrStreamDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, VidmgrStreamManage_VidmgrStreamDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vidmgrStreamManageClient) VidmgrStreamIndex(ctx context.Context, in *VidmgrStreamIndexReq, opts ...grpc.CallOption) (*VidmgrStreamIndexResp, error) {
	out := new(VidmgrStreamIndexResp)
	err := c.cc.Invoke(ctx, VidmgrStreamManage_VidmgrStreamIndex_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vidmgrStreamManageClient) VidmgrStreamRead(ctx context.Context, in *VidmgrStreamReadReq, opts ...grpc.CallOption) (*VidmgrStream, error) {
	out := new(VidmgrStream)
	err := c.cc.Invoke(ctx, VidmgrStreamManage_VidmgrStreamRead_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vidmgrStreamManageClient) VidmgrStreamCount(ctx context.Context, in *VidmgrStreamCountReq, opts ...grpc.CallOption) (*VidmgrStreamCountResp, error) {
	out := new(VidmgrStreamCountResp)
	err := c.cc.Invoke(ctx, VidmgrStreamManage_VidmgrStreamCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VidmgrStreamManageServer is the server API for VidmgrStreamManage service.
// All implementations must embed UnimplementedVidmgrStreamManageServer
// for forward compatibility
type VidmgrStreamManageServer interface {
	// 流添加
	VidmgrStreamCreate(context.Context, *VidmgrStream) (*Response, error)
	// 流更新
	VidmgrStreamUpdate(context.Context, *VidmgrStream) (*Response, error)
	// 删除流
	VidmgrStreamDelete(context.Context, *VidmgrStreamDeleteReq) (*Response, error)
	// 获取流列表
	VidmgrStreamIndex(context.Context, *VidmgrStreamIndexReq) (*VidmgrStreamIndexResp, error)
	// 获取流信息详情
	VidmgrStreamRead(context.Context, *VidmgrStreamReadReq) (*VidmgrStream, error)
	// 统计流 在线，离线，未激活
	VidmgrStreamCount(context.Context, *VidmgrStreamCountReq) (*VidmgrStreamCountResp, error)
	mustEmbedUnimplementedVidmgrStreamManageServer()
}

// UnimplementedVidmgrStreamManageServer must be embedded to have forward compatible implementations.
type UnimplementedVidmgrStreamManageServer struct {
}

func (UnimplementedVidmgrStreamManageServer) VidmgrStreamCreate(context.Context, *VidmgrStream) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VidmgrStreamCreate not implemented")
}
func (UnimplementedVidmgrStreamManageServer) VidmgrStreamUpdate(context.Context, *VidmgrStream) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VidmgrStreamUpdate not implemented")
}
func (UnimplementedVidmgrStreamManageServer) VidmgrStreamDelete(context.Context, *VidmgrStreamDeleteReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VidmgrStreamDelete not implemented")
}
func (UnimplementedVidmgrStreamManageServer) VidmgrStreamIndex(context.Context, *VidmgrStreamIndexReq) (*VidmgrStreamIndexResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VidmgrStreamIndex not implemented")
}
func (UnimplementedVidmgrStreamManageServer) VidmgrStreamRead(context.Context, *VidmgrStreamReadReq) (*VidmgrStream, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VidmgrStreamRead not implemented")
}
func (UnimplementedVidmgrStreamManageServer) VidmgrStreamCount(context.Context, *VidmgrStreamCountReq) (*VidmgrStreamCountResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VidmgrStreamCount not implemented")
}
func (UnimplementedVidmgrStreamManageServer) mustEmbedUnimplementedVidmgrStreamManageServer() {}

// UnsafeVidmgrStreamManageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VidmgrStreamManageServer will
// result in compilation errors.
type UnsafeVidmgrStreamManageServer interface {
	mustEmbedUnimplementedVidmgrStreamManageServer()
}

func RegisterVidmgrStreamManageServer(s grpc.ServiceRegistrar, srv VidmgrStreamManageServer) {
	s.RegisterService(&VidmgrStreamManage_ServiceDesc, srv)
}

func _VidmgrStreamManage_VidmgrStreamCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VidmgrStream)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VidmgrStreamManageServer).VidmgrStreamCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VidmgrStreamManage_VidmgrStreamCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VidmgrStreamManageServer).VidmgrStreamCreate(ctx, req.(*VidmgrStream))
	}
	return interceptor(ctx, in, info, handler)
}

func _VidmgrStreamManage_VidmgrStreamUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VidmgrStream)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VidmgrStreamManageServer).VidmgrStreamUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VidmgrStreamManage_VidmgrStreamUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VidmgrStreamManageServer).VidmgrStreamUpdate(ctx, req.(*VidmgrStream))
	}
	return interceptor(ctx, in, info, handler)
}

func _VidmgrStreamManage_VidmgrStreamDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VidmgrStreamDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VidmgrStreamManageServer).VidmgrStreamDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VidmgrStreamManage_VidmgrStreamDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VidmgrStreamManageServer).VidmgrStreamDelete(ctx, req.(*VidmgrStreamDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VidmgrStreamManage_VidmgrStreamIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VidmgrStreamIndexReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VidmgrStreamManageServer).VidmgrStreamIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VidmgrStreamManage_VidmgrStreamIndex_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VidmgrStreamManageServer).VidmgrStreamIndex(ctx, req.(*VidmgrStreamIndexReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VidmgrStreamManage_VidmgrStreamRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VidmgrStreamReadReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VidmgrStreamManageServer).VidmgrStreamRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VidmgrStreamManage_VidmgrStreamRead_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VidmgrStreamManageServer).VidmgrStreamRead(ctx, req.(*VidmgrStreamReadReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VidmgrStreamManage_VidmgrStreamCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VidmgrStreamCountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VidmgrStreamManageServer).VidmgrStreamCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VidmgrStreamManage_VidmgrStreamCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VidmgrStreamManageServer).VidmgrStreamCount(ctx, req.(*VidmgrStreamCountReq))
	}
	return interceptor(ctx, in, info, handler)
}

// VidmgrStreamManage_ServiceDesc is the grpc.ServiceDesc for VidmgrStreamManage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VidmgrStreamManage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vid.VidmgrStreamManage",
	HandlerType: (*VidmgrStreamManageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "vidmgrStreamCreate",
			Handler:    _VidmgrStreamManage_VidmgrStreamCreate_Handler,
		},
		{
			MethodName: "vidmgrStreamUpdate",
			Handler:    _VidmgrStreamManage_VidmgrStreamUpdate_Handler,
		},
		{
			MethodName: "vidmgrStreamDelete",
			Handler:    _VidmgrStreamManage_VidmgrStreamDelete_Handler,
		},
		{
			MethodName: "vidmgrStreamIndex",
			Handler:    _VidmgrStreamManage_VidmgrStreamIndex_Handler,
		},
		{
			MethodName: "vidmgrStreamRead",
			Handler:    _VidmgrStreamManage_VidmgrStreamRead_Handler,
		},
		{
			MethodName: "vidmgrStreamCount",
			Handler:    _VidmgrStreamManage_VidmgrStreamCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/vid.proto",
}

const (
	VidmgrConfigManage_VidmgrConfigCreate_FullMethodName = "/vid.VidmgrConfigManage/vidmgrConfigCreate"
	VidmgrConfigManage_VidmgrConfigDelete_FullMethodName = "/vid.VidmgrConfigManage/vidmgrConfigDelete"
	VidmgrConfigManage_VidmgrConfigUpdate_FullMethodName = "/vid.VidmgrConfigManage/vidmgrConfigUpdate"
	VidmgrConfigManage_VidmgrConfigIndex_FullMethodName  = "/vid.VidmgrConfigManage/vidmgrConfigIndex"
	VidmgrConfigManage_VidmgrConfigRead_FullMethodName   = "/vid.VidmgrConfigManage/vidmgrConfigRead"
)

// VidmgrConfigManageClient is the client API for VidmgrConfigManage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VidmgrConfigManageClient interface {
	// 新建配置
	VidmgrConfigCreate(ctx context.Context, in *VidmgrConfig, opts ...grpc.CallOption) (*Response, error)
	// 删除配置
	VidmgrConfigDelete(ctx context.Context, in *VidmgrConfigDeleteReq, opts ...grpc.CallOption) (*Response, error)
	// 更新配置
	VidmgrConfigUpdate(ctx context.Context, in *VidmgrConfig, opts ...grpc.CallOption) (*Response, error)
	// 配置列表
	VidmgrConfigIndex(ctx context.Context, in *VidmgrConfigIndexReq, opts ...grpc.CallOption) (*VidmgrConfigIndexResp, error)
	// 获取配置信息详情
	VidmgrConfigRead(ctx context.Context, in *VidmgrConfigReadReq, opts ...grpc.CallOption) (*VidmgrConfig, error)
}

type vidmgrConfigManageClient struct {
	cc grpc.ClientConnInterface
}

func NewVidmgrConfigManageClient(cc grpc.ClientConnInterface) VidmgrConfigManageClient {
	return &vidmgrConfigManageClient{cc}
}

func (c *vidmgrConfigManageClient) VidmgrConfigCreate(ctx context.Context, in *VidmgrConfig, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, VidmgrConfigManage_VidmgrConfigCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vidmgrConfigManageClient) VidmgrConfigDelete(ctx context.Context, in *VidmgrConfigDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, VidmgrConfigManage_VidmgrConfigDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vidmgrConfigManageClient) VidmgrConfigUpdate(ctx context.Context, in *VidmgrConfig, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, VidmgrConfigManage_VidmgrConfigUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vidmgrConfigManageClient) VidmgrConfigIndex(ctx context.Context, in *VidmgrConfigIndexReq, opts ...grpc.CallOption) (*VidmgrConfigIndexResp, error) {
	out := new(VidmgrConfigIndexResp)
	err := c.cc.Invoke(ctx, VidmgrConfigManage_VidmgrConfigIndex_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vidmgrConfigManageClient) VidmgrConfigRead(ctx context.Context, in *VidmgrConfigReadReq, opts ...grpc.CallOption) (*VidmgrConfig, error) {
	out := new(VidmgrConfig)
	err := c.cc.Invoke(ctx, VidmgrConfigManage_VidmgrConfigRead_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VidmgrConfigManageServer is the server API for VidmgrConfigManage service.
// All implementations must embed UnimplementedVidmgrConfigManageServer
// for forward compatibility
type VidmgrConfigManageServer interface {
	// 新建配置
	VidmgrConfigCreate(context.Context, *VidmgrConfig) (*Response, error)
	// 删除配置
	VidmgrConfigDelete(context.Context, *VidmgrConfigDeleteReq) (*Response, error)
	// 更新配置
	VidmgrConfigUpdate(context.Context, *VidmgrConfig) (*Response, error)
	// 配置列表
	VidmgrConfigIndex(context.Context, *VidmgrConfigIndexReq) (*VidmgrConfigIndexResp, error)
	// 获取配置信息详情
	VidmgrConfigRead(context.Context, *VidmgrConfigReadReq) (*VidmgrConfig, error)
	mustEmbedUnimplementedVidmgrConfigManageServer()
}

// UnimplementedVidmgrConfigManageServer must be embedded to have forward compatible implementations.
type UnimplementedVidmgrConfigManageServer struct {
}

func (UnimplementedVidmgrConfigManageServer) VidmgrConfigCreate(context.Context, *VidmgrConfig) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VidmgrConfigCreate not implemented")
}
func (UnimplementedVidmgrConfigManageServer) VidmgrConfigDelete(context.Context, *VidmgrConfigDeleteReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VidmgrConfigDelete not implemented")
}
func (UnimplementedVidmgrConfigManageServer) VidmgrConfigUpdate(context.Context, *VidmgrConfig) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VidmgrConfigUpdate not implemented")
}
func (UnimplementedVidmgrConfigManageServer) VidmgrConfigIndex(context.Context, *VidmgrConfigIndexReq) (*VidmgrConfigIndexResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VidmgrConfigIndex not implemented")
}
func (UnimplementedVidmgrConfigManageServer) VidmgrConfigRead(context.Context, *VidmgrConfigReadReq) (*VidmgrConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VidmgrConfigRead not implemented")
}
func (UnimplementedVidmgrConfigManageServer) mustEmbedUnimplementedVidmgrConfigManageServer() {}

// UnsafeVidmgrConfigManageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VidmgrConfigManageServer will
// result in compilation errors.
type UnsafeVidmgrConfigManageServer interface {
	mustEmbedUnimplementedVidmgrConfigManageServer()
}

func RegisterVidmgrConfigManageServer(s grpc.ServiceRegistrar, srv VidmgrConfigManageServer) {
	s.RegisterService(&VidmgrConfigManage_ServiceDesc, srv)
}

func _VidmgrConfigManage_VidmgrConfigCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VidmgrConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VidmgrConfigManageServer).VidmgrConfigCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VidmgrConfigManage_VidmgrConfigCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VidmgrConfigManageServer).VidmgrConfigCreate(ctx, req.(*VidmgrConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _VidmgrConfigManage_VidmgrConfigDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VidmgrConfigDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VidmgrConfigManageServer).VidmgrConfigDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VidmgrConfigManage_VidmgrConfigDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VidmgrConfigManageServer).VidmgrConfigDelete(ctx, req.(*VidmgrConfigDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VidmgrConfigManage_VidmgrConfigUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VidmgrConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VidmgrConfigManageServer).VidmgrConfigUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VidmgrConfigManage_VidmgrConfigUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VidmgrConfigManageServer).VidmgrConfigUpdate(ctx, req.(*VidmgrConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _VidmgrConfigManage_VidmgrConfigIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VidmgrConfigIndexReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VidmgrConfigManageServer).VidmgrConfigIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VidmgrConfigManage_VidmgrConfigIndex_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VidmgrConfigManageServer).VidmgrConfigIndex(ctx, req.(*VidmgrConfigIndexReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VidmgrConfigManage_VidmgrConfigRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VidmgrConfigReadReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VidmgrConfigManageServer).VidmgrConfigRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VidmgrConfigManage_VidmgrConfigRead_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VidmgrConfigManageServer).VidmgrConfigRead(ctx, req.(*VidmgrConfigReadReq))
	}
	return interceptor(ctx, in, info, handler)
}

// VidmgrConfigManage_ServiceDesc is the grpc.ServiceDesc for VidmgrConfigManage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VidmgrConfigManage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vid.VidmgrConfigManage",
	HandlerType: (*VidmgrConfigManageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "vidmgrConfigCreate",
			Handler:    _VidmgrConfigManage_VidmgrConfigCreate_Handler,
		},
		{
			MethodName: "vidmgrConfigDelete",
			Handler:    _VidmgrConfigManage_VidmgrConfigDelete_Handler,
		},
		{
			MethodName: "vidmgrConfigUpdate",
			Handler:    _VidmgrConfigManage_VidmgrConfigUpdate_Handler,
		},
		{
			MethodName: "vidmgrConfigIndex",
			Handler:    _VidmgrConfigManage_VidmgrConfigIndex_Handler,
		},
		{
			MethodName: "vidmgrConfigRead",
			Handler:    _VidmgrConfigManage_VidmgrConfigRead_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/vid.proto",
}
