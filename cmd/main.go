package main

import (
	"log"

	"github.com/AbdRofiqMm/go-simple-rest-api/internal/handlers/memberships"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	// Define a simple GET endpoint
	membershipsHandler := memberships.NewHandler(r)
	membershipsHandler.RegisterRoute()

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
