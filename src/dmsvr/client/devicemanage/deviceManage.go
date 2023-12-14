// Code generated by goctl. DO NOT EDIT.
// Source: dm.proto

package devicemanage

import (
	"context"

	"github.com/i-Things/things/src/dmsvr/internal/svc"
	"github.com/i-Things/things/src/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CustomTopic                 = dm.CustomTopic
	DeleteOtaFirmwareReq        = dm.DeleteOtaFirmwareReq
	DeviceCore                  = dm.DeviceCore
	DeviceGatewayBindDevice     = dm.DeviceGatewayBindDevice
	DeviceGatewayIndexReq       = dm.DeviceGatewayIndexReq
	DeviceGatewayIndexResp      = dm.DeviceGatewayIndexResp
	DeviceGatewayMultiCreateReq = dm.DeviceGatewayMultiCreateReq
	DeviceGatewayMultiDeleteReq = dm.DeviceGatewayMultiDeleteReq
	DeviceGatewaySign           = dm.DeviceGatewaySign
	DeviceInfo                  = dm.DeviceInfo
	DeviceInfoCountReq          = dm.DeviceInfoCountReq
	DeviceInfoCountResp         = dm.DeviceInfoCountResp
	DeviceInfoDeleteReq         = dm.DeviceInfoDeleteReq
	DeviceInfoIndexReq          = dm.DeviceInfoIndexReq
	DeviceInfoIndexResp         = dm.DeviceInfoIndexResp
	DeviceInfoReadReq           = dm.DeviceInfoReadReq
	DeviceTypeCountReq          = dm.DeviceTypeCountReq
	DeviceTypeCountResp         = dm.DeviceTypeCountResp
	EventIndex                  = dm.EventIndex
	EventIndexResp              = dm.EventIndexResp
	EventLogIndexReq            = dm.EventLogIndexReq
	Firmware                    = dm.Firmware
	FirmwareFile                = dm.FirmwareFile
	FirmwareInfo                = dm.FirmwareInfo
	FirmwareInfoDeleteReq       = dm.FirmwareInfoDeleteReq
	FirmwareInfoDeleteResp      = dm.FirmwareInfoDeleteResp
	FirmwareInfoIndexReq        = dm.FirmwareInfoIndexReq
	FirmwareInfoIndexResp       = dm.FirmwareInfoIndexResp
	FirmwareInfoReadReq         = dm.FirmwareInfoReadReq
	FirmwareInfoReadResp        = dm.FirmwareInfoReadResp
	FirmwareResp                = dm.FirmwareResp
	GetPropertyReplyReq         = dm.GetPropertyReplyReq
	GetPropertyReplyResp        = dm.GetPropertyReplyResp
	GroupDeviceIndexReq         = dm.GroupDeviceIndexReq
	GroupDeviceIndexResp        = dm.GroupDeviceIndexResp
	GroupDeviceMultiCreateReq   = dm.GroupDeviceMultiCreateReq
	GroupDeviceMultiDeleteReq   = dm.GroupDeviceMultiDeleteReq
	GroupInfo                   = dm.GroupInfo
	GroupInfoCreateReq          = dm.GroupInfoCreateReq
	GroupInfoDeleteReq          = dm.GroupInfoDeleteReq
	GroupInfoIndexReq           = dm.GroupInfoIndexReq
	GroupInfoIndexResp          = dm.GroupInfoIndexResp
	GroupInfoReadReq            = dm.GroupInfoReadReq
	GroupInfoUpdateReq          = dm.GroupInfoUpdateReq
	HubLogIndex                 = dm.HubLogIndex
	HubLogIndexReq              = dm.HubLogIndexReq
	HubLogIndexResp             = dm.HubLogIndexResp
	ListOtaFirmwareReq          = dm.ListOtaFirmwareReq
	ListOtaFirmwareResp         = dm.ListOtaFirmwareResp
	ModifyOtaFirmwareReq        = dm.ModifyOtaFirmwareReq
	MultiSendPropertyReq        = dm.MultiSendPropertyReq
	MultiSendPropertyResp       = dm.MultiSendPropertyResp
	OtaCommonResp               = dm.OtaCommonResp
	OtaFirmwareDeviceInfoReq    = dm.OtaFirmwareDeviceInfoReq
	OtaFirmwareDeviceInfoResp   = dm.OtaFirmwareDeviceInfoResp
	OtaFirmwareFile             = dm.OtaFirmwareFile
	OtaFirmwareFileIndexReq     = dm.OtaFirmwareFileIndexReq
	OtaFirmwareFileIndexResp    = dm.OtaFirmwareFileIndexResp
	OtaFirmwareFileInfo         = dm.OtaFirmwareFileInfo
	OtaFirmwareFileReq          = dm.OtaFirmwareFileReq
	OtaFirmwareFileResp         = dm.OtaFirmwareFileResp
	OtaFirmwareInfo             = dm.OtaFirmwareInfo
	OtaFirmwareReq              = dm.OtaFirmwareReq
	OtaFirmwareResp             = dm.OtaFirmwareResp
	OtaPageInfo                 = dm.OtaPageInfo
	OtaPromptIndexReq           = dm.OtaPromptIndexReq
	OtaPromptIndexResp          = dm.OtaPromptIndexResp
	OtaTaskBatchReq             = dm.OtaTaskBatchReq
	OtaTaskBatchResp            = dm.OtaTaskBatchResp
	OtaTaskCancleReq            = dm.OtaTaskCancleReq
	OtaTaskCreatResp            = dm.OtaTaskCreatResp
	OtaTaskCreateReq            = dm.OtaTaskCreateReq
	OtaTaskDeviceCancleReq      = dm.OtaTaskDeviceCancleReq
	OtaTaskDeviceIndexReq       = dm.OtaTaskDeviceIndexReq
	OtaTaskDeviceIndexResp      = dm.OtaTaskDeviceIndexResp
	OtaTaskDeviceInfo           = dm.OtaTaskDeviceInfo
	OtaTaskDeviceProcessReq     = dm.OtaTaskDeviceProcessReq
	OtaTaskDeviceReadReq        = dm.OtaTaskDeviceReadReq
	OtaTaskIndexReq             = dm.OtaTaskIndexReq
	OtaTaskIndexResp            = dm.OtaTaskIndexResp
	OtaTaskInfo                 = dm.OtaTaskInfo
	OtaTaskReadReq              = dm.OtaTaskReadReq
	OtaTaskReadResp             = dm.OtaTaskReadResp
	PageInfo                    = dm.PageInfo
	PageInfo_OrderBy            = dm.PageInfo_OrderBy
	Point                       = dm.Point
	ProductCustom               = dm.ProductCustom
	ProductCustomReadReq        = dm.ProductCustomReadReq
	ProductInfo                 = dm.ProductInfo
	ProductInfoDeleteReq        = dm.ProductInfoDeleteReq
	ProductInfoIndexReq         = dm.ProductInfoIndexReq
	ProductInfoIndexResp        = dm.ProductInfoIndexResp
	ProductInfoReadReq          = dm.ProductInfoReadReq
	ProductRemoteConfig         = dm.ProductRemoteConfig
	ProductSchemaCreateReq      = dm.ProductSchemaCreateReq
	ProductSchemaDeleteReq      = dm.ProductSchemaDeleteReq
	ProductSchemaIndexReq       = dm.ProductSchemaIndexReq
	ProductSchemaIndexResp      = dm.ProductSchemaIndexResp
	ProductSchemaInfo           = dm.ProductSchemaInfo
	ProductSchemaTslImportReq   = dm.ProductSchemaTslImportReq
	ProductSchemaTslReadReq     = dm.ProductSchemaTslReadReq
	ProductSchemaTslReadResp    = dm.ProductSchemaTslReadResp
	ProductSchemaUpdateReq      = dm.ProductSchemaUpdateReq
	PropertyIndex               = dm.PropertyIndex
	PropertyIndexResp           = dm.PropertyIndexResp
	PropertyLatestIndexReq      = dm.PropertyLatestIndexReq
	PropertyLogIndexReq         = dm.PropertyLogIndexReq
	ProtocolInfo                = dm.ProtocolInfo
	ProtocolInfoIndexReq        = dm.ProtocolInfoIndexReq
	ProtocolInfoIndexResp       = dm.ProtocolInfoIndexResp
	QueryOtaFirmwareReq         = dm.QueryOtaFirmwareReq
	QueryOtaFirmwareResp        = dm.QueryOtaFirmwareResp
	RemoteConfigCreateReq       = dm.RemoteConfigCreateReq
	RemoteConfigIndexReq        = dm.RemoteConfigIndexReq
	RemoteConfigIndexResp       = dm.RemoteConfigIndexResp
	RemoteConfigLastReadReq     = dm.RemoteConfigLastReadReq
	RemoteConfigLastReadResp    = dm.RemoteConfigLastReadResp
	RemoteConfigPushAllReq      = dm.RemoteConfigPushAllReq
	RespActionReq               = dm.RespActionReq
	RespReadReq                 = dm.RespReadReq
	Response                    = dm.Response
	RootCheckReq                = dm.RootCheckReq
	SdkLogIndex                 = dm.SdkLogIndex
	SdkLogIndexReq              = dm.SdkLogIndexReq
	SdkLogIndexResp             = dm.SdkLogIndexResp
	SendActionReq               = dm.SendActionReq
	SendActionResp              = dm.SendActionResp
	SendMsgReq                  = dm.SendMsgReq
	SendMsgResp                 = dm.SendMsgResp
	SendOption                  = dm.SendOption
	SendPropertyMsg             = dm.SendPropertyMsg
	SendPropertyReq             = dm.SendPropertyReq
	SendPropertyResp            = dm.SendPropertyResp
	ShadowIndex                 = dm.ShadowIndex
	ShadowIndexResp             = dm.ShadowIndexResp
	Tag                         = dm.Tag
	VerifyOtaFirmwareReq        = dm.VerifyOtaFirmwareReq

	DeviceManage interface {
		// 鉴定是否是root账号(提供给mqtt broker)
		RootCheck(ctx context.Context, in *RootCheckReq, opts ...grpc.CallOption) (*Response, error)
		// 新增设备
		DeviceInfoCreate(ctx context.Context, in *DeviceInfo, opts ...grpc.CallOption) (*Response, error)
		// 更新设备
		DeviceInfoUpdate(ctx context.Context, in *DeviceInfo, opts ...grpc.CallOption) (*Response, error)
		// 删除设备
		DeviceInfoDelete(ctx context.Context, in *DeviceInfoDeleteReq, opts ...grpc.CallOption) (*Response, error)
		// 获取设备信息列表
		DeviceInfoIndex(ctx context.Context, in *DeviceInfoIndexReq, opts ...grpc.CallOption) (*DeviceInfoIndexResp, error)
		// 获取设备信息详情
		DeviceInfoRead(ctx context.Context, in *DeviceInfoReadReq, opts ...grpc.CallOption) (*DeviceInfo, error)
		// 绑定网关下子设备设备
		DeviceGatewayMultiCreate(ctx context.Context, in *DeviceGatewayMultiCreateReq, opts ...grpc.CallOption) (*Response, error)
		// 获取绑定信息的设备信息列表
		DeviceGatewayIndex(ctx context.Context, in *DeviceGatewayIndexReq, opts ...grpc.CallOption) (*DeviceGatewayIndexResp, error)
		// 删除网关下子设备
		DeviceGatewayMultiDelete(ctx context.Context, in *DeviceGatewayMultiDeleteReq, opts ...grpc.CallOption) (*Response, error)
		// 设备计数
		DeviceInfoCount(ctx context.Context, in *DeviceInfoCountReq, opts ...grpc.CallOption) (*DeviceInfoCountResp, error)
		// 设备类型
		DeviceTypeCount(ctx context.Context, in *DeviceTypeCountReq, opts ...grpc.CallOption) (*DeviceTypeCountResp, error)
	}

	defaultDeviceManage struct {
		cli zrpc.Client
	}

	directDeviceManage struct {
		svcCtx *svc.ServiceContext
		svr    dm.DeviceManageServer
	}
)

func NewDeviceManage(cli zrpc.Client) DeviceManage {
	return &defaultDeviceManage{
		cli: cli,
	}
}

func NewDirectDeviceManage(svcCtx *svc.ServiceContext, svr dm.DeviceManageServer) DeviceManage {
	return &directDeviceManage{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

// 鉴定是否是root账号(提供给mqtt broker)
func (m *defaultDeviceManage) RootCheck(ctx context.Context, in *RootCheckReq, opts ...grpc.CallOption) (*Response, error) {
	client := dm.NewDeviceManageClient(m.cli.Conn())
	return client.RootCheck(ctx, in, opts...)
}

// 鉴定是否是root账号(提供给mqtt broker)
func (d *directDeviceManage) RootCheck(ctx context.Context, in *RootCheckReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.RootCheck(ctx, in)
}

// 新增设备
func (m *defaultDeviceManage) DeviceInfoCreate(ctx context.Context, in *DeviceInfo, opts ...grpc.CallOption) (*Response, error) {
	client := dm.NewDeviceManageClient(m.cli.Conn())
	return client.DeviceInfoCreate(ctx, in, opts...)
}

// 新增设备
func (d *directDeviceManage) DeviceInfoCreate(ctx context.Context, in *DeviceInfo, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.DeviceInfoCreate(ctx, in)
}

// 更新设备
func (m *defaultDeviceManage) DeviceInfoUpdate(ctx context.Context, in *DeviceInfo, opts ...grpc.CallOption) (*Response, error) {
	client := dm.NewDeviceManageClient(m.cli.Conn())
	return client.DeviceInfoUpdate(ctx, in, opts...)
}

// 更新设备
func (d *directDeviceManage) DeviceInfoUpdate(ctx context.Context, in *DeviceInfo, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.DeviceInfoUpdate(ctx, in)
}

// 删除设备
func (m *defaultDeviceManage) DeviceInfoDelete(ctx context.Context, in *DeviceInfoDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	client := dm.NewDeviceManageClient(m.cli.Conn())
	return client.DeviceInfoDelete(ctx, in, opts...)
}

// 删除设备
func (d *directDeviceManage) DeviceInfoDelete(ctx context.Context, in *DeviceInfoDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.DeviceInfoDelete(ctx, in)
}

// 获取设备信息列表
func (m *defaultDeviceManage) DeviceInfoIndex(ctx context.Context, in *DeviceInfoIndexReq, opts ...grpc.CallOption) (*DeviceInfoIndexResp, error) {
	client := dm.NewDeviceManageClient(m.cli.Conn())
	return client.DeviceInfoIndex(ctx, in, opts...)
}

// 获取设备信息列表
func (d *directDeviceManage) DeviceInfoIndex(ctx context.Context, in *DeviceInfoIndexReq, opts ...grpc.CallOption) (*DeviceInfoIndexResp, error) {
	return d.svr.DeviceInfoIndex(ctx, in)
}

// 获取设备信息详情
func (m *defaultDeviceManage) DeviceInfoRead(ctx context.Context, in *DeviceInfoReadReq, opts ...grpc.CallOption) (*DeviceInfo, error) {
	client := dm.NewDeviceManageClient(m.cli.Conn())
	return client.DeviceInfoRead(ctx, in, opts...)
}

// 获取设备信息详情
func (d *directDeviceManage) DeviceInfoRead(ctx context.Context, in *DeviceInfoReadReq, opts ...grpc.CallOption) (*DeviceInfo, error) {
	return d.svr.DeviceInfoRead(ctx, in)
}

// 绑定网关下子设备设备
func (m *defaultDeviceManage) DeviceGatewayMultiCreate(ctx context.Context, in *DeviceGatewayMultiCreateReq, opts ...grpc.CallOption) (*Response, error) {
	client := dm.NewDeviceManageClient(m.cli.Conn())
	return client.DeviceGatewayMultiCreate(ctx, in, opts...)
}

// 绑定网关下子设备设备
func (d *directDeviceManage) DeviceGatewayMultiCreate(ctx context.Context, in *DeviceGatewayMultiCreateReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.DeviceGatewayMultiCreate(ctx, in)
}

// 获取绑定信息的设备信息列表
func (m *defaultDeviceManage) DeviceGatewayIndex(ctx context.Context, in *DeviceGatewayIndexReq, opts ...grpc.CallOption) (*DeviceGatewayIndexResp, error) {
	client := dm.NewDeviceManageClient(m.cli.Conn())
	return client.DeviceGatewayIndex(ctx, in, opts...)
}

// 获取绑定信息的设备信息列表
func (d *directDeviceManage) DeviceGatewayIndex(ctx context.Context, in *DeviceGatewayIndexReq, opts ...grpc.CallOption) (*DeviceGatewayIndexResp, error) {
	return d.svr.DeviceGatewayIndex(ctx, in)
}

// 删除网关下子设备
func (m *defaultDeviceManage) DeviceGatewayMultiDelete(ctx context.Context, in *DeviceGatewayMultiDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	client := dm.NewDeviceManageClient(m.cli.Conn())
	return client.DeviceGatewayMultiDelete(ctx, in, opts...)
}

// 删除网关下子设备
func (d *directDeviceManage) DeviceGatewayMultiDelete(ctx context.Context, in *DeviceGatewayMultiDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.DeviceGatewayMultiDelete(ctx, in)
}

// 设备计数
func (m *defaultDeviceManage) DeviceInfoCount(ctx context.Context, in *DeviceInfoCountReq, opts ...grpc.CallOption) (*DeviceInfoCountResp, error) {
	client := dm.NewDeviceManageClient(m.cli.Conn())
	return client.DeviceInfoCount(ctx, in, opts...)
}

// 设备计数
func (d *directDeviceManage) DeviceInfoCount(ctx context.Context, in *DeviceInfoCountReq, opts ...grpc.CallOption) (*DeviceInfoCountResp, error) {
	return d.svr.DeviceInfoCount(ctx, in)
}

// 设备类型
func (m *defaultDeviceManage) DeviceTypeCount(ctx context.Context, in *DeviceTypeCountReq, opts ...grpc.CallOption) (*DeviceTypeCountResp, error) {
	client := dm.NewDeviceManageClient(m.cli.Conn())
	return client.DeviceTypeCount(ctx, in, opts...)
}

// 设备类型
func (d *directDeviceManage) DeviceTypeCount(ctx context.Context, in *DeviceTypeCountReq, opts ...grpc.CallOption) (*DeviceTypeCountResp, error) {
	return d.svr.DeviceTypeCount(ctx, in)
}
