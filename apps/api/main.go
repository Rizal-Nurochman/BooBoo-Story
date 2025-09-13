package main

import (
	"os"

	"github.com/BooBooStory/config/database"
	"github.com/BooBooStory/v1/users/handler"
	"github.com/BooBooStory/v1/users/repository"
	"github.com/BooBooStory/v1/users/service"
	"github.com/gin-gonic/gin"
)

func main() {
		database.ConnectDatabase()
		DB := database.DB
		
		// Domain User
		userRepository := repository.NewUserRepository(DB)
		userService := service.NewUserService(userRepository)
		userHandler := handler.NewUserHandler(userService)

		router := gin.Default()

		// Grouping route versi 1
		api := router.Group("/api/v1")
		{
			api.POST("/register", userHandler.RegisterHandler)
		}

		router.Run(":"+os.Getenv("PORT"))
}