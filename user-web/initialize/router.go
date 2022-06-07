package initialize

import (
	"github.com/gin-gonic/gin"

	"github.com/liuyongbing/hello-go-web/user-web/middlewares"
	"github.com/liuyongbing/hello-go-web/user-web/router"
)

func Routers() *gin.Engine {
	r := gin.Default()

	// 配置跨域
	r.Use(middlewares.Cors())

	ApiGroup := r.Group("/u/v1")
	router.InitUserRouter(ApiGroup)

	return r
}
