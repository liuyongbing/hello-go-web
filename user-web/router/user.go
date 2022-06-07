package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/liuyongbing/hello-go-web/user-web/api"
	"github.com/liuyongbing/hello-go-web/user-web/middlewares"
)

func InitUserRouter(Router *gin.RouterGroup) {
	// 多个模块都会使用 router 实例，因此全局化，让调用者传入使用
	// router := gin.Default()
	// router.GET("/ping")

	zap.S().Infof("配置 User API Url")

	// API: User
	UserRouter := Router.Group("user")
	{
		// UserRouter.GET("list", api.GetUserList) // 获取用户列表
		UserRouter.GET("", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetUserList) // 获取用户列表
		UserRouter.POST("pwd_login", api.PassWordLogin)                                       // 用户登录
	}
}
