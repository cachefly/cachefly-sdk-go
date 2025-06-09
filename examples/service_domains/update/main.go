// Example demonstrates updating a custom domain for a CacheFly service.
//
// This example shows:
// - Client initialization with API token
// - Updating service domain description
// - Error handling and response formatting
//
// Usage:
//
//	export CACHEFLY_API_TOKEN="your-token"
//	go run main.go <service_id> <domain_id>
//
// Example:
//
//	go run main.go srv_123456789 dom_987654321

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
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("⚠️ Warning: unable to load .env file: %v", err)
	}

	// Read API token
	token := os.Getenv("CACHEFLY_API_TOKEN")
	if token == "" {
		log.Fatal("❌ CACHEFLY_API_TOKEN environment variable is required")
	}

	// Read service ID and domain ID arguments
	if len(os.Args) < 3 {
		log.Println("⚠️ Usage: go run main.go <service_id> <domain_id>")
		return
	}
	serviceID := os.Args[1]
	domainID := os.Args[2]

	// Initialize CacheFly client
	client := cachefly.NewClient(
		cachefly.WithToken(token),
	)

	// Prepare update payload for service domain
	payload := api.UpdateServiceDomainRequest{
		//Name: "updated.example.com",
		Description: "update service domain from SDK",
	}

	// Call Update service domain by ID
	updatedDomain, err := client.ServiceDomains.UpdateByID(context.Background(), serviceID, domainID, payload)
	if err != nil {
		log.Fatalf("❌ Failed to update domain %s for service %s: %v", domainID, serviceID, err)
	}

	// Pretty-print the updated domain
	out, err := json.MarshalIndent(updatedDomain, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting updated domain JSON: %v", err)
	}

	fmt.Println("\n✅ Service domain updated successfully:")
	fmt.Println(string(out))
}
