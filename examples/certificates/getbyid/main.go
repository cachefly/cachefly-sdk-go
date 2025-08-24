// Example demonstrates retrieving a specific TLS/SSL certificate from CacheFly.
//
// This example shows:
// - Client initialization with API token
// - Fetching certificate details by ID
// - Error handling and response formatting
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
	"encoding/json"
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

	// Fetch certificate by ID (GET /certificates/{id})
	cert, err := client.Certificates.GetByID(context.Background(), certID, "")
	if err != nil {
		log.Fatalf("❌ Failed to get certificate %s: %v", certID, err)
	}

	// Pretty-print the certificate details
	out, err := json.MarshalIndent(cert, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting certificate JSON: %v", err)
	}

	fmt.Println("\n✅ Certificate fetched successfully:")
	fmt.Println(string(out))
}
