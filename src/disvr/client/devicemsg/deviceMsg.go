// Code generated by goctl. DO NOT EDIT!
// Source: di.proto

package client

import (
	"context"

	"github.com/i-Things/things/src/disvr/internal/svc"
	"github.com/i-Things/things/src/disvr/pb/di"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	HubLogIndex          = di.HubLogIndex
	HubLogIndexReq       = di.HubLogIndexReq
	HubLogIndexResp      = di.HubLogIndexResp
	PageInfo             = di.PageInfo
	Response             = di.Response
	SchemaIndex          = di.SchemaIndex
	SchemaIndexResp      = di.SchemaIndexResp
	SchemaLatestIndexReq = di.SchemaLatestIndexReq
	SchemaLogIndexReq    = di.SchemaLogIndexReq
	SdkLogIndex          = di.SdkLogIndex
	SdkLogIndexReq       = di.SdkLogIndexReq
	SdkLogIndexResp      = di.SdkLogIndexResp
	SendActionReq        = di.SendActionReq
	SendActionResp       = di.SendActionResp
	SendMsgReq           = di.SendMsgReq
	SendMsgResp          = di.SendMsgResp
	SendPropertyReq      = di.SendPropertyReq
	SendPropertyResp     = di.SendPropertyResp

	DeviceMsg interface {
		// 获取设备sdk调试日志
		SdkLogIndex(ctx context.Context, in *SdkLogIndexReq, opts ...grpc.CallOption) (*SdkLogIndexResp, error)
		// 获取设备调试信息记录登入登出,操作
		HubLogIndex(ctx context.Context, in *HubLogIndexReq, opts ...grpc.CallOption) (*HubLogIndexResp, error)
		// 获取设备数据信息
		SchemaLatestIndex(ctx context.Context, in *SchemaLatestIndexReq, opts ...grpc.CallOption) (*SchemaIndexResp, error)
		// 获取设备数据信息
		SchemaLogIndex(ctx context.Context, in *SchemaLogIndexReq, opts ...grpc.CallOption) (*SchemaIndexResp, error)
	}

	defaultDeviceMsg struct {
		cli zrpc.Client
	}

	directDeviceMsg struct {
		svcCtx *svc.ServiceContext
		svr    di.DeviceMsgServer
	}
)

func NewDeviceMsg(cli zrpc.Client) DeviceMsg {
	return &defaultDeviceMsg{
		cli: cli,
	}
}

func NewDirectDeviceMsg(svcCtx *svc.ServiceContext, svr di.DeviceMsgServer) DeviceMsg {
	return &directDeviceMsg{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

// 获取设备sdk调试日志
func (m *defaultDeviceMsg) SdkLogIndex(ctx context.Context, in *SdkLogIndexReq, opts ...grpc.CallOption) (*SdkLogIndexResp, error) {
	client := di.NewDeviceMsgClient(m.cli.Conn())
	return client.SdkLogIndex(ctx, in, opts...)
}

// 获取设备sdk调试日志
func (d *directDeviceMsg) SdkLogIndex(ctx context.Context, in *SdkLogIndexReq, opts ...grpc.CallOption) (*SdkLogIndexResp, error) {
	return d.svr.SdkLogIndex(ctx, in)
}

// 获取设备调试信息记录登入登出,操作
func (m *defaultDeviceMsg) HubLogIndex(ctx context.Context, in *HubLogIndexReq, opts ...grpc.CallOption) (*HubLogIndexResp, error) {
	client := di.NewDeviceMsgClient(m.cli.Conn())
	return client.HubLogIndex(ctx, in, opts...)
}

// 获取设备调试信息记录登入登出,操作
func (d *directDeviceMsg) HubLogIndex(ctx context.Context, in *HubLogIndexReq, opts ...grpc.CallOption) (*HubLogIndexResp, error) {
	return d.svr.HubLogIndex(ctx, in)
}

// 获取设备数据信息
func (m *defaultDeviceMsg) SchemaLatestIndex(ctx context.Context, in *SchemaLatestIndexReq, opts ...grpc.CallOption) (*SchemaIndexResp, error) {
	client := di.NewDeviceMsgClient(m.cli.Conn())
	return client.SchemaLatestIndex(ctx, in, opts...)
}

// 获取设备数据信息
func (d *directDeviceMsg) SchemaLatestIndex(ctx context.Context, in *SchemaLatestIndexReq, opts ...grpc.CallOption) (*SchemaIndexResp, error) {
	return d.svr.SchemaLatestIndex(ctx, in)
}

// 获取设备数据信息
func (m *defaultDeviceMsg) SchemaLogIndex(ctx context.Context, in *SchemaLogIndexReq, opts ...grpc.CallOption) (*SchemaIndexResp, error) {
	client := di.NewDeviceMsgClient(m.cli.Conn())
	return client.SchemaLogIndex(ctx, in, opts...)
}

// 获取设备数据信息
func (d *directDeviceMsg) SchemaLogIndex(ctx context.Context, in *SchemaLogIndexReq, opts ...grpc.CallOption) (*SchemaIndexResp, error) {
	return d.svr.SchemaLogIndex(ctx, in)
}
