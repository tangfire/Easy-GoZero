// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.1
// Source: user_01.proto

package server

import (
	"context"

	"Easy-GoZero/demo_02/user/rpc/internal/logic"
	"Easy-GoZero/demo_02/user/rpc/internal/svc"
	"Easy-GoZero/demo_02/user/rpc/types/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) GetUser(ctx context.Context, in *user.IdRequest) (*user.UserResponse, error) {
	l := logic.NewGetUserLogic(ctx, s.svcCtx)
	return l.GetUser(in)
}
