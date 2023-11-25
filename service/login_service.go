package service

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
	"todo-app/dto"
)

var jwtKey = []byte("your_secret_key_here")

func Login(c *gin.Context) {
	var loginRes dto.LoginRes
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username != "example" || password != "password" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = 1                             // Replace with your user ID or relevant information
	expirationTime := time.Now().Add(24 * time.Hour) // Change the duration as needed
	claims["exp"] = expirationTime.Unix()
	claims["role"] = []string{"admin", "todo_user"}

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	loginRes.Token = tokenString
	res := dto.BaseRes{Data: loginRes}

	c.JSON(http.StatusOK, res)
}
