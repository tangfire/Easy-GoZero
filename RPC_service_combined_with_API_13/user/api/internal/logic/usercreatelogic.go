package logic

import (
	"Easy-GoZero/RPC_service_combined_with_API_13/user/rpc/types/user"
	"context"
	"errors"

	"Easy-GoZero/RPC_service_combined_with_API_13/user/api/internal/svc"
	"Easy-GoZero/RPC_service_combined_with_API_13/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCreateLogic {
	return &UserCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCreateLogic) UserCreate(req *types.UserCreateRequest) (resp string, err error) {
	response, err := l.svcCtx.UserRpc.UserCreate(l.ctx, &user.UserCreateRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return "", err
	}
	if response.Err != "" {
		return "", errors.New(response.Err)
	}
	return
}
