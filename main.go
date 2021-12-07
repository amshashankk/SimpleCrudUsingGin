package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/amshashankk/database"
	"github.com/amshashankk/routes"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	app := gin.Default()

	godotenv.Load()

	//InitPostgres()
	database.ConnectDB()

	fmt.Println("Server started at port 4000")

	
	routes.SetupRoutes(app)

	// Start and run the server
	app.Run(":5000")

}
