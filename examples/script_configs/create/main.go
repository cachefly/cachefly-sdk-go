package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/avvvet/cachefly-sdk-go/pkg/cachefly"
	"github.com/avvvet/cachefly-sdk-go/pkg/cachefly/api"
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
		Name:                   "example-script",
		Services:               []string{"681b3dc52715310035cb75d4"},
		ScriptConfigDefinition: "771b3dc52715310035cb75d4",
		MimeType:               "json",
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
