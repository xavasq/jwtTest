package middlewares

import (
	"ECCO2K/internal/security"
	"github.com/gin-gonic/gin"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		if tokenString == authHeader || tokenString == "" {
			c.JSON(401, gin.H{"error": "неверный токен"})
			c.Abort()
			return
		}

		claims, err := security.ValidateToken(tokenString)
		if err != nil {
			c.JSON(401, gin.H{"error": "токен не валидный"})
			c.Abort()
			return
		}

		if claims.UserID == 0 || claims.Name == "" {
			c.JSON(401, gin.H{"error": "недостаточно данных в токене"})
			c.Abort()
			return
		}

		c.Set("UserID", claims.UserID)
		c.Set("Name", claims.Name)

		c.Next()
	}
}
