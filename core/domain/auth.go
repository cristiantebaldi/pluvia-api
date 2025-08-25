package domain

import (
	"time"
)

type Auth struct {
	ID              int32     `json:"id"             swaggerignore:"true" db:"id"`
	Type            string    `json:"login"                               db:"type"`
	Hash            string    `json:"senha"                            db:"hash"`
	Token           string    `json:"email"          swaggerignore:"true" db:"token"`
	AdministradorID int32     `json:"accountID"      swaggerignore:"true" db:"admin_id"`
	Revoked         bool      `json:"revoked" swaggerignore:"true" db:"revoked"`
	CreatedDate     time.Time `json:"createdDate"                         db:"created_date"`
}

type AuthRepository interface {
	GetByRefreshTypeToken(string, string) (*Auth, error)
	GetByAdministradorID(int32) (*Auth, error)
	Create(Auth) error
	Delete(int32) error
	Update(int32, Auth) error
}

type AuthUseCase interface {
	RefreshToken(string) (*JwtAuthToken, error)
}
