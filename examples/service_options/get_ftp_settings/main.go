// Example demonstrates retrieving FTP settings for a CacheFly service.
//
// This example shows:
// - Client initialization with API token
// - Fetching FTP configuration including password
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

	// Call GetFTPSettings (GET /services/{id}/options/ftp)
	ftpSettings, err := client.ServiceOptions.GetFTPSettings(context.Background(), serviceID, false)
	if err != nil {
		log.Fatalf("❌ Failed to fetch FTP settings for service %s: %v", serviceID, err)
	}

	// Pretty-print the FTP settings
	out, err := json.MarshalIndent(ftpSettings, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting FTP settings JSON: %v", err)
	}

	fmt.Println("\n✅ FTP settings retrieved successfully:")
	fmt.Println(string(out))
}
