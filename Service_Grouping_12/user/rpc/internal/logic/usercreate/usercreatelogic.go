package usercreatelogic

import (
	"context"
	"fmt"

	"Easy-GoZero/Service_Grouping_12/user/rpc/internal/svc"
	"Easy-GoZero/Service_Grouping_12/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCreateLogic {
	return &UserCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserCreateLogic) UserCreate(in *user.UserCreateRequest) (*user.UserCreateResponse, error) {
	// todo: add your logic here and delete this line
	fmt.Println(in.Username, in.Password)
	return &user.UserCreateResponse{}, nil
}
