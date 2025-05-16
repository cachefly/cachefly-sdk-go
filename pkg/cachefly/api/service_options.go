package api

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/avvvet/cachefly-sdk-go/internal/httpclient"
)

// ServiceOptions represents basic service options.
type ServiceOptions struct {
	CORS                   bool               `json:"cors"`
	AutoRedirect           bool               `json:"autoRedirect"`
	ReverseProxy           ReverseProxyConfig `json:"reverseProxy"`
	MimeTypesOverrides     []MimeTypeOverride `json:"mimeTypesOverrides"`
	ExpiryHeaders          []ExpiryHeader     `json:"expiryHeaders"`
	APIKeyEnabled          bool               `json:"apiKeyEnabled"`
	ProtectServeKeyEnabled bool               `json:"protectServeKeyEnabled"`
}

// ReverseProxyConfig configures reverse proxy behavior.
type ReverseProxyConfig struct {
	Enabled           bool   `json:"enabled"`
	Hostname          string `json:"hostname,omitempty"`
	Prepend           string `json:"prepend,omitempty"`
	TTL               int    `json:"ttl,omitempty"`
	CacheByQueryParam bool   `json:"cacheByQueryParam,omitempty"`
	OriginScheme      string `json:"originScheme,omitempty"`
	UseRobotsTXT      bool   `json:"useRobotsTxt,omitempty"`
	Mode              string `json:"mode,omitempty"`
}

// MimeTypeOverride overrides MIME types.
type MimeTypeOverride struct {
	Extension string `json:"extension"`
	MimeType  string `json:"mimeType"`
}

// ExpiryHeader sets expiry headers for paths.
type ExpiryHeader struct {
	Path       string `json:"path"`
	Extension  string `json:"extension"`
	ExpiryTime int    `json:"expiryTime"`
}

// LegacyAPIKeyResponse represents API key payload.
type LegacyAPIKeyResponse struct {
	APIKey string `json:"apiKey"`
}

// ProtectServeKeyResponse for protectserve.
type ProtectServeKeyResponse struct {
	ProtectServeKey   string `json:"protectServeKey"`
	ForceProtectServe string `json:"forceProtectserve"`
}

// UpdateProtectServeRequest updates protectserve options.
type UpdateProtectServeRequest struct {
	ForceProtectServe string `json:"forceProtectServe"`
	ProtectServeKey   string `json:"protectServeKey"`
}

// FTPSettingsResponse represents FTP settings.
type FTPSettingsResponse struct {
	FTPPassword string `json:"ftpPassword"`
}

// ServiceOptionsService handles service options endpoints.
type ServiceOptionsService struct {
	Client *httpclient.Client
}

// GetBasicOptions retrieves basic options for a service.
func (s *ServiceOptionsService) GetBasicOptions(ctx context.Context, id string) (*ServiceOptions, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/services/%s/options", id)

	var opts ServiceOptions
	if err := s.Client.Get(ctx, endpoint, &opts); err != nil {
		return nil, err
	}
	return &opts, nil
}

// SaveBasicOptions updates basic service options.
func (s *ServiceOptionsService) SaveBasicOptions(ctx context.Context, id string, req ServiceOptions) (*ServiceOptions, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/services/%s/options", id)

	var updated ServiceOptions
	if err := s.Client.Put(ctx, endpoint, req, &updated); err != nil {
		return nil, err
	}
	return &updated, nil
}

// GetLegacyAPIKey returns the legacy API key for a service.
func (s *ServiceOptionsService) GetLegacyAPIKey(ctx context.Context, id string) (*LegacyAPIKeyResponse, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/services/%s/options/apikey", id)

	var res LegacyAPIKeyResponse
	if err := s.Client.Get(ctx, endpoint, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// RegenerateLegacyAPIKey creates a new legacy API key.
func (s *ServiceOptionsService) RegenerateLegacyAPIKey(ctx context.Context, id string) (*LegacyAPIKeyResponse, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/services/%s/options/apikey", id)

	var res LegacyAPIKeyResponse
	if err := s.Client.Post(ctx, endpoint, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// DeleteLegacyAPIKey deletes the legacy API key for a service.
func (s *ServiceOptionsService) DeleteLegacyAPIKey(ctx context.Context, id string) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/services/%s/options/apikey", id)
	return s.Client.Delete(ctx, endpoint, nil)
}

// GetProtectServeKey retrieves the protectserve key (optional hideSecrets).
func (s *ServiceOptionsService) GetProtectServeKey(ctx context.Context, id string, hideSecrets bool) (*ProtectServeKeyResponse, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/services/%s/options/protectserve", id)
	params := url.Values{}
	params.Set("hideSecrets", strconv.FormatBool(hideSecrets))
	fullURL := fmt.Sprintf("%s?%s", endpoint, params.Encode())

	var res ProtectServeKeyResponse
	if err := s.Client.Get(ctx, fullURL, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// RecreateProtectServeKey regenerates or reverts the protectserve key.
func (s *ServiceOptionsService) RecreateProtectServeKey(ctx context.Context, id, action string) (*ProtectServeKeyResponse, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/services/%s/options/protectserve", id)
	params := url.Values{}
	if action != "" {
		params.Set("action", action)
	}
	fullURL := fmt.Sprintf("%s?%s", endpoint, params.Encode())

	var res ProtectServeKeyResponse
	if err := s.Client.Post(ctx, fullURL, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// UpdateProtectServeOptions updates protectserve key and options.
func (s *ServiceOptionsService) UpdateProtectServeOptions(ctx context.Context, id string, req UpdateProtectServeRequest) (*ProtectServeKeyResponse, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/services/%s/options/protectserve", id)

	var res ProtectServeKeyResponse
	if err := s.Client.Put(ctx, endpoint, req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetFTPSettings retrieves FTP settings for a service (optional hideSecrets).
func (s *ServiceOptionsService) GetFTPSettings(ctx context.Context, id string, hideSecrets bool) (*FTPSettingsResponse, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/services/%s/options/ftp", id)
	params := url.Values{}
	params.Set("hideSecrets", strconv.FormatBool(hideSecrets))
	fullURL := fmt.Sprintf("%s?%s", endpoint, params.Encode())

	var res FTPSettingsResponse
	if err := s.Client.Get(ctx, fullURL, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// RegenerateFTPPassword regenerates the FTP password for a service.
func (s *ServiceOptionsService) RegenerateFTPPassword(ctx context.Context, id string, hideSecrets bool) (*FTPSettingsResponse, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/services/%s/options/ftp", id)
	params := url.Values{}
	params.Set("hideSecrets", strconv.FormatBool(hideSecrets))
	fullURL := fmt.Sprintf("%s?%s", endpoint, params.Encode())

	var res FTPSettingsResponse
	if err := s.Client.Post(ctx, fullURL, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
