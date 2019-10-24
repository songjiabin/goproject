package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"linux_apiserver/conf"
	"linux_apiserver/model"
	"linux_apiserver/router"
)

var (
	// apiservice.exe -c config.yaml
	cfg = pflag.StringP("config", "c", "", "apiserver config file path")
)

func main() {
	//设置配置文件的读取
	pflag.Parse()
	if err := conf.Init(*cfg); err != nil {
		panic(err)
	}

	//连接数据库
	model.DB.Init()

	//new engine
	engine := gin.New()
	router.Load(engine)
	engine.Run()
}
