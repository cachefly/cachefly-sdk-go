package v2_6

import (
	"context"
	"fmt"
)

// PurgeRequest represents a request to purge cached content for a service.
// Either set All=true to purge everything, or provide one or more Paths.
// Directory paths should end with a trailing slash.
type PurgeRequest struct {
	Paths []string `json:"paths,omitempty"`
	All   bool     `json:"all,omitempty"`
}

// Purge triggers a cache purge for a service.
// Provide either All=true to purge everything, or a list of Paths to purge specific objects/directories.
func (s *ServicesService) Purge(ctx context.Context, id string, req PurgeRequest) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}
	if !req.All && len(req.Paths) == 0 {
		return fmt.Errorf("either 'all' must be true or 'paths' must be provided")
	}

	endpoint := fmt.Sprintf("/services/%s/purge", id)

	// Purge does not need a response body; just check for HTTP errors
	if err := s.Client.Put(ctx, endpoint, req, nil); err != nil {
		return err
	}
	return nil
}
