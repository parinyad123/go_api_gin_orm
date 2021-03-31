package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	controllers "createrestful/controllers"
)

func Routes(router *gin.Engine) {
	router.GET("/", welcome) 
	router.GET("/todo", controllers.GetAllTodos)
	router.POST("/todo", controllers.CreateTodo)
	router.GET("todo/:todoId", controllers.GetSingleTodo)
	router.PUT("todo/:todoId", controllers.EditTodo)
	router.DELETE("todo/:todoId", controllers.DeleteTodo)
	router.NoRoute(notFound)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"message": "Welcome To Api",
	})
	return
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status": 404,
		"message": "Route Not Found",
	})
	return
}