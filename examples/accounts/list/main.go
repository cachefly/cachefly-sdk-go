package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/avvvet/cachefly-sdk-go/pkg/cachefly"
	"github.com/avvvet/cachefly-sdk-go/pkg/cachefly/api"
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

	// Initialize CacheFly client
	client := cachefly.NewClient(
		cachefly.WithToken(token),
	)

	// Prepare list options
	opts := api.ListAccountsOptions{
		Offset:       0,
		Limit:        10,
		Status:       "ACTIVE",
		IsChild:      false,
		IsParent:     false,
		ResponseType: "shallow",
	}

	// Call List endpoint
	resp, err := client.Accounts.List(context.Background(), opts)
	if err != nil {
		log.Fatalf("❌ Failed to list accounts: %v", err)
	}

	// Print out the retrieved accounts
	if len(resp.Accounts) == 0 {
		fmt.Println("✗ No accounts found.")
	} else {
		fmt.Println("✅ Accounts retrieved successfully:")
		for _, acct := range resp.Accounts {
			fmt.Printf("- ID: %s, Name: %s\n", acct.ID, acct.CompanyName)
		}
	}
}
