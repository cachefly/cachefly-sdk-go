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

	// Prepare update payload for service rules, Refer API doc
	payload := api.UpdateServiceRulesRequest{
		Rules: []api.ServiceRule{
			{
				ID: "",
			},
			{
				ID: "",
			},
		},
	}

	// Call Update service rules (PUT /services/{id}/rules)
	updated, err := client.ServiceRules.Update(context.Background(), serviceID, payload)
	if err != nil {
		log.Fatalf("❌ Failed to update service rules for %s: %v", serviceID, err)
	}

	// Pretty-print the updated rules response
	out, err := json.MarshalIndent(updated, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting updated rules JSON: %v", err)
	}

	fmt.Println("\n✅ Service rules updated successfully:")
	fmt.Println(string(out))
}
