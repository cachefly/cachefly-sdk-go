package main

import (
	"context"
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

	// Read API token from environment
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

	// Delete the referer rule by ID
	if err := client.ServiceOptionsRefererRules.Delete(context.Background(), serviceID, ruleID); err != nil {
		log.Fatalf("❌ Failed to delete referer rule %s for service %s: %v", ruleID, serviceID, err)
	}

	fmt.Printf("✅ Service referer rule %s deleted successfully for service %s\n", ruleID, serviceID)
}
