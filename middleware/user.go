package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tests/lib"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(lib.KeyXToken)
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": http.StatusText(http.StatusUnauthorized)})
			c.Abort()
			return
		}

		claims, err := lib.NewJWT().ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": http.StatusText(http.StatusUnauthorized)})
			c.Abort()
			return
		}

		c.Set(lib.KeyXClaims, claims)
		c.Next()
	}
}
