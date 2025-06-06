// Example demonstrates updating basic options for a CacheFly service.
//
// This example shows:
// - Client initialization with API token
// - Configuring reverse proxy settings
// - Setting error TTL and connection timeout
// - Error handling and response formatting
//
// Usage:
//
//	export CACHEFLY_API_TOKEN="your-token"
//	go run main.go <service_id>
//
// Example:
//
//	go run main.go srv_123456789

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
		log.Printf("⚠️  Warning: unable to load .env file: %v", err)
	}

	token := os.Getenv("CACHEFLY_API_TOKEN")
	if token == "" {
		log.Fatal("❌ CACHEFLY_API_TOKEN environment variable is required")
	}

	if len(os.Args) < 2 {
		log.Fatalf("⚠️  Usage: go run main.go <service_id>")
	}
	serviceID := os.Args[1]

	client := cachefly.NewClient(cachefly.WithToken(token))

	// Configure service options
	opts := api.ServiceOptions{
		ReverseProxy: api.ReverseProxyConfig{
			Mode:              "WEB",
			Enabled:           true,
			CacheByQueryParam: true,
			Hostname:          "www.example.com",
			OriginScheme:      "FOLLOW",
			TTL:               2678400,
			UseRobotsTXT:      true,
		},
		ErrorTTL: api.Option{
			Enabled: true,
			Value:   120,
		},
		ConTimeout: api.Option{
			Enabled: true,
			Value:   5,
		},
		NoCache:            true,
		MimeTypesOverrides: []api.MimeTypeOverride{},
		ExpiryHeaders:      []api.ExpiryHeader{},
	}

	// Save service options
	updated, err := client.ServiceOptions.SaveBasicOptions(context.Background(), serviceID, opts)
	if err != nil {
		log.Fatalf("❌ Failed to save basic service options for %s: %v", serviceID, err)
	}

	out, _ := json.MarshalIndent(updated, "", "  ")
	fmt.Println("✅ Basic service options saved successfully:")
	fmt.Println(string(out))
}
