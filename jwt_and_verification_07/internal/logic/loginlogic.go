package logic

import (
	"Easy-GoZero/common/jwts"
	"context"

	"Easy-GoZero/jwt_and_verification_07/internal/svc"
	"Easy-GoZero/jwt_and_verification_07/internal/types"

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
	auth := l.svcCtx.Config.Auth
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		UserID:   1,
		Username: "tangfire",
		Role:     1,
	}, auth.AccessSecret, auth.AccessExpire)
	if err != nil {
		return "", err
	}
	return token, err
}
