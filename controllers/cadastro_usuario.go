package controllers

import (
	"html/template"
	"net/http"
)

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/contact.html"))
	tmpl.ExecuteTemplate(w, "base", map[string]string{
		"Title": "Contato",
		"Body":  "Envie-nos uma mensagem!",
	})
}
