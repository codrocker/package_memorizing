package handler

import (
	"net/http"
	"package_memorizing/internal/logic"
	"package_memorizing/internal/svc"
	"package_memorizing/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CommenceHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommenceRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCommenceLogic(r.Context(), svcCtx)
		resp, err := l.Commence(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
