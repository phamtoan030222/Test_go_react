package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phamtoan030222/test/backend/storage"
)

func CreateTaskHandler(c *gin.Context) {
	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&req); err != nil || req.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error" : "Invalid input"})
		return 
	}

	task := storage.CreateTask(req.Title, req.Description)
	c.JSON(http.StatusCreated, gin.H{"id":task.ID})
}

func GetTaskHandler(c *gin.Context) {
	tasks := storage.GetAllTasks()
	c.JSON(http.StatusOK, tasks)
}

func UpdateTaskHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Invalid task ID"})
		return 
	}

	var req struct {
		Completed bool `json:"completed"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Invalid input"})
		return 
	}

	task, err := storage.UpdateTaskStatus(id, req.Completed)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error":"Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)

}

func DeleteTaskHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Invalid task ID"})
		return 
	}

	if err := storage.DeleteTask(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error":"Task not found"})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}