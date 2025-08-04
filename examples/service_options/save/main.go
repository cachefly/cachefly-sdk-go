// Example demonstrates updating service options for a CacheFly service with validation.
//
// Usage:
//
// export CACHEFLY_API_TOKEN="your-token"
// go run main.go <service_id>
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/cachefly/cachefly-go-sdk/pkg/cachefly"
	api "github.com/cachefly/cachefly-go-sdk/pkg/cachefly/api/v2_5"
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
		log.Fatalf("⚠️ Usage: go run main.go <service_id>")
	}
	serviceID := os.Args[1]

	client := cachefly.NewClient(cachefly.WithToken(token))
	ctx := context.Background()

	// First, let's see what the current service options look like
	fmt.Printf("🔍 Getting current service options to understand the format...\n")
	currentOptions, err := client.ServiceOptions.GetOptions(ctx, serviceID)
	if err != nil {
		log.Fatalf("❌ Failed to get current options: %v", err)
	}

	// Show current format
	currentJSON, _ := json.MarshalIndent(currentOptions, "", "  ")
	fmt.Printf("📄 Current service options format:\n%s\n\n", string(currentJSON))

	enableAllOptions := api.ServiceOptions{

		"allowretry":       true,
		"forceorigqstring": true,
		"send-xff":         true,
		"brotli_support":   true,

		// Purge options
		"purgenoquery": true,
		// uncomment to test unsupported feature error handling
		/*
			"unsupported_feature": map[string]interface{}{
				"enabled": true,
				"value":   "test",
			},
		*/

		"reverseProxy": map[string]interface{}{
			"enabled":           true,
			"mode":              "WEB",
			"hostname":          "www.example.com",
			"cacheByQueryParam": true,
			"originScheme":      "FOLLOW",
			"ttl":               2678400,
			"useRobotsTxt":      true,
		},

		// Standard enabled/value structure options
		"error_ttl": map[string]interface{}{
			"enabled": true,
			"value":   700,
		},

		"ttfb_timeout": map[string]interface{}{
			"enabled": true,
			"value":   30,
		},

		"contimeout": map[string]interface{}{
			"enabled": true,
			"value":   10,
		},

		"maxcons": map[string]interface{}{
			"enabled": true,
			"value":   100,
		},

		"sharedshield": map[string]interface{}{
			"enabled": true,
			"value":   "ORD", // Chicago data center
		},

		"purgemode": map[string]interface{}{
			"enabled": true,
			"value":   "2",
		},

		/*
			"slice": map[string]interface{}{
				"enabled": true,
				"value":   true,
			},
		*/

		"originhostheader": map[string]interface{}{
			"enabled": true,
			"value":   []string{"origin.example.com", "backup.example.com"},
		},

		/*
			"bwthrottlequery": map[string]interface{}{
				"enabled": true,
				"value":   []string{"limit", "throttle"},
			},
		*/
		"dirpurgeskip": map[string]interface{}{
			"enabled": true,
			"value":   1,
		},

		// Caching options
		"nocache":              true,
		"cachebygeocountry":    true,
		"cachebyregion":        true,
		"normalizequerystring": true,
		"servestale":           true,
		"cachebyreferer":       true,
		"expiryHeaders":        []interface{}{},

		// Delivery options
		"cors":          true,
		"autoRedirect":  true,
		"livestreaming": true,
		"linkpreheat":   true,
		"redirect": map[string]interface{}{
			"enabled": true,
			"value":   "https://www.newdomain.com/",
		},
		"skip_encoding_ext": map[string]interface{}{
			"enabled": true,
			"value":   []string{".zip", ".gz", ".tar", ".rar"},
		},
		"bwthrottle": map[string]interface{}{
			"enabled": true,
			"value":   70656,
		},
		"httpmethods": map[string]interface{}{
			"enabled": true,
			"value": map[string]interface{}{
				"GET":     true,
				"POST":    true,
				"PUT":     false,
				"DELETE":  false,
				"HEAD":    true,
				"OPTIONS": true,
				"PATCH":   false,
			},
		},

		// Security options
		"protectServeKeyEnabled": true,
		"skip_pserve_ext": map[string]interface{}{
			"enabled": true,
			"value":   []string{".jpg", ".png", ".gif", ".css", ".js"},
		},
	}

	opts := enableAllOptions

	fmt.Printf("🔄 Updating service options...\n")

	updated, err := client.ServiceOptions.UpdateOptions(ctx, serviceID, opts)
	if err != nil {
		if validationErr, ok := err.(api.ServiceOptionsValidationError); ok {
			fmt.Printf("❌ Validation failed: %s\n", validationErr.Message)
			for _, fieldErr := range validationErr.Errors {
				fmt.Printf("   • %s: %s (%s)\n", fieldErr.Field, fieldErr.Message, fieldErr.Code)
			}
			log.Fatalf("❌ Please fix the validation errors above")
		}
		log.Fatalf("❌ Failed to update service options for %s: %v", serviceID, err)
	}

	out, _ := json.MarshalIndent(updated, "", "  ")
	fmt.Println("✅ Service options updated successfully:")
	fmt.Println(string(out))
}
