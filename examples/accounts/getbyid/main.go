package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/avvvet/cachefly-sdk-go/pkg/cachefly"
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
		log.Println("⚠️ Usage: go run main.go <account_id> [responseType]")
		return
	}
	accountID := os.Args[1]
	responseType := ""
	if len(os.Args) > 2 {
		responseType = os.Args[2]
	}

	client := cachefly.NewClient(
		cachefly.WithToken(token),
	)

	account, err := client.Accounts.GetByID(context.Background(), accountID, responseType)
	if err != nil {
		log.Fatalf("❌ Failed to get account by ID: %v", err)
	}

	resp, err := json.MarshalIndent(account, "", " ")
	if err != nil {
		log.Fatalf("❌ Error formatting account: %v", err)
	}

	fmt.Println("\n ✅ Account By Id fetched successfully:")
	fmt.Println(string(resp))

}
