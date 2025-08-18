package adminrepository

import (
	"github.com/pluvia/pluvia-api/core/domain"
	"github.com/pluvia/pluvia-api/util"
)

func (repository repository) Delete(id int32) (*domain.Admin, error) {
	adminDeleted := domain.Admin{}

	_, err := repository.db.Exec(
		"DELETE FROM auth WHERE admin_id = $1;",
		id,
	)

	if err != nil {
		return nil, err
	}

	err = repository.db.QueryRowx(
		"DELETE FROM admin WHERE id = $1 returning *;",
		id,
	).StructScan(&adminDeleted)

	if err != nil {
		return nil, util.GetError(err)
	}

	return &adminDeleted, nil
}
