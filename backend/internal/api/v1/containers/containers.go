package containers

import (
	"github.com/gin-gonic/gin"
	"github.com/solomonjdavid001/Dockernetes/backend/internal/docker"
)

func GetContainers(c *gin.Context) {
	containerList, err := docker.GetContainers()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"containers": containerList,
	})
}
