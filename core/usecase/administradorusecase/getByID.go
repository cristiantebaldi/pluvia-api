package administradorusecase

import "github.com/pluvia/pluvia-api/core/domain"

func (usecase usecase) GetByID(id int32) (*domain.Administrador, error) {
	administrador, err := usecase.repository.GetByID(id)

	if err != nil {
		return nil, err
	}

	return administrador, nil
}
