package userdevicelogic

import (
	"context"
	"gitee.com/i-Things/core/service/syssvr/pb/sys"
	"gitee.com/i-Things/share/ctxs"
	"gitee.com/i-Things/share/errors"
	"github.com/i-Things/things/service/dmsvr/internal/logic"
	"github.com/i-Things/things/service/dmsvr/internal/repo/relationDB"

	"github.com/i-Things/things/service/dmsvr/internal/svc"
	"github.com/i-Things/things/service/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDeviceShareIndexLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserDeviceShareIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDeviceShareIndexLogic {
	return &UserDeviceShareIndexLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取设备分享列表(只有)
func (l *UserDeviceShareIndexLogic) UserDeviceShareIndex(in *dm.UserDeviceShareIndexReq) (*dm.UserDeviceShareIndexResp, error) {
	uc := ctxs.GetUserCtx(l.ctx)
	di, err := relationDB.NewDeviceInfoRepo(l.ctx).FindOneByFilter(l.ctx, relationDB.DeviceFilter{ProductID: in.Device.ProductID, DeviceName: in.Device.DeviceName})
	if err != nil {
		return nil, err
	}
	pi, err := l.svcCtx.ProjectM.ProjectInfoRead(l.ctx, &sys.ProjectWithID{ProjectID: int64(di.ProjectID)})
	if err != nil {
		return nil, err
	}
	if pi.AdminUserID != uc.UserID { //只有所有者和被分享者才有权限操作
		return nil, errors.Permissions
	}
	f := relationDB.UserDeviceShareFilter{ProductID: di.ProductID, DeviceName: di.DeviceName}
	total, err := relationDB.NewUserDeviceShareRepo(l.ctx).CountByFilter(l.ctx, f)
	if err != nil {
		return nil, err
	}
	list, err := relationDB.NewUserDeviceShareRepo(l.ctx).FindByFilter(l.ctx, f, logic.ToPageInfo(in.Page))
	if err != nil {
		return nil, err
	}
	return &dm.UserDeviceShareIndexResp{Total: total, List: ToUserDeviceSharePbs(list)}, nil
}