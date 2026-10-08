// Example demonstrates working with the edge control script library.
//
// This example shows:
// - Creating a USER script in the library
// - Listing USER and SYSTEM request scripts
//
// Usage:
//
//	export CACHEFLY_API_TOKEN="your-token"
//	go run main.go
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
		log.Printf("Warning: unable to load .env file: %v", err)
	}

	token := os.Getenv("CACHEFLY_API_TOKEN")
	if token == "" {
		log.Fatal("CACHEFLY_API_TOKEN environment variable is required")
	}

	client := cachefly.NewClient(
		cachefly.WithToken(token),
	)
	ctx := context.Background()

	created, err := client.EdgeControlLibrary.Create(ctx, api.EdgeControlLibraryScriptRequest{
		Name: "pass-through",
		Kind: api.EdgeControlScriptKindRequest,
		Code: "function handler(event) {\n  return event;\n}",
	})
	if err != nil {
		log.Fatalf("Failed to create library script: %v", err)
	}
	fmt.Printf("Created library script %s (%s)\n", created.ID, created.Name)

	list, err := client.EdgeControlLibrary.List(ctx, api.ListEdgeControlLibraryOptions{
		Kind:  api.EdgeControlScriptKindRequest,
		Limit: 50,
	})
	if err != nil {
		log.Fatalf("Failed to list library scripts: %v", err)
	}

	fmt.Printf("\n%d request scripts in the library:\n", list.Meta.Count)
	for _, script := range list.Scripts {
		fmt.Printf("  %s  %-6s  %s\n", script.ID, script.Type, script.Name)
	}
}
