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
		"root:123root@A@tcp(192.144.238.85:3306)/test?charset=utf8&loc=Asia%2FShanghai")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("connection succedssed")
	}
	defer db.Close()
	createAndAdd(db)
}

type myselfModel struct {
	Id        int        `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"` //主键
	CreatedAt time.Time  `gorm:"column:createdAt" json:"-"`
	UpdatedAt time.Time  `gorm:"column:updatedAt" json:"-"`
	DeletedAt *time.Time `gorm:"column:deletedAt" sql:"index" json:"-"`
}

type User struct {
	myselfModel
	Name     string `gorm:"column:name;not null"`
	Passwrod string `gorm:"column:password;not null"`
}

//关联数据库的名字
func (UserModel *User) TableName() string {
	return "users"
}

//创建表并添加数据
func createAndAdd(db *gorm.DB) {
	 db.AutoMigrate(&User{})

	user := &User{Name: "蜡笔小新2"}

	err := db.Create(user).Error
	logs.Info(err)

	//u := []User{}
	//db.Exec("select * from users").Find(&u)
	//logs.Info(u[0].Id, u[0].Name, u[0].CreatedAt)

}
