package v2_5

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/cachefly/cachefly-go-sdk/internal/httpclient"
)

// ServiceRule represents a rule configuration for a service.
type ServiceRule struct {
	ID        string `json:"_id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updateAt"`
}

// ListServiceRulesResponse contains paginated service rule results.
type ListServiceRulesResponse struct {
	Meta  MetaInfo      `json:"meta"`
	Rules []ServiceRule `json:"data"`
}

// ListServiceRulesOptions specifies filters and pagination for listing service rules.
type ListServiceRulesOptions struct {
	Offset       int
	Limit        int
	ResponseType string
}

// ServiceRulesService handles service rule operations.
type ServiceRulesService struct {
	Client *httpclient.Client
}

// UpdateServiceRulesRequest contains rules for bulk update operations.
type UpdateServiceRulesRequest struct {
	Rules []ServiceRule `json:"rules"`
}

// List retrieves rules for a service with optional filtering and pagination.
func (s *ServiceRulesService) List(ctx context.Context, serviceID string, opts ListServiceRulesOptions) (*ListServiceRulesResponse, error) {
	if serviceID == "" {
		return nil, fmt.Errorf("serviceID is required")
	}

	endpoint := fmt.Sprintf("/services/%s/rules", serviceID)
	params := url.Values{}

	if opts.ResponseType != "" {
		params.Set("responseType", opts.ResponseType)
	}
	if opts.Offset >= 0 {
		params.Set("offset", strconv.Itoa(opts.Offset))
	}
	if opts.Limit > 0 {
		params.Set("limit", strconv.Itoa(opts.Limit))
	}

	fullURL := fmt.Sprintf("%s?%s", endpoint, params.Encode())

	var resp ListServiceRulesResponse
	if err := s.Client.Get(ctx, fullURL, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Update performs a bulk update of rules for a service.
func (s *ServiceRulesService) Update(ctx context.Context, serviceID string, req UpdateServiceRulesRequest) (*ListServiceRulesResponse, error) {
	if serviceID == "" {
		return nil, fmt.Errorf("serviceID is required")
	}

	endpoint := fmt.Sprintf("/services/%s/rules", serviceID)

	var resp ListServiceRulesResponse
	if err := s.Client.Put(ctx, endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSchema retrieves the JSON schema for service rules.
func (s *ServiceRulesService) GetSchema(ctx context.Context, serviceID string) (map[string]interface{}, error) {
	if serviceID == "" {
		return nil, fmt.Errorf("serviceID is required")
	}

	endpoint := fmt.Sprintf("/services/%s/rules/schema", serviceID)

	var schema map[string]interface{}
	if err := s.Client.Get(ctx, endpoint, &schema); err != nil {
		return nil, err
	}
	return schema, nil
}
