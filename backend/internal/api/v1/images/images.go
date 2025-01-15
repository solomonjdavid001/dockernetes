package images

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/solomonjdavid001/Dockernetes/backend/internal/docker"
)

func GetImages(c *gin.Context) {
	imageList, err := docker.ListImages()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"images":  json.RawMessage(imageList),
	})
}
