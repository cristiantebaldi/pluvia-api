package controllers

import (
	"html/template"
	"net/http"
)

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/cadastro_usuario/verificacao_numero.html"))
	tmpl.ExecuteTemplate(w, "base", nil)
}
