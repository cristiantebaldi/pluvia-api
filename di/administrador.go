package di

import (
	"github.com/pluvia/pluvia-api/adapter/repository/postgres"
	"github.com/pluvia/pluvia-api/core/domain"
	"github.com/pluvia/pluvia-api/core/usecase"
)

func ConfigAdministradorDI(conn postgres.DatabaseConfig) domain.AdministradorUseCase {
	administradorRepository := postgres.NewAdministradorRepository(conn)
	authRepository := postgres.NewAuthRepository(conn)
	administradorUseCase := usecase.NewAdministradorUseCase(administradorRepository, authRepository)

	return administradorUseCase
}
