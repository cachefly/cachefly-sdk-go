package v2_6

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/cachefly/cachefly-sdk-go/internal/httpclient"
)

type MetaInfoDeliveryRegions struct {
	Limit  int      `json:"limit"`
	Offset int      `json:"offset"`
	Count  int      `json:"count"`
	SortBy []string `json:"sortBy"`
}

type DeliveryRegion struct {
	ID          string `json:"_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ListDeliveryRegionsResponse contains paginated delivery region results.
type ListDeliveryRegionsResponse struct {
	Meta    MetaInfoDeliveryRegions `json:"meta"`
	Regions []DeliveryRegion        `json:"data"`
}

// ListDeliveryRegionsOptions specifies filtering and pagination for listing delivery regions.
type ListDeliveryRegionsOptions struct {
	Search string
	SortBy []string
	Offset int
	Limit  int
}

// DeliveryRegionsService handles delivery region operations.
type DeliveryRegionsService struct {
	Client *httpclient.Client
}

// List retrieves delivery regions with optional sorting, grouping, and pagination.
func (s *DeliveryRegionsService) List(ctx context.Context, opts ListDeliveryRegionsOptions) (*ListDeliveryRegionsResponse, error) {
	endpoint := "/deliveryregions"
	params := url.Values{}

	if len(opts.SortBy) > 0 {
		for _, v := range opts.SortBy {
			params.Add("sortBy", v)
		}
	}
	if opts.Search != "" {
		if len(opts.Search) < 2 {
			return nil, fmt.Errorf("search must be at least 2 characters")
		}
		params.Set("search", opts.Search)
	}

	if opts.Offset >= 0 {
		params.Set("offset", strconv.Itoa(opts.Offset))
	}
	if opts.Limit > 0 {
		params.Set("limit", strconv.Itoa(opts.Limit))
	}

	fullURL := fmt.Sprintf("%s?%s", endpoint, params.Encode())

	var resp ListDeliveryRegionsResponse
	if err := s.Client.Get(ctx, fullURL, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
