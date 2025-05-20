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
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("⚠️ Warning: unable to load .env file: %v", err)
	}

	// Read API token
	token := os.Getenv("CACHEFLY_API_TOKEN")
	if token == "" {
		log.Fatal("❌ CACHEFLY_API_TOKEN environment variable is required")
	}

	// Initialize CacheFly client
	client := cachefly.NewClient(
		cachefly.WithToken(token),
	)

	// Prepare payload for creating a certificate
	opts := api.CreateCertificateRequest{
		Certificate:    os.Getenv("CERT_PEM"),      // Set in .env or replace with literal
		CertificateKey: os.Getenv("CERT_KEY"),      // Set in .env or replace with literal
		Password:       os.Getenv("CERT_PASSWORD"), // Optional
	}

	// Call Create (POST /certificates)
	cert, err := client.Certificates.Create(context.Background(), opts)
	if err != nil {
		log.Fatalf("❌ Failed to create certificate: %v", err)
	}

	// Pretty-print the created certificate details
	out, err := json.MarshalIndent(cert, "", "  ")
	if err != nil {
		log.Fatalf("❌ Error formatting certificate JSON: %v", err)
	}

	fmt.Println("\n✅ Certificate created successfully:")
	fmt.Println(string(out))
}
