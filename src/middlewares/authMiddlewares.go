package middlewares

import (
	"strings"

	"segmenta/src/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.SendErrorResponse(c, "Authorization header missing", 401)
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			utils.SendErrorResponse(c, "Invalid authorization format, expected 'Bearer <token>'", 401)
			c.Abort()
			return
		}

		tokenString := authHeader[len("Bearer "):]
		claims, errorHandler := utils.ValidateJWT(tokenString)
		if errorHandler != nil {
			utils.SendErrorResponse(c, "Invalid or expired token", 401)
			c.Abort()
			return
		}

		c.Set("user_id", claims["user_id"])
		c.Set("email", claims["email"])
		c.Set("name", claims["name"])

		c.Next()
	}
}