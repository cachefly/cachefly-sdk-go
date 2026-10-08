package v2_6

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/cachefly/cachefly-sdk-go/internal/httpclient"
)

// EdgeControlScriptsService handles edge control script operations
// (/edgecontrol/services/{sid}/{kind}).
//
// Each service has one editable draft per kind. Publishing creates an
// immutable, numbered version; at most one version per kind is active.
// Versions cannot be deleted.
type EdgeControlScriptsService struct {
	Client *httpclient.Client
}

// EdgeControlScriptKind is the kind of an edge control script.
type EdgeControlScriptKind string

const (
	// EdgeControlScriptKindRequest scripts run when a request is received.
	EdgeControlScriptKindRequest EdgeControlScriptKind = "REQUEST"
	// EdgeControlScriptKindResponse scripts run before the response is returned.
	EdgeControlScriptKindResponse EdgeControlScriptKind = "RESPONSE"
)

// Edge control script statuses.
const (
	EdgeControlScriptStatusActive      = "ACTIVE"
	EdgeControlScriptStatusDeactivated = "DEACTIVATED"
	EdgeControlScriptStatusDraft       = "DRAFT"
)

// EdgeControlScript is an edge control script draft or version.
type EdgeControlScript struct {
	ID   string `json:"_id"`
	Kind string `json:"kind"`
	// Version is 0 for drafts.
	Version    int    `json:"version"`
	Script     string `json:"script"`
	ScriptSize int    `json:"scriptSize"`
	Status     string `json:"status"`
	// Filename is empty for drafts.
	Filename string `json:"filename"`
	// Dirty is only returned for drafts: true when the draft has unpublished changes.
	Dirty           bool   `json:"dirty"`
	LastActivatedAt string `json:"lastActivatedAt"`
	CreatedAt       string `json:"createdAt"`
	UpdatedAt       string `json:"updatedAt"`
}

type edgeControlScriptBody struct {
	Script string `json:"script"`
}

// GetDraft returns the draft for the given kind. The API creates a draft with a
// default script if none exists yet.
func (s *EdgeControlScriptsService) GetDraft(ctx context.Context, serviceID string, kind EdgeControlScriptKind) (*EdgeControlScript, error) {
	endpoint, err := edgeControlScriptEndpoint(serviceID, kind, "")
	if err != nil {
		return nil, err
	}

	var draft EdgeControlScript
	if err := s.Client.Get(ctx, endpoint, &draft); err != nil {
		return nil, err
	}
	return &draft, nil
}

// SaveDraft replaces the draft for the given kind without affecting traffic.
func (s *EdgeControlScriptsService) SaveDraft(ctx context.Context, serviceID string, kind EdgeControlScriptKind, script string) (*EdgeControlScript, error) {
	if script == "" {
		return nil, fmt.Errorf("script is required")
	}
	endpoint, err := edgeControlScriptEndpoint(serviceID, kind, "")
	if err != nil {
		return nil, err
	}

	var draft EdgeControlScript
	if err := s.Client.Put(ctx, endpoint, edgeControlScriptBody{Script: script}, &draft); err != nil {
		return nil, err
	}
	return &draft, nil
}

// ListVersions returns all published versions for the given kind, newest first.
func (s *EdgeControlScriptsService) ListVersions(ctx context.Context, serviceID string, kind EdgeControlScriptKind) ([]EdgeControlScript, error) {
	endpoint, err := edgeControlScriptEndpoint(serviceID, kind, "/versions")
	if err != nil {
		return nil, err
	}

	var versions []EdgeControlScript
	if err := s.Client.Get(ctx, endpoint, &versions); err != nil {
		return nil, err
	}
	return versions, nil
}

// GetActiveVersion returns the active version for the given kind. The API
// responds with 404 when no version is active.
func (s *EdgeControlScriptsService) GetActiveVersion(ctx context.Context, serviceID string, kind EdgeControlScriptKind) (*EdgeControlScript, error) {
	endpoint, err := edgeControlScriptEndpoint(serviceID, kind, "/versions/active")
	if err != nil {
		return nil, err
	}

	var active EdgeControlScript
	if err := s.Client.Get(ctx, endpoint, &active); err != nil {
		return nil, err
	}
	return &active, nil
}

// GetVersion returns a specific version for the given kind.
func (s *EdgeControlScriptsService) GetVersion(ctx context.Context, serviceID string, kind EdgeControlScriptKind, version int) (*EdgeControlScript, error) {
	if version < 1 {
		return nil, fmt.Errorf("version must be 1 or greater")
	}
	endpoint, err := edgeControlScriptEndpoint(serviceID, kind, fmt.Sprintf("/versions/%d", version))
	if err != nil {
		return nil, err
	}

	var result EdgeControlScript
	if err := s.Client.Get(ctx, endpoint, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateVersion publishes the given script as a new, deactivated version. The
// script must define a "handler" function.
func (s *EdgeControlScriptsService) CreateVersion(ctx context.Context, serviceID string, kind EdgeControlScriptKind, script string) (*EdgeControlScript, error) {
	if script == "" {
		return nil, fmt.Errorf("script is required")
	}
	endpoint, err := edgeControlScriptEndpoint(serviceID, kind, "/versions")
	if err != nil {
		return nil, err
	}

	var created EdgeControlScript
	if err := s.Client.Post(ctx, endpoint, edgeControlScriptBody{Script: script}, &created); err != nil {
		return nil, err
	}
	return &created, nil
}

// CreateVersionFromDraft publishes the current draft as a new, deactivated
// version. When resetDraft is true the draft is reset to the default script
// afterwards.
func (s *EdgeControlScriptsService) CreateVersionFromDraft(ctx context.Context, serviceID string, kind EdgeControlScriptKind, resetDraft bool) (*EdgeControlScript, error) {
	endpoint, err := edgeControlScriptEndpoint(serviceID, kind, "/versions")
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Set("fromDraft", "true")
	if resetDraft {
		params.Set("resetDraft", "true")
	}

	// The API rejects a script body together with fromDraft, and its JSON body
	// parser rejects a bare null, so send an empty object.
	var created EdgeControlScript
	if err := s.Client.Post(ctx, endpoint+"?"+params.Encode(), struct{}{}, &created); err != nil {
		return nil, err
	}
	return &created, nil
}

// ActivateVersion makes the given version the only active version for the kind.
func (s *EdgeControlScriptsService) ActivateVersion(ctx context.Context, serviceID string, kind EdgeControlScriptKind, version int) (*EdgeControlScript, error) {
	if version < 1 {
		return nil, fmt.Errorf("version must be 1 or greater")
	}
	endpoint, err := edgeControlScriptEndpoint(serviceID, kind, fmt.Sprintf("/versions/%d/activate", version))
	if err != nil {
		return nil, err
	}

	var activated EdgeControlScript
	if err := s.Client.Put(ctx, endpoint, struct{}{}, &activated); err != nil {
		return nil, err
	}
	return &activated, nil
}

// ActivateLastVersion re-activates the most recently activated version for the
// kind. The API responds with 400 when no version was ever activated.
func (s *EdgeControlScriptsService) ActivateLastVersion(ctx context.Context, serviceID string, kind EdgeControlScriptKind) (*EdgeControlScript, error) {
	endpoint, err := edgeControlScriptEndpoint(serviceID, kind, "/activate")
	if err != nil {
		return nil, err
	}

	var activated EdgeControlScript
	if err := s.Client.Put(ctx, endpoint, struct{}{}, &activated); err != nil {
		return nil, err
	}
	return &activated, nil
}

// Deactivate deactivates every version for the kind, leaving no script running.
func (s *EdgeControlScriptsService) Deactivate(ctx context.Context, serviceID string, kind EdgeControlScriptKind) error {
	endpoint, err := edgeControlScriptEndpoint(serviceID, kind, "/deactivate")
	if err != nil {
		return err
	}

	return s.Client.Put(ctx, endpoint, struct{}{}, nil)
}

// edgeControlScriptEndpoint builds /edgecontrol/services/{sid}/{kind}{suffix}.
// The kind is accepted in either case; the API expects it lowercase in the path.
func edgeControlScriptEndpoint(serviceID string, kind EdgeControlScriptKind, suffix string) (string, error) {
	if serviceID == "" {
		return "", fmt.Errorf("service ID is required")
	}
	normalized := EdgeControlScriptKind(strings.ToUpper(string(kind)))
	if normalized != EdgeControlScriptKindRequest && normalized != EdgeControlScriptKindResponse {
		return "", fmt.Errorf("invalid edge control script kind %q: must be %s or %s", kind, EdgeControlScriptKindRequest, EdgeControlScriptKindResponse)
	}
	return fmt.Sprintf("/edgecontrol/services/%s/%s%s", url.PathEscape(serviceID), strings.ToLower(string(normalized)), suffix), nil
}
