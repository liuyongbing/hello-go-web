package main

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/liuyongbing/hello-go-web/user-web/initinalize"
)

func main() {
	fmt.Println("Hello, this is for Go srvs of user srv API web.")

	port := 8021

	// Logger 初始化交由初始化层处理, 此处只负责调用
	// logger, _ := zap.NewProduction()
	// zap.ReplaceGlobals(logger)
	initinalize.InitLogger()

	// router := gin.Default()
	// 1. 路由配置交给专门的路由配置层处理
	// router.GET("/ping")

	// 2. 路由初始化交给初始化层处理
	// ApiGroup := router.Group("/v1")
	// iRouter.InitUserRouter(ApiGroup)

	// 3. 初始化 Router
	Router := initinalize.Routers()
	/*
		1. zap.S()可以获取一个全局的 sugar,可以让我们自己设置一个全局的 logger
		2. 日志的分级: debug, info, warning, error, fetal
		3. zap.S() & zap.L() 提供了一个全局的可安全访问 logger 的途径
	*/
	zap.S().Infof("启动服务器，端口：%d", port)

	err := Router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		zap.S().Panic("启动服务器失败：", err.Error())
	}
}
