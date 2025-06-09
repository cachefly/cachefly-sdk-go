// Example demonstrates listing custom domains for a CacheFly service.
//
// This example shows:
// - Client initialization with API token
// - Listing service domains with pagination
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

	if err := godotenv.Load(); err != nil {
		log.Printf("⚠️ Warning: unable to load .env file: %v", err)
	}

	token := os.Getenv("CACHEFLY_API_TOKEN")
	if token == "" {
		log.Fatal("❌ CACHEFLY_API_TOKEN environment variable is required")
	}

	if len(os.Args) < 2 {
		log.Println("⚠️ Usage: go run main.go <service_id>")
		return
	}
	serviceID := os.Args[1]

	client := cachefly.NewClient(
		cachefly.WithToken(token),
	)

	opts := api.ListServiceDomainsOptions{
		Search:       "",
		Offset:       0,
		Limit:        10,
		ResponseType: "",
	}

	resp, err := client.ServiceDomains.List(context.Background(), serviceID, opts)
	if err != nil {
		log.Fatalf("❌ Failed to list domains for service %s: %v", serviceID, err)
	}

	out, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting domains list JSON: %v", err)
	}

	fmt.Println("\n✅ Service domains retrieved successfully:")
	fmt.Println(string(out))
}
