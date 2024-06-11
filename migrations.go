package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Car struct {
	ID           int    `json:"id"`
	BrandName    string `json:"brand_name"`
	Name         string `json:"name"`
	AveragePrice int    `json:"average_price"`
}

func RunMigrations(db *sql.DB) {
	jsonFile, err := os.Open("models.json")
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
