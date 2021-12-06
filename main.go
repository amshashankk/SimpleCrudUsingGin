package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/amshashankk/database"
	"github.com/amshashankk/routes"
)

func main() {

	//InitPostgres()
	database.Connect()

	fmt.Println("Server started at port 4000")
	defer database.DB.Close()

	// Set the router as the default one shipped with Gin
	gin.SetMode(gin.ReleaseMode)
	routes.SetupRoutes()

	app := gin.Default()
	// Start and run the server
	app.Run(":4000")

}
