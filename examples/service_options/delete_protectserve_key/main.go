package main

import (
	"context"
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
		log.Fatal("⚠️ Usage: go run main.go <service_id>")
	}
	serviceID := os.Args[1]

	// Initialize CacheFly client
	client := cachefly.NewClient(cachefly.WithToken(token))

	// Delete the ProtectServe key (DELETE /services/{id}/options/protectserveKey)
	if err := client.ServiceOptions.DeleteProtectServeKey(context.Background(), serviceID); err != nil {
		log.Fatalf("❌ Failed to delete ProtectServe key for service %s: %v", serviceID, err)
	}

	fmt.Printf("✅ ProtectServe key deleted successfully for service %s\n", serviceID)
}
