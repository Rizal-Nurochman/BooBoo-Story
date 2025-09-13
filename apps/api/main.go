package main

import (
	"log"
	"os"

	"github.com/BooBooStory/config/database"
	"github.com/BooBooStory/v1/users/usershandler"
	"github.com/BooBooStory/v1/users/usersrepository"
	"github.com/BooBooStory/v1/users/usersservice"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauthConfig *oauth2.Config

func init() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	googleOauthConfig = &oauth2.Config{
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
}

func main() {
		database.ConnectDatabase()
		DB := database.DB
		
		// Domain User
		userRepository := usersrepository.NewUserRepository(DB)
		userService := usersservice.NewUserService(userRepository)
		userHandler := usershandler.NewUserHandler(userService)

		router := gin.Default()

		// Grouping route versi 1
		api := router.Group("/api/v1")
		{
			api.POST("/register", userHandler.RegisterHandler)
			api.POST("/login", userHandler.LoginHandler)

			// Rute baru untuk Google OAuth
			auth := api.Group("/auth/google")
			{
				auth.GET("/login", func(c *gin.Context) {
					userHandler.GoogleLoginHandler(c, googleOauthConfig)
				})
				auth.GET("/callback", func(c *gin.Context) {
					userHandler.GoogleCallbackHandler(c, googleOauthConfig)
				})
			}
		}

		port := os.Getenv("PORT")
		router.Run(":" + port)
}