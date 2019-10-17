package model

import (
	"gopkg.in/go-playground/validator.v9"
	"myapiserver/pkg/auth"
)

type UserModel struct {
	BaseModel
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

//删除数据
func DeleteUser(id uint) error {
	user := UserModel{}
	user.Id = id
	return DB.Self.Delete(&user).Error
}

//更新数据
func UpdateUser(model *UserModel) error {
	return DB.Self.Save(model).Error
}

// 验证数据  根据约束 validate
func (u *UserModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

// 对密码加密
func (u *UserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}
