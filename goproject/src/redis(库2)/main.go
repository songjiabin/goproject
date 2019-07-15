package main

import (
	"github.com/gomodule/redigo/redis"
	"fmt"
	"os"
)

func main() {

	//创建连接
	con, error := redis.Dial("tcp", "127.0.0.1:6379",
		//redis.DialDatabase(1), //dialOption参数可以配置选择数据库、连接密码、心跳检测等等
		redis.DialPassword("123456"))
	defer con.Close()
	errCheck(error)

	//对本次连接进行set操作
	_, setErr := con.Do("SET", "url", "https://songjiabin.github.io/")
	errCheck(setErr)

	//检测值是否存在
	exists, errExists := redis.Bool(con.Do("EXISTS", "url"))
	errCheck(errExists)
	fmt.Println(exists)

	// 对本次的get
	r, getErr := redis.String(con.Do("GET", "url"))
	errCheck(getErr)
	fmt.Println(r)

	//给定一个kv的过期时间  ex过期时间为5秒
	_, e := con.Do("set", "sex", "男", "ex", "5")
	errCheck(e)

	//得到刚才设置的数据
	result, er := redis.String(con.Do("get", "sex"))
	errCheck(er)
	fmt.Println(result)

	//删除key
	isDel,_:=redis.Bool(con.Do("del", "name"))
	fmt.Println("是否删除成功",isDel)

}

func errCheck(err error) {
	if err != nil {
		fmt.Println("sorry,has some error:", err)
		os.Exit(-1)
	}
}
