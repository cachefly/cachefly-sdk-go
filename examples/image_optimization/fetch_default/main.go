package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/cachefly/cachefly-go-sdk/pkg/cachefly"
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

	// Fetch default image optimization configuration
	defaultCfg, err := client.ServiceImageOptimization.GetDefaults(context.Background(), serviceID)
	if err != nil {
		log.Fatalf("❌ Failed to fetch default configuration for service %s: %v", serviceID, err)
	}

	// Pretty-print the default configuration
	out, err := json.MarshalIndent(defaultCfg, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting default configuration JSON: %v", err)
	}

	fmt.Println("\n✅ Default image optimization configuration retrieved successfully:")
	fmt.Println(string(out))
}
