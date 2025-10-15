package creators

import (
	"github.com/BooBooStory/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreatorRouter(api *gin.RouterGroup, db *gorm.DB) {
	// Inisialisasi semua lapisan
	creatorRepository := NewRepository(db)
	creatorService := NewService(creatorRepository)
	creatorHandler := NewCreatorHandler(creatorService)

	creatorRouter := api.Group("/creators")

	creatorRouter.Use(middleware.RequireAuth())
	{
		creatorRouter.GET("/me", creatorHandler.GetMyProfile)
	}


}