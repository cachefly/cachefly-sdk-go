package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/avvvet/cachefly-sdk-go/pkg/cachefly"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file (optional)
	if err := godotenv.Load(); err != nil {
		log.Printf("⚠️ Warning: unable to load .env file: %v", err)
	}

	// Read API token
	token := os.Getenv("CACHEFLY_API_TOKEN")
	if token == "" {
		log.Fatal("❌ CACHEFLY_API_TOKEN environment variable is required")
	}

	// Read Script Config ID argument
	if len(os.Args) < 2 {
		log.Fatalf("⚠️ Usage: go run main.go <config_id>")
	}
	configID := os.Args[1]

	// Initialize CacheFly client
	client := cachefly.NewClient(cachefly.WithToken(token))

	// Fetch the script config file content
	data, err := client.ScriptConfigs.GetValueAsFile(context.Background(), configID)
	if err != nil {
		log.Fatalf("❌ Failed to fetch script config file for %s: %v", configID, err)
	}

	// Write the content to a local file
	fileName := fmt.Sprintf("%s.js", configID)
	if err := ioutil.WriteFile(fileName, data, 0644); err != nil {
		log.Fatalf("❌ Failed to write script to %s: %v", fileName, err)
	}

	fmt.Printf("✅ Script configuration file saved to %s\n", fileName)
}
