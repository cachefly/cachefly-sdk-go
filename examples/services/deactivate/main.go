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

	if len(os.Args) < 2 {
		log.Println("⚠️ Usage: go run main.go <service_id>")
		return
	}
	serviceID := os.Args[1]

	client := cachefly.NewClient(
		cachefly.WithToken(token),
	)

	resp, err := client.Services.DeactivateServiceByID(context.Background(), serviceID)
	if err != nil {
		log.Fatalf("❌ Failed to deactivate service: %v", err)
	}

	dataJSON, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatalf("Error formatting MarshalIndent [DeactivateServiceByID]: %v", err)
	}

	fmt.Println("\n ✅ Service deactivated successfully.")
	fmt.Println(string(dataJSON))
}
