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

	// Prepare list options (offset, limit)
	opts := api.ListCertificatesOptions{
		Offset: 0,
		Limit:  10,
	}

	// Call List (GET /certificates)
	resp, err := client.Certificates.List(context.Background(), opts)
	if err != nil {
		log.Fatalf("❌ Failed to list certificates: %v", err)
	}

	// Pretty-print the certificates list
	out, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting certificates list JSON: %v", err)
	}

	fmt.Println("\n✅ Certificates retrieved successfully:")
	fmt.Println(string(out))
}
