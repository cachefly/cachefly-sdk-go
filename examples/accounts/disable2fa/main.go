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

	// Enable 2FA for the current account
	updatedAccount, err := client.Accounts.Disable2FAForCurrentAccount(context.Background())
	if err != nil {
		log.Fatalf("❌ Failed to disable 2FA: %v", err)
	}

	listJSON, err := json.MarshalIndent(updatedAccount, "", "  ")
	if err != nil {
		log.Fatalf("Error formatting MarshalIndent: %v", err)
	}

	fmt.Println("\n ✅ Two-factor authentication disabled successfully!")
	fmt.Println(string(listJSON))
}
