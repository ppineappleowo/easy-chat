// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"easy-chat/apps/social/api/internal/config"
	"easy-chat/apps/social/api/internal/middleware"
	"easy-chat/apps/social/rpc/socialclient"
	"easy-chat/apps/user/rpc/userclient"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config                config.Config
	IdempotenceMiddleware rest.Middleware
	LimitMiddleware       rest.Middleware

	socialclient.Social
	userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                c,
		IdempotenceMiddleware: middleware.NewIdempotenceMiddleware().Handle,
		LimitMiddleware:       middleware.NewLimitMiddleware().Handle,

		Social: socialclient.NewSocial(zrpc.MustNewClient(c.SocialRpc)),
		User:   userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
