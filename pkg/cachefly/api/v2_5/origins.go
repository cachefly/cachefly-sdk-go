package v2_5

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/cachefly/cachefly-go-sdk/internal/httpclient"
)

// OriginsService handles origin configuration operations.
type OriginsService struct {
	Client *httpclient.Client
}

// Origin represents an origin configuration in CacheFly.
type Origin struct {
	ID                     string `json:"_id"`
	UpdatedAt              string `json:"updateAt"`
	CreatedAt              string `json:"createdAt"`
	Type                   string `json:"type,omitempty"`
	Name                   string `json:"name,omitempty"`
	Hostname               string `json:"hostname"`
	CacheByQueryParam      bool   `json:"cacheByQueryParam"`
	Gzip                   bool   `json:"gzip"`
	Scheme                 string `json:"scheme"`
	TTL                    int    `json:"ttl"`
	MissedTTL              int    `json:"missedTtl"`
	ConnectionTimeout      int    `json:"connectionTimeout,omitempty"`
	TimeToFirstByteTimeout int    `json:"timeToFirstByteTimeout,omitempty"`
	AccessKey              string `json:"accessKey,omitempty"`
	SecretKey              string `json:"secretKey,omitempty"`
	Region                 string `json:"region,omitempty"`
	SignatureVersion       string `json:"signatureVersion,omitempty"`
}

// ListOriginsResponse wraps paginated origin list.
type ListOriginsResponse struct {
	Meta    MetaInfo `json:"meta"`
	Origins []Origin `json:"data"`
}

// ListOriginsOptions configures filtering and pagination.
type ListOriginsOptions struct {
	Type         string
	Offset       int
	Limit        int
	ResponseType string
}

// CreateOriginRequest is the payload for creating a new origin.
type CreateOriginRequest struct {
	Type                   string `json:"type"`
	Name                   string `json:"name,omitempty"`
	Hostname               string `json:"hostname"`
	Gzip                   bool   `json:"gzip,omitempty"`
	CacheByQueryParam      bool   `json:"cacheByQueryParam,omitempty"`
	Scheme                 string `json:"scheme,omitempty"`
	TTL                    int    `json:"ttl,omitempty"`
	MissedTTL              int    `json:"missedTtl,omitempty"`
	ConnectionTimeout      int    `json:"connectionTimeout,omitempty"`
	TimeToFirstByteTimeout int    `json:"timeToFirstByteTimeout,omitempty"`
	AccessKey              string `json:"accessKey,omitempty"`
	SecretKey              string `json:"secretKey,omitempty"`
	Region                 string `json:"region,omitempty"`
	SignatureVersion       string `json:"signatureVersion,omitempty"`
}

// UpdateOriginRequest is the payload for updating an existing origin.
type UpdateOriginRequest struct {
	Type                   string `json:"type,omitempty"`
	Name                   string `json:"name,omitempty"`
	Hostname               string `json:"hostname,omitempty"`
	Gzip                   bool   `json:"gzip,omitempty"`
	CacheByQueryParam      bool   `json:"cacheByQueryParam,omitempty"`
	Scheme                 string `json:"scheme,omitempty"`
	TTL                    int    `json:"ttl,omitempty"`
	MissedTTL              int    `json:"missedTtl,omitempty"`
	ConnectionTimeout      int    `json:"connectionTimeout,omitempty"`
	TimeToFirstByteTimeout int    `json:"timeToFirstByteTimeout,omitempty"`
	AccessKey              string `json:"accessKey,omitempty"`
	SecretKey              string `json:"secretKey,omitempty"`
	Region                 string `json:"region,omitempty"`
	SignatureVersion       string `json:"signatureVersion,omitempty"`
}

// List retrieves all origins with optional filters.
func (s *OriginsService) List(ctx context.Context, opts ListOriginsOptions) (*ListOriginsResponse, error) {
	endpoint := "/origins"
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

	var resp ListOriginsResponse
	if err := s.Client.Get(ctx, fullURL, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Create adds a new origin.
func (s *OriginsService) Create(ctx context.Context, req CreateOriginRequest) (*Origin, error) {
	endpoint := "/origins"
	var created Origin
	if err := s.Client.Post(ctx, endpoint, req, &created); err != nil {
		return nil, err
	}
	return &created, nil
}

// GetByID fetches a single origin by its ID.
func (s *OriginsService) GetByID(ctx context.Context, id, responseType string) (*Origin, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/origins/%s", id)
	params := url.Values{}
	if responseType != "" {
		params.Set("responseType", responseType)
	}
	fullURL := fmt.Sprintf("%s?%s", endpoint, params.Encode())
	var origin Origin
	if err := s.Client.Get(ctx, fullURL, &origin); err != nil {
		return nil, err
	}
	return &origin, nil
}

// UpdateByID modifies an existing origin.
func (s *OriginsService) UpdateByID(ctx context.Context, id string, req UpdateOriginRequest) (*Origin, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/origins/%s", id)
	var updated Origin
	if err := s.Client.Put(ctx, endpoint, req, &updated); err != nil {
		return nil, err
	}
	return &updated, nil
}

// Delete removes an origin by ID.
func (s *OriginsService) Delete(ctx context.Context, id string) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/origins/%s", id)
	return s.Client.Delete(ctx, endpoint, nil)
}
