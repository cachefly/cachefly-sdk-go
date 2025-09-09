package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/cachefly/cachefly-sdk-go/pkg/cachefly"
	api "github.com/cachefly/cachefly-sdk-go/pkg/cachefly/api/v2_6"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	token := os.Getenv("CACHEFLY_API_TOKEN")
	if token == "" {
		log.Fatal("CACHEFLY_API_TOKEN is required")
	}

	client := cachefly.NewClient(cachefly.WithToken(token))

	req := api.CreateCacheWarmingTaskRequest{
		Name:    "Warm marko-service.cachefly.net",
		Targets: []string{"https://marko-service.cachefly.net/", "https://marko-service.cachefly.net/assets/logo.png"},
		Regions: []string{"global"},
	}

	task, err := client.CacheWarming.Create(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to create cache warming task: %v", err)
	}

	b, _ := json.MarshalIndent(task, "", "  ")
	fmt.Println(string(b))
}
