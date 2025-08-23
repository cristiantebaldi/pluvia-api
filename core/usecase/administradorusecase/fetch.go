package administradorusecase

import "github.com/pluvia/pluvia-api/core/domain"

func (usecase usecase) Fetch(
	administrador domain.Administrador,
) (*[]domain.Administrador, error) {
	adminCreated, err := usecase.repository.Fetch()

	if err != nil {
		return nil, err
	}

	return adminCreated, nil
}
