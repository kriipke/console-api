package main

import (
	"log"
	"strings"
	"net/http"	
	//"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kriipke/console-api/pkg/controllers"
	"github.com/kriipke/console-api/pkg/initializers"
	"github.com/kriipke/console-api/pkg/routes"
	"github.com/kriipke/console-api/pkg/middleware"
)

var (
	server              *gin.Engine
	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController

	UserController      controllers.UserController
	UserRouteController routes.UserRouteController

	PostController      controllers.PostController
	PostRouteController routes.PostRouteController

	ClusterController      controllers.ClusterController
	ClusterRouteController routes.ClusterRouteController
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

	ClusterController = controllers.NewClusterController(initializers.DB)
	ClusterRouteController = routes.NewRouteClusterController(ClusterController)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig("configs")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	// https://github.com/gin-gonic/gin/issues/2801 
	// corsConfig := cors.DefaultConfig()
	// corsConfig.AllowCredentials = true
	// corsConfig.AllowOrigins = []string{"http://localhost:8080", "http://localhost:3000", config.ClientOrigin}
	// corsConfig.AllowMethods = []string{"OPTIONS", "POST", "GET", "PUT"}
	// corsConfig.AllowHeader = []string{"Content-Type", "Authorization", "Origin"}


	// server.Use(cors.New(corsConfig))

	server.Use(middleware.CORSMiddleware())
	server.Use(middleware.ResponseLogger())
	server.Use(middleware.RequestLogger())

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	AuthRouteController.AuthRoute(router)
	UserRouteController.UserRoute(router)
	PostRouteController.PostRoute(router)
	ClusterRouteController.ClusterRoute(router)
	log.Fatal(server.Run(":" + config.ServerPort))
}
