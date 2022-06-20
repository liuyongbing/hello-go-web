package initialize

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/liuyongbing/hello-go-web/goods-web/middlewares"
	"github.com/liuyongbing/hello-go-web/goods-web/router"
)

/*
pong
default response for get "/"
*/
func pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, Welcome to go web.",
		"time":    time.Now(),
	})
}

/*
health
服务的健康检查
*/
func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, service is healthly.",
		"time":    time.Now(),
	})
}

func Routers() *gin.Engine {
	r := gin.Default()
	r.GET("/", pong)
	r.GET("/health", health)

	// 配置跨域
	r.Use(middlewares.Cors())

	ApiGroup := r.Group("/g/v1")
	router.InitBannerRouter(ApiGroup)   // Banner
	router.InitBrandRouter(ApiGroup)    // Brand
	router.InitCategoryRouter(ApiGroup) // Category
	router.InitGoodsRouter(ApiGroup)    // Goods

	return r
}
