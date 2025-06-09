// Example demonstrates listing TLS profiles in CacheFly.
//
// This example shows:
// - Client initialization with API token
// - Listing TLS profiles with pagination
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
	// Load environment variables (optional)
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

	// Prepare list options for TLS profiles
	opts := api.ListTLSProfilesOptions{
		Offset: 0,
		Limit:  10,
	}

	// Call List (GET /tlsProfiles)
	resp, err := client.TLSProfiles.List(context.Background(), opts)
	if err != nil {
		log.Fatalf("❌ Failed to list TLS profiles: %v", err)
	}

	// Pretty-print the TLS profiles list
	out, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting TLS profiles list JSON: %v", err)
	}

	fmt.Println("\n✅ TLS profiles retrieved successfully:")
	fmt.Println(string(out))
}
