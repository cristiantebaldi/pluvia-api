package dto

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/pluvia/pluvia-api/util"
)

type AdministradorLoginRequestBody struct {
	Email string `json:"email" example:"email" validate:"required"`
	Senha string `json:"senha" example:"the senha" validate:"required,gte=5,lte=20"`
}

type AdministradorRequestBody struct {
	Usuario        string    `json:"usuario" validate:"required,gte=5,lte=150"`
	Email          string    `json:"email" validate:"required,email"`
	Senha          string    `json:"senha"`
	NumeroTelefone string    `json:"numeroTelefone" validate:"required,e164"`
	Acesso         int       `json:"acesso" validate:"required"`
	CreatedDate    time.Time `json:"createdDate"`
	AtualizadoEm   time.Time `json:"atualizadoEm"`
}

type AdministradorFormRequest struct {
    Usuario        string `validate:"required,gte=3,lte=100"`
    Email          string `validate:"required,email"`
    Senha          string `validate:"required,gte=6,lte=50"`
    ConfirmarSenha string `validate:"required,eqfield=Senha"`
    NumeroTelefone string `validate:"required,e164"`
    Acesso         int    `validate:"required,oneof=0 1"`
}

// Converte dados do formulário HTML
func FromFormToAdministradorRequest(r *http.Request, validate *validator.Validate) (*AdministradorFormRequest, error) {
    if err := r.ParseForm(); err != nil {
        return nil, err
    }
    
    acesso, _ := strconv.Atoi(r.FormValue("acesso"))
    
    adminReq := &AdministradorFormRequest{
        Usuario:        strings.TrimSpace(r.FormValue("usuario")),
        Email:          strings.TrimSpace(r.FormValue("email")),
        Senha:          r.FormValue("senha"),
        ConfirmarSenha: r.FormValue("confirmar_senha"),
        NumeroTelefone: strings.TrimSpace(r.FormValue("numero_telefone")),
        Acesso:         acesso,
    }
    
    if err := validate.Struct(adminReq); err != nil {
        return adminReq, util.HandleValidatorFieldError(err)
    }
    
    return adminReq, nil
}

func FromJsonToAdminLoginRequestBody(
	body io.ReadCloser,
	validate *validator.Validate,
) (*AdministradorLoginRequestBody, error) {
	administrador := AdministradorLoginRequestBody{}

	if err := json.NewDecoder(body).Decode(&administrador); err != nil {
		return nil, err
	}
	if err := validate.Struct(administrador); err != nil {
		return nil, util.HandleValidatorFieldError(err)
	}

	return &administrador, nil
}

func FromJsonToAdminRequestBody(
	body io.ReadCloser,
	validate *validator.Validate,
) (*AdministradorRequestBody, error) {
	administrador := AdministradorRequestBody{}

	if err := json.NewDecoder(body).Decode(&administrador); err != nil {
		return nil, err
	}

	if err := validate.Struct(administrador); err != nil {
		return nil, util.HandleValidatorFieldError(err)
	}

	return &administrador, nil
}
