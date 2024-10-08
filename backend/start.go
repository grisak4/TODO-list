package main

import (
	db "todo-app/database"
	"todo-app/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	serverAddr = "127.0.0.1:8081"
)

func main() {
	// Инициализация базы данных
	db.InitDB()
	defer db.CloseDB()

	// Конфигурация CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "ngrok-skip-browser-warning"}

	// Запуск
	r := gin.Default()
	r.Use(cors.New(config))
	routes.Initialize(r)
	r.Run(serverAddr)
}
