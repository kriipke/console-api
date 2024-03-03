package main

import (
	"log"
	"time"
	"strings"
	"net/http"	
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kriipke/console-api/pkg/controllers"
	"github.com/kriipke/console-api/pkg/initializers"
	"github.com/kriipke/console-api/pkg/routes"
	"github.com/charmbracelet/lipgloss"
	"github.com/fatih/color"
	//"github.com/charmbracelet/lipgloss/table"

)

var style_keys = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("5"))
var style_values = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))

var (
	server              *gin.Engine
	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController

	UserController      controllers.UserController
	UserRouteController routes.UserRouteController

	PostController      controllers.PostController
	PostRouteController routes.PostRouteController
)
// function to convert an array to string
func arrayToString(arr []string) string {

   // seperating string elements with -
   return strings.Join([]string(arr), ", ")
}

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

	PostController = controllers.NewPostController(initializers.DB)
	PostRouteController = routes.NewRoutePostController(PostController)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig("configs")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	//consoleApiHost := os.Getenv("CONSOLE_API_HOST")
	//consoleApiPort := os.Getenv("CONSOLE_API_PORT")

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://0.0.0.0:8080", config.ClientOrigin}
	//corsConfig.AllowAllOrigins = true
	//corsConfig.AllowCredentials = true
	// https://github.com/gin-gonic/gin/issues/2801 
	//corsConfig.AddAllowMethods("OPTIONS")

	//server.Use(cors.New(corsConfig))
	server.Use(corsMiddleware())
	server.Use(ResponseLogger())
	server.Use(RequestLogger())

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	AuthRouteController.AuthRoute(router)
	UserRouteController.UserRoute(router)
	PostRouteController.PostRoute(router)
	log.Fatal(server.Run(":" + config.ServerPort))
}


// https://jwstanly.com/blog/article/How+to+solve+SAM+403+Errors+on+CORS+Preflight
func RequestLogger() gin.HandlerFunc {
    return func(c *gin.Context) {
        t := time.Now()

				 c1 := color.New(color.FgMagenta, color.Bold)
				 c2 := color.New(color.FgWhite)
				 c3 := color.New(color.FgYellow)
				 c4 := color.New(color.FgCyan, color.Bold)
				 c5 := color.New(color.FgWhite, color.Bold)

        c.Next()

				latency := time.Since(t)

        c4.Printf("\n  %s", c.Request.Method)
				c5.Printf("%s  ", c.Request.RequestURI)
				c1.Printf("%s  ", c.Request.Proto)
				c2.Printf("%s\n",  latency)

				for name, values := range c.Request.Header {
						for _, value := range values {
							c3.Printf("   %-20s:", name)
							c2.Printf("%-40s\n", value)
						}
				}
    }
}

func ResponseLogger() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("X-Content-Type-Options", "nosniff")

        c.Next()

				//c1 := color.New(color.FgMagenta, color.Bold)
				c2 := color.New(color.FgWhite)
				c3 := color.New(color.FgYellow)
			  c4 := color.New(color.FgCyan, color.Bold)
			  c5 := color.New(color.FgWhite, color.Bold)

        c4.Printf("\n   %s", c.Writer.Status())
				c4.Printf("%s  ", c.Request.Method)
				c5.Printf("%s\n",  c.Request.RequestURI)

				for name, values := range c.Writer.Header() {
						for _, value := range values {
							c3.Printf("    %-40s:", name)
							c2.Printf("%-80s\n", value)
						}
				}

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

