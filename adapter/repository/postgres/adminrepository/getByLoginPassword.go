package adminrepository

import (
	"database/sql"
	"fmt"

	"github.com/pluvia/pluvia-api/core/domain"
	"github.com/pluvia/pluvia-api/core/dto"
	"github.com/pluvia/pluvia-api/util"
)

func (repository repository) GetByLoginPassword(
	loginRequest *dto.AdminLoginRequestBody,
) (*domain.Admin, error) {
	admin := domain.Admin{}

	err := repository.db.Get(
		&admin,
		`SELECT * FROM admin where email = $1 and password = crypt($2, password);`,
		loginRequest.Email,
		loginRequest.Password,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("admin not exists")
	}

	if err != nil {
		return nil, util.GetError(err)
	}

	return &admin, nil
}
