package errors

import "errors"

var ErrBrandAlreadyExists = errors.New("la marca ya existe")
var ErrBrandNotFound = errors.New("la marca no existe")
var ErrModelAlreadyExists = errors.New("el modelo ya existe")
var ErrModelPriceTooLow = errors.New("el precio promedio debe ser mayor a 100,000")
