package domain

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pluvia/pluvia-api/core/dto"
)

type Admin struct {
	ID          int32     `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Email       string    `json:"email"  db:"email"`
	Password    string    `json:"-"  db:"password"`
	Phone       string    `json:"phone" db:"phone"`
	Enable      int       `json:"enable" db:"enable"`
	CreatedDate time.Time `json:"createdDate" db:"created_date"`
	UpdatedDate time.Time `json:"updatedDate" db:"update_date"`
}

type AdminHTTPService interface {
	Create(*gin.Context)
	Update(*gin.Context)
	Fetch(*gin.Context)
	GetByID(*gin.Context)
	Delete(*gin.Context)
}

type AdminUseCase interface {
	Create(*dto.AdminRequestBody) (*Admin, error)
	Update(int32, *dto.AdminRequestBody) (*Admin, error)
	UpdatePassword(int32, *dto.AdminRequestBody) (*Admin, error)
	Delete(int32) (*Admin, error)
	GetByLoginPassword(*dto.AdminLoginRequestBody) (*JwtAuthToken, error)
	Fetch() (*[]Admin, error)
	RefreshToken(string, string) (*JwtAuthToken, error)
}

type AdminRepository interface {
	Create(*dto.AdminRequestBody) (*Admin, error)
	Update(int32, *dto.AdminRequestBody) (*Admin, error)
	UpdatePassword(int32, *dto.AdminRequestBody) (*Admin, error)
	Delete(int32) (*Admin, error)
	GetByID(int32) (*Admin, error)
	GetByLoginPassword(*dto.AdminLoginRequestBody) (*Admin, error)
	Fetch() (*[]Admin, error)
}