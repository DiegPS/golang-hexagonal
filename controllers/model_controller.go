package controllers

import (
	"golang-hexagonal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ModelController struct {
	Service *services.ModelService
}

func (ctrl *ModelController) GetAllModelsWithGreaterOrLowerPrice(c *gin.Context) {
	greater := c.Query("greater")
	lower := c.Query("lower")

	models, err := ctrl.Service.GetAllModelsWithGreaterOrLowerPrice(greater, lower)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, models)
}

func (ctrl *ModelController) UpdateModelAveragePriceByID(c *gin.Context) {
	modelID := c.Param("id")
	modelIDInt, err := strconv.Atoi(modelID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid model ID"})
		return
	}

	var model struct {
		AveragePrice int `json:"average_price" binding:"required"`
	}
	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if model.AveragePrice < 100000 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El precio promedio debe ser mayor a 100,000"})
		return
	}

	err = ctrl.Service.UpdateModelAveragePriceByID(model.AveragePrice, modelIDInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": modelID})
}
