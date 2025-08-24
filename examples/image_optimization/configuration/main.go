package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/cachefly/cachefly-sdk-go/pkg/cachefly"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables (optional)
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

	// Fetch image optimization validation schema
	schema, err := client.ServiceImageOptimization.GetDetail(context.Background(), serviceID)
	if err != nil {
		log.Fatalf("❌ Failed to fetch validation schema for service %s: %v", serviceID, err)
	}

	// Pretty-print the schema JSON
	out, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting schema JSON: %v", err)
	}

	fmt.Println("\n✅ Image optimization validation schema retrieved successfully:")
	fmt.Println(string(out))
}
