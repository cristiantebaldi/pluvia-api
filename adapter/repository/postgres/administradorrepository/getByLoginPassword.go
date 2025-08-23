package administradorrepository

import (
	"database/sql"
	"fmt"

	"github.com/pluvia/pluvia-api/core/domain"
	"github.com/pluvia/pluvia-api/core/dto"
	"github.com/pluvia/pluvia-api/util"
)

func (repository repository) GetByLoginPassword(
	loginRequest *dto.AdministradorLoginRequestBody,
) (*domain.Administrador, error) {
	administrador := domain.Administrador{}

	err := repository.db.Get(
		&administrador,
		`SELECT * FROM administrador where email = $1 and senha = crypt($2, senha);`,
		loginRequest.Email,
		loginRequest.Senha,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("administrador not exists")
	}

	if err != nil {
		return nil, util.GetError(err)
	}

	return &administrador, nil
}
