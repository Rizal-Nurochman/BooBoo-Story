package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/BooBooStory/v1/users/usersrepository"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(userRepo usersrepository.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Ambil cookie dari request
		tokenString, err := c.Cookie("Authorization")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Token not provided"})
			return
		}

		// 2. Parse dan validasi token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Invalid token"})
			return
		}

		// 3. Ambil claims dan periksa apakah token valid
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			// Periksa waktu kedaluwarsa (exp)
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Token expired"})
				return
			}

			// 4. Ambil data user dari database berdasarkan ID dari token (sub)
			userID := uint(claims["user_id"].(float64))
			user, err := userRepo.FindByID(userID)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
				return
			}

			// 5. Lampirkan data user ke context request
			c.Set("user", user)

			// Lanjutkan ke handler berikutnya
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Invalid claims"})
		}
	}
}