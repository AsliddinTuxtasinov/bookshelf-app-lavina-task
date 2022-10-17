package main

import (
	"bookshelf-app/initializers"
	"bookshelf-app/models"
	"fmt"
	"log"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	if err := initializers.DB.AutoMigrate(&models.Book{}); err != nil {
		log.Fatal(err)
	}
	if err := initializers.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal(err)
	}

	fmt.Println("All models migrated ...")
}
