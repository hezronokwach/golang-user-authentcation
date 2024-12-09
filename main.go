package main

import (
	"log"

	"authorization/controllers"
	"authorization/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	log.Println("Initializing database connection...")
	initializers.ConnectToDb()
	log.Println("Running database migrations...")
	initializers.SyncDb()
}

func main() {
    r := gin.Default()
    r.LoadHTMLGlob("templates/*")
    
    // Public routes
    r.POST("/signup", controllers.SignUp)
    r.GET("/signup", controllers.SignUpPage)
    r.POST("/login", controllers.Login)
    r.GET("/login", controllers.LoginPage)
    
    r.Run()
}
