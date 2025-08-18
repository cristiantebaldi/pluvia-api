package authrepository

import "github.com/pluvia/pluvia-api/core/domain"

// GetByAccountID implements domain.AuthRepository.
func (repository repository) GetByAccountID(id int32) (*domain.Auth, error) {
	panic("unimplemented")
}