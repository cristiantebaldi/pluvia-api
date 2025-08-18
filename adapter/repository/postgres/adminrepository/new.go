package adminrepository

import (
	"github.com/pluvia/pluvia-api/adapter/repository/postgres"
	"github.com/pluvia/pluvia-api/core/domain"
)

type repository struct {
	db postgres.DatabaseConfig
}

// New returns contract implementation of AdminRepository
func New(db postgres.DatabaseConfig) domain.AdminRepository {
	return &repository{
		db: db,
	}
}
