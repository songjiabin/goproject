package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"config"
	"err"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"time"
)

// 程序运行起来后，提示输入命令 输入”getall“ 查询并显示所有人员信息

// 第一次 从 mysql中查询并将结果缓存到redis中，并设置过期时间为60秒

// 以后每次查询 如果redis有数据就从redis中加载，没有则重复上一步操纵

var redisClient *redis.Pool

func init() {
	redisClient = &redis.Pool{
		MaxIdle:     16,                //最初的连接数量
		MaxActive:   0,                 //接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300 * time.Second, //超时时间 连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) {
			//连接
			con, error := redis.Dial(config.RedisConfig["type"], config.RedisConfig["address"])
			err.PrintErr(error)
			if _, e := con.Do("auth", "123456"); e != nil {
				return nil, e
			}
			return con, nil
		},
	}
}

func main() {

	for {
		fmt.Printf("请输入命令:")
		var cmd string
		fmt.Scan(&cmd)
		switch cmd {
		case "getall":
			getAllPersonForMysql()
		case "quit":
			//如果是break的话会直接退出switch 不会退出for的
			goto GAMEOVER
		default:
			fmt.Printf("命令不对！")
		}
	}

GAMEOVER:
	fmt.Println("结束了")

}

//从数据库中拿数据
func getAllPersonForMysql() {

	//先尝试从redis拿到缓存如果没有那么拿mysql中数据
	dataForRdis := getAllPeronForRedis()
	if dataForRdis != nil && len(dataForRdis) > 0 {
		fmt.Println("读取缓存成功,数据是：", dataForRdis)
		return
	}

	//连接mysql
	db, e := sqlx.Connect(config.MySqlConfig["mysql"], config.MySqlConfig["allconfig"])
	err.PrintErr(e)
	defer db.Close()

	var person []Person

	e2 := db.Select(&person, "select * from person;")
	err.PrintErr(e2)
	setAllPeronToRedis(&person)
}

//从redis中拿数据
func getAllPeronForRedis() (result []string) {
	rc := redisClient.Get()
	defer rc.Close()
	reply, err1 := rc.Do("lrange", "person", "0", "-1")
	err.PrintErr(err1)
	result, eeee := redis.Strings(reply, err1)
	err.PrintErr(eeee)
	return
}

//将得到是数据存储到redis
func setAllPeronToRedis(person *[]Person) {
	//连接redis 使用连接池
	rc := redisClient.Get()
	defer rc.Close()

	for _, value := range *person {
		///拿到实体对应的字符型
		v := fmt.Sprint(value)
		_, errs := rc.Do("rpush", "person", v)
		if errs != nil {
			err.PrintErr(errs)
			return
		}
	}

	//设置过期时间 为 60秒
	_, errs := rc.Do("expire", "person", 60)
	if errs != nil {
		err.PrintErr(errs)
	}

	fmt.Println("缓存成功！")
}

//用于接受对象的结构体
type Person struct {
	ID       int     `db:"id"`
	Name     string  `db:"name"`
	Age      int     `db:"age"`
	Money    float64 `db:"rmb"`
	Gender   bool    `db:"gender"`
	Birthday string  `db:"birthday"`
}
