package middleware

import (
	"net/http"
	"strings"

	"github.com/Sina-Mahmoodmoradi/task-manager/internal/core/port"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtManager port.TokenManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			return
		}

		tokenString := parts[1]
		userId, err := jwtManager.ParseToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// Set user ID from token in context
		c.Set("userID", userId)

		c.Next()
	}
}
