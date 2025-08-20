package adminusecase

import "github.com/pluvia/pluvia-api/core/domain"

func (usecase usecase) GetByID(id int32) (*domain.Admin, error) {
	admin, err := usecase.repository.GetByID(id)

	if err != nil {
		return nil, err
	}

	return admin, nil
}
