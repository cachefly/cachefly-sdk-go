// Example demonstrates listing origin servers in CacheFly.
//
// This example shows:
// - Client initialization with API token
// - Listing origins with pagination
// - Error handling and response formatting
//
// Usage:
//
//	export CACHEFLY_API_TOKEN="your-token"
//	go run main.go

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

	// Initialize CacheFly client
	client := cachefly.NewClient(
		cachefly.WithToken(token),
	)

	// Prepare list options
	opts := api.ListOriginsOptions{
		Offset: 0,
		Limit:  10,
	}

	// Call List account origins (GET /origins)
	resp, err := client.Origins.List(context.Background(), opts)
	if err != nil {
		log.Fatalf("❌ Failed to list account origins: %v", err)
	}

	// Pretty-print the origins list
	out, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting origins list JSON: %v", err)
	}

	fmt.Println("\n✅ Account origins retrieved successfully:")
	fmt.Println(string(out))
}
