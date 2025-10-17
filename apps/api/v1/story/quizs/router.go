package quizs

import (
	"github.com/BooBooStory/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func QuizsRouter(router *gin.RouterGroup, db *gorm.DB) {
	quizRepository := NewQuizRepository(db)
	quizService := NewService(quizRepository)
	quizHandler := NewHandler(quizService)

	quizRouter := router.Group("/quizs")
	quizRouter.Use(middleware.RequireAuth())
	{
		quizRouter.GET("/", quizHandler.GetAll)
		quizRouter.GET("/:id", quizHandler.GetByID)
		quizRouter.POST("/", quizHandler.Create)
		quizRouter.DELETE("/:id", quizHandler.Delete)
	}
}