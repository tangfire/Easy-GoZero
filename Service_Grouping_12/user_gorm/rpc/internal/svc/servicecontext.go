package svc

import (
	"Easy-GoZero/Combined_with_gorm_10/user/models"
	"Easy-GoZero/Service_Grouping_12/user_gorm/rpc/internal/config"
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
