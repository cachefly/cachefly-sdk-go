package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/avvvet/cachefly-sdk-go/pkg/cachefly"
	"github.com/avvvet/cachefly-sdk-go/pkg/cachefly/api"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func LoadEnv() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Info(".env file loaded.")
}

func main() {

	LoadEnv()
	client := cachefly.NewClient(
		cachefly.WithToken(os.Getenv("API_TOKEN")),
	)

	ctx := context.Background()

	account, err := client.Accounts.Get(ctx, "")
	if err != nil {
		log.Fatalf("Error fetching account: %v", err)
	}

	j, err := json.MarshalIndent(account, "", " ")
	if err != nil {
		log.Fatalf("Error formating account %v", err)
	}
	fmt.Print("\n Accounts")
	fmt.Println(string(j))

	// --- Create a new service  ---
	id := uuid.New().String()[:8]
	name := "my-test-service-" + id

	newService := api.CreateServiceRequest{
		Name:        name,
		UniqueName:  name,
		Description: "This is a test service created from SDK",
	}

	created, err := client.Services.Create(ctx, newService)
	if err != nil {
		log.Fatalf("Error creating service: %v", err)
	}

	createdJSON, err := json.MarshalIndent(created, "", "  ")
	if err != nil {
		log.Fatalf("Error formatting created service: %v", err)
	}

	fmt.Println("Created Service:")
	fmt.Println(string(createdJSON))

	// --- Get service by Id ---
	serviceID := created.ID

	service, err := client.Services.Get(ctx, serviceID, "shallow", false)
	if err != nil {
		log.Fatalf("Error getting service: %v", err)
	}

	data, _ := json.MarshalIndent(service, "", "  ")
	fmt.Println("Fetched Service:")
	fmt.Println(string(data))

	// --- List active services ---
	resp, err := client.Services.List(ctx, api.ListOptions{
		ResponseType:    "shallow",
		IncludeFeatures: false,
		Status:          "ACTIVE",
		Offset:          0,
		Limit:           5,
	})
	if err != nil {
		log.Fatalf("Error fetching services: %v", err)
	}

	listJSON, err := json.MarshalIndent(resp.Services, "", "  ")
	if err != nil {
		log.Fatalf("Error formatting service list: %v", err)
	}

	fmt.Println("\nActive Services:")
	fmt.Println(string(listJSON))
}
