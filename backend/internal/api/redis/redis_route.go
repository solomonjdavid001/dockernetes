package redis

import (
	"github.com/gin-gonic/gin"
)

func RegisterRedisRoutes(group *gin.RouterGroup) {
	v1 := group.Group("/redis/v1")

	routes := []struct {
		path    string
		handler func(*gin.RouterGroup)
	}{
		{path: "/commands", handler: registerCommandRoutes},
	}
	for _, route := range routes {
		routeGroup := v1.Group(route.path)
		route.handler(routeGroup)
	}
}

func registerCommandRoutes(group *gin.RouterGroup) {
	group.POST("/set", SetRedisValue)
}
