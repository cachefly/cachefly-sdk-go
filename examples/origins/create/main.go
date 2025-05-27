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

	// Prepare payload for creating a new origin
	opts := api.CreateOriginRequest{
		Name:     "example-origin",
		Hostname: "origin.example.com",
		Type:     "WEB",
		Scheme:   "HTTP",
		TTL:      2678400,
	}

	// Call Create (POST /origins)
	origin, err := client.Origins.Create(context.Background(), opts)
	if err != nil {
		log.Fatalf("❌ Failed to create origin: %v", err)
	}

	// Pretty-print the created origin
	out, err := json.MarshalIndent(origin, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting origin JSON: %v", err)
	}

	fmt.Println("\n✅ Origin created successfully:")
	fmt.Println(string(out))
}
