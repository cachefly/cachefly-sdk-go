// Example demonstrates updating an origin server in CacheFly.
//
// This example shows:
// - Client initialization with API token
// - Updating origin configuration including hostname and TTL
// - Error handling and response formatting
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

	// Read Origin ID argument
	if len(os.Args) < 2 {
		log.Fatalf("⚠️ Usage: go run main.go <origin_id>")
	}
	originID := os.Args[1]

	// Initialize CacheFly client
	client := cachefly.NewClient(
		cachefly.WithToken(token),
	)

	// Prepare payload for updating the origin
	name := "updated-origin"
	hostname := "updated-new-origin.example.com"
	originType := "WEB"
	scheme := "HTTP"
	ttl := int32(2678400)

	opts := api.UpdateOriginRequest{
		Name:     &name,
		Hostname: &hostname,
		Type:     &originType,
		Scheme:   &scheme,
		TTL:      &ttl,
	}

	// Call Update (PUT /origins/{id})
	updatedOrigin, err := client.Origins.UpdateByID(context.Background(), originID, opts)
	if err != nil {
		log.Fatalf("❌ Failed to update origin %s: %v", originID, err)
	}

	// Pretty-print the updated origin
	out, err := json.MarshalIndent(updatedOrigin, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting updated origin JSON: %v", err)
	}

	fmt.Println("\n✅ Origin updated successfully:")
	fmt.Println(string(out))
}
