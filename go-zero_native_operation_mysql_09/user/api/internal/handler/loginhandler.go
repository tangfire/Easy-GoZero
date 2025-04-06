package handler

import (
	"net/http"

	"Easy-GoZero/go-zero_native_operation_mysql_09/user/api/internal/logic"
	"Easy-GoZero/go-zero_native_operation_mysql_09/user/api/internal/svc"
	"Easy-GoZero/go-zero_native_operation_mysql_09/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
