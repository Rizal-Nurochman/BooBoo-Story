package stories

import (
	"github.com/BooBooStory/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func StoriesRouter(api *gin.RouterGroup, DB *gorm.DB) {
	// story
	storyRepo := NewRepository(DB)
	storyService := NewService(storyRepo)
	storyHandler := NewHandler(storyService)

	public := api.Group("/stories")
	{
		public.GET("/", storyHandler.GetAll)
		public.GET("/:id", storyHandler.GetByID)
	}
	
	protected := api.Group("/stories")
	protected.Use(middleware.RequireAuth(), middleware.RequireAuthorization("Admin", "Creator"))
	{
		protected.POST("/", storyHandler.Create)
		protected.PUT("/:id", storyHandler.Update)
		protected.DELETE("/:id", storyHandler.Delete)
		protected.POST("/contents", )
	}
}