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

	// Read API token
	token := os.Getenv("CACHEFLY_API_TOKEN")
	if token == "" {
		log.Fatal("❌ CACHEFLY_API_TOKEN environment variable is required")
	}

	// Read Certificate ID argument
	if len(os.Args) < 2 {
		log.Fatalf("⚠️ Usage: go run main.go <certificate_id>")
	}
	certID := os.Args[1]

	// Initialize CacheFly client
	client := cachefly.NewClient(
		cachefly.WithToken(token),
	)

	// Delete certificate by ID (DELETE /certificates/{id})
	if err := client.Certificates.Delete(context.Background(), certID); err != nil {
		log.Fatalf("❌ Failed to delete certificate %s: %v", certID, err)
	}

	fmt.Printf("✅ Certificate %s deleted successfully\n", certID)
}
