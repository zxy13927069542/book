package config

import (
	"flag"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
	"reflect"
)

type Config struct {
	rest.RestConf

	Mysql struct {
		DataSource string
	}
}

var configFile = flag.String("f", "config/book-api.yaml", "the config file")

var c Config

/**
提供一个接口给其他包来获取配置参数
*/
func GetConfig() Config {
	flag.Parse()
	//如果已有配置，则不再覆盖
	equal := reflect.DeepEqual(c, Config{})
	if equal {
		//从运行指令获取参数
		conf.MustLoad(*configFile, &c) //从默认配置文件或者指令的配置文件读取配置
	}
	return c
}

func SetConfig() *Config {
	return &c
}
