package administradorusecase

import (
	"github.com/pluvia/pluvia-api/core/domain"
	"github.com/pluvia/pluvia-api/core/dto"
)

func (usecase usecase) Create(administrador *dto.AdministradorRequestBody) (*domain.Administrador, error) {
	adminCreated, err := usecase.repository.Create(administrador)

	if err != nil {
		return nil, err
	}

	return adminCreated, nil
}
