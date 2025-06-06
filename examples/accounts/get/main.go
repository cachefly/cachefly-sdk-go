// Example demonstrates retrieving the current CacheFly account information.
//
// This example shows:
// - Client initialization with API token
// - Fetching current account details
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

	if err := godotenv.Load(); err != nil {
		log.Printf("⚠️ Warning: unable to load .env file: %v", err)
	}

	token := os.Getenv("CACHEFLY_API_TOKEN")
	if token == "" {
		log.Fatal("❌ CACHEFLY_API_TOKEN environment variable is required")
	}

	client := cachefly.NewClient(
		cachefly.WithToken(token),
	)

	resp, err := client.Accounts.Get(context.Background(), "")
	if err != nil {
		log.Fatalf("❌ Failed to get account: %v", err)
	}

	listJSON, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatalf("Error formatting MarshalIndent [account]: %v", err)
	}

	fmt.Println("\n ✅ Current Account:")
	fmt.Println(string(listJSON))
}
