// Example demonstrates retrieving a specific TLS profile from CacheFly.
//
// This example shows:
// - Client initialization with API token
// - Fetching TLS profile details by ID
// - Error handling and response formatting
//
// Usage:
//
//	export CACHEFLY_API_TOKEN="your-token"
//	go run main.go <profile_id>
//
// Example:
//
//	go run main.go tls_123456789

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/cachefly/cachefly-sdk-go/pkg/cachefly"
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

	// Read TLS Profile ID argument
	if len(os.Args) < 2 {
		log.Fatalf("⚠️ Usage: go run main.go <tls_profile_id>")
	}
	profileID := os.Args[1]

	// Initialize CacheFly client
	client := cachefly.NewClient(
		cachefly.WithToken(token),
	)

	// Fetch TLS profile by ID (GET /tlsProfiles/{id})
	profile, err := client.TLSProfiles.GetByID(context.Background(), profileID)
	if err != nil {
		log.Fatalf("❌ Failed to get TLS profile %s: %v", profileID, err)
	}

	// Pretty-print the TLS profile details
	out, err := json.MarshalIndent(profile, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting TLS profile JSON: %v", err)
	}

	fmt.Println("\n✅ TLS profile fetched successfully:")
	fmt.Println(string(out))
}
