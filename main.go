package main

import (
	"authorization/controllers"
	"authorization/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDb()
	initializers.SyncDb()
}

func main() {
	r := gin.Default()
    r.LoadHTMLGlob("templates/*") // Load HTML templates
	r.POST("/signup", controllers.SignUp)
    r.GET("/signup", controllers.SignUpPage) // Add a route to serve the signup page
	r.Run() // listen and serve on 0.0.0.0:8080
}