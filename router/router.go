package router

import (
	"net/http"

	"github.com/pluvia/pluvia-api/adapter/handler"
)

func SetupRoutes(adminHandler *handler.AdministradorHandler) *http.ServeMux {
    mux := http.NewServeMux()
    
    // Páginas HTML (Server-side Rendered)
    mux.HandleFunc("/admin/create", adminHandler.CreateAdministrador)
    mux.HandleFunc("/admin/list", adminHandler.ShowList)
    
    // Arquivos estáticos
    mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    
    return mux
}
