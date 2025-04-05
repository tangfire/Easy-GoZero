package handler

import (
	"net/http"

	"Easy-GoZero/demo_02/video/api/internal/logic"
	"Easy-GoZero/demo_02/video/api/internal/svc"
	"Easy-GoZero/demo_02/video/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func getVideoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VideoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetVideoLogic(r.Context(), svcCtx)
		resp, err := l.GetVideo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
