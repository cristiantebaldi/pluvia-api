package administradorusecase

import (
	"github.com/pluvia/pluvia-api/core/domain"
	"github.com/pluvia/pluvia-api/core/dto"
)

func (usecase usecase) Update(
	id int32,
	administrador *dto.AdministradorRequestBody,
) (adminCreated *domain.Administrador, err error) {
	adminCreated, err = usecase.repository.Update(id, administrador)

	if err != nil {
		return nil, err
	}

	return adminCreated, nil
}
