// Example demonstrates creating a new script configuration in CacheFly.
//
// This example shows:
// - Client initialization with API token
// - Creating a script config with service associations
// - Setting script definition and MIME type
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

	// Prepare payload for creating a new script configuration
	opts := api.CreateScriptConfigRequest{
		Name:                   "advanced-cache-config-from-sdk",
		Services:               []string{"681e55412d0cc10041680dc7"},
		ScriptConfigDefinition: "64de825631861c0035b75708",
		MimeType:               "text/yaml",
	}

	// Call Create (POST /services/{id}/scriptConfigs)
	config, err := client.ScriptConfigs.Create(context.Background(), opts)
	if err != nil {
		log.Fatalf("❌ Failed to create script config: %v", err)
	}

	// Pretty-print the created script configuration
	out, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting script config JSON: %v", err)
	}

	fmt.Println("\n✅ Script configuration created successfully")
	fmt.Println(string(out))
}
