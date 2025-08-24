// Example demonstrates listing script configurations in CacheFly.
//
// This example shows:
// - Client initialization with API token
// - Listing script configs with pagination
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

	"github.com/cachefly/cachefly-sdk-go/pkg/cachefly"
	api "github.com/cachefly/cachefly-sdk-go/pkg/cachefly/api/v2_6"
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

	// Prepare list options for script configs
	opts := api.ListScriptConfigsOptions{
		Offset: 0,
		Limit:  10,
	}

	// Call List script configurations (GET /services/{id}/scriptConfigs)
	resp, err := client.ScriptConfigs.List(context.Background(), opts)
	if err != nil {
		log.Fatalf("❌ Failed to list script configs for service: %v", err)
	}

	// Pretty-print the script configs list
	out, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting script configs list JSON: %v", err)
	}

	fmt.Println("\n✅ Script configurations retrieved successfully:")
	fmt.Println(string(out))
}
