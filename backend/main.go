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
	// Khởi tạo router Gin
	router := gin.Default()

	// Fix cảnh báo proxy (không trust proxy nào)
	if err := router.SetTrustedProxies(nil); err != nil {
		log.Fatal(err)
	}

	// Lấy domain frontend từ biến môi trường (production)
	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "https://test-go-react.vercel.app" // fallback dev local
	}

	// CORS config
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{frontendURL}, // frontend cloud URL
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
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

	// Lấy port từ Railway env
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback local
	}

	log.Println("✅ Backend server running on port", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("❌ Failed to start server: ", err)
	}
}
