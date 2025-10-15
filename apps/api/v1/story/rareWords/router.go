package rarewords

import (
	"github.com/BooBooStory/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RareWordsRouter(api *gin.RouterGroup, DB *gorm.DB) {
	rareWordRepo := NewRepository(DB)
	rareWordService := NewService(rareWordRepo)
	rareWordHandler := NewHandler(rareWordService)

	public := api.Group("/rare-words")
	{
		public.GET("/", rareWordHandler.GetAll)
	}

	protected := api.Group("/rare-words")
	protected.Use(middleware.RequireAuth(), middleware.RequireAuthorization("Admin", "Creator"))
	{
		protected.POST("/", rareWordHandler.Create)
		protected.PUT("/:id", rareWordHandler.Update)
		protected.DELETE("/:id", rareWordHandler.Delete)
	}	
}