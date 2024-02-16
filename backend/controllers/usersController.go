package controllers

import (
	"log"
	"main/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	// Get the email pass off req body
	var body struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.BindJSON(&body); err != nil {
		log.Println("Error binding JSON:", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	log.Printf("Received signup request: email=%s, password=***\n", body.Email) // body.Password

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error generating password hash:", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to hash password"})
		return
	}

	// Create a save the user
	err = db.AddUser(body.Email, string(hash))
	if err != nil {
		log.Println("Error creating user:", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Failed to create user"})
		return
	}

	// Respond
	c.IndentedJSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
