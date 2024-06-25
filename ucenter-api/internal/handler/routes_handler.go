package handler

import (
	"ucenter-api/internal/svc"
)

func RegisterHandlers(r *Router, serverCtx *svc.ServiceContext) {
	// 如果要有中间件 怎么办？
	register := NewRegisterHandler(serverCtx)
	registerGroup := r.Group()
	registerGroup.Post("/uc/register/phone", register.Register)
	registerGroup.Post("/uc/mobile/code", register.SendCode)
	loginGroup := r.Group()
	login := NewLoginHandler(serverCtx)
	loginGroup.Post("/uc/login", login.Login)
	loginGroup.Post("/uc/check/login", login.CheckLogin)
}
