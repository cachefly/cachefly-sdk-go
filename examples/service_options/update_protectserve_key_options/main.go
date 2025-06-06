// Example demonstrates updating ProtectServe options for a CacheFly service.
//
// This example shows:
// - Client initialization with API token
// - Updating ProtectServe key and force settings
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

	// Prepare payload for updating ProtectServe key options
	opts := api.UpdateProtectServeRequest{
		ForceProtectServe: "OPTIONAL",
		ProtectServeKey:   "1921f7aae1200a5e9a3de74d4b85ed4b",
	}

	// Call UpdateProtectServeKeyOptions (PUT /services/{id}/options/protectserveKeyOptions)
	updatedOpts, err := client.ServiceOptions.UpdateProtectServeOptions(context.Background(), serviceID, opts)
	if err != nil {
		log.Fatalf("❌ Failed to update ProtectServe key options for service %s: %v", serviceID, err)
	}

	// Pretty-print the updated options
	out, err := json.MarshalIndent(updatedOpts, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting response JSON [UpdateProtectServeOptions]: %v", err)
	}

	fmt.Println("\n✅ ProtectServe key options updated successfully:")
	fmt.Println(string(out))
}
