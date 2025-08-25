package router

import (
	"net/http"

	"github.com/pluvia/pluvia-api/adapter/handler"
	"github.com/pluvia/pluvia-api/adapter/http/middleware"
)

func SetupRoutes(adminHandler *handler.AdministradorHandler, authMiddleware *middleware.AuthMiddleware) *http.ServeMux {
    mux := http.NewServeMux()

    // Rota principal redireciona para o login
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
    })

    // Páginas HTML (Server-side Rendered)
    mux.HandleFunc("/admin/login", adminHandler.Login)
    mux.HandleFunc("/admin/logout", adminHandler.Logout)
    mux.Handle("/admin/create", authMiddleware.Authenticate(http.HandlerFunc(adminHandler.CreateAdministrador)))
    mux.Handle("/admin/list", authMiddleware.Authenticate(http.HandlerFunc(adminHandler.ShowList)))

    // Arquivos estáticos
    mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("templates"))))

    return mux
}
