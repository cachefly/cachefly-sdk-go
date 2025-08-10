package v2_5

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/cachefly/cachefly-go-sdk/internal/httpclient"
)

// ScriptDefinitionsService handles /scriptConfigDefinitions endpoints.
type ScriptDefinitionsService struct {
	Client *httpclient.Client
}

// ScriptDefinition represents a script config definition resource.
//
// Reference: Script Definitions in API docs
// GET /scriptConfigDefinitions
// GET /scriptConfigDefinitions/{id}
type ScriptDefinition struct {
	ID               string                 `json:"_id"`
	Name             string                 `json:"name"`
	HelpText         string                 `json:"helpText"`
	DocsLink         string                 `json:"docsLink"`
	ScriptConfigs    []string               `json:"scriptConfigs"`
	Available        bool                   `json:"available"`
	LinkWithService  bool                   `json:"linkWithService"`
	RequiresOptions  bool                   `json:"requiresOptions"`
	RequiresRules    bool                   `json:"requiresRules"`
	RequiresPlugin   bool                   `json:"requiresPlugin"`
	CanCreate        bool                   `json:"canCreate"`
	Purpose          string                 `json:"purpose"`
	Meta             map[string]interface{} `json:"meta"`
	AllowedMimeTypes []string               `json:"allowedMimeTypes"`
	DefaultMimeType  string                 `json:"defaultMimeType"`
	ValueSchema      map[string]interface{} `json:"valueSchema"`
	DataMode         string                 `json:"dataMode"`
	DefaultValue     interface{}            `json:"defaultValue"`
}

type MetaInfoScriptDefinitions struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Count  int `json:"count"`
}

// ListScriptDefinitionsOptions holds filters & pagination for listing definitions.
type ListScriptDefinitionsOptions struct {
	IncludeFeatures bool
	IncludeHidden   bool
	Offset          int
	Limit           int
	ResponseType    string
}

// ListScriptDefinitionsResponse wraps a paged list of script definitions.
type ListScriptDefinitionsResponse struct {
	Meta        MetaInfoScriptDefinitions `json:"meta"`
	Definitions []ScriptDefinition        `json:"data"`
}

// List retrieves account-level script config definitions with optional filters.
// GET /scriptConfigDefinitions
func (s *ScriptDefinitionsService) List(ctx context.Context, opts ListScriptDefinitionsOptions) (*ListScriptDefinitionsResponse, error) {
	endpoint := "/scriptConfigDefinitions"
	params := url.Values{}

	// Only include boolean params if explicitly set; default is false
	if opts.IncludeFeatures {
		params.Set("includeFeatures", strconv.FormatBool(true))
	}
	if opts.IncludeHidden {
		params.Set("includeHidden", strconv.FormatBool(true))
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
	var resp ListScriptDefinitionsResponse
	if err := s.Client.Get(ctx, fullURL, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetByID retrieves a script config definition by its ID.
// GET /scriptConfigDefinitions/{id}
func (s *ScriptDefinitionsService) GetByID(ctx context.Context, id string) (*ScriptDefinition, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("/scriptConfigDefinitions/%s", url.PathEscape(id))
	var def ScriptDefinition
	if err := s.Client.Get(ctx, endpoint, &def); err != nil {
		return nil, err
	}
	return &def, nil
}
