package handler

import (
	"html/template"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/pluvia/pluvia-api/core/domain"
)

type AdministradorHandler struct {
	usecase   domain.AdministradorUseCase
	validator *validator.Validate
}

// ShowCreateForm implements domain.AdministradorHTTPHandler.
func (h *AdministradorHandler) ShowCreateForm(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles(
        "templates/layouts/base.html",
        "templates/admin/create_form.html",
    ))
    
    data := domain.AdministradorFormData{
        Administrador: domain.Administrador{},
        Errors:        make(map[string]string),
        Success:       false,
        Message:       "",
    }
    
    if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
        http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
        return
    }
}

func NewAdministradorHandler(usecase domain.AdministradorUseCase, validator *validator.Validate) domain.AdministradorHTTPHandler {
	return &AdministradorHandler{
		usecase:   usecase,
		validator: validator,
	}
}
