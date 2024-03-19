package main

import (
	"log"
	"os"

	"assignment_2/pkg/database"
	"assignment_2/pkg/routers"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Load database
	db := database.GetConnection()
	gorm := database.GormInit(db)
	defer db.Close()

	port := ":" + os.Getenv("APP_PORT")
	start := routers.StartServer(db, gorm)
	start.Run(port)
}
