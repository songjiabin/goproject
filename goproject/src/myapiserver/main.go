package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"myapiserver/config"
	"myapiserver/model"
	"myapiserver/router"
	"net/http"
	"time"
)

var (
	//命令行参数
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)

func main() {
	//初始化配置 读取配置文件
	//解析函数  将会在碰到第一个非flag命令行参数时停止
	pflag.Parse()
	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	//初始化数据库操作
	model.DB.Init()
	defer model.DB.Close()

	// Set gin mode.
	gin.SetMode(viper.GetString("runmode"))

	//Create the Gin engine.
	g := gin.New()
	// Set gin mode.
	//gin 有 3 种运行模式：debug、release 和 test，其中 debug 模式会打印很多 debug 信息。
	gin.SetMode(viper.GetString("runmode"))
	// gin middlewares
	middlewares := []gin.HandlerFunc{}
	router.Load(g, middlewares...)
	go func() {
		if err := pingServer(); err != nil {
			// 路由器无响应，或者启动时间可能太长。
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		//路由器已成功部署
		log.Infof("The router has been deployed successfully.")
	}()
	//监听端口 8080
	log.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())

}

//验证请求
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		//等待一秒钟
		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}

	return errors.New("Cannot connect to the router.")

}
