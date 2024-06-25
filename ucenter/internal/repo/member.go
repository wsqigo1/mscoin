package repo

import (
	"context"
	"ucenter/internal/model"
	"ucenter/internal/repo/dao"
)

type MemberRepo interface {
	FindByPhone(ctx context.Context, phone string) (*model.Member, error)
	CreateMember(ctx context.Context, mem *model.Member) error
	UpdateLoginCount(ctx context.Context, id int64, step int) error
	FindMemberById(ctx context.Context, memberId int64) (*model.Member, error)
}

type memberRepo struct {
	dao *dao.MemberDAO
}

func (r *memberRepo) FindByPhone(ctx context.Context, phone string) (*model.Member, error) {
	return r.dao.FindByPhone(ctx, phone)
}

func (r *memberRepo) CreateMember(ctx context.Context, mem *model.Member) error {
	return r.dao.InsertMember(ctx, mem)
}

func (r *memberRepo) UpdateLoginCount(ctx context.Context, id int64, step int) error {
	return r.dao.UpdateLoginCount(ctx, id, step)
}

func (r *memberRepo) FindMemberById(ctx context.Context, memberId int64) (*model.Member, error) {
	return r.dao.FindMemberById(ctx, memberId)
}

func NewMemberRepo(dao *dao.MemberDAO) MemberRepo {
	return &memberRepo{
		dao: dao,
	}
}
