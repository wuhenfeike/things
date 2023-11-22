package info

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/i-Things/things/shared/def"
	"github.com/i-Things/things/shared/errors"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/apisvr/internal/logic/things/vidmgr/indexapi"
	"github.com/i-Things/things/src/vidsvr/pb/vid"
	"time"

	"github.com/i-Things/things/src/apisvr/internal/svc"
	"github.com/i-Things/things/src/apisvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActiveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewActiveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActiveLogic {
	return &ActiveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}


func (l *ActiveLogic) Active(req *types.VidmgrInfoActiveReq) error {
	// todo: add your logic here and delete this line
	//read VidmgrInfo table and update table
	vidTmp, err := l.svcCtx.VidmgrM.VidmgrInfoRead(l.ctx, &vid.VidmgrInfoReadReq{
		VidmgrtID: req.VidmgrID,
	})
	if err != nil {
		er := errors.Fmt(err)
		l.Errorf("%s.rpc.ManageVidmgr req=%v err=%v", utils.FuncName(), req, er)
		return er
	}
	//STEP1  需要获取流服务的配置
	getServerConfig := indexapi.NewGetServerConfigLogic(l.ctx, l.svcCtx)
	setServerConfig := indexapi.NewSetServerConfigLogic(l.ctx, l.svcCtx)
	type IndexApiReq struct {
		VidmgrID string `json:"vidmgrID"`
		Data     string `json:"data"`
	}
	dataConfig, error := getServerConfig.GetServerConfig(&types.IndexApiReq{
		VidmgrID: req.VidmgrID,
		Data:     "",
	})
	if error != nil {
		fmt.Println("Server Activer failed")
		return error
	}
	fmt.Println("dataConfig:", dataConfig)
	//STEP2 修改流媒体服务  //set default
	if len(dataConfig.Data) > 0 {
		indexapi.SetDefaultConfig(l.svcCtx, &dataConfig.Data[0])
		dataConfig.Data[0].GeneralMediaServerId = req.VidmgrID
	}
	byteConfig, _ := json.Marshal(dataConfig.Data[0])
	//STEP3 配置流服务
	configReq := &types.IndexApiSetServerConfigReq{
		VidmgrID: req.VidmgrID,
		Data:     string(byteConfig),
	}
	setDataConfig, error := setServerConfig.SetServerConfig(configReq)
	if setDataConfig.Code != 0 {
		fmt.Println("setServerConfig  配置流服务出错")
		return error
	}
	//STEP3 insert配置
	_, err = l.svcCtx.VidmgrC.VidmgrConfigCreate(l.ctx, ToVidmgrConfigRpc(&dataConfig.Data[0]))
	if err != nil {
		er := errors.Fmt(err)
		l.Errorf("%s.rpc.ManageVidmgr req=%v err=%v", utils.FuncName(), req, er)
		return er
	}
	//STEP4 更新状态
	if vidTmp.VidmgrStatus == def.DeviceStatusInactive {
		//UPDATE
		vidReq := &vid.VidmgrInfo{
			VidmgrID:     vidTmp.VidmgrID,
			VidmgrStatus: def.DeviceStatusOnline,
			FirstLogin:   time.Now().Unix(),
			LastLogin:    time.Now().Unix(),
		}
		vidTmp.VidmgrStatus = def.DeviceStatusOnline

		_, err := l.svcCtx.VidmgrM.VidmgrInfoUpdate(l.ctx, vidReq)
		if err != nil {
			er := errors.Fmt(err)
			l.Errorf("%s.rpc.ManageVidmgr req=%v err=%v", utils.FuncName(), req, er)
			return er
		}
	}
	return nil
}