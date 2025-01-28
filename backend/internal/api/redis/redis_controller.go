package redis

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetStringRedisValue(c *gin.Context) {

	var body SetRedisStringValueRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
			"error":   err.Error(),
		})
		return
	}

	err := SetStringValue(body.Key, body.Value)
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
func SetListRedisValue(c *gin.Context) {

	var body SetRedisListRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
			"error":   err.Error(),
		})
		return
	}

	err := SetListValue(body)
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

func SetHashRedisValue(c *gin.Context) {

	var body SetRedisHashRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
			"error":   err.Error(),
		})
		return
	}

	err := SetHashValue(body)
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

func SetAddSetRedisValue(c *gin.Context) {

	var body SetRedisSetRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
			"error":   err.Error(),
		})
		return
	}

	err := SetAddSetValue(body)
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

func SetAddSortedSetRedisValue(c *gin.Context) {

	var body SetRedisSortedSetRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
			"error":   err.Error(),
		})
		return
	}

	err := SetAddSortedSetValue(body)
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

func UpdateElementInRedisList(c *gin.Context) {

	var body UpdateRedisListRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
			"error":   err.Error(),
		})
		return
	}

	err := UpdateRedisList(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error updating redis value",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func DeleteElementInRedisSet(c *gin.Context) {

	key := c.Query("key")
	member := c.Query("member")

	err := DeleteRedisSet(key, member)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error deleting redis set",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func DeleteElementInRedisSortedSet(c *gin.Context) {

	key := c.Query("key")
	member := c.Query("member")

	err := DeleteRedisSortedSet(key, member)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error deleting redis sorted set",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func DeleteElementInRedisHash(c *gin.Context) {

	key := c.Query("key")
	member := c.Query("member")

	err := DeleteRedisHash(key, member)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error deleting redis hash",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func GetRedisValue(c *gin.Context) {
	key := c.Param("key")

	val, err := GetValue(key)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "error getting redis value",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"value": val,
	})
}

func GetAllKeys(c *gin.Context) {

	keyInfo, err := GetAll()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"keys":    json.RawMessage(keyInfo),
	})
}

func DeleteRedisKey(c *gin.Context) {

	key := c.Param("key")
	err := DeleteKey(key)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func SetRedisKeyExpiry(c *gin.Context) {

	var body SetRedisExpiryRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
			"error":   err.Error(),
		})
		return
	}
	
	err := SetKeyExpiry(body)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}
