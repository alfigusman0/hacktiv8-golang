package main

import (
	"log"
	"os"

	"final_project/pkg/database"
	"final_project/pkg/routers"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
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
