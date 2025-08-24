package handler

import (
	"html/template"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/pluvia/pluvia-api/core/domain"
	"github.com/pluvia/pluvia-api/core/dto"
)

type AdministradorHandler struct {
	UseCase   domain.AdministradorUseCase
	Validator *validator.Validate
}

func NewAdministradorHandler(useCase domain.AdministradorUseCase, validator *validator.Validate) *AdministradorHandler {
	return &AdministradorHandler{
		UseCase:   useCase,
		Validator: validator,
	}
}

func (h *AdministradorHandler) CreateAdministrador(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		h.ProcessCreate(w, r)
	} else {
		h.ShowCreateForm(w, r)
	}
}

func (h *AdministradorHandler) ShowCreateForm(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"templates/base.html",
		"templates/admin/create_form.html",
	))

	data := domain.AdministradorFormData{
		Title:         "Cadastrar Novo Administrador",
		Administrador: domain.Administrador{},
		Errors:        make(map[string]string),
		Success:       false,
		Message:       "",
	}

	
	if err := tmpl.ExecuteTemplate(w, "base.html", data); err != nil {
		http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
		return
	}
}

func (h *AdministradorHandler) ProcessCreate(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Erro ao processar formulário", http.StatusBadRequest)
		return
	}

	ar := &dto.AdministradorRequestBody{
		Usuario:        r.FormValue("usuario"),
		Email:          r.FormValue("email"),
		Senha:          r.FormValue("senha"),
		NumeroTelefone: r.FormValue("numero_telefone"),
	}

	_, err := h.UseCase.Create(ar)

	if err != nil {
		http.Error(w, "Erro ao criar administrador", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/list", http.StatusSeeOther)
}

func (h *AdministradorHandler) ShowList(w http.ResponseWriter, r *http.Request) {
	administradores, err := h.UseCase.Fetch()
	if err != nil {
		http.Error(w, "Erro ao buscar administradores", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles(
		"templates/base.html",
		"templates/admin/list.html",
	))

	data := struct {
		Title           string
		Administradores []domain.Administrador
		Message         string
	}{
		Title:           "Lista de Administradores",
		Administradores: *administradores,
		Message:         r.URL.Query().Get("message"),
	}

	if err := tmpl.ExecuteTemplate(w, "base.html", data); err != nil {
		http.Error(w, "Erro ao renderizar a página", http.StatusInternalServerError)
	}
}