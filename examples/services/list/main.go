// Example demonstrates listing CacheFly services with filtering options.
//
// This example shows:
// - Client initialization with API token
// - Listing services with pagination and status filter
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
	if err := godotenv.Load(); err != nil {
		log.Printf("⚠️ Warning: unable to load .env file: %v", err)
	}

	token := os.Getenv("CACHEFLY_API_TOKEN")
	if token == "" {
		log.Fatal("❌ CACHEFLY_API_TOKEN environment variable is required")
	}

	client := cachefly.NewClient(
		cachefly.WithToken(token),
	)

	opts := api.ListOptions{
		Offset:          0,
		Limit:           10,
		Status:          "ACTIVE",
		IncludeFeatures: false,
		ResponseType:    "",
	}

	resp, err := client.Services.List(context.Background(), opts)
	if err != nil {
		log.Fatalf("❌ Failed to list services: %v", err)
	}

	listJSON, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatalf("Error formatting service list: %v", err)
	}

	fmt.Println("\n ✅ Current Account:")
	fmt.Println(string(listJSON))

}
