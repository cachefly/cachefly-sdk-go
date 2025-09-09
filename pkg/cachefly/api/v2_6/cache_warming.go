package v2_6

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/cachefly/cachefly-sdk-go/internal/httpclient"
)

// CacheWarmingTask represents a cache warming task.
type CacheWarmingTask struct {
	ID               string      `json:"_id"`
	Name             string      `json:"name"`
	Targets          []string    `json:"targets"`
	Regions          []string    `json:"regions,omitempty"`
	ContentTypes     []string    `json:"contentTypes,omitempty"`
	ContentEncodings []string    `json:"contentEncodings,omitempty"`
	ContentLanguages []string    `json:"contentLanguages,omitempty"`
	Status           string      `json:"status,omitempty"`
	StartedAt        string      `json:"startedAt,omitempty"`
	StoppedAt        string      `json:"stoppedAt,omitempty"`
	TaskType         interface{} `json:"taskType,omitempty"`
	Properties       interface{} `json:"properties,omitempty"`
	CreatedAt        string      `json:"createdAt,omitempty"`
	UpdatedAt        string      `json:"updatedAt,omitempty"`
}

// CreateCacheWarmingTaskRequest is the payload for creating a task.
type CreateCacheWarmingTaskRequest struct {
	Name             string   `json:"name,omitempty"`
	Targets          []string `json:"targets"`
	Regions          []string `json:"regions"`
	ContentTypes     []string `json:"contentTypes,omitempty"`
	ContentEncodings []string `json:"contentEncodings,omitempty"`
	ContentLanguages []string `json:"contentLanguages,omitempty"`
}

// ListCacheWarmingTasksOptions allows pagination of tasks list.
type ListCacheWarmingTasksOptions struct {
	Offset       int
	Limit        int
	ResponseType string
}

// ListCacheWarmingTasksResponse wraps paginated tasks.
type ListCacheWarmingTasksResponse struct {
	Meta MetaInfo           `json:"meta"`
	Data []CacheWarmingTask `json:"data"`
}

// CacheWarmingService manages cache warming endpoints.
type CacheWarmingService struct {
	Client *httpclient.Client
}

// List returns cache warming tasks with optional pagination.
func (s *CacheWarmingService) List(ctx context.Context, opts ListCacheWarmingTasksOptions) (*ListCacheWarmingTasksResponse, error) {
	endpoint := "/cachewarming"

	params := url.Values{}
	if opts.Offset >= 0 {
		params.Set("offset", strconv.Itoa(opts.Offset))
	}

	if opts.Limit > 0 {
		params.Set("limit", strconv.Itoa(opts.Limit))
	}

	fullURL := fmt.Sprintf("%s?%s", endpoint, params.Encode())

	var resp ListCacheWarmingTasksResponse
	if err := s.Client.Get(ctx, fullURL, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Create creates a new cache warming task.
func (s *CacheWarmingService) Create(ctx context.Context, req CreateCacheWarmingTaskRequest) (*CacheWarmingTask, error) {
	endpoint := "/cachewarming"
	if len(req.Targets) == 0 {
		return nil, fmt.Errorf("at least one target is required")
	}

	if len(req.Regions) == 0 {
		return nil, fmt.Errorf("at least one region is required")
	}

	var created CacheWarmingTask
	if err := s.Client.Post(ctx, endpoint, req, &created); err != nil {
		return nil, err
	}
	return &created, nil
}

// GetByID returns info about a single cache warming task.
func (s *CacheWarmingService) GetByID(ctx context.Context, id string) (*CacheWarmingTask, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("/cachewarming/%s", id)

	var task CacheWarmingTask
	if err := s.Client.Get(ctx, endpoint, &task); err != nil {
		return nil, err
	}
	return &task, nil
}

// DeleteByID deletes a cache warming task.
func (s *CacheWarmingService) DeleteByID(ctx context.Context, id string) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/cachewarming/%s", id)
	return s.Client.Delete(ctx, endpoint, nil)
}
