package di

import (
	"github.com/go-playground/validator/v10"
	"github.com/pluvia/pluvia-api/adapter/http/administradorservice"
	"github.com/pluvia/pluvia-api/adapter/repository/postgres"
	"github.com/pluvia/pluvia-api/adapter/repository/postgres/administradorrepository"
	"github.com/pluvia/pluvia-api/adapter/repository/postgres/authrepository"
	"github.com/pluvia/pluvia-api/core/domain"
	"github.com/pluvia/pluvia-api/core/usecase/administradorusecase"
)

func ConfigAdministradorDI(conn postgres.DatabaseConfig, validate *validator.Validate) domain.AdministradorHTTPService {
	administradorRepository := administradorrepository.New(conn)
	authRepository := authrepository.New(conn)
	administradorUseCase := administradorusecase.New(administradorRepository, authRepository)
	administradorService := administradorservice.New(validate, administradorUseCase)

	return administradorService
}

func ConfigAdministradorDIUsecase(conn postgres.DatabaseConfig) domain.AdministradorUseCase {
	administradorRepository := administradorrepository.New(conn)
	authRepository := authrepository.New(conn)
	administradorUseCase := administradorusecase.New(administradorRepository, authRepository)

	return administradorUseCase
}
