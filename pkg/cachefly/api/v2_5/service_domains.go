package v2_5

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/cachefly/cachefly-go-sdk/internal/httpclient"
)

// ServiceDomain represents a domain attached to a service.
type ServiceDomain struct {
	ID               string   `json:"_id"`
	UpdatedAt        string   `json:"updateAt"`
	CreatedAt        string   `json:"createdAt"`
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	Service          string   `json:"service"`
	Certificates     []string `json:"certificates"`
	ValidationMode   string   `json:"validationMode"`
	ValidationTarget string   `json:"validationTarget"`
	ValidationStatus string   `json:"validationStatus"`
}

// ListServiceDomainsResponse wraps the paged list of domains.
type ListServiceDomainsResponse struct {
	Meta    MetaInfo        `json:"meta"`
	Domains []ServiceDomain `json:"data"`
}

// ListServiceDomainsOptions allows filtering & pagination.
type ListServiceDomainsOptions struct {
	Search       string
	Offset       int
	Limit        int
	ResponseType string
}

// CreateServiceDomainRequest is the payload to add a domain.
type CreateServiceDomainRequest struct {
	Name           string `json:"name"`
	Description    string `json:"description,omitempty"`
	ValidationMode string `json:"validationMode,omitempty"`
}

// UpdateServiceDomainRequest is the payload to update a domain.
type UpdateServiceDomainRequest struct {
	Name           string `json:"name,omitempty"`
	Description    string `json:"description,omitempty"`
	ValidationMode string `json:"validationMode,omitempty"`
}

// ServiceDomainsService handles Service Domains endpoints.
type ServiceDomainsService struct {
	Client *httpclient.Client
}

// List returns all domains for a given service ID.
func (s *ServiceDomainsService) List(ctx context.Context, sid string, opts ListServiceDomainsOptions) (*ListServiceDomainsResponse, error) {
	if sid == "" {
		return nil, fmt.Errorf("service ID is required")
	}
	endpoint := fmt.Sprintf("/services/%s/domains", sid)

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
	var resp ListServiceDomainsResponse
	if err := s.Client.Get(ctx, fullURL, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Create adds a new domain to the service.
func (s *ServiceDomainsService) Create(ctx context.Context, sid string, req CreateServiceDomainRequest) (*ServiceDomain, error) {
	if sid == "" {
		return nil, fmt.Errorf("service ID is required")
	}
	endpoint := fmt.Sprintf("/services/%s/domains", sid)

	var created ServiceDomain
	if err := s.Client.Post(ctx, endpoint, req, &created); err != nil {
		return nil, err
	}
	return &created, nil
}

// GetByID fetches a single domain by its ID.
func (s *ServiceDomainsService) GetByID(ctx context.Context, sid, id, responseType string) (*ServiceDomain, error) {
	if sid == "" || id == "" {
		return nil, fmt.Errorf("service ID and domain ID are required")
	}
	endpoint := fmt.Sprintf("/services/%s/domains/%s", sid, id)

	params := url.Values{}
	if responseType != "" {
		params.Set("responseType", responseType)
	}
	fullURL := fmt.Sprintf("%s?%s", endpoint, params.Encode())

	var domain ServiceDomain
	if err := s.Client.Get(ctx, fullURL, &domain); err != nil {
		return nil, err
	}
	return &domain, nil
}

// UpdateByID updates an existing service domain.
func (s *ServiceDomainsService) UpdateByID(ctx context.Context, sid, id string, req UpdateServiceDomainRequest) (*ServiceDomain, error) {
	if sid == "" || id == "" {
		return nil, fmt.Errorf("service ID and domain ID are required")
	}
	endpoint := fmt.Sprintf("/services/%s/domains/%s", sid, id)

	var updated ServiceDomain
	if err := s.Client.Put(ctx, endpoint, req, &updated); err != nil {
		return nil, err
	}
	return &updated, nil
}

// DeleteByID removes a domain from the service.
func (s *ServiceDomainsService) DeleteByID(ctx context.Context, sid, id string) error {
	if sid == "" || id == "" {
		return fmt.Errorf("service ID and domain ID are required")
	}
	endpoint := fmt.Sprintf("/services/%s/domains/%s", sid, id)
	return s.Client.Delete(ctx, endpoint, nil)
}

// ValidationReady signals that the domain is ready for validation.
func (s *ServiceDomainsService) ValidationReady(ctx context.Context, sid, id string) (*ServiceDomain, error) {
	if sid == "" || id == "" {
		return nil, fmt.Errorf("service ID and domain ID are required")
	}
	endpoint := fmt.Sprintf("/services/%s/domains/%s/validationReady", sid, id)

	var result ServiceDomain
	// Empty JSON body ensures Content-Type header is set
	if err := s.Client.Put(ctx, endpoint, struct{}{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
