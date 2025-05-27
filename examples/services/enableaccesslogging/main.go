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
		log.Println("⚠️ Usage: go run main.go <service_id>")
		return
	}
	serviceID := os.Args[1]

	payload := api.EnableAccessLogsRequest{
		LogTarget: "66320d4208158b00411703e4",
	}

	client := cachefly.NewClient(
		cachefly.WithToken(token),
	)

	resp, err := client.Services.EnableAccessLogging(context.Background(), serviceID, payload)
	if err != nil {
		log.Fatalf("❌ Failed to enable access logging: %v", err)
	}

	dataJSON, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatalf("Error formatting MarshalIndent [EnableAccessLogging]: %v", err)
	}
	fmt.Println("\n ✅ Access logging enabled successfully.")
	fmt.Println(string(dataJSON))
}
