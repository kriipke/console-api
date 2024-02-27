package main

import (
	"log"
	"fmt"
	"time"
	"net/http"	
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kriipke/console-api/pkg/controllers"
	"github.com/kriipke/console-api/pkg/initializers"
	"github.com/kriipke/console-api/pkg/routes"
)

var (
	server              *gin.Engine
	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController

	UserController      controllers.UserController
	UserRouteController routes.UserRouteController
)

func init() {
	config, err := initializers.LoadConfig("configs")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	AuthController = controllers.NewAuthController(initializers.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController(initializers.DB)
	UserRouteController = routes.NewRouteUserController(UserController)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig("configs")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://0.0.0.0:8080", config.ClientOrigin}
	//corsConfig.AllowAllOrigins = true
	//corsConfig.AllowCredentials = true
	// https://github.com/gin-gonic/gin/issues/2801 
	//corsConfig.AddAllowMethods("OPTIONS")

	//server.Use(cors.New(corsConfig))
	server.Use(corsMiddleware())
	server.Use(RequestLogger())
	server.Use(ResponseLogger())

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	AuthRouteController.AuthRoute(router)
	UserRouteController.UserRoute(router)
	log.Fatal(server.Run(":" + config.ServerPort))
}


// https://jwstanly.com/blog/article/How+to+solve+SAM+403+Errors+on+CORS+Preflight
func RequestLogger() gin.HandlerFunc {
    return func(c *gin.Context) {
        t := time.Now()

        c.Next()

        latency := time.Since(t)

        fmt.Printf("%s %s %s %s\n",
            c.Request.Method,
            c.Request.RequestURI,
            c.Request.Proto,
            latency,
        )
    }
}

func ResponseLogger() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("X-Content-Type-Options", "nosniff")

        c.Next()

        fmt.Printf("%d %s %s\n",
            c.Writer.Status(),
            c.Request.Method,
            c.Request.RequestURI,
        )
    }
}

func corsMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", 
				"http://localhost:3000, http://localhost:8080")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}

