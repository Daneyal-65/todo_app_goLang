package main

import (
	"net/http"
	"todo-app/database"
	"todo-app/routes"

	"github.com/gin-gonic/gin"
)

//  main funtion and Entry point
func main() {
	// connect to db
	database.InitDB()
// load html or screens 
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")
// screen routes
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", nil)
	})
	r.GET("/singin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	r.GET("/app", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	routes.RegisterRoutes(r)

	r.Run(":8080")
}
