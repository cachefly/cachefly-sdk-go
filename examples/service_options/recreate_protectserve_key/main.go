// Example demonstrates regenerating the ProtectServe key for a CacheFly service.
//
// This example shows:
// - Client initialization with API token
// - Regenerating a new ProtectServe key
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
		log.Fatal("⚠️ Usage: go run main.go <service_id>")
	}
	serviceID := os.Args[1]

	// Initialize CacheFly client
	client := cachefly.NewClient(cachefly.WithToken(token))

	// Regenerate ProtectServe key (PUT /services/{id}/options/protectserveKey)
	newKey, err := client.ServiceOptions.RecreateProtectServeKey(context.Background(), serviceID, "")
	if err != nil {
		log.Fatalf("❌ Failed to regenerate ProtectServe key for service %s: %v", serviceID, err)
	}

	// Print the new key
	fmt.Println("✅ ProtectServe key regenerated successfully:")
	jsonData, err := json.MarshalIndent(newKey, "", " ")
	if err != nil {
		log.Fatalf("Error formatting MarshalIndent [RecreateProtectServeKey]: %v", err)
	}
	fmt.Println(string(jsonData))
}
