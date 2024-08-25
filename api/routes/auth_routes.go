package routes

import (
	"crm-system/api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(r *gin.Engine) {
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/login", controllers.Login)
		authRoutes.POST("/register", controllers.Register)
	}
}
