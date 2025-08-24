// Example usage to demonstrates how to activate a CacheFly child account using the SDK.
//
// This example shows:
//   - Client initialization with environment variables
//   - Error handling best practices
//   - Using the Accounts service to activate an account
//   - Pretty-printing JSON responses
//
// Usage:
//
//	export CACHEFLY_API_TOKEN="your-token"
//	go run main.go <account_id>
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
	// Load environment variables from .env file if available
	if err := godotenv.Load(); err != nil {
		log.Printf("⚠️  Warning: unable to load .env file: %v", err)
	}

	// Get API token from environment
	token := os.Getenv("CACHEFLY_API_TOKEN")
	if token == "" {
		log.Fatal("❌ CACHEFLY_API_TOKEN environment variable is required")
	}

	// Validate command line arguments
	if len(os.Args) < 2 {
		log.Println("⚠️  Usage: go run main.go <account_id>")
		log.Println("   Example: go run main.go acc_123456789")
		return
	}

	accountID := os.Args[1]

	// Initialize the CacheFly client
	client := cachefly.NewClient(
		cachefly.WithToken(token),
	)

	fmt.Printf("🔄 Activating account: %s\n", accountID)

	// Activate the child account
	resp, err := client.Accounts.ActivateAccountByID(context.Background(), accountID)
	if err != nil {
		log.Fatalf("❌ Failed to activate child account: %v", err)
	}

	// Format and display the response
	dataJSON, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting response: %v", err)
	}

	fmt.Println("✅ Child account activated successfully!")
	fmt.Println("\nResponse:")
	fmt.Println(string(dataJSON))
}
