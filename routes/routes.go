package routes

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Public auth routes
	router.POST("/auth/register", controllers.Register)
	router.POST("/auth/login", controllers.Login)
	router.POST("/auth/logout", controllers.Logout)

	// Protected auth routes
	router.GET("/auth/me", middleware.AuthMiddleware(), controllers.GetMe)

	// Public project routes (read only)
	router.GET("/projects", controllers.GetProjects)
	router.GET("/projects/:id", controllers.GetProjectID)

	// Protected project routes (require authentication)
	protected := router.Group("/projects")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("", controllers.CreateProjects)
		protected.PUT("/:id", controllers.UpdateProjects)
		protected.DELETE("/:id", controllers.DeleteProjects)
	}
}