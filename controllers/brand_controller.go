package controllers

import (
	"golang-hexagonal/models"
	"golang-hexagonal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BrandController struct {
	Service *services.BrandService
}

func (ctrl *BrandController) GetAllBrands(c *gin.Context) {
	brands, err := ctrl.Service.GetAllBrands()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, brands)
}

func (ctrl *BrandController) GetModelsByBrandID(c *gin.Context) {
	brandID := c.Param("id")
	brandIDInt, err := strconv.Atoi(brandID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid brand ID"})
		return
	}
	models, err := ctrl.Service.GetModelsByBrandID(brandIDInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, models)
}

func (ctrl *BrandController) CreateBrand(c *gin.Context) {
	var brand models.Brand
	if err := c.ShouldBindJSON(&brand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	brandID, err := ctrl.Service.CreateBrand(brand)
	if err != nil {
		println(err.Error())

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": brandID})
}

func (ctrl *BrandController) CreateModel(c *gin.Context) {
	brandID := c.Param("id")
	var model models.ModelCar
	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	modelID, err := ctrl.Service.CreateModel(brandID, model)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": modelID})
}
