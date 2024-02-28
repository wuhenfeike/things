package protocolmanagelogic

import (
	"context"
	"github.com/i-Things/things/service/dmsvr/internal/repo/relationDB"

	"github.com/i-Things/things/service/dmsvr/internal/svc"
	"github.com/i-Things/things/service/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProtocolInfoReadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProtocolInfoReadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProtocolInfoReadLogic {
	return &ProtocolInfoReadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 协议详情
func (l *ProtocolInfoReadLogic) ProtocolInfoRead(in *dm.WithIDCode) (*dm.ProtocolInfo, error) {
	f := relationDB.ProtocolInfoFilter{
		ID:   in.Id,
		Code: in.Code,
	}
	po, err := relationDB.NewProtocolInfoRepo(l.ctx).FindOneByFilter(l.ctx, f)
	return ToProtocolInfoPb(po), err
}