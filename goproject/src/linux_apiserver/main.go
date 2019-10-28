package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"linux_apiserver/conf"
	"linux_apiserver/model"
	"linux_apiserver/router"
	"net/http"
	"time"
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
	gin.SetMode(viper.GetString("runmode"))
	router.Load(engine)

	go func() {
		if err := pingServer(); err != nil {
			//直接退出程序了
			log.Fatal("程序退出了...", err)
		}
	}()
	log.Infof("正在连接，端口是：%s", viper.GetString("port"))
	log.Info(http.ListenAndServe(viper.GetString("port"), engine).Error())
}

//配置自连接
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		resp, err := http.Get(viper.GetString("url") + "/test")
		if err == nil && resp.StatusCode == 200 {
			//连接成功
			log.Info("连接成功了....")
			return nil
		}
		log.Info("等待一秒钟，重新连接")
		time.Sleep(time.Second)
	}
	return errors.New("不能连接上服务器")
}
