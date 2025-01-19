package docker

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetImages(c *gin.Context) {
	imageList, err := FetchImages()
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

func GetContainers(c *gin.Context) {
	containerList, err := FetchContainers()
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

func GetSystemInfo(c *gin.Context) {
	systemInfo := SystemInfo{}

	// Get CPU Usage
	systemInfo.CalculateCpuUsage()

	// Get RAM Usage
	systemInfo.CalculateRamUsage()

	c.JSON(http.StatusOK, gin.H{
		"systemInfo": systemInfo,
	})
}
