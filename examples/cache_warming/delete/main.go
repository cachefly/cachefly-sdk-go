package main

import (
	"context"
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

	if err := client.CacheWarming.DeleteByID(context.Background(), id); err != nil {
		log.Fatalf("failed to delete task: %v", err)
	}
	fmt.Println("deleted")
}
