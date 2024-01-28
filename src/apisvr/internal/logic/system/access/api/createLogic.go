package api

import (
	"context"
	"github.com/i-Things/things/src/apisvr/internal/svc"
	"github.com/i-Things/things/src/apisvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.ApiInfo) (resp *types.WithID, err error) {
	rst, err := l.svcCtx.AccessRpc.ApiInfoCreate(l.ctx, ToApiInfoPb(req))
	if err != nil {
		return nil, err
	}
	return &types.WithID{ID: rst.Id}, nil
}
