package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

func main() {
	db, err := gorm.Open("mysql",
		"root:123root@A@tcp(192.144.238.85:3306)/test?charset=utf8&parseTime=true&loc=Asia%2FShanghai")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("connection succedssed")
	}
	defer db.Close()
	createAndAdd(db)
}

type BaseModel struct {
	Id        uint       `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"` //主键
	CreatedAt time.Time  `gorm:"column:createdAt" json:"-"`
	UpdatedAt time.Time  `gorm:"column:updatedAt" json:"-"`
	DeletedAt *time.Time `gorm:"column:deletedAt" sql:"index" json:"-"`
}

type User struct {
	BaseModel
	Name     string `gorm:"column:name;not null"`
	Passwrod string `gorm:"column:password;not null"`
}

//关联数据库的名字
func (UserModel *User) TableName() string {
	return "users"
}

//创建表并添加数据
func createAndAdd(db *gorm.DB) {

	/*user := &User{Name: "蜡笔小新", Passwrod: "w3ewe"}
	err := db.Create(user).Error
	logs.Info(err)*/

	user := User{}
	db.First(&user)
	logs.Info("name-->", user.Name)
	logs.Info("password-->", user.Passwrod)
	logs.Info("createTime-->", user.CreatedAt.Format("2006-01-02 15:04:05"))
}
