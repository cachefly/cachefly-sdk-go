// Example demonstrates checking domain name availability.
//
// Usage:
//
//	export CACHEFLY_API_TOKEN="your-token"
//	go run main.go example.com
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/cachefly/cachefly-sdk-go/pkg/cachefly"
	api "github.com/cachefly/cachefly-sdk-go/pkg/cachefly/api/v2_6"
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
		log.Fatal("Usage: go run main.go <domain>")
	}
	name := os.Args[1]

	client := cachefly.NewClient(cachefly.WithToken(token))

	available, err := client.Availability.Domains(context.Background(), api.CheckDomainRequest{Name: name})
	if err != nil {
		log.Fatalf("❌ Availability check failed: %v", err)
	}

	if available {
		fmt.Printf("✅ Domain '%s' is available\n", name)
	} else {
		fmt.Printf("❌ Domain '%s' is not available\n", name)
	}
}
