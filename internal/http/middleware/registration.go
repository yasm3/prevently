package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
)

var allowRegistration = os.Getenv("ALLOW_REGISTRATION") == "true"

func RegistrationEnabled() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !allowRegistration {
			c.AbortWithStatusJSON(403, gin.H{
				"error": "registration disabled",
			})
		}

		c.Next()
	}
}
