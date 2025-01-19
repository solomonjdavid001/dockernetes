package redis

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetRedisValue(c *gin.Context) {

	var body SetRedisValueRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
			"error":   err.Error(),
		})
		return
	}

	err := SetValue(body.Key, body.Value)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error setting redis value",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}
