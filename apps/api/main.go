package main

import (
	"os"

	"github.com/BooBooStory/config"
	"github.com/BooBooStory/config/database"
	"github.com/BooBooStory/v1/auth"
	"github.com/gin-gonic/gin"
)



func main() {
		database.ConnectDatabase()
		DB := database.DB
		
		config.LoadEnv()
		router := gin.Default()

		api := router.Group("/api/v1")

		//all router
		auth.AuthRouter(api, DB)

		port := os.Getenv("PORT")
		router.Run(":" + port)
}