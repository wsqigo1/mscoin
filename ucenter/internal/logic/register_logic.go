package logic

import (
	"context"
	"grpc-common/ucenter/types/register"
	"ucenter/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterByPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) RegisterByPhone(in *register.RegReq) (*register.ReqRes, error) {
	// todo: add your logic here and delete this line

	return &register.ReqRes{}, nil
}
