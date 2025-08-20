package adminusecase

import "github.com/pluvia/pluvia-api/core/domain"

type usecase struct {
	repository     domain.AdminRepository
	authRepository domain.AuthRepository
}

func New(
	repository domain.AdminRepository,
	authRepository domain.AuthRepository,
) domain.AdminUseCase {
	return &usecase{
		repository:     repository,
		authRepository: authRepository,
	}
}
