package cmd

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/solomonjdavid001/Dockernetes/backend/internal/api"
)

func StartServer() {
	router := gin.Default()
	router.Use(cors.Default())
	api.RegisterRoutes(router)

	fmt.Println("Server Running on Port: 8080")
	router.Run(":8080")
}
