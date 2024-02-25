package main

import (
	"fmt"
	"log"

	"github.com/kriipke/golang-gorm-postgres/initializers"
	"github.com/kriipke/golang-gorm-postgres/models"
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
