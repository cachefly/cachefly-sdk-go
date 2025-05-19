package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/avvvet/cachefly-sdk-go/pkg/cachefly"
	"github.com/avvvet/cachefly-sdk-go/pkg/cachefly/api"
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

	// Initialize CacheFly client
	client := cachefly.NewClient(
		cachefly.WithToken(token),
	)

	payload := api.CreateChildAccountRequest{
		CompanyName: "string",
		Username:    "stringstring",
		Password:    "stringstring",
		FullName:    "string",
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

	// Call Create (POST /accounts)
	account, err := client.Accounts.CreateChildAccount(context.Background(), payload)
	if err != nil {
		log.Fatalf("❌ Failed to create child account: %v", err)
	}

	// Marshal and print the created account as indented JSON
	out, err := json.MarshalIndent(account, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting account JSON: %v", err)
	}

	fmt.Println("\n✅ Child account created successfully:")
	fmt.Println(string(out))
}
