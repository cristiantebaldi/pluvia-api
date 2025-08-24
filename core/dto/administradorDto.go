package dto

import (
	"html/template"
)

type AdministradorRequestBody struct {
	Usuario        string `json:"usuario" form:"usuario"`
	Email          string `json:"email" form:"email"`
	Senha          string `json:"senha" form:"senha"`
	NumeroTelefone string `json:"numeroTelefone" form:"numeroTelefone"`
	Acesso         int    `json:"acesso" form:"acesso"`
}

type AdministradorLoginRequestBody struct {
	Usuario string `json:"usuario"`
	Senha   string `json:"senha"`
}

type AdministradorTemplateData struct {
	Form    template.HTML
	Success bool
	Message string
}
