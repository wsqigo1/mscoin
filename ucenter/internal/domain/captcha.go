package domain

import (
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"mscoin-common/tools"
)

type captchaReq struct {
	Id        string `json:"id"`
	SecretKey string `json:"secretKey"`
	Scene     int    `json:"scene"`
	Token     string `json:"token"`
	Ip        string `json:"ip"`
}

type captchaRsp struct {
	Success int    `json:"success"`
	Score   int    `json:"score"`
	Msg     string `json:"msg"`
}

type CaptchaDomain struct {
	vid string
	key string
}

func NewCaptchaDomain(vid string, key string) *CaptchaDomain {
	return &CaptchaDomain{
		vid: vid,
		key: key,
	}
}

// 人机校验的ip地址
func (d *CaptchaDomain) Verify(server string,
	token string, scene int, ip string) bool {
	// 发送一个 post 请求
	resp, err := tools.Post(server, &captchaReq{
		Id:        d.vid,
		SecretKey: d.key,
		Token:     token,
		Scene:     scene,
		Ip:        ip,
	})
	if err != nil {
		logx.Error(err)
		return false
	}
	res := &captchaRsp{}
	err = json.Unmarshal(resp, res)
	if err != nil {
		logx.Error(err)
		return false
	}
	return res.Success == 1
}
