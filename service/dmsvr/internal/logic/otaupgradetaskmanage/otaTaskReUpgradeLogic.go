package otaupgradetaskmanagelogic

import (
	"context"
	"github.com/i-Things/things/service/dmsvr/internal/repo/relationDB"

	"github.com/i-Things/things/service/dmsvr/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type OtaTaskReUpgradeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	OtDB *relationDB.OtaFirmwareDeviceRepo
}

func NewOtaTaskReUpgradeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OtaTaskReUpgradeLogic {
	return &OtaTaskReUpgradeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		OtDB:   relationDB.NewOtaFirmwareDeviceRepo(ctx),
	}
}

//
//// 重新升级指定批次下升级失败或升级取消的设备升级作业
//func (l *OtaTaskReUpgradeLogic) OtaTaskReUpgrade(in *dm.OTATaskReUpgradeReq) (*dm.Empty, error) {
//	taskStatusList := []int{msgOta.UpgradeStatusCanceled, msgOta.UpgradeStatusFailed}
//	filter := relationDB.OtaFirmwareDeviceFilter{
//		Statues: taskStatusList,
//		JobID:          in.JobID,
//		IDs:            in.TaskIDs,
//	}
//	updateData := make(map[string]interface{})
//	updateData["task_status"] = msgOta.UpgradeStatusQueued
//	err := l.OtDB.BatchUpdateField(l.ctx, filter, updateData)
//	if err != nil {
//		l.Errorf("%s.TaskInfo.TaskInfo BatchUpdate failure err=%+v", utils.FuncName(), err)
//		return nil, err
//	}
//	return &dm.Empty{}, nil
//}
