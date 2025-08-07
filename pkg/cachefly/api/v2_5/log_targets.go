package v2_5

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/cachefly/cachefly-go-sdk/internal/httpclient"
)

// LogTarget represents a CacheFly log target configuration.
type LogTarget struct {
	ID                         string   `json:"_id"`
	UpdatedAt                  string   `json:"updatedAt"`
	CreatedAt                  string   `json:"createdAt"`
	Name                       string   `json:"name"`
	Type                       string   `json:"type"`
	Endpoint                   string   `json:"endpoint,omitempty"`
	Region                     string   `json:"region,omitempty"`
	Bucket                     string   `json:"bucket,omitempty"`
	AccessKey                  string   `json:"accessKey,omitempty"`
	SecretKey                  string   `json:"secretKey,omitempty"`
	SignatureVersion           string   `json:"signatureVersion,omitempty"`
	JsonKey                    string   `json:"jsonKey,omitempty"`
	Hosts                      []string `json:"hosts,omitempty"`
	SSL                        bool     `json:"ssl,omitempty"`
	SSLCertificateVerification bool     `json:"sslCertificateVerification,omitempty"`
	Index                      string   `json:"index,omitempty"`
	User                       string   `json:"user,omitempty"`
	Password                   string   `json:"password,omitempty"`
	ApiKey                     string   `json:"apiKey,omitempty"`
	AccessLogsServices         []string `json:"accessLogsServices"`
	OriginLogsServices         []string `json:"originLogsServices"`
}

// ListLogTargetsResponse contains paginated log target results.
type ListLogTargetsResponse struct {
	Meta       MetaInfo    `json:"meta"`
	LogTargets []LogTarget `json:"data"`
}

// ListLogTargetsOptions allows filtering & pagination for log targets.
type ListLogTargetsOptions struct {
	Type         string
	Offset       int
	Limit        int
	ResponseType string
}

// CreateLogTargetRequest contains the required fields for creating a new log target.
type CreateLogTargetRequest struct {
	Name                       string   `json:"name"`
	Type                       string   `json:"type"`
	Endpoint                   string   `json:"endpoint,omitempty"`
	Region                     string   `json:"region,omitempty"`
	Bucket                     string   `json:"bucket,omitempty"`
	AccessKey                  string   `json:"accessKey,omitempty"`
	SecretKey                  string   `json:"secretKey,omitempty"`
	SignatureVersion           string   `json:"signatureVersion,omitempty"`
	JsonKey                    string   `json:"jsonKey,omitempty"`
	Hosts                      []string `json:"hosts,omitempty"`
	SSL                        bool     `json:"ssl,omitempty"`
	SSLCertificateVerification bool     `json:"sslCertificateVerification,omitempty"`
	Index                      string   `json:"index,omitempty"`
	User                       string   `json:"user,omitempty"`
	Password                   string   `json:"password,omitempty"`
	ApiKey                     string   `json:"apiKey,omitempty"`
	AccessLogsServices         []string `json:"accessLogsServices"`
	OriginLogsServices         []string `json:"originLogsServices"`
}

// UpdateLogTargetRequest contains the fields for updating an existing log target.
type UpdateLogTargetRequest struct {
	Name                       string   `json:"name,omitempty"`
	Type                       string   `json:"type,omitempty"`
	Endpoint                   string   `json:"endpoint,omitempty"`
	Region                     string   `json:"region,omitempty"`
	Bucket                     string   `json:"bucket,omitempty"`
	AccessKey                  string   `json:"accessKey,omitempty"`
	SecretKey                  string   `json:"secretKey,omitempty"`
	SignatureVersion           string   `json:"signatureVersion,omitempty"`
	JsonKey                    string   `json:"jsonKey,omitempty"`
	Hosts                      []string `json:"hosts,omitempty"`
	SSL                        bool     `json:"ssl,omitempty"`
	SSLCertificateVerification bool     `json:"sslCertificateVerification,omitempty"`
	Index                      string   `json:"index,omitempty"`
	User                       string   `json:"user,omitempty"`
	Password                   string   `json:"password,omitempty"`
	ApiKey                     string   `json:"apiKey,omitempty"`
	AccessLogsServices         []string `json:"accessLogsServices,omitempty"`
	OriginLogsServices         []string `json:"originLogsServices,omitempty"`
}

// SetLoggingRequest contains the services to set logging for.
type SetLoggingRequest struct {
	AccessLogsServices []string `json:"accessLogsServices"`
	OriginLogsServices []string `json:"originLogsServices"`
}

// LogTargetsService handles log target-related API operations.
type LogTargetsService struct {
	Client *httpclient.Client
}

// List returns all log targets for the current account.
func (s *LogTargetsService) List(ctx context.Context, opts ListLogTargetsOptions) (*ListLogTargetsResponse, error) {
	endpoint := "/logtargets"

	params := url.Values{}

	if opts.Type != "" {
		params.Set("type", opts.Type)
	}

	if opts.Offset >= 0 {
		params.Set("offset", strconv.Itoa(opts.Offset))
	}

	if opts.Limit > 0 {
		params.Set("limit", strconv.Itoa(opts.Limit))
	}

	if opts.ResponseType != "" {
		params.Set("responseType", opts.ResponseType)
	}

	fullURL := fmt.Sprintf("%s?%s", endpoint, params.Encode())
	var resp ListLogTargetsResponse
	if err := s.Client.Get(ctx, fullURL, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Create creates a new log target.
func (s *LogTargetsService) Create(ctx context.Context, req CreateLogTargetRequest) (*LogTarget, error) {
	endpoint := "/logtargets"

	var created LogTarget
	if err := s.Client.Post(ctx, endpoint, req, &created); err != nil {
		return nil, err
	}

	return &created, nil
}

// UpdateByID updates an existing log target by its ID.
func (s *LogTargetsService) UpdateByID(ctx context.Context, id string, req UpdateLogTargetRequest) (*LogTarget, error) {
	if id == "" {
		return nil, fmt.Errorf("log target ID is required")
	}
	endpoint := fmt.Sprintf("/logtargets/%s", id)

	var updated LogTarget
	if err := s.Client.Put(ctx, endpoint, req, &updated); err != nil {
		return nil, err
	}
	return &updated, nil
}

// GetByID retrieves a log target by its ID.
func (s *LogTargetsService) GetByID(ctx context.Context, id string) (*LogTarget, error) {
	if id == "" {
		return nil, fmt.Errorf("log target ID is required")
	}
	endpoint := fmt.Sprintf("/logtargets/%s", id)

	var logTarget LogTarget
	if err := s.Client.Get(ctx, endpoint, &logTarget); err != nil {
		return nil, err
	}
	return &logTarget, nil
}

// DeleteByID deletes a log target by its ID.
func (s *LogTargetsService) DeleteByID(ctx context.Context, id string) error {
	if id == "" {
		return fmt.Errorf("log target ID is required")
	}
	endpoint := fmt.Sprintf("/logtargets/%s", id)
	return s.Client.Delete(ctx, endpoint, nil)
}

// SetLogging sets services logging for a log target.
func (s *LogTargetsService) SetLogging(ctx context.Context, id string, req SetLoggingRequest) (*LogTarget, error) {
	if id == "" {
		return nil, fmt.Errorf("log target ID is required")
	}
	endpoint := fmt.Sprintf("/logtargets/%s/logging", id)

	var result LogTarget
	if err := s.Client.Put(ctx, endpoint, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
