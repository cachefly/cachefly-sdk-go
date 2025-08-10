// Example demonstrates retrieving a specific script definition from CacheFly.
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

	// Read API token
	token := os.Getenv("CACHEFLY_API_TOKEN")
	if token == "" {
		log.Fatal("❌ CACHEFLY_API_TOKEN environment variable is required")
	}

	// Read Script Definition ID argument
	if len(os.Args) < 2 {
		log.Println("⚠️ Usage: go run main.go <definition_id>")
		return
	}
	definitionID := os.Args[1]

	// Initialize CacheFly client
	client := cachefly.NewClient(
		cachefly.WithToken(token),
	)

	// Fetch specific script definition by ID
	def, err := client.ScriptDefinitions.GetByID(context.Background(), definitionID)
	if err != nil {
		log.Fatalf("❌ Failed to get script definition %s: %v", definitionID, err)
	}

	// Pretty-print the script definition
	out, err := json.MarshalIndent(def, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting script definition JSON: %v", err)
	}

	fmt.Println("\n✅ Script definition fetched successfully:")
	fmt.Println(string(out))
}
