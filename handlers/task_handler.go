package handlers

import (
	"net/http"
	"strconv"
	"todo-app/database"
	"todo-app/models"

	"github.com/gin-gonic/gin"
)

// CreateTask creates a new task
func CreateTask(c *gin.Context) {
	userID, err := getUserIDFromCookie(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil || task.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
		return
	}

	task.UserID = userID
	database.DB.Create(&task)
	c.JSON(http.StatusCreated, task)
}

// GetTasks fetches all tasks for the logged-in user
func GetTasks(c *gin.Context) {
	userID, err := getUserIDFromCookie(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var tasks []models.Task
	database.DB.Where("user_id = ?", userID).Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}

// UpdateTask updates a task by ID
func UpdateTask(c *gin.Context) {
	userID, err := getUserIDFromCookie(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	var task models.Task
	if err := database.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	if task.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized"})
		return
	}

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&task)
	c.JSON(http.StatusOK, task)
}

// DeleteTask deletes a task by ID
func DeleteTask(c *gin.Context) {
	userID, err := getUserIDFromCookie(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	var task models.Task
	if err := database.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	if task.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized"})
		return
	}

	database.DB.Delete(&task)
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}

// secure routes


// Helper function to get user ID from cookie
func getUserIDFromCookie(c *gin.Context) (uint, error) {
	userToken, err := c.Cookie("user_token")
	if err != nil {
		return 0, err
	}

	userID, err := strconv.ParseUint(userToken, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(userID), nil
}
