// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package handler

import (
	"net/http"

	"go-zero-mall/service/order/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/order/create",
				Handler: CreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/order/detail",
				Handler: DetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/order/list",
				Handler: ListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/order/remove",
				Handler: RemoveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/order/update",
				Handler: UpdateHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
