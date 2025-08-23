package authrepository

import (
	"github.com/pluvia/pluvia-api/core/domain"
	"github.com/pluvia/pluvia-api/util"
)

// Create implements domain.AuthRepository.
func (repository repository) Create(auth domain.Auth) error {
	_, err := repository.db.Exec(
		`INSERT INTO auth (
			type, 
			hash, 
			token, 
			admin_id, 
			revoked
		) VALUES ($1, $2, $3, $4, $5) returning *;`,
		auth.Type,
		auth.Hash,
		auth.Token,
		auth.AdministradorID,
		auth.Revoked,
	)

	if err != nil {
		return util.GetError(err)
	}

	return nil
}
