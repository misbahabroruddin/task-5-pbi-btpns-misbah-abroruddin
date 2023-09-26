package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/misbahabroruddin/task-5-pbi-btpns-misbah-abroruddin/controllers"
	"github.com/misbahabroruddin/task-5-pbi-btpns-misbah-abroruddin/middleware"
)

func Routes() *gin.Engine {

	router := gin.Default()

	app := router.Group("/api/v1")

	app.GET("/ping", controllers.Ping)
	app.POST("/auth/login", controllers.Login)
	app.POST("/auth/register", controllers.Register)
	app.Static("/photos", "./uploads")
	app.Use(middleware.AuthMiddleware())
	app.PUT("/users/:id", controllers.UpdateUser)
	app.DELETE("/users/:id", controllers.DeleteUser)
	app.GET("/photos", controllers.GetPhoto)
	app.POST("/photos", controllers.UploadPhoto)
	app.PUT("/photos/:id", controllers.UpdatePhoto)
	app.DELETE("/photos/:id", controllers.DeletePhoto)

	return router
}