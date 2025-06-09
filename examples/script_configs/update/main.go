// Example demonstrates updating a script configuration in CacheFly.
//
// This example shows:
// - Client initialization with API token
// - Updating script config with URL redirects
// - Configuring 301 redirect mappings
// - Error handling and response formatting
//
// Usage:
//
//	export CACHEFLY_API_TOKEN="your-token"
//	go run main.go <config_id>
//
// Example:
//
//	go run main.go cfg_123456789

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

	// Read Script Config ID argument
	if len(os.Args) < 2 {
		log.Fatalf("⚠️ Usage: go run main.go <config_id>")
	}
	configID := os.Args[1]

	// Initialize CacheFly client
	client := cachefly.NewClient(
		cachefly.WithToken(token),
	)

	// Prepare payload for updating the script configuration
	opts := api.UpdateScriptConfigRequest{
		Name:                   "url-redirects-updated-sdk",
		Services:               []string{"681b3dc52715310035cb75d4"},
		ScriptConfigDefinition: "63fcfcc58a797a005f2ad04e",
		MimeType:               "text/json",
		Value: map[string]map[string]string{
			"301": {
				"/old/path/to/file.jpg":  "https://www.sdk.com/path/to/new/file.jpg",
				"/old/path/to/file2.jpg": "https://www.sdk.com/path/to/some/other/file.jpg",
			},
		},
	}

	// Call Update (PUT /scriptConfigs/{id})
	updatedConfig, err := client.ScriptConfigs.UpdateByID(context.Background(), configID, opts)
	if err != nil {
		log.Fatalf("❌ Failed to update script config %s: %v", configID, err)
	}

	// Pretty-print the updated script configuration
	out, err := json.MarshalIndent(updatedConfig, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting updated script config JSON: %v", err)
	}

	fmt.Println("\n✅ Script configuration updated successfully (ID:", configID, "):")
	fmt.Println(string(out))
}
