package routes

import (
	"crm-system/api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupCustomerRoutes(r *gin.Engine) {
	customerRoutes := r.Group("/customers")
	{
		customerRoutes.GET("/", controllers.GetCustomers)
		customerRoutes.POST("/", controllers.CreateCustomer)
		customerRoutes.PUT("/:id", controllers.UpdateCustomer)
		customerRoutes.DELETE("/:id", controllers.DeleteCustomer)
	}
}
