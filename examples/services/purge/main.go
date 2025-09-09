// Example demonstrates purging cached content for a CacheFly service.
//
// Usage:
//
//	export CACHEFLY_API_TOKEN="your-token"
//	go run main.go <service_id> [all]
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/cachefly/cachefly-sdk-go/pkg/cachefly"
	api "github.com/cachefly/cachefly-sdk-go/pkg/cachefly/api/v2_6"
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
		log.Println("⚠️ Usage: go run main.go <service_id> [all]")
		return
	}
	serviceID := os.Args[1]
	purgeAll := len(os.Args) >= 3 && os.Args[2] == "all"

	client := cachefly.NewClient(cachefly.WithToken(token))
	ctx := context.Background()

	var req api.PurgeRequest
	if purgeAll {
		req = api.PurgeRequest{All: true}
	} else {
		// Example of purging specific objects/directories
		req = api.PurgeRequest{Paths: []string{"/index.html", "/images/"}}
	}

	fmt.Println("🔄 Submitting purge request...")
	if err := client.Services.Purge(ctx, serviceID, req); err != nil {
		log.Fatalf("❌ Failed to purge service cache: %v", err)
	}

	if purgeAll {
		fmt.Println("✅ Purge-all request accepted.")
	} else {
		fmt.Println("✅ Purge-paths request accepted.")
	}
}
