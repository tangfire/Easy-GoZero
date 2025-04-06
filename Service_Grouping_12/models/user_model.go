package models

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Username string `grom:"size:32" json:"username"`
	Password string `grom:"size:64" json:"password"`
}
