package logic

import (
	"context"

	"Easy-GoZero/Single_rpc_service_mode_11/user/rpc/internal/svc"
	"Easy-GoZero/Single_rpc_service_mode_11/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	// todo: add your logic here and delete this line

	return &user.UserInfoResponse{
		UserId:   in.UserId,
		Username: "tangfire",
	}, nil
}
