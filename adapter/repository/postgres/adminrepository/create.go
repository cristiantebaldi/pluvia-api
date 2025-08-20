package adminrepository

import (
	"github.com/pluvia/pluvia-api/core/domain"
	"github.com/pluvia/pluvia-api/core/dto"
	"github.com/pluvia/pluvia-api/util"
)

// Create implements domain.AdminRepository.
func (repository repository) Create(admin *dto.AdminRequestBody) (*domain.Admin, error) {
	adminCreated := domain.Admin{}

	err := repository.db.QueryRowx(
		`INSERT INTO admin (
			name, 
			email, 
			password,
			phone,
			enable
		) VALUES ($1, $2, crypt($3, gen_salt('bf', 8)), $4, $5) returning *;`,
		admin.Name,
		admin.Email,
		admin.Password,
		admin.Phone,
		admin.Enable,
	).StructScan(&adminCreated)

	if err != nil {
		return nil, util.GetError(err)
	}

	return &adminCreated, nil
}