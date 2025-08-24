// Create a new CacheFly child account using the SDK.
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

	"github.com/cachefly/cachefly-sdk-go/pkg/cachefly"
	api "github.com/cachefly/cachefly-sdk-go/pkg/cachefly/api/v2_6"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("⚠️  Warning: unable to load .env file: %v", err)
	}

	// Read API token from environment
	token := os.Getenv("CACHEFLY_API_TOKEN")
	if token == "" {
		log.Fatal("❌ CACHEFLY_API_TOKEN environment variable is required")
	}

	// Initialize CacheFly client
	client := cachefly.NewClient(
		cachefly.WithToken(token),
	)

	// Define account details
	payload := api.CreateChildAccountRequest{
		CompanyName: "child-01-account-sdk",
		Username:    "childchildsdk001",
		Password:    "stringstring",
		FullName:    "SDK GO",
		Email:       "user@example.com",
		Website:     "http://example.com",
		Address1:    "string",
		Address2:    "string",
		City:        "string",
		Country:     "string",
		State:       "string",
		Phone:       "string",
		Zip:         "string",
	}

	// Create child account
	account, err := client.Accounts.CreateChildAccount(context.Background(), payload)
	if err != nil {
		log.Fatalf("❌ Failed to create child account: %v", err)
	}

	// Display result
	out, err := json.MarshalIndent(account, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting account JSON: %v", err)
	}

	fmt.Println("✅ Child account created successfully:")
	fmt.Println(string(out))
}
