package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/avvvet/cachefly-sdk-go/pkg/cachefly"
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

	// Fetch the JSON schema for service rules
	schemaResp, err := client.ServiceRules.GetSchema(context.Background(), serviceID)
	if err != nil {
		log.Fatalf("❌ Failed to fetch service rules JSON schema for service %s: %v", serviceID, err)
	}

	// Pretty-print the schema
	out, err := json.MarshalIndent(schemaResp, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting schema JSON: %v", err)
	}

	fmt.Println("\n✅ Service rules JSON schema retrieved successfully:")
	fmt.Println(string(out))
}
