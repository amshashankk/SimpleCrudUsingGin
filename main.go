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

	//Loading .env file
	godotenv.Load()

	//Starting database connection
	database.ConnectDB()

	fmt.Println("Server started at port 4000")

	//Setting up routes for the api
	routes.SetupRoutes(app)

	// Start and run the server
	app.Run(":5000")

}
