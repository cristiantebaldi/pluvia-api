package adminusecase

import "github.com/pluvia/pluvia-api/core/domain"

func (usecase usecase) Delete(id int32) (*domain.Admin, error) {
	adminDeleted, err := usecase.repository.Delete(id)

	if err != nil {
		return nil, err
	}

	return adminDeleted, nil
}
