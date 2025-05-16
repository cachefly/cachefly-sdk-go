package api

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/avvvet/cachefly-sdk-go/internal/httpclient"
)

type Service struct {
	ID                string `json:"_id"`
	UpdatedAt         string `json:"updateAt"`
	CreatedAt         string `json:"createdAt"`
	Name              string `json:"name"`
	UniqueName        string `json:"uniqueName"`
	AutoSSL           bool   `json:"autoSsl"`
	ConfigurationMode string `json:"configurationMode"`
	Status            string `json:"status"`
}

type CreateServiceRequest struct {
	Name        string `json:"name"`
	UniqueName  string `json:"uniqueName"`
	Description string `json:"description"`
}

type ListServicesResponse struct {
	Meta     MetaInfo  `json:"meta"`
	Services []Service `json:"data"`
}

type MetaInfo struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Count  int `json:"count"`
}

type ListOptions struct {
	ResponseType    string
	IncludeFeatures bool
	Status          string
	Offset          int
	Limit           int
}

// UpdateServiceRequest is the payload for updating a service by ID.
type UpdateServiceRequest struct {
	Description    string `json:"description"`
	TLSProfile     string `json:"tlsProfile"`
	AutoSSL        bool   `json:"autoSsl"`
	DeliveryRegion string `json:"deliveryRegion"`
}

// communication with the services endpoint.
type ServicesService struct {
	Client *httpclient.Client
}

// payload for enabling access logging on a service.
type EnableAccessLogsRequest struct {
	LogTarget string `json:"logTarget"`
}

type EnableOriginLogsRequest struct {
	LogTarget string `json:"logTarget"`
}

// Create a new service
func (s *ServicesService) Create(ctx context.Context, req CreateServiceRequest) (*Service, error) {
	endpoint := "/services"

	var created Service
	err := s.Client.Post(ctx, endpoint, req, &created)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

// Get service by ID
func (s *ServicesService) Get(ctx context.Context, id string, responseType string, includeFeatures bool) (*Service, error) {
	endpoint := fmt.Sprintf("/services/%s", id)

	params := url.Values{}
	if responseType != "" {
		params.Set("responseType", responseType)
	}
	params.Set("includeFeatures", strconv.FormatBool(includeFeatures))

	fullURL := fmt.Sprintf("%s?%s", endpoint, params.Encode())

	var result Service
	err := s.Client.Get(ctx, fullURL, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// List services
func (s *ServicesService) List(ctx context.Context, opts ListOptions) (*ListServicesResponse, error) {
	endpoint := "/services"
	params := url.Values{}

	if opts.ResponseType != "" {
		params.Set("responseType", opts.ResponseType)
	}
	params.Set("includeFeatures", strconv.FormatBool(opts.IncludeFeatures))
	if opts.Status != "" {
		params.Set("status", opts.Status)
	}
	if opts.Offset >= 0 {
		params.Set("offset", strconv.Itoa(opts.Offset))
	}
	if opts.Limit > 0 {
		params.Set("limit", strconv.Itoa(opts.Limit))
	}

	fullURL := fmt.Sprintf("%s?%s", endpoint, params.Encode())

	var result ListServicesResponse
	err := s.Client.Get(ctx, fullURL, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateServiceByID updates an existing service.
// id is required.
func (s *ServicesService) UpdateServiceByID(ctx context.Context, id string, req UpdateServiceRequest) (*Service, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/services/%s", id)

	var updated Service
	if err := s.Client.Put(ctx, endpoint, req, &updated); err != nil {
		return nil, err
	}
	return &updated, nil
}

// ActivateServiceByID
// id is required. Returns the updated Service.
func (s *ServicesService) ActivateServiceByID(ctx context.Context, id string) (*Service, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/services/%s/activate", id)

	var updated Service
	if err := s.Client.Put(ctx, endpoint, struct{}{}, &updated); err != nil {
		return nil, err
	}
	return &updated, nil
}

// id is required. Returns the updated Service.
func (s *ServicesService) DeactivateServiceByID(ctx context.Context, id string) (*Service, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/services/%s/deactivate", id)

	var updated Service
	if err := s.Client.Put(ctx, endpoint, struct{}{}, &updated); err != nil {
		return nil, err
	}
	return &updated, nil
}

// id is required. Returns the updated Service.
func (s *ServicesService) EnableAccessLogging(ctx context.Context, id string, req EnableAccessLogsRequest) (*Service, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/services/%s/accessLogs", id)

	var updated Service
	if err := s.Client.Put(ctx, endpoint, req, &updated); err != nil {
		return nil, err
	}
	return &updated, nil
}

// id is required.
func (s *ServicesService) DeleteAccessLoggingByID(ctx context.Context, id string) (*Service, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/services/%s/accessLogs", id)

	var updated Service
	if err := s.Client.Delete(ctx, endpoint, &updated); err != nil {
		return nil, err
	}
	return &updated, nil
}

// id is required.
func (s *ServicesService) EnableOriginLogging(ctx context.Context, id string, req EnableOriginLogsRequest) (*Service, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/services/%s/originLogs", id)

	var updated Service
	if err := s.Client.Put(ctx, endpoint, req, &updated); err != nil {
		return nil, err
	}
	return &updated, nil
}

// id is required.
func (s *ServicesService) DeleteOriginLoggingByID(ctx context.Context, id string) (*Service, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/services/%s/originLogs", id)

	var updated Service
	if err := s.Client.Delete(ctx, endpoint, &updated); err != nil {
		return nil, err
	}
	return &updated, nil
}
