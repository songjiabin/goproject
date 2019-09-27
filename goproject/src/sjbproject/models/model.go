package models

import (
	"config"
	"github.com/astaxie/beego/orm"
	"github.com/gomodule/redigo/redis"
	"time"
)
import _ "github.com/go-sql-driver/mysql"

var (
	//定义常量
	RedisClient *redis.Pool //redis 连接池
	REDIS_DB    int
)

//用户表
type User struct {
	Id       int
	UserName string `orm:"unique"`
	Passwd   string
	Article  []*Article `orm:"reverse(many)"` //多对多
}

//文章表
type Article struct {
	Id          int          `orm:"pk;auto"`                          //主键 自增
	Title       string       `orm:"size(20)""`                        //标题
	Content     string       `orm:"size(200)"`                        //内容
	Img         string       `orm:"size(100)"`                        // 图片
	Time        time.Time    `orm:"type(datetime);auto_now_add;null"` //创建时间
	Count       int          `orm:"default(0);null"`                  //阅读次数
	Num         float64      `orm:"digits(12);decimals(4);null"`      //  浮点型 一共12位数，小数点后面4位
	ArticleType *ArticleType `orm:"rel(fk)"`                          //外键  设置一对多的关系
	User        []*User      `orm:"rel(m2m)"`                         ////多对多
}

//类型表   一个文章只属于一个类型 ，但是一个类型中有多个文章
type ArticleType struct {
	Id       int        `orm:"pk;auto"`
	TypeName string     `orm:"size(20);unique"`
	Article  []*Article `orm:"reverse(many)"` //一对多的反向关系
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:123root@A@tcp(192.144.238.85:3306)/mydb?charset=utf8&loc=Asia%2FShanghai")
	orm.RegisterModel(new(User), new(Article), new(ArticleType))
	orm.RunSyncdb("default", false, true)

	//连接Redis
	connReidis()

}

func connReidis() {
	RedisClient = &redis.Pool{
		MaxIdle:     16,                //最大的连接数量
		MaxActive:   0,                 //接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300 * time.Second, //超时时间 连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (conn redis.Conn, e error) {
			con, err := redis.Dial(config.RedisConfig["type"], "192.144.238.85:6379")
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
