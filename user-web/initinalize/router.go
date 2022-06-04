package initinalize

import (
	"github.com/gin-gonic/gin"

	"github.com/liuyongbing/hello-go-web/user-web/router"
)

func Routers() *gin.Engine {
	r := gin.Default()

	ApiGroup := r.Group("/v1")
	router.InitUserRouter(ApiGroup)

	return r
}
