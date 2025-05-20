// Entry point: HTTP client setup, auth, base URL
package cachefly

import (
	"github.com/avvvet/cachefly-sdk-go/internal/httpclient"
	"github.com/avvvet/cachefly-sdk-go/pkg/cachefly/api"
)

// main struct to interact with CacheFly APIs.
type Client struct {
	httpClient *httpclient.Client

	// API groups
	Services                   *api.ServicesService
	Accounts                   *api.AccountsService
	ServiceDomains             *api.ServiceDomainsService
	ServiceRules               *api.ServiceRulesService
	ServiceOptions             *api.ServiceOptionsService
	ServiceOptionsRefererRules *api.ServiceOptionsRefererRulesService
	ServiceImageOptimization   *api.ServiceImageOptimizationService
	Certificates               *api.CertificatesService
	Origins                    *api.OriginsService
	Users                      *api.UsersService
	ScriptConfigs              *api.ScriptConfigsService
	TLSProfiles                *api.TLSProfilesService
}

// for configuring the Client.
type Option func(*ClientConfig)

type ClientConfig struct {
	Token   string
	BaseURL string
}

// sets the Bearer token.
func WithToken(token string) Option {
	return func(c *ClientConfig) {
		c.Token = token
	}
}

// WithBaseURL overrides the API base URL.
func WithBaseURL(url string) Option {
	return func(c *ClientConfig) {
		c.BaseURL = url
	}
}

// initializes and returns a CacheFly API client.
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
	}
}
