package authrepository

import (
	"github.com/pluvia/pluvia-api/adapter/repository/postgres"
	"github.com/pluvia/pluvia-api/core/domain"
)

type repository struct {
	db postgres.DatabaseConfig
}

// Update implements domain.AuthRepository.
func (r *repository) Update(int32, domain.Auth) error {
	panic("unimplemented")
}

// New returns contract implementation of AuthRepository
func New(db postgres.DatabaseConfig) domain.AuthRepository {
	return &repository{
		db: db,
	}
}
