package postgres

import (
	"database/sql"
	"fmt"

	"github.com/pluvia/pluvia-api/core/domain"
	"github.com/pluvia/pluvia-api/core/dto"
	"github.com/pluvia/pluvia-api/util"
)

type administradorRepository struct {
	db DatabaseConfig
}

func NewAdministradorRepository(db DatabaseConfig) domain.AdministradorRepository {
	return &administradorRepository{
		db: db,
	}
}

func (r *administradorRepository) Create(administrador *dto.AdministradorRequestBody) (*domain.Administrador, error) {
	adminCreated := domain.Administrador{}
	
	err := r.db.QueryRowx(
		`INSERT INTO administradores (
			usuario, 
			email, 
			senha,
			numero_telefone,
			acesso
		) VALUES ($1, $2, crypt($3, gen_salt('bf', 8)), $4, $5) returning *;`,
		administrador.Usuario,
		administrador.Email,
		administrador.Senha,
		administrador.NumeroTelefone,
		administrador.Acesso,
	).StructScan(&adminCreated)

	if err != nil {
		return nil, util.GetError(err)
	}

	return &adminCreated, nil
}

func (r *administradorRepository) Fetch() (*[]domain.Administrador, error) {
	admins := []domain.Administrador{}

	err := r.db.Select(&admins, `SELECT 
		id, 
		usuario,
		email, 
		senha, 
		numero_telefone,
		acesso,
		criado_em,
		atualizado_em
        FROM administradores
        ORDER BY criado_em ASC`)
	if err != nil {
		return nil, util.GetError(err)
	}

	return &admins, nil
}

func (r *administradorRepository) GetByLoginPassword(loginRequest *dto.AdministradorLoginRequestBody) (*domain.Administrador, error) {
	administrador := domain.Administrador{}

	err := r.db.Get(
		&administrador,
		`SELECT * FROM administradores where usuario = $1 and senha = crypt($2, senha);`,
		loginRequest.Usuario,
		loginRequest.Senha,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("administrador not exists")
	}

	if err != nil {
		return nil, util.GetError(err)
	}

	return &administrador, nil
}