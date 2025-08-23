package administradorrepository

import (
	"github.com/pluvia/pluvia-api/core/domain"
	"github.com/pluvia/pluvia-api/core/dto"
)

// Update implements domain.AdministradorRepository.
func (repository *repository) Update(int32, *dto.AdministradorRequestBody) (*domain.Administrador, error) {
	panic("unimplemented")
}
