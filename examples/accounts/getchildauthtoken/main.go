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

	if len(os.Args) < 2 {
		log.Println("⚠️ Usage: go run main.go <CACHEFLY_CHILD_ID> [responseType]")
		return
	}
	childID := os.Args[1]
	if childID == "" {
		log.Fatal("⚠️ CACHEFLY_CHILD_ID environment variable is required")
	}

	// Initialize CacheFly client
	client := cachefly.NewClient(
		cachefly.WithToken(token),
	)

	// Call Get Child Account Auth Token endpoint
	authResp, err := client.Accounts.GetChildAccountAuthToken(context.Background(), childID)
	if err != nil {
		log.Fatalf("❌ Failed to get child account auth token: %v", err)
	}

	resp, err := json.MarshalIndent(authResp, "", " ")
	if err != nil {
		log.Fatalf("❌ Error formatting child auth token: %v", err)
	}

	fmt.Println("\n ✅ Child account auth token retrieved successfully:")
	fmt.Println(string(resp))

}
