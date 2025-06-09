// Example demonstrates enabling origin logging for a CacheFly service.
//
// This example shows:
// - Client initialization with API token
// - Configuring origin logs with target destination
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

	"github.com/cachefly/cachefly-go-sdk/pkg/cachefly"
	api "github.com/cachefly/cachefly-go-sdk/pkg/cachefly/api/v2_5"
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

	payload := api.EnableOriginLogsRequest{
		LogTarget: "stringstringstringstring",
	}

	updatedService, err := client.Services.EnableOriginLogging(context.Background(), serviceID, payload)
	if err != nil {
		log.Fatalf("❌ Failed to enable origin logging: %v", err)
	}

	out, err := json.MarshalIndent(updatedService, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting response JSON [EnableOriginLogging]: %v", err)
	}

	fmt.Println("\n✅ Origin logging enabled successfully:")
	fmt.Println(string(out))
}
