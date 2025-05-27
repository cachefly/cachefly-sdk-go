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

	// Read Origin ID argument
	if len(os.Args) < 2 {
		log.Fatalf("⚠️ Usage: go run main.go <origin_id>")
	}
	originID := os.Args[1]

	// Initialize CacheFly client
	client := cachefly.NewClient(
		cachefly.WithToken(token),
	)

	// Prepare payload for updating the origin
	opts := api.UpdateOriginRequest{
		Name:     "updated-origin",
		Hostname: "updated-new-origin.example.com",
		Type:     "WEB",
		Scheme:   "HTTP",
		TTL:      2678400,
	}

	// Call Update (PUT /origins/{id})
	updatedOrigin, err := client.Origins.UpdateByID(context.Background(), originID, opts)
	if err != nil {
		log.Fatalf("❌ Failed to update origin %s: %v", originID, err)
	}

	// Pretty-print the updated origin
	out, err := json.MarshalIndent(updatedOrigin, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting updated origin JSON: %v", err)
	}

	fmt.Println("\n✅ Origin updated successfully:")
	fmt.Println(string(out))
}
