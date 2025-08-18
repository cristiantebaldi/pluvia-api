package adminrepository

import (
	"github.com/pluvia/pluvia-api/core/domain"
	"github.com/pluvia/pluvia-api/util"
)

// Fetch implements domain.AdminRepository.
func (repository repository) Fetch() (*[]domain.Admin, error) {
	admins := []domain.Admin{}

	err := repository.db.Select(&admins, `SELECT 
		id, 
		name,
		email, 
		password, 
		phone,
		enable,
		created_date,
		updated_date
        FROM admin
        ORDER BY created_date ASC`)
	if err != nil {
		return nil, util.GetError(err)
	}

	return &admins, nil
}