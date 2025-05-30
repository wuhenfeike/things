// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.1
// Source: dm.proto

package server

import (
	"context"

	"gitee.com/unitedrhino/things/service/dmsvr/internal/logic/devicegroup"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/svc"
	"gitee.com/unitedrhino/things/service/dmsvr/pb/dm"
)

type DeviceGroupServer struct {
	svcCtx *svc.ServiceContext
	dm.UnimplementedDeviceGroupServer
}

func NewDeviceGroupServer(svcCtx *svc.ServiceContext) *DeviceGroupServer {
	return &DeviceGroupServer{
		svcCtx: svcCtx,
	}
}

// 创建分组
func (s *DeviceGroupServer) GroupInfoCreate(ctx context.Context, in *dm.GroupInfo) (*dm.WithID, error) {
	l := devicegrouplogic.NewGroupInfoCreateLogic(ctx, s.svcCtx)
	return l.GroupInfoCreate(in)
}

func (s *DeviceGroupServer) GroupInfoMultiCreate(ctx context.Context, in *dm.GroupInfoMultiCreateReq) (*dm.Empty, error) {
	l := devicegrouplogic.NewGroupInfoMultiCreateLogic(ctx, s.svcCtx)
	return l.GroupInfoMultiCreate(in)
}

// 获取分组信息列表
func (s *DeviceGroupServer) GroupInfoIndex(ctx context.Context, in *dm.GroupInfoIndexReq) (*dm.GroupInfoIndexResp, error) {
	l := devicegrouplogic.NewGroupInfoIndexLogic(ctx, s.svcCtx)
	return l.GroupInfoIndex(in)
}

// 获取分组信息详情
func (s *DeviceGroupServer) GroupInfoRead(ctx context.Context, in *dm.GroupInfoReadReq) (*dm.GroupInfo, error) {
	l := devicegrouplogic.NewGroupInfoReadLogic(ctx, s.svcCtx)
	return l.GroupInfoRead(in)
}

// 更新分组
func (s *DeviceGroupServer) GroupInfoUpdate(ctx context.Context, in *dm.GroupInfo) (*dm.Empty, error) {
	l := devicegrouplogic.NewGroupInfoUpdateLogic(ctx, s.svcCtx)
	return l.GroupInfoUpdate(in)
}

// 删除分组
func (s *DeviceGroupServer) GroupInfoDelete(ctx context.Context, in *dm.WithID) (*dm.Empty, error) {
	l := devicegrouplogic.NewGroupInfoDeleteLogic(ctx, s.svcCtx)
	return l.GroupInfoDelete(in)
}

// 创建分组设备
func (s *DeviceGroupServer) GroupDeviceMultiCreate(ctx context.Context, in *dm.GroupDeviceMultiSaveReq) (*dm.Empty, error) {
	l := devicegrouplogic.NewGroupDeviceMultiCreateLogic(ctx, s.svcCtx)
	return l.GroupDeviceMultiCreate(in)
}

// 更新分组设备
func (s *DeviceGroupServer) GroupDeviceMultiUpdate(ctx context.Context, in *dm.GroupDeviceMultiSaveReq) (*dm.Empty, error) {
	l := devicegrouplogic.NewGroupDeviceMultiUpdateLogic(ctx, s.svcCtx)
	return l.GroupDeviceMultiUpdate(in)
}

// 删除分组设备
func (s *DeviceGroupServer) GroupDeviceMultiDelete(ctx context.Context, in *dm.GroupDeviceMultiDeleteReq) (*dm.Empty, error) {
	l := devicegrouplogic.NewGroupDeviceMultiDeleteLogic(ctx, s.svcCtx)
	return l.GroupDeviceMultiDelete(in)
}
