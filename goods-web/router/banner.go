package router

import (
	"github.com/gin-gonic/gin"

	"github.com/liuyongbing/hello-go-web/goods-web/api/banner"
)

func InitBannerRouter(Router *gin.RouterGroup) {
	BannerRouter := Router.Group("banner")
	{
		BannerRouter.GET("", banner.List) // 轮播图列表页

		// BannerRouter.POST("", middlewares.JWTAuth(), middlewares.IsAdminAuth(), banner.Create)       // 新建轮播图
		// BannerRouter.PUT("/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), banner.Update)    // 修改轮播图信息
		// BannerRouter.DELETE("/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), banner.Delete) // 删除轮播图

		BannerRouter.POST("", banner.Create)       // 新建轮播图
		BannerRouter.PUT("/:id", banner.Update)    // 修改轮播图信息
		BannerRouter.DELETE("/:id", banner.Delete) // 删除轮播图
	}
}
