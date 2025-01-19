package cmd

import (
	"embed"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/solomonjdavid001/Dockernetes/backend/config"
	"github.com/solomonjdavid001/Dockernetes/backend/internal/api"
	"github.com/solomonjdavid001/Dockernetes/backend/pkg"
)

func StartServer(configFile embed.FS) {
	err := config.LoadConfig(configFile)
	if err != nil {
		fmt.Println("Error loading config file:", err)
	}
	pkg.LoadFiglet()

	router := gin.Default()
	router.Use(cors.Default())
	api.RegisterRoutes(router)

	PORT := config.GlobalConfig.App.Port
	fmt.Printf("Server Running on Port: %s\n", PORT)
	router.Run(PORT)
}
