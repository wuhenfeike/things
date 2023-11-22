package indexapi

import (
	"context"
	"encoding/json"

	"github.com/i-Things/things/src/apisvr/internal/svc"
	"github.com/i-Things/things/src/apisvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StopSendRtpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStopSendRtpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StopSendRtpLogic {
	return &StopSendRtpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StopSendRtpLogic) StopSendRtp(req *types.IndexApiReq) (resp *types.IndexApiStopSendRtpResp, err error) {
	// todo: add your logic here and delete this line
	data, err := proxyMediaServer(l.ctx, l.svcCtx, STOPSENDRTP, req.VidmgrID)
	dataRecv := new(types.IndexApiStopSendRtpResp)
	json.Unmarshal(data, dataRecv)
	return dataRecv, err
}