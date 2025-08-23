package administradorrepository

import (
	"github.com/pluvia/pluvia-api/core/domain"
	"github.com/pluvia/pluvia-api/core/dto"
	"github.com/pluvia/pluvia-api/util"
)

// Create implements domain.AdministradorRepository.
func (repository repository) Create(administrador *dto.AdministradorRequestBody) (*domain.Administrador, error) {
	adminCreated := domain.Administrador{}

	err := repository.db.QueryRowx(
		`INSERT INTO administrador (
			usuario, 
			email, 
			senha,
			numeroTelefone,
			acesso
		) VALUES ($1, $2, crypt($3, gen_salt('bf', 8)), $4, $5) returning *;`,
		administrador.Usuario,
		administrador.Email,
		administrador.Senha,
		administrador.NumeroTelefone,
		administrador.Acesso,
	).StructScan(&adminCreated)

	if err != nil {
		return nil, util.GetError(err)
	}

	return &adminCreated, nil
}
