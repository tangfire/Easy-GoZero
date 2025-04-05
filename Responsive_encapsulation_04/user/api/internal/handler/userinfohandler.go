package handler

import (
	"Easy-GoZero/common/response"
	"net/http"

	"Easy-GoZero/Responsive_encapsulation_04/user/api/internal/logic"
	"Easy-GoZero/Responsive_encapsulation_04/user/api/internal/svc"
)

func userInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
		response.Response(r, w, resp, err)
	}
}
