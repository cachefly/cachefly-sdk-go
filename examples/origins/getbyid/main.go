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

	// Read Origin ID argument
	if len(os.Args) < 2 {
		log.Fatalf("⚠️ Usage: go run main.go <origin_id>")
	}
	originID := os.Args[1]

	// Initialize CacheFly client
	client := cachefly.NewClient(
		cachefly.WithToken(token),
	)

	// Fetch origin by ID (GET /origins/{id})
	origin, err := client.Origins.GetByID(context.Background(), originID, "")
	if err != nil {
		log.Fatalf("❌ Failed to get origin %s: %v", originID, err)
	}

	// Pretty-print the origin details
	out, err := json.MarshalIndent(origin, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting origin JSON: %v", err)
	}

	fmt.Println("\n✅ Origin fetched successfully:")
	fmt.Println(string(out))
}
