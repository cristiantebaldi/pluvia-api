package administradorusecase

import "github.com/pluvia/pluvia-api/core/domain"

type usecase struct {
	repository     domain.AdministradorRepository
	authRepository domain.AuthRepository
}

func New(
	repository domain.AdministradorRepository,
	authRepository domain.AuthRepository,
) domain.AdministradorUseCase {
	return &usecase{
		repository:     repository,
		authRepository: authRepository,
	}
}
