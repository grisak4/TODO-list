package db

import (
	"database/sql"
	"fmt"
	"todo-app/config"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() {
	var err error

	config.Load() // Загрузка конфигурации из config.json
	dsn := config.AppConfig.DatabaseURL

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Error with database: ", err)
		return
	}

	// Проверка соединения с базой данных
	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging database: ", err)
		return
	}

	fmt.Println("Database has successfully connected!")
}

func GetDB() *sql.DB {
	return db
}

func CloseDB() {
	db.Close()
}
