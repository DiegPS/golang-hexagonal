package services

import (
	"database/sql"
	"golang-hexagonal/errors"
	"golang-hexagonal/models"
	"golang-hexagonal/repositories"
)

type BrandService struct {
	Repo *repositories.BrandRepository
}

func (s *BrandService) GetAllBrands() ([]models.Brand, error) {
	return s.Repo.GetAllBrands()
}

func (s *BrandService) GetModelsByBrandID(brandID int) ([]models.ModelCar, error) {
	return s.Repo.GetModelsByBrandID(brandID)
}

func (s *BrandService) CreateBrand(brand models.Brand) (int, error) {
	_, err := s.Repo.GetBrandByName(brand.Name)
	if err == nil {
		return 0, errors.ErrBrandAlreadyExists
	} else if err != sql.ErrNoRows {
		return 0, err
	}

	brandID, err := s.Repo.CreateBrand(brand)
	if err != nil {
		return 0, err
	}

	return brandID, nil
}

func (s *BrandService) CreateModel(brandID string, model models.ModelCar) (int, error) {
	brandExists, err := s.Repo.CheckBrandExists(brandID)
	if err != nil {
		return 0, err
	}
	if !brandExists {
		return 0, errors.ErrBrandNotFound
	}

	modelExists, err := s.Repo.CheckModelExists(model.Name, brandID)
	if err != nil {
		return 0, err
	}
	if modelExists {
		return 0, errors.ErrModelAlreadyExists
	}

	if model.AveragePrice < 100000 {
		return 0, errors.ErrModelPriceTooLow
	}

	modelID, err := s.Repo.InsertModel(model.Name, int(model.AveragePrice), brandID)
	if err != nil {
		return 0, err
	}
	return modelID, nil
}
