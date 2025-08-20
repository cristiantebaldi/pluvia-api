package adminrepository

import (
	"database/sql"
	"fmt"

	"github.com/pluvia/pluvia-api/core/domain"
	"github.com/pluvia/pluvia-api/util"
)

func (repository repository) GetByID(id int32) (*domain.Admin, error) {
	admin := domain.Admin{}

	err := repository.db.Get(
		&admin,
		`SELECT * FROM admin where id = $1;`,
		id,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("admin not exists")
	}

	if err != nil {
		return nil, util.GetError(err)
	}

	return &admin, nil
}
