package bookmarks

import (
	"github.com/BooBooStory/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BookmarkRouter(router *gin.RouterGroup, db *gorm.DB) {
	bookmarkRepository := NewRepository(db)
	bookmarkService := NewService(bookmarkRepository)
	bookmarkHandler := NewBookmarkHandler(bookmarkService)

	bookmarkRouter := router.Group("/api/v1/story-bookmarks")
	bookmarkRouter.Use(middleware.RequireAuth())
	{
		bookmarkRouter.POST("/", bookmarkHandler.CreateBookmark)
		bookmarkRouter.GET("/", bookmarkHandler.GetBookmarks)
		bookmarkRouter.GET("/:id", bookmarkHandler.GetBookmarkByID)
		bookmarkRouter.DELETE("/:id", bookmarkHandler.DeleteBookmark)
	}
}