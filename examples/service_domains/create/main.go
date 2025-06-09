// Example demonstrates adding a custom domain to a CacheFly service.
//
// This example shows:
// - Client initialization with API token
// - Creating a service domain with description
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
	// Load environment variables
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

	// Prepare create payload for new domain
	payload := api.CreateServiceDomainRequest{
		Name:        "www.example.com",
		Description: "example description from SDK domain create",
	}

	// Call Create service domain
	domain, err := client.ServiceDomains.Create(context.Background(), serviceID, payload)
	if err != nil {
		log.Fatalf("❌ Failed to create domain for service %s: %v", serviceID, err)
	}

	// Pretty-print the created domain
	out, err := json.MarshalIndent(domain, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting domain JSON: %v", err)
	}

	fmt.Println("\n✅ Service domain created successfully:")
	fmt.Println(string(out))
}
