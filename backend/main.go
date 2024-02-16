package main

import (
	"log"
	"main/controllers"
	"main/initializers"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.InitializeDB()

	// Register keyboard interrupt handler
	registerKeyboardInterruptHandler()
}

func main() {
	log.Print("Starting server...")
	gin.SetMode(gin.DebugMode)

	router := gin.Default()
	router.GET("/", homePage)

	router.GET("/status", getStatus)

	router.GET("/view", grafanaRedirect)

	router.POST("/signup", controllers.Signup)

	log.Println("done.")
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

		// Close the channel
		close(c)

		// Exit the application
		os.Exit(0)
	}()
}
