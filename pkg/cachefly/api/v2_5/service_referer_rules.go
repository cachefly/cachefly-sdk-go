// Package v2_5 provides types and services for CacheFly API v2.5.
package v2_5

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/cachefly/cachefly-go-sdk/internal/httpclient"
)

// ServiceOptionsRefererRulesService handles referer rule API operations..
type ServiceOptionsRefererRulesService struct {
	Client *httpclient.Client
}

// RefererRule represents a referer access control rule for a service.
type RefererRule struct {
	ID            string   `json:"_id"`
	Directory     string   `json:"directory"`
	Extension     string   `json:"extension,omitempty"`
	Exceptions    []string `json:"exceptions"`
	DefaultAction string   `json:"defaultAction"`
	Order         int      `json:"order,omitempty"`
}

// ListRefererRulesOptions specifies pagination for listing referer rules.
type ListRefererRulesOptions struct {
	Offset int
	Limit  int
}

// ListRefererRulesResponse contains paginated referer rule results.
type ListRefererRulesResponse struct {
	Meta  MetaInfo      `json:"meta"`
	Rules []RefererRule `json:"data"`
}

// CreateRefererRuleRequest contains the required fields for creating a referer rule.
type CreateRefererRuleRequest struct {
	Directory     string   `json:"directory"`
	Extension     string   `json:"extension,omitempty"`
	Exceptions    []string `json:"exceptions"`
	DefaultAction string   `json:"defaultAction"`
}

// UpdateRefererRuleRequest contains fields for updating an existing referer rule.
type UpdateRefererRuleRequest struct {
	Directory     string   `json:"directory,omitempty"`
	Extension     string   `json:"extension,omitempty"`
	Exceptions    []string `json:"exceptions,omitempty"`
	DefaultAction string   `json:"defaultAction,omitempty"`
	Order         int      `json:"order,omitempty"`
}

// List retrieves referer rules for a service with optional pagination.
func (s *ServiceOptionsRefererRulesService) List(ctx context.Context, sid string, opts ListRefererRulesOptions) (*ListRefererRulesResponse, error) {
	if sid == "" {
		return nil, fmt.Errorf("service ID is required")
	}

	endpoint := fmt.Sprintf("/services/%s/options/refererrules", sid)
	params := url.Values{}

	if opts.Offset >= 0 {
		params.Set("offset", strconv.Itoa(opts.Offset))
	}
	if opts.Limit > 0 {
		params.Set("limit", strconv.Itoa(opts.Limit))
	}

	fullURL := fmt.Sprintf("%s?%s", endpoint, params.Encode())

	var resp ListRefererRulesResponse
	if err := s.Client.Get(ctx, fullURL, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Create adds a new referer rule to a service.
func (s *ServiceOptionsRefererRulesService) Create(ctx context.Context, sid string, req CreateRefererRuleRequest) (*RefererRule, error) {
	if sid == "" {
		return nil, fmt.Errorf("service ID is required")
	}

	endpoint := fmt.Sprintf("/services/%s/options/refererrules", sid)

	var created RefererRule
	if err := s.Client.Post(ctx, endpoint, req, &created); err != nil {
		return nil, err
	}
	return &created, nil
}

// GetByID retrieves a specific referer rule by service ID and rule ID.
func (s *ServiceOptionsRefererRulesService) GetByID(ctx context.Context, sid, id string) (*RefererRule, error) {
	if sid == "" || id == "" {
		return nil, fmt.Errorf("service ID and rule ID are required")
	}

	endpoint := fmt.Sprintf("/services/%s/options/refererrules/%s", sid, id)

	var rule RefererRule
	if err := s.Client.Get(ctx, endpoint, &rule); err != nil {
		return nil, err
	}
	return &rule, nil
}

// Update modifies an existing referer rule.
func (s *ServiceOptionsRefererRulesService) Update(ctx context.Context, sid, id string, req UpdateRefererRuleRequest) (*RefererRule, error) {
	if sid == "" || id == "" {
		return nil, fmt.Errorf("service ID and rule ID are required")
	}

	endpoint := fmt.Sprintf("/services/%s/options/refererrules/%s", sid, id)

	var updated RefererRule
	if err := s.Client.Put(ctx, endpoint, req, &updated); err != nil {
		return nil, err
	}
	return &updated, nil
}

// Delete removes a referer rule from a service.
func (s *ServiceOptionsRefererRulesService) Delete(ctx context.Context, sid, id string) error {
	if sid == "" || id == "" {
		return fmt.Errorf("service ID and rule ID are required")
	}

	endpoint := fmt.Sprintf("/services/%s/options/refererrules/%s", sid, id)
	return s.Client.Delete(ctx, endpoint, nil)
}
