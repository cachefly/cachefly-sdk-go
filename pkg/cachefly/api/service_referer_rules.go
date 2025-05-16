package api

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/avvvet/cachefly-sdk-go/internal/httpclient"
)

type ServiceOptionsRefererRulesService struct {
	Client *httpclient.Client
}

type RefererRule struct {
	ID            string   `json:"_id"`
	Directory     string   `json:"directory"`
	Extension     string   `json:"extension,omitempty"`
	Exceptions    []string `json:"exceptions"`
	DefaultAction string   `json:"defaultAction"`
	Order         int      `json:"order,omitempty"`
}

type ListRefererRulesOptions struct {
	Offset int
	Limit  int
}

type ListRefererRulesResponse struct {
	Meta  MetaInfo      `json:"meta"`
	Rules []RefererRule `json:"data"`
}

type CreateRefererRuleRequest struct {
	Directory     string   `json:"directory"`
	Extension     string   `json:"extension,omitempty"`
	Exceptions    []string `json:"exceptions"`
	DefaultAction string   `json:"defaultAction"`
}

type UpdateRefererRuleRequest struct {
	Directory     string   `json:"directory,omitempty"`
	Extension     string   `json:"extension,omitempty"`
	Exceptions    []string `json:"exceptions,omitempty"`
	DefaultAction string   `json:"defaultAction,omitempty"`
	Order         int      `json:"order,omitempty"`
}

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

func (s *ServiceOptionsRefererRulesService) Delete(ctx context.Context, sid, id string) error {
	if sid == "" || id == "" {
		return fmt.Errorf("service ID and rule ID are required")
	}
	endpoint := fmt.Sprintf("/services/%s/options/refererrules/%s", sid, id)
	return s.Client.Delete(ctx, endpoint, nil)
}
