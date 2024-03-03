package main

import (
	"fmt"
	"log"

	"github.com/kriipke/console-api/pkg/initializers"
	"github.com/kriipke/console-api/pkg/models"
)

func init() {
	fmt.Println("starting init ")
	config, err := initializers.LoadConfig("configs")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Cluster{})
	fmt.Println("? Migration complete")
}
