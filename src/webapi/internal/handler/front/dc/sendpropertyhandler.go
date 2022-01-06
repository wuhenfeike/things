package dc

import (
	"github.com/go-things/things/src/webapi/internal/logic/front/dc"
	"net/http"

	"github.com/go-things/things/src/webapi/internal/svc"
	"github.com/go-things/things/src/webapi/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func SendPropertyHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendDcPropertyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := dc.NewSendPropertyLogic(r.Context(), ctx)
		resp, err := l.SendProperty(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
