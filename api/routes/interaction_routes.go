package routes

import (
	"crm-system/api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupInteractionRoutes(r *gin.Engine) {
	interactionRoutes := r.Group("/interactions")
	{
		interactionRoutes.GET("/", controllers.GetInteractions)
		interactionRoutes.POST("/", controllers.CreateInteraction)
		interactionRoutes.PUT("/:id", controllers.UpdateInteraction)
		interactionRoutes.DELETE("/:id", controllers.DeleteInteraction)
	}
}
