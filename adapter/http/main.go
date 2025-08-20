package main

import (
	"fmt"
	"net/http"

	"github.com/pluvia/pluvia-api/router"
)

func main() {
	r := router.SetupRoutes()

	fmt.Println("Servidor rodando em http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
