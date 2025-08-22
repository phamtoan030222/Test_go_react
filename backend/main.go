package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/phamtoan030222/test/backend/handlers"
)

func main() {
	router := gin.Default()

	// CORS config
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",            // khi dev
			"https://test-go-react.vercel.app", // frontend trên Vercel
		},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// API routes
	api := router.Group("/api")
	{
		api.POST("/tasks", handlers.CreateTaskHandler)
		api.GET("/tasks", handlers.GetTaskHandler)
		api.PATCH("/tasks/:id", handlers.UpdateTaskHandler)
		api.DELETE("/tasks/:id", handlers.DeleteTaskHandler)
	}

	// Lấy PORT từ Railway
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	log.Println("Server running on port", port)
	router.Run(":" + port)
}
