package handler

import (
	"html/template"
	"net/http"

	"github.com/pluvia/pluvia-api/core/domain"
	"github.com/pluvia/pluvia-api/core/dto"
	"github.com/pluvia/pluvia-api/util"
)

func (h *AdministradorHandler) ProcessCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/admin/create", http.StatusSeeOther)
		return
	}

	// Converter dados do formulário
	adminReq, err := dto.FromFormToAdministradorRequest(r, h.validator)

	tmpl := template.Must(template.ParseFiles(
		"templates/layouts/base.html",
		"templates/admin/create_form.html",
	))

	data := domain.AdministradorFormData{
		Administrador: domain.Administrador{
			Usuario:        adminReq.Usuario,
			Email:          adminReq.Email,
			NumeroTelefone: adminReq.NumeroTelefone,
			Acesso:         adminReq.Acesso,
		},
		Errors:  make(map[string]string),
		Success: false,
		Message: "",
	}

	// Validação falhou
	if err != nil {
		if reqErr, ok := err.(util.RequestError); ok {
			for _, field := range reqErr.Fields {
				data.Errors[field.Field] = field.Message
			}
		} else {
			data.Errors["general"] = "Erro na validação dos dados"
		}

		tmpl.ExecuteTemplate(w, "base", data)
		return
	}

	// Criar administrador
	createdAdmin, err := h.usecase.Create(&dto.AdministradorRequestBody{
		Usuario:        adminReq.Usuario,
		Email:          adminReq.Email,
		Senha:          adminReq.Senha,
		NumeroTelefone: adminReq.NumeroTelefone,
		Acesso:         adminReq.Acesso,
	})

	if err != nil {
		data.Errors["general"] = "Erro ao criar administrador: " + err.Error()
		tmpl.ExecuteTemplate(w, "base", data)
		return
	}

	// Sucesso - redirecionar ou mostrar mensagem
	data.Success = true
	data.Message = "Administrador criado com sucesso!"
	data.Administrador = *createdAdmin

	tmpl.ExecuteTemplate(w, "base", data)
}
