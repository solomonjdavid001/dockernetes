package api

import (
	"github.com/gin-gonic/gin"
	"github.com/solomonjdavid001/Dockernetes/backend/internal/api/v1/containers"
	"github.com/solomonjdavid001/Dockernetes/backend/internal/api/v1/images"
	"github.com/solomonjdavid001/Dockernetes/backend/internal/api/v1/system"
)

func RegisterRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	registerV1Routes(v1)
}

func registerV1Routes(v1 *gin.RouterGroup) {
	registerImageRoutes(v1)
	registerContainerRoutes(v1)
	registerSystemRoutes(v1)
}

func registerImageRoutes(v1 *gin.RouterGroup) {
	imagesRoutes := v1.Group("/images")
	imagesRoutes.GET("", images.GetImages)
}

func registerContainerRoutes(v1 *gin.RouterGroup) {
	containersRoutes := v1.Group("/containers")
	containersRoutes.GET("", containers.GetContainers)
}

func registerSystemRoutes(v1 *gin.RouterGroup) {
	systemRoutes := v1.Group("/system")
	systemRoutes.GET("", system.GetSystemInfo)
}
