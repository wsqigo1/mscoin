package svc

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"mscoin-common/msdb"
	"ucenter/internal/config"
	"ucenter/internal/database"
)

type ServiceContext struct {
	Config config.Config
	Cache  cache.Cache
	DB     *msdb.MsDB
	//MarketR
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisCache := cache.New(
		c.CacheRedis,
		nil,
		cache.NewStat("mscoin"),
		nil,
	)
	mysql := database.ConnMysql(c.Mysql.DataSource)
	//cli := database.NewKafkaClient(c.Kafka)
	//cli.StartRead("add-exchange-order")
	//order := eclient.NewOrder(zrpc.MustNewClient(c.ExchangeRpc))
	return &ServiceContext{
		Config: c,
		Cache:  redisCache,
		DB:     mysql,
	}
}
