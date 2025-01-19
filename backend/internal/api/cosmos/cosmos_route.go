package cosmos

import (
	"github.com/gin-gonic/gin"
)

func RegisterCosmosRoutes(group *gin.RouterGroup) {
	v1 := group.Group("/cosmos/v1")

	routes := []struct {
		path    string
		handler func(*gin.RouterGroup)
	}{
		{path: "/db", handler: registerDatabaseRoutes},
	}
	for _, route := range routes {
		routeGroup := v1.Group(route.path)
		route.handler(routeGroup)
	}
}

func registerDatabaseRoutes(group *gin.RouterGroup) {
	group.GET("", fetchDatabases)
}
