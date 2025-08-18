package authrepository

import (
	"database/sql"
	"fmt"

	"github.com/pluvia/pluvia-api/core/domain"
	"github.com/pluvia/pluvia-api/util"
)

// GetByID implements domain.AuthRepository.
func (repository repository) GetByID(id int32) (*domain.Auth, error) {
	auth := domain.Auth{}

	err := repository.db.Get(
		&auth,
		`SELECT * FROM auth where id = $1;`,
		id,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("admin not exists")
	}

	if err != nil {
		return nil, util.GetError(err)
	}

	return &auth, nil
}