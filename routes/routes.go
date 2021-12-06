package routes

import (
	"github.com/amshashankk/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() {
	router := gin.Default()
	api := router.Group("/api")

	api.POST("/add", controllers.AddEmployee)
	api.POST("/details/delete", controllers.DeleteEmployee)
	api.POST("/details/edit", controllers.UpdateEmpDetails)
	api.GET("/details/show", controllers.GetEmployeeById)
}
