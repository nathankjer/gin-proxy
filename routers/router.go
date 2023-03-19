package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/nathankjer/gin-proxy/controllers"
	"go.uber.org/ratelimit"
)

func InitRouterRequests() *gin.Engine {
	router := gin.Default()
	router.Any("/*path", controllers.Request)
	return router
}

func InitRouterGateway() *gin.Engine {
	router := gin.Default()
	limit := ratelimit.New(1) // Limit to 1 requests per second
	router.Use(func(ctx *gin.Context) { limit.Take() })
	router.Any("/*path", controllers.Proxy)
	return router
}
