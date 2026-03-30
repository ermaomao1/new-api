package middleware

import (
	"github.com/QuantumNous/new-api/common"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	config := cors.DefaultConfig()
	// SECURITY FIX: AllowAllOrigins=false prevents CSRF attacks
	config.AllowAllOrigins = false
	config.AllowOrigins = []string{
		"https://tokene.top",
		"https://www.tokene.top",
		"http://localhost:3000",
		"http://127.0.0.1:3000",
	}
	config.AllowCredentials = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"*"}
	return cors.New(config)
}

func PoweredBy() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-New-Api-Version", common.Version)
		c.Next()
	}
}
