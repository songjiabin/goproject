package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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

type User struct {
	gorm.Model
	Name string `gorm:"not null"`
}

//创建表并添加数据
func createAndAdd(db *gorm.DB) {
	db.AutoMigrate(&User{})
	user := &User{}
	err := db.Create(user).Error
	logs.Info(err)

}
