// Example demonstrates listing access control rules for a CacheFly service.
//
// This example shows:
// - Client initialization with API token
// - Listing service rules with pagination
// - Using shallow response type for efficiency
// - Error handling and response formatting
//
// Usage:
//
//	export CACHEFLY_API_TOKEN="your-token"
//	go run main.go <service_id>
//
// Example:
//
//	go run main.go srv_123456789

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
	// Load .env (optional)
	if err := godotenv.Load(); err != nil {
		log.Printf("⚠️ Warning: unable to load .env file: %v", err)
	}

	// Ensure your API token is set
	token := os.Getenv("CACHEFLY_API_TOKEN")
	if token == "" {
		log.Fatal("❌ CACHEFLY_API_TOKEN environment variable is required")
	}

	// Read service ID
	if len(os.Args) < 2 {
		log.Fatalf("⚠️ Usage: go run main.go <service_id>")
	}
	serviceID := os.Args[1]

	// Initialize client
	client := cachefly.NewClient(cachefly.WithToken(token))

	// Set listing options with offset & limit
	opts := api.ListServiceRulesOptions{
		Offset:       5,  // skip the first 5 rules
		Limit:        10, // retrieve up to 10 rules
		ResponseType: "shallow",
	}

	// Fetch the rules
	resp, err := client.ServiceRules.List(context.Background(), serviceID, opts)
	if err != nil {
		log.Fatalf("❌ Failed to list service rules: %v", err)
	}

	// Pretty-print the list
	out, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting JSON: %v", err)
	}

	fmt.Println("\n✅ Service rules retrieved successfully:")
	fmt.Println(string(out))
}
