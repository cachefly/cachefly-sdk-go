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

	client := cachefly.NewClient(
		cachefly.WithToken(token),
	)

	payload := api.UpdateServiceRequest{
		Description:       "updated service from SDK",
		TLSProfile:        "66320d4208158b00411703e4",
		AutoSSL:           false,
		DeliveryRegion:    "673f01735a5ddf015fc46997",
		ConfigurationMode: "API_RULES",
	}

	service, err := client.Services.UpdateServiceByID(context.Background(), serviceID, payload)
	if err != nil {
		log.Fatalf("❌ Failed to update service by ID: %v", err)
	}

	out, err := json.MarshalIndent(service, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting service JSON: %v", err)
	}

	fmt.Println("\n✅ Service updated successfully:")
	fmt.Println(string(out))
}
