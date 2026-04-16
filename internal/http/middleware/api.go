package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yasm3/prevently/internal/db"
	"github.com/yasm3/prevently/internal/domain"
	"github.com/yasm3/prevently/internal/security"
)

const UserContextKey = "user"

func APIKeyMiddleware(q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-Key")
		if apiKey == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Missing API key",
			})
			return
		}

		u, err := q.GetUserByAPIKey(c.Request.Context(), security.HashAPIKey(apiKey))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid API Key",
			})
		}

		user := domain.User{
			ID:    u.ID,
			Email: u.Email,
		}

		c.Set(UserContextKey, user)
		c.Next()
	}
}
