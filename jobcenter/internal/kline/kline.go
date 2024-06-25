package kline

import (
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"jobcenter/internal/database"
	"jobcenter/internal/domain"
	"log"
	"mscoin-common/tools"
	"sync"
	"time"
)

type OkxConfig struct {
	ApiKey    string
	SecretKey string
	Pass      string
	Host      string
	Proxy     string
}

type OkxResult struct {
	Code string     `json:"code"`
	Msg  string     `json:"msg"`
	Data [][]string `json:"data"`
}

type Kline struct {
	wg          sync.WaitGroup
	c           OkxConfig
	klineDomain *domain.KlineDomain
	//queueDomain *domain.QueueDomain
}

func NewKline(c OkxConfig, client *database.MongoClient) *Kline {
	return &Kline{
		c:           c,
		klineDomain: domain.NewKlineDomain(client),
	}
}

func (k *Kline) Do(period string) {
	k.wg.Add(2)
	// 获取某个币 BTC-USDT ETH-USDT  BTC/USDT
	go k.getKlineData("BTC-USDT", "BTC/USDT", period)
	go k.getKlineData("ETH-USDT", "ETH/USDT", period)
	k.wg.Wait()
}

func (k *Kline) getKlineData(instId string, symbol string, period string) {
	defer k.wg.Done()
	// 发起http请求 获取数据
	api := k.c.Host + "/api/v5/market/candles?instId=" + instId + "&bar=" + period
	timestamp := tools.ISO(time.Now())
	sign := tools.ComputeHmacSha256(timestamp+"GET"+"/api/v5/market/candles?instId="+instId+"&bar="+period, k.c.SecretKey)
	header := make(map[string]string)
	header["OK-ACCESS-KEY"] = k.c.ApiKey
	header["OK-ACCESS-SIGN"] = sign
	header["OK-ACCESS-TIMESTAMP"] = timestamp
	header["OK-ACCESS-PASSPHRASE"] = k.c.Pass
	resp, err := tools.GetWithHeader(api, header, k.c.Proxy)
	if err != nil {
		log.Println(err)
		return
	}
	var result = &OkxResult{}
	err = json.Unmarshal(resp, result)
	if err != nil {
		log.Println(err)
		return
	}
	logx.Info(result)
	log.Println("==================执行存储mongo====================")
	if result.Code == "0" {
		//代表成功
		k.klineDomain.SaveBatch(result.Data, symbol, period)
		if "1m" == period {
			//把这个最新的数据result.Data[0] 推送到market服务，推送到前端页面，实时进行变化
			//->kafka->market kafka消费者进行数据消费-> 通过websocket通道发送给前端 ->前端更新数据
			//if len(result.Data) > 0 {
			//	k.queueDomain.Send1mKline(result.Data[0], symbol)
			//}
		}
	}
	log.Println("==================End====================")
}
