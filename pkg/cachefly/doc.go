// Package cachefly provides a client library for the CacheFly CDN API.
//
// The CacheFly SDK allows you to programmatically manage your CDN resources
// including services, domains, SSL certificates, and caching rules.
//
// # Getting Started
//
// To use this SDK, you'll need a CacheFly API token. You can obtain one from
// your CacheFly dashboard.
//
// Create a new client:
//
//	client := cachefly.NewClient(
//	    cachefly.WithToken("your-api-token"),
//	)
//
// # Making API Calls
//
// The client provides access to different resource services:
//
//	// List all services
//	services, err := client.Services.List(context.Background())
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// Get a specific service
//	service, err := client.Services.GetByID(context.Background(), "srv_123")
//
//	// Manage accounts
//	account, err := client.Accounts.Get(context.Background(), "acc_456")
//
// # Available Services
//
// The client provides access to the following services:
//
//   - Services: CDN service configuration and management
//   - Accounts: Account management and billing
//   - ServiceDomains: Domain configuration for services
//   - ServiceRules: Caching and delivery rules
//   - Certificates: SSL/TLS certificate management
//   - Origins: Origin server configuration
//   - Users: User access and permissions
//
// # Error Handling
//
// All methods return an error as the last value. Check these errors to handle
// API failures gracefully:
//
//	service, err := client.Services.GetByID(ctx, "invalid-id")
//	if err != nil {
//	    // Handle error - could be network, auth, or API error
//	    log.Printf("Failed to get service: %v", err)
//	}
//
// # Configuration Options
//
// The client supports several configuration options:
//
//	client := cachefly.NewClient(
//	    cachefly.WithToken("your-token"),           // API authentication
//	    cachefly.WithBaseURL("https://api.example"), // Custom API endpoint
//	)
//
// # Examples
//
// See the examples directory for complete working examples of common use cases.
package cachefly
