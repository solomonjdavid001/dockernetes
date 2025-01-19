package cosmos

import (
	"github.com/gin-gonic/gin"
)

func fetchDatabases(c *gin.Context) {
	dbNames := ListDatabases(c.Query("subscriptionId"), c.Query("resourceGroupName"), c.Query("accountName"))

	c.JSON(200, gin.H{
		"dbNames": dbNames,
		"message": "success",
	})
}
