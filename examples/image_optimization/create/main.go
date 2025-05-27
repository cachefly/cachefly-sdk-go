package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/cachefly/cachefly-go-sdk/pkg/cachefly"
	api "github.com/cachefly/cachefly-go-sdk/pkg/cachefly/api/v2_5"
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

	// Prepare payload for creating image optimization configuration
	opts := api.CreateImageOptimizationOptions{
		Enabled:        true,
		Formats:        []string{"webp", "avif"},
		DefaultQuality: 85,
		// Add other fields as needed
	}

	// Call CreateConfiguration (POST /services/{id}/image-optimization/configuration)
	cfg, err := client.ServiceImageOptimization.CreateConfiguration(context.Background(), serviceID, opts)
	if err != nil {
		log.Fatalf("❌ Failed to create image optimization configuration for service %s: %v", serviceID, err)
	}

	// Pretty-print the created configuration
	out, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting configuration JSON: %v", err)
	}

	fmt.Println("\n✅ Image optimization configuration created successfully:")
	fmt.Println(string(out))
}
