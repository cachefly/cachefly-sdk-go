package v2_6

import (
	"context"
	"fmt"

	"github.com/cachefly/cachefly-sdk-go/internal/httpclient"
)

// AccountStatsService exposes account-level stats endpoints.
type AccountStatsService struct {
	Client *httpclient.Client
}

// internal helper for GET /stats endpoints under account scope
func (s *AccountStatsService) get(ctx context.Context, endpoint string, opts StatsQueryOptions) (*StatsResponse, error) {
	params := opts.toURLValues()
	url := endpoint
	if len(params) > 0 {
		url = fmt.Sprintf("%s?%s", endpoint, params.Encode())
	}

	var out StatsResponse
	if err := s.Client.Get(ctx, url, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// POP returns account POP stats.
// Docs: https://portal.cachefly.com/api/2.6/docs/#tag/Regular-Account-Stats
func (s *AccountStatsService) POP(ctx context.Context, opts StatsQueryOptions) (*StatsResponse, error) {
	if opts.From == "" || opts.To == "" {
		return nil, fmt.Errorf("'from' and 'to' parameters are required")
	}
	return s.get(ctx, "/stats/pop", opts)
}

// Country returns account country stats.
func (s *AccountStatsService) Country(ctx context.Context, opts StatsQueryOptions) (*StatsResponse, error) {
	if opts.From == "" || opts.To == "" {
		return nil, fmt.Errorf("'from' and 'to' parameters are required")
	}
	return s.get(ctx, "/stats/country", opts)
}

// Cache returns account cache stats.
func (s *AccountStatsService) Cache(ctx context.Context, opts StatsQueryOptions) (*StatsResponse, error) {
	if opts.From == "" || opts.To == "" {
		return nil, fmt.Errorf("'from' and 'to' parameters are required")
	}
	return s.get(ctx, "/stats/cache", opts)
}

// Status returns account status stats.
func (s *AccountStatsService) Status(ctx context.Context, opts StatsQueryOptions) (*StatsResponse, error) {
	if opts.From == "" || opts.To == "" {
		return nil, fmt.Errorf("'from' and 'to' parameters are required")
	}
	return s.get(ctx, "/stats/status", opts)
}

// Origin returns account origin stats.
func (s *AccountStatsService) Origin(ctx context.Context, opts StatsQueryOptions) (*StatsResponse, error) {
	if opts.From == "" || opts.To == "" {
		return nil, fmt.Errorf("'from' and 'to' parameters are required")
	}
	return s.get(ctx, "/stats/origin", opts)
}

// Storage returns account storage stats.
func (s *AccountStatsService) Storage(ctx context.Context, opts StatsQueryOptions) (*StatsResponse, error) {
	if opts.From == "" || opts.To == "" {
		return nil, fmt.Errorf("'from' and 'to' parameters are required")
	}
	return s.get(ctx, "/stats/storage", opts)
}

// Realtime returns account realtime stats.
func (s *AccountStatsService) Realtime(ctx context.Context, opts StatsQueryOptions) (*StatsResponse, error) {
	return s.get(ctx, "/stats/realtime", opts)
}

// Path returns account path stats.
func (s *AccountStatsService) Path(ctx context.Context, opts StatsQueryOptions) (*StatsResponse, error) {
	if opts.From == "" || opts.To == "" {
		return nil, fmt.Errorf("'from' and 'to' parameters are required")
	}
	return s.get(ctx, "/stats/path", opts)
}

// Referer returns account referer stats.
func (s *AccountStatsService) Referer(ctx context.Context, opts StatsQueryOptions) (*StatsResponse, error) {
	if opts.From == "" || opts.To == "" {
		return nil, fmt.Errorf("'from' and 'to' parameters are required")
	}
	return s.get(ctx, "/stats/referer", opts)
}
