package v2_6

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/cachefly/cachefly-sdk-go/internal/httpclient"
)

// EdgeControlLibraryService handles the edge control script library
// (/edgecontrol/library).
//
// SYSTEM scripts are curated by CacheFly and visible to every account; USER
// scripts are private to the account that created them. Only USER scripts can
// be created, updated or deleted.
type EdgeControlLibraryService struct {
	Client *httpclient.Client
}

// Edge control library script ownership types.
const (
	EdgeControlLibraryTypeUser   = "USER"
	EdgeControlLibraryTypeSystem = "SYSTEM"
)

// EdgeControlLibraryScript is a reusable edge control script.
type EdgeControlLibraryScript struct {
	ID        string `json:"_id"`
	Type      string `json:"type"`
	Kind      string `json:"kind"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// ListEdgeControlLibraryOptions filters and paginates library scripts.
type ListEdgeControlLibraryOptions struct {
	// Type is USER or SYSTEM.
	Type string
	Kind EdgeControlScriptKind
	// Search matches script names (at least 2 characters).
	Search string
	SortBy []string
	Offset int
	Limit  int
}

// ListEdgeControlLibraryResponse contains paginated library scripts.
type ListEdgeControlLibraryResponse struct {
	Meta    MetaInfo                   `json:"meta"`
	Scripts []EdgeControlLibraryScript `json:"data"`
}

// EdgeControlLibraryScriptRequest is the payload for creating or updating a USER
// script. The API requires every field on both create and update, and the code
// must define a "handler" function.
type EdgeControlLibraryScriptRequest struct {
	Name string                `json:"name"`
	Kind EdgeControlScriptKind `json:"kind"`
	Code string                `json:"code"`
}

// List returns SYSTEM scripts and the account's own USER scripts.
func (s *EdgeControlLibraryService) List(ctx context.Context, opts ListEdgeControlLibraryOptions) (*ListEdgeControlLibraryResponse, error) {
	params := url.Values{}
	if opts.Type != "" {
		params.Set("type", opts.Type)
	}
	if opts.Kind != "" {
		params.Set("kind", string(opts.Kind))
	}
	if opts.Search != "" {
		params.Set("search", opts.Search)
	}
	for _, sortBy := range opts.SortBy {
		params.Add("sortBy", sortBy)
	}
	if opts.Offset >= 0 {
		params.Set("offset", strconv.Itoa(opts.Offset))
	}
	if opts.Limit > 0 {
		params.Set("limit", strconv.Itoa(opts.Limit))
	}

	fullURL := fmt.Sprintf("/edgecontrol/library?%s", params.Encode())
	var resp ListEdgeControlLibraryResponse
	if err := s.Client.Get(ctx, fullURL, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetByID returns a SYSTEM script or one of the account's own USER scripts.
func (s *EdgeControlLibraryService) GetByID(ctx context.Context, id string) (*EdgeControlLibraryScript, error) {
	endpoint, err := edgeControlLibraryEndpoint(id)
	if err != nil {
		return nil, err
	}

	var script EdgeControlLibraryScript
	if err := s.Client.Get(ctx, endpoint, &script); err != nil {
		return nil, err
	}
	return &script, nil
}

// Create creates a USER script owned by the account.
func (s *EdgeControlLibraryService) Create(ctx context.Context, req EdgeControlLibraryScriptRequest) (*EdgeControlLibraryScript, error) {
	var created EdgeControlLibraryScript
	if err := s.Client.Post(ctx, "/edgecontrol/library", req, &created); err != nil {
		return nil, err
	}
	return &created, nil
}

// Update replaces the name, kind and code of one of the account's USER scripts.
func (s *EdgeControlLibraryService) Update(ctx context.Context, id string, req EdgeControlLibraryScriptRequest) (*EdgeControlLibraryScript, error) {
	endpoint, err := edgeControlLibraryEndpoint(id)
	if err != nil {
		return nil, err
	}

	var updated EdgeControlLibraryScript
	if err := s.Client.Put(ctx, endpoint, req, &updated); err != nil {
		return nil, err
	}
	return &updated, nil
}

// Delete deletes one of the account's USER scripts and returns it.
func (s *EdgeControlLibraryService) Delete(ctx context.Context, id string) (*EdgeControlLibraryScript, error) {
	endpoint, err := edgeControlLibraryEndpoint(id)
	if err != nil {
		return nil, err
	}

	var deleted EdgeControlLibraryScript
	if err := s.Client.Delete(ctx, endpoint, &deleted); err != nil {
		return nil, err
	}
	return &deleted, nil
}

func edgeControlLibraryEndpoint(id string) (string, error) {
	if id == "" {
		return "", fmt.Errorf("library script ID is required")
	}
	return fmt.Sprintf("/edgecontrol/library/%s", url.PathEscape(id)), nil
}
