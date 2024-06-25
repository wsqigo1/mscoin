package domain

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"mscoin-common/tools"
	"ucenter/internal/model"
	"ucenter/internal/repo"
)

type MemberDomain struct {
	memberRepo repo.MemberRepo
}

func NewMemberDomain(repo repo.MemberRepo) *MemberDomain {
	return &MemberDomain{
		memberRepo: repo,
	}
}

func (d *MemberDomain) FindByPhone(ctx context.Context, phone string) (*model.Member, error) {
	//涉及到数据库的查询
	mem, err := d.memberRepo.FindByPhone(ctx, phone)
	if err != nil {
		logx.Error(err)
		return nil, errors.New("数据库异常")
	}
	return mem, nil
}

func (d *MemberDomain) Register(
	ctx context.Context,
	phone string,
	password string,
	username string,
	country string,
	partner string,
	// 邀请码
	promotion string) error {
	mem := model.NewMember()
	//password 密码 进行md5加密 加盐 md5加密不安全（通过彩虹表进行破解）
	//member表字段比较多，所有的字段都不为null，也就是很多字段要填写默认值 写一个工具类 通过反射填充默认值即可
	_ = tools.Default(mem)
	salt, pwd := tools.Encode(password, nil)
	mem.Username = username
	mem.Country = country
	mem.Password = pwd
	mem.MobilePhone = phone
	mem.FillSuperPartner(partner)
	mem.PromotionCode = promotion
	mem.MemberLevel = model.GENERAL
	mem.Salt = salt
	mem.Avatar = "https://mszlu.oss-cn-beijing.aliyuncs.com/mscoin/defaultavatar.png"
	err := d.memberRepo.CreateMember(ctx, mem)
	if err != nil {
		logx.Error(err)
		return errors.New("数据库异常")
	}
	return nil
}

func (d *MemberDomain) UpdateLoginCount(ctx context.Context, id int64, step int) {
	err := d.memberRepo.UpdateLoginCount(ctx, id, step)
	if err != nil {
		logx.Error(err)
	}
}
