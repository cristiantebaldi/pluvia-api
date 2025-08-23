package administradorrepository

import (
	"github.com/pluvia/pluvia-api/core/domain"
	"github.com/pluvia/pluvia-api/util"
)

// Fetch implements domain.AdministradorRepository.
func (repository repository) Fetch() (*[]domain.Administrador, error) {
	admins := []domain.Administrador{}

	err := repository.db.Select(&admins, `SELECT 
		id, 
		usuario,
		email, 
		senha, 
		numeroTelefone,
		acesso,
		created_date,
		updated_date
        FROM administrador
        ORDER BY created_date ASC`)
	if err != nil {
		return nil, util.GetError(err)
	}

	return &admins, nil
}
