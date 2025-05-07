package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/avvvet/cachefly-sdk-go/pkg/cachefly"
	"github.com/avvvet/cachefly-sdk-go/pkg/cachefly/api"
)

func main() {
	client := cachefly.NewClient(
		cachefly.WithToken("YOUR-API-TOKEN"),
	)

	ctx := context.Background()

	resp, err := client.Services.List(ctx, api.ListOptions{
		ResponseType:    "shallow",
		IncludeFeatures: false,
		Status:          "ACTIVE",
		Offset:          0,
		Limit:           1,
	})
	if err != nil {
		log.Fatalf("Error fetching services: %v", err)
	}

	data, err := json.MarshalIndent(resp.Services, "", "  ")
	if err != nil {
		fmt.Printf("Error formatting JSON: %v\n", err)
		return
	}

	fmt.Println("Services:")
	fmt.Println(string(data))
}
