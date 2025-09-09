package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/cachefly/cachefly-sdk-go/pkg/cachefly"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	token := os.Getenv("CACHEFLY_API_TOKEN")
	if token == "" {
		log.Fatal("CACHEFLY_API_TOKEN is required")
	}
	if len(os.Args) < 2 {
		log.Fatalf("usage: go run main.go <task_id>")
	}

	id := os.Args[1]
	client := cachefly.NewClient(cachefly.WithToken(token))

	task, err := client.CacheWarming.GetByID(context.Background(), id)
	if err != nil {
		log.Fatalf("failed to fetch task: %v", err)
	}

	b, _ := json.MarshalIndent(task, "", "  ")
	fmt.Println(string(b))
}
