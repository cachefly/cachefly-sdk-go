// Example demonstrates creating a new user in CacheFly.
//
// This example shows:
// - Client initialization with API token
// - Creating a user with permissions and service access
// - Setting admin and billing permissions
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
	// Load environment variables for token and defaults
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

	// Prepare payload for creating a new user
	opts := api.CreateUserRequest{
		Username:    "yellowgreen02",
		Password:    "yellowyellow",
		Email:       "jdoe@example.com",
		FullName:    "Yellow Green",
		Services:    []string{"681b3dc52715310035cb75d4"},
		Permissions: []string{"P_ADMIN_VIEW", "P_ADMIN_MANAGE", "P_ADMIN_BILLING", "P_ADMIN_STATS", "P_ACCOUNT_ADMIN"},
	}

	// Call Create to add the new user (POST /account/users)
	user, err := client.Users.Create(context.Background(), opts)
	if err != nil {
		log.Fatalf("❌ Failed to create user: %v", err)
	}

	// Pretty-print the created user
	out, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting user JSON: %v", err)
	}

	fmt.Println("\n✅ User created successfully:")
	fmt.Println(string(out))
}
