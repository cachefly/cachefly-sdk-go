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

	// Read API token from environment
	token := os.Getenv("CACHEFLY_API_TOKEN")
	if token == "" {
		log.Fatal("❌ CACHEFLY_API_TOKEN environment variable is required")
	}

	// Read Service ID argument
	if len(os.Args) < 2 {
		log.Println("⚠️ Usage: go run main.go <service_id>")
		return
	}
	serviceID := os.Args[1]

	// Initialize CacheFly client
	client := cachefly.NewClient(
		cachefly.WithToken(token),
	)

	// Prepare list options
	opts := api.ListServiceRulesOptions{
		Offset:       0,
		Limit:        10,
		ResponseType: "", // "shallow" or "full"
	}

	// Call List service rules
	resp, err := client.ServiceRules.List(context.Background(), serviceID, opts)
	if err != nil {
		log.Fatalf("❌ Failed to list service rules for service %s: %v", serviceID, err)
	}

	// Pretty-print the result
	out, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting service rules list JSON: %v", err)
	}

	fmt.Println("\n✅ Service rules retrieved successfully:")
	fmt.Println(string(out))
}
