package docker

import (
	"github.com/gin-gonic/gin"
)

func RegisterDockerRoutes(group *gin.RouterGroup) {
	v1 := group.Group("/docker/v1")

	routes := []struct {
		path    string
		handler func(*gin.RouterGroup)
	}{
		{path: "/images", handler: registerImageRoutes},
		{path: "/containers", handler: registerContainerRoutes},
		{path: "/system", handler: registerSystemRoutes},
	}
	for _, route := range routes {
		routeGroup := v1.Group(route.path)
		route.handler(routeGroup)
	}
}

func registerImageRoutes(group *gin.RouterGroup) {
	group.GET("", GetImages)
}

func registerContainerRoutes(group *gin.RouterGroup) {
	group.GET("", GetContainers)
}

func registerSystemRoutes(group *gin.RouterGroup) {
	group.GET("", GetSystemInfo)
}
