// Example demonstrates retrieving an authentication token for a CacheFly child account.
//
// This example shows:
// - Client initialization with API token
// - Fetching auth token for a child account
// - Error handling and JSON response formatting
//
// Usage:
//
//	export CACHEFLY_API_TOKEN="your-token"
//	go run main.go <child_account_id> [responseType]
//
// Example:
//
//	go run main.go acc_child_123456789
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/cachefly/cachefly-sdk-go/pkg/cachefly"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file in project root
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: unable to load .env file: %v", err)
	}

	// Read API token from environment
	token := os.Getenv("CACHEFLY_API_TOKEN")
	if token == "" {
		log.Fatal("⚠️ CACHEFLY_API_TOKEN environment variable is required")
	}

	if len(os.Args) < 2 {
		log.Println("⚠️ Usage: go run main.go <CACHEFLY_CHILD_ID> [responseType]")
		return
	}
	childID := os.Args[1]
	if childID == "" {
		log.Fatal("⚠️ CACHEFLY_CHILD_ID environment variable is required")
	}

	// Initialize CacheFly client
	client := cachefly.NewClient(
		cachefly.WithToken(token),
	)

	// Call Get Child Account Auth Token endpoint
	authResp, err := client.Accounts.GetChildAccountAuthToken(context.Background(), childID)
	if err != nil {
		log.Fatalf("❌ Failed to get child account auth token: %v", err)
	}

	resp, err := json.MarshalIndent(authResp, "", " ")
	if err != nil {
		log.Fatalf("❌ Error formatting child auth token: %v", err)
	}

	fmt.Println("\n ✅ Child account auth token retrieved successfully:")
	fmt.Println(string(resp))

}
