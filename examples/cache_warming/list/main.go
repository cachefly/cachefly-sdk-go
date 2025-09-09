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

	client := cachefly.NewClient(cachefly.WithToken(token))

	resp, err := client.CacheWarming.List(context.Background(), struct {
		Offset, Limit int
		ResponseType  string
	}{Offset: 0, Limit: 10})
	if err != nil {
		log.Fatalf("failed to list cache warming tasks: %v", err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
