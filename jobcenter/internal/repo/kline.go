package repo

import (
	"context"
	"jobcenter/internal/model"
	"jobcenter/internal/repo/dao"
)

type KlineRepo interface {
	SaveBatch(ctx context.Context, data []*model.Kline, symbol, period string) error
	DeleteGtTime(ctx context.Context, time int64, symbol string, period string) error
}

type klineRepo struct {
	dao *dao.KlineDao
}

func NewKlineRepo(dao *dao.KlineDao) KlineRepo {
	return &klineRepo{
		dao: dao,
	}
}

func (k *klineRepo) SaveBatch(ctx context.Context, data []*model.Kline, symbol, period string) error {
	return k.dao.SaveBatch(ctx, data, symbol, period)
}

func (k *klineRepo) DeleteGtTime(ctx context.Context, time int64, symbol string, period string) error {
	return k.dao.DeleteGtTime(ctx, time, symbol, period)
}
