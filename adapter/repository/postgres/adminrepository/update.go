package adminrepository

import (
	"github.com/pluvia/pluvia-api/core/domain"
	"github.com/pluvia/pluvia-api/core/dto"
)

// Update implements domain.AdminRepository.
func (repository *repository) Update(int32, *dto.AdminRequestBody) (*domain.Admin, error) {
	panic("unimplemented")
}