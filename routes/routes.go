package routes

import (
	"database/sql"
	"golang-hexagonal/controllers"
	"golang-hexagonal/repositories"
	"golang-hexagonal/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	brandRepo := &repositories.BrandRepository{DB: db}
	brandService := &services.BrandService{Repo: brandRepo}
	brandController := &controllers.BrandController{Service: brandService}

	r.GET("/brands", brandController.GetAllBrands)
	r.GET("/brands/:id/models", brandController.GetModelsByBrandID)
	r.POST("/brands", brandController.CreateBrand)
	r.POST("/brands/:id/models", brandController.CreateModel)

	modelRepo := &repositories.ModelRepository{DB: db}
	modelService := &services.ModelService{Repo: modelRepo}
	modelController := &controllers.ModelController{Service: modelService}

	r.GET("/models", modelController.GetAllModelsWithGreaterOrLowerPrice)
	r.PUT("/models/:id", modelController.UpdateModelAveragePriceByID)

	return r
}
