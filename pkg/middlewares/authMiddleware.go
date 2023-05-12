package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("X-Auth-Token")
		if tokenString == "" {
			log.Print("empty token")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		// jwt is in tokenString

		// do rpc call

		c.Next()
	}
}
