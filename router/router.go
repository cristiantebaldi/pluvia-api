package router

import (
	"net/http"

	"github.com/pluvia/pluvia-api/controllers"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", controllers.HomeHandler)
	mux.HandleFunc("/cadastro-usuario", controllers.ContactHandler)

	// Arquivos estáticos
	mux.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	return mux
}
