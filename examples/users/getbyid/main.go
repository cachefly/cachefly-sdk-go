// Example demonstrates retrieving a specific user from CacheFly.
//
// This example shows:
// - Client initialization with API token
// - Fetching user details by ID
// - Error handling and response formatting
//
// Usage:
//
//	export CACHEFLY_API_TOKEN="your-token"
//	go run main.go <user_id>
//
// Example:
//
//	go run main.go usr_123456789

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

	// Read User ID argument
	if len(os.Args) < 2 {
		log.Fatalf("⚠️ Usage: go run main.go <user_id>")
	}
	userID := os.Args[1]

	// Initialize CacheFly client
	client := cachefly.NewClient(
		cachefly.WithToken(token),
	)

	// Call Get User by ID (GET /account/users/{id})
	user, err := client.Users.GetByID(context.Background(), userID, "")
	if err != nil {
		log.Fatalf("❌ Failed to get user %s: %v", userID, err)
	}

	// Pretty-print the user details
	out, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting user JSON: %v", err)
	}

	fmt.Println("\n✅ User fetched successfully:")
	fmt.Println(string(out))
}
