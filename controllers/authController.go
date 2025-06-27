package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/moeefa/serenity/db"
	"github.com/moeefa/serenity/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var signUpInput struct {
		Email    string `json:"email" binding:"required"`
		Name     string `json:"name" binding:"required"`
		Phone 	 string `json:"phone"`
		Password string `json:"password" binding:"required"`
	}

	db := db.GetDB()

	if err := c.ShouldBindJSON(&signUpInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound models.User
	db.Where("email=?", signUpInput.Email).Find(&userFound)

	if userFound.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email already exists"})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(signUpInput.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Email: signUpInput.Email,
		Name:  signUpInput.Name,
		Phone: signUpInput.Phone,
		Password: string(passwordHash),
	}

	db.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// curl -X POST http://localhost:8080/v1/auth/signup -H "Content-Type: application/json" -d '{"email":"testuser","name":"Test User","phone":"1234567890","password":"testpassword"}'

func LoginUser(c *gin.Context) {
	var loginInput struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound models.User
	db.GetDB().Where("email=?", loginInput.Email).Find(&userFound)

	if userFound.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(loginInput.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
		return
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userFound.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}
// curl -X POST http://localhost:8080/v1/auth/login -H "Content-Type: application/json" -d '{"email":"testuser","password":"testpassword"}'

func VerifyUser(c *gin.Context) {
	user, _ := c.Get("user")
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	currentUser, ok := user.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id":    currentUser.ID,
			"name":  currentUser.Name,
			"email": currentUser.Email,
			"phone": currentUser.Phone,
		},
	})
}