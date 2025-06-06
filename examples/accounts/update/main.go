// Example demonstrates updating the current CacheFly account information.
//
// This example shows:
// - Client initialization with API token
// - Updating account details including company info and settings
// - Configuring SAML and 2FA grace period
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

	// Prepare update payload object
	updatePayload := api.UpdateAccountRequest{
		CompanyName:              "parent-company-updated-sdk",
		Website:                  "http://exammple.com",
		Address1:                 "string",
		Address2:                 "string",
		City:                     "string",
		Country:                  "string",
		State:                    "string",
		Phone:                    "string",
		Email:                    "user@example.com",
		TwoFactorAuthGracePeriod: 1,
		SAMLRequired:             true,
		DefaultDeliveryRegion:    "673f01735a5ddf015fc46997",
	}

	// Call UpdateCurrent (PUT /accounts/me)
	account, err := client.Accounts.UpdateCurrentAccount(context.Background(), updatePayload)
	if err != nil {
		log.Fatalf("❌ Failed to update account: %v", err)
	}

	// Marshal and print the updated account as indented JSON
	out, err := json.MarshalIndent(account, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting account JSON: %v", err)
	}

	fmt.Println("\n✅ Account updated successfully:")
	fmt.Println(string(out))
}
