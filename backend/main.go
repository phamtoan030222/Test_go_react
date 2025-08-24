package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/phamtoan030222/test/backend/handlers"
)

func main() {
	router := gin.Default()

	// CORS middleware đơn giản nhưng hiệu quả
	allowedOrigins := map[string]bool{
    "https://test-go-react.vercel.app": true,
    "https://test-go-react-git-master-phamtoan-s-projects.vercel.app": true,
    "http://localhost:5173": true, // cho dev local
    }

    router.Use(func(c *gin.Context) {
        origin := c.Request.Header.Get("Origin")
        if allowedOrigins[origin] {
            c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
            c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
            c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
            c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        }

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        c.Next()
    })

	// API routes
	api := router.Group("/api")
	{
		api.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status":    "OK",
				"service":   "Backend API",
				"platform":  "Railway",
			})
		})
		
		api.POST("/tasks", handlers.CreateTaskHandler)
		api.GET("/tasks", handlers.GetTaskHandler)
		api.PATCH("/tasks/:id", handlers.UpdateTaskHandler)
		api.DELETE("/tasks/:id", handlers.DeleteTaskHandler)
	}

	// Root endpoint
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Backend Server is Running",
			"endpoints": map[string]string{
				"health": "/api/health",
				"tasks":  "/api/tasks",
			},
		})
	})

	// Lấy port từ environment (Railway tự set)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default for local
	}

	log.Printf("🚀 Server starting on port %s", port)
	
	// Khởi động server
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}