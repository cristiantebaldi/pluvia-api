package administradorrepository

import (
	"github.com/pluvia/pluvia-api/adapter/repository/postgres"
	"github.com/pluvia/pluvia-api/core/domain"
)

type repository struct {
	db postgres.DatabaseConfig
}

// New returns contract implementation of AdministradorRepository
func New(db postgres.DatabaseConfig) domain.AdministradorRepository {
	return &repository{
		db: db,
	}
}
