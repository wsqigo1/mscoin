package handler

import (
	"github.com/zeromicro/go-zero/rest"
	"net/http"
)

type Routes struct {
	server *rest.Server
}

func NewRouters(server *rest.Server) *Routes {
	return &Routes{
		server: server,
	}
}

func (r *Routes) Get(path string, handler http.HandlerFunc) {
	r.server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Handler: handler,
		Path:    path,
	})
}

func (r *Routes) Post(path string, handler http.HandlerFunc) {
	r.server.AddRoute(rest.Route{
		Method:  http.MethodPost,
		Handler: handler,
		Path:    path,
	})
}
