// Example demonstrates deactivating a SAML configuration.
//
// Usage:
//
//	export CACHEFLY_API_TOKEN="your-token"
//	go run main.go <saml_id>
//
// Example:
//
//	go run main.go 0123456789abcdef01234567
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/cachefly/cachefly-sdk-go/pkg/cachefly"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("⚠️  Warning: unable to load .env file: %v", err)
	}

	token := os.Getenv("CACHEFLY_API_TOKEN")
	if token == "" {
		log.Fatal("❌ CACHEFLY_API_TOKEN environment variable is required")
	}

	if len(os.Args) < 2 {
		log.Println("⚠️  Usage: go run main.go <saml_id>")
		return
	}
	samlID := os.Args[1]

	client := cachefly.NewClient(
		cachefly.WithToken(token),
	)

	fmt.Printf("🔄 Deactivating SAML configuration: %s\n", samlID)
	if err := client.SAML.DeactivateByID(context.Background(), samlID); err != nil {
		log.Fatalf("❌ Failed to deactivate SAML configuration: %v", err)
	}

	fmt.Println("✅ SAML configuration deactivated successfully.")
}
