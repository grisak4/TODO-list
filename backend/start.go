package main

import (
	db "todo-app/database"
	"todo-app/routes"

	"github.com/gin-gonic/gin"
)

const (
	serverAddr = "127.0.0.1:8080"
)

func main() {
	// Инициализация базы данных
	db.InitDB()
	defer db.CloseDB()

	router := gin.Default()
	routes.Initialize(router)
	router.Run(serverAddr)
}
