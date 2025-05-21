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

	// Call ListPromo with includeFeatures = true
	defs, err := client.ScriptConfigs.ListPromo(context.Background(), true)
	if err != nil {
		log.Fatalf("❌ Failed to list promo script config definitions: %v", err)
	}

	// Pretty-print the definitions
	out, err := json.MarshalIndent(defs, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting definitions JSON: %v", err)
	}

	fmt.Println("\n✅ Promo script config definitions retrieved successfully:")
	fmt.Println(string(out))
}
