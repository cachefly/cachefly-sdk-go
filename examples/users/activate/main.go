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

	// Read User ID argument
	if len(os.Args) < 2 {
		log.Fatalf("⚠️ Usage: go run main.go <user_id>")
	}
	userID := os.Args[1]

	// Initialize CacheFly client
	client := cachefly.NewClient(cachefly.WithToken(token))

	// Activate user by ID (PUT /account/users/{id}/activate)
	resp, err := client.Users.ActivateByID(context.Background(), userID)
	if err != nil {
		log.Fatalf("❌ Failed to activate user %s: %v", userID, err)
	}

	// Pretty-print
	out, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting  JSON: %v", err)
	}

	fmt.Printf("✅ User %s activated successfully\n", userID)
	fmt.Println(string(out))
}
