package main

import (
	"context"
	"log"
	"net/http"

	"github.com/pluvia/pluvia-api/adapter/handler"
	"github.com/pluvia/pluvia-api/adapter/http/middleware"
	"github.com/pluvia/pluvia-api/adapter/repository/postgres"
	"github.com/pluvia/pluvia-api/core/domain"
	"github.com/pluvia/pluvia-api/di"
	"github.com/pluvia/pluvia-api/util"
	"github.com/spf13/viper"

	// Swagger
	"github.com/gin-gonic/gin"
	"github.com/pluvia/pluvia-api/adapter/http/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

// @title Pluvia API Docs - Híbrida
// @version 2025.8.4.0
// @description API híbrida do Pluvia - Server-side Rendered + REST API
// @host pluvia.api.com.br
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /
func main() {
	// Configurações básicas
	validate := util.NewValidator()
	ctx := context.Background()
	conn := postgres.Initialize(ctx)
	defer conn.Close()

	// Dependency Injection - renomeando para administrator
	administradorUseCase := di.ConfigAdministradorDIUsecase(conn)

	// Criando handlers
	administradorHandler := handler.NewAdministradorHandler(administradorUseCase, validate)

	// Setup do router principal (net/http)
	mux := setupMainRoutes(administradorHandler)

	// Setup do Gin para Swagger (apenas para documentação)
	ginRouter := setupSwaggerRoutes(administradorUseCase, validate)

	// Servidor híbrido
	port := viper.GetString("server.http.port")
	
	log.Printf("🚀 Servidor Pluvia iniciando...")
	log.Printf("📱 Interface Web: http://localhost:%s", port)
	log.Printf("📚 Documentação Swagger: http://localhost:%s/swagger/index.html", port)
	log.Printf("🔧 API REST: http://localhost:%s/api/", port)

	// Servidor principal com roteamento híbrido
	server := &http.Server{
		Addr:    ":" + port,
		Handler: createHybridHandler(mux, ginRouter),
	}

	log.Fatal(server.ListenAndServe())
}

// setupMainRoutes configura as rotas principais (Server-side Rendered)
func setupMainRoutes(
	adminHandler handler.AdministradorHandler,
) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/admin/administrators", withAuth(handleMethodRouter(
		adminHandler.ShowList,           // GET
		nil,                            // POST não usado nesta rota
	)))
	mux.HandleFunc("/admin/administrators/create", withAuth(handleMethodRouter(
		adminHandler.ShowCreateForm,     // GET
		adminHandler.ProcessCreate,      // POST
	)))

	mux.Handle("/static/", http.StripPrefix("/static/", 
		middleware.CorsSwagger(http.FileServer(http.Dir("static")))))

	return mux
}

// setupSwaggerRoutes configura as rotas do Swagger (usando Gin apenas para isso)
func setupSwaggerRoutes(administratorUseCase interface{}, validate interface{}) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())

	// Configurar Swagger
	docs.SwaggerInfo.BasePath = "/api"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// APIs REST (para integrações e AJAX)
	api := router.Group("/api")
	{
		// Autenticação API
		api.POST("/auth/login", func(c *gin.Context) {
			// Implementação da API de login
		})
		api.GET("/auth/refresh", func(c *gin.Context) {
			// Implementação do refresh token
		})

		// Administradores API
		adminAPI := api.Group("/administrators")
		{
			adminAPI.POST("/", func(c *gin.Context) {
				// API para criar administrador
			})
			adminAPI.GET("/", func(c *gin.Context) {
				// API para listar administradores
			})
			adminAPI.GET("/:id", func(c *gin.Context) {
				// API para buscar administrador por ID
			})
			adminAPI.PUT("/:id", func(c *gin.Context) {
				// API para atualizar administrador
			})
			adminAPI.DELETE("/:id", func(c *gin.Context) {
				// API para deletar administrador
			})
		}

		// Usuários API
		userAPI := api.Group("/users")
		{
			userAPI.GET("/", func(c *gin.Context) {
				// API para listar usuários
			})
			userAPI.GET("/:id", func(c *gin.Context) {
				// API para buscar usuário por ID
			})
		}
	}

	return router
}

// createHybridHandler combina o handler principal (net/http) com Gin (para swagger/api)
func createHybridHandler(mainHandler http.Handler, ginHandler *gin.Engine) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Se a rota começar com /swagger ou /api, usar Gin
		if isAPIRoute(r.URL.Path) {
			ginHandler.ServeHTTP(w, r)
			return
		}
		
		// Caso contrário, usar o handler principal (Server-side Rendered)
		mainHandler.ServeHTTP(w, r)
	})
}

// isAPIRoute verifica se a rota é para API/Swagger
func isAPIRoute(path string) bool {
	return len(path) >= 4 && (path[:4] == "/api" || 
		   len(path) >= 8 && path[:8] == "/swagger")
}

// handleMethodRouter facilita o roteamento baseado no método HTTP
func handleMethodRouter(getHandler, postHandler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			if getHandler != nil {
				getHandler(w, r)
			} else {
				http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			}
		case http.MethodPost:
			if postHandler != nil {
				postHandler(w, r)
			} else {
				http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			}
		default:
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	}
}

// withAuth é um middleware simples para proteger rotas administrativas
func withAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Verificar se há session/cookie de autenticação
		cookie, err := r.Cookie("admin_session")
		if err != nil || cookie.Value == "" {
			// Redirecionar para login se não autenticado
			http.Redirect(w, r, "/admin/login?redirect="+r.URL.Path, http.StatusSeeOther)
			return
		}

		// Validar session/token aqui
		// Por simplicidade, assumindo que existe
		
		handler(w, r)
	}
}

// logRequest é um middleware para logging (opcional)
func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}