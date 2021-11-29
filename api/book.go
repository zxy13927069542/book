package main

import (
	"book/api/internal/handler"
	"book/api/internal/svc"
	"book/config"
	"book/models"
	"fmt"
	"github.com/tal-tech/go-zero/rest"
)

func main() {
	//从config包获取配置
	c := config.GetConfig()
	//初始化数据库
	models.SetupDataBase()

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
