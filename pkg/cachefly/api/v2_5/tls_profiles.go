package v2_5

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/cachefly/cachefly-go-sdk/internal/httpclient"
)

// TLSProfile represents a TLS configuration profile in CacheFly.
type TLSProfile struct {
	ID        string `json:"_id"`
	UpdatedAt string `json:"updateAt"`
	CreatedAt string `json:"createdAt"`
	Name      string `json:"name"`
	// Add more fields here once the schema is known
}

// ListTLSProfilesResponse contains paginated TLS profile results.
type ListTLSProfilesResponse struct {
	Meta     MetaInfo     `json:"meta"`
	Profiles []TLSProfile `json:"data"`
}

// ListTLSProfilesOptions specifies filtering and pagination for listing TLS profiles.
type ListTLSProfilesOptions struct {
	SortBy []string
	Group  string
	Offset int
	Limit  int
}

// TLSProfilesService handles TLS profile operations.
type TLSProfilesService struct {
	Client *httpclient.Client
}

// List retrieves TLS profiles with optional sorting, grouping, and pagination.
func (s *TLSProfilesService) List(ctx context.Context, opts ListTLSProfilesOptions) (*ListTLSProfilesResponse, error) {
	endpoint := "/tlsprofiles"
	params := url.Values{}

	if len(opts.SortBy) > 0 {
		for _, v := range opts.SortBy {
			params.Add("sortBy", v)
		}
	}
	if opts.Group != "" {
		params.Set("group", opts.Group)
	}
	if opts.Offset >= 0 {
		params.Set("offset", strconv.Itoa(opts.Offset))
	}
	if opts.Limit > 0 {
		params.Set("limit", strconv.Itoa(opts.Limit))
	}

	fullURL := fmt.Sprintf("%s?%s", endpoint, params.Encode())

	var resp ListTLSProfilesResponse
	if err := s.Client.Get(ctx, fullURL, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetByID retrieves a TLS profile by its ID.
func (s *TLSProfilesService) GetByID(ctx context.Context, id string) (*TLSProfile, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("/tlsprofiles/%s", id)

	var p TLSProfile
	if err := s.Client.Get(ctx, endpoint, &p); err != nil {
		return nil, err
	}
	return &p, nil
}
