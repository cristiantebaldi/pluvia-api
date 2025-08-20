package controllers

import (
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/home.html"))
	tmpl.ExecuteTemplate(w, "base", map[string]string{
		"Title": "Página Inicial",
		"Body":  "Bem-vindo ao site em Go!",
	})
}
