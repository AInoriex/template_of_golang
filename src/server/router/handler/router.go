package handler

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(g *gin.Engine) {
	g.Use(Logger(), gin.Recovery())
	routerApi := g.Group("/api")
	{
		routerApi.POST("/hello_world", HelloWorld)
	}
}
