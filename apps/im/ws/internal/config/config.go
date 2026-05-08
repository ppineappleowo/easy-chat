package config

import (
	"database/sql"
	"easy-chat/apps/im/ws/internal/handler/user"
	"github.com/zeromicro/go-zero/core/service"
)

type Config struct {
	service.ServiceConf

	ListenOn string

	JwtAuth struct {
		AccessSecret string
	}

	Mongo struct {
		Url string
		Db  string
	}
}
