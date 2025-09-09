package v2_6

import (
	"context"
	"fmt"
	"net/url"

	"github.com/cachefly/cachefly-sdk-go/internal/httpclient"
)

// SAMLService provides methods to manage SAML configurations.
type SAMLService struct {
	Client *httpclient.Client
}

// ActivateByID activates a SAML configuration by its id.
// PUT /saml/{id}/activate
func (s *SAMLService) ActivateByID(ctx context.Context, id string) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/saml/%s/activate", url.PathEscape(id))
	// No response body defined; send empty body and ignore response
	return s.Client.Put(ctx, endpoint, struct{}{}, nil)
}

// DeactivateByID deactivates a SAML configuration by its id.
// PUT /saml/{id}/deactivate
func (s *SAMLService) DeactivateByID(ctx context.Context, id string) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/saml/%s/deactivate", url.PathEscape(id))
	// No response body defined; send empty body and ignore response
	return s.Client.Put(ctx, endpoint, struct{}{}, nil)
}
