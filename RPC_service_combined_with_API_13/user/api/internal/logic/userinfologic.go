package logic

import (
	"Easy-GoZero/RPC_service_combined_with_API_13/user/rpc/types/user"
	"context"

	"Easy-GoZero/RPC_service_combined_with_API_13/user/api/internal/svc"
	"Easy-GoZero/RPC_service_combined_with_API_13/user/api/internal/types"

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

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (resp *types.UserInfoResponse, err error) {
	response, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		UserId: uint32(req.ID),
	})

	if err != nil {
		return nil, err
	}

	return &types.UserInfoResponse{UserId: uint(response.UserId), Username: response.Username}, nil
}
