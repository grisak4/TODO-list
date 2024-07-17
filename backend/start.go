package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

const (
	serverAddr = "127.0.0.1:8080"
)

type Task struct {
	TaskID    int    `json:"id"`
	TaskTitle string `json:"title"`
}

func dbConnect() {
	var err error

	dsn := "root:12345@tcp(127.0.0.1:3306)/todoapp"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Error with database: ", err)
		return
	}
	fmt.Println("Database has successfully connected!")
}

func main() {
	go dbConnect()
	defer db.Close()

	router := gin.Default()

	router.GET("/api/v1/tasks", getAllTasks)
	router.POST("/api/v1/createTask", postCreateTask)
	router.DELETE("/api/v1/deleteTask/:id", deleteTask)

	router.Run(serverAddr)
}

func getAllTasks(c *gin.Context) {
	rows, err := db.Query("SELECT * FROM Tasks")
	if err != nil {
		fmt.Println("Error with getTasks: ", err)
		c.IndentedJSON(http.StatusConflict, "Error with get")
		return
	}

	var tasks []Task

	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.TaskID, &task.TaskTitle); err != nil {
			fmt.Println("Error with scan")
			c.IndentedJSON(http.StatusConflict, "Error with get")
			return
		}
		tasks = append(tasks, task)
	}

	c.IndentedJSON(http.StatusOK, tasks)
}

func postCreateTask(c *gin.Context) {
	var newTask Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		fmt.Println("Error with creating: ", err)
		c.IndentedJSON(http.StatusInternalServerError, "Error with post")
		return
	}

	result, err := db.Exec("INSERT INTO tasks (TitleTask) VALUES (?)", newTask.TaskTitle)
	if err != nil {
		fmt.Println("Error with insert: ", err)
		c.IndentedJSON(http.StatusInternalServerError, "Error with post")
		return
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Error getting last insert ID:", err)
		c.IndentedJSON(http.StatusInternalServerError, "Error with post")
		return
	}

	c.IndentedJSON(http.StatusCreated, lastInsertID)
}

func deleteTask(c *gin.Context) {
	taskID := c.Param("id")

	id, err := strconv.Atoi(taskID)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	result, err := db.Exec("DELETE FROM tasks WHERE IDTask = ?", id)
	if err != nil {
		fmt.Println("Error with deleting task: ", err)
		c.IndentedJSON(http.StatusConflict, gin.H{"error": "Error with database"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error getting rows affected:", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving affected rows"})
		return
	}

	if rowsAffected == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}

	fmt.Printf("Deleted row with ID: %d\n", id)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Deleted task", "task_id": id})
}
