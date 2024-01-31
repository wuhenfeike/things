package otafirmwaremanagelogic

import (
	"context"
	"gitee.com/i-Things/share/utils"
	"github.com/i-Things/things/service/dmsvr/internal/repo/relationDB"
	"github.com/jinzhu/copier"

	"github.com/i-Things/things/service/dmsvr/internal/svc"
	"github.com/i-Things/things/service/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/core/logx"
)

type OtaFirmwareReadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	PiDB  *relationDB.ProductInfoRepo
	OfDB  *relationDB.OtaFirmwareRepo
	OffDB *relationDB.OtaFirmwareFileRepo
}

func NewOtaFirmwareReadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OtaFirmwareReadLogic {
	return &OtaFirmwareReadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		PiDB:   relationDB.NewProductInfoRepo(ctx),
		OfDB:   relationDB.NewOtaFirmwareRepo(ctx),
		OffDB:  relationDB.NewOtaFirmwareFileRepo(ctx),
	}
}

// 查询升级包
func (l *OtaFirmwareReadLogic) OtaFirmwareRead(in *dm.OtaFirmwareReadReq) (*dm.OtaFirmwareReadResp, error) {
	otaFirmware, err := l.OfDB.FindOneByFilter(l.ctx, relationDB.OtaFirmwareFilter{FirmwareID: in.FirmwareId})
	if err != nil {
		l.Errorf("%s.Query OTAFirmware err=%v", utils.FuncName(), err)
		return nil, err
	}
	otaFirmwareList, err := l.OffDB.FindByFilter(l.ctx, relationDB.OtaFirmwareFileFilter{FirmwareID: in.FirmwareId}, nil)
	if err != nil {
		l.Errorf("%s.Query OTAFirmwareFile err=%v", utils.FuncName(), err)
		return nil, err
	}
	var result dm.OtaFirmwareReadResp
	err = copier.Copy(&result, &otaFirmware)
	result.FirmwareId = otaFirmware.ID
	result.FirmwareName = otaFirmware.Name
	result.DestVersion = otaFirmware.Version
	result.FirmwareDesc = otaFirmware.Desc
	result.FirmwareUdi = otaFirmware.Extra
	if err != nil {
		l.Errorf("%s.Copy OTAFirmware err=%v", utils.FuncName(), err)
		return nil, err
	}
	err = copier.Copy(&result.FirmwareFileList, &otaFirmwareList)
	if err != nil {
		l.Errorf("%s.Copy OTAFirmwareFile err=%v", utils.FuncName(), err)
		return nil, err
	}
	logx.Infof("result:%+v", &result)
	return &result, nil
}
