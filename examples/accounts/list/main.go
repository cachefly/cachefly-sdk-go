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
		Offset: 0,
		Limit:  10,
		Status: "ACTIVE",
		//IsChild:      false,
		IsParent:     true,
		ResponseType: "shallow",
	}

	// Call List endpoint
	resp, err := client.Accounts.List(context.Background(), opts)
	if err != nil {
		log.Fatalf("❌ Failed to list accounts: %v", err)
	}

	listJSON, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatalf("Error formatting MarshalIndent [account]: %v", err)
	}

	fmt.Println("\n ✅ Accounts retrieved successfully:")
	fmt.Println(string(listJSON))
}
