package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

//数据库
/**
因为一个 API 服务器可能需要同时访问多个数据库，
为了对多个数据库进行初始化和连接管理，这里定义了一个叫 `Database` 的 struct：
*/
type Database struct {
	Self   *gorm.DB
	Docker *gorm.DB
}

var DB *Database

//每个数据库有个连接的初始化方法
func (db *Database) Init() {
	DB = &Database{
		Self:   getSelfDB(),
		Docker: getDockerDB(),
	}
}

//获取 gorm连接的 DB
func getSelfDB() *gorm.DB {
	return initSelfDB()
}

func getDockerDB() *gorm.DB {
	return initDockerDB()
}

//初始化self类型的连接
func initSelfDB() *gorm.DB {
	return openDB(
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"),
	)
}

//初始化docker类型的连接
func initDockerDB() *gorm.DB {
	return openDB(
		viper.GetString("docker.username"),
		viper.GetString("docker.password"),
		viper.GetString("docker.addr"),
		viper.GetString("docker.name"),
	)

}

/**
username: 用户名
password: 密码
addr:    地址
name     是用哪个数据库
*/
func openDB(username, password, addr, name string) *gorm.DB {
	//root:123root@A@tcp(192.144.238.85:3306)/mydb?charset=utf8&loc=Asia%2FShanghai
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		"Local") //   Asia%2FShanghai  Local
		fmt.Println("config----->",config)

	db, err := gorm.Open("mysql", config)
	if err != nil {
		log.Errorf(err, "Database connection failed. Database name: %s", name)
	}

	//设置db连接的
	setupDB(db)
	return db
}

//设置db连接
func setupDB(db *gorm.DB) {
	//是否设置gorm的log打印
	db.LogMode(viper.GetBool("gormlog"))
	//db.DB().SetMaxOpenConns(20000) // 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	db.DB().SetMaxIdleConns(0) // 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
}

func (db *Database) Close() {
	db.Self.Close()
	db.Docker.Close()
}
