// Example demonstrates updating a specific CacheFly account by ID.
//
// This example shows:
// - Client initialization with API token
// - Updating account details for a specific account ID
// - Modifying company information and settings
// - Error handling and response formatting
//
// Usage:
//
//	export CACHEFLY_API_TOKEN="your-token"
//	go run main.go <account_id>
//
// Example:
//
//	go run main.go acc_123456789

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/cachefly/cachefly-sdk-go/pkg/cachefly"
	api "github.com/cachefly/cachefly-sdk-go/pkg/cachefly/api/v2_6"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("⚠️ Warning: unable to load .env file: %v", err)
	}

	// Read API token from environment
	token := os.Getenv("CACHEFLY_API_TOKEN")
	if token == "" {
		log.Fatal("❌ CACHEFLY_API_TOKEN environment variable is required")
	}

	// Read account ID argument
	if len(os.Args) < 2 {
		log.Println("⚠️ Usage: go run main.go <account_id>")
		return
	}
	accountID := os.Args[1]

	// Initialize CacheFly client
	client := cachefly.NewClient(
		cachefly.WithToken(token),
	)

	// Prepare the update payload object
	payload := api.UpdateAccountRequest{
		CompanyName:              "new-company-name",
		Website:                  "http://new-example.com",
		Address1:                 "123 New St",
		Address2:                 "Suite 100",
		City:                     "New York",
		Country:                  "US",
		State:                    "AA",
		Phone:                    "+151900000000",
		Email:                    "new-user@example.com",
		TwoFactorAuthGracePeriod: 2,
		SAMLRequired:             false,
		DefaultDeliveryRegion:    "673f01735a5ddf015fc46997",
	}

	// Call UpdateByID (PUT /accounts/{id})
	account, err := client.Accounts.UpdateAccountByID(context.Background(), accountID, payload)
	if err != nil {
		log.Fatalf("❌ Failed to update account by ID: %v", err)
	}

	// Marshal and print the updated account as indented JSON
	out, err := json.MarshalIndent(account, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting account JSON: %v", err)
	}

	fmt.Println("\n✅ Account updated by ID successfully:")
	fmt.Println(string(out))
}
