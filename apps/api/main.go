package main

import (
	"os"

	"github.com/BooBooStory/config/database"
	"github.com/gin-gonic/gin"
)

func main() {
		database.ConnectDatabase()
		router := gin.Default()

		router.Run(":"+os.Getenv("PORT"))
}