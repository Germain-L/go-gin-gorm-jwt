package middlewares

import (
	"net/http"
	"os"
	"server/security"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	secret = os.Getenv("secret")
)

func (m Middleware) Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header not provided"})
			return
		}

		splitToken := strings.Split(authHeader, "Bearer ")
		if len(splitToken) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token format"})
			return
		}

		claims, err := security.ValidateToken(splitToken[1], secret, security.AccessToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid access token"})
			return
		}

		// Store user information in the context
		c.Set("email", claims.Email)
		c.Next()
	}
}
