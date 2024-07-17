package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"todo-app/models"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func Initialize(database *sql.DB) {
	db = database
}

func GetAllTasks(c *gin.Context) {
	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		fmt.Println("Error with getTasks: ", err)
		c.IndentedJSON(http.StatusConflict, "Error with get")
		return
	}

	var tasks []models.Task

	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.TaskID, &task.TaskTitle); err != nil {
			fmt.Println("Error with scan")
			c.IndentedJSON(http.StatusConflict, "Error with get")
			return
		}
		tasks = append(tasks, task)
	}

	c.IndentedJSON(http.StatusOK, tasks)
}

func CreateTask(c *gin.Context) {
	var newTask models.Task
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

func DeleteTask(c *gin.Context) {
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
