package conf

import (
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	name string
}

//初始化config
func Init(cfg string) error {
	config := Config{name: cfg}
	if err := config.InitConfit(); err != nil {
		return err
	}
	return nil
}

// 通过指定配置文件可以很方便地连接不同的环境（开发环境、测试环境）并加载不同的配置，方便开发和测试。
func (c *Config) InitConfit() error {
	name := c.name
	viper.SetConfigType("yaml") //配置类型
	if name == "" {
		viper.AddConfigPath("./conf")
		viper.SetConfigName("config")
	} else {
		viper.SetConfigName(name)
	}
	viper.AutomaticEnv()            //读取匹配的环境变量
	viper.SetEnvPrefix("APISERVER") //// 读取环境变量的前缀为APISERVER
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {
		log.Infof("错误是：", err)
		return err
	}

	return nil
}
