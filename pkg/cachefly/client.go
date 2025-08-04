// Package cachefly provides the main client interface for CacheFly SDK.
//
// This package serves as the entry point for CacheFly API interactions,
// providing client initialization and access to all API service groups.
//
// The client is organized into service groups that correspond to different
// areas of the CacheFly API, such as Services, Accounts, Certificates, etc.
//
// Basic Usage:
//
//	client := cachefly.NewClient(
//		cachefly.WithToken("your-bearer-token"),
//	)
//
//	// Access different service groups
//	services, err := client.Services.List()
//	accounts, err := client.Accounts.Get("account-id")
//
// Custom Configuration:
//
//	client := cachefly.NewClient(
//		cachefly.WithToken("your-bearer-token"),
//		cachefly.WithBaseURL("https://custom-api.cachefly.com/api/2.5"),
//	)
package cachefly

import (
	"github.com/cachefly/cachefly-go-sdk/internal/httpclient"
	api "github.com/cachefly/cachefly-go-sdk/pkg/cachefly/api/v2_5"
)

// Client is the main struct to interact with CacheFly APIs.
//
// It provides access to all CacheFly API service groups through
// organized service properties. Each service group handles
// specific aspects of the CacheFly platform.
type Client struct {
	httpClient *httpclient.Client

	// API service groups

	// Services manages CacheFly services (CDN configurations)
	Services *api.ServicesService

	// Accounts manages account-level operations and settings
	Accounts *api.AccountsService

	// ServiceDomains manages domain configurations for services
	ServiceDomains *api.ServiceDomainsService

	// ServiceRules manages caching and delivery rules
	ServiceRules *api.ServiceRulesService

	// ServiceOptions manages service-level configuration options
	ServiceOptions *api.ServiceOptionsService

	// ServiceOptionsRefererRules manages referer-based access rules
	ServiceOptionsRefererRules *api.ServiceOptionsRefererRulesService

	// ServiceImageOptimization manages image optimization settings
	ServiceImageOptimization *api.ServiceImageOptimizationService

	// Certificates manages SSL/TLS certificates
	Certificates *api.CertificatesService

	// Origins manages origin server configurations
	Origins *api.OriginsService

	// Users manages user accounts and permissions
	Users *api.UsersService

	// ScriptConfigs manages edge script configurations
	ScriptConfigs *api.ScriptConfigsService

	// TLSProfiles manages TLS security profiles
	TLSProfiles *api.TLSProfilesService

	// DeliveryRegions manages delivery regions
	DeliveryRegions *api.DeliveryRegionsService
}

// Option is a functional option for configuring the Client.
type Option func(*ClientConfig)

// ClientConfig holds configuration options for the CacheFly client.
type ClientConfig struct {
	// Token is the Bearer token for API authentication
	Token string

	// BaseURL overrides the default API base URL
	BaseURL string
}

// WithToken sets the Bearer token for API authentication.
//
// This token is required for all API calls and should be obtained
// from your CacheFly dashboard.
//
// Example:
//
//	client := cachefly.NewClient(cachefly.WithToken("your-bearer-token"))
func WithToken(token string) Option {
	return func(c *ClientConfig) {
		c.Token = token
	}
}

// WithBaseURL overrides the default API base URL.
//
// This is useful for testing against different environments
// or when using a custom API endpoint.
//
// Example:
//
//	client := cachefly.NewClient(
//		cachefly.WithToken("token"),
//		cachefly.WithBaseURL("https://staging-api.cachefly.com/api/2.5"),
//	)
func WithBaseURL(url string) Option {
	return func(c *ClientConfig) {
		c.BaseURL = url
	}
}

// NewClient initializes and returns a new CacheFly API client.
//
// The client is configured with functional options and provides
// access to all CacheFly API service groups.
//
// Authentication is required - use WithToken() to provide your API token.
//
// Example:
//
//	// Basic client with authentication
//	client := cachefly.NewClient(cachefly.WithToken("your-token"))
//
//	// Client with custom configuration
//	client := cachefly.NewClient(
//		cachefly.WithToken("your-token"),
//		cachefly.WithBaseURL("https://api.cachefly.com/api/2.5"),
//	)
//
//	// Use the client
//	services, err := client.Services.List()
//	if err != nil {
//		log.Fatal(err)
//	}
func NewClient(opts ...Option) *Client {
	cfg := &ClientConfig{
		BaseURL: "https://api.cachefly.com/api/2.5",
	}

	for _, opt := range opts {
		opt(cfg)
	}

	hc := httpclient.New(httpclient.Config{
		BaseURL:   cfg.BaseURL,
		AuthToken: cfg.Token,
	})

	return &Client{
		httpClient:                 hc,
		Services:                   &api.ServicesService{Client: hc},
		Accounts:                   &api.AccountsService{Client: hc},
		ServiceDomains:             &api.ServiceDomainsService{Client: hc},
		ServiceRules:               &api.ServiceRulesService{Client: hc},
		ServiceOptions:             &api.ServiceOptionsService{Client: hc},
		ServiceOptionsRefererRules: &api.ServiceOptionsRefererRulesService{Client: hc},
		ServiceImageOptimization:   &api.ServiceImageOptimizationService{Client: hc},
		Certificates:               &api.CertificatesService{Client: hc},
		Origins:                    &api.OriginsService{Client: hc},
		Users:                      &api.UsersService{Client: hc},
		ScriptConfigs:              &api.ScriptConfigsService{Client: hc},
		TLSProfiles:                &api.TLSProfilesService{Client: hc},
		DeliveryRegions:            &api.DeliveryRegionsService{Client: hc},
	}
}
