package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/liuyongbing/hello-go-web/goods-web/api/goods"
	"github.com/liuyongbing/hello-go-web/goods-web/middlewares"
)

func InitGoodsRouter(Router *gin.RouterGroup) {
	// 多个模块都会使用 router 实例，因此全局化，让调用者传入使用
	// router := gin.Default()
	// router.GET("/ping")

	zap.S().Infof("配置 Goods API Url")

	// API: Goods
	GoodsRouter := Router.Group("goods")
	{
		GoodsRouter.GET("ping", goods.Pong)                                                  // Demo
		GoodsRouter.GET("", goods.List)                                                      // 列表
		GoodsRouter.POST("", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.Create) // 创建
		GoodsRouter.GET("/:id", goods.Detail)                                                // 详情
	}
}
