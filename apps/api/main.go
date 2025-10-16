package main

import (
	"os"
	"time"

	"github.com/BooBooStory/config"
	"github.com/BooBooStory/config/database"
	"github.com/BooBooStory/middleware"
	"github.com/BooBooStory/utils"
	"github.com/BooBooStory/v1/auth"
	"github.com/BooBooStory/v1/categories"
	"github.com/BooBooStory/v1/story"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	database.ConnectDatabase()
	utils.InitLogger()
	middleware.InitMiddleware(database.DB)

	DB := database.DB
	router := gin.Default()

	corsConfig := cors.Config{
		AllowOrigins:    []string{config.Envs.FE_URL},
		AllowOriginFunc: func(origin string) bool {
			return origin == config.Envs.FE_URL
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	router.Use(cors.New(corsConfig))

	api := router.Group("/api/" + config.Envs.APP_VERSION)
	{
		auth.AuthRouter(api, DB)
		categories.CategoryRouter(api, DB)
		story.StoryRouter(api, DB)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	router.Run(":" + port)
}
