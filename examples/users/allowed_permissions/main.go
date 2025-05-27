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
	// Load environment variables (optional)
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

	// Fetch allowed permissions for the given user (GET /account/users/{id}/allowedPermissions)
	resp, err := client.Users.GetAllowedPermissions(context.Background(), userID)
	if err != nil {
		log.Fatalf("❌ Failed to list allowed permissions for user %s: %v", userID, err)
	}

	// Pretty-print the created user
	out, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting user permissions JSON: %v", err)
	}

	fmt.Println("\n✅ Allowed permissions for user:", userID)
	fmt.Println(string(out))

}
