package svc

import (
	"Easy-GoZero/go-zero_native_operation_mysql_09/user/api/internal/config"
	"Easy-GoZero/go-zero_native_operation_mysql_09/user/models"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	UsersModel models.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		UsersModel: models.NewUserModel(mysqlConn),
	}
}
