package handler

import (
	"net/http"

	"book/api/internal/logic"
	"book/api/internal/svc"
	"book/api/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func bookUpdateHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewBookUpdateLogic(r.Context(), ctx)
		resp, err := l.BookUpdate(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
