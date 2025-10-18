package main

import (
	"log"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Initialize database
	InitDB()
	defer DB.Close()

	// Initialize OIDC
	InitOIDC()

	// Create Gin router
	r := gin.Default()

	// Setup sessions
	store := cookie.NewStore([]byte("some-very-secret-key")) // Replace with secure key in production
	r.Use(sessions.Sessions("session", store))

	// Setup routes
	SetupRoutes(r)

	// Serve static files
	r.Static("/static", "./frontend")

	// Load HTML templates
	r.LoadHTMLGlob("templates/*")

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("🚀 Server started on http://localhost:%s", port)
	r.Run(":" + port)
}
