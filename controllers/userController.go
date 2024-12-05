package controllers

import (
	"net/http"

	"authorization/initializers"
	"authorization/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}
	if c.Bind(&body) != nil {
		c.HTML(http.StatusBadRequest, "signup", gin.H{
			"error": "Failed signup",
		})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.HTML(http.StatusBadRequest, "signup", gin.H{
			"error": "Wrong password",
		})
		return

	}
	user := models.User{Email: body.Email, Password: string(hash)}

	result := initializers.DB.Create(&user) // pass pointer of data to Create
	if result.Error != nil {
		c.HTML(http.StatusBadRequest, "signup", gin.H{
			"error": "Failed to create user",
		})
		return
	}
	c.HTML(http.StatusOK, "signup", gin.H{
		"message": "Success",
	})
}

func SignUpPage(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", gin.H{}) // Render the signup.html template
}
