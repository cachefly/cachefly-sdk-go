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

	// Prepare payload for new referer rule
	payload := api.CreateRefererRuleRequest{
		Directory:     "/images",
		Extension:     "jpg",
		Exceptions:    []string{"trusted.example.com"},
		DefaultAction: "ALLOW", // or "DENY"
	}

	// Call Create referer rule (POST /services/{id}/options/refererRules)
	rule, err := client.ServiceOptionsRefererRules.Create(context.Background(), serviceID, payload)
	if err != nil {
		log.Fatalf("❌ Failed to create referer rule for service %s: %v", serviceID, err)
	}

	// Pretty-print the created rule
	out, err := json.MarshalIndent(rule, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting referer rule JSON: %v", err)
	}

	fmt.Println("\n✅ Service referer rule created successfully:")
	fmt.Println(string(out))
}
