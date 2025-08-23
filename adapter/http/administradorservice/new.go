package administradorservice

import (
	"github.com/go-playground/validator/v10"
	"github.com/pluvia/pluvia-api/core/domain"
)

type service struct {
	validator *validator.Validate
	usecase   domain.AdministradorUseCase
}

func New(validator *validator.Validate, usecase domain.AdministradorUseCase) domain.AdministradorHTTPService {
	return &service{
		validator: validator,
		usecase:   usecase,
	}
}
