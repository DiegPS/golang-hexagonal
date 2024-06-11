package main

import (
	"fmt"
	"golang-hexagonal/database"
	"log"
	"os"

	"golang-hexagonal/routes"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	if os.Getenv("PROD") != "true" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	port := os.Getenv("PORT")

	db := database.Connect()
	defer db.Close()

	if shouldRunMigrations := os.Getenv("RUN_MIGRATIONS"); shouldRunMigrations == "true" {
		fmt.Printf("Running migrations\n")
		RunMigrations(db)
	}

	r := routes.SetupRouter(db)

	panic(r.Run(":" + port))
}
