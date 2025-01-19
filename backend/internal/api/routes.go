package api

import (
	"github.com/gin-gonic/gin"
	"github.com/solomonjdavid001/Dockernetes/backend/internal/api/cosmos"
	"github.com/solomonjdavid001/Dockernetes/backend/internal/api/docker"
	"github.com/solomonjdavid001/Dockernetes/backend/internal/api/redis"
)

func RegisterRoutes(router *gin.Engine) {
	root := router.Group("/api")
	docker.RegisterDockerRoutes(root)
	cosmos.RegisterCosmosRoutes(root)
	redis.RegisterRedisRoutes(root)
}
