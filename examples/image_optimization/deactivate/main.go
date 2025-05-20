package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/avvvet/cachefly-sdk-go/pkg/cachefly"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("⚠️ Warning: unable to load .env file: %v", err)
	}

	// Read API token
	token := os.Getenv("CACHEFLY_API_TOKEN")
	if token == "" {
		log.Fatal("❌ CACHEFLY_API_TOKEN environment variable is required")
	}

	// Read service ID argument
	if len(os.Args) < 2 {
		log.Fatalf("⚠️ Usage: go run main.go <service_id>")
	}
	serviceID := os.Args[1]

	// Initialize CacheFly client
	client := cachefly.NewClient(cachefly.WithToken(token))

	// Deactivate image optimization configuration for the service
	if err := client.ServiceImageOptimization.DeactivateConfiguration(context.Background(), serviceID); err != nil {
		log.Fatalf("❌ Failed to deactivate image optimization for service %s: %v", serviceID, err)
	}

	fmt.Println("✅ Image optimization deactivated successfully for service:", serviceID)
}
