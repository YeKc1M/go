package handler

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gozerodemo/testapi/internal/logic"
	"gozerodemo/testapi/internal/svc"
	"gozerodemo/testapi/internal/types"
)

func TestapiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewTestapiLogic(r.Context(), svcCtx)
		resp, err := l.Testapi(&req)
		logx.Infof("%v", *resp)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
