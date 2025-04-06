package svc

import (
	"Easy-GoZero/RPC_service_combined_with_API_13/user/models"
	"Easy-GoZero/RPC_service_combined_with_API_13/user/rpc/internal/config"
	"Easy-GoZero/common/init_gorm"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := init_gorm.InitGorm(c.Mysql.DataSource)
	db.AutoMigrate(&models.UserModel{})
	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
