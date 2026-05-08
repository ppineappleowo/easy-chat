package svc

import (
	"easy-chat/apps/user/models"
	"easy-chat/apps/user/rpc/internal/config"
	"easy-chat/pkg/constants"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"time"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	*redis.Redis
	models.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:     c,
		Redis:      redis.MustNewRedis(c.RedisX),
		UsersModel: models.NewUsersModel(sqlConn, c.Cache),
	}
}

func (svc *ServiceContext) SetRootToken() error {
	//生成jwt
	systemToken, err := cxtdata.GetJwtToken(svc.Config.Jwt.AccessSecret, time.Now(), Unix(), constants.SYSTEM_ROOT_UID)
	if err != nil {
		return err
	}
	//写入redis
}
