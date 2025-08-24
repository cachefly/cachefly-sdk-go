// Example demonstrates retrieving a specific script configuration from CacheFly.
//
// This example shows:
// - Client initialization with API token
// - Fetching script config details by ID
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

	// Read Service ID and Script Config ID arguments
	if len(os.Args) < 1 {
		log.Fatalf("⚠️ Usage: go run main.go <config_id>")
	}
	configID := os.Args[1]

	// Initialize CacheFly client
	client := cachefly.NewClient(
		cachefly.WithToken(token),
	)

	// Fetch specific script configuration by ID
	cfg, err := client.ScriptConfigs.GetByID(context.Background(), configID, "")
	if err != nil {
		log.Fatalf("❌ Failed to get script config %s : %v", configID, err)
	}

	// Pretty-print the script configuration
	out, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting script config JSON: %v", err)
	}

	fmt.Println("\n✅ Script configuration fetched successfully:")
	fmt.Println(string(out))
}
