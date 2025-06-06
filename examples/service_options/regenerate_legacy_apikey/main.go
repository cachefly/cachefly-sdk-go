// Example demonstrates regenerating the legacy API key for a CacheFly service.
//
// This example shows:
// - Client initialization with API token
// - Creating a new legacy API key for a service
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

	// Regenerate legacy API key for the service (PUT /services/{id}/options/apikey)
	newKey, err := client.ServiceOptions.RegenerateLegacyAPIKey(context.Background(), serviceID)
	if err != nil {
		log.Fatalf("❌ Failed to regenerate legacy API key for service %s: %v", serviceID, err)
	}

	// Print the new key
	fmt.Println("✅ Legacy API key regenerated successfully:")
	jsonData, err := json.MarshalIndent(newKey, "", " ")
	if err != nil {
		log.Fatalf("Error formatting MarshalIndent [GetLegacyAPIKey]: %v", err)
	}
	fmt.Println(string(jsonData))
}
