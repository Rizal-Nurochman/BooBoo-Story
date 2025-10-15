package creators

import (
	"github.com/BooBooStory/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreatorRouter(router *gin.RouterGroup, db *gorm.DB) {
	// Inisialisasi semua lapisan
	creatorRepository := NewRepository(db)
	creatorService := NewService(creatorRepository)
	creatorHandler := NewCreatorHandler(creatorService)

	// Grup utama untuk /api/v1/creators
	creatorRouter := router.Group("/api/v1/creators")

	// 1. Grup untuk kreator mengelola profil PRIBADI
	profileRouter := creatorRouter.Group("/profile")
	// Semua rute di sini memerlukan pengguna untuk login
	profileRouter.Use(middleware.RequireAuth())
	{
		profileRouter.GET("/me", creatorHandler.GetMyProfile)
		profileRouter.PUT("/me", creatorHandler.UpdateMyProfile)
	}

	// 2. Grup untuk ADMIN mengelola data kreator
	// Rute ini juga memerlukan login, dan idealnya memerlukan cek peran "admin"
	// Contoh: adminRouter.Use(middleware.RequireAuthorization("admin"))
	adminRouter := creatorRouter.Group("/")
	adminRouter.Use(middleware.RequireAuth())
	{
		// :id di sini adalah ID dari tabel creators, bukan user_id
		adminRouter.GET("/:id", creatorHandler.GetCreatorByID)
		adminRouter.DELETE("/:id", creatorHandler.DeleteCreatorByID)
		// Anda juga bisa menambahkan PUT /:id di sini untuk admin mengedit profil kreator lain
	}
}