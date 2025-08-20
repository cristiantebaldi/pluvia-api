package main

import (
// 	"context"

// 	"github.com/gin-gonic/gin"
// 	"github.com/pluvia/pluvia-api/adapter/http/docs"
// 	"github.com/pluvia/pluvia-api/adapter/http/middleware"
// 	"github.com/pluvia/pluvia-api/adapter/repository/postgres"
// 	"github.com/pluvia/pluvia-api/di"
// 	"github.com/pluvia/pluvia-api/util"
// 	"github.com/spf13/viper"

// 	swaggerfiles "github.com/swaggo/files"
// 	ginSwagger "github.com/swaggo/gin-swagger"
// )

// func init() {
// 	viper.SetConfigFile(`config.json`)
// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		panic(err)
// 	}
// }

	"fmt"
	"net/http"

	"github.com/pluvia/pluvia-api/router"
)

// @title Pluvia API Docs
// @version 2025.8.4.0
// @host pluvia.api.com.br
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /
func main() {
	// validate := util.NewValidator()
	// ctx := context.Background()
	// conn := postgres.Initialize(ctx)
	// defer conn.Close()

	// adminService := di.ConfigAdminDI(conn, validate)
	
	// router := gin.Default()

	// docs.SwaggerInfo.BasePath = "/"
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// router.Use(middleware.Cors())
	// router.POST("/session", adminService.GetByLoginPassword)
	// router.GET("/session", adminService.RefreshToken)

	// jsonApi := router.Group("/")
	// {
	// 	jsonApi.POST("/admin", adminService.Create)
	// 	jsonApi.PUT("/admin/:id", adminService.Update)
	// 	jsonApi.GET("/admin", adminService.Fetch)
	// 	jsonApi.GET("/admin/:id", adminService.GetByID)
	// 	jsonApi.DELETE("/admin/:id", adminService.Delete)
	// 	}

	// port := viper.GetString("server.http.port")

	// router.Run(":" + port)

	// r := router.SetupRoutes()

	fmt.Println("Servidor rodando em http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
