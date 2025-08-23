package administradorrepository

import (
	"database/sql"
	"fmt"

	"github.com/pluvia/pluvia-api/core/domain"
	"github.com/pluvia/pluvia-api/util"
)

func (repository repository) GetByID(id int32) (*domain.Administrador, error) {
	administrador := domain.Administrador{}

	err := repository.db.Get(
		&administrador,
		`SELECT * FROM administrador where id = $1;`,
		id,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("administrador not exists")
	}

	if err != nil {
		return nil, util.GetError(err)
	}

	return &administrador, nil
}
