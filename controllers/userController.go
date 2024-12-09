package controllers

import (
	"net/http"
	"regexp"

	"authorization/initializers"
	"authorization/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type SignUpInput struct {
	Email       string `form:"email" binding:"required"`
	Password    string `form:"password" binding:"required"`
	PhoneNumber string `form:"phone_number" binding:"required"`
	FirstName   string `form:"first_name" binding:"required"`
	LastName    string `form:"last_name" binding:"required"`
}

func SignUp(c *gin.Context) {
	var input SignUpInput
	if err := c.ShouldBind(&input); err != nil {
		c.HTML(http.StatusBadRequest, "signup.html", gin.H{"message": "Invalid input"})
		return
	}

	// Validate email format
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if matched, _ := regexp.MatchString(emailRegex, input.Email); !matched {
		c.HTML(http.StatusBadRequest, "signup.html", gin.H{"message": "Invalid email format"})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "signup.html", gin.H{"message": "Failed to hash password"})
		return
	}

	// Create the user
	user := models.User{
		Email:       input.Email,
		Password:    string(hashedPassword),
		PhoneNumber: input.PhoneNumber,
		FirstName:   input.FirstName,
		LastName:    input.LastName,
	}

	if err := initializers.DB.Create(&user).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "signup.html", gin.H{"message": "Failed to create user"})
		return
	}

	c.HTML(http.StatusOK, "login.html", gin.H{"message": "User created successfully"})
}

func SignUpPage(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", gin.H{}) // Render the signup.html template
}

type LoginInput struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBind(&input); err != nil {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"message": "Invalid input"})
		return
	}

	// Find user by email
	var user models.User
	if err := initializers.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"message": "Invalid credentials"})
		return
	}

	// Compare passwords
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"message": "Invalid credentials"})
		return
	}
	c.HTML(http.StatusOK, "login.html", gin.H{"message": "Login successful"})
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}
