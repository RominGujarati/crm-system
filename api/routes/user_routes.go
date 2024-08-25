package routes

import (
	"crm-system/api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine) {
	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", controllers.GetUsers)
		userRoutes.POST("/", controllers.CreateUser)
		userRoutes.PUT("/:id", controllers.UpdateUser)
		userRoutes.DELETE("/:id", controllers.DeleteUser)
	}
}
