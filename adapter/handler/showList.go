package handler

import (
	"html/template"
	"net/http"

	"github.com/pluvia/pluvia-api/core/domain"
)

func (h *AdministradorHandler) ShowList(w http.ResponseWriter, r *http.Request) {
	administrators, err := h.usecase.Fetch(domain.Administrador{})
	if err != nil {
		http.Error(w, "Erro ao buscar administradores", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles(
		"templates/layouts/base.html",
		"templates/admin/list.html",
	))

	data := struct {
		Administrators []domain.Administrador
		Message        string
	}{
		Administrators: *administrators,
		Message:        r.URL.Query().Get("message"),
	}

	tmpl.ExecuteTemplate(w, "base", data)
}