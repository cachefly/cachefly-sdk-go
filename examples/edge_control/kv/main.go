// Example demonstrates managing the edge control KV store of a service.
//
// This example shows:
// - Replacing the service-level KV store (values: strings, numbers, booleans)
// - Reading the account and service stores merged, as edge scripts see them
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
		log.Printf("Warning: unable to load .env file: %v", err)
	}

	token := os.Getenv("CACHEFLY_API_TOKEN")
	if token == "" {
		log.Fatal("CACHEFLY_API_TOKEN environment variable is required")
	}

	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <service_id>")
	}
	serviceID := os.Args[1]

	client := cachefly.NewClient(
		cachefly.WithToken(token),
	)
	ctx := context.Background()

	kv, err := client.EdgeControlKV.ReplaceService(ctx, serviceID, map[string]interface{}{
		"region":      "eu-west",
		"maxAge":      3600,
		"maintenance": false,
	})
	if err != nil {
		log.Fatalf("Failed to replace service KV: %v", err)
	}
	fmt.Printf("Service KV now has %d keys (updated %s)\n", len(kv.Data), kv.UpdatedAt)

	merged, err := client.EdgeControlKV.GetMerged(ctx, serviceID)
	if err != nil {
		log.Fatalf("Failed to read merged KV: %v", err)
	}

	out, err := json.MarshalIndent(merged, "", "  ")
	if err != nil {
		log.Fatalf("Error formatting KV JSON: %v", err)
	}
	fmt.Println("\nMerged KV (service keys override account keys):")
	fmt.Println(string(out))
}
