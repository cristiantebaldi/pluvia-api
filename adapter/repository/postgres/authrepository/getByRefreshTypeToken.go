package authrepository

import (
	"database/sql"
	"fmt"

	"github.com/pluvia/pluvia-api/core/domain"
	"github.com/pluvia/pluvia-api/util"
)

// GetByRefreshTypeToken implements domain.AuthRepository.
func (repository repository) GetByRefreshTypeToken(refreshToken string, typeToken string) (*domain.Auth, error) {
	auth := domain.Auth{}

	err := repository.db.QueryRowx(
		`SELECT * FROM auth where hash = $1 and type = $2;`,
		refreshToken, typeToken,
	).StructScan(&auth)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("this token not exists Please, verify your token or do login again")
	}

	if err != nil {
		return nil, util.GetError(err)
	}

	return &auth, nil
}