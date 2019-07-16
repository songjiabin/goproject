package main

import (
	"github.com/gomodule/redigo/redis"
	"fmt"
	"os"
	"encoding/json"
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
	isDel, _ := redis.Bool(con.Do("del", "name"))
	fmt.Println("是否删除成功", isDel)

	//读取josn到redis
	handlerJson(con)

	//将json读取
	jsonToBean(con)

	//列表操作
	redisList(con)
}

//操作json
func handlerJson(con redis.Conn) {
	m := make(map[string]string, 10)
	m["usename"] = "666"
	m["phoneNum"] = "8888"
	j, _ := json.Marshal(m)

	//将json存到redis中去
	// setnx 当可以存在时啥也不做 当不存在的时候 操作
	n, _ := con.Do("setnx", "myJson", j)
	if n == int64(1) {
		fmt.Println("存储json成功")
	}

	//得到刚才的json
	result, error := redis.String(con.Do("get", "myJson"))
	errCheck(error)
	fmt.Println(result)
}

//将读取到的json转为bean
func jsonToBean(con redis.Conn) {
	m := make(map[string]interface{})
	b, e := redis.Bytes(con.Do("get", "myJson"))
	errCheck(e)
	merr := json.Unmarshal(b, &m)
	errCheck(merr)

	fmt.Println(m)
}

//redis 操作 List
func redisList(con redis.Conn) {

	//插入数据
	//_, e := con.Do("lpush", "listDemo", 100)
	//errCheck(e)

	v, _ := redis.Values(con.Do("lrange", "listDemo", "0", "-1"))

	for _, value := range v {
		//类型断言
		x, ok := value.([]byte)
		if ok {
			fmt.Println(string(x))
		}
	}

}

func errCheck(err error) {
	if err != nil {
		fmt.Println("sorry,has some error:", err)
		os.Exit(-1)
	}
}
