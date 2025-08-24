// Example demonstrates listing users in a CacheFly account.
//
// This example shows:
// - Client initialization with API token
// - Listing users with pagination
// - Using shallow response type for efficiency
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

	// Prepare list options
	opts := api.ListUsersOptions{
		Offset:       0,
		Limit:        10,
		ResponseType: "shallow",
	}

	// Call List users (GET /account/users)
	resp, err := client.Users.List(context.Background(), opts)
	if err != nil {
		log.Fatalf("❌ Failed to list users: %v", err)
	}

	// Pretty-print the users list
	out, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting users list JSON: %v", err)
	}

	fmt.Println("\n✅ Users retrieved successfully:")
	fmt.Println(string(out))
}
