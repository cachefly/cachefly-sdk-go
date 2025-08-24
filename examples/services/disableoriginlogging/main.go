// Example demonstrates disabling origin logging for a CacheFly service.
//
// This example shows:
// - Client initialization with API token
// - Removing origin logging configuration
// - Error handling for delete operations
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

	resp, err := client.Services.DeleteOriginLoggingByID(context.Background(), serviceID)
	if err != nil {
		log.Fatalf("❌ Failed to disable origin logging: %v", err)
	}

	dataJSON, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatalf("Error formatting MarshalIndent [DeleteOriginLoggingByID]: %v", err)
	}
	fmt.Println("\n ✅ Origin logging disabled successfully.")
	fmt.Println(string(dataJSON))
}
