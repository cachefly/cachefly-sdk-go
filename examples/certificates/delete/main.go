// Example demonstrates deleting a TLS/SSL certificate from CacheFly.
//
// This example shows:
// - Client initialization with API token
// - Deleting a certificate by ID
// - Error handling for delete operations
//
// Usage:
//
//	export CACHEFLY_API_TOKEN="your-token"
//	go run main.go <certificate_id>
//
// Example:
//
//	go run main.go cert_123456789

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/cachefly/cachefly-go-sdk/pkg/cachefly"
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

	// Read Certificate ID argument
	if len(os.Args) < 2 {
		log.Fatalf("⚠️ Usage: go run main.go <certificate_id>")
	}
	certID := os.Args[1]

	// Initialize CacheFly client
	client := cachefly.NewClient(
		cachefly.WithToken(token),
	)

	// Delete certificate by ID (DELETE /certificates/{id})
	if err := client.Certificates.Delete(context.Background(), certID); err != nil {
		log.Fatalf("❌ Failed to delete certificate %s: %v", certID, err)
	}

	fmt.Printf("✅ Certificate %s deleted successfully\n", certID)
}
