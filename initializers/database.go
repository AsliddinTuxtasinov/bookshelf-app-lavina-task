package initializers

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {

	var err error
	dsn := "host=localhost user=postgres password=postgres dbname=test port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Database connection error: %v", err)
		panic(err)
	}

	fmt.Println("Connect to database ...")
}
