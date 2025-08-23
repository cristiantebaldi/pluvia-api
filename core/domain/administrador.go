package domain

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pluvia/pluvia-api/core/dto"
)

type Administrador struct {
	ID             int32     `json:"id" db:"id"`
	Usuario        string    `json:"usuario" db:"usuario"`
	Email          string    `json:"email"  db:"email"`
	Senha          string    `json:"-"  db:"senha"`
	NumeroTelefone string    `json:"numeroTelefone" db:"numeroTelefone"`
	Acesso         int      `json:"acesso" db:"acesso"`
	CreatedDate    time.Time `json:"createdDate" db:"created_date"`
	AtualizadoEm   time.Time `json:"atualizadoEm" db:"updated_date"`
}

type AdministradorFormData struct {
	Administrador Administrador
	Errors        map[string]string
	Success       bool
	Message       string
}

type AdministradorHTTPHandler interface {
	ShowCreateForm(http.ResponseWriter, *http.Request)
	ProcessCreate(http.ResponseWriter, *http.Request)
	ShowList(http.ResponseWriter, *http.Request)
	//ShowEditForm(http.ResponseWriter, *http.Request)
	//ProcessUpdate(http.ResponseWriter, *http.Request)
	//ProcessDelete(http.ResponseWriter, *http.Request)
}

type AdministradorHTTPService interface {
	Create(*gin.Context)
	Update(*gin.Context)
	Fetch(*gin.Context)
	GetByID(*gin.Context)
	Delete(*gin.Context)
	UpdatePassword(*gin.Context)

	GetByLoginPassword(*gin.Context)
	RefreshToken(*gin.Context)
}

type AdministradorUseCase interface {
	Create(*dto.AdministradorRequestBody) (*Administrador, error)
	Update(int32, *dto.AdministradorRequestBody) (*Administrador, error)
	UpdatePassword(int32, *dto.AdministradorRequestBody) (*Administrador, error)
	GetByID(int32) (*Administrador, error)
	Delete(int32) (*Administrador, error)
	GetByLoginPassword(*dto.AdministradorLoginRequestBody) (*JwtAuthToken, error)
	Fetch(Administrador) (*[]Administrador, error)
	RefreshToken(string, string) (*JwtAuthToken, error)
}

type AdministradorRepository interface {
	Create(*dto.AdministradorRequestBody) (*Administrador, error)
	Update(int32, *dto.AdministradorRequestBody) (*Administrador, error)
	UpdatePassword(int32, *dto.AdministradorRequestBody) (*Administrador, error)
	Delete(int32) (*Administrador, error)
	GetByID(int32) (*Administrador, error)
	GetByLoginPassword(*dto.AdministradorLoginRequestBody) (*Administrador, error)
	Fetch() (*[]Administrador, error)
}
