package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/cachefly/cachefly-go-sdk/pkg/cachefly"
	api "github.com/cachefly/cachefly-go-sdk/pkg/cachefly/api/v2_5"
	"github.com/google/uuid"
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

	client := cachefly.NewClient(
		cachefly.WithToken(token),
	)

	id := uuid.New().String()[:8]
	name := "sdk-test-service-" + id

	newService := api.CreateServiceRequest{
		Name:        name,
		UniqueName:  name,
		Description: "This is a test service created from SDK",
	}

	ctx := context.Background()
	resp, err := client.Services.Create(ctx, newService)
	if err != nil {
		log.Fatalf("❌ Error creating service: %v", err)
	}

	dataJSON, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatalf("Error formatting [service]: %v", err)
	}

	fmt.Println("\n ✅ Created Service:")
	fmt.Println(string(dataJSON))

}
