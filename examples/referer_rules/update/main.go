// Example demonstrates updating a referer rule for a CacheFly service.
//
// This example shows:
// - Client initialization with API token
// - Updating referer rule configuration
// - Modifying allowed domains and file extensions
// - Error handling and response formatting
//
// Usage:
//
//	export CACHEFLY_API_TOKEN="your-token"
//	go run main.go <service_id> <rule_id>
//
// Example:
//
//	go run main.go srv_123456789 rule_987654321

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

	// Prepare payload for updating the referer rule
	payload := api.UpdateRefererRuleRequest{
		Directory:     "/images",
		Extension:     "png",
		Exceptions:    []string{"trusted.example.com"},
		DefaultAction: "ALLOW", // or "DENY"
	}

	// Update the referer rule by ID
	updatedRule, err := client.ServiceOptionsRefererRules.Update(context.Background(), serviceID, ruleID, payload)
	if err != nil {
		log.Fatalf("❌ Failed to update referer rule %s for service %s: %v", ruleID, serviceID, err)
	}

	// Pretty-print the updated rule
	out, err := json.MarshalIndent(updatedRule, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting updated referer rule JSON: %v", err)
	}

	fmt.Println("\n✅ Service referer rule updated successfully:")
	fmt.Println(string(out))
}
