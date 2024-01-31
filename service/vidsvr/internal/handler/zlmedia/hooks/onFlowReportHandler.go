package hooks

import (
	"gitee.com/i-Things/share/errors"
	"gitee.com/i-Things/share/result"
	"github.com/i-Things/things/service/vidsvr/internal/logic/zlmedia/hooks"
	"github.com/i-Things/things/service/vidsvr/internal/svc"
	"github.com/i-Things/things/service/vidsvr/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func OnFlowReportHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HooksApiFlowReportReq
		if err := httpx.Parse(r, &req); err != nil {
			result.Http(w, r, nil, errors.Parameter.WithMsg("入参不正确:"+err.Error()))
			return
		}

		l := hooks.NewOnFlowReportLogic(r.Context(), svcCtx)
		resp, err := l.OnFlowReport(&req)
		result.HttpWithoutWrap(w, r, resp, err)
	}
}
