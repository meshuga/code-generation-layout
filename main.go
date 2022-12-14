package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/meshuga/code-generation-playground/api"
	"github.com/meshuga/code-generation-playground/internal/service"
)

//go:generate swag init

// @title          Swagger Example API
// @version        1.0
// @description    This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name  API Support
// @contact.url   http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @host     localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth
func main() {
	svc := service.NewService()
	r := setupRouter(svc)
	r.Run(":8080")
}

func setupRouter(svc service.Service) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		var a = api.NewApi(svc)
		accounts := v1.Group("/users")
		{
			accounts.GET("", a.ListUsers)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
