// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.1

package handler

import (
	"net/http"

	"Easy-GoZero/api_syntax/user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/users/info",
				Handler: userInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/users/login",
				Handler: loginHandler(serverCtx),
			},
		},
	)
}
