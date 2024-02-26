package main

import (
	"fmt"
	"log"

	"github.com/kriipke/console-api/initializers"
	"github.com/kriipke/console-api/models"
)

func init() {
	fmt.Println("starting init ")
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
	fmt.Println("? Migration complete")
}
