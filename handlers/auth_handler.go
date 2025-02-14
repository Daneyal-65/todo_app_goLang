package handlers

import (
	"net/http"
	"strconv"
	"todo-app/database"
	"todo-app/models"

	"github.com/gin-gonic/gin"
)

// Register handles user registration
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Check if username already exists
	var existingUser models.User
	if err := database.DB.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already taken"})
		return
	}

	// Save new user to database
	database.DB.Create(&user)

	// Set user ID as a cookie
	c.SetCookie("user_token", strconv.FormatUint(uint64(user.ID), 10), 3600, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

//  for login
func Login(c *gin.Context) {
	var user models.User
	var foundUser models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Check if user exists
	if err := database.DB.Where("username = ?", user.Username).First(&foundUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Set user ID as a cookie
	c.SetCookie("user_token", strconv.FormatUint(uint64(foundUser.ID), 10), 3600, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

// Logout handles user logout
func Logout(c *gin.Context) {
	// Clear the cookie
	c.SetCookie("user_token", "", -1, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
