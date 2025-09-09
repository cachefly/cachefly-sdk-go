package v2_6

import (
	"context"
	"fmt"

	"github.com/cachefly/cachefly-sdk-go/internal/httpclient"
)

// ServiceStatsService exposes service-level stats endpoints.
type ServiceStatsService struct {
	Client *httpclient.Client
}

func (s *ServiceStatsService) get(ctx context.Context, sid string, endpoint string, opts StatsQueryOptions) (*StatsResponse, error) {
	if sid == "" {
		return nil, fmt.Errorf("service id is required")
	}
	params := opts.toURLValues()
	url := fmt.Sprintf("/services/%s/stats/%s", sid, endpoint)
	if len(params) > 0 {
		url = fmt.Sprintf("%s?%s", url, params.Encode())
	}
	var out StatsResponse
	if err := s.Client.Get(ctx, url, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// POP returns service POP stats.
// Docs: https://portal.cachefly.com/api/2.6/docs/#tag/Service-Stats
func (s *ServiceStatsService) POP(ctx context.Context, sid string, opts StatsQueryOptions) (*StatsResponse, error) {
	if opts.From == "" || opts.To == "" {
		return nil, fmt.Errorf("'from' and 'to' parameters are required")
	}
	return s.get(ctx, sid, "pop", opts)
}

// Country returns service country stats.
func (s *ServiceStatsService) Country(ctx context.Context, sid string, opts StatsQueryOptions) (*StatsResponse, error) {
	if opts.From == "" || opts.To == "" {
		return nil, fmt.Errorf("'from' and 'to' parameters are required")
	}
	return s.get(ctx, sid, "country", opts)
}

// Cache returns service cache stats.
func (s *ServiceStatsService) Cache(ctx context.Context, sid string, opts StatsQueryOptions) (*StatsResponse, error) {
	if opts.From == "" || opts.To == "" {
		return nil, fmt.Errorf("'from' and 'to' parameters are required")
	}
	return s.get(ctx, sid, "cache", opts)
}

// Status returns service status stats.
func (s *ServiceStatsService) Status(ctx context.Context, sid string, opts StatsQueryOptions) (*StatsResponse, error) {
	if opts.From == "" || opts.To == "" {
		return nil, fmt.Errorf("'from' and 'to' parameters are required")
	}
	return s.get(ctx, sid, "status", opts)
}

// Realtime returns service realtime stats.
func (s *ServiceStatsService) Realtime(ctx context.Context, sid string, opts StatsQueryOptions) (*StatsResponse, error) {
	return s.get(ctx, sid, "realtime", opts)
}

// Path returns service path stats.
func (s *ServiceStatsService) Path(ctx context.Context, sid string, opts StatsQueryOptions) (*StatsResponse, error) {
	if opts.From == "" || opts.To == "" {
		return nil, fmt.Errorf("'from' and 'to' parameters are required")
	}
	return s.get(ctx, sid, "path", opts)
}

// Referer returns service referer stats.
func (s *ServiceStatsService) Referer(ctx context.Context, sid string, opts StatsQueryOptions) (*StatsResponse, error) {
	if opts.From == "" || opts.To == "" {
		return nil, fmt.Errorf("'from' and 'to' parameters are required")
	}
	return s.get(ctx, sid, "referer", opts)
}

// Origin returns service origin stats.
func (s *ServiceStatsService) Origin(ctx context.Context, sid string, opts StatsQueryOptions) (*StatsResponse, error) {
	if opts.From == "" || opts.To == "" {
		return nil, fmt.Errorf("'from' and 'to' parameters are required")
	}
	return s.get(ctx, sid, "origin", opts)
}
