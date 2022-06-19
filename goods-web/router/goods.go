package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/liuyongbing/hello-go-web/goods-web/api"
)

func InitGoodsRouter(Router *gin.RouterGroup) {
	// 多个模块都会使用 router 实例，因此全局化，让调用者传入使用
	// router := gin.Default()
	// router.GET("/ping")

	zap.S().Infof("配置 Goods API Url")

	// API: User
	UserRouter := Router.Group("goods")
	{
		UserRouter.GET("list", api.GetGoodsList) // 获取用户列表
		// UserRouter.GET("", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetGoodsList) // 获取用户列表
	}
}
