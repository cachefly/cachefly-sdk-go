// Example demonstrates working with Script Definitions via the CacheFly SDK.
//
// This example shows:
// - Client initialization with API token
// - Listing account Script Definitions
//
// Usage:
//
//	export CACHEFLY_API_TOKEN="your-token"
//	go run main.go
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
	// Load environment variables (optional)
	if err := godotenv.Load(); err != nil {
		log.Printf("⚠️ Warning: unable to load .env file: %v", err)
	}

	// Ensure API token is set
	token := os.Getenv("CACHEFLY_API_TOKEN")
	if token == "" {
		log.Fatal("❌ CACHEFLY_API_TOKEN environment variable is required")
	}

	// Initialize CacheFly client
	client := cachefly.NewClient(cachefly.WithToken(token))

	// List definitions
	opts := api.ListScriptDefinitionsOptions{
		IncludeFeatures: true,
		IncludeHidden:   false,
		Offset:          0,
		Limit:           20,
		ResponseType:    "", // optional
	}

	resp, err := client.ScriptDefinitions.List(context.Background(), opts)
	if err != nil {
		log.Fatalf("❌ Failed to list script definitions: %v", err)
	}

	out, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting list JSON: %v", err)
	}

	fmt.Println("\n✅ Script definitions retrieved successfully:")
	fmt.Println(string(out))
}
