package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/avvvet/cachefly-sdk-go/pkg/cachefly"
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

	// Fetch the current authenticated user (GET /account/users/me)
	user, err := client.Users.GetCurrentUser(context.Background())
	if err != nil {
		log.Fatalf("❌ Failed to fetch current user: %v", err)
	}

	// Pretty-print the user details
	out, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting user JSON: %v", err)
	}

	fmt.Println("\n✅ Current user retrieved successfully:")
	fmt.Println(string(out))
}
