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

	if err := godotenv.Load(); err != nil {
		log.Printf("⚠️ Warning: unable to load .env file: %v", err)
	}

	token := os.Getenv("CACHEFLY_API_TOKEN")
	if token == "" {
		log.Fatal("❌ CACHEFLY_API_TOKEN environment variable is required")
	}

	if len(os.Args) < 2 {
		log.Println("⚠️ Usage: go run main.go <service_id>")
		return
	}
	serviceID := os.Args[1]

	client := cachefly.NewClient(
		cachefly.WithToken(token),
	)

	opts := api.ListServiceDomainsOptions{
		Search:       "",
		Offset:       0,
		Limit:        10,
		ResponseType: "",
	}

	resp, err := client.ServiceDomains.List(context.Background(), serviceID, opts)
	if err != nil {
		log.Fatalf("❌ Failed to list domains for service %s: %v", serviceID, err)
	}

	out, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting domains list JSON: %v", err)
	}

	fmt.Println("\n✅ Service domains retrieved successfully:")
	fmt.Println(string(out))
}
