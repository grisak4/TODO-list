package routes

import (
	"todo-app/controllers"
	db "todo-app/database"

	"github.com/gin-gonic/gin"
)

func Initialize(router *gin.Engine) {
	db := db.GetDB()

	// Подключение к контроллерам
	controllers.Initialize(db)

	// Определение маршрутов
	router.GET("/api/v1/tasks", controllers.GetAllTasks)
	router.POST("/api/v1/createTask", controllers.CreateTask)
	router.DELETE("/api/v1/deleteTask/:id", controllers.DeleteTask)
}
