package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/avvvet/cachefly-sdk-go/pkg/cachefly"
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
	client := cachefly.NewClient(cachefly.WithToken(token))

	// Fetch the JSON schema for the specified script config
	schema, err := client.ScriptConfigs.GetSchemaByID(context.Background(), configID)
	if err != nil {
		log.Fatalf("❌ Failed to fetch schema for script config %s: %v", configID, err)
	}

	// Pretty-print the schema
	out, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting schema JSON: %v", err)
	}

	fmt.Println("\n✅ Script config schema retrieved successfully for ID:", configID)
	fmt.Println(string(out))
}
