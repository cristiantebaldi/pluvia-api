package adminusecase

import (
	"github.com/pluvia/pluvia-api/core/domain"
	"github.com/pluvia/pluvia-api/core/dto"
)

func (usecase usecase) Create(admin *dto.AdminRequestBody) (*domain.Admin, error) {
	adminCreated, err := usecase.repository.Create(admin)

	if err != nil {
		return nil, err
	}

	return adminCreated, nil
}
