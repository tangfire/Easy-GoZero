package logic

import (
	"Easy-GoZero/go-zero_native_operation_mysql_09/user/models"
	"context"
	"fmt"

	"Easy-GoZero/go-zero_native_operation_mysql_09/user/api/internal/svc"
	"Easy-GoZero/go-zero_native_operation_mysql_09/user/api/internal/types"

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

func (l *LoginLogic) Login(req *types.LoginRequest) (resp string, err error) {
	// 增
	l.svcCtx.UsersModel.Insert(context.Background(), &models.User{
		Username: "tangfire",
		Password: "123456",
	})

	// 查
	user, err := l.svcCtx.UsersModel.FindOne(context.Background(), 1)
	fmt.Println(user, err)
	// 查
	user, err = l.svcCtx.UsersModel.FindOneByUsername(context.Background(), "枫枫")
	fmt.Println(user, err)

	// 改
	l.svcCtx.UsersModel.Update(context.Background(), &models.User{
		Username: "tangfire1",
		Password: "1234567",
		Id:       1,
	})
	user, err = l.svcCtx.UsersModel.FindOne(context.Background(), 1)
	fmt.Println(user, err)
	// 删
	l.svcCtx.UsersModel.Delete(context.Background(), 1)
	user, err = l.svcCtx.UsersModel.FindOne(context.Background(), 1)
	fmt.Println(user, err)
	return
}
