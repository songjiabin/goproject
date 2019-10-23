package model

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"gopkg.in/go-playground/validator.v9"
	"myapiserver/constvar"
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
func DeleteUser(id uint64) error {
	user := UserModel{}
	user.Id = id
	return DB.Self.Delete(&user).Error
}

//更新数据
func (model *UserModel) UpdateUser() error {
	return DB.Self.Save(model).Error
}

//获取所有的用户信息
func ListUser(username string, offset, limit int) ([]*UserModel, uint, error) {
	//获取默认值
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	var count uint
	users := make([]*UserModel, 0)
	where := fmt.Sprintf("username like '%%%s%%'", username)

	//先用 db.Model() 选择一个表  并计算数量
	if err := DB.Self.Model(&UserModel{}).Where(where).Count(&count).Error; err != nil {
		return users, count, err
	}

	logs.Info(">>>>>>>>>>>>>>>>", users)

	//可以使用 db.Find(&Likes) 或者只需要查一条记录 db.First(&Like)
	if err := DB.Self.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users, count, err
	}

	logs.Info("<<<<<<<<<<<<<<<<<", users[0].Username, users[0].Password)

	var u []UserModel
	DB.Self.Exec("select * from tb_users").Find(&u)
	logs.Info("<<<<<<<<<<<<<<<<<", u[0].Username, u[0].Password, u[0].CreatedAt)

	return users, count, nil

}

//从数据库获取user
func GetUser(username string) (*UserModel, error) {
	u := &UserModel{}
	result := DB.Self.Where("username = ?", username).First(&u)
	return u, result.Error
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
