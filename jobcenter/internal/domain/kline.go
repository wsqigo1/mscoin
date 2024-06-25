package domain

import (
	"context"
	"jobcenter/internal/database"
	"jobcenter/internal/model"
	"jobcenter/internal/repo"
	"jobcenter/internal/repo/dao"
	"log"
)

type KlineDomain struct {
	klineRepo repo.KlineRepo
}

func NewKlineDomain(client *database.MongoClient) *KlineDomain {
	d := dao.NewKlineDao(client.Db)
	return &KlineDomain{
		klineRepo: repo.NewKlineRepo(d),
	}
}

func (d *KlineDomain) SaveBatch(data [][]string, symbol string, period string) {
	klines := make([]*model.Kline, len(data))
	for i, v := range data {
		klines[i] = model.NewKline(v, period)
	}
	err := d.klineRepo.DeleteGtTime(context.Background(), klines[len(data)-1].Time, symbol, period)
	if err != nil {
		log.Println(err)
		return
	}
	err = d.klineRepo.SaveBatch(context.Background(), klines, symbol, period)
	if err != nil {
		log.Println(err)
	}
}
