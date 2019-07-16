package main

import (
	"github.com/gomodule/redigo/redis"
	"config"
	"time"
	"fmt"
)

var (
	//定义常量
	redisClient *redis.Pool //redis 连接池
	REDIS_DB    int
)

//初始化连接池
func init() {
	// 从配置文件获取redis的ip以及db
	redisClient = &redis.Pool{
		MaxIdle:     16,                //最初的连接数量
		MaxActive:   0,                 //接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300 * time.Second, //超时时间 连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial(config.RedisConfig["type"], config.RedisConfig["address"])
			if err != nil {
				return nil, err
			}

			if _, errAuth := con.Do("auth", config.RedisConfig["auth"]); err != nil {
				return nil, errAuth
			}

			return con, nil
		},
	}
}

//连接池
func main() {
	//获取连接池i
	rc := redisClient.Get()
	//用完连接放回到连接池
	defer rc.Close()
	r, e := redis.String(rc.Do("get", "name"))
	fmt.Println(r, e)

	for i := 0; i < 10; i++ {
		//开辟并发任务执行任务
		go func() {

		}()
	}

}
