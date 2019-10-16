package model

import (
	"gopkg.in/go-playground/validator.v9"
	"myapiserver/pkg/auth"
)

type UserModel struct {
	Username string `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
}

//关联数据库的名字
func (UserModel *UserModel) TableName() string {
	return "tb_users"
}

//插入数据到数据库中
func (user *UserModel) Create() error {
	return DB.Self.Create(&user).Error
}

// 验证数据
func (u *UserModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

// 对密码加密
func (u *UserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}
