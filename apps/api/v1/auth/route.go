package auth

import (
	"github.com/BooBooStory/config"
	"github.com/BooBooStory/v1/users"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gorm.io/gorm"
)

func AuthRouter(api *gin.RouterGroup, DB *gorm.DB) {
	googleOauthConfig := &oauth2.Config{
		RedirectURL:  config.Envs.GoogleRedirect,
		ClientID:     config.Envs.GoogleClientID,
		ClientSecret: config.Envs.GoogleSecret,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	userRepo := users.NewRepository(DB)
	authService := NewService(userRepo)
	authHandler := NewHandler(authService, googleOauthConfig)

	authGroup := api.Group("/auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
		authGroup.POST("/logout", authHandler.Logout)
		authGroup.GET("/google/login", authHandler.GoogleLoginHandler)
		authGroup.GET("/google/callback", authHandler.GoogleCallbackHandler)
	}
}
