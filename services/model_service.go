package services

import (
	"golang-hexagonal/models"
	"golang-hexagonal/repositories"
)

type ModelService struct {
	Repo *repositories.ModelRepository
}

func (s *ModelService) GetAllModelsWithGreaterOrLowerPrice(greater, lower string) ([]models.ModelCar, error) {
	return s.Repo.GetAllModelsWithGreaterOrLowerPrice(greater, lower)
}

func (s *ModelService) UpdateModelAveragePriceByID(averagePrice, modelID int) error {
	return s.Repo.UpdateModelAveragePriceByID(averagePrice, modelID)
}
