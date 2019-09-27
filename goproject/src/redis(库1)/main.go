package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func handlerError(err error, when string) {

}

func main() {
	//连接到服务器
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "123456",
		DB:       0, // // use default DB
	})

	defer client.Close()

	pong, error := client.Ping().Result()
	fmt.Println(pong, error)

	//设置key
	er := client.Set("name", "songjiabin", 0).Err()
	if (er != nil) {

		return
	}

	//获取key对应的值
	val, e := client.Get("name").Result()
	if e != nil {
		fmt.Println(e)
		return
	}

	fmt.Println(val)

}
