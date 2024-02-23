package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"main/controllers"
	"main/finalizers"
	"main/initializers"
	"main/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.InitializeDB()

	// Register keyboard interrupt handler
	registerKeyboardInterruptHandler()

	// Initialize the logger
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Log the start of the application
	log.Println("Starting application...")

	// Set the Gin mode to debug or release
	if os.Getenv("DEBUG") == "false" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
}

func main() {
	router := gin.Default()

	router.GET("/", homePage)
	router.GET("/status", getStatus)
	router.GET("/view", grafanaRedirect)
	router.GET("/validate", middleware.RequireAuth, controllers.ValidateToken)

	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)

	router.Run(":8080")
}

func homePage(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Home page", "data": "Welcome to the FlightStream AI"})
}

func getStatus(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Server is running"})
}

// Redirect to grafana
func grafanaRedirect(c *gin.Context) {
	c.Redirect(http.StatusPermanentRedirect, "http://localhost:3000")
}

func registerKeyboardInterruptHandler() {
	// Create a channel to receive OS signals
	c := make(chan os.Signal, 1)

	// Notify the channel on interrupt or SIGTERM signals
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Start a goroutine that will wait for a signal
	go func() {
		// Wait for a signal
		sig := <-c

		// Log the received signal
		log.Printf("Received signal: %s. Shutting down...\n", sig)

		// Run the finalizers
		finalizers.ShutdownDB()

		// Close the channel
		close(c)

		// Exit the application
		os.Exit(0)
	}()
}
