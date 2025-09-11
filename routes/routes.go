package routes

import (
	"unbound/auth_data_processing/middleware"
	"unbound/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {

	publicRoutes := map[string]bool{
		"POST /api/v1/auth/register": true,
		"POST /api/v1/auth/login":    true,
	}

	api := r.Group("/api/v1")
	api.Use(middleware.JWTAuth(publicRoutes))
	{

		// grupo de autenticacao
		auth := r.Group("/auth")
		{
			auth.POST("/register", controllers.Register)
			auth.POST("/login", controllers.Login)
			// auth.POST("/logout", controllers.Logout)
		}

	}

}
