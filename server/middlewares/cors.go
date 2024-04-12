package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	)

	func CORSMiddleware() gin.HandlerFunc{
		config := cors.DefaultConfig()
		config.AllowOrigins = []string{"*"} // Permitir solicitações de qualquer origem
		config.AllowCredentials = true
		config.AllowHeaders = []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"}
		config.AllowMethods = []string{"POST", "OPTIONS", "GET", "PUT", "DELETE"}

		// Retorna o middleware CORS
		return cors.New(config)
	}

	func Cors(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
