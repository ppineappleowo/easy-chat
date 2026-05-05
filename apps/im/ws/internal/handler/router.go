package handler

import (
	"easy-chat/apps/im/ws/internal/handler/user"
	"easy-chat/apps/im/ws/internal/svc"
	"easy-chat/apps/im/ws/websocket"
)

func RegisterHandlers(srv *websocket.Server, svc *svc.ServiceContext) {
	srv.AddRoute([]websocket.Route{
		{
			Method:  "user.Online",
			Handler: user.OnLine(svc),
		},
	})
}
