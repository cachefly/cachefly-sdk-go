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
	if len(os.Args) < 2 {
		log.Fatal("usage: go run main.go <service_id>")
	}
	sid := os.Args[1]

	client := cachefly.NewClient(cachefly.WithToken(token))

	opts := api.StatsQueryOptions{
		Limit:   10,
		From:    "2024-09-01",
		To:      "2025-09-02",
		GroupBy: []string{"pop", "date"},
	}
	resp, err := client.ServiceStats.POP(context.Background(), sid, opts)
	if err != nil {
		log.Fatalf("failed to fetch service POP stats: %v", err)
	}
	out, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatalf("failed to format JSON: %v", err)
	}
	fmt.Println(string(out))
}
