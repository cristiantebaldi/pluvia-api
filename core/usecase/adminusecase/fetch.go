package adminusecase

import "github.com/pluvia/pluvia-api/core/domain"

func (usecase usecase) Fetch(
	admin domain.Admin,
) (*[]domain.Admin, error) {
	adminCreated, err := usecase.repository.Fetch()

	if err != nil {
		return nil, err
	}

	return adminCreated, nil
}
