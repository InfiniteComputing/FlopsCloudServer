package handler

import (
	"net/http"

	"gzfcserver/internal/logic"
	"gzfcserver/internal/svc"
	"gzfcserver/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GzfcserverHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGzfcserverLogic(r.Context(), svcCtx)
		resp, err := l.Gzfcserver(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
