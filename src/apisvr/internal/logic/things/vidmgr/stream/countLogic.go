package stream

import (
	"context"
	"github.com/i-Things/things/shared/errors"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/vidsvr/pb/vid"

	"github.com/i-Things/things/src/apisvr/internal/svc"
	"github.com/i-Things/things/src/apisvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CountLogic {
	return &CountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CountLogic) Count(req *types.VidmgrStreamCountReq) (resp *types.VidmgrStreamCountResp, err error) {
	// todo: add your logic here and delete this line
	vidReq := &vid.VidmgrStreamCountReq{
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	}
	vidResp, err := l.svcCtx.VidmgrS.VidmgrStreamCount(l.ctx, vidReq)
	if err != nil {
		er := errors.Fmt(err)
		l.Errorf("%s.rpc.VidmgrStreamCount req=%v err=%+v", utils.FuncName(), req, er)
		return nil, er
	}
	return &types.VidmgrStreamCountResp{
		Online:  vidResp.Online,
		Offline: vidResp.Offline,
	}, nil

}