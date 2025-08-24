// Example demonstrates downloading script configuration content as a file from CacheFly.
//
// This example shows:
// - Client initialization with API token
// - Retrieving script config content by ID
// - Error handling and file content processing
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

	"github.com/cachefly/cachefly-sdk-go/pkg/cachefly"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file (optional)
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
	client := cachefly.NewClient(cachefly.WithToken(token))

	// Fetch the script config file content
	data, err := client.ScriptConfigs.GetValueAsFile(context.Background(), configID)
	if err != nil {
		log.Fatalf("❌ Failed to fetch script config file for %s: %v", configID, err)
	}

	// Pretty-print the script configuration
	out, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting script config JSON: %v", err)
	}

	fmt.Println("\n✅ Script configuration fetched successfully:")
	fmt.Println(string(out))
}
