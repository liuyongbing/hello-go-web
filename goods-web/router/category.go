package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/liuyongbing/hello-go-web/goods-web/api/category"
)

func InitCategoryRouter(Router *gin.RouterGroup) {

	zap.S().Infof("配置 Category API Url")

	// API: Category
	GoodsRouter := Router.Group("cate")
	{
		GoodsRouter.GET("/tree", category.Tree)  // 分类树
		GoodsRouter.GET("/:id", category.Detail) // 详情

		// GoodsRouter.POST("", middlewares.JWTAuth(), middlewares.IsAdminAuth(), category.Create)       // 创建
		// GoodsRouter.PUT("/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), category.Update)    // 修改
		// GoodsRouter.DELETE("/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), category.Delete) // 删除

		GoodsRouter.POST("", category.Create)       // 创建
		GoodsRouter.PUT("/:id", category.Update)    // 修改
		GoodsRouter.DELETE("/:id", category.Delete) // 删除
	}
}
