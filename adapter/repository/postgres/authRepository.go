package postgres

import (
	"database/sql"
	"fmt"

	"github.com/pluvia/pluvia-api/core/domain"
	"github.com/pluvia/pluvia-api/util"
)

type authRepository struct {
	db DatabaseConfig
}

// Delete implements domain.AuthRepository.
func (r *authRepository) Delete(int32) error {
	panic("unimplemented")
}

// GetByRefreshTypeToken implements domain.AuthRepository.
func (r *authRepository) GetByRefreshTypeToken(string, string) (*domain.Auth, error) {
	panic("unimplemented")
}

// Update implements domain.AuthRepository.
func (r *authRepository) Update(int32, domain.Auth) error {
	panic("unimplemented")
}

func NewAuthRepository(db DatabaseConfig) domain.AuthRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) Create(auth domain.Auth) error {
	_, err := r.db.Exec(
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

func (r *authRepository) GetByAdministradorID(id int32) (*domain.Auth, error) {
	auth := domain.Auth{}

	err := r.db.QueryRow(
		`SELECT * FROM auth where administradores_id = $1;`,
		id,
	).Scan(
		&auth.ID,
		&auth.Type,
		&auth.Hash,
		&auth.Token,
		&auth.AdministradorID,
		&auth.Revoked,
		&auth.CreatedDate,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("token not created yet")
	}

	return &auth, nil
}
