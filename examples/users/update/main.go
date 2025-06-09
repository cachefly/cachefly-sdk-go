// Example demonstrates updating a user account in CacheFly.
//
// This example shows:
// - Client initialization with API token
// - Updating user details including password and permissions
// - Modifying service access and email
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

	// Read User ID argument
	if len(os.Args) < 2 {
		log.Fatalf("⚠️ Usage: go run main.go <user_id>")
	}
	userID := os.Args[1]

	// Initialize CacheFly client
	client := cachefly.NewClient(
		cachefly.WithToken(token),
	)

	// Prepare payload for updating the user
	opts := api.UpdateUserRequest{
		Password:    "yellowyellow",
		Email:       "updated_by_sdk@example.com",
		FullName:    "Updated Yellow Green",
		Services:    []string{"681b3dc52715310035cb75d4"},
		Permissions: []string{"P_ADMIN_VIEW", "P_ADMIN_MANAGE", "P_ADMIN_BILLING", "P_ADMIN_STATS", "P_ACCOUNT_ADMIN"},
	}

	// Call Update (PUT /account/users/{id})
	updatedUser, err := client.Users.UpdateByID(context.Background(), userID, opts)
	if err != nil {
		log.Fatalf("❌ Failed to update user %s: %v", userID, err)
	}

	// Pretty-print the updated user
	out, err := json.MarshalIndent(updatedUser, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting user JSON: %v", err)
	}

	fmt.Println("\n✅ User updated successfully:")
	fmt.Println(string(out))
}
