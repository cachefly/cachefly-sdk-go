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
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("⚠️ Warning: unable to load .env file: %v", err)
	}

	// Read API token
	token := os.Getenv("CACHEFLY_API_TOKEN")
	if token == "" {
		log.Fatal("❌ CACHEFLY_API_TOKEN environment variable is required")
	}

	// Read Script Config ID and file path arguments
	if len(os.Args) < 3 {
		log.Fatalf("⚠️ Usage: go run main.go <config_id> <script_file_path>")
	}
	configID := os.Args[1]
	filePath := os.Args[2]

	// Read the script content from file
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("❌ Failed to read script file %s: %v", filePath, err)
	}

	// Initialize CacheFly client
	client := cachefly.NewClient(
		cachefly.WithToken(token),
	)

	// Update the script config value from file
	updatedConfig, err := client.ScriptConfigs.UpdateValueAsFile(context.Background(), configID, content)
	if err != nil {
		log.Fatalf("❌ Failed to update script config %s from file: %v", configID, err)
	}

	// Pretty-print the updated configuration
	out, err := json.MarshalIndent(updatedConfig, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting updated script config JSON: %v", err)
	}

	fmt.Println("\n✅ Script configuration value updated from file successfully:")
	fmt.Println(string(out))
}
