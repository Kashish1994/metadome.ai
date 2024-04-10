package middleware

import (
	"github.com/eduhub/helper"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// ValidateToken --> Middleware to make sure the API's are Authenticated
func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/metadome-api/user/login" || c.Request.URL.Path == "/metadome-api/user/" {
			c.Set("user", "Basic")
			c.Next()
			return
		}
		myHeader := c.GetHeader("Authorization")
		authHeader := strings.Split(myHeader, " ")
		if len(authHeader) == 0 || authHeader[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		}
		token, err := helper.DecodeToken(authHeader[1])
		if err != nil || token == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "could not decode token"})
		}
		sub, _ := token.Claims.GetSubject()
		c.Set("user", sub)
		c.Next()
	}
}
