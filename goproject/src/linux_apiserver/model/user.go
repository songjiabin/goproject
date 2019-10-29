package model

import (
	"gopkg.in/go-playground/validator.v9"
	"linux_apiserver/utils"
)

//和数据库中的表进行映射的实体类
type User struct {
	BaseModel
	Username string `gorm:"column:username" json:"username"  form:"username" binding:"required" validate="min=1,max=32"`
	Password string `gorm:"column:password" json:"password" form:"password"  binding:"required" validate="min=5,max=128"`
}

//关联数据库的名字
func (user *User) TableName() string {
	return "tb_users"
}

//创建用户
func CreateUser(user *User) error {
	return DB.Self.Create(user).Error
}

// 验证数据  根据约束 validate
func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

//对数据进行加密
func (u *User) Encrypt() (err error) {
	u.Password, err = utils.Encrypt(u.Password)
	return
}
