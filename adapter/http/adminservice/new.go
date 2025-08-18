package adminservice

import (
	"github.com/go-playground/validator/v10"
	"github.com/pluvia/pluvia-api/core/domain"
)

type service struct {
	validator *validator.Validate
	usecase   domain.AdminUseCase
}


func New(validator *validator.Validate, usecase domain.AdminUseCase) domain.AdminHTTPService {
	return &service{
		validator: validator,
		usecase:   usecase,
	}
}
