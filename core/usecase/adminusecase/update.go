package adminusecase

import (
	"github.com/pluvia/pluvia-api/core/domain"
	"github.com/pluvia/pluvia-api/core/dto"
)

func (usecase usecase) Update(
	id int32,
	admin *dto.AdminRequestBody,
) (adminCreated *domain.Admin, err error) {
	currentAdmin, err := usecase.repository.GetByID(id)

	if err != nil {
		return nil, err
	}

	adminCreated, err = usecase.repository.Update(id, admin)
	
	if err != nil {
		return nil, err
	}

	if admin.Enable != currentAdmin.Enable {
		auth, err := usecase.authRepository.GetByAdminID(id)
		if err != nil {
			if err.Error() == "Token not created yet" {
				return adminCreated, nil
			} else {
				return nil, err
			}
		}

		auth.Revoked = !admin.Enable

		err = usecase.authRepository.Update(auth.ID, *auth)

		if err != nil {
			return nil, err
		}
	}

	return adminCreated, nil
}
