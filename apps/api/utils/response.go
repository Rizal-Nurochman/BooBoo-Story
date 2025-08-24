package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Struct response standar
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Meta    interface{} `json:"meta,omitempty"`
}

// Middleware untuk intercept response
func ResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// jalanin handler
		c.Next()

		// cek apakah handler sudah nulis response
		if c.Writer.Written() {
			return
		}

		// kalau handler belum kirim response, kasih default
		c.JSON(http.StatusOK, Response{
			Status:  "success",
			Message: "OK",
			Data:    nil,
		})
	}
}

// Helper biar handler lebih singkat
func JSON(c *gin.Context, status string, message string, data interface{}, meta interface{}) {
	c.JSON(http.StatusOK, Response{
		Status:  status,
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}