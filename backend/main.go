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

	// Fix cảnh báo proxy (không trust proxy nào)
	if err := router.SetTrustedProxies(nil); err != nil {
		log.Fatal(err)
	}

	// Lấy domain frontend từ biến môi trường (nếu có)
	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:8081" // default cho dev local
	}

	// CORS config
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			frontendURL,
			"https://test-go-react.vercel.app", // fallback cho Vercel
		},
		AllowMethods: []string{
			"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS",
		},
		AllowHeaders: []string{
			"Origin", "Content-Type", "Accept", "Authorization",
		},
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

	// Railway sẽ truyền PORT qua biến môi trường
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("✅ Server running on port", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("❌ Failed to start server: ", err)
	}
}
