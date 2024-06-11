package models

type ModelCar struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	AveragePrice float64 `json:"average_price"`
}

type Brand struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	AveragePrice float64 `json:"average_price"`
}
