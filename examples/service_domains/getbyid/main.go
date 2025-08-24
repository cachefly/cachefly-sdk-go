// Example demonstrates retrieving a specific domain from a CacheFly service.
//
// This example shows:
// - Client initialization with API token
// - Fetching service domain details by ID
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

	"github.com/cachefly/cachefly-sdk-go/pkg/cachefly"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("⚠️ Warning: unable to load .env file: %v", err)
	}

	// Read token
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

	// Call GetByID endpoint for service domains
	domain, err := client.ServiceDomains.GetByID(context.Background(), serviceID, domainID, "")
	if err != nil {
		log.Fatalf("❌ Failed to get domain %s for service %s: %v", domainID, serviceID, err)
	}

	// Marshal and print the domain
	out, err := json.MarshalIndent(domain, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting domain JSON: %v", err)
	}

	fmt.Println("\n✅ Service domain fetched successfully:")
	fmt.Println(string(out))
}
