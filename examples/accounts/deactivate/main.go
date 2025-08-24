// Example usage to demonstrates how to deactivate a CacheFly child account using the SDK.
//
// This example shows:
//   - Client initialization with environment variables
//   - Error handling best practices
//   - Using the Accounts service to deactivate an account
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

	if err := godotenv.Load(); err != nil {
		log.Printf("⚠️ Warning: unable to load .env file: %v", err)
	}

	token := os.Getenv("CACHEFLY_API_TOKEN")
	if token == "" {
		log.Fatal("❌ CACHEFLY_API_TOKEN environment variable is required")
	}

	if len(os.Args) < 2 {
		log.Println("⚠️ Usage: go run main.go <account_id>")
		return
	}
	serviceID := os.Args[1]

	client := cachefly.NewClient(
		cachefly.WithToken(token),
	)

	resp, err := client.Accounts.DeactivateAccountByID(context.Background(), serviceID)
	if err != nil {
		log.Fatalf("❌ Failed to deactivate child account: %v", err)
	}

	dataJSON, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatalf("Error formatting MarshalIndent: %v", err)
	}

	fmt.Println("\n ✅ Child account deactivated successfully.")
	fmt.Println(string(dataJSON))
}
