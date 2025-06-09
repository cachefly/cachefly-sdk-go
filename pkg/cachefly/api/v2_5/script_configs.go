package v2_5

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/cachefly/cachefly-go-sdk/internal/httpclient"
)

// ScriptConfigsService handles /scriptConfigs endpoints.
type ScriptConfigsService struct {
	Client *httpclient.Client
}

// ScriptConfig represents a script config resource.
type ScriptConfig struct {
	ID                     string                 `json:"_id"`
	Name                   string                 `json:"name"`
	Services               []string               `json:"services"`
	ScriptConfigDefinition string                 `json:"scriptConfigDefinition"`
	Purpose                string                 `json:"purpose"`
	UseSchema              bool                   `json:"useSchema"`
	Meta                   map[string]interface{} `json:"meta"`
	MimeType               string                 `json:"mimeType"`
	DataMode               string                 `json:"dataMode"`
	Value                  interface{}            `json:"value"`
	CreatedAt              string                 `json:"createdAt"`
	UpdatedAt              string                 `json:"updateAt"`
}

// ListScriptConfigsOptions holds filters & pagination.
type ListScriptConfigsOptions struct {
	IncludeFeatures bool
	IncludeHidden   bool
	Status          string
	Offset          int
	Limit           int
	ResponseType    string
	Search          string
	SortBy          []string
}

// ListScriptConfigsResponse wraps a paged list.
type ListScriptConfigsResponse struct {
	Meta    MetaInfo       `json:"meta"`
	Configs []ScriptConfig `json:"data"`
}

// CreateScriptConfigRequest is the payload for creating a config.
type CreateScriptConfigRequest struct {
	Name                   string      `json:"name"`
	Services               []string    `json:"services"`
	ScriptConfigDefinition string      `json:"scriptConfigDefinition"`
	MimeType               string      `json:"mimeType"`
	Value                  interface{} `json:"value"`
}

// UpdateScriptConfigRequest is the payload for updating a config.
type UpdateScriptConfigRequest struct {
	Name                   string      `json:"name,omitempty"`
	MimeType               string      `json:"mimeType,omitempty"`
	Services               []string    `json:"services,omitempty"`
	ScriptConfigDefinition string      `json:"scriptConfigDefinition"`
	Value                  interface{} `json:"value,omitempty"`
}

// List returns script configs with optional filters.
func (s *ScriptConfigsService) List(ctx context.Context, opts ListScriptConfigsOptions) (*ListScriptConfigsResponse, error) {
	endpoint := "/scriptConfigs"
	params := url.Values{}
	params.Set("includeFeatures", strconv.FormatBool(opts.IncludeFeatures))
	params.Set("includeHidden", strconv.FormatBool(opts.IncludeHidden))
	if opts.Status != "" {
		params.Set("status", opts.Status)
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
	if opts.Search != "" {
		params.Set("search", opts.Search)
	}
	if len(opts.SortBy) > 0 {
		params.Set("sortBy", fmt.Sprintf("%v", opts.SortBy))
	}

	fullURL := fmt.Sprintf("%s?%s", endpoint, params.Encode())
	var resp ListScriptConfigsResponse
	if err := s.Client.Get(ctx, fullURL, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Create posts a new script config.
func (s *ScriptConfigsService) Create(ctx context.Context, req CreateScriptConfigRequest) (*ScriptConfig, error) {
	var created ScriptConfig
	if err := s.Client.Post(ctx, "/scriptConfigs", req, &created); err != nil {
		return nil, err
	}
	return &created, nil
}

// GetByID fetches a single config by ID.
func (s *ScriptConfigsService) GetByID(ctx context.Context, id, responseType string) (*ScriptConfig, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/scriptConfigs/%s", id)
	params := url.Values{}
	if responseType != "" {
		params.Set("responseType", responseType)
	}
	fullURL := fmt.Sprintf("%s?%s", endpoint, params.Encode())

	var cfg ScriptConfig
	if err := s.Client.Get(ctx, fullURL, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

// UpdateByID modifies an existing config.
func (s *ScriptConfigsService) UpdateByID(ctx context.Context, id string, req UpdateScriptConfigRequest) (*ScriptConfig, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/scriptConfigs/%s", id)

	var updated ScriptConfig
	if err := s.Client.Put(ctx, endpoint, req, &updated); err != nil {
		return nil, err
	}
	return &updated, nil
}

// GetSchemaByID retrieves the JSON schema for a config.
func (s *ScriptConfigsService) GetSchemaByID(ctx context.Context, id string) (map[string]interface{}, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/scriptConfigs/%s/schema", id)

	var schema map[string]interface{}
	if err := s.Client.Get(ctx, endpoint, &schema); err != nil {
		return nil, err
	}
	return schema, nil
}

// ActivateByID activates a script config.
func (s *ScriptConfigsService) ActivateByID(ctx context.Context, id string) (*ScriptConfig, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/scriptConfigs/%s/activate", id)

	var cfg ScriptConfig
	if err := s.Client.Put(ctx, endpoint, struct{}{}, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

// DeactivateByID deactivates a script config.
func (s *ScriptConfigsService) DeactivateByID(ctx context.Context, id string) (*ScriptConfig, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/scriptConfigs/%s/deactivate", id)

	var cfg ScriptConfig
	if err := s.Client.Put(ctx, endpoint, struct{}{}, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

// GetValueAsFile retrieves the raw script configuration file content for the given config ID.
// It calls GET /scriptConfigs/{id}/file and returns the file bytes.
func (s *ScriptConfigsService) GetValueAsFile(ctx context.Context, configID string) (*interface{}, error) {
	if configID == "" {
		return nil, fmt.Errorf("config ID is required")
	}
	endpoint := fmt.Sprintf("/scriptConfigs/%s/file", url.PathEscape(configID))

	// The client should return raw bytes for non-JSON endpoints.
	var content interface{}
	if err := s.Client.Get(ctx, endpoint, &content); err != nil {
		return nil, err
	}
	return &content, nil
}

// UpdateScriptConfigValue updates the script configuration content using raw file data.
func (s *ScriptConfigsService) UpdateValueAsFile(ctx context.Context, configID string, content []byte) (*ScriptConfig, error) {
	if configID == "" {
		return nil, fmt.Errorf("config ID is required")
	}
	endpoint := fmt.Sprintf("/scriptConfigs/%s/value", url.PathEscape(configID))

	var updated ScriptConfig
	// Pass raw bytes as body; Client.Put must handle []byte by sending as-is
	if err := s.Client.Put(ctx, endpoint, content, &updated); err != nil {
		return nil, err
	}
	return &updated, nil
}

// ListPromo retrieves promo script config definitions.
// GET /scriptConfigDefinitions/promo
func (s *ScriptConfigsService) ListPromo(ctx context.Context, includeFeatures bool) ([]ScriptConfig, error) {
	endpoint := "/scriptConfigDefinitions/promo"
	params := url.Values{}
	// only send includeFeatures when true
	if includeFeatures {
		params.Set("includeFeatures", strconv.FormatBool(true))
	}
	fullURL := fmt.Sprintf("%s?%s", endpoint, params.Encode())

	var defs []ScriptConfig
	if err := s.Client.Get(ctx, fullURL, &defs); err != nil {
		return nil, err
	}
	return defs, nil
}

// GetDefinitionByID retrieves definition script config.
func (s *ScriptConfigsService) GetDefinitionByID(ctx context.Context, id string) (*ScriptConfig, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/scriptConfigDefinitions/%s", url.PathEscape(id))

	var def ScriptConfig
	if err := s.Client.Get(ctx, endpoint, &def); err != nil {
		return nil, err
	}
	return &def, nil
}

// List returns account-level script config definitions with optional filters.
// GET /scriptConfigDefinitions
func (s *ScriptConfigsService) ListAccountScriptConfigDefinitions(ctx context.Context, opts ListScriptConfigsOptions) (*ListScriptConfigsResponse, error) {
	endpoint := "/scriptConfigDefinitions"
	params := url.Values{}
	params.Set("includeFeatures", strconv.FormatBool(opts.IncludeFeatures))
	params.Set("includeHidden", strconv.FormatBool(opts.IncludeHidden))
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

	var resp ListScriptConfigsResponse
	if err := s.Client.Get(ctx, fullURL, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
