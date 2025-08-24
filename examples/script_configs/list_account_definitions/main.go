// Example demonstrates listing script configuration definitions for a CacheFly account.
//
// This example shows:
// - Client initialization with API token
// - Listing available script definitions with features
// - Filtering options for hidden definitions
// - Error handling and response formatting
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

	// Initialize CacheFly client
	client := cachefly.NewClient(
		cachefly.WithToken(token),
	)

	// Prepare list options for account-level script config definitions
	opts := api.ListScriptConfigsOptions{
		IncludeFeatures: true,
		IncludeHidden:   false,
		Offset:          0,
		Limit:           20,
		ResponseType:    "", // e.g., "shallow" or "deep"
	}

	// Call ListAccountScriptConfigDefinitions (GET /scriptConfigDefinitions)
	resp, err := client.ScriptConfigs.ListAccountScriptConfigDefinitions(context.Background(), opts)
	if err != nil {
		log.Fatalf("❌ Failed to list script config definitions: %v", err)
	}

	// Pretty-print the definitions list
	out, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting definitions JSON: %v", err)
	}

	fmt.Println("\n✅ Script config definitions retrieved successfully:")
	fmt.Println(string(out))
}
