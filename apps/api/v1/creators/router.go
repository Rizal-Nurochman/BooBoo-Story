package creators

import (
	"github.com/BooBooStory/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreatorRouter(api *gin.RouterGroup, DB *gorm.DB) {
	creatorRepo := NewRepository(DB)
	creatorService := NewService(creatorRepo)
	creatorHandler := NewHandler(creatorService)

	public := api.Group("/creators")
	{
		public.GET("/", creatorHandler.GetAll)
		public.GET("/:id", creatorHandler.GetByID)
	}

	protected := api.Group("/creators")
	protected.Use(middleware.RequireAuth(), middleware.RequireAuthorization("Admin", "Creator"))
	{
		protected.POST("/", creatorHandler.Create)
		protected.PUT("/:id", creatorHandler.Update)
		protected.DELETE("/:id", creatorHandler.Delete)
	}
}