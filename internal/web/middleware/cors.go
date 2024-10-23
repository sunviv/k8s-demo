package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func CORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		//AllowAllOrigins: true
		//AllowOrigins: []string{"*"},
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				return true
			}
			return strings.Contains(origin, "sun.dev.com")
		},
		AllowCredentials: true,
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"jwt"},
		//AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		MaxAge: 12 * time.Hour,
	})
}
