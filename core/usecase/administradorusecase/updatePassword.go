package administradorusecase

import (
	"github.com/pluvia/pluvia-api/core/domain"
	"github.com/pluvia/pluvia-api/core/dto"
)

func (usecase usecase) UpdatePassword(
	id int32,
	administrador *dto.AdministradorRequestBody,
) (*domain.Administrador, error) {
	adminCreated, err := usecase.repository.UpdatePassword(id, administrador)

	if err != nil {
		return nil, err
	}

	return adminCreated, nil
}
