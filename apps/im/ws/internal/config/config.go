package config

import (
	"easy-chat/apps/im/ws/internal/handler/user"
	"github.com/zeromicro/go-zero/core/service"
)

type Config struct {
	service.ServiceConf

	ListenOn string

	JwtAuth struct {
		AccessSecret string
	}
}
