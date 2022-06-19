package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	uuid "github.com/satori/go.uuid"

	"github.com/liuyongbing/hello-go-web/goods-web/global"
	"github.com/liuyongbing/hello-go-web/goods-web/initialize"
	"github.com/liuyongbing/hello-go-web/goods-web/utils/register/consul"
	myvalidator "github.com/liuyongbing/hello-go-web/goods-web/validator"
)

func main() {
	fmt.Println("Hello, this is for Go srvs of user srv API web.")

	// Logger 初始化交由初始化层处理, 此处只负责调用
	// logger, _ := zap.NewProduction()
	// zap.ReplaceGlobals(logger)
	initialize.InitLogger()

	// 初始化配置加载
	initialize.InitConfig()

	// 载入语言包
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}

	// 注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", myvalidator.ValidateMobile)
		_ = v.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0} 非法的手机号码!", true) // see universal-translator for details
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
	}

	// 初始化 GRPC Client
	initialize.InitSrvConn()

	// router := gin.Default()
	// 1. 路由配置交给专门的路由配置层处理
	// router.GET("/ping")

	// 2. 路由初始化交给初始化层处理
	// ApiGroup := router.Group("/v1")
	// iRouter.InitUserRouter(ApiGroup)

	// 3. 初始化 Router
	port := global.ServerConfig.Port
	Router := initialize.Routers()
	/*
		1. zap.S()可以获取一个全局的 sugar,可以让我们自己设置一个全局的 logger
		2. 日志的分级: debug, info, warning, error, fetal
		3. zap.S() & zap.L() 提供了一个全局的可安全访问 logger 的途径
	*/
	zap.S().Infof("启动服务器，端口：%d", port)

	// 服务注册
	addr := global.ServerConfig.Host
	name := global.ServerConfig.Name
	tags := global.ServerConfig.Tags
	id := uuid.NewV4().String()

	regClient := consul.NewRegistryClient(global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	err := regClient.Register(addr, port, name, tags, id)
	if err != nil {
		zap.S().Panic("服务注册失败：", err.Error())
	}

	go func() {
		if err := Router.Run(fmt.Sprintf(":%d", port)); err != nil {
			zap.S().Panic("启动服务器失败：", err.Error())
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	if err := regClient.DeRegister(id); err != nil {
		zap.S().Info("注销失败", err.Error())
	} else {
		zap.S().Info("注销成功")
	}
}
