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
	// String-related routes
	registerStringRoutes(group)

	// List-related routes
	registerListRoutes(group)

	// Hash-related routes
	registerHashRoutes(group)

	// Set-related routes
	registerSetRoutes(group)

	// Sorted Set-related routes
	registerSortedSetRoutes(group)

	// General routes
	group.GET("/get", GetAllKeys)
	group.POST("/set-expiry", SetRedisKeyExpiry)
	group.DELETE("/del/:key", DeleteRedisKey)
}

// String-related routes
func registerStringRoutes(group *gin.RouterGroup) {
	group.POST("/set/string", SetStringRedisValue)
	group.GET("/get/:key", GetRedisValue)
}

// List-related routes
func registerListRoutes(group *gin.RouterGroup) {
	group.POST("/set/list", SetListRedisValue)
	group.PUT("/list/update", UpdateElementInRedisList)
}

// Hash-related routes
func registerHashRoutes(group *gin.RouterGroup) {
	group.POST("/set/hash", SetHashRedisValue)
	group.DELETE("/hash/delete-key", DeleteElementInRedisHash)
}

// Set-related routes
func registerSetRoutes(group *gin.RouterGroup) {
	group.POST("/set/add-set", SetAddSetRedisValue)
	group.DELETE("/set/delete-element", DeleteElementInRedisSet)
}

// Sorted Set-related routes
func registerSortedSetRoutes(group *gin.RouterGroup) {
	group.POST("/set/add-sortedset", SetAddSortedSetRedisValue)
	group.DELETE("/sorted-set/delete-element", DeleteElementInRedisSortedSet)
}
