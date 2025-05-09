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

// communication with the services endpoint.
type ServicesService struct {
	Client *httpclient.Client
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
