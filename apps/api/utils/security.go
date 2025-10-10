package utils

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	configCORS "github.com/BooBooStory/config"
)

func LoadCors(router *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://127.0.0.1:3000", configCORS.Envs.FE_URL}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	router.Use(cors.New(config))
}