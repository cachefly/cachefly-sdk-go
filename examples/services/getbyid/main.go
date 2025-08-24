// Example demonstrates retrieving a specific CacheFly service by ID.
//
// This example shows:
// - Client initialization with API token
// - Fetching service details by ID
// - Error handling and response formatting
//
// Usage:
//
//	export CACHEFLY_API_TOKEN="your-token"
//	go run main.go <service_id>
//
// Example:
//
//	go run main.go srv_123456789

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/cachefly/cachefly-sdk-go/pkg/cachefly"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Printf("⚠️ Warning: unable to load .env file: %v", err)
	}

	token := os.Getenv("CACHEFLY_API_TOKEN")
	if token == "" {
		log.Fatal("❌ CACHEFLY_API_TOKEN environment variable is required")
	}

	if len(os.Args) < 2 {
		log.Println("⚠️ Usage: go run main.go <service_id>")
		return
	}
	serviceID := os.Args[1]

	client := cachefly.NewClient(
		cachefly.WithToken(token),
	)

	service, err := client.Services.GetByID(context.Background(), serviceID)
	if err != nil {
		log.Fatalf("❌ Failed to get service by ID: %v", err)
	}

	out, err := json.MarshalIndent(service, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting service JSON: %v", err)
	}

	fmt.Println("\n✅ Service fetched successfully:")
	fmt.Println(string(out))
}
