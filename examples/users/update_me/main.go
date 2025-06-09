// Example demonstrates updating the current authenticated user in CacheFly.
//
// This example shows:
// - Client initialization with API token
// - Updating current user profile
// - Modifying email and full name
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

	// Prepare payload to update the current authenticated user
	opts := api.UpdateUserRequest{
		Email:    "updated@example.com",
		FullName: "Updated User",
	}

	// Call UpdateCurrent to modify the current user (PUT /account/users/me)
	updatedUser, err := client.Users.UpdateCurrentUser(context.Background(), opts)
	if err != nil {
		log.Fatalf("❌ Failed to update current user: %v", err)
	}

	// Pretty-print the updated user details
	out, err := json.MarshalIndent(updatedUser, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting user JSON: %v", err)
	}

	fmt.Println("\n✅ Current user updated successfully:")
	fmt.Println(string(out))
}
