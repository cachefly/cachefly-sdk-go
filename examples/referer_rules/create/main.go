// Example demonstrates creating a referer rule for a CacheFly service.
//
// This example shows:
// - Client initialization with API token
// - Creating referer access control for specific directory
// - Setting allowed domains and default action
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

	"github.com/cachefly/cachefly-sdk-go/pkg/cachefly"
	api "github.com/cachefly/cachefly-sdk-go/pkg/cachefly/api/v2_6"
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

	// Read Service ID argument
	if len(os.Args) < 2 {
		log.Println("⚠️ Usage: go run main.go <service_id>")
		return
	}
	serviceID := os.Args[1]

	// Initialize CacheFly client
	client := cachefly.NewClient(
		cachefly.WithToken(token),
	)

	// Prepare payload for new referer rule
	payload := api.CreateRefererRuleRequest{
		Directory:     "/images",
		Extension:     "jpg",
		Exceptions:    []string{"trusted.example.com"},
		DefaultAction: "ALLOW", // or "DENY"
	}

	// Call Create referer rule (POST /services/{id}/options/refererRules)
	rule, err := client.ServiceOptionsRefererRules.Create(context.Background(), serviceID, payload)
	if err != nil {
		log.Fatalf("❌ Failed to create referer rule for service %s: %v", serviceID, err)
	}

	// Pretty-print the created rule
	out, err := json.MarshalIndent(rule, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting referer rule JSON: %v", err)
	}

	fmt.Println("\n✅ Service referer rule created successfully:")
	fmt.Println(string(out))
}
