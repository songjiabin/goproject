package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

type Database struct {
	Self   *gorm.DB
	Docker *gorm.DB
}

var DB *Database

//初始化数据库的一些连接
func (db *Database) Init() {
	db = &Database{
		Self:   getSelfDB(),
		Docker: nil,
	}
}

func getSelfDB() *gorm.DB {
	return openDB(
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"),
	)
}

func getDockerDB() *gorm.DB {
	return openDB(
		viper.GetString("docker.username"),
		viper.GetString("docker.password"),
		viper.GetString("docker.addr"),
		viper.GetString("docker.name"),
	)
}

func openDB(username, password, addr, name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		true,
		"local",
	)
	db, err := gorm.Open("mysql", config)
	if err != nil {
		log.Errorf(err, "Database connection failed. Database name: %s", name)
	}
	return db
}
