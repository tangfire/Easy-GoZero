package logic

import (
	"Easy-GoZero/Combined_with_gorm_10/user/models"
	"context"
	"fmt"

	"Easy-GoZero/Combined_with_gorm_10/user/api/internal/svc"
	"Easy-GoZero/Combined_with_gorm_10/user/api/internal/types"

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
	// todo: add your logic here and delete this line
	err = l.svcCtx.DB.Create(&models.UserModel{
		Username: "fireshine",
		Password: "123456",
	}).Error
	fmt.Println(err)

	return "", nil
}
