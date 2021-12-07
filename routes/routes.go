package routes

import (
	"github.com/amshashankk/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")

	api.POST("/add", controllers.AddEmployee)
	api.DELETE("/delete", controllers.DeleteEmployee)
	api.POST("/update", controllers.UpdateEmpDetails)
	api.GET("/show/:id", controllers.GetEmployeeById)
	api.GET("/showall", controllers.GetEmployees)
}
