package di

import (
	"github.com/go-playground/validator/v10"
	"github.com/pluvia/pluvia-api/adapter/http/adminservice"
	"github.com/pluvia/pluvia-api/adapter/repository/postgres"
	"github.com/pluvia/pluvia-api/adapter/repository/postgres/adminrepository"
	"github.com/pluvia/pluvia-api/adapter/repository/postgres/authrepository"
	"github.com/pluvia/pluvia-api/core/domain"
	"github.com/pluvia/pluvia-api/core/usecase/adminusecase"
)

func ConfigAdminDI(conn postgres.DatabaseConfig, validate *validator.Validate) domain.AdminHTTPService {
	adminRepository := adminrepository.New(conn)
	authRepository := authrepository.New(conn)
	adminUseCase := adminusecase.New(adminRepository, authRepository)
	adminService := adminservice.New(validate, adminUseCase)

	return adminService
}

func ConfigAdminDIUsecase(conn postgres.DatabaseConfig) domain.AdminUseCase {
	adminRepository := adminrepository.New(conn)
	authRepository := authrepository.New(conn)
	adminUseCase := adminusecase.New(adminRepository, authRepository)

	return adminUseCase
}
