// Example demonstrates enabling access logging for a CacheFly service.
//
// This example shows:
// - Client initialization with API token
// - Configuring access logs with target destination
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
	api "github.com/cachefly/cachefly-sdk-go/pkg/cachefly/api/v2_6"
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

	payload := api.EnableAccessLogsRequest{
		LogTarget: "66320d4208158b00411703e4",
	}

	client := cachefly.NewClient(
		cachefly.WithToken(token),
	)

	resp, err := client.Services.EnableAccessLogging(context.Background(), serviceID, payload)
	if err != nil {
		log.Fatalf("❌ Failed to enable access logging: %v", err)
	}

	dataJSON, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatalf("Error formatting MarshalIndent [EnableAccessLogging]: %v", err)
	}
	fmt.Println("\n ✅ Access logging enabled successfully.")
	fmt.Println(string(dataJSON))
}
