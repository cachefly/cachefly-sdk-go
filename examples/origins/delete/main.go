// Example demonstrates deleting an origin server from CacheFly.
//
// This example shows:
// - Client initialization with API token
// - Deleting an origin by ID
// - Error handling for delete operations
//
// Usage:
//
//	export CACHEFLY_API_TOKEN="your-token"
//	go run main.go <origin_id>
//
// Example:
//
//	go run main.go org_123456789

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

	// Read Origin ID argument
	if len(os.Args) < 2 {
		log.Fatalf("⚠️ Usage: go run main.go <origin_id>")
	}
	originID := os.Args[1]

	// Initialize CacheFly client
	client := cachefly.NewClient(
		cachefly.WithToken(token),
	)

	// Delete origin by ID (DELETE /origins/{id})
	if err := client.Origins.Delete(context.Background(), originID); err != nil {
		log.Fatalf("❌ Failed to delete origin %s: %v", originID, err)
	}

	fmt.Printf("✅ Origin %s deleted successfully\n", originID)
}
