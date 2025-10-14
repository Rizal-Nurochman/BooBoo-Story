package categories

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CategoryRouter mengatur semua endpoint kategori
func CategoryRouter(api *gin.RouterGroup, DB *gorm.DB) {
	categoryRepo := NewRepository(DB)
	categoryService := NewService(categoryRepo)
	categoryHandler := NewHandler(categoryService)

	categoryGroup := api.Group("/categories")
	{
		// CRUD utama
		categoryGroup.POST("/", categoryHandler.Create)
		categoryGroup.GET("/", categoryHandler.GetAll)
		categoryGroup.GET("/:id", categoryHandler.GetByID)
		categoryGroup.PUT("/:id", categoryHandler.Update)
		categoryGroup.DELETE("/:id", categoryHandler.Delete)

		// Hierarki (parent-child)
		categoryGroup.GET("/:id/children", categoryHandler.GetChildren)
		categoryGroup.GET("/:id/parent", categoryHandler.GetParent)
		categoryGroup.POST("/:id/subcategory", categoryHandler.CreateSubCategory)
	}
}
