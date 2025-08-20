package authrepository

import (
	"database/sql"
	"fmt"

	"github.com/pluvia/pluvia-api/core/domain"
)

func (repository repository) GetByAdminID(id int32) (*domain.Auth, error) {
	auth := domain.Auth{}

	err := repository.db.QueryRow(
		`SELECT * FROM auth where account_id = $1;`,
		id,
	).Scan(
		&auth.ID,
		&auth.Type,
		&auth.Hash,
		&auth.Token,
		&auth.AdminID,
		&auth.Revoked,
		&auth.CreatedDate,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("token not created yet")
	}

	return &auth, nil
}
