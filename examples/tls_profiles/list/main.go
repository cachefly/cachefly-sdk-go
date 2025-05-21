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
	// Load environment variables (optional)
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

	// Prepare list options for TLS profiles
	opts := api.ListTLSProfilesOptions{
		Offset: 0,
		Limit:  10,
	}

	// Call List (GET /tlsProfiles)
	resp, err := client.TLSProfiles.List(context.Background(), opts)
	if err != nil {
		log.Fatalf("❌ Failed to list TLS profiles: %v", err)
	}

	// Pretty-print the TLS profiles list
	out, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting TLS profiles list JSON: %v", err)
	}

	fmt.Println("\n✅ TLS profiles retrieved successfully:")
	fmt.Println(string(out))
}
