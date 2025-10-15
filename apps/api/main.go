package main

import (
	"os"

	"github.com/BooBooStory/config"
	"github.com/BooBooStory/config/database"
	"github.com/BooBooStory/middleware"
	"github.com/BooBooStory/utils"
	"github.com/BooBooStory/v1/auth"
	"github.com/BooBooStory/v1/categories"
	"github.com/BooBooStory/v1/story"
	"github.com/gin-gonic/gin"
)


func main() {
		database.ConnectDatabase()
		utils.InitLogger()
		middleware.InitMiddleware(database.DB)
		
		DB := database.DB
		
		config.LoadEnv()
		router := gin.Default()

		utils.LoadCors(router)

		api := router.Group("/api/v1")

		// all router
		auth.AuthRouter(api, DB)
		categories.CategoryRouter(api, DB)
		story.StoryRouter(api, DB)

		port := os.Getenv("PORT")
		router.Run(":" + port)
}