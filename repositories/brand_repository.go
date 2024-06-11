package repositories

import (
	"database/sql"
	"golang-hexagonal/models"
)

type BrandRepository struct {
	DB *sql.DB
}

func (r *BrandRepository) GetAllBrands() ([]models.Brand, error) {
	rows, err := r.DB.Query("SELECT brand_id AS id, (SELECT name FROM brands WHERE id = models.brand_id) AS nombre, AVG(average_price) AS average_price FROM models GROUP BY brand_id ORDER BY nombre;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var brands []models.Brand
	for rows.Next() {
		var brand models.Brand
		if err := rows.Scan(&brand.ID, &brand.Name, &brand.AveragePrice); err != nil {
			return nil, err
		}
		brands = append(brands, brand)
	}
	return brands, nil
}

func (r *BrandRepository) GetModelsByBrandID(brandID int) ([]models.ModelCar, error) {
	rows, err := r.DB.Query("SELECT id, name, average_price FROM models WHERE brand_id = $1", brandID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var modelsCars []models.ModelCar
	for rows.Next() {
		var modelCar models.ModelCar
		if err := rows.Scan(&modelCar.ID, &modelCar.Name, &modelCar.AveragePrice); err != nil {
			return nil, err
		}
		modelsCars = append(modelsCars, modelCar)
	}
	return modelsCars, nil
}

func (r *BrandRepository) GetBrandByName(name string) (models.Brand, error) {
	var brand models.Brand
	err := r.DB.QueryRow("SELECT id, name FROM brands WHERE name = $1", name).Scan(&brand.ID, &brand.Name)
	if err != nil {
		return models.Brand{}, err
	}
	return brand, nil
}

func (r *BrandRepository) CreateBrand(brand models.Brand) (int, error) {
	var brandID int
	err := r.DB.QueryRow("INSERT INTO brands (name) VALUES ($1) RETURNING id", brand.Name).Scan(&brandID)
	if err != nil {
		return 0, err
	}
	return brandID, nil
}

// r.POST("/brands/:id/models", func(c *gin.Context) {
// 	brandID := c.Param("id")
// 	var model struct {
// 		Name         string `json:"name" binding:"required"`
// 		AveragePrice int    `json:"average_price"`
// 	}
// 	if err := c.ShouldBindJSON(&model); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	var brandIDExists int
// 	err := db.QueryRow("SELECT id FROM brands WHERE id = $1", brandID).Scan(&brandIDExists)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "La marca no existe"})
// 		return
// 	}

// 	var modelID int
// 	err = db.QueryRow("SELECT id FROM models WHERE name = $1 AND brand_id = $2", model.Name, brandID).Scan(&modelID)
// 	if err == nil {
// 		c.JSON(http.StatusConflict, gin.H{"error": "El modelo ya existe"})
// 		return
// 	}

// 	if model.AveragePrice < 100000 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "El precio promedio debe ser mayor a 100,000"})
// 		return
// 	}

// 	_, err = db.Exec("INSERT INTO models (name, average_price, brand_id) VALUES ($1, $2, $3)",
// 		model.Name, model.AveragePrice, brandID)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al insertar modelo"})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, gin.H{"id": modelID})
// })

func (r *BrandRepository) CheckModelExists(name string, brandID string) (bool, error) {
	var modelID int
	err := r.DB.QueryRow("SELECT id FROM models WHERE name = $1 AND brand_id = $2", name, brandID).Scan(&modelID)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *BrandRepository) InsertModel(name string, averagePrice int, brandID string) (int, error) {
	var modelID int
	err := r.DB.QueryRow("INSERT INTO models (name, average_price, brand_id) VALUES ($1, $2, $3) RETURNING id",
		name, averagePrice, brandID).Scan(&modelID)
	if err != nil {
		return 0, err
	}
	return modelID, nil
}

func (r *BrandRepository) CheckBrandExists(brandID string) (bool, error) {
	var brandIDExists int
	err := r.DB.QueryRow("SELECT id FROM brands WHERE id = $1", brandID).Scan(&brandIDExists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
