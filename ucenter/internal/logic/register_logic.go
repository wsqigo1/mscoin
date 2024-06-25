package logic

import (
	"context"
	"errors"
	"grpc-common/ucenter/types/register"
	"ucenter/internal/domain"
	"ucenter/internal/repo"
	"ucenter/internal/repo/dao"
	"ucenter/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx           context.Context
	svcCtx        *svc.ServiceContext
	CaptchaDomain *domain.CaptchaDomain
	MemberDomain  *domain.MemberDomain
	logx.Logger
}

func NewRegisterByPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	// todo: wire注册
	db := dao.NewMemberDAO(svcCtx.DB)
	r := repo.NewMemberRepo(db)
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		CaptchaDomain: domain.NewCaptchaDomain(svcCtx.Config.Captcha.Vid,
			svcCtx.Config.Captcha.Key),
		MemberDomain: domain.NewMemberDomain(r),
	}
}

// RegisterByPhone {code:111,msg:xxx}
func (l *RegisterLogic) RegisterByPhone(in *register.RegReq) (*register.ReqRes, error) { //1. 先校验人机是否通过
	// 1. 先校验人机是否通过
	if !l.CaptchaDomain.Verify(in.Captcha.Server, in.Captcha.Token, 2, in.Ip) {
		return nil, errors.New("人机校验不通过")
	}
	// 2. 校验验证码
	redisValue := ""
	err := l.svcCtx.Cache.GetCtx(context.Background(),
		RegisterCacheKey+in.Phone, &redisValue)
	if err != nil {
		return nil, errors.New("验证码获取错误")
	}
	if in.Code != redisValue {
		return nil, errors.New("验证码输入错误")
	}
	//3.验证码通过 进行注册即可 手机号首先验证此手机号是否注册过
	mem, err := l.MemberDomain.FindByPhone(context.Background(), in.Phone)
	if err != nil {
		return nil, errors.New("服务异常，请联系管理员")
	}
	if mem != nil {
		return nil, errors.New("此手机号已经被注册")
	}
	//4. 生成member模型，存入数据库
	err = l.MemberDomain.Register(
		context.Background(),
		in.Phone,
		in.Password,
		in.Username,
		in.Country,
		in.SuperPartner,
		in.Promotion)
	if err != nil {
		return nil, errors.New("注册失败")
	}
	return &register.ReqRes{}, nil
}
