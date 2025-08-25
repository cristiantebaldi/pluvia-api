package main

import (
	"context"
	"log"
	"net/http"

	"github.com/pluvia/pluvia-api/adapter/http/middleware"
	"github.com/pluvia/pluvia-api/core/usecase"
	"github.com/pluvia/pluvia-api/adapter/handler"
	"github.com/pluvia/pluvia-api/adapter/repository/postgres"
	"github.com/pluvia/pluvia-api/di"
	"github.com/pluvia/pluvia-api/router"
	"github.com/pluvia/pluvia-api/util"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	// Configurações básicas
	validate := util.NewValidator()
	ctx := context.Background()
	conn := postgres.Initialize(ctx)
	defer conn.Close()

	// Dependency Injection
	administradorUseCase := di.ConfigAdministradorDI(conn)
	authRepository := postgres.NewAuthRepository(conn)
	authUseCase := usecase.NewAuthUseCase(authRepository)

	// Criando handlers
	administradorHandler := handler.NewAdministradorHandler(administradorUseCase, validate)

	// Setup do router
	authMiddleware := middleware.NewAuthMiddleware(authUseCase)
	mux := router.SetupRoutes(administradorHandler, authMiddleware)

	// Servidor
	port := viper.GetString("server.http.port")

	log.Printf("🚀 Servidor Pluvia iniciando...")
	log.Printf("🌐 Interface Web: http://localhost:%s/admin/login", port)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}