package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/avvvet/cachefly-sdk-go/pkg/cachefly"
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

	// Ensure service ID is provided
	if len(os.Args) < 2 {
		log.Println("⚠️ Usage: go run main.go <service_id>")
		return
	}
	serviceID := os.Args[1]

	client := cachefly.NewClient(
		cachefly.WithToken(token),
	)

	resp, err := client.Services.DeleteAccessLoggingByID(context.Background(), serviceID)
	if err != nil {
		log.Fatalf("❌ Failed to disabled access logging: %v", err)
	}

	// Success message
	dataJSON, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatalf("Error formatting MarshalIndent [DeleteAccessLoggingByID]: %v", err)
	}
	fmt.Println("\n ✅ Access logging disabled successfully.")
	fmt.Println(string(dataJSON))
}
