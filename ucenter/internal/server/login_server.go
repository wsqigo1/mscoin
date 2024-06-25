package server

import (
	"context"
	"grpc-common/ucenter/types/login"
	"ucenter/internal/logic"
	"ucenter/internal/svc"
)

type LoginServer struct {
	svcCtx *svc.ServiceContext
	login.UnimplementedLoginServer
}

func NewLoginServer(svcCtx *svc.ServiceContext) *LoginServer {
	return &LoginServer{
		svcCtx: svcCtx,
	}
}

func (s *LoginServer) Login(ctx context.Context, req *login.LoginReq) (*login.LoginRes, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(req)
}
