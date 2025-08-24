// Example demonstrates signaling that a domain is ready for validation in CacheFly.
//
// This example shows:
// - Client initialization with API token
// - Triggering domain validation process
// - Error handling for validation operations
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
	"fmt"
	"log"
	"os"

	"github.com/cachefly/cachefly-sdk-go/pkg/cachefly"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("⚠️ Warning: unable to load .env file: %v", err)
	}

	// Read API token from environment
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

	// Signal domain validation ready
	if _, err := client.ServiceDomains.ValidationReady(context.Background(), serviceID, domainID); err != nil {
		log.Fatalf("❌ Failed to signal validation ready for domain %s on service %s: %v", domainID, serviceID, err)
	}

	fmt.Printf("✅ Domain %s on service %s signaled validation ready successfully.\n", domainID, serviceID)
}
