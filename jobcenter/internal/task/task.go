package task

import (
	"github.com/robfig/cron/v3"
	"jobcenter/internal/kline"
	"jobcenter/internal/svc"
)

type Task struct {
	c      *cron.Cron
	svcCtx *svc.ServiceContext
}

func NewTask(svcCtx *svc.ServiceContext) *Task {
	return &Task{
		c:      cron.New(cron.WithSeconds()),
		svcCtx: svcCtx,
	}
}

func (t *Task) Run() {
	_, _ = t.c.AddFunc("@every 1m", func() {
		kline.NewKline(t.svcCtx.Config.Okx, t.svcCtx.MongoClient).Do("1m")
	})
	_, _ = t.c.AddFunc("@every 3m", func() {
		kline.NewKline(t.svcCtx.Config.Okx, t.svcCtx.MongoClient).Do("3m")
	})
	_, _ = t.c.AddFunc("@every 5m", func() {
		kline.NewKline(t.svcCtx.Config.Okx, t.svcCtx.MongoClient).Do("5m")
	})
	_, _ = t.c.AddFunc("@every 15m", func() {
		kline.NewKline(t.svcCtx.Config.Okx, t.svcCtx.MongoClient).Do("15m")
	})
	_, _ = t.c.AddFunc("@every 30m", func() {
		kline.NewKline(t.svcCtx.Config.Okx, t.svcCtx.MongoClient).Do("30m")
	})
	_, _ = t.c.AddFunc("@every 1H", func() {
		kline.NewKline(t.svcCtx.Config.Okx, t.svcCtx.MongoClient).Do("1H")
	})
	_, _ = t.c.AddFunc("@every 2H", func() {
		kline.NewKline(t.svcCtx.Config.Okx, t.svcCtx.MongoClient).Do("2H")
	})
	_, _ = t.c.AddFunc("@every 1D", func() {
		kline.NewKline(t.svcCtx.Config.Okx, t.svcCtx.MongoClient).Do("1D")
	})
	_, _ = t.c.AddFunc("@every 1W", func() {
		kline.NewKline(t.svcCtx.Config.Okx, t.svcCtx.MongoClient).Do("1W")
	})
	_, _ = t.c.AddFunc("@every 1M", func() {
		kline.NewKline(t.svcCtx.Config.Okx, t.svcCtx.MongoClient).Do("1M")
	})

	t.c.Start()
}

func (t *Task) Start() {
	t.c.Start()
}

func (t *Task) Stop() {
	ctx := t.c.Stop()
	<-ctx.Done()
}
