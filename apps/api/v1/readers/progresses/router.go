package progresses

import (
	"github.com/BooBooStory/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProgressReadRouter(rg *gin.RouterGroup, db *gorm.DB) {
	progressRepo := NewRepository(db)
	progressService := NewService(progressRepo)
	progressHandler := NewHandler(progressService)

	public := rg.Group("/progresses")
	{
		public.GET("/", progressHandler.GetAll)
		public.GET("/:id", progressHandler.GetByID)
	}

	protected := rg.Group("/progresses")
	protected.Use(middleware.RequireAuth())
	{
		protected.POST("/", progressHandler.Create)
		protected.PUT("/:id", progressHandler.Update)
		protected.DELETE("/:id", progressHandler.Delete)
	}
}