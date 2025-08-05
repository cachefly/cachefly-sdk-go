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
	ID          string                 `json:"_id"`
	UpdatedAt   string                 `json:"updatedAt"`
	CreatedAt   string                 `json:"createdAt"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Type        string                 `json:"type"`
	Enabled     bool                   `json:"enabled"`
	Config      map[string]interface{} `json:"config"`
	Services    []string               `json:"services"`
}

// ListLogTargetsResponse contains paginated log target results.
type ListLogTargetsResponse struct {
	Meta       MetaInfo    `json:"meta"`
	LogTargets []LogTarget `json:"data"`
}

// ListLogTargetsOptions allows filtering & pagination for log targets.
type ListLogTargetsOptions struct {
	Search       string
	Offset       int
	Limit        int
	ResponseType string
}

// CreateLogTargetRequest contains the required fields for creating a new log target.
type CreateLogTargetRequest struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description,omitempty"`
	Type        string                 `json:"type"`
	Config      map[string]interface{} `json:"config"`
}

// UpdateLogTargetRequest contains the fields for updating an existing log target.
type UpdateLogTargetRequest struct {
	Name        string                 `json:"name,omitempty"`
	Description string                 `json:"description,omitempty"`
	Type        string                 `json:"type,omitempty"`
	Config      map[string]interface{} `json:"config,omitempty"`
	Enabled     *bool                  `json:"enabled,omitempty"`
}

// EnableLoggingRequest contains the services to enable logging for.
type EnableLoggingRequest struct {
	Services []string `json:"services"`
}

// LogTargetsService handles log target-related API operations.
type LogTargetsService struct {
	Client *httpclient.Client
}

// List returns all log targets for the current account.
func (s *LogTargetsService) List(ctx context.Context, opts ListLogTargetsOptions) (*ListLogTargetsResponse, error) {
	endpoint := "/logTargets"

	params := url.Values{}
	if opts.Search != "" {
		params.Set("search", opts.Search)
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
	endpoint := "/logTargets"

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
	endpoint := fmt.Sprintf("/logTargets/%s", id)

	var updated LogTarget
	if err := s.Client.Put(ctx, endpoint, req, &updated); err != nil {
		return nil, err
	}
	return &updated, nil
}

// DeleteByID deletes a log target by its ID.
func (s *LogTargetsService) DeleteByID(ctx context.Context, id string) error {
	if id == "" {
		return fmt.Errorf("log target ID is required")
	}
	endpoint := fmt.Sprintf("/logTargets/%s", id)
	return s.Client.Delete(ctx, endpoint, nil)
}

// EnableLogging enables services logging for a log target.
func (s *LogTargetsService) EnableLogging(ctx context.Context, id string, req EnableLoggingRequest) (*LogTarget, error) {
	if id == "" {
		return nil, fmt.Errorf("log target ID is required")
	}
	endpoint := fmt.Sprintf("/logTargets/%s/enableLogging", id)

	var result LogTarget
	if err := s.Client.Put(ctx, endpoint, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
