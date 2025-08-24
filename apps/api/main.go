package main

import (
	"net/http"

	"github.com/BooBooStory/config/database"
	"github.com/gin-gonic/gin"
)

func helloHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello from the Golang API!")
}

func main() {
		database.ConnectDatabase()
		router := gin.Default()
		router.GET("/hello", helloHandler)
		router.Run(":8080")
}