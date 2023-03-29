package main

import (
	"wsgin/controllers"
	"wsgin/middlewares"
	"wsgin/models"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDataBase()

	router := gin.Default()

	public := router.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	public.GET("/albums", controllers.GetAlbums)
	public.GET("/time", controllers.GetTime)

	protected := router.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)

	router.Run("localhost:8080")
}
