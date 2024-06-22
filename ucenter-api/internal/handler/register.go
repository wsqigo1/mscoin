package handler

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"

	"ucenter-api/internal/svc"
)

type RegisterHandler struct {
	svcCtx *svc.ServiceContext
}

func NewRegisterHandler(svcCtx *svc.ServiceContext) *RegisterHandler {
	return &RegisterHandler{
		svcCtx: svcCtx,
	}
}

func (h *RegisterHandler) Register(w http.ResponseWriter, r *http.Request) {
	logx.Info("api register")
	return
}
