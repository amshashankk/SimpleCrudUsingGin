package routes

import (
	"github.com/amshashankk/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	//Grouping the router apis with "/api" path
	api := router.Group("/api")

	api.POST("/add", controllers.AddEmployee)
	api.DELETE("/delete/:id", controllers.DeleteEmployee)
	api.POST("/update/:id", controllers.UpdateEmpDetails)
	api.GET("/show/:id", controllers.GetEmployeeById)
	api.GET("/showall", controllers.GetEmployees)
}
