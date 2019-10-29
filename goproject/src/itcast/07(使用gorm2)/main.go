package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type BaseModel struct {
	Id int `gorm:"column:primary_key;AUTO_INCREMENT;column:id" json:"-"`


	CreatedAt time.Time  `gorm:"column:createdAt" json:"createdtime"`
	UpdatedAt time.Time  `gorm:"column:updatedAt" json:"updatedtime"`
	DeletedAt *time.Time `gorm:"column:deletedAt" json:"deletedtime" sql:"index"`
}

type User struct {
	BaseModel
	Username string `gorm:"column:username" json:"username"  form:"username" binding:"required" validate="min=1,max=32"`
	Password string `gorm:"column:password" json:"password" form:"password"  binding:"required" validate="min=5,max=128"`

	//Name     string `gorm:"column:name;not null"`
	//Passwrod string `gorm:"column:password;not null"`
}

func (user *User) TableName() string {
	return "tb_users"
}

func main() {
	db, err := gorm.Open("mysql",
		"root:123root@A@tcp(47.95.5.232:3306)/db_apiserver?charset=utf8&parseTime=true&loc=Asia%2FShanghai")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("connection succedssed")
	}
	defer db.Close()

	u := &User{
		Username: "宋佳宾",
		Password: "111111",
	}

	errs := db.Create(u).Error
	fmt.Println(errs)

}
