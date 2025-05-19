package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/avvvet/cachefly-sdk-go/pkg/cachefly"
	"github.com/avvvet/cachefly-sdk-go/pkg/cachefly/api"
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

	payload := api.EnableOriginLogsRequest{
		LogTarget: "stringstringstringstring",
	}

	updatedService, err := client.Services.EnableOriginLogging(context.Background(), serviceID, payload)
	if err != nil {
		log.Fatalf("❌ Failed to enable origin logging: %v", err)
	}

	out, err := json.MarshalIndent(updatedService, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting response JSON [EnableOriginLogging]: %v", err)
	}

	fmt.Println("\n✅ Origin logging enabled successfully:")
	fmt.Println(string(out))
}
