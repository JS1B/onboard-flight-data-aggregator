package controllers

import (
	"log"
	"main/db"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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

func Login(c *gin.Context) {
	// Get the email pass off req body
	var body struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// Bind the JSON
	err := c.BindJSON(&body)
	if err != nil {
		log.Println("Error binding JSON:", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	log.Printf("Received login request: email=%s, password=***\n", body.Email) // body.Password

	// Get the user from the database
	user, err := db.GetUser(body.Email)
	if err != nil {
		log.Println("Error getting user:", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Failed to get user"})
		return
	}

	// Compare the password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(body.Password))
	if err != nil {
		log.Println("Error comparing password:", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid password"})
		return
	}

	// Generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_SALT")))
	if err != nil {
		log.Println("Error signing token:", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to sign token"})
		return
	}

	// Set the token in a cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "", "", true, true)

	// Respond
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func ValidateToken(c *gin.Context) {
	// Get the user from the context
	user, exists := c.Get("user")
	if !exists {
		log.Println("Error getting user from context")
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to get user from context"})
		return
	}

	// Respond
	c.IndentedJSON(http.StatusOK, gin.H{"message": user})
}
