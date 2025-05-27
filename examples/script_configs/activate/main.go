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

	resp, err := client.ScriptConfigs.ActivateByID(context.Background(), configID)
	if err != nil {
		log.Fatalf("❌ Failed to activate script config %s: %v", configID, err)
	}
	// Pretty-print
	out, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting script config JSON: %v", err)
	}

	fmt.Printf("✅ Script configuration %s activated successfully\n", configID)
	fmt.Println(string(out))
}
