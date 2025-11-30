package main

import (
	"backend/connection"
	"backend/models"
	"backend/routes"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	connection.ConnectDB()

	connection.DB.AutoMigrate(&models.Projects{}, &models.User{})

	fmt.Println("Server running....")

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	routes.SetupRoutes(router)


	router.Run(":8080")

}