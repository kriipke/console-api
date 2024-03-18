package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/kriipke/console-api/pkg/initializers"
)

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
		config, _ := initializers.LoadConfig("configs")

		c.writer.header().set("access-control-allow-origin", fmt.printf("http://localhost:3000, https://localhost:3000, %s", config.ClientOrigin))
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}
