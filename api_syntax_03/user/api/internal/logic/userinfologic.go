package logic

import (
	"context"

	"Easy-GoZero/api_syntax_03/user/api/internal/svc"
	"Easy-GoZero/api_syntax_03/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserResponse, err error) {

	return &types.UserResponse{Code: 0, Data: types.UserInfo{
		UserId:   1,
		Username: "tangfire",
	}, Msg: "成功"}, nil
}
