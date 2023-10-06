// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: proto/rule.proto

package rule

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

// RuleEngineClient is the client API for RuleEngine service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RuleEngineClient interface {
	FlowInfoCreate(ctx context.Context, in *FlowInfo, opts ...grpc.CallOption) (*WithID, error)
	FlowInfoUpdate(ctx context.Context, in *FlowInfo, opts ...grpc.CallOption) (*Empty, error)
	FlowInfoDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error)
	FlowInfoIndex(ctx context.Context, in *FlowInfoIndexReq, opts ...grpc.CallOption) (*FlowInfoIndexResp, error)
	FlowInfoRead(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*FlowInfo, error)
}

type ruleEngineClient struct {
	cc grpc.ClientConnInterface
}

func NewRuleEngineClient(cc grpc.ClientConnInterface) RuleEngineClient {
	return &ruleEngineClient{cc}
}

func (c *ruleEngineClient) FlowInfoCreate(ctx context.Context, in *FlowInfo, opts ...grpc.CallOption) (*WithID, error) {
	out := new(WithID)
	err := c.cc.Invoke(ctx, "/rule.RuleEngine/flowInfoCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ruleEngineClient) FlowInfoUpdate(ctx context.Context, in *FlowInfo, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/rule.RuleEngine/flowInfoUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ruleEngineClient) FlowInfoDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/rule.RuleEngine/flowInfoDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ruleEngineClient) FlowInfoIndex(ctx context.Context, in *FlowInfoIndexReq, opts ...grpc.CallOption) (*FlowInfoIndexResp, error) {
	out := new(FlowInfoIndexResp)
	err := c.cc.Invoke(ctx, "/rule.RuleEngine/flowInfoIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ruleEngineClient) FlowInfoRead(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*FlowInfo, error) {
	out := new(FlowInfo)
	err := c.cc.Invoke(ctx, "/rule.RuleEngine/flowInfoRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RuleEngineServer is the server API for RuleEngine service.
// All implementations must embed UnimplementedRuleEngineServer
// for forward compatibility
type RuleEngineServer interface {
	FlowInfoCreate(context.Context, *FlowInfo) (*WithID, error)
	FlowInfoUpdate(context.Context, *FlowInfo) (*Empty, error)
	FlowInfoDelete(context.Context, *WithID) (*Empty, error)
	FlowInfoIndex(context.Context, *FlowInfoIndexReq) (*FlowInfoIndexResp, error)
	FlowInfoRead(context.Context, *WithID) (*FlowInfo, error)
	mustEmbedUnimplementedRuleEngineServer()
}

// UnimplementedRuleEngineServer must be embedded to have forward compatible implementations.
type UnimplementedRuleEngineServer struct {
}

func (UnimplementedRuleEngineServer) FlowInfoCreate(context.Context, *FlowInfo) (*WithID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FlowInfoCreate not implemented")
}
func (UnimplementedRuleEngineServer) FlowInfoUpdate(context.Context, *FlowInfo) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FlowInfoUpdate not implemented")
}
func (UnimplementedRuleEngineServer) FlowInfoDelete(context.Context, *WithID) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FlowInfoDelete not implemented")
}
func (UnimplementedRuleEngineServer) FlowInfoIndex(context.Context, *FlowInfoIndexReq) (*FlowInfoIndexResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FlowInfoIndex not implemented")
}
func (UnimplementedRuleEngineServer) FlowInfoRead(context.Context, *WithID) (*FlowInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FlowInfoRead not implemented")
}
func (UnimplementedRuleEngineServer) mustEmbedUnimplementedRuleEngineServer() {}

// UnsafeRuleEngineServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RuleEngineServer will
// result in compilation errors.
type UnsafeRuleEngineServer interface {
	mustEmbedUnimplementedRuleEngineServer()
}

func RegisterRuleEngineServer(s grpc.ServiceRegistrar, srv RuleEngineServer) {
	s.RegisterService(&RuleEngine_ServiceDesc, srv)
}

func _RuleEngine_FlowInfoCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlowInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuleEngineServer).FlowInfoCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rule.RuleEngine/flowInfoCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuleEngineServer).FlowInfoCreate(ctx, req.(*FlowInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _RuleEngine_FlowInfoUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlowInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuleEngineServer).FlowInfoUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rule.RuleEngine/flowInfoUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuleEngineServer).FlowInfoUpdate(ctx, req.(*FlowInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _RuleEngine_FlowInfoDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuleEngineServer).FlowInfoDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rule.RuleEngine/flowInfoDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuleEngineServer).FlowInfoDelete(ctx, req.(*WithID))
	}
	return interceptor(ctx, in, info, handler)
}

func _RuleEngine_FlowInfoIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlowInfoIndexReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuleEngineServer).FlowInfoIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rule.RuleEngine/flowInfoIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuleEngineServer).FlowInfoIndex(ctx, req.(*FlowInfoIndexReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RuleEngine_FlowInfoRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuleEngineServer).FlowInfoRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rule.RuleEngine/flowInfoRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuleEngineServer).FlowInfoRead(ctx, req.(*WithID))
	}
	return interceptor(ctx, in, info, handler)
}

// RuleEngine_ServiceDesc is the grpc.ServiceDesc for RuleEngine service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RuleEngine_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rule.RuleEngine",
	HandlerType: (*RuleEngineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "flowInfoCreate",
			Handler:    _RuleEngine_FlowInfoCreate_Handler,
		},
		{
			MethodName: "flowInfoUpdate",
			Handler:    _RuleEngine_FlowInfoUpdate_Handler,
		},
		{
			MethodName: "flowInfoDelete",
			Handler:    _RuleEngine_FlowInfoDelete_Handler,
		},
		{
			MethodName: "flowInfoIndex",
			Handler:    _RuleEngine_FlowInfoIndex_Handler,
		},
		{
			MethodName: "flowInfoRead",
			Handler:    _RuleEngine_FlowInfoRead_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/rule.proto",
}

// SceneLinkageClient is the client API for SceneLinkage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SceneLinkageClient interface {
	SceneInfoCreate(ctx context.Context, in *SceneInfo, opts ...grpc.CallOption) (*WithID, error)
	SceneInfoUpdate(ctx context.Context, in *SceneInfo, opts ...grpc.CallOption) (*Empty, error)
	SceneInfoDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error)
	SceneInfoIndex(ctx context.Context, in *SceneInfoIndexReq, opts ...grpc.CallOption) (*SceneInfoIndexResp, error)
	SceneInfoRead(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*SceneInfo, error)
	SceneManuallyTrigger(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error)
}

type sceneLinkageClient struct {
	cc grpc.ClientConnInterface
}

func NewSceneLinkageClient(cc grpc.ClientConnInterface) SceneLinkageClient {
	return &sceneLinkageClient{cc}
}

func (c *sceneLinkageClient) SceneInfoCreate(ctx context.Context, in *SceneInfo, opts ...grpc.CallOption) (*WithID, error) {
	out := new(WithID)
	err := c.cc.Invoke(ctx, "/rule.SceneLinkage/sceneInfoCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sceneLinkageClient) SceneInfoUpdate(ctx context.Context, in *SceneInfo, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/rule.SceneLinkage/sceneInfoUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sceneLinkageClient) SceneInfoDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/rule.SceneLinkage/sceneInfoDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sceneLinkageClient) SceneInfoIndex(ctx context.Context, in *SceneInfoIndexReq, opts ...grpc.CallOption) (*SceneInfoIndexResp, error) {
	out := new(SceneInfoIndexResp)
	err := c.cc.Invoke(ctx, "/rule.SceneLinkage/sceneInfoIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sceneLinkageClient) SceneInfoRead(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*SceneInfo, error) {
	out := new(SceneInfo)
	err := c.cc.Invoke(ctx, "/rule.SceneLinkage/sceneInfoRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sceneLinkageClient) SceneManuallyTrigger(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/rule.SceneLinkage/sceneManuallyTrigger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SceneLinkageServer is the server API for SceneLinkage service.
// All implementations must embed UnimplementedSceneLinkageServer
// for forward compatibility
type SceneLinkageServer interface {
	SceneInfoCreate(context.Context, *SceneInfo) (*WithID, error)
	SceneInfoUpdate(context.Context, *SceneInfo) (*Empty, error)
	SceneInfoDelete(context.Context, *WithID) (*Empty, error)
	SceneInfoIndex(context.Context, *SceneInfoIndexReq) (*SceneInfoIndexResp, error)
	SceneInfoRead(context.Context, *WithID) (*SceneInfo, error)
	SceneManuallyTrigger(context.Context, *WithID) (*Empty, error)
	mustEmbedUnimplementedSceneLinkageServer()
}

// UnimplementedSceneLinkageServer must be embedded to have forward compatible implementations.
type UnimplementedSceneLinkageServer struct {
}

func (UnimplementedSceneLinkageServer) SceneInfoCreate(context.Context, *SceneInfo) (*WithID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SceneInfoCreate not implemented")
}
func (UnimplementedSceneLinkageServer) SceneInfoUpdate(context.Context, *SceneInfo) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SceneInfoUpdate not implemented")
}
func (UnimplementedSceneLinkageServer) SceneInfoDelete(context.Context, *WithID) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SceneInfoDelete not implemented")
}
func (UnimplementedSceneLinkageServer) SceneInfoIndex(context.Context, *SceneInfoIndexReq) (*SceneInfoIndexResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SceneInfoIndex not implemented")
}
func (UnimplementedSceneLinkageServer) SceneInfoRead(context.Context, *WithID) (*SceneInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SceneInfoRead not implemented")
}
func (UnimplementedSceneLinkageServer) SceneManuallyTrigger(context.Context, *WithID) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SceneManuallyTrigger not implemented")
}
func (UnimplementedSceneLinkageServer) mustEmbedUnimplementedSceneLinkageServer() {}

// UnsafeSceneLinkageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SceneLinkageServer will
// result in compilation errors.
type UnsafeSceneLinkageServer interface {
	mustEmbedUnimplementedSceneLinkageServer()
}

func RegisterSceneLinkageServer(s grpc.ServiceRegistrar, srv SceneLinkageServer) {
	s.RegisterService(&SceneLinkage_ServiceDesc, srv)
}

func _SceneLinkage_SceneInfoCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SceneInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SceneLinkageServer).SceneInfoCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rule.SceneLinkage/sceneInfoCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SceneLinkageServer).SceneInfoCreate(ctx, req.(*SceneInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _SceneLinkage_SceneInfoUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SceneInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SceneLinkageServer).SceneInfoUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rule.SceneLinkage/sceneInfoUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SceneLinkageServer).SceneInfoUpdate(ctx, req.(*SceneInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _SceneLinkage_SceneInfoDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SceneLinkageServer).SceneInfoDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rule.SceneLinkage/sceneInfoDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SceneLinkageServer).SceneInfoDelete(ctx, req.(*WithID))
	}
	return interceptor(ctx, in, info, handler)
}

func _SceneLinkage_SceneInfoIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SceneInfoIndexReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SceneLinkageServer).SceneInfoIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rule.SceneLinkage/sceneInfoIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SceneLinkageServer).SceneInfoIndex(ctx, req.(*SceneInfoIndexReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SceneLinkage_SceneInfoRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SceneLinkageServer).SceneInfoRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rule.SceneLinkage/sceneInfoRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SceneLinkageServer).SceneInfoRead(ctx, req.(*WithID))
	}
	return interceptor(ctx, in, info, handler)
}

func _SceneLinkage_SceneManuallyTrigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SceneLinkageServer).SceneManuallyTrigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rule.SceneLinkage/sceneManuallyTrigger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SceneLinkageServer).SceneManuallyTrigger(ctx, req.(*WithID))
	}
	return interceptor(ctx, in, info, handler)
}

// SceneLinkage_ServiceDesc is the grpc.ServiceDesc for SceneLinkage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SceneLinkage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rule.SceneLinkage",
	HandlerType: (*SceneLinkageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "sceneInfoCreate",
			Handler:    _SceneLinkage_SceneInfoCreate_Handler,
		},
		{
			MethodName: "sceneInfoUpdate",
			Handler:    _SceneLinkage_SceneInfoUpdate_Handler,
		},
		{
			MethodName: "sceneInfoDelete",
			Handler:    _SceneLinkage_SceneInfoDelete_Handler,
		},
		{
			MethodName: "sceneInfoIndex",
			Handler:    _SceneLinkage_SceneInfoIndex_Handler,
		},
		{
			MethodName: "sceneInfoRead",
			Handler:    _SceneLinkage_SceneInfoRead_Handler,
		},
		{
			MethodName: "sceneManuallyTrigger",
			Handler:    _SceneLinkage_SceneManuallyTrigger_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/rule.proto",
}

// AlarmCenterClient is the client API for AlarmCenter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AlarmCenterClient interface {
	AlarmInfoCreate(ctx context.Context, in *AlarmInfo, opts ...grpc.CallOption) (*WithID, error)
	AlarmInfoUpdate(ctx context.Context, in *AlarmInfo, opts ...grpc.CallOption) (*Empty, error)
	AlarmInfoDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error)
	AlarmInfoIndex(ctx context.Context, in *AlarmInfoIndexReq, opts ...grpc.CallOption) (*AlarmInfoIndexResp, error)
	AlarmInfoRead(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*AlarmInfo, error)
	//告警关联场景联动
	AlarmSceneMultiUpdate(ctx context.Context, in *AlarmSceneMultiUpdateReq, opts ...grpc.CallOption) (*Empty, error)
	AlarmSceneDelete(ctx context.Context, in *AlarmSceneDeleteReq, opts ...grpc.CallOption) (*Empty, error)
	//告警记录
	AlarmRecordIndex(ctx context.Context, in *AlarmRecordIndexReq, opts ...grpc.CallOption) (*AlarmRecordIndexResp, error)
	//告警触发
	AlarmTrigger(ctx context.Context, in *AlarmTriggerReq, opts ...grpc.CallOption) (*WithID, error)
	//告警解除
	AlarmRelieve(ctx context.Context, in *AlarmRelieveReq, opts ...grpc.CallOption) (*WithID, error)
	//告警流水日志
	AlarmLogIndex(ctx context.Context, in *AlarmLogIndexReq, opts ...grpc.CallOption) (*AlarmLogIndexResp, error)
	//告警处理记录
	AlarmDealRecordCreate(ctx context.Context, in *AlarmDealRecordCreateReq, opts ...grpc.CallOption) (*WithID, error)
	AlarmDealRecordIndex(ctx context.Context, in *AlarmDealRecordIndexReq, opts ...grpc.CallOption) (*AlarmDealRecordIndexResp, error)
}

type alarmCenterClient struct {
	cc grpc.ClientConnInterface
}

func NewAlarmCenterClient(cc grpc.ClientConnInterface) AlarmCenterClient {
	return &alarmCenterClient{cc}
}

func (c *alarmCenterClient) AlarmInfoCreate(ctx context.Context, in *AlarmInfo, opts ...grpc.CallOption) (*WithID, error) {
	out := new(WithID)
	err := c.cc.Invoke(ctx, "/rule.alarmCenter/alarmInfoCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alarmCenterClient) AlarmInfoUpdate(ctx context.Context, in *AlarmInfo, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/rule.alarmCenter/alarmInfoUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alarmCenterClient) AlarmInfoDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/rule.alarmCenter/alarmInfoDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alarmCenterClient) AlarmInfoIndex(ctx context.Context, in *AlarmInfoIndexReq, opts ...grpc.CallOption) (*AlarmInfoIndexResp, error) {
	out := new(AlarmInfoIndexResp)
	err := c.cc.Invoke(ctx, "/rule.alarmCenter/alarmInfoIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alarmCenterClient) AlarmInfoRead(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*AlarmInfo, error) {
	out := new(AlarmInfo)
	err := c.cc.Invoke(ctx, "/rule.alarmCenter/alarmInfoRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alarmCenterClient) AlarmSceneMultiUpdate(ctx context.Context, in *AlarmSceneMultiUpdateReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/rule.alarmCenter/alarmSceneMultiUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alarmCenterClient) AlarmSceneDelete(ctx context.Context, in *AlarmSceneDeleteReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/rule.alarmCenter/alarmSceneDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alarmCenterClient) AlarmRecordIndex(ctx context.Context, in *AlarmRecordIndexReq, opts ...grpc.CallOption) (*AlarmRecordIndexResp, error) {
	out := new(AlarmRecordIndexResp)
	err := c.cc.Invoke(ctx, "/rule.alarmCenter/alarmRecordIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alarmCenterClient) AlarmTrigger(ctx context.Context, in *AlarmTriggerReq, opts ...grpc.CallOption) (*WithID, error) {
	out := new(WithID)
	err := c.cc.Invoke(ctx, "/rule.alarmCenter/alarmTrigger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alarmCenterClient) AlarmRelieve(ctx context.Context, in *AlarmRelieveReq, opts ...grpc.CallOption) (*WithID, error) {
	out := new(WithID)
	err := c.cc.Invoke(ctx, "/rule.alarmCenter/alarmRelieve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alarmCenterClient) AlarmLogIndex(ctx context.Context, in *AlarmLogIndexReq, opts ...grpc.CallOption) (*AlarmLogIndexResp, error) {
	out := new(AlarmLogIndexResp)
	err := c.cc.Invoke(ctx, "/rule.alarmCenter/alarmLogIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alarmCenterClient) AlarmDealRecordCreate(ctx context.Context, in *AlarmDealRecordCreateReq, opts ...grpc.CallOption) (*WithID, error) {
	out := new(WithID)
	err := c.cc.Invoke(ctx, "/rule.alarmCenter/alarmDealRecordCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alarmCenterClient) AlarmDealRecordIndex(ctx context.Context, in *AlarmDealRecordIndexReq, opts ...grpc.CallOption) (*AlarmDealRecordIndexResp, error) {
	out := new(AlarmDealRecordIndexResp)
	err := c.cc.Invoke(ctx, "/rule.alarmCenter/alarmDealRecordIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlarmCenterServer is the server API for AlarmCenter service.
// All implementations must embed UnimplementedAlarmCenterServer
// for forward compatibility
type AlarmCenterServer interface {
	AlarmInfoCreate(context.Context, *AlarmInfo) (*WithID, error)
	AlarmInfoUpdate(context.Context, *AlarmInfo) (*Empty, error)
	AlarmInfoDelete(context.Context, *WithID) (*Empty, error)
	AlarmInfoIndex(context.Context, *AlarmInfoIndexReq) (*AlarmInfoIndexResp, error)
	AlarmInfoRead(context.Context, *WithID) (*AlarmInfo, error)
	//告警关联场景联动
	AlarmSceneMultiUpdate(context.Context, *AlarmSceneMultiUpdateReq) (*Empty, error)
	AlarmSceneDelete(context.Context, *AlarmSceneDeleteReq) (*Empty, error)
	//告警记录
	AlarmRecordIndex(context.Context, *AlarmRecordIndexReq) (*AlarmRecordIndexResp, error)
	//告警触发
	AlarmTrigger(context.Context, *AlarmTriggerReq) (*WithID, error)
	//告警解除
	AlarmRelieve(context.Context, *AlarmRelieveReq) (*WithID, error)
	//告警流水日志
	AlarmLogIndex(context.Context, *AlarmLogIndexReq) (*AlarmLogIndexResp, error)
	//告警处理记录
	AlarmDealRecordCreate(context.Context, *AlarmDealRecordCreateReq) (*WithID, error)
	AlarmDealRecordIndex(context.Context, *AlarmDealRecordIndexReq) (*AlarmDealRecordIndexResp, error)
	mustEmbedUnimplementedAlarmCenterServer()
}

// UnimplementedAlarmCenterServer must be embedded to have forward compatible implementations.
type UnimplementedAlarmCenterServer struct {
}

func (UnimplementedAlarmCenterServer) AlarmInfoCreate(context.Context, *AlarmInfo) (*WithID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlarmInfoCreate not implemented")
}
func (UnimplementedAlarmCenterServer) AlarmInfoUpdate(context.Context, *AlarmInfo) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlarmInfoUpdate not implemented")
}
func (UnimplementedAlarmCenterServer) AlarmInfoDelete(context.Context, *WithID) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlarmInfoDelete not implemented")
}
func (UnimplementedAlarmCenterServer) AlarmInfoIndex(context.Context, *AlarmInfoIndexReq) (*AlarmInfoIndexResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlarmInfoIndex not implemented")
}
func (UnimplementedAlarmCenterServer) AlarmInfoRead(context.Context, *WithID) (*AlarmInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlarmInfoRead not implemented")
}
func (UnimplementedAlarmCenterServer) AlarmSceneMultiUpdate(context.Context, *AlarmSceneMultiUpdateReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlarmSceneMultiUpdate not implemented")
}
func (UnimplementedAlarmCenterServer) AlarmSceneDelete(context.Context, *AlarmSceneDeleteReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlarmSceneDelete not implemented")
}
func (UnimplementedAlarmCenterServer) AlarmRecordIndex(context.Context, *AlarmRecordIndexReq) (*AlarmRecordIndexResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlarmRecordIndex not implemented")
}
func (UnimplementedAlarmCenterServer) AlarmTrigger(context.Context, *AlarmTriggerReq) (*WithID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlarmTrigger not implemented")
}
func (UnimplementedAlarmCenterServer) AlarmRelieve(context.Context, *AlarmRelieveReq) (*WithID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlarmRelieve not implemented")
}
func (UnimplementedAlarmCenterServer) AlarmLogIndex(context.Context, *AlarmLogIndexReq) (*AlarmLogIndexResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlarmLogIndex not implemented")
}
func (UnimplementedAlarmCenterServer) AlarmDealRecordCreate(context.Context, *AlarmDealRecordCreateReq) (*WithID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlarmDealRecordCreate not implemented")
}
func (UnimplementedAlarmCenterServer) AlarmDealRecordIndex(context.Context, *AlarmDealRecordIndexReq) (*AlarmDealRecordIndexResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlarmDealRecordIndex not implemented")
}
func (UnimplementedAlarmCenterServer) mustEmbedUnimplementedAlarmCenterServer() {}

// UnsafeAlarmCenterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AlarmCenterServer will
// result in compilation errors.
type UnsafeAlarmCenterServer interface {
	mustEmbedUnimplementedAlarmCenterServer()
}

func RegisterAlarmCenterServer(s grpc.ServiceRegistrar, srv AlarmCenterServer) {
	s.RegisterService(&AlarmCenter_ServiceDesc, srv)
}

func _AlarmCenter_AlarmInfoCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlarmInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlarmCenterServer).AlarmInfoCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rule.alarmCenter/alarmInfoCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlarmCenterServer).AlarmInfoCreate(ctx, req.(*AlarmInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlarmCenter_AlarmInfoUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlarmInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlarmCenterServer).AlarmInfoUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rule.alarmCenter/alarmInfoUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlarmCenterServer).AlarmInfoUpdate(ctx, req.(*AlarmInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlarmCenter_AlarmInfoDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlarmCenterServer).AlarmInfoDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rule.alarmCenter/alarmInfoDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlarmCenterServer).AlarmInfoDelete(ctx, req.(*WithID))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlarmCenter_AlarmInfoIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlarmInfoIndexReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlarmCenterServer).AlarmInfoIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rule.alarmCenter/alarmInfoIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlarmCenterServer).AlarmInfoIndex(ctx, req.(*AlarmInfoIndexReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlarmCenter_AlarmInfoRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlarmCenterServer).AlarmInfoRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rule.alarmCenter/alarmInfoRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlarmCenterServer).AlarmInfoRead(ctx, req.(*WithID))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlarmCenter_AlarmSceneMultiUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlarmSceneMultiUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlarmCenterServer).AlarmSceneMultiUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rule.alarmCenter/alarmSceneMultiUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlarmCenterServer).AlarmSceneMultiUpdate(ctx, req.(*AlarmSceneMultiUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlarmCenter_AlarmSceneDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlarmSceneDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlarmCenterServer).AlarmSceneDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rule.alarmCenter/alarmSceneDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlarmCenterServer).AlarmSceneDelete(ctx, req.(*AlarmSceneDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlarmCenter_AlarmRecordIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlarmRecordIndexReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlarmCenterServer).AlarmRecordIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rule.alarmCenter/alarmRecordIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlarmCenterServer).AlarmRecordIndex(ctx, req.(*AlarmRecordIndexReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlarmCenter_AlarmTrigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlarmTriggerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlarmCenterServer).AlarmTrigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rule.alarmCenter/alarmTrigger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlarmCenterServer).AlarmTrigger(ctx, req.(*AlarmTriggerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlarmCenter_AlarmRelieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlarmRelieveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlarmCenterServer).AlarmRelieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rule.alarmCenter/alarmRelieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlarmCenterServer).AlarmRelieve(ctx, req.(*AlarmRelieveReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlarmCenter_AlarmLogIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlarmLogIndexReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlarmCenterServer).AlarmLogIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rule.alarmCenter/alarmLogIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlarmCenterServer).AlarmLogIndex(ctx, req.(*AlarmLogIndexReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlarmCenter_AlarmDealRecordCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlarmDealRecordCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlarmCenterServer).AlarmDealRecordCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rule.alarmCenter/alarmDealRecordCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlarmCenterServer).AlarmDealRecordCreate(ctx, req.(*AlarmDealRecordCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlarmCenter_AlarmDealRecordIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlarmDealRecordIndexReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlarmCenterServer).AlarmDealRecordIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rule.alarmCenter/alarmDealRecordIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlarmCenterServer).AlarmDealRecordIndex(ctx, req.(*AlarmDealRecordIndexReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AlarmCenter_ServiceDesc is the grpc.ServiceDesc for AlarmCenter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AlarmCenter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rule.alarmCenter",
	HandlerType: (*AlarmCenterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "alarmInfoCreate",
			Handler:    _AlarmCenter_AlarmInfoCreate_Handler,
		},
		{
			MethodName: "alarmInfoUpdate",
			Handler:    _AlarmCenter_AlarmInfoUpdate_Handler,
		},
		{
			MethodName: "alarmInfoDelete",
			Handler:    _AlarmCenter_AlarmInfoDelete_Handler,
		},
		{
			MethodName: "alarmInfoIndex",
			Handler:    _AlarmCenter_AlarmInfoIndex_Handler,
		},
		{
			MethodName: "alarmInfoRead",
			Handler:    _AlarmCenter_AlarmInfoRead_Handler,
		},
		{
			MethodName: "alarmSceneMultiUpdate",
			Handler:    _AlarmCenter_AlarmSceneMultiUpdate_Handler,
		},
		{
			MethodName: "alarmSceneDelete",
			Handler:    _AlarmCenter_AlarmSceneDelete_Handler,
		},
		{
			MethodName: "alarmRecordIndex",
			Handler:    _AlarmCenter_AlarmRecordIndex_Handler,
		},
		{
			MethodName: "alarmTrigger",
			Handler:    _AlarmCenter_AlarmTrigger_Handler,
		},
		{
			MethodName: "alarmRelieve",
			Handler:    _AlarmCenter_AlarmRelieve_Handler,
		},
		{
			MethodName: "alarmLogIndex",
			Handler:    _AlarmCenter_AlarmLogIndex_Handler,
		},
		{
			MethodName: "alarmDealRecordCreate",
			Handler:    _AlarmCenter_AlarmDealRecordCreate_Handler,
		},
		{
			MethodName: "alarmDealRecordIndex",
			Handler:    _AlarmCenter_AlarmDealRecordIndex_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/rule.proto",
}