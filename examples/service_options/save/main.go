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
	if err := godotenv.Load(); err != nil {
		log.Printf("⚠️ Warning: unable to load .env file: %v", err)
	}

	token := os.Getenv("CACHEFLY_API_TOKEN")
	if token == "" {
		log.Fatal("❌ CACHEFLY_API_TOKEN environment variable is required")
	}
	if len(os.Args) < 2 {
		log.Fatalf("⚠️ Usage: go run main.go <service_id>")
	}
	serviceID := os.Args[1]

	client := cachefly.NewClient(cachefly.WithToken(token))

	// Build your options payload to match exactly what works in Postman
	opts := api.ServiceOptions{
		ReverseProxy: api.ReverseProxyConfig{
			Mode:              "WEB",
			Enabled:           false,
			CacheByQueryParam: true,
			Hostname:          "www.example.com",
			OriginScheme:      "FOLLOW",
			TTL:               2678400,
			UseRobotsTXT:      true,
		},
		// Explicitly send an empty array (not null) for this field:
		MimeTypesOverrides: []api.MimeTypeOverride{},
		ExpiryHeaders:      []api.ExpiryHeader{},
	}

	updated, err := client.ServiceOptions.SaveBasicOptions(context.Background(), serviceID, opts)
	if err != nil {
		log.Fatalf("❌ Failed to save basic service options for %s: %v", serviceID, err)
	}

	out, _ := json.MarshalIndent(updated, "", "  ")
	fmt.Println("\n✅ Basic service options saved successfully:")
	fmt.Println(string(out))
}
