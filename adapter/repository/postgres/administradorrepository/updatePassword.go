package administradorrepository

import (
	"github.com/pluvia/pluvia-api/core/domain"
	"github.com/pluvia/pluvia-api/core/dto"
	"github.com/pluvia/pluvia-api/util"
)

func (repository repository) UpdatePassword(
	id int32,
	administrador *dto.AdministradorRequestBody,
) (*domain.Administrador, error) {
	adminUpdated := domain.Administrador{}

	err := repository.db.QueryRowx(
		`UPDATE administrador SET 
			senha = crypt($1, gen_salt('bf', 8)), 
			change_password = false 
		WHERE id = $2 returning *;`,
		administrador.Senha,
		id,
	).StructScan(&adminUpdated)

	if err != nil {
		return nil, util.GetError(err)
	}

	return &adminUpdated, nil
}
