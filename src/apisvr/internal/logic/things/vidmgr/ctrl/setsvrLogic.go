package ctrl

import (
	"context"

	"github.com/i-Things/things/src/apisvr/internal/svc"
	"github.com/i-Things/things/src/apisvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetsvrLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetsvrLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetsvrLogic {
	return &SetsvrLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetsvrLogic) Setsvr(req *types.CtrlApiReq) (resp *types.CtrlApiResp, err error) {
	// todo: add your logic here and delete this line

	return
}