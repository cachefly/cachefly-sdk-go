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
	// Load .env variables (optional)
	if err := godotenv.Load(); err != nil {
		log.Printf("⚠️ Warning: unable to load .env file: %v", err)
	}

	// Ensure API token is set
	token := os.Getenv("CACHEFLY_API_TOKEN")
	if token == "" {
		log.Fatal("❌ CACHEFLY_API_TOKEN environment variable is required")
	}

	// Read service ID and rule ID arguments
	if len(os.Args) < 3 {
		log.Fatalf("⚠️ Usage: go run main.go <service_id> <rule_id>")
	}
	serviceID := os.Args[1]
	ruleID := os.Args[2]

	// Initialize CacheFly client
	client := cachefly.NewClient(cachefly.WithToken(token))

	// Fetch the specific referer rule by ID using GetByID
	rule, err := client.ServiceOptionsRefererRules.GetByID(context.Background(), serviceID, ruleID)
	if err != nil {
		log.Fatalf("❌ Failed to get referer rule %s for service %s: %v", ruleID, serviceID, err)
	}

	// Pretty-print the rule
	out, err := json.MarshalIndent(rule, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting referer rule JSON: %v", err)
	}

	fmt.Println("\n✅ Service referer rule fetched successfully:")
	fmt.Println(string(out))
}
