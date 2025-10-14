package categories

import (
	"github.com/BooBooStory/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CategoryRouter(api *gin.RouterGroup, DB *gorm.DB,) {
	categoryRepo := NewRepository(DB)
	categoryService := NewService(categoryRepo)
	categoryHandler := NewHandler(categoryService)

	public := api.Group("/categories")
	{
		public.GET("/", categoryHandler.GetAll)
		public.GET("/:id", categoryHandler.GetByID)
		public.GET("/:id/children", categoryHandler.GetChildren)
		public.GET("/:id/parent", categoryHandler.GetParent)
	}

	protected := api.Group("/categories")
	protected.Use(middleware.RequireAuth(), middleware.RequireAuthorization("Admin", "Creator"))
	{
		protected.POST("/", categoryHandler.Create)
		protected.POST("/:id/subcategory", categoryHandler.CreateSubCategory)
		protected.PUT("/:id", categoryHandler.Update)
		protected.DELETE("/:id", categoryHandler.Delete)
	}
}
