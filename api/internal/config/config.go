package config

import (
	"flag"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

type Config struct {
	rest.RestConf

	Mysql struct {
		DataSource string
	}
}

var configFile = flag.String("f", "etc/book-api.yaml", "the config file")

/**
提供一个接口给其他包来获取配置参数
*/
func GetConfig() Config {
	flag.Parse() //从运行指令获取参数
	var c Config
	conf.MustLoad(*configFile, &c) //从默认配置文件或者指令的配置文件读取配置
	return c
}
