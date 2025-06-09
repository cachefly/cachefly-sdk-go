// Example demonstrates disabling two-factor authentication for the current user in CacheFly.
//
// This example shows:
// - Client initialization with API token
// - Disabling 2FA for authenticated user
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

	user, err := client.Users.DisableTwoFactorAuth(context.Background())
	if err != nil {
		log.Fatalf("❌ Failed to disable 2FA %v", err)
	}

	// Pretty-print
	out, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting JSON: %v", err)
	}

	fmt.Printf("✅ Two-factor authentication disabled for user %s\n", user.ID)
	fmt.Println(string(out))
}
