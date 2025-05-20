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
	// Load .env file (optional)
	if err := godotenv.Load(); err != nil {
		log.Printf("⚠️ Warning: unable to load .env file: %v", err)
	}

	// Ensure API token is set
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

	// Fetch image optimization configuration
	config, err := client.ServiceImageOptimization.GetConfiguration(context.Background(), serviceID)
	if err != nil {
		log.Fatalf("❌ Failed to fetch configuration for service %s: %v", serviceID, err)
	}

	// Print the raw configuration (YAML or JSON)
	fmt.Println("\n✅ Configuration fetched successfully:")
	fmt.Println(config)
}
