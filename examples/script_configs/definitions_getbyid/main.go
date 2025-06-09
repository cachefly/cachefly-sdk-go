// Example demonstrates retrieving a script configuration definition from CacheFly.
//
// This example shows:
// - Client initialization with API token
// - Fetching script definition details by ID
// - Error handling and response formatting
//
// Usage:
//
//	export CACHEFLY_API_TOKEN="your-token"
//	go run main.go <definition_id>
//
// Example:
//
//	go run main.go def_123456789

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
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("⚠️ Warning: unable to load .env file: %v", err)
	}

	// Ensure API token is set
	token := os.Getenv("CACHEFLY_API_TOKEN")
	if token == "" {
		log.Fatal("❌ CACHEFLY_API_TOKEN environment variable is required")
	}

	// Read Definition ID argument
	if len(os.Args) < 2 {
		log.Fatalf("⚠️ Usage: go run main.go <definition_id>")
	}
	defID := os.Args[1]

	// Initialize CacheFly client
	client := cachefly.NewClient(cachefly.WithToken(token))

	// Fetch the specific script config definition by ID
	def, err := client.ScriptConfigs.GetDefinitionByID(context.Background(), defID)
	if err != nil {
		log.Fatalf("❌ Failed to get script config definition %s: %v", defID, err)
	}

	// Pretty-print the definition
	out, err := json.MarshalIndent(def, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting definition JSON: %v", err)
	}

	fmt.Println("\n✅ Script config definition fetched successfully:")
	fmt.Println(string(out))
}
