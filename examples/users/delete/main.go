// Example demonstrates deleting a user account from CacheFly.
//
// This example shows:
// - Client initialization with API token
// - Deleting a user by ID
// - Error handling for delete operations
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
	client := cachefly.NewClient(cachefly.WithToken(token))

	// Delete user by ID (DELETE /account/users/{id})
	if err := client.Users.DeleteByID(context.Background(), userID); err != nil {
		log.Fatalf("❌ Failed to delete user %s: %v", userID, err)
	}

	fmt.Printf("✅ User %s deleted successfully\n", userID)
}
