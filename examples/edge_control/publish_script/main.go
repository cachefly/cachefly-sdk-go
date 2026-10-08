// Example demonstrates publishing and activating an edge control script.
//
// This example shows:
// - Publishing a script as a new, immutable version
// - Activating that version so it handles live traffic
// - Listing all published versions
//
// Usage:
//
//	export CACHEFLY_API_TOKEN="your-token"
//	go run main.go <service_id> [request|response]
//
// Example:
//
//	go run main.go srv_123456789 request
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

// Every edge control script must define a handler function.
const script = `function handler(event) {
  return event;
}`

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: unable to load .env file: %v", err)
	}

	token := os.Getenv("CACHEFLY_API_TOKEN")
	if token == "" {
		log.Fatal("CACHEFLY_API_TOKEN environment variable is required")
	}

	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <service_id> [request|response]")
	}
	serviceID := os.Args[1]
	kind := api.EdgeControlScriptKindRequest
	if len(os.Args) > 2 {
		kind = api.EdgeControlScriptKind(os.Args[2])
	}

	client := cachefly.NewClient(
		cachefly.WithToken(token),
	)
	ctx := context.Background()

	version, err := client.EdgeControlScripts.CreateVersion(ctx, serviceID, kind, script)
	if err != nil {
		log.Fatalf("Failed to publish edge control script: %v", err)
	}
	fmt.Printf("Published version %d\n", version.Version)

	active, err := client.EdgeControlScripts.ActivateVersion(ctx, serviceID, kind, version.Version)
	if err != nil {
		log.Fatalf("Failed to activate version %d: %v", version.Version, err)
	}

	out, err := json.MarshalIndent(active, "", "  ")
	if err != nil {
		log.Fatalf("Error formatting edge control script JSON: %v", err)
	}
	fmt.Println("\nActive edge control script:")
	fmt.Println(string(out))

	versions, err := client.EdgeControlScripts.ListVersions(ctx, serviceID, kind)
	if err != nil {
		log.Fatalf("Failed to list versions: %v", err)
	}
	fmt.Println("\nPublished versions:")
	for _, v := range versions {
		fmt.Printf("  version %d: %s (%d bytes)\n", v.Version, v.Status, v.ScriptSize)
	}
}
