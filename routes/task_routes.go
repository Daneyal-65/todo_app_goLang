package routes

import (
	"todo-app/handlers"

	"github.com/gin-gonic/gin"
)

// routes ragistrations
func RegisterRoutes(r *gin.Engine) {
	 r.POST("/login", handlers.Login)
	 r.POST("/register", handlers.Register)
	r.GET("/todos", handlers.GetTasks)
    r.POST("/logout", handlers.Logout)
    r.POST("/tasks", handlers.CreateTask)
    r.PUT("/tasks/:id", handlers.UpdateTask)
    r.DELETE("/tasks/:id", handlers.DeleteTask)
}
