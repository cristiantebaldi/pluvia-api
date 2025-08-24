package domain

import (
	"time"

	"github.com/pluvia/pluvia-api/core/dto"
)

type Administrador struct {
	ID             int32     `json:"id" db:"id"`
	Usuario        string    `json:"usuario" db:"usuario"`
	Email          string    `json:"email"  db:"email"`
	Senha          string    `json:"-"  db:"senha"`
	NumeroTelefone string    `json:"numeroTelefone" db:"numero_telefone"`
	Acesso         int       `json:"acesso" db:"acesso"`
	CriadoEm       time.Time `json:"criadoEm" db:"criado_em"`
	AtualizadoEm   time.Time `json:"atualizadoEm" db:"atualizado_em"`
}

type AdministradorFormData struct {
	Title         string
	Administrador Administrador
	Errors        map[string]string
	Success       bool
	Message       string
}

type AdministradorUseCase interface {
	Create(*dto.AdministradorRequestBody) (*Administrador, error)
	GetByLoginPassword(*dto.AdministradorLoginRequestBody) (*JwtAuthToken, error)
	Fetch() (*[]Administrador, error)
}

type AdministradorRepository interface {
	Create(*dto.AdministradorRequestBody) (*Administrador, error)
	GetByLoginPassword(*dto.AdministradorLoginRequestBody) (*Administrador, error)
	Fetch() (*[]Administrador, error)
}
