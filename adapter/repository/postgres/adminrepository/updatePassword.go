package adminrepository

import (
	"github.com/pluvia/pluvia-api/core/domain"
	"github.com/pluvia/pluvia-api/core/dto"
	"github.com/pluvia/pluvia-api/util"
)

func (repository repository) UpdatePassword(
	id int32,
	admin *dto.AdminRequestBody,
) (*domain.Admin, error) {
	adminUpdated := domain.Admin{}

	err := repository.db.QueryRowx(
		`UPDATE admin SET 
			password = crypt($1, gen_salt('bf', 8)), 
			change_password = false 
		WHERE id = $2 returning *;`,
		admin.Password,
		id,
	).StructScan(&adminUpdated)

	if err != nil {
		return nil, util.GetError(err)
	}

	return &adminUpdated, nil
}
