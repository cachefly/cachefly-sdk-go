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

	opts := api.StatsQueryOptions{
		From:    "2024-09-01",
		To:      "2025-09-02",
		Limit:   10,
		GroupBy: []string{"pop", "date"},
	}
	resp, err := client.AccountStats.POP(context.Background(), opts)
	if err != nil {
		log.Fatalf("failed to fetch account POP stats: %v", err)
	}

	out, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatalf("failed to format JSON: %v", err)
	}
	fmt.Println(string(out))
}
