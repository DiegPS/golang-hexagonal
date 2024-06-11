package repositories

import (
	"database/sql"
	"golang-hexagonal/models"
)

type ModelRepository struct {
	DB *sql.DB
}

func (r *ModelRepository) GetAllModelsWithGreaterOrLowerPrice(greater, lower string) ([]models.ModelCar, error) {
	var rows *sql.Rows
	var err error
	if greater != "" && lower != "" {
		rows, err = r.DB.Query("SELECT id, name, average_price FROM models WHERE average_price > $1 AND average_price < $2", greater, lower)
	} else if greater != "" {
		rows, err = r.DB.Query("SELECT id, name, average_price FROM models WHERE average_price > $1", greater)
	} else if lower != "" {
		rows, err = r.DB.Query("SELECT id, name, average_price FROM models WHERE average_price < $1", lower)
	} else {
		rows, err = r.DB.Query("SELECT id, name, average_price FROM models")
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var modelsCar []models.ModelCar
	for rows.Next() {
		var modelCar models.ModelCar
		if err := rows.Scan(&modelCar.ID, &modelCar.Name, &modelCar.AveragePrice); err != nil {
			return nil, err
		}
		modelsCar = append(modelsCar, modelCar)
	}

	return modelsCar, nil
}

func (r *ModelRepository) UpdateModelAveragePriceByID(modelID, AveragePrice int) error {
	_, err := r.DB.Exec("UPDATE models SET average_price = $1 WHERE id = $2", AveragePrice, modelID)
	if err != nil {
		return err
	}
	return nil
}
