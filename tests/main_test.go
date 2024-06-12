package tests

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"

	"golang-hexagonal/database"
	"golang-hexagonal/routes"
)

func TestGetAllBrandsRouteOk(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("An error '%s' occurred while connecting to the test database", err)
	}
	defer db.Close()
	defer teardownTestDB(db)

	router := routes.SetupRouter(db)

	req, err := http.NewRequest("GET", "/brands", nil)
	if err != nil {
		t.Fatalf("An error occurred while creating request: %s", err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "expected status code 200, got %d", w.Code)
}

type Car struct {
	ID           int    `json:"id"`
	BrandName    string `json:"brand_name"`
	Name         string `json:"name"`
	AveragePrice int    `json:"average_price"`
}

func RunMigrations(db *sql.DB) {
	createTablesIfNotExist(db)

	jsonFile, err := os.Open("../models.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	var cars []Car
	if err := json.Unmarshal(byteValue, &cars); err != nil {
		log.Fatal(err)
	}

	for _, car := range cars {
		var brandID int
		err := db.QueryRow("SELECT id FROM brands WHERE name = $1", car.BrandName).Scan(&brandID)
		if err != nil {
			err = db.QueryRow("INSERT INTO brands (name) VALUES ($1) RETURNING id", car.BrandName).Scan(&brandID)
			if err != nil {
				log.Fatal(err)
			}
		}

		_, err = db.Exec("INSERT INTO models (name, average_price, brand_id) VALUES ($1, $2, $3)",
			car.Name, car.AveragePrice, brandID)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Datos insertados exitosamente")
}

func createTablesIfNotExist(db *sql.DB) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS brands (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) UNIQUE
		);
	`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS models (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255),
			average_price INTEGER,
			brand_id INTEGER REFERENCES brands(id)
		);
	`)
	if err != nil {
		log.Fatal(err)
	}
}

func setupTestDB() (*sql.DB, error) {
	err := godotenv.Load("../.env.test")
	if err != nil {
		log.Fatal("Error loading .env file in test mode")
	}

	db := database.Connect()

	if shouldRunMigrations := os.Getenv("RUN_MIGRATIONS"); shouldRunMigrations == "true" {
		fmt.Printf("Running migrations\n")
		RunMigrations(db)
	}

	return db, nil
}

func teardownTestDB(db *sql.DB) {
	db.Exec("TRUNCATE TABLE models CASCADE")
	db.Exec("TRUNCATE TABLE brands CASCADE")
}
