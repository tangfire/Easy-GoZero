package logic

import (
	"context"
	"fmt"

	"Easy-GoZero/api_syntax/user/api/internal/svc"
	"Easy-GoZero/api_syntax/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (
	resp *types.LoginResponse, err error) {
	fmt.Println(req.Username, req.Password)

	return &types.LoginResponse{Code: 0, Data: "tangfire", Msg: "成功"}, nil
}
