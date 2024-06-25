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
	_, _ = t.c.AddFunc("every 1m", func() {
		kline.NewKline(t.svcCtx.Config.Okx).Do("1m")
	})
	_, _ = t.c.AddFunc("every 1h", func() {
		kline.NewKline(t.svcCtx.Config.Okx).Do("1h")
	})
	t.c.Start()
}

func (t *Task) Start() {
	t.c.Start()
}

func (t *Task) Stop() {
	t.c.Stop()
}
