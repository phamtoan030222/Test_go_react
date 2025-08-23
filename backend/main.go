package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/phamtoan030222/test/backend/handlers"
)

func main() {
	router := gin.Default()

	if err := router.SetTrustedProxies(nil); err != nil {
		log.Fatal(err)
	}

	// Cấu hình CORS chi tiết
	config := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS", "PUT"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	// Xử lý multiple origins từ environment variable
	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		log.Println("⚠️  FRONTEND_URL not set, allowing all origins for development")
		config.AllowAllOrigins = true
	} else {
		origins := strings.Split(frontendURL, ",")
		for i, origin := range origins {
			origins[i] = strings.TrimSpace(origin)
		}
		config.AllowOrigins = origins
		log.Printf("✅ CORS configured for origins: %v", origins)
	}

	router.Use(cors.New(config))

	// Middleware để log CORS headers (debug)
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		if len(config.AllowOrigins) > 0 {
			origin := c.Request.Header.Get("Origin")
			for _, allowedOrigin := range config.AllowOrigins {
				if origin == allowedOrigin {
					c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
					break
				}
			}
		}
		c.Next()
	})

	// API routes
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "API is working!"})
		})
		
		api.POST("/tasks", handlers.CreateTaskHandler)
		api.GET("/tasks", handlers.GetTaskHandler)
		api.PATCH("/tasks/:id", handlers.UpdateTaskHandler)
		api.DELETE("/tasks/:id", handlers.DeleteTaskHandler)
	}

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":    "OK",
			"timestamp": time.Now().Format(time.RFC3339),
			"service":   "Go Backend API",
		})
	})

	// Root endpoint
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":   "Welcome to Go Backend API",
			"endpoints": []string{"/api", "/api/tasks", "/health"},
			"timestamp": time.Now().Format(time.RFC3339),
		})
	})

	// Handle OPTIONS requests for CORS preflight
	router.OPTIONS("/*path", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server starting on port %s", port)
	log.Printf("🌐 Frontend URL: %s", frontendURL)
	
	if err := router.Run(":" + port); err != nil {
		log.Fatal("❌ Failed to start server: ", err)
	}
}